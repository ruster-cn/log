/* Copyright 2019 mingzhuyang All Rights Reserved. */

/* config.go - log config  and factory function */
/*
modification history
--------------------
2019-05-13   mingzhuyang   create


*/

package log

import (
	"fmt"
)

//imports

//Constants

//Typedefs
type LogConfiguration struct {
	Level string `yaml:"level"`
	//Console is true，log will be output to console
	Console bool `yaml:"console"`
	//File is true, log will be output to file
	File    bool   `yaml:"isfile"`
	Service string `yaml:"service"`

	//FileConfig if log output to file,FileConfig must be set
	FileConfig *LogFileConfig `yaml:"file"`
}

type LogFileConfig struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Split bool   `yaml:"split"`
}

type LogFactory struct {
	LogConf *LogConfiguration
}

//Globals

//functions
func (l *LogConfiguration) Validate() []error {
	if l == nil {
		return nil
	}
	errs := []error{}
	if l.Level == "" {
		errs = append(errs, fmt.Errorf("Must specify log output level"))
	}
	if !l.Console && !l.File {
		errs = append(errs, fmt.Errorf("Must specify a log output method, console or file"))
	}
	if l.File && l.FileConfig == nil {
		errs = append(errs, fmt.Errorf("If the log is output to a file, the file path and file name must be configured"))
	}
	if err := l.FileConfig.validate(); err != nil {
		errs = append(errs, err)
	}
	return errs
}

func (f *LogFileConfig) validate() error {
	if f == nil {
		return nil
	}
	if f.Name == "" || f.Path == "" {
		return fmt.Errorf("Please specify file output path and file name")
	}

	return nil
}

func NewLogFactory(l *LogConfiguration) *LogFactory {
	return &LogFactory{
		LogConf: l,
	}
}

func (f *LogFactory) Build() error {
	if f.LogConf.Console {
		logger := NewXConsoleLog()
		err := RegisterLogger("console_log", logger)
		if err != nil {
			return err
		}
		err = InitLoggerFromConfiguration("console_log", f.LogConf)
		if err != nil {
			return err
		}
	}

	if f.LogConf.File {
		logger := NewXFileLog()
		err := RegisterLogger("file_log", logger)
		if err != nil {
			return err
		}
		err = InitLoggerFromConfiguration("file_log", f.LogConf)
		if err != nil {
			return err
		}
	}
	return nil
}
