import type { AxiosError, AxiosInstance, AxiosResponse } from 'axios'
import axios from 'axios'

import type { AdmiralError } from './errors.ts'
import { grpcResponseToError } from './errors.ts'

/**
 * HTTP response status.
 *
 * Responses are grouped in five classes:
 *  - Informational responses (100–199)
 *  - Successful responses (200–299)
 *  - Redirects (300–399)
 *  - Client errors (400–499)
 *  - Server errors (500–599)
 */
export interface HttpStatus {
  code: number
  text: string
}

const successInterceptor = (response: AxiosResponse): AxiosResponse => {
  return response
}

const errorInterceptor = async (error: AxiosError): Promise<AdmiralError> => {
  /* eslint-disable  @typescript-eslint/no-explicit-any */
  const response: AxiosResponse<any, any> = error?.response as AxiosResponse
  if (response === undefined) {
    const clientError: AdmiralError = {
      status: {
        code: 500,
        text: 'Client Error',
      },
      message: error.message,
      name: 'Client Error',
    }
    return await Promise.reject(clientError)
  }

  // This section handles authentication redirects.
  if (response?.status === 401) {
    const redirectUrl: string = window.location.pathname + window.location.search
    window.location.href = `/auth/login?redirect_url=${encodeURIComponent(redirectUrl)}`
  }

  // we are guaranteed to have a response object on the error from this point on
  // since we have already accounted for axios errors.
  const responseData = response?.data

  // if the response data has a code on it we know it's a gRPC response.
  let err
  if (responseData?.code !== undefined) {
    err = grpcResponseToError(error)
  } else {
    const message =
      typeof error.response?.data === 'string'
        ? error.response.data
        : // eslint-disable-next-line no-unsafe-optional-chaining
          (error?.message).length > 0 || error.response?.statusText

    // eslint-disable-next-line @typescript-eslint/consistent-type-assertions
    err = {
      // eslint-disable-next-line @typescript-eslint/consistent-type-assertions
      status: {
        code: error.response?.status,
        text: error.response?.statusText,
      } as HttpStatus,
      message,
      data: responseData,
    } as AdmiralError
  }
  return await Promise.reject(err)
}

const createClient = (): AxiosInstance => {
  const axiosClient: AxiosInstance = axios.create({
    // n.b. the client will treat any response code >= 400 as an error and apply the error interceptor.
    validateStatus: (status: number): boolean => {
      return status < 400
    },
  })
  axiosClient.interceptors.response.use(successInterceptor, errorInterceptor)

  return axiosClient
}

const client: AxiosInstance = createClient()

export { client as default, errorInterceptor, successInterceptor }
