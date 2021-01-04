package db

import (
	"errors"
	"simple-tracking/support"
	"time"

	"github.com/jmoiron/sqlx"
)
import _ "github.com/go-sql-driver/mysql"

var conn  *connection
func init()  { conn = &connection{} }
func Get() *connection  { return conn }


var (
	ERROR_NOT_CONNECTION = errors.New("db not connection")
)

type connection struct {
	dsn    string
	conn    *sqlx.DB
}

func (this *connection) Init()  {
	dsn:= support.Config().DSN
	if dsn == "" {
		support.Config().SetEnableTracking(false)
		support.Log("dsn is not setting .... ")
		return
	}
	this.dsn = dsn
	this.connect()
}

func (this *connection) connect() {
	db,err:=sqlx.Open("mysql",this.dsn)
	if err!=nil{
		support.Log("mysql Connection error ",err.Error())
		support.Config().SetEnableTracking(false)
		return
	}
	if maxOpen := support.Config().MaxOpen; maxOpen > 0{
		db.SetMaxOpenConns( maxOpen )
	}else{
		db.SetMaxOpenConns( 300 )
	}

	db.SetConnMaxLifetime(10)
	db.SetConnMaxLifetime( time.Minute * 3)

	if err:= db.Ping();err!=nil{
		support.Log("mysql connection error",err)
		support.Config().SetEnableTracking(false)
		return
	}
	this.conn = db
}

func (this *connection) DB() *sqlx.DB  {
	return this.conn
}

func (this *connection) IsConnection() bool {
	return this.conn!=nil && this.conn.Ping() ==nil
}

func (this *connection) Close()  {
	if this.IsConnection(){
		if err:=this.conn.Close();err!=nil{
			support.Log("mysql close error ",err.Error())
			return
		}
		support.Log("mysql close success!")
	}
	support.Config().SetEnableTracking(false)
}