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
	Fping Fping `mapstructure:"fping"`
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
	Fping Fping `mapstructure:"fping"`
}

type Host struct {
	Address string `mapstructure:"address"`
	Descr string `mapstructure:"description"`
	Fping Fping `mapstructure:"fping"`
}

func ReadAndParseConfig(conf *Config) {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Unable to read config file", err)
	}
	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}

	if ( ! viper.IsSet("src_host")) {
		src_host, err := os.Hostname()
		if err != nil {
			panic(err)
		}
	conf.Src_host = src_host
	}

	if ( (! viper.IsSet("influx.host")) || (! viper.IsSet("influx.port")) || (! viper.IsSet("influx.user")) || (! viper.IsSet("influx.pass")) || (! viper.IsSet("influx.ssl")) || (! viper.IsSet("influx.db")) ) {
		panic("Missing values for influx!")
	}

	if ( (! viper.IsSet("fping.backoff")) || (! viper.IsSet("fping.retries")) || (! viper.IsSet("fping.tos")) || (! viper.IsSet("fping.summary")) || (! viper.IsSet("fping.period")) ) {
		panic("Missing default values for fping!")
	}

	// setting default fping to group fping and hosts fping values if not set
	for i := range conf.Groups {
		x := conf.Groups[i].Fping
		log.Printf("%v: ", i)
		log.Printf("Groups Fping (old): %+v", conf.Groups[i].Fping)
		if ( ! viper.IsSet("hostgroups." + i + ".fping.backoff")) {
			x.Backoff = conf.Fping.Backoff
		}
		if ( ! viper.IsSet("hostgroups." + i + ".fping.retries")) {
			x.Retries = conf.Fping.Retries
		}
		if ( ! viper.IsSet("hostgroups." + i + ".fping.tos")) {
			x.TOS = conf.Fping.TOS
		}
		if ( ! viper.IsSet("hostgroups." + i + ".fping.summary")) {
			x.Summary = conf.Fping.Summary
		}
		if ( ! viper.IsSet("hostgroups." + i + ".fping.period")) {
			x.Period = conf.Fping.Period
		}
		if ( ! viper.IsSet("hostgroups." + i + ".fping.custom")) {
			x.Custom = conf.Fping.Custom
		}
		conf.Groups[i].Fping = x
		log.Printf("Groups Fping (new): %+v", conf.Groups[i].Fping)

//		if ( ! viper.IsSet("hostgroups." + i + ".fping.backoff")) {
//			x["Backoff"] = conf.Fping.Backoff
//			conf.Groups[i].Fping.Backoff = x["Backoff"]
//			log.Printf("%v", conf.Groups[i].Fping)
//		}
		log.Printf("%v", conf.Groups[i].Fping)
	}
}



func Parse(conf *Config) {
	src_host := conf.src_host
	log.printf("%v", &src_host)
}

func main() {
	var conf Config
	viper.SetConfigName(infping)
	viper.AddConfigPath(".")

	ReadConfig(&conf)
	log.printf("%v", conf)
}