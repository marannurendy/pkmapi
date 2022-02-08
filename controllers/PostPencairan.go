package controllers

import (
	"pkmapi/models"
	"pkmapi/global"
	"github.com/astaxie/beego"
	// "fmt"
	"encoding/json"	
	// "github.com/astaxie/beego/validation"
)

// Post Pencairan PKM Mobile
type PostPencairanController struct {
	beego.Controller
}


// @Title PKM INISIASI
// @Description PKM PROSPEK DAN UK
// @Param   data body []models.SP_ADD_PENCAIRAN true "simpan data"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql / validasi"}
// @router /post_pencairan [post]
func (c *PostInisiasiController) PostPencairan() {	
	var data []models.SP_ADD_PENCAIRAN		

    err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {		
		global.Logging("ERROR","controllers.PostPencairan json.Unmarshal ---> " + err.Error())		
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}	

	err = models.POST_PENCAIRAN(data)
	if err != nil {
		global.Logging("ERROR","controllers.PostPencairan models.POST_PENCAIRAN ---> " + err.Error())
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}


	global.Logging("INFO","controllers.PostPencairan SUCCESS ")	
	
	c.Data["json"] = global.APIResponse{Code: 200, Message: "Data berhasil disimpan"}
    c.ServeJSON()
}