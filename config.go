/*
The MIT License (MIT)

Copyright (c) 2018 Bostjan Bele           https://github.com/MuadDib81/infping

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal 
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"log"
	"github.com/spf13/viper"
)

//structures to unmarshal the config
type Config struct {
	Src_host string `mapstructure:"src_host"`
	Influx Influx `mapstructure:"influx"`
	Infping Infping `mapstructure:"infping"`
	Groups map[string]Group `mapstructure:"hostgroups"`
}

type Influx struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
	SSL string `mapstructure:"ssl"`
	DB string `mapstructure:"db"`
}

type Infping struct {
	Backoff int `mapstructure:"backoff"`
	Retries int `mapstructure:"retries"`
	TOS int `mapstructure:"tos"`
	Summary int `mapstructure:"summary"`
	Period int `mapstructure:"period"`
	Custom map[string]string `mapstructure:"custom"`
}

type Group struct {
	Hosts []Host `mapstructure:"hosts"`
}

type Host struct {
	Address string `mapstructure:"address"`
	Descr string `mapstructure:"description"`
	Infping Infping `mapstructure:"infping"`
}

func ReadConfig(conf *Config) {
	if err := viper.ReadInConfig(); err != nil {
        log.Fatal("Unable to read config file", err)
    }
	
	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}
	
}

func main() {
	var conf Config
	viper.SetConfigName(infping)
	viper.AddConfigPath(".")

	ReadConfig(&conf)
	log.printf("%v", conf)
}