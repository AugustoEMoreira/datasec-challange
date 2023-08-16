package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/AugustoEMoreira/datasec-challange/adapter/mysql"
	"github.com/AugustoEMoreira/datasec-challange/di"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn, err := mysql.GetConnection(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	datastorageservice := di.ConfigDatastorage(conn)

	router := mux.NewRouter()
	router.Handle("/api/v1/database", http.HandlerFunc(datastorageservice.Create)).Methods("POST")
	router.Handle("/api/v1/database/scan/{id}", http.HandlerFunc(datastorageservice.Fetch)).Methods("GET")
	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
