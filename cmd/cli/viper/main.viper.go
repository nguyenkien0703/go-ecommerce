package main

//
//import (
//	"fmt"
//
//	"github.com/spf13/viper"
//)
//
//type Config struct {
//	Server struct {
//		Port int `mapstructure:"port"`
//	} `mapstructure:"server"`
//
//	Databases []struct {
//		User     string `mapstructure:"user"`
//		Password string `mapstructure:"password"`
//		Host     string `mapstructure:"host"`
//	} `mapstructure:"databases"`
//}
//
//func main() {
//	viper := viper.New()
//	viper.AddConfigPath(("./config/")) // path to config
//	viper.SetConfigName("local")       //ten file
//
//	viper.SetConfigType("yaml") // loai file
//
//	// read configuration
//	err := viper.ReadInConfig()
//	if err != nil {
//		panic(fmt.Errorf("Fatal error config file: %s \n", err))
//	}
//
//	// read server configuration
//	// fmt.Println("Server port is: ", viper.GetInt("server.port"))
//	// fmt.Println("jwt key is: ", viper.GetString("security.jwt.key"))
//
//	//configure structure
//	var config Config
//	if err := viper.Unmarshal(&config); err != nil {
//		fmt.Printf("Unable to decode into struct, %v", err)
//	}
//
//	fmt.Println("Server port is: ", config.Server.Port)
//	fmt.Println("Database host is: ", config.Databases)
//	for _, db := range config.Databases {
//		fmt.Printf("datbaseUser: %s, password: %s, host: %s \n", db.User, db.Password, db.Host)
//
//	}
//
//}
