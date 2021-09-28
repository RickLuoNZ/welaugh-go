package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	dbUser         string
	dbPswd         string
	dbHost         string
	dbPort         string
	dbName         string
	apiServerPort  string
}

func New() *Config {
	conf := &Config{}

	flag.StringVar(&conf.dbUser, "dbuser", os.Getenv("DB_USER"), "DB user name")
	flag.StringVar(&conf.dbPswd, "dbpswd", os.Getenv("DB_PSWD"), "DB pass")
	flag.StringVar(&conf.dbPort, "dbport", os.Getenv("DB_PORT"), "DB port")
	flag.StringVar(&conf.dbHost, "dbhost", os.Getenv("DB_HOST"), "DB host")
	flag.StringVar(&conf.dbName, "dbname", os.Getenv("DB_NAME"), "DB name")
	flag.StringVar(&conf.apiServerPort, "apiServerPort", os.Getenv("API_SERVER_PORT"), "API Server Port")

	flag.Parse()

	return conf
}

func (c *Config) GetDBConnStr() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.dbUser,
		c.dbPswd,
		c.dbHost,
		c.dbPort,
		c.dbName,
	)
}

func (c *Config) GetApiServerPort() string {
	return ":" + c.apiServerPort
}
