package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/iAmPlus/microservice/models/apimodels"
)

var (
	errMissingDBCredentials = fmt.Errorf("Database Credentials Missing")
	maxDuration             = 100 * time.Second
	indexDuration           = 4 * time.Minute
)

func NewManager(DBUrl string, DBname string) (Manager, error) {
	dbOpts := CreateOptsFromConfig(DBUrl, DBname)
	if dbOpts.DatabaseURI == "" || dbOpts.DatabaseName == "" {
		return nil, errMissingDBCredentials
	}
	db, err := sql.Open("mysql", "root:pass123@tcp(localhost:3306)/school")
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Unable to create MYSQL DB client, URL %s, opts : %v : %v", dbOpts.DatabaseURI, dbOpts, err)
	}
	defer db.Close()

	um := &manager{dbOpts: dbOpts, client: db}
	err = um.initiateTables()
	if err != nil {
		return nil, fmt.Errorf("Unable to create MYSQL DB teacher manager URL %s, opts : %v : %v", dbOpts.DatabaseURI, dbOpts, err)
	}

	return um, nil
}
func CreateOptsFromConfig(DBUrl string, DBname string) DatabaseOpts {
	return DatabaseOpts{
		DatabaseURI:       DBUrl,
		DatabaseName:      DBname,
		Users_coll:        "users",
		Posts_coll:        "posts",
		Related_post_coll: "related",
	}
}

type Opts struct {
	// time in seconds for considering dialogue records for context.
	RecentDialogueThreshold float64
	// time in seconds for considering dialogue continuation
	NewDialogueThreshold float64
}
type manager struct {
	dbOpts DatabaseOpts
	client *sql.DB
}

// DatabaseOpts options for passing database details
type DatabaseOpts struct {
	// Uri of databse to connect to
	DatabaseURI string
	// Name of database to connect
	DatabaseName      string
	Users_coll        string
	Posts_coll        string
	Related_post_coll string
}

func (um *manager) initiateTables() error {
	var err error
	err = um.createTeacherTable()
	if err != nil {
		return err
	}
	err = um.createStudentTable()
	if err != nil {
		return err
	}
	return nil
}

func (u *manager) createStudentTable() error {
	query := `CREATE TABLE IF NOT EXISTS Student(student_unique_id int primary key auto_increment, student_id text, student_email_id text,student_status text,teacher_unique_id int,FOREIGN KEY (teacher_unique_id) REFERENCES Teacher(teacher_unique_id))`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := u.client.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}
func (u *manager) createTeacherTable() error {
	query := `CREATE TABLE IF NOT EXISTS Teacher(teacher_unique_id int primary key auto_increment, teacher_id text, teacher_email_id text)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := u.client.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

func (u *manager) Register(teacher_ID string, register apimodels.Register) error {
	return nil
}
func (u *manager) Getcommonstudents(teacher_ID string) (apimodels.CommonStudents, error) {
	return apimodels.CommonStudents{}, nil
}
func (u *manager) SuspendStudent(teacher_ID string, Pagestate apimodels.SuspendStudents) error {
	return nil
}
func (u *manager) Retrievefornotifications(teacher_ID apimodels.RetrieveForNotifications) (apimodels.Recipients, error) {
	return apimodels.Recipients{}, nil
}
