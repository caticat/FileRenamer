package main

import (
	"conf"
	"file"
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
)

var (
	configPath = flag.String("configPath", "FileRenamer.ini", "Configureation file.")
)

type Config struct {
	path  string
	index int
}

func main() {
	var fileV []string
	var err *error
	conf := LoadConfig()
	fileV, err = file.GetFilelist(conf.path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for _, v := range fileV {
		tar := fmt.Sprintf("%s/%d%s", path.Dir(file.FixPath(v)), conf.index, path.Ext(v))
		os.Rename(v, tar)
		fmt.Println("file: ", v, "\t=>\t", tar)
		conf.index++
	}
}

func LoadConfig() Config {
	pfp := file.FixPath(*configPath)
	config := conf.LoadConfig(&pfp)
	index, _ := strconv.Atoi(config.Read("basic", "index"))
	return Config{config.Read("basic", "path"), index}
}
