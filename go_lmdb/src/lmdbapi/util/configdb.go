package util

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	MovieDbFile     = "moviedata.json"
	MovieFilterFile = "moviefilter.json"
)

var (
	MovieDataViper   = viper.New()
	MovieFilterViper = viper.New()
)

func InitViper() {
	MovieDataViper.SetConfigFile(MovieDbFile)
	if err := MovieDataViper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	MovieFilterViper.SetConfigFile(MovieFilterFile)
	if err := MovieFilterViper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
