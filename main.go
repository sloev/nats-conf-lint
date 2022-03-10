package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/nats-io/nats-server/v2/conf"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error, missing filename arg for nats.conf")
		os.Exit(1)
	}

	natsConfFile := os.Args[len(os.Args)-1]
	data, err := ioutil.ReadFile(natsConfFile)

	if err != nil {
		panic(fmt.Errorf("error opening config file: %v", err))
	}
	r := regexp.MustCompile(`\$[A-Za-z0-9_]+[^\s]`)

	for _, variableName := range r.FindAllString(string(data), -1) {
		fmt.Printf("setting variable '%s' to 'test'\n", variableName)
		err := os.Setenv(variableName[1:], "test")
		if err != nil {
			panic(err)
		}
	}

	_, err = conf.Parse(string(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("lint ok")
}
