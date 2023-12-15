package config

import (
	"context"
	"log"
	"server04/ent"

	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

var (
	client *ent.Client
)

func GetClient() *ent.Client {
	return client
}

func SetClient(newClient *ent.Client) {
	client = newClient
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
)

func NewEntClient() (*ent.Client, error) {

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
	// 	os.Getenv("DB_USERNAME"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	// client, err := ent.Open("mysql", dsn, ent.Debug(), ent.Log(func(i ...interface{}) {
	// 	for _, v := range i {
	// 		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), v)
	// 		fmt.Print("\n")
	// 	}
	// }))
	////option 1
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+	"password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	client, err := ent.Open("postgres", "dbname=video_sharing  sslmode=disable")
	////option 2
	//client, err := ent.Open("postgres", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, err
}
