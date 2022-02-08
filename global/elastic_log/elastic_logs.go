package elastic_log

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	// "github.com/elastic/go-elasticsearch/v6"	
	// "github.com/elastic/go-elasticsearch/v6/esapi"
	
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"

	"github.com/astaxie/beego/logs"

	"net"
	"github.com/beevik/guid"
)

// NewES return a LoggerInterface
func NewES() logs.Logger {
	cw := &esLogger{
		Level: logs.LevelDebug,
	}
	return cw
}

// esLogger will log msg into ES
// before you using this implementation,
// please import this package
// usually means that you can import this package in your main package
// for example, anonymous:
// import _ "github.com/astaxie/beego/logs/es"
type esLogger struct {
	*elasticsearch.Client
	DSN   string `json:"dsn"`
	Level int    `json:"level"`
}

// {"dsn":"http://localhost:9200/","level":1}
func (el *esLogger) Init(jsonconfig string) error {
	err := json.Unmarshal([]byte(jsonconfig), el)
	if err != nil {
		return err
	}
	if el.DSN == "" {
		return errors.New("empty dsn")
	} else if u, err := url.Parse(el.DSN); err != nil {
		return err
	} else if u.Path == "" {
		return errors.New("missing prefix")
	} else {
		conn, err := elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{el.DSN},
		})
		if err != nil {
			return err
		}
		el.Client = conn
	}
	return nil
}

// WriteMsg will write the msg and level into es
func (el *esLogger) WriteMsg(when time.Time, msg string, level int) error {
	if level > el.Level {
		return nil
	}

	ip,_ := externalIP()

	var jenis,nama_index string

	msg_array := strings.Split(msg, "||")

	msg = msg_array[0]		

	if len(msg_array)>1{
		jenis = msg_array[1]
	}else{
		jenis = "STDOUT"
	}

	if jenis == "SIMPLE_SYNC_TO_BRNET" {		
		nama_index = fmt.Sprintf("simple.sync.brnet.log.%04d.%02d", when.Year(), when.Month())
	}else if jenis == "PAYLOAD" {				
		nama_index = fmt.Sprintf("payload.pkm.inisiasi.log.%04d.%02d", when.Year(), when.Month())
	}else{		
		nama_index = fmt.Sprintf("pkm.inisiasi.log.%04d.%02d", when.Year(), when.Month())
	}

	idx := LogDocument{
		Ip:        ip,
		Type:      jenis,
		Timestamp: when.Format(time.RFC3339),
		Msg:       msg,
	}		

	body, err := json.Marshal(idx)
	if err != nil {
		return err
	}
	req := esapi.IndexRequest{
		Index:      nama_index,
		DocumentID: guid.New().String(),
		Body:       strings.NewReader(string(body)),
	}
	_, err = req.Do(context.Background(), el.Client)
	return err
}

// Destroy is a empty method
func (el *esLogger) Destroy() {

}

// Flush is a empty method
func (el *esLogger) Flush() {

}

type LogDocument struct {
	Ip		  string `json:"ip"`
	Type	  string `json:"type"`	
	Timestamp string `json:"timestamp"`
	Msg       string `json:"msg"`
}

func init() {
	logs.Register(logs.AdapterEs, NewES)
}


func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

