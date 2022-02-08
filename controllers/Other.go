package controllers

import (
	"pkmapi/models"
	"pkmapi/global"
	"github.com/astaxie/beego"
	// "fmt"
	"encoding/json"	
	// "github.com/astaxie/beego/validation"
	"io"
	"io/ioutil"
	"bytes"

	// "context"
	// "strconv"
	// elastic "github.com/olivere/elastic"
)

// Post Audit PKM Mobile
type OtherController struct {
	beego.Controller
}



// @Title PKM INISIASI
// @Description PKM AUDIT TRAIL
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body models.AUDIT_TRAILT true "simpan data"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /post_audit_trail [post]
func (c *OtherController) PostAuditTrail() {	
	var data models.AUDIT_TRAILT
	// var data map[string]interface{}

    err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {
		global.Logging("ERROR","JSON controller PostAuditTrail ---> " + err.Error())			
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}	

	
	err = models.POST_AUDIT_TRAIL(data)
	if err != nil {
		global.Logging("ERROR","models.POST_AUDIT_TRAIL ---> " + err.Error())			
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}
	

	c.Data["json"] = global.APIResponse{Code: 200, Message: "Data berhasil disimpan"}
    c.ServeJSON()
}



// Post Audit PKM Mobile
type ImagesController struct {
	beego.Controller
}

// @Title PkmImages
// @Description PkmImages
// @Param	Authorization header string  false "Authorization Token"
// @Param   data path string true "path file"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /home/PKM/:data [get]
func (c *ImagesController) PkmImages() {	

	conn,err := global.ConnFTP()
    if err != nil {
		global.Logging("ERROR","global.ConnFTP controller PkmImages ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }    

	// fmt.Println(beego.AppConfig.String("pathimages")+c.Ctx.Input.Param(":data"))

    r, err := conn.Retr(beego.AppConfig.String("pathimages")+c.Ctx.Input.Param(":data"))	
    // r, err := conn.Retr("home/PKM/coba.png")	
    if err != nil {
		global.Logging("ERROR","conn.Retr controller PkmImages ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }    


	if err := conn.Quit(); err != nil {
		global.Logging("ERROR","conn.Quit controller PkmImages ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }	
	defer r.Close()

    buf, err := ioutil.ReadAll(r)	
	// println(string(buf))
	// c.Ctx.Output.Header("Content-Type", "application/octet-stream")	
	if _, err := io.Copy(c.Ctx.ResponseWriter,bytes.NewReader(buf)); err != nil {
	// if _, err := io.CopyBuffer(c.Ctx.ResponseWriter,bytes.NewReader(buf),buf); err != nil {
		global.Logging("ERROR","io.Copy controller PkmImages ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }		
}


// @Title PkmImages
// @Description PkmImages
// @Param	Authorization header string  false "Authorization Token"
// @Param   data path string true "path file"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /home/PKM_PROD/:data [get]
func (c *ImagesController) PkmImagesProd() {	

	conn,err := global.ConnFTP()
    if err != nil {
		global.Logging("ERROR","global.ConnFTP controller PkmImages ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }    

	// fmt.Println(beego.AppConfig.String("pathimages")+c.Ctx.Input.Param(":data"))

    r, err := conn.Retr(beego.AppConfig.String("pathimages")+c.Ctx.Input.Param(":data"))	
    // r, err := conn.Retr("home/PKM/coba.png")	
    if err != nil {
		global.Logging("ERROR","conn.Retr controller PkmImages ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }    


	if err := conn.Quit(); err != nil {
		global.Logging("ERROR","conn.Quit controller PkmImages ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }	
	defer r.Close()

    buf, err := ioutil.ReadAll(r)	
	// println(string(buf))
	// c.Ctx.Output.Header("Content-Type", "application/octet-stream")	
	if _, err := io.Copy(c.Ctx.ResponseWriter,bytes.NewReader(buf)); err != nil {
	// if _, err := io.CopyBuffer(c.Ctx.ResponseWriter,bytes.NewReader(buf),buf); err != nil {
		global.Logging("ERROR","io.Copy controller PkmImages ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }		
}



// @Title Logout
// @Description Logout Hapus Prospek yang ke ambil oleh user
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body models.LogoutReleaseProspek true "data LogoutReleaseProspek"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /AuthLogout [post]
func (c *OtherController) AuthLogout() {
	var params models.LogoutReleaseProspek	


    err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	if err != nil {
		global.Logging("ERROR","controllers.AuthLogout json.Unmarshal ---> " + err.Error())
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}			

	for _, value := range params.ID_Prospek { 			
		err := models.DELETE_RKH_ID_PROSPEK_TERPAKAI(value) //modelnya di Other
		if err != nil {		
			global.Logging("ERROR","controllers.PersetujuanPembiayaan models.DELETE_RKH_ID_PROSPEK_TERPAKAI ---> " + err.Error())		
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
			c.ServeJSON()
		}		
	}

	c.Data["json"] = global.APIResponse{Code: 200, Message: "Berhasil Logout"}
    c.ServeJSON()	
}





/*
// @Title GetIndexElastic
// @Description Get Data Elastic 
// @Param	Authorization header string  false "Authorization Token"
// @Param   INDEX path string true "INDEX"
// @Param   START path string true "START"
// @Param   LIMIT path string true "LIMIT"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /GetIndexElastic/:INDEX/:START/:LIMIT [get]
func (c *OtherController) GetIndexElastic() {
	ctx := context.Background()
	esclient, err := models.GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}	

	start,_ := strconv.Atoi(c.Ctx.Input.Param(":START"))
	limit,_ := strconv.Atoi(c.Ctx.Input.Param(":LIMIT"))


	var dataraw []models.RawDataElastic

	// searchSource := elastic.NewSearchSource()
	// var q,t string
	// c.Ctx.Input.Bind(&q, "q") //string yang di cari
	// c.Ctx.Input.Bind(&t, "t") //nama kolom yang di cari
	// searchSource.Query(elastic.NewMatchQuery(t, q))	
	// queryStr, err1 := searchSource.Source()
	// queryJs, err2 := json.Marshal(queryStr)
	// if err1 != nil || err2 != nil {
	// 	fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	// }
	// fmt.Println(string(queryJs))

	searchService := esclient.Search().Index(c.Ctx.Input.Param(":INDEX")).From(start).Size(limit)
	
	// searchService = searchService.SearchSource(searchSource)
    
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var row models.RawDataElastic

		err := json.Unmarshal(hit.Source, &row)
		if err != nil {
			fmt.Println(err)
		}
		dataraw = append(dataraw, row)
	}

	Response := models.JsonElastic{TotalCount:10000, IncompleteResults: false, Items: dataraw }

	c.Data["json"] = Response
	c.ServeJSON()
}
*/

/*
// @Title TESTING
// @Description TESTING
// @Param	Authorization header string  false "Authorization Token"
// @Param   data formData string true "simpan data"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /testing_simpan_file [post]
func (c *ImagesController) TestingSimpanFile() {	
	var fileimages = make(map[string]string)
	fileimages["TTD_AO"] = c.GetString("data")

	fileimages,err := global.MapB64SaveFile(fileimages,"PKM/images/")
	// fileimages,err = global.MapSaveFile(fileimages,pathimages)
	if err != nil { 				
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}
	

	c.Data["json"] = global.APIResponse{Code: 200, Message: "Data berhasil disimpan"}
    c.ServeJSON()
}
*/