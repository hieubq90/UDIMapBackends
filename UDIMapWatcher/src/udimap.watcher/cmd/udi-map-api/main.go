package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"

	"udimap.watcher/config"
	"udimap.watcher/models"
)

func main() {
	if config.InitFromYAML() {
		fmt.Printf("UDIMap API - Start Listen On %s\n", config.AppConfig.ServerListenEndpoint)
		db, err := gorm.Open("postgres", config.AppConfig.DBConnection)
		db.DB().SetMaxIdleConns(config.AppConfig.MaxIdleConns)
		db.DB().SetMaxOpenConns(config.AppConfig.MaxOpenConns)
		db.DB().SetConnMaxLifetime(time.Duration(config.AppConfig.ConnMaxLifetime) * time.Minute)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		db.AutoMigrate(&models.Camera{})
		db.AutoMigrate(&models.TramDoMua{})
		db.AutoMigrate(&models.TramDoTrieu{})

		router := httprouter.New()
		server := NewServer(db)
		server.RegisterRouter(router)
		log.Fatal(http.ListenAndServe(config.AppConfig.ServerListenEndpoint, router))
	}

}
