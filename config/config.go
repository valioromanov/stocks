package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/magiconair/properties"
)

type Config struct {
	Host   string `properties:"host"`
	Port   string `properties:"port"`
	DbHost string `properties:"dbHost"`
	DbPort string `properties:"dbPort"`
	DbName string `properties:"dbName"`
	DbUser string `properties:"dbUser"`
	DbPass string `properties:"dbPass"`
}

var Cfg Config

func NewConfig() {
	p := properties.MustLoadFile("config.properties", properties.UTF8)

	if err := p.Decode(&Cfg); err != nil {
		log.Fatal(err)
	}
}

func NewDbConnection() *pgx.Conn {
	dbURL := "host=" + Cfg.DbHost + " dbname=" + Cfg.DbName + " user=" + Cfg.DbUser + " password=" + Cfg.DbPass + " port=" + Cfg.DbPort
	fmt.Println(dbURL)
	client, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}

	return client
}
