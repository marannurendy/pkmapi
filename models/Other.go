package models

import (
	"pkmapi/global"
	// "strconv"
	// "github.com/astaxie/beego/validation"	
	// "time"	
	"fmt"	
	// "strings"
	// "errors"		

	"encoding/json"	
	elastic "github.com/olivere/elastic"
	"github.com/astaxie/beego"
)

type AUDIT_TRAILT struct {
	Data []SP_ADD_AUDIT_TRAILT `json:"data"`
}

type SP_ADD_AUDIT_TRAILT struct {
	ID_Prospek string `json:"ID_Prospek" db:"ID_Prospek"`
	Field string `json:"Field" db:"Field"`
	Change_Log string `json:"Change_Log" db:"Change_Log"`
	Change_By string `json:"Change_By" db:"Change_By"`
}


func POST_AUDIT_TRAIL(params AUDIT_TRAILT)error {
	db,err := global.Conn()
	if err != nil {
		return err
	}	
	var query string

	err = db.Begin() //Begin Transaction
	if err != nil {return err}		

	for _, data := range params.Data {

		query = "EXEC ADD_AUDIT_TRAILT "+
		",@ID_Prospek = '"+data.ID_Prospek+"' "+
		",@Field = '"+data.Field+"' "+
		",@Change_Log = '"+data.Change_Log+"' "+
		",@Change_By = '"+data.Change_By+"' "			

		_ , err = db.RawSQL(query).DoWithIterator()
		if err != nil {		
			errsql := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {return err}
			return errsql
		}			

	}

	err = db.Commit() //Commit Transaction
	if err != nil {return err}

	err = db.Close()
	if err != nil {return err}	

	return nil	

}


func INSERT_RKH_ID_PROSPEK_TERPAKAI(ID_Prospek []string,By string)([]string,error){
	var ID_Prospek_Allow []string
	var query string	

	db,err := global.Conn()
	if err != nil {
		return ID_Prospek_Allow,err
	}			

	for _, v := range ID_Prospek {
		query = "INSERT INTO TEMP_RKH_PROSPEK WITH (ROWLOCK) (ID_Prospek_Terpakai,Terpakai_Oleh) values ('"+v+"','"+By+"') "
		_ , err = db.RawSQL(query).DoWithIterator()
		if err == nil {		
			ID_Prospek_Allow = append(ID_Prospek_Allow,v)
		}else{
			global.Logging("ERROR","models.INSERT_RKH_ID_PROSPEK_TERPAKAI ID_Prospek_Terpakai "+v+" ---> " + err.Error())		
		}		
	}	


	err = db.Close()
	if err != nil {return ID_Prospek_Allow,err}	

	return ID_Prospek_Allow,nil	
}


type LogoutReleaseProspek struct {
	ID_Prospek []string `json:"ID_Prospek"`
}


func DELETE_RKH_ID_PROSPEK_TERPAKAI(ID_Prospek string)error{
	db,err := global.Conn()
	if err != nil {
		return err
	}	
	var query string

	err = db.Begin() //Begin Transaction
	if err != nil {return err}		

	query = "DELETE TEMP_RKH_PROSPEK WHERE ID_Prospek_Terpakai='"+ID_Prospek+"' "
	_ , err = db.RawSQL(query).DoWithIterator()
	if err != nil {		
		errsql := err
		err = db.Rollback() //Rollback Transaction
		if err != nil {return err}
		return errsql
	}		


	err = db.Commit() //Commit Transaction
	if err != nil {return err}

	err = db.Close()
	if err != nil {return err}	

	return nil		
}



type JsonElastic struct {
	TotalCount int `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`	
	Items []RawDataElastic `json:"items"`	
}

type RawDataElastic struct {
	json.RawMessage
}

func GetESClient() (*elastic.Client, error) {
	client, err :=  elastic.NewClient(elastic.SetURL(beego.AppConfig.String("Elastic")),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err
}