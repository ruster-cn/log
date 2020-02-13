package main

import "github.com/ruster-cn/log"

func main(){
	log.NewLogFactory(&log.LogConfiguration{
		Level:      "Debug",
		Console:    true,
		File:       true,
		Service:    "main",
		FileConfig: &log.LogFileConfig{
			Path:  "./",
			Name:  "main.log",
			Split: false,
		},
	}).Build()
	log.Warn("test log %s","log")
}
