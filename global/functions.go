package global

import (
    "encoding/base64"
    "strings"    
    "io/ioutil"
    "bytes"
    "os"
    "github.com/beevik/guid"


    // "fmt"
	// cc "golang.org/x/oauth2/clientcredentials"
	// goctx "golang.org/x/net/context"

    "io"
    "image"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"    
    // "golang.org/x/image/tiff"
    // "golang.org/x/image/bmp"
    // "github.com/nfnt/resize"

    "time"
    "crypto/md5"
    "log"
    "github.com/dgrijalva/jwt-go"
    "strconv"
    "github.com/astaxie/beego"    
    "path/filepath"

    // "errors"
)

func MapB64SaveFile(data_arr map[string]string,path string)(map[string]string,error) {

    var fileName = make(map[string]string)
    var err error    
    for key, data := range data_arr {	             
        if (data != ""){             
            fileName[key],err = B64SaveFile(data,path)                            
            if err != nil {
                // Logging("ERROR MapB64SaveFile "+key+" ---> "+data+" ---> ")  
                break
                return fileName,err
            }  
        }
    }
    return fileName,err
}

func B64SaveFile(data string,pathnya string)(string,error) {
    err := os.MkdirAll(pathnya, os.ModePerm)
    if err != nil {
        return "",err
    }

    fileNameBase := guid.New().String()+guid.New().String()    
	
	// klw string data menggunakan data:image/png;base64,
	// idx := strings.Index(data , ";base64,")
    // if idx < 0 {
    //     return ErrInvalidImage
    // }

    // reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))

    reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
    buff := bytes.Buffer{}
    _,err = buff.ReadFrom(reader)
    if err != nil {
        return "",err
    }

    _, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
    if err != nil {
        return "",err
    }

    fileName := pathnya+fileNameBase + "." + fm

	// fmt.Println(fileName)
    err = ioutil.WriteFile(fileName, buff.Bytes(), 0600)
    if err != nil {
        return "",err
    }


    



    // fmt.Println("======= UPLOAD FILE TO FTP SERVER =======")

    conn,err := ConnFTP()
    if err != nil {
        return "",err
    }    

    // err = conn.MakeDir("pkm/images/")
    // if err != nil {
    //     return "",err
    // }
    

    sourceFile := filepath.Clean(fileName)
    f, err := os.Open(sourceFile)
    if err != nil {
        return "",err
    }
        
    destinationFile := fileName
    err = conn.Stor(destinationFile, f)
    if err != nil {
        return "",err
    }
    err = f.Close()
    if err != nil {
        return "",err
    }

    if err := conn.Quit(); err != nil {
        return "",err
    }
    
    // fmt.Println("======= UPLOAD FILE TO FTP SERVER SELESAI =======")


	return fileName,nil
}


func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func FileB64(pathfile string) (string,error) {
    var base64Encoding string

    // fmt.Println("function file b 64")

    // bytes, err := ioutil.ReadFile(pathfile)    
	// if err != nil {
	// 	return base64Encoding,err
	// }
	
	// Append the base64 encoded output
	// base64Encoding = toBase64(bytes)
	

	// Print the full base64 representation of the image
	// fmt.Println(base64Encoding)
    // return base64Encoding,nil


   //=================FTP======================//
   conn,err := ConnFTP()
   if err != nil {
       return base64Encoding,err
   }    
    //    r, err := conn.Retr("home/PKM/coba.png")
   r, err := conn.Retr(pathfile)

   if err != nil {
       return base64Encoding,err
   }   
   
   buf, err := ioutil.ReadAll(r)
   if err != nil {
        return base64Encoding,err
   }   
   // println(string(buf))   

   if err := conn.Quit(); err != nil {
        return base64Encoding,err
   }
   defer r.Close()

    //    fmt.Println(buf)

   base64Encoding = toBase64(buf)

    //    fmt.Println(base64Encoding)
   return base64Encoding,nil
   //=================FTP======================//    
}

// func FileB64Resize(pathfile string) (string,error) {

