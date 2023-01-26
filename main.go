package main

import (
	"fmt"
	"log"
	"net/http"

	middleware "file_storage_service/internal/handler"
	handlerFile "file_storage_service/internal/handler/file"
	handlerUser "file_storage_service/internal/handler/user"
	driver "file_storage_service/internal/repository/database"
	fileDriver "file_storage_service/internal/repository/local_storage"
	usecaseFile "file_storage_service/internal/usecase/file"
	usecaseUser "file_storage_service/internal/usecase/user"

	"github.com/gorilla/mux"
	config "github.com/spf13/viper"
)

// main is where the program started
func main() {
	var (
		err error
	)

	// read configuration
	config.SetConfigFile("configurations/App.yaml")
	config.SetConfigType("yaml")
	err = config.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}

	// set database attributes
	conf := driver.MySQLConfig{
		Host:     config.GetString("database.host"),
		User:     config.GetString("database.user"),
		Password: config.GetString("database.password"),
		Port:     config.GetString("database.port"),
		Db:       config.GetString("database.name"),
	}

	db, err := driver.ConnectToMySQL(conf)
	if err != nil {
		log.Println("error when connect to DB", err)
		return
	}

	datastore := driver.NewRepository(db)
	userUsecase := usecaseUser.NewUsecase(datastore)
	userHandler := handlerUser.NewHandler(userUsecase)

	fileDrive := fileDriver.NewRepository()
	fileUsecase := usecaseFile.NewUsecase(fileDrive, datastore)
	fileHandler := handlerFile.NewHandler(fileUsecase)

	// TODO: move this router to route.go file
	r := mux.NewRouter()

	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("/register", userHandler.Register).Methods("POST")
	user.HandleFunc("/login", userHandler.Login).Methods("POST")

	file := r.PathPrefix("/file").Subrouter()
	file.HandleFunc("/upload", fileHandler.UploadFile).Methods("POST")
	file.HandleFunc("/", fileHandler.GetAllFiles).Methods("GET") // for internal API (using authorization)
	file.Use(middleware.JWTMiddleware)

	// for public API
	// TODO: make another handler function
	publicFile := r.PathPrefix("/ex/file").Subrouter()
	publicFile.HandleFunc("/download", fileHandler.GetAllFiles).Methods("GET")

	fmt.Println("Running..")
	fmt.Println(http.ListenAndServe(":"+config.GetString("app.port"), r))
}
