package controllers

import (
	"pkmapi/models"
	"pkmapi/global"
	"github.com/astaxie/beego"
	// "fmt"
	"encoding/json"	
	// "github.com/astaxie/beego/validation"
	"time"
)

// Inisiasi Mobile
type MasterDataController struct {
	beego.Controller
}

// @Title MASTER_DATA
// @Description master_data
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body models.PARAM_DATA_SOSIALISASI true "parameter get data"
// @Success 200 {object} models.GET_DATA_SOSIALISASI
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /GetSosialisasiMobile [post]
func (c *MasterDataController) Get_Sosialisasi_Mobile() {
	var params models.PARAM_DATA_SOSIALISASI	

	global.Logging("INFO_HIT"," ================ v1/inisiasi/GetSosialisasiMobile ================ ")	

	start_Get_Sosialisasi_Mobile := time.Now()
	global.Logging("PAYLOAD","INFO controllers.Get_Sosialisasi_Mobile RequestBody ---> " +string(c.Ctx.Input.RequestBody))		

    err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	if err != nil {
		global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile json.Unmarshal ---> " + err.Error())
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}		

	if (params.Check_ID_Prospek=="1"){
		params.ID_Prospek,err = models.INSERT_RKH_ID_PROSPEK_TERPAKAI(params.ID_Prospek,params.Check_ID_By) //modelnya di Other
		if err != nil {		
			global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile models.INSERT_RKH_ID_PROSPEK_TERPAKAI ---> " + err.Error())		
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
			c.ServeJSON()	
		}	
	}

	data_sosialisasi,err := models.GetSosialisasiMobile(params)	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile models.GetSosialisasiMobile ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}
	data_uk,err := models.GetUkMobile(params)
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile models.GetUkMobile ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	
	data_uk_detail,err := models.GetUkDetailMobile(params)
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile models.GetUkDetailMobile ---> " + err.Error())			
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}
	data_uk_client_data,err := models.GetUkClientDataMobile(params)
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile models.GetUkClientDataMobile ---> " + err.Error())					
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	

	data_pp_kelompok,err := models.GetPpKelompokMobile(params)
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile models.GetPpKelompokMobile ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}			
	data_pp_2_kelompok,err := models.GetPp2KelompokMobile(params)
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile models.GetPp2KelompokMobile ---> " + err.Error())			
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	
	data_pp_3_kelompok,err := models.GetPp3KelompokMobile(params)
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile models.GetPp3KelompokMobile ---> " + err.Error())					
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	
	data_persetujuan_pembiayaan_kelompok,err := models.GetPersetujuanPembiayaanKelompokMobile(params)
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile models.GetPersetujuanPembiayaanKelompokMobile ---> " + err.Error())							
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}				
	data_persetujuan_pembiayaan_client_kelompok,err := models.GetPersetujuanPembiayaanKelompokClientMobile(params)
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile models.GetPersetujuanPembiayaanKelompokClientMobile ---> " + err.Error())									
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	
	
	
	data_prospek_lama,err := models.GetSosialisasiMobileProspekLama(params)
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile models.GetPersetujuanPembiayaanKelompokClientMobile ---> " + err.Error())									
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}		

	data := models.GET_DATA_SOSIALISASI{ 
		Sosialisai: data_sosialisasi,
		Uk: data_uk,
		Uk_detail: data_uk_detail,
		Uk_client_data : data_uk_client_data,
		Pp_kelompok : data_pp_kelompok,
		Pp_2_kelompok : data_pp_2_kelompok,
		Pp_3_kelompok : data_pp_3_kelompok,
		Persetujuan_pembiayaan_kelompok : data_persetujuan_pembiayaan_kelompok,
		Persetujuan_pembiayaan_client_kelompok : data_persetujuan_pembiayaan_client_kelompok,
		Sosialisasi_prospek_lama : data_prospek_lama,
	}

	jsondata, err := json.Marshal(&data)
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Sosialisasi_Mobile json.Marshal ---> " + err.Error())
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = global.APIGetResponse{Code: 403, Message: err.Error()}
		c.ServeJSON()
	}		


	elapsed_Get_Sosialisasi_Mobile := time.Since(start_Get_Sosialisasi_Mobile)
	global.Logging("INFO","controllers.Get_Sosialisasi_Mobile process_time "+elapsed_Get_Sosialisasi_Mobile.String()+" ---> ")

	c.Data["json"] = global.APIGetResponse{Code: 200, Message: "Data Berhasil dimuat", Data: jsondata}
	c.ServeJSON()	

}



