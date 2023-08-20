package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/AugustoEMoreira/datasec-challange/adapter/mysql"
	dependencyinjector "github.com/AugustoEMoreira/datasec-challange/dependencyInjector"
	"github.com/gorilla/handlers"
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

	fmt.Println("server starting")
	ctx := context.Background()
	fmt.Println("Connecting into the main database")
	conn, err := mysql.GetMainDBConnection(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	datastorageservice := dependencyinjector.ConfigDatastorage(conn)
	scanservice := dependencyinjector.ConfigScan(conn)

	fmt.Println("Seting up routes")
	router := mux.NewRouter()
	router.Handle("/api/v1/database", http.HandlerFunc(datastorageservice.Create)).Methods("POST")
	router.Handle("/api/v1/database/scan/{id}", http.HandlerFunc(datastorageservice.Fetch)).Methods("GET")
	router.Handle("/api/v1/database/scan/{id}", http.HandlerFunc(scanservice.Create)).Methods("POST")

	port := viper.GetString("server.port")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	fmt.Println("Server running on port: " + port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), loggedRouter)
}
