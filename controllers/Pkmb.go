package controllers

import (
	"pkmapi/models"
	"pkmapi/global"
	"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
)

type PkmbController struct {
	beego.Controller
}

type GetSurvey struct {
	Survey []models.PkmbSurvey `json:"survey"`
}

type ResultGetSurvey struct {
	Survey bool `json:"survey"`
	Data []GetSurvey `json:"data"`
}

type ResultGetMateri struct {
	Survey bool `json:"survey"`
	Data models.PkmbMateri `json:"data"`
}

// @Title PKM Bersama (menu divisi PKU)
// @Description pkmb untuk data
// @Param	Authorization header string  false "Authorization Token"
// @Param	Username path string  true "Username"
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /get_pkmb/:Username [get]
func (c *PkmbController) GetPkmb() {
	var check = models.Select_PkmbMingguKe()	
	fmt.Println(check)
	// if check == 3 || check == 4 {
		var group = models.Select_Ao_Group(c.Ctx.Input.Param(":Username"))
		survey := make([]GetSurvey, 0)

		for _,val := range group{
			var client = models.Client_By_Group(val.GroupID)
			
			for _,value := range client{
				var each []models.PkmbSurvey
				
				each = models.Select_PkmbSurvey(value.AccountID,value.ClientID,val.GroupID)				

				survey = append(survey, GetSurvey{Survey: each})				
			}		
		}
		// fmt.Println(survey)
		c.Data["json"] = ResultGetSurvey{Survey : true,Data : survey}
	// }else{
	// 	c.Data["json"] = ResultGetMateri{W4: false,Data : models.Select_PkmbMateri()}
	// }

	c.ServeJSON()
}


// @Title PKM Bersama (menu divisi PKU)
// @Description pkmb untuk data
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body []models.PostPkmSurvey true "simpan data pkmb"
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /post_pkmb [post]
func (c *PkmbController) PostPkmb() {
	var data []models.PostPkmSurvey
    err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)

	err = models.Post_Pkm_Survey(data)
	if err != nil {
		global.Logging("ERROR","models.Post_Pkm_Survey ---> " + err.Error())
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}
	
	c.Data["json"] = global.APIResponse{Code: 200, Message: "Data berhasil disimpan"}
    c.ServeJSON()

}