// @Title MASTER_DATA
// @Description master_data
// @Param	Authorization header string  false "Authorization Token"
// @Param	OurBranchID path string  true "OurBranchID"
// @Param	CreatedBy path string  false "CreatedBy"
// @Param	Search path string  false "Pencarian"
// @Param	PAGE path string  true "Halaman berapa"
// @Param	LIMIT path string  true "Jumlah baris per Halaman"
// @Success 200 {object} models.SP_GET_LIST_OF_CLIENT
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /GetListClient/:OurBranchID/:CreatedBy/:Search/:PAGE/:LIMIT [get]
func (c *MasterDataController) GetListClient() {
	start_GetListClient := time.Now()

	global.Logging("INFO_HIT"," ================ v1/inisiasi/GetListClient ================ ")	

	var params = make(map[string]string)
	params["OurBranchID"] = c.Ctx.Input.Param(":OurBranchID")
	params["CreatedBy"] = c.Ctx.Input.Param(":CreatedBy")	
	params["Search"] = c.Ctx.Input.Param(":Search")	
	params["PAGE"] = c.Ctx.Input.Param(":PAGE")	
	params["LIMIT"] = c.Ctx.Input.Param(":LIMIT")	


	data,err := models.GetListOfClient(params)
	if err != nil {		
		global.Logging("ERROR","controllers.GetListClient models.GetListOfClient ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}		

	elapsed_GetListClient := time.Since(start_GetListClient)
	global.Logging("INFO","controllers.GetListClient process_time "+elapsed_GetListClient.String()+" ---> ")


	c.Data["json"] = data
	c.ServeJSON()	
}



// @Title MASTER_DATA
// @Description master_data
// @Param	Authorization header string  false "Authorization Token"
// @Param	OurBranchID path string  true "OurBranchID"
// @Param	CreatedBy path string  false "CreatedBy"
// @Param	Search path string  false "Pencarian"
// @Param	PAGE path string  true "Halaman berapa"
// @Param	LIMIT path string  true "Jumlah baris per Halaman"
// @Success 200 {object} models.SP_GET_LIST_OF_CLIENT_BRNET
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /GetListClientBRNET/:OurBranchID/:CreatedBy/:Search/:PAGE/:LIMIT [get]
func (c *MasterDataController) GetListClientBRNET() {
	start_GetListClientBRNET := time.Now()

	global.Logging("INFO_HIT"," ================ v1/inisiasi/GetListClientBRNET ================ ")	

	var params = make(map[string]string)
	params["OurBranchID"] = c.Ctx.Input.Param(":OurBranchID")
	params["CreatedBy"] = c.Ctx.Input.Param(":CreatedBy")	
	params["Search"] = c.Ctx.Input.Param(":Search")	
	params["PAGE"] = c.Ctx.Input.Param(":PAGE")	
	params["LIMIT"] = c.Ctx.Input.Param(":LIMIT")	


	data,err := models.GetListOfClientBRNET(params)
	if err != nil {		
		global.Logging("ERROR","controllers.GetListClientBRNET models.GetListOfClientBRNET ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}		

	elapsed_GetListClientBRNET := time.Since(start_GetListClientBRNET)
	global.Logging("INFO","controllers.GetListClientBRNET process_time "+elapsed_GetListClientBRNET.String()+" ---> ")	

	c.Data["json"] = data
	c.ServeJSON()	
}





// @Title Scoring
// @Description Status Layak
// @Param	Authorization header string  false "Authorization Token"
// @Param	prospek_id path string  true "prospek_id"
// @Success 200 {object} models.SP_GET_SCORING_STATUS
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /ScoringStatus/:prospek_id [get]
func (c *MasterDataController) ScoringStatus() {
	start_ScoringStatus := time.Now()

	global.Logging("INFO_HIT"," ================ v1/inisiasi/ScoringStatus ================ ")	

	var params string
	params = c.Ctx.Input.Param(":prospek_id")

	data,err := models.GetScoringStatus(params)
	if err != nil {		
		global.Logging("ERROR","controllers.ScoringStatus models.GetScoringStatus ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}		

	elapsed_ScoringStatus := time.Since(start_ScoringStatus)
	global.Logging("INFO","controllers.ScoringStatus process_time "+elapsed_ScoringStatus.String()+" ---> ")	

	c.Data["json"] = data
	c.ServeJSON()
}



// @Title MASTER_DATA
// @Description master_data
// @Param	Authorization header string  false "Authorization Token"
// @Param	OurBranchID path string  true "OurBranchID"
// @Param	CreatedBy path string  false "CreatedBy"
// @Success 200 {object} models.GET_DATA_PENCAIRAN
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /GetPencairanMobile/:OurBranchID/:CreatedBy [get]
func (c *MasterDataController) Get_Pencairan_Mobile() {
	start_Get_Pencairan_Mobile := time.Now()

	global.Logging("INFO_HIT"," ================ v1/inisiasi/GetPencairanMobile ================ ")	

	var params = make(map[string]string)
	params["OurBranchID"] = c.Ctx.Input.Param(":OurBranchID")
	params["CreatedBy"] = c.Ctx.Input.Param(":CreatedBy")	

	data_pencairan,err := models.GetPencairanMobile(params)	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Pencairan_Mobile models.GetPencairanMobile ---> " + err.Error())				
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}
	data_pencairan_nasabah,err := models.GetPencairanNasabahMobile(params)	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Pencairan_Mobile models.GetPencairanNasabahMobile ---> " + err.Error())						
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	

	data := models.GET_DATA_PENCAIRAN{ 
		Pencairan: data_pencairan,	
		Pencairan_nasabah: data_pencairan_nasabah,
	}

	jsondata, err := json.Marshal(&data)
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Pencairan_Mobile json.Marshal ---> " + err.Error())
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = global.APIGetResponse{Code: 403, Message: err.Error()}
		c.ServeJSON()
	}		

	elapsed_Get_Pencairan_Mobile := time.Since(start_Get_Pencairan_Mobile)
	global.Logging("INFO","controllers.Get_Pencairan_Mobile process_time "+elapsed_Get_Pencairan_Mobile.String()+" ---> ")	

	c.Data["json"] = global.APIGetResponse{Code: 200, Message: "Data Berhasil dimuat", Data: jsondata}
	c.ServeJSON()		
}




// @Title MASTER_DATA
// @Description master_data
// @Param	Authorization header string  false "Authorization Token"
// @Param	OurBranchID path string  true "OurBranchID"
// @Success 200 {object} models.GetMasterDataResponse
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /GetMasterData/:OurBranchID [get]
func (c *MasterDataController) Get_Master_Data() {

	global.Logging("INFO_HIT"," ================ v1/inisiasi/GetMasterData ================ ")		

	start_Get_Master_Data := time.Now()

	OurBranchID := c.Ctx.Input.Param(":OurBranchID")

	data_absent,err := models.GetMasterAbsent()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterAbsent ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_religion,err := models.GetMasterReligion()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterReligion ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	// data_branch,err := models.GetMasterBranch()	
	// if err != nil {		
	// 	global.Logging("ERROR models.GetMasterReligion ---> " + err.Error())		
	// 	c.Ctx.Output.SetStatus(500)
	// 	c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
	// 	c.ServeJSON()
	// }
	
	data_living_type,err := models.GetMasterLivingType()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterLivingType ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_identity_type,err := models.GetMasterIdentityType()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterIdentityType ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_partner_job,err := models.GetMasterPartnerJob()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterPartnerJob ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_dwelling_condition,err := models.GetMasterDwellingCondition()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterDwellingCondition ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_residence_location,err := models.GetMasterResidenceLocation()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterResidenceLocation ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_pembiayaan_lain,err := models.GetMasterPembiayaanLain()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterPembiayaanLain ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_education,err := models.GetMasterEducation()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterEducation ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_product,err := models.GetMasterProduct()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterEducation ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_economic_sector,err := models.GetMasterEconomicSector()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterEconomicSector ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_relation_status,err := models.GetMasterRelationStatus()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterRelationStatus ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_marriage_status,err := models.GetMasterMarriageStatus()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterMarriageStatus ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_home_status,err := models.GetMasterHomeStatus()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterHomeStatus ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_referral,err := models.GetMasterReferral()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterReferral ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data_trans_fund,err := models.GetMasterTransFund()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterTransFund ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	

	data_jenis_pembiayaan,err := models.GetMasterJenisPembiayaan()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterJenisPembiayaan ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	

	data_sub_jenis_pembiayaan,err := models.GetMasterSubJenisPembiayaan()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterSubJenisPembiayaan ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	

	data_tujuan_pembiayaan,err := models.GetMasterTujuanPembiayaan()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterTujuanPembiayaan ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	

	data_kategori_tujuan_pembiayaan,err := models.GetMasterKategoriTujuanPembiayaan()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterKategoriTujuanPembiayaan ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	

	data_frekuensi,err := models.GetMasterFrekuensi()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterFrekuensi ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	

	data_wilayah_mobile,err := models.GetMasterWilayahMobile(OurBranchID)	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterWilayahMobile ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}		

	data_available_sub_group,err := models.GetMasterAvailableSubGroupBRNET(OurBranchID)	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterAvailableSubGroupBRNET ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	
	
	
	data_group_product,err := models.GetMasterGroupProduct()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetMasterGroupProduct ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}			

	data := models.GetMasterDataResponse{
		ResponseCode:        "200",
		ResponseDescription: "Data Berhasil dimuat",
		Status:              "sukses",
		Data: &models.MasterData{
			Absent:            			data_absent,
			Religion:          			data_religion,
			// Branch:						data_branch,
			LivingType:        			data_living_type,
			IdentityType:      			data_identity_type,
			PartnerJob:        			data_partner_job,
			DwellingCondition: 			data_dwelling_condition,
			ResidenceLocation: 			data_residence_location,
			PembiayaanLain:				data_pembiayaan_lain,
			Product:					data_product,
			Education:					data_education,
			EconomicSector:				data_economic_sector,
			RelationStatus:				data_relation_status,
			MarriageStatus:				data_marriage_status,
			HomeStatus:					data_home_status,
			Referral:					data_referral,
			TransFund:					data_trans_fund,
			JenisPembiayaan:			data_jenis_pembiayaan,
			SubJenisPembiayaan: 		data_sub_jenis_pembiayaan,
			TujuanPembiayaan:			data_tujuan_pembiayaan,
			KategoriTujuanPembiayaan:	data_kategori_tujuan_pembiayaan,
			Frekuensi:					data_frekuensi,
			WilayahMobile:				data_wilayah_mobile,
			AvailableSubGroup:			data_available_sub_group,
			GroupProduct:			    data_group_product,
			SetUKtimeOut:				beego.AppConfig.String("SetUKtimeOut"),
		},
	}

	elapsed_Get_Master_Data := time.Since(start_Get_Master_Data)
	global.Logging("INFO","controllers.Get_Master_Data process_time "+elapsed_Get_Master_Data.String()+" ---> ")

	c.Data["json"] = data
	c.ServeJSON()	
}


// @Title SIKLUS NASABAH
// @Description GET SIKLUS NASABAH
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body models.PARAM_GET_SIKLUS_NASABAH_BRNET true "parameter get data"
// @Success 200 {object} []models.SP_GET_SIKLUS_NASABAH_BRNET
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /GetSiklusNasabahBrnet [post]
func (c *MasterDataController) Get_Siklus_Nasabah_Brnet() {
	var params models.PARAM_GET_SIKLUS_NASABAH_BRNET

	global.Logging("INFO_HIT"," ================ v1/inisiasi/GetSiklusNasabahBrnet ================ ")			

	start_Get_Siklus_Nasabah_Brnet := time.Now()

    err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	if err != nil {
		global.Logging("ERROR","controllers.Get_Siklus_Nasabah_Brnet json.Unmarshal ---> " + err.Error())
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}	


	data,err := models.GetSiklusNasabahBRNET(params)	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Siklus_Nasabah_Brnet models.GetSiklusNasabahBRNET ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	


	jsondata, err := json.Marshal(&data)
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Siklus_Nasabah_Brnet json.Marshal ---> " + err.Error())
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = global.APIGetResponse{Code: 403, Message: err.Error()}
		c.ServeJSON()
	}		

	elapsed_Get_Siklus_Nasabah_Brnet := time.Since(start_Get_Siklus_Nasabah_Brnet)
	global.Logging("INFO","controllers.Get_Siklus_Nasabah_Brnet process_time "+elapsed_Get_Siklus_Nasabah_Brnet.String()+" ---> ")

	c.Data["json"] = global.APIGetResponse{Code: 200, Message: "Data Berhasil dimuat", Data: jsondata}
	c.ServeJSON()		
}




// @Title Get Semua WIlayah PKM
// @Description Semua wilayah
// @Param	Authorization header string  false "Authorization Token"
// @Success 200 {object} []models.MasterWilayahMobile
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /GetAllWilayah [get]
func (c *MasterDataController) Get_All_Wilayah() {

	data_wilayah_mobile,err := models.GetAllMasterWilayahMobile()	
	if err != nil {		
		global.Logging("ERROR","controllers.Get_Master_Data models.GetAllMasterWilayahMobile ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	

	c.Data["json"] = data_wilayah_mobile
	c.ServeJSON()		
}