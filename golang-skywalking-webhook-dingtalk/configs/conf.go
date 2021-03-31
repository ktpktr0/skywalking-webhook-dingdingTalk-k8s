package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var V *viper.Viper

const (
	productionConfName = "conf.yml"
)


func InitConf() {

	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("failed to get app path,err:", err)
		os.Exit(1)
	}

	fileName := filepath.Join(path, "configs", productionConfName)

	tempV := viper.New()
	tempV.SetConfigFile(fileName)
	err = tempV.ReadInConfig()
	if err != nil {
		fmt.Println("failed to init config,err:", err)
		os.Exit(1)
	}
	V = tempV
}
