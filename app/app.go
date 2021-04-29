package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/niluwats/covid_log/domain"
	"github.com/niluwats/covid_log/service"
)

func Start() {
	sanityCheck()
	router := mux.NewRouter()
	dbClient := getDbClient()

	visitorRepoDb := domain.NewVisitorRepositoryDb(dbClient)
	vh := VisitorHandlers{service.NewVisitorService(visitorRepoDb)}

	companyRepoDb := domain.NewCompanyRepositoryDb(dbClient)
	ch := CompanyHandlers{service.NewCompanyService(companyRepoDb)}

	logRepoDb := domain.NewLogRepositoryDb(dbClient)
	lh := LogHandlers{service.NewLogService(logRepoDb)}

	router.HandleFunc("/visitors", vh.getAllVisitors).Methods(http.MethodGet)
	router.HandleFunc("/visitors/{nic}", vh.getVisitor).Methods(http.MethodGet)
	router.HandleFunc("/visitors", vh.newVisitor).Methods(http.MethodPost)
	router.HandleFunc("/visitors/{nic}", vh.editVisitor).Methods(http.MethodPatch)
	router.HandleFunc("/visitors/{nic}", vh.removeVisitor).Methods(http.MethodDelete)

	router.HandleFunc("/company", ch.getAllCompanies).Methods(http.MethodGet)
	router.HandleFunc("/company/{id}", ch.getCompany).Methods(http.MethodGet)
	router.HandleFunc("/company", ch.newCompany).Methods(http.MethodPost)
	router.HandleFunc("/company/{id}", ch.editCompany).Methods(http.MethodPatch)
	router.HandleFunc("/company/{id}", ch.removeCompany).Methods(http.MethodDelete)

	router.HandleFunc("/company/{id}/log", lh.newLog).Methods(http.MethodPost)
	router.HandleFunc("/company/{id}/log", lh.getAllLogs).Methods(http.MethodGet)
	router.HandleFunc("/company/{id}/log/{date}", lh.getLogsByDate).Methods(http.MethodGet)
	router.HandleFunc("/visitors/log/{nic}", lh.getLogsByNic).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	server := fmt.Sprintf("%s:%s", address, port)
	log.Fatal(http.ListenAndServe(server, router))
}
func getDbClient() *sqlx.DB {
	db_user := os.Getenv("DB_USER")
	db_pwd := os.Getenv("DB_PASSWORD")
	db_addrs := os.Getenv("DB_ADDRESS")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_pwd, db_addrs, db_port, db_name)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
func sanityCheck() {
	if os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable SERVER PORT not defined")
	}
	if os.Getenv("SERVER_ADDRESS") == "" {
		log.Fatal("Environment variable SERVER ADDRESS not defined")
	}
	if os.Getenv("DB_PORT") == "" {
		log.Fatal("Environment variable DB PORT not defined")
	}
	if os.Getenv("DB_ADDRESS") == "" {
		log.Fatal("Environment variable DB ADDRESS not defined")
	}
	if os.Getenv("DB_NAME") == "" {
		log.Fatal("Environment variable DB NAME not defined")
	}
	if os.Getenv("DB_USER") == "" {
		log.Fatal("Environment variable DB USER not defined")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		log.Fatal("Environment variable DB PASSWORD not defined")
	}
}
