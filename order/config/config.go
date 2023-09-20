package config

import "github.com/jinzhu/configor"

type Config struct {
	AppConfig          AppConfig
	DBConfig           DBConfig
	JWTConfig          JWTConfig
	InventoryAPIConfig InventoryAPIConfig
}

type AppConfig struct {
	APPName string `default:"order-api"`
	Port    int    `env:"PORT" default:"3000"`
	Version string `default:"x.x.x"`
}

type DBConfig struct {
	Host     string `default:"localhost" env:"DBHOST"`
	DataBase string `default:"order" env:"DBNAME"`
	User     string `default:"order" env:"DBUSERNAME"`
	Password string `required:"true" env:"DBPASSWORD" default:"mysecretpassword"`
	Port     uint   `default:"3306" env:"DBPORT"`
}

type JWTConfig struct {
	PrivateKey string `default:"mysecretkey" env:"JWTPRIVATEKEY"`
	PublicKey  string `default:"publickey" env:"JWTPUBLICKEY"`
	Expire     uint64 `default:"60" env:"JWTEXPIRE"`
	Issuer     string `default:"user" env:"JWTISSUER"`
	Audience   string `default:"user" env:"JWTAUDIENCE"`
}

type InventoryAPIConfig struct {
	Host string `default:"apollo-router" env:"INVENTORYAPIHOST"`
	Port int    `default:"4000" env:"INVENTORYAPIPORT"`
}

func LoadConfigOrPanic() Config {
	var config = Config{}
	configor.Load(&config, "config/config.dev.json")

	return config
}
