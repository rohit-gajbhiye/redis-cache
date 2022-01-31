package main

import (
	"fmt"
	"log"

	"com.hook.cache/config"
)

func main() {
	// Register default env properties
	appConfig := config.NewAppConfig()
	err := config.LoadEnvFile(appConfig)
	if err != nil {
		log.Fatal("Error loading config file", err)
	}
	// configure redis datastore
	//redisStore := client.NewRedisStoreImpl(appConfig)
	fmt.Print("Just Ran .. ")
}