//     file, err := os.Open(pathfile)
//     if err != nil {
//         return "",err
//     }        
//     defer file.Close()

//     img, filetype, err := image.Decode(file)
//     if err != nil {
//         return "",err
//     }        

//     sizenya64, err := strconv.ParseUint(beego.AppConfig.String("RESIZE_GAMBAR"),10,32)
//     if err != nil {
//         return "",err
//     }        
//     sizenya := uint(sizenya64)

//     newImage := resize.Resize(sizenya, sizenya, img, resize.Lanczos3)

//     fileTemp := "temp/"+guid.New().String()+"."+filetype    

//     w, err := os.Create(fileTemp)
//     if err != nil {
//         return "",err
//     }
//     defer w.Close()        
    


//     switch filetype {
//     case "png":
//         err = png.Encode(w, newImage)
//     case "gif":
//         err = gif.Encode(w, newImage, nil)
//     case "jpeg", "jpg":
//         err = jpeg.Encode(w, newImage, nil)
//     case "bmp":
//         err = bmp.Encode(w, newImage)
//     case "tiff":
//         err = tiff.Encode(w, newImage, nil)
//     default:
//         // not sure how you got here but what are we going to do with you?
//         fmt.Println("Unknown image type: ", filetype)
//         io.Copy(w, file)
//     }        



//     var base64Encoding string

//     bytes, err := ioutil.ReadFile(fileTemp)    
//     if err != nil {
//         return base64Encoding,err
//     }                
//     base64Encoding = toBase64(bytes)    
//     return base64Encoding,nil        
// }

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

// func Oauth2() error {

// 	fmt.Println("cek kesini dlu kan")
// 	// ctx.Output.Header("Content-Type", "application/json")

// 	clientID := "PKM"
// 	clientSecret := "b53d714e-8e21-4271-9eb6-d43a8505b936"


// 	config := &cc.Config{
// 		ClientID:     clientID,
// 		ClientSecret: clientSecret,
// 		TokenURL: "http://127.0.0.1:8080/v1/oauth2/token",
// 		Scopes:   []string{"all"},
// 		AuthStyle: 1,
// 	}

// 	goctx := goctx.Background()
// 	// 	client := config.Client(goctx)
// 	// fmt.Println(client)

// 	token, err := config.Token(goctx)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	fmt.Println(token)


// 	return nil
// }

func GenerateTokenJWT(d string) string {
	var uid int = 0
	currentTimestamp := time.Now().UTC().Unix()
	//var ttl int64 = 3600 //satu jam
	var ttl int64 = 43200 //dua belas jam
	h := md5.New()
	_,err := io.WriteString(h, strconv.Itoa(uid))

	if err != nil {
    	log.Fatal(err)
	}

	_,err = io.WriteString(h, strconv.FormatInt(int64(currentTimestamp), 10))

	if err != nil {
    	log.Fatal(err)
	}	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": uid,
		"iat": currentTimestamp,
		"exp": currentTimestamp + ttl,
		"nbf": currentTimestamp,
		"iss": d,
		"jti": h.Sum(nil),
	})

	tokenString, err := token.SignedString([]byte(beego.AppConfig.String("KEY")))

	if err != nil {
    	log.Fatal(err)
	}

	return (tokenString)
}






// func S3storage(){
//     ctx := context.Background()
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

// 	// Make a new bucket called mymusic.
// 	bucketName := "PKM"
// 	location := "us-east-1"

// 	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
// 	if err != nil {
// 		// Check to see if we already own this bucket (which happens if you run this twice)
// 		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
// 		if errBucketExists == nil && exists {
// 			log.Printf("We already own %s\n", bucketName)
// 		} else {
// 			log.Fatalln(err)
// 		}
// 	} else {
// 		log.Printf("Successfully created %s\n", bucketName)
// 	}

// 	// Upload the zip file
// 	objectName := "absensi-fandi-zainal-january-2020.pdf"
// 	filePath := "/images/golden-oldies.zip"
// 	contentType := "application/pdf"

// 	// Upload the zip file with FPutObject
// 	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
// }