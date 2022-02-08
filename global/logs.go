package global

import (
	beego "github.com/astaxie/beego"
	// "fmt"
	"github.com/astaxie/beego/logs"
	"pkmapi/global/elastic_log"
	"log"
	"time"	
	// "strings"
)

// Delete on Kibana Dev Tools
// POST pkm.inisiasi.log/_search
// {
//   "query": {
//     "range": {
//       "timestamp":{
//         "lt":"2021-12-30T03:00:00.000Z"
//       }
//     }
//   }
// }

func Logging(jenis string,message string) {	
	ElasticLog := elastic_log.NewES()
	err := ElasticLog.Init(`{"dsn":"`+beego.AppConfig.String("Elastic")+`","level":7}`)
	if err != nil {
    	log.Fatal(err)
	}

	// Keterangan	
	// 	jenis = "SIMPLE_SYNC_TO_BRNET"
	// 	jenis = "INFO_HIT"
	// 	jenis = "SYNC_TO_BRNET"	
	// 	jenis = "INFO_PKM"					
	// 	jenis = "ERROR_PKM"					
	// 	jenis = "INFO"
	// 	jenis = "ERROR"	
	//  jenis = "PAYLOAD"

	message = message+"||"+jenis

	err = ElasticLog.WriteMsg(time.Now(),message,7)
	if err != nil {
    	log.Fatal(err)
	}
}

func MainLog(){	
	err := logs.SetLogger("es",`{"dsn":"`+beego.AppConfig.String("Elastic")+`","level":7}`)
	if err != nil {
    	log.Fatal(err)
	}
}



// func Logging(message string) {	
	// l := logs.NewLogger(10000)
	// err := l.SetLogger(logs.AdapterFile,`{"filename":"logs/pkmmobile.log","level":7,"maxsize":1000000,"daily":true,"maxdays":30,"color":true , "perm":"777"}`)
	// if err != nil {
    // 	log.Fatal(err)
	// }
	// l.Error(message)	
// }

// func LogSyncBRNET(message string) {	
// 	l := logs.NewLogger(10000)
// 	err := l.SetLogger(logs.AdapterFile,`{"filename":"logs/SYNC_BR_NET.log","level":7,"maxsize":1000000,"daily":true,"maxdays":30,"color":true , "perm":"777"}`)
// 	if err != nil {
//     	log.Fatal(err)
// 	}
// 	l.Info(message)
// }