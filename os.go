package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	//viper.SetConfigFile("config.yml")      // name of config file (without extension)
	//	//	//viper.SetConfigType("yml")        // config file extension
	//	//	//viper.AddConfigPath("$GOPATH/github.com/tppgit/we_service") // optionally look for config in the working directory
	//	//	//err := viper.ReadInConfig()        // Find and read the config file
	//	//	//if err != nil {                    // Handle errors reading the config file
	//	//	//	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	//	//	//}
	//	//	//
	//	//	////fmt.Printf("%s lives in %s.\n", os.Getenv("USER"), os.Getenv("HOME"))
	//	//	//fmt.Printf("%s", viper.GetString("database.host"))
	viper.BindEnv("name", "123")
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	fmt.Println("asd")
	fmt.Println(viper.GetInt("Foo"))
}