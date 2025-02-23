package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func flattenMap(prefix string, settings map[string]interface{}) map[string]string {
	flat := make(map[string]string)
	for key, value := range settings {
		fullKey := strings.ToUpper(key)
		fmt.Println("fullKey----", fullKey)
		if prefix != "" {
			fullKey = strings.ToUpper(prefix) + "_" + strings.ToUpper(fullKey)

		}
		switch v := value.(type) {
		case map[string]interface{}:
			// recurse for nested maps
			nested := flattenMap(fullKey, v)
			for nestedKey, nestedValue := range nested {
				flat[nestedKey] = nestedValue
			}

		default:
			flat[fullKey] = fmt.Sprintf("%v", value)
		}
	}
	return flat

}

func main() {
	// Load the YAML file
	viper.SetConfigName("local") // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") //look for config in the current directory

	//read the yarml configuration
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return
	}
	// Flatten the configuration
	allSettings := viper.AllSettings()
	log.Println("allSeetingss------", allSettings)
	flatConfig := flattenMap("", allSettings)

	//write to .env file
	envFile, err := os.Create("./environment/.env")
	if err != nil {
		fmt.Printf("Error creating .env file: %v\n", err)
		return
	}

	defer envFile.Close()
	for key, value := range flatConfig {
		//fmt.Fprintf(w io.Writer, format string, args ...interface{}) (n int, err error)
		_, err := fmt.Fprintf(envFile, "%s=%s\n", key, value)
		if err != nil {
			fmt.Printf("Error writing to .env file: %v\n", err)
			return
		}

	}
	fmt.Println(".env file created successfully")
}
