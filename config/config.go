package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var Config Configuration

type Configuration struct {
	Mongo MongoConf `mapstructure:"mongo"`
}

type MongoConf struct {
	User    string        `mapstructure:"user"`
	Pass    string        `mapstructure:"password"`
	Host    string        `mapstructure:"host"`
	Port    int           `mapstructure:"port"`
	DbName  string        `mapstructure:"dbName"`
	TimeOut time.Duration `mapstructure:"timeOut"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/Users/jcpalma/go/src/admin-server")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("No se encontró el archivo de configuración\n", err)
		} else {
			log.Fatal("Error al cargar archivo de configuración\n", err)
		}
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatal("No se pudo decodificar el archivo de configuración\n", err)
	}

}

func Get(key string) interface{} {
	return viper.Get(key)
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
}
