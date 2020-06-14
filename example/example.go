package main

import "github.com/ruster-cn/log"

func main() {
	if err := log.NewLogFactory(&log.LogConfiguration{
		Level:   "Debug",
		Console: true,
		File:    true,
		Service: "main",
		FileConfig: &log.LogFileConfig{
			Path:  "./",
			Name:  "main.log",
			Split: false,
		},
	}).Build(); err != nil {
		panic(err)
	}
	log.Warn("test log %s", "log")
}
