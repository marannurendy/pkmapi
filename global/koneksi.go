package global

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/denisenkom/go-mssqldb"	
	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/mssql"

	"github.com/jlaffaye/ftp"
	"errors"


	
	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"
)

// var Db *sql.DB

// var Db2 *sql.DB

// var Db3 *sql.DB

// func init() {
// 	Db, _ = sql.Open("mssql", beego.AppConfig.String("SqlConnPKM"))

// 	Db2, _ = sql.Open("mssql", beego.AppConfig.String("SqlConnPelatihan"))

// 	Db3, _ = sql.Open("mssql", beego.AppConfig.String("SqlConnInisiasi"))
// }


func ConnPKM() (*sql.DB,error) {
	conn,err := sql.Open("mssql", beego.AppConfig.String("SqlConnPKM"))
	if err != nil {
		return nil,err
	}

	err = conn.Ping()
	if err != nil {
		return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}	
	
	return conn,nil	
}

func Conn() (*godb.DB,error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("SqlConnInisiasi"))
	if err != nil {
		return nil,err
	}

	err = conn.CobaPing()
	if err != nil {
		return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}

	return conn,nil
}

func ConnINISIASI_WITHTIME() (*godb.DB,error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("SqlConnInisiasi_WITHTIME"))
	if err != nil {
		return nil,err
	}

	err = conn.CobaPing()
	if err != nil {
		return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}


	return conn,nil
}

func ConnPKU() (*godb.DB,error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("SqlConnPelatihan"))
	if err != nil {
		return nil,err
	}

	err = conn.CobaPing()
	if err != nil {
		return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}


	return conn,nil
}

func ConnPKUold() (*sql.DB,error) {
	conn,err := sql.Open("mssql", beego.AppConfig.String("SqlConnPelatihan"))
	if err != nil {
		return nil,err
	}

	err = conn.Ping()
	if err != nil {
		return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}	

	return conn,nil	
}

func ConnBRNET_GET() (*godb.DB,error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("SqlConnBRNET_GET"))
	if err != nil {
		return nil,err
	}

	err = conn.CobaPing()
	if err != nil {
		return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}


	return conn,nil
}


func ConnBRNET_POST() (*godb.DB,error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("SqlConnBRNET_POST"))
	if err != nil {
		return nil,err
	}

	err = conn.CobaPing()
	if err != nil {
		return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}

	return conn,nil
}

func ConnSCORING() (*godb.DB,error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("SqlConnSCORING"))
	if err != nil {
		return nil,err
	}

	err = conn.CobaPing()
	if err != nil {
		return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}


	return conn,nil
}

func ConnFTP()(*ftp.ServerConn,error){
	
	const FTP_ADDR = "10.61.4.110:21"
	const FTP_USERNAME = "itd"
	const FTP_PASSWORD = "pnm123#"

	conn, err := ftp.Dial(FTP_ADDR)
	if err != nil {
		return nil,err
	}

	err = conn.Login(FTP_USERNAME, FTP_PASSWORD)
	if err != nil {
		return nil,err
	}

	return conn,nil
}


func ConnS3Storage()(*minio.Client,error){
    
	endpoint := "pnmdc-cluster-cohesity.pnm.co.id:3000"
	accessKeyID := "MFwH9ODLzSKpPmFiymfZSFSwahkJqYNF-fMcR0C4hnY"
	secretAccessKey := "B07xJAOJHYyOxk0cdA1TI9lBcMSA7dCASEcBvPRfNlk"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return minioClient,err
	}

	return minioClient,nil
}