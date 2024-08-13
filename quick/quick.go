package quick

import (
	"app/quick/current"
	"app/quick/db"
	"app/quick/entdb"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var plugins *Plugins

func init() {
	// fmt.Println("quick initialization")

	rootdir, err := current.Dirname()
	if err != nil {
		panic(fmt.Errorf("error getting directory: %w", err))
	}

	rootdir = getPathBeforeQuick(rootdir)

	getAndSetConfig(rootdir)

	db := db.New(rootdir)
	plugins = &Plugins{
		DB: db,
		EntDB: entdb.New(entdb.Config{
			DB: *db,
		}),
	}
}

// fix task quick:migrate it will take the quick directory as root
func getPathBeforeQuick(fullPath string) string {
	// Find the index of "/quick"
	index := strings.Index(fullPath, "/quick")
	if index == -1 {
		return fullPath // "/quick" not found; return original path
	}
	// Return the substring from the beginning to the index of "/quick"
	return fullPath[:index]
}

type Plugins struct {
	DB    *db.DB
	EntDB entdb.EntDB
}

func GetPlugins() *Plugins {
	return plugins
}

func getAndSetConfig(rootdir string) {
	viper.SetConfigName("default") // name of config file (without extension)
	viper.SetConfigType("yaml")

	viper.AddConfigPath(fmt.Sprintf("%s/config", rootdir))

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
