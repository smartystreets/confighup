// After starting this command, update a value in config.json and
// send a SIGHUP signal to the running process. You will see (based
// on the recurring log messages) that the updates values from the
// config file were brought into the process.
package main

import (
	"log"
	"os"
	"path"
	"time"
)

func main() {
	configFile, _ := os.Getwd()
	configFile = path.Join(configFile, "config.json")
	manager := NewConfigManager(configFile)

	go pollingExample(manager)
	channelExample(manager)
}

func pollingExample(manager *ConfigManager) {
	for {
		log.Println("POLLING EXAMPLE:", manager.Config())
		time.Sleep(time.Second * 1)
	}
}

func channelExample(manager *ConfigManager) {
	for {
		log.Println("CHANNEL EXAMPLE:", manager.Notifications())
	}
}
