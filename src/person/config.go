package person

import (
	"os"
	"strconv"
)

type Config struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Host string `json:"host"`
	Port int    `json:"port"`
	DB   string `json:"db"`
}

var config Config

func init() {

	if os.Getenv("MYSQLUSER") == "" {
		panic("No \"MYSQLUSER\" environment variable has been configured")
	}
	config.User = os.Getenv("MYSQLUSER")
	config.Pass = os.Getenv("MYSQLPASS")

	if os.Getenv("MYSQLHOST") == "" {
		panic("No \"MYSQLHOST\" environment variable has been configured")
	}
	config.Host = os.Getenv("MYSQLHOST")

	if os.Getenv("MYSQLPORT") == "" {
		panic("No \"MYSQLPORT\" environment variable has been configured")
	}
	var err error
	config.Port, err = strconv.Atoi(os.Getenv("MYSQLPORT"))
	if err != nil {
		panic("Environment variable \"MYSQLPORT\" needs to be an integer")
	}

	if os.Getenv("MYSQLDB") == "" {
		panic("No \"MYSQLDB\" environment variable has been configured")
	}
	config.DB = os.Getenv("MYSQLDB")
}
