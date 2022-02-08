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

// Post Inisiasi PKM Mobile
type PostInisiasiController struct {
	beego.Controller
}

// @Title PKM INISIASI
// @Param	Authorization header string  false "Authorization Token"
// @Description PKM PROSPEK DAN UK
// @Param   data body models.SP_ADD_PROSPEK true "simpan data"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql / validasi"}
// @router /post_prospek [post]
func (c *PostInisiasiController) PostProspek() {	
	var data models.SP_ADD_PROSPEK		

	global.Logging("INFO_HIT"," ================ v1/post_inisiasi/post_prospek ================ ")	

    err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {		
		global.Logging("ERROR","controllers.PostProspek json.Unmarshal ---> " + err.Error())
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}	

	err = models.POST_PROSPEK(data)
	if err != nil {
		global.Logging("ERROR","controllers.PostProspek ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}
	
	c.Data["json"] = global.APIResponse{Code: 200, Message: "Data berhasil disimpan"}
    c.ServeJSON()
}

// @Title PKM INISIASI
// @Description PKM PROSPEK DAN UK
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body models.SP_ADD_PROSPEK_UK true "simpan data"
// @Success 200 {object} global.APIGetResponse {"code": 200,"message": "Data berhasil disimpan", "data":"prospek id"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql / validasi"}
// @router /post_prospek_uk [post]
func (c *PostInisiasiController) PostProspekUK() {	
	var data models.SP_ADD_PROSPEK_UK

	global.Logging("INFO_HIT"," ================ v1/post_inisiasi/post_prospek_uk ================ ")			

	start_PostProspekUK := time.Now()

	// global.Logging("PAYLOAD","INFO controllers.PostProspekUK RequestBody ---> " +string(c.Ctx.Input.RequestBody))		

    err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {		
		global.Logging("ERROR","controllers.PostProspekUK json.Unmarshal ---> " + err.Error())				
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIGetResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}	

	global.Logging("PAYLOAD",`INFO controllers.PostProspekUK RequestBody --->
	ID_Prospek=`+data.ID_Prospek+`
	,OurBranchID=`+data.OurBranchID+`
	,Sumber=`+data.Sumber+`
	,NoHP=`+data.NoHP+`
	,Sisipan=`+data.Sisipan+`
	,LokasiSos=`+data.LokasiSos+`
	,LokasiUK=`+data.LokasiUK+`
	,CreatedBy=`+data.CreatedBy+`
	,CreatedNIP=`+data.CreatedNIP+`
	,JenisIdentitas=`+data.JenisIdentitas+`
	,NomorIdentitas=`+data.NomorIdentitas+`
	,IsSesuaiDukcapil=`+data.IsSesuaiDukcapil+`
	,NamaLengkap=`+data.NamaLengkap+`
	,TempatLahir=`+data.TempatLahir+`
	,TanggalLahir=`+data.TanggalLahir+`
	,IDAgama=`+data.IDAgama+`
	,Status_Perkawinan=`+data.Status_Perkawinan+`
	,Alamat=`+data.Alamat+`
	,AlamatDomisili=`+data.AlamatDomisili+`
	,Latitude=`+data.Latitude+`
	,Longitude=`+data.Longitude+`
	,Provinsi=`+data.Provinsi+`
	,Kabupaten=`+data.Kabupaten+`
	,Kecamatan=`+data.Kecamatan+`
	,Kelurahan=`+data.Kelurahan+`
	,NoKK=`+data.NoKK+`
	,NamaAyah=`+data.NamaAyah+`
	,NamaGadisIBU=`+data.NamaGadisIBU+`
	,JmlhAnak=`+data.JmlhAnak+`
	,JmlhTanggungan=`+data.JmlhTanggungan+`
	,JmlTenagaKerja=`+data.JmlTenagaKerja+`
	,StatusRumah=`+data.StatusRumah+`
	,LamaTinggal=`+data.LamaTinggal+`
	,IsAdaRekening=`+data.IsAdaRekening+`
	,NamaBank=`+data.NamaBank+`
	,NomorRekening=`+data.NomorRekening+`
	,NamaPemilikRekening=`+data.NamaPemilikRekening+`
	,IsPernyataanDibaca=`+data.IsPernyataanDibaca+`
	,NamaSuami=`+data.NamaSuami+`
	,PekerjaanSuami=`+data.PekerjaanSuami+`
	,IsSuamiDitempat=`+data.IsSuamiDitempat+`
	,StatusPenjamin=`+data.StatusPenjamin+`
	,NamaPenjamin=`+data.NamaPenjamin+`
	,LuasBangunan=`+data.LuasBangunan+`
	,KondisiBangunan=`+data.KondisiBangunan+`
	,JenisAtap=`+data.JenisAtap+`
	,Dinding=`+data.Dinding+`
	,Lantai=`+data.Lantai+`
	,Is_AdaAksesAirBersih=`+data.Is_AdaAksesAirBersih+`
	,Is_AdaAdaToiletPribadi=`+data.Is_AdaAdaToiletPribadi+`
	,Siklus=`+data.Siklus+`
	,JenisPembiayaan=`+data.JenisPembiayaan+`
	,IDProduk=`+data.IDProduk+`
	,NamaProduk=`+data.NamaProduk+`
	,IDProdukPembiayaan=`+data.IDProdukPembiayaan+`
	,ProdukPembiayaan=`+data.ProdukPembiayaan+`
	,JumlahPinjaman=`+data.JumlahPinjaman+`
	,TermPembiayaan=`+data.TermPembiayaan+`
	,KategoriTujuanPembiayaan=`+data.KategoriTujuanPembiayaan+`
	,TujuanPembiayaan=`+data.TujuanPembiayaan+`
	,TypePencairan=`+data.TypePencairan+`
	,FrekuensiPembiayaan=`+data.FrekuensiPembiayaan+`
	,NoRekening=`+data.NoRekening+`
	,PemilikRekening=`+data.PemilikRekening+`
	,ID_SektorEkonomi=`+data.ID_SektorEkonomi+`
	,ID_SubSektorEkonomi=`+data.ID_SubSektorEkonomi+`
	,Jenis_Usaha=`+data.Jenis_Usaha+`
	,PendapatanKotor_perHari=`+data.PendapatanKotor_perHari+`
	,PengeluaranKel_perHari=`+data.PengeluaranKel_perHari+`
	,PendapatanBersih_perHari=`+data.PendapatanBersih_perHari+`
	,JmlHariUsaha_perBulan=`+data.JmlHariUsaha_perBulan+`
	,PendapatanBersih_perBulan=`+data.PendapatanBersih_perBulan+`
	,PendapatanBersih_perMinggu=`+data.PendapatanBersih_perMinggu+`
	,Is_AdaPembiayaanLain=`+data.Is_AdaPembiayaanLain+`
	,Nama_Pembiayaan_Lembaga_Lain=`+data.Nama_Pembiayaan_Lembaga_Lain+`
	,Kemampuan_Angsuran=`+data.Kemampuan_Angsuran+`
	,Berdasarkan_Tingkat_Pendapatan=`+data.Berdasarkan_Tingkat_Pendapatan+`
	,Berdasarkan_Kemampuan_Angsuran=`+data.Berdasarkan_Kemampuan_Angsuran+`
	,Berdasarkan_Lembaga_Lain=`+data.Berdasarkan_Lembaga_Lain+`
	,PendapatanKotor_perHari_Suami=`+data.PendapatanKotor_perHari_Suami+`
	,PengeluaranKel_perHari_Suami=`+data.PengeluaranKel_perHari_Suami+`
	,PendapatanBersih_perHari_Suami=`+data.PendapatanBersih_perHari_Suami+`
	,JmlHariUsaha_perBulan_Suami=`+data.JmlHariUsaha_perBulan_Suami+`
	,PendapatanBersih_perBulan_Suami=`+data.PendapatanBersih_perBulan_Suami+`
	,PendapatanBersih_perMinggu_Suami=`+data.PendapatanBersih_perMinggu_Suami+`
	,Nama_TTD_Nasabah=`+data.Nama_TTD_Nasabah+`
	,Nama_TTD_Penjamin=`+data.Nama_TTD_Penjamin+`
	,Nama_TTD_KSK=`+data.Nama_TTD_KSK+`
	,Nama_TTD_KK=`+data.Nama_TTD_KK+`
	,Nama_TTD_AO=`+data.Nama_TTD_AO+`
	,Kelompok_ID=`+data.Kelompok_ID+`
	,Nama_Kelompok=`+data.Nama_Kelompok+`
	,Sub_Kelompok=`+data.Sub_Kelompok+`
	,ClientID=`+data.ClientID+`
	,Kehadiran_pkm=`+data.Kehadiran_pkm+`
	,Angsuran_Pada_Saat_PKM=`+data.Angsuran_Pada_Saat_PKM+`	
	`)

	response,err := models.POST_PROSPEK_UK(data)
	if err != nil {
		global.Logging("ERROR","controllers.PostProspekUK models.POST_PROSPEK_UK ---> " + err.Error())						
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	jsondata, err := json.Marshal(&response)
	if err != nil {		
		global.Logging("ERROR","controllers.PostProspekUK json.Marshal ---> " + err.Error())								
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = global.APIGetResponse{Code: 403, Message: err.Error()}
		c.ServeJSON()
	}		


	err = models.DELETE_RKH_ID_PROSPEK_TERPAKAI(data.ID_Prospek) //modelnya di Other
	if err != nil {		
		global.Logging("ERROR","controllers.PostProspekUK models.DELETE_RKH_ID_PROSPEK_TERPAKAI ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}		

	elapsed_PostProspekUK := time.Since(start_PostProspekUK)
	global.Logging("INFO","controllers.PostProspekUK process_time "+elapsed_PostProspekUK.String()+" ---> ")
	
	c.Data["json"] = global.APIGetResponse{Code: 200, Message: "Data berhasil disimpan",Data: jsondata}
    c.ServeJSON()
}

// @Title PKM INISIASI
// @Description PKM PROSPEK LAMA
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body []models.SP_ADD_PERSETUJUAN_KELOMPOK true "simpan data"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /post_prospek_lama [post]
func (c *PostInisiasiController) PostProspekLama() {	
	var data []models.SP_ADD_PERSETUJUAN_KELOMPOK

	global.Logging("INFO_HIT"," ================ v1/post_inisiasi/post_prospek_lama ================ ")	

	// global.Logging("PAYLOAD","controllers.PostProspekLama RequestBody ---> " +string(c.Ctx.Input.RequestBody))			
	
	start_PostProspekLama := time.Now()

    err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {
		global.Logging("ERROR","controllers.PostProspekLama  json.Unmarshal ---> " + err.Error())										
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}	

	for _, params := range data {	
	global.Logging("PAYLOAD",`controllers.PostProspekLama RequestBody ---> 
	ID_Prospek=`+params.ID_Prospek+`
	,No_Identitas=`+params.No_Identitas+`
	,ID_Kelompok=`+params.ID_Kelompok+`
	,Tahap_Pembiayaan=`+params.Tahap_Pembiayaan+`
	,Plafon_Diajukan=`+params.Plafon_Diajukan+`
	,Tenor=`+params.Tenor+`
	,Status_Tempat_Tinggal=`+params.Status_Tempat_Tinggal+`
	,Is_Perkawinan_Berubah=`+params.Is_Perkawinan_Berubah+`
	,Is_Tanggungan_Berubah=`+params.Is_Tanggungan_Berubah+`
	,Keterangan_Tanggungan=`+params.Keterangan_Tanggungan+`
	,Jml_Kehadiran_PKM=`+params.Jml_Kehadiran_PKM+`
	,Jml_Pembayaran=`+params.Jml_Pembayaran+`
	,Is_Usaha_Berubah=`+params.Is_Usaha_Berubah+`
	,Keterangan_Usaha=`+params.Keterangan_Usaha+`
	,Lokasi_Persetujuan=`+params.Lokasi_Persetujuan+`
	,Tanggal_Persetujuan=`+params.Tanggal_Persetujuan)
	}
	
	err = models.POST_PROSPEK_LAMA(data)
	if err != nil {
		global.Logging("ERROR","controllers.PostProspekLama models.POST_PROSPEK_LAMA ---> " + err.Error())												
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	elapsed_PostProspekLama := time.Since(start_PostProspekLama)
	global.Logging("INFO","controllers.PostProspekLama process_time "+elapsed_PostProspekLama.String()+" ---> ")

	c.Data["json"] = global.APIResponse{Code: 200, Message: "Data berhasil disimpan"}
    c.ServeJSON()
}

// @Title PKM INISIASI
// @Description PKM DATA KELOMPOK
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body models.SP_ADD_DATA_KELOMPOK true "simpan data"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /post_data_kelompok [post]
func (c *PostInisiasiController) PostDataKelompok() {	
	var data models.SP_ADD_DATA_KELOMPOK

	global.Logging("INFO_HIT"," ================ v1/post_inisiasi/post_data_kelompok ================ ")	

	global.Logging("PAYLOAD","controllers.PostDataKelompok RequestBody ---> " +string(c.Ctx.Input.RequestBody))		

	start_PostDataKelompok := time.Now()

    err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {
		global.Logging("ERROR","controllers.PostDataKelompok json.Unmarshal ---> " + err.Error())		
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}	

	
	err = models.POST_DATA_KELOMPOK(data)
	if err != nil {
		global.Logging("ERROR","controllers.PostDataKelompok models.POST_DATA_KELOMPOK ---> " + err.Error())				
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	

	elapsed_PostDataKelompok := time.Since(start_PostDataKelompok)
	global.Logging("INFO","controllers.PostDataKelompok process_time "+elapsed_PostDataKelompok.String()+" ---> ")	

	c.Data["json"] = global.APIResponse{Code: 200, Message: "Data berhasil disimpan"}
    c.ServeJSON()
}

// @Title PKM INISIASI
// @Description PKM PP
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body models.SP_ADD_INISIASI_PP true "simpan data"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /post_pp [post]
func (c *PostInisiasiController) PostPP() {	
	var data models.SP_ADD_INISIASI_PP

	global.Logging("INFO_HIT"," ================ v1/post_inisiasi/post_pp ================ ")	

	global.Logging("PAYLOAD","controllers.PostPP RequestBody ---> " +string(c.Ctx.Input.RequestBody))		
	
	start_PostPP := time.Now()	
	
    err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {
		global.Logging("ERROR","controllers.PostPP json.Unmarshal ---> " + err.Error())						
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}							
	
	err = models.POST_PP(data)
	if err != nil {
		global.Logging("ERROR","controllers.PostPP models.POST_PP ---> " + err.Error())								
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}
	
	elapsed_PostPP := time.Since(start_PostPP)
	global.Logging("INFO","controllers.PostPP process_time "+elapsed_PostPP.String()+" ---> ")

	c.Data["json"] = global.APIResponse{Code: 200, Message: "Data berhasil disimpan"}
    c.ServeJSON()
}

// @Title PKM INISIASI
// @Description PKM PP
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body models.SP_SET_VERIF_STATUS true "simpan data"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /post_verif_status [post]
func (c *PostInisiasiController) PostVerifStatus() {	
	var data models.SP_SET_VERIF_STATUS

	global.Logging("INFO_HIT"," ================ v1/post_inisiasi/post_verif_status ================ ")	
	
	// global.Logging("PAYLOAD","controllers.PostVerifStatus RequestBody ---> " +string(c.Ctx.Input.RequestBody))	

	start_PostVerifStatus := time.Now()	

    err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {
		global.Logging("ERROR","controllers.PostVerifStatus json.Unmarshal ---> " + err.Error())										
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}	

	global.Logging("PAYLOAD",`controllers.PostVerifStatus RequestBody ---> 
	PostStatus=`+data.PostStatus+`
	,Reason=`+data.Reason+`
	,ID_PROSPEK=`+data.ID_PROSPEK+`
	,VerifBy=`+data.VerifBy+`
	,URL_FP4=`+data.URL_FP4)
	
	err = models.POST_VERIF_STATUS(data)
	if err != nil {
		global.Logging("ERROR","controllers.PostVerifStatus models.POST_VERIF_STATUS ---> " + err.Error())			
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	if (data.PostStatus=="1") {
		err = models.POST_SET_DATA_FROM_INISIASI_MOBILE_BRNET(data)
		if err != nil {
			global.Logging("ERROR","controllers.PostVerifStatus models.POST_SET_DATA_FROM_INISIASI_MOBILE_BRNET ---> " + err.Error())			
			c.Ctx.Output.SetStatus(500)	
			c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
			c.ServeJSON()
		}	
	}

	err = models.DELETE_RKH_ID_PROSPEK_TERPAKAI(data.ID_PROSPEK) //modelnya di Other
	if err != nil {		
		global.Logging("ERROR","controllers.PostVerifStatus models.DELETE_RKH_ID_PROSPEK_TERPAKAI ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}	


	elapsed_PostVerifStatus := time.Since(start_PostVerifStatus)
	global.Logging("INFO","controllers.PostVerifStatus process_time "+elapsed_PostVerifStatus.String()+" ---> ")

	c.Data["json"] = global.APIResponse{Code: 200, Message: "Data berhasil disimpan"}
    c.ServeJSON()
}

// @Title INISIASI PERSETUJUAN PEMBIAYAAN
// @Description SP_ADD_PERSETUJUAN_PEMBIAYAAN
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body models.SP_ADD_PERSETUJUAN_PEMBIAYAAN true "simpan data"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /persetujuan_pembiayaan [post]
func (c *PostInisiasiController) PersetujuanPembiayaan() {	
	var data models.SP_ADD_PERSETUJUAN_PEMBIAYAAN	

	global.Logging("INFO_HIT"," ================ v1/post_inisiasi/persetujuan_pembiayaan ================ ")	

	start_PersetujuanPembiayaan := time.Now()	

	// global.Logging("PAYLOAD","controllers.PersetujuanPembiayaan RequestBody ---> " +string(c.Ctx.Input.RequestBody))

    err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {
		global.Logging("ERROR","controllers.PersetujuanPembiayaan json.Unmarshal ---> " + err.Error())										
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}		
	
	global.Logging("PAYLOAD",`controllers.PersetujuanPembiayaan RequestBody --->
	ID_Prospek=`+data.ID_Prospek+`
	Keterangan_Pembelian=`+data.Keterangan_Pembelian)

	err = models.POST_PERSETUJUAN_PEMBIAYAAN(data)
	if err != nil {
		global.Logging("ERROR","controllers.PersetujuanPembiayaan models.POST_PERSETUJUAN_PEMBIAYAAN ---> " + err.Error())			
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}


	err = models.DELETE_RKH_ID_PROSPEK_TERPAKAI(data.ID_Prospek) //modelnya di Other
	if err != nil {		
		global.Logging("ERROR","controllers.PersetujuanPembiayaan models.DELETE_RKH_ID_PROSPEK_TERPAKAI ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}		

	elapsed_PersetujuanPembiayaan := time.Since(start_PersetujuanPembiayaan)
	global.Logging("INFO","controllers.PersetujuanPembiayaan process_time "+elapsed_PersetujuanPembiayaan.String()+" ---> ")	
	
	c.Data["json"] = global.APIResponse{Code: 200, Message: "Data berhasil disimpan"}
    c.ServeJSON()
}

// @Title PKM INISIASI
// @Description PKM DATA KELOMPOK
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body models.SP_ADD_KK_KSK_DATA true "simpan data"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /post_ketua_subketua [post]
func (c *PostInisiasiController) PostKetuaSubKetua() {	
	var data models.SP_ADD_KK_KSK_DATA

	global.Logging("INFO_HIT"," ================ v1/post_inisiasi/post_ketua_subketua ================ ")	

	global.Logging("PAYLOAD","controllers.PostKetuaSubKetua RequestBody ---> " +string(c.Ctx.Input.RequestBody))

	start_PostKetuaSubKetua := time.Now()	

    err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {
		global.Logging("ERROR","controllers.PostKetuaSubKetua json.Unmarshal ---> " + err.Error())		
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APIResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}	

	
	err = models.POST_KK_KSK_DATA(data)
	if err != nil {
		global.Logging("ERROR","controllers.PostKetuaSubKetua models.POST_KK_KSK_DATA ---> " + err.Error())				
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	elapsed_PostKetuaSubKetua := time.Since(start_PostKetuaSubKetua)
	global.Logging("INFO","controllers.PostKetuaSubKetua process_time "+elapsed_PostKetuaSubKetua.String()+" ---> ")	
	
	c.Data["json"] = global.APIResponse{Code: 200, Message: "Data berhasil disimpan"}
    c.ServeJSON()
}