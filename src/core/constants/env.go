package constants

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
)

var Env env

type env struct {
	AppPort       string
	RedisUsername string
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisTTL      string
	RedisDB       string
	Salt          string
	BaseURL       string
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	loadEnv()
}

func loadEnv() {
	Env = env{}
	var unsetVars []string

	envmap := map[string]*string{
		"APP_PORT":       &Env.AppPort,
		"REDIS_USER":     &Env.RedisUsername,
		"REDIS_HOST":     &Env.RedisHost,
		"REDIS_PORT":     &Env.RedisPort,
		"REDIS_DB":       &Env.RedisDB,
		"REDIS_TTL":      &Env.RedisTTL,
		"REDIS_PASSWORD": &Env.RedisPassword,
		"SALT":           &Env.Salt,
		"BASE_URL":       &Env.BaseURL,
	}

	for key, value := range envmap {
		val := os.Getenv(key)
		*value = val
		if len(*value) == 0 || *value == "" {
			unsetVars = append(unsetVars, key)
			continue
		}

	}

	if len(unsetVars) != 0 {
		slog.Error(fmt.Sprintf("Required envs : %v", unsetVars))
	}
}

func Int(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		slog.Error(err.Error())
		return 0
	}
	return val
}
