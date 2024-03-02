package cmd

import (
	"context"
	"crypto/rand"
	_ "embed"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/cli/browser"
	"github.com/coreos/go-oidc/v3/oidc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"

	"go.datalift.io/admiral/common/client"
)

var (
	//go:embed login.gohtml
	resultPage []byte
	ports      = []int{1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368}
)

type LoginCmd struct {
	Cmd *cobra.Command
}

func NewLoginCmd(clientOpts *client.Options) *LoginCmd {
	root := &LoginCmd{}

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Log in to Admiral and retrieve credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			// set up a listener on the redirect port
			var host string
			var ln net.Listener
			for _, port := range ports {
				var err error

				host = fmt.Sprintf("localhost:%d", port)
				ln, err = net.Listen("tcp", host)
				if err == nil {
					break
				}
			}

			if ln == nil {
				return errors.New("unable to find an available port to create callback listener")
			}

			// create an admiral client and retrieve oauth2 config
			admiral, err := client.NewClient(clientOpts)
			if err != nil {
				return err
			}
			httpClient, err := admiral.HTTPClient()
			if err != nil {
				return err
			}
			oauth2conf, _, err := admiral.OIDCConfig(oidc.ClientContext(ctx, httpClient))
			if err != nil {
				return err
			}

			// create redirect URL
			oauth2conf.RedirectURL = fmt.Sprintf("http://%s/", host)

			// generate authorization code URL
			state, err := state(24)
			verifier := oauth2.GenerateVerifier()
			authCodeURL := oauth2conf.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier), oauth2.AccessTypeOffline)

			// start a web server to listen on a callback URL
			server := &http.Server{Addr: host}
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/" {
					http.Error(w, "file not found", http.StatusNotFound)
					return
				}

				// get the authorization code
				code := r.URL.Query().Get("code")
				if code == "" {
					msg := "the provided authorization response is invalid. expect a url with query parameters of code."
					http.Error(w, msg, http.StatusBadRequest)
					log.Errorf(msg)

					cleanup(server)
					return
				}

				s := r.URL.Query().Get("state")
				if state != s {
					msg := "the provided authorization response is invalid. expect a url with query parameters of state to match provided."
					http.Error(w, msg, http.StatusBadRequest)
					log.Errorf(msg)

					cleanup(server)
					return
				}

				// exchange the authorization code for oauth2 token
				token, err := oauth2conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
				if err != nil {
					msg := "failure to exchange authorization code for oauth tokens."
					http.Error(w, msg, http.StatusInternalServerError)
					log.Errorf(msg)

					cleanup(server)
					return
				}

				// save the token to the config file
				config := admiral.ClientConfig()
				config.Token = *token

				err = config.Save()
				if err != nil {
					msg := "failure to save oauth token to configuration file."
					http.Error(w, msg, http.StatusInternalServerError)
					log.Errorf(msg)

					cleanup(server)
					return
				}

				// return an indication of success to the caller
				_, err = io.WriteString(w, string(resultPage))
				if err != nil {
					log.Error(err.Error())
					return
				}

				log.Info("successfully logged into admiral")
				cleanup(server)
			})

			log.Infof("browser has been opened to visit: %v", authCodeURL)

			// open authorization URL in a browser window
			err = browser.OpenURL(authCodeURL)
			if err != nil {
				return fmt.Errorf("unable to open browser to URL %s: %s", authCodeURL, err)
			}

			_ = server.Serve(ln)

			return nil
		},
	}

	root.Cmd = cmd
	return root
}

func state(n int) (string, error) {
	data := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func cleanup(server *http.Server) {
	go func() {
		_ = server.Close()
	}()
}
