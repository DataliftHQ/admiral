package env

import (
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func ParseNumFromEnv(env string, defaultValue, min, max int) int {
	str := os.Getenv(env)
	if str == "" {
		return defaultValue
	}
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		log.Warnf("Could not parse '%s' as a number from environment %s", str, env)
		return defaultValue
	}
	if num > math.MaxInt || num < math.MinInt {
		log.Warnf("Value in %s is %d is outside of the min and max %d allowed values. Using default %d", env, num, min, defaultValue)
		return defaultValue
	}
	if int(num) < min {
		log.Warnf("Value in %s is %d, which is less than minimum %d allowed", env, num, min)
		return defaultValue
	}
	if int(num) > max {
		log.Warnf("Value in %s is %d, which is greater than maximum %d allowed", env, num, max)
		return defaultValue
	}
	return int(num)
}

func ParseInt64FromEnv(env string, defaultValue, min, max int64) int64 {
	str := os.Getenv(env)
	if str == "" {
		return defaultValue
	}

	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Warnf("Could not parse '%s' as a int64 from environment %s", str, env)
		return defaultValue
	}
	if num < min {
		log.Warnf("Value in %s is %d, which is less than minimum %d allowed", env, num, min)
		return defaultValue
	}
	if num > max {
		log.Warnf("Value in %s is %d, which is greater than maximum %d allowed", env, num, max)
		return defaultValue
	}
	return num
}

func ParseFloatFromEnv(env string, defaultValue, min, max float32) float32 {
	str := os.Getenv(env)
	if str == "" {
		return defaultValue
	}

	num, err := strconv.ParseFloat(str, 32)
	if err != nil {
		log.Warnf("Could not parse '%s' as a float32 from environment %s", str, env)
		return defaultValue
	}
	if float32(num) < min {
		log.Warnf("Value in %s is %f, which is less than minimum %f allowed", env, num, min)
		return defaultValue
	}
	if float32(num) > max {
		log.Warnf("Value in %s is %f, which is greater than maximum %f allowed", env, num, max)
		return defaultValue
	}
	return float32(num)
}

func ParseFloat64FromEnv(env string, defaultValue, min, max float64) float64 {
	str := os.Getenv(env)
	if str == "" {
		return defaultValue
	}

	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Warnf("Could not parse '%s' as a float32 from environment %s", str, env)
		return defaultValue
	}
	if num < min {
		log.Warnf("Value in %s is %f, which is less than minimum %f allowed", env, num, min)
		return defaultValue
	}
	if num > max {
		log.Warnf("Value in %s is %f, which is greater than maximum %f allowed", env, num, max)
		return defaultValue
	}
	return num
}

func ParseDurationFromEnv(env string, defaultValue, min, max time.Duration) time.Duration {
	str := os.Getenv(env)
	if str == "" {
		return defaultValue
	}
	dur, err := time.ParseDuration(str)
	if err != nil {
		log.Warnf("Could not parse '%s' as a duration string from environment %s", str, env)
		return defaultValue
	}

	if dur < min {
		log.Warnf("Value in %s is %s, which is less than minimum %s allowed", env, dur, min)
		return defaultValue
	}
	if dur > max {
		log.Warnf("Value in %s is %s, which is greater than maximum %s allowed", env, dur, max)
		return defaultValue
	}
	return dur
}

func StringFromEnv(env string, defaultValue string) string {
	if str := os.Getenv(env); str != "" {
		return str
	}
	return defaultValue
}

func StringsFromEnv(env string, defaultValue []string, separator string) []string {
	if str := os.Getenv(env); str != "" {
		ss := strings.Split(str, separator)
		for i, s := range ss {
			ss[i] = strings.TrimSpace(s)
		}
		return ss
	}
	return defaultValue
}

func ParseBoolFromEnv(envVar string, defaultValue bool) bool {
	if val := os.Getenv(envVar); val != "" {
		if strings.ToLower(val) == "true" {
			return true
		} else if strings.ToLower(val) == "false" {
			return false
		}
	}
	return defaultValue
}

func ParseStringToStringFromEnv(envVar string, defaultValue map[string]string, seperator string) map[string]string {
	str := os.Getenv(envVar)
	str = strings.TrimSpace(str)
	if str == "" {
		return defaultValue
	}

	parsed := make(map[string]string)
	for _, pair := range strings.Split(str, seperator) {
		keyvalue := strings.Split(pair, "=")
		if len(keyvalue) != 2 {
			log.Warnf("Invalid key-value pair when parsing environment '%s' as a string map", str)
			return defaultValue
		}
		key := strings.TrimSpace(keyvalue[0])
		value := strings.TrimSpace(keyvalue[1])
		if _, ok := parsed[key]; ok {
			log.Warnf("Duplicate key '%s' when parsing environment '%s' as a string map", key, str)
			return defaultValue
		}
		parsed[key] = value
	}
	return parsed
}
