package main

import (
	"time"

	"git.querycap.com/tools/confredis"
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
	"github.com/garyburd/redigo/redis"
	"github.com/go-courier/envconf"
	"github.com/sirupsen/logrus"
)

var (
	myredis = &confredis.Redis{
		Host:           "127.0.0.1",
		Port:           6379,
		Password:       "",
		DB:             3,
		ConnectTimeout: envconf.Duration(30 * time.Minute),
		ReadTimeout:    envconf.Duration(30 * time.Minute),
	}
)

func main() {
	start()
}

func start() {
	color.Green("start:", time.Now().String())

	myredis.SetDefaults()
	myredis.Init()
	redisConn := myredis.Get()

	defer func() {
		err := redisConn.Close()
		if err != nil {
			logrus.Warnf("redis conn close error,[err:%s]", err.Error())
		}
	}()

	color.Green("enter for:", time.Now().String())

	for {

		vs, err := redis.Values(redisConn.Do("brpop", myredis.Prefix("tangxin"), 10))
		color.Green("redisConn.Do(brpop):", time.Now().String())

		if err != nil {
			// logrus.Errorf("read list err,[err:%s]", err.Error())
			color.Red("read list err,[err:%s]", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		spew.Dump(vs)
		time.Sleep(3 * time.Second)

	}
}
