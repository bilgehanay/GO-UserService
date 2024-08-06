package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	configFile string
	configType string
	config     ConfigModel
	userDB     *mongo.Collection
	cl         = make(map[string]*mongo.Collection)
	ctx        context.Context
)

func init() {
	flag.StringVar(&configFile, "c", "config", "Config File Name")
	flag.StringVar(&configType, "t", "json", "Config File Type")
	flag.Parse()

	viper.SetConfigName(configFile)
	viper.SetConfigType(configType)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	JWTKey = []byte(config.JWT)

	for _, mcnf := range config.Mongo {
		mongoConn := options.Client().ApplyURI(mcnf.ConnectionString)
		mongoConn.SetAppName(mcnf.ConnectionName)
		mc, err := mongo.Connect(ctx, mongoConn)
		if err != nil {
			panic(err)
		}
		for _, dc := range mcnf.Collection {
			cl[dc.N] = mc.Database(dc.D).Collection(dc.C)
		}
	}
	userDB = cl[config.Mongo["users"].Collection["users"].N]
	if userDB == nil {
		fmt.Println("Db can not initilazied")
	}
	fmt.Println("Db initilazied")

}
