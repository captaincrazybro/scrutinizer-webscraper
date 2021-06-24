package sw

import (
	"os"
	"log"
)

// GetEnvVariables gets the environment variables required 
// for sending data to ga and email
func GetEnvVariables(env []string) {
	var value string

	for i, _ := range env {
		value = os.Getenv(env[i])

		if value == "" {
			log.Fatalln("error: please set environment variable " + env[i])
		} else {
			env[i] = value
		}
	}
}
