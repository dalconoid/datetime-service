package main

import (
	"fmt"
	"github.com/dalconoid/datetime-service/server"
	"github.com/jessevdk/go-flags"
	"log"
)

type envVars struct {
	Port string `short:"p" long:"port" env:"PORT" description:"microservice port" default:"8080"`
}


// getEnvVars - parse environment or command line arguments
func getEnvVars() (*envVars, error) {
	env := &envVars{}
	if _, err := flags.Parse(env); err != nil {
		return nil, err
	}
	return env, nil
}

func main() {
	env, err := getEnvVars()
	if err != nil {
		log.Fatal(err)
	}

	s := server.New()
	s.ConfigureRouter()
	log.Fatal(s.Start(fmt.Sprintf(":%s", env.Port)))
}
