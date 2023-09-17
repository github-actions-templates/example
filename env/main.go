package main

import "os"

func main() {
	if env, ok := os.LookupEnv("ENVIRONMENT"); !ok {
		panic("ENVIRONMENT is not set")
	} else {
		println(env)
	}
}
