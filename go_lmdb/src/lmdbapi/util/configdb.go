package util

import (
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

func init() {
	MovieDataViper.SetConfigName(MovieDbFile)
	MovieDataViper.AddConfigPath(".")

	MovieFilterViper.SetConfigName(MovieFilterFile)
	MovieFilterViper.AddConfigPath(".")
}
