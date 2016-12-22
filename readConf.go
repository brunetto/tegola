package tegola

import (
	"log"
	"github.com/brunetto/goutils/conf"
)

func LoadBot (fileName string) Bot {
	var (
		err error
		c = Bot{}
	)
	err = conf.LoadJsonConf(fileName, &c)
	if err != nil {
		log.Fatal("Error reading JSON config file: ", err)
	}
	return c
}
