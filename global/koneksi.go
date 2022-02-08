package global

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/denisenkom/go-mssqldb"	
	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/mssql"

	"github.com/jlaffaye/ftp"
	// "errors"
	// "fmt"
	// "github.com/minio/minio-go"
	// "github.com/minio/minio-go/pkg/credentials"
	// "golang.org/x/net/context"
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
	return conn,nil	
}

func Conn() (*godb.DB,error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("SqlConnInisiasi"))
	if err != nil {
		return nil,err
	}
	return conn,nil
}

func ConnPKU() (*godb.DB,error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("SqlConnPelatihan"))
	if err != nil {
		return nil,err
	}
	return conn,nil
}

func ConnPKUold() (*sql.DB,error) {
	conn,err := sql.Open("mssql", beego.AppConfig.String("SqlConnPelatihan"))
	if err != nil {
		return nil,err
	}
	return conn,nil	
}

func ConnBRNET_GET() (*godb.DB,error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("SqlConnBRNET_GET"))
	if err != nil {
		return nil,err
	}
	return conn,nil
}


func ConnBRNET_POST() (*godb.DB,error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("SqlConnBRNET_POST"))
	if err != nil {
		return nil,err
	}
	return conn,nil
}

func ConnSCORING() (*godb.DB,error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("SqlConnSCORING"))
	if err != nil {
		return nil,err
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


// func init() {
// 	endpoint := "pnmdc-cluster-cohesity.pnm.co.id:3000"
// 	accessKeyID := "21jEi3Oaa-90SqunmmBVAgR6qKYsah-iDM2N9ewRKcw"
// 	secretAccessKey := "FLCnMUmpIg8qepylN8c7nVmop_vX2dW1XFBlvQXmaTo"
// 	useSSL := false

// 	// Initialize minio client object.
// 	minioClient, err := minio.New(endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
// 		Secure: useSSL,
// 	})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	log.Printf("%#v\n", minioClient) // minioClient is now set up
// }




// func Upload_file_s3()error{
// 	fmt.Println("serius nih")

//     ctx := context.Background()
// 	endpoint := "pnmdc-cluster-cohesity.pnm.co.id:3000"
// 	accessKeyID := "MFwH9ODLzSKpPmFiymfZSFSwahkJqYNF-fMcR0C4hnY"
// 	secretAccessKey := "B07xJAOJHYyOxk0cdA1TI9lBcMSA7dCASEcBvPRfNlk"
// 	useSSL := true

// 	// Initialize minio client object.
// 	minioClient, err := minio.New(endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
// 		Secure: useSSL,
// 	})
// 	if err != nil {
// 		// log.Fatalln(err)
// 		return err
// 	}

	
// 	fmt.Println("minioClient set ")

// 	// Make a new bucket called mymusic.
// 	bucketName := "mekdi"
// 	// location := "us-east-1"
// 	// location := ""

// 	// err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
// 	// if err != nil {
// 	// 	// Check to see if we already own this bucket (which happens if you run this twice)
// 	// 	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
// 	// 	if errBucketExists == nil && exists {
// 	// 		// log.Printf("We already own %s\n", bucketName)
// 	// 		fmt.Println("We already own %s\n", bucketName)			
// 	// 	} else {
// 	// 		// log.Fatalln(err)
// 	// 		return 
// 	// 	}
// 	// } else {
// 	// 	// log.Printf("Successfully created %s\n", bucketName)
// 	// 	fmt.Println("Successfully created %s\n", bucketName)
// 	// }


// 	// fmt.Println("selesai bucket")

// 	// Upload the zip file
// 	objectName := "absensi-fandi-zainal-january-2020.pdf"
// 	filePath := "./images/absensi-fandi-zainal-january-2020.pdf"
// 	contentType := "application/pdf"

// 	// Upload the zip file with FPutObject
// 	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
// 	if err != nil {
// 		// log.Fatalln(err)
// 		return err
// 	}

// 	// log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
// 	fmt.Println("Successfully uploaded %s of size %d\n", objectName, info.Size)
// 	return nil
// }