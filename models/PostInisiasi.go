package models

import (
	"pkmapi/global"
	// "strconv"
	"github.com/astaxie/beego/validation"	
	"time"		
	// "fmt"	
	"net/http"
	"encoding/json"
	"strings"

	// "errors"
	"github.com/astaxie/beego"	
)

type SP_ADD_PROSPEK struct {
	OurBranchID string `json:"OurBranchID"`
	Sumber string `json:"Sumber"`
	NamaCalonNasabah string `json:"NamaCalonNasabah"`
	NoHP string `json:"NoHP"`
	Sisipan string `json:"Sisipan"`
	TanggalSos string `json:"TanggalSos"`
	LokasiSos string `json:"LokasiSos"`
	LanjutUK string `json:"LanjutUK"`
	CreatedBy string `json:"CreatedBy"`
	CreatedNIP string `json:"CreatedNIP"`
	Kelompok_ID string `json:"Kelompok_ID" db:"Kelompok_ID"`
	Nama_Kelompok string `json:"Nama_Kelompok" db:"Nama_Kelompok"`
	Sub_Kelompok string `json:"Sub_Kelompok" db:"Sub_Kelompok"`
	ClientID string `json:"ClientID" db:"ClientID"`
}

func POST_PROSPEK(data SP_ADD_PROSPEK)error {
	db,err := global.Conn()
	if err != nil {
		return err
	}	
	var query string

	err = db.Begin() //Begin Transaction
	if err != nil {return err}		

	query = "EXEC ADD_PROSPEK "+
	"@OurBranchID = '"+data.OurBranchID+"' "+
	",@Sumber = '"+data.Sumber+"' "+
	",@NamaCalonNasabah = '"+data.NamaCalonNasabah+"' "+
	",@NoHP = '"+data.NoHP+"' "+
	",@Sisipan = '"+data.Sisipan+"' "+
	",@TanggalSos = '"+data.TanggalSos+"' "+
	",@LokasiSos = '"+data.LokasiSos+"' "+
	",@LanjutUK = '"+data.LanjutUK+"' "+
	",@CreatedBy = '"+data.CreatedBy+"' "+
	",@CreatedNIP = '"+data.CreatedNIP+"' "+
	",@Kelompok_ID = '"+data.Kelompok_ID+"' "+
	",@Nama_Kelompok = '"+data.Nama_Kelompok+"' "+
	",@Sub_Kelompok = '"+data.Sub_Kelompok+"' "+
	",@ClientID = '"+data.ClientID+"' "
	

	_ , err = db.RawSQL(query).DoWithIterator()
	if err != nil {	
		global.Logging("ERROR","models POST_PROSPEK ---> "+query+" ---> " + err.Error())		
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

type PATH struct {
	FILE string `db:"PATH_FILE"`
}

type SP_ADD_PROSPEK_UK struct {
	ID_Prospek string `json:"ID_Prospek"`

	OurBranchID string `json:"OurBranchID" valid:"Required"`
	Sumber string `json:"Sumber"`
	// NamaCalonNasabah string `json:"NamaCalonNasabah" valid:"Required"` 
	NoHP string `json:"NoHP" valid:"Required"`
	Sisipan string `json:"Sisipan"`
	// TanggalSos string `json:"TanggalSos" valid:"Required"`
	LokasiSos string `json:"LokasiSos" valid:"Required"`
	// LanjutUK string `json:"LanjutUK"`
	LokasiUK string `json:"LokasiUK" valid:"Required`
	CreatedBy string `json:"CreatedBy" valid:"Required"`
	CreatedNIP string `json:"CreatedNIP" valid:"Required"`
	FotoKartuIdentitas string `json:"FotoKartuIdentitas"` 
	JenisIdentitas string `json:"JenisIdentitas" valid:"Numeric"` 
	NomorIdentitas string `json:"NomorIdentitas" valid:"Required"` 
	IsSesuaiDukcapil string `json:"IsSesuaiDukcapil"`
	NamaLengkap string `json:"NamaLengkap" valid:"Required"` 
	TempatLahir string `json:"TempatLahir" valid:"Required"` 
	TanggalLahir string `json:"TanggalLahir" valid:"Required"` 
	// Umur string `json:"Umur" valid:"Required"`
	// JenisKelamin string `json:"JenisKelamin" valid:"Required"`
	IDAgama string `json:"IDAgama" valid:"Required"`
	Status_Perkawinan string `json:"Status_Perkawinan" valid:"Numeric"` 
	Alamat string `json:"Alamat" valid:"Required"` 
	AlamatDomisili string `json:"AlamatDomisili" valid:"Required"` 
	Latitude string `json:"Latitude" valid:"Required"`
	Longitude string `json:"Longitude" valid:"Required"`
	FotoSuketDomisili string `json:"FotoSuketDomisili"` 
	// RT string `json:"RT" valid:"Required"`
	// RW string `json:"RW" valid:"Required"`
	Provinsi string `json:"Provinsi" valid:"Numeric"` 
	Kabupaten string `json:"Kabupaten" valid:"Numeric"` 
	Kecamatan string `json:"Kecamatan" valid:"Numeric"` 
	Kelurahan string `json:"Kelurahan" valid:"Numeric"` 
	// KodePOS string `json:"KodePOS" valid:"Required"`
	// LokasiTinggal string `json:"LokasiTinggal"`
	// DaerahTinggal string `json:"DaerahTinggal"`
	// Pekerjaan string `json:"Pekerjaan" valid:"Required"`
	// TanggalReleaseID string `json:"TanggalReleaseID" valid:"Required"`
	// Pendidikan string `json:"Pendidikan"`
	// FotoProspek string `json:"FotoProspek" `
	FotoKK string `json:"FotoKK" valid:"Required"` 
	NoKK string `json:"NoKK" valid:"Required"` 
	NamaAyah string `json:"NamaAyah" valid:"Required"` 
	NamaGadisIBU string `json:"NamaGadisIBU" valid:"Required"`
	// NoTelp string `json:"NoTelp" valid:"Required"` 
	JmlhAnak string `json:"JmlhAnak" valid:"Numeric"` 
	JmlhTanggungan string `json:"JmlhTanggungan" valid:"Numeric"` 
	JmlTenagaKerja string `json:"JmlTenagaKerja"` 
	// PendidikanAnak string `json:"PendidikanAnak"`
	StatusRumah string `json:"StatusRumah" valid:"Numeric"` 
	LamaTinggal string `json:"LamaTinggal" valid:"Numeric"` 
	IsAdaRekening string `json:"IsAdaRekening" valid:"Numeric"`
	NamaBank string `json:"NamaBank" valid:"Required"`
	NomorRekening string `json:"NomorRekening" valid:"Required"` 
	NamaPemilikRekening string `json:"NamaPemilikRekening" valid:"Required"` 
	IsPernyataanDibaca string `json:"IsPernyataanDibaca" valid:"Required"`
	// Is_DPVerified string `json:"Is_DPVerified"`
	NamaSuami string `json:"NamaSuami"` 
	// TglLahirSuami string `json:"TglLahirSuami" valid:"Required"`
	// TempatLahirSuami string `json:"TempatLahirSuami" valid:"Required"`
	// PendidikanTerkahirSuami string `json:"PendidikanTerkahirSuami"`
	PekerjaanSuami string `json:"PekerjaanSuami"`
	// IDJenisPekerjaanSuami string `json:"IDJenisPekerjaanSuami"`
	// FotoSuami string `json:"FotoSuami" `
	FotoKTPSuami string `json:"FotoKTPSuami"` 
	IsSuamiDitempat string `json:"IsSuamiDitempat" valid:"Numeric"` 
	// IsSuami string `json:"IsSuami"`
	StatusPenjamin string `json:"StatusPenjamin" valid:"Numeric"` 
	NamaPenjamin string `json:"NamaPenjamin" valid:"Required"` 
	// NoTelpPenjamin string `json:"NoTelpPenjamin" valid:"Required"`
	// NoIdentitasPenjamin string `json:"NoIdentitasPenjamin" valid:"Required"`
	// JenisIdentitasPenjamin string `json:"JenisIdentitasPenjamin"`
	// AlamatPenjamin string `json:"AlamatPenjamin" valid:"Required"`
	// FotoPenjamin string `json:"FotoPenjamin" `
	FotoKTPPenjamin string `json:"FotoKTPPenjamin"` 
	LuasBangunan string `json:"LuasBangunan" valid:"Numeric"` 
	KondisiBangunan string `json:"KondisiBangunan" valid:"Numeric"` 
	JenisAtap string `json:"JenisAtap" valid:"Numeric"` 
	Dinding string `json:"Dinding" valid:"Numeric"` 
	Lantai string `json:"Lantai" valid:"Numeric"` 
	// FotoRumah1 string `json:"FotoRumah1"`
	// FotoRumah2 string `json:"FotoRumah2"`
	// Foto_Jenis_Atap string `json:"Foto_Jenis_Atap"`
	// Foto_Dinding string `json:"Foto_Dinding"`
	// Foto_Lantai string `json:"Foto_Lantai"`
	// Nilai_Total string `json:"Nilai_Total"`
	Is_AdaAksesAirBersih string `json:"Is_AdaAksesAirBersih" valid:"Numeric"` 
	Is_AdaAdaToiletPribadi string `json:"Is_AdaAdaToiletPribadi" valid:"Numeric"` 
	
	Siklus string `json:"Siklus"`
	JenisPembiayaan string `json:"JenisPembiayaan" valid:"Numeric"` 
	IDProduk string `json:"IDProduk" db:"IDProduk"`
	NamaProduk string `json:"NamaProduk" valid:"Required"` 
	IDProdukPembiayaan string `json:"IDProdukPembiayaan" db:"IDProdukPembiayaan"`
	ProdukPembiayaan string `json:"ProdukPembiayaan" valid:"Required"` 
	JumlahPinjaman string `json:"JumlahPinjaman" valid:"Numeric"` 
	TermPembiayaan string `json:"TermPembiayaan" valid:"Numeric"` 
	KategoriTujuanPembiayaan string `json:"KategoriTujuanPembiayaan" valid:"Numeric"` 
	TujuanPembiayaan string `json:"TujuanPembiayaan" valid:"Numeric"` 
	TypePencairan string `json:"TypePencairan" valid:"Numeric"` 
	FrekuensiPembiayaan string `json:"FrekuensiPembiayaan" valid:"Numeric"` 
	// IsProductVerified string `json:"IsProductVerified"`
	// RekeningBank string `json:"RekeningBank"`
	// NamaBankPemilik string `json:"NamaBankPemilik" valid:"Required"`
	NoRekening string `json:"NoRekening" valid:"Required"` 
	PemilikRekening string `json:"PemilikRekening" valid:"Required"` 
	ID_SektorEkonomi string `json:"ID_SektorEkonomi" valid:"Numeric"` 
	ID_SubSektorEkonomi string `json:"ID_SubSektorEkonomi" valid:"Numeric"` 
	Jenis_Usaha string `json:"Jenis_Usaha" valid:"Required"` 
	// Lokasi_Usaha_1 string `json:"Lokasi_Usaha_1"`
	// Lokasi_Usaha_2 string `json:"Lokasi_Usaha_2"`
	// Foto_Usaha string `json:"Foto_Usaha" `
	// Is_SEVerified string `json:"Is_SEVerified"`
	PendapatanKotor_perHari string `json:"PendapatanKotor_perHari" valid:"Numeric"` 
	PengeluaranKel_perHari string `json:"PengeluaranKel_perHari" valid:"Numeric"` 
	PendapatanBersih_perHari string `json:"PendapatanBersih_perHari" valid:"Numeric"` 
	JmlHariUsaha_perBulan string `json:"JmlHariUsaha_perBulan" valid:"Numeric"` 
	PendapatanBersih_perBulan string `json:"PendapatanBersih_perBulan" valid:"Numeric"` 
	PendapatanBersih_perMinggu string `json:"PendapatanBersih_perMinggu" valid:"Numeric"`
	Is_AdaPembiayaanLain string `json:"Is_AdaPembiayaanLain" valid:"Numeric"` 
	// Nama_Tipe string `json:"Nama_Tipe"`
	Nama_Pembiayaan_Lembaga_Lain string `json:"Nama_Pembiayaan_Lembaga_Lain" valid:"Required"` 
	// Jumlah_Pendapatan_Bersih string `json:"Jumlah_Pendapatan_Bersih"`
	// Tingkat_Pendapatan_Bersih string `json:"Tingkat_Pendapatan_Bersih"`
	Kemampuan_Angsuran string `json:"Kemampuan_Angsuran" valid:"Numeric"` 
	Berdasarkan_Tingkat_Pendapatan string `json:"Berdasarkan_Tingkat_Pendapatan" valid:"Numeric"`
	Berdasarkan_Kemampuan_Angsuran string `json:"Berdasarkan_Kemampuan_Angsuran" valid:"Numeric"`
	Berdasarkan_Lembaga_Lain string `json:"Berdasarkan_Lembaga_Lain" valid:"Numeric"`
	PendapatanKotor_perHari_Suami string `json:"PendapatanKotor_perHari_Suami" valid:"Numeric"` 
	PengeluaranKel_perHari_Suami string `json:"PengeluaranKel_perHari_Suami" valid:"Numeric"` 
	PendapatanBersih_perHari_Suami string `json:"PendapatanBersih_perHari_Suami" valid:"Numeric"` 
	JmlHariUsaha_perBulan_Suami string `json:"JmlHariUsaha_perBulan_Suami" valid:"Numeric"` 
	PendapatanBersih_perBulan_Suami string `json:"PendapatanBersih_perBulan_Suami" valid:"Numeric` 
	PendapatanBersih_perMinggu_Suami string `json:"PendapatanBersih_perMinggu_Suami" valid:"Numeric"` 
	// Is_TPVerified string `json:"Is_TPVerified"`
	TTD_Nasabah string `json:"TTD_Nasabah" valid:"Required"`
	// TTD_Suami string `json:"TTD_Suami" `
	TTD_Penjamin string `json:"TTD_Penjamin" valid:"Required"`
	TTD_KSK string `json:"TTD_KSK" valid:"Required"`
	TTD_KK string `json:"TTD_KK" valid:"Required"`
	TTD_AO string `json:"TTD_AO" valid:"Required"`
	// TTD_KC_SAO string `json:"TTD_KC_SAO"`

	Nama_TTD_Nasabah string `json:"Nama_TTD_Nasabah" db:"Nama_TTD_Nasabah"`
	Nama_TTD_Penjamin string `json:"Nama_TTD_Penjamin" db:"Nama_TTD_Penjamin"`
	Nama_TTD_KSK string `json:"Nama_TTD_KSK" db:"Nama_TTD_KSK"`
	Nama_TTD_KK string `json:"Nama_TTD_KK" db:"Nama_TTD_KK"`
	Nama_TTD_AO string `json:"Nama_TTD_AO" db:"Nama_TTD_AO"`
	// Nama_TTD_KC_SAO string `json:"Nama_TTD_KC_SAO" db:"Nama_TTD_KC_SAO"`

	Kelompok_ID string `json:"Kelompok_ID" db:"Kelompok_ID"` 
	Nama_Kelompok string `json:"Nama_Kelompok" db:"Nama_Kelompok"`
	Sub_Kelompok string `json:"Sub_Kelompok" db:"Sub_Kelompok"`
	ClientID string `json:"ClientID" db:"ClientID"`

	// Nama_KK string `json:"Nama_KK" db:"Nama_KK"`
	// Nama_KSK string `json:"Nama_KSK" db:"Nama_KSK"`

	Kehadiran_pkm string `json:"kehadiran_pkm" db:"kehadiran_pkm"`
	Angsuran_Pada_Saat_PKM string `json:"Angsuran_Pada_Saat_PKM" db:"Angsuran_Pada_Saat_PKM"`
}

type CHECK_ACCOUNT_BRI struct {
	Data 				DATA_ACCOUNT_BRI `json:"Data"`
	ErrorDescription 	string `json:"errorDescription"`
	ResponseCode 		string `json:"responseCode"`
	ResponseDescription string `json:"responseDescription"`
}

type DATA_ACCOUNT_BRI struct {
	RegistrationStatus 	string `json:"registrationStatus"`
	SourceAccount 		string `json:"sourceAccount"`
	SourceAccountBalace string `json:"sourceAccountBalace"`
	SourceAccountName 	string `json:"sourceAccountName"`
	SourceAccountStatus string `json:"sourceAccountStatus"`
}

type RESPONSE_ADD_PROSPEK_UK struct {
	ID_Prospek 	string `json:"ID_Prospek" db:"ID_Prospek"`
	Status_Kelayakan 	string `json:"Status_Kelayakan" db:"Status_Kelayakan"`
}

func POST_PROSPEK_UK(data SP_ADD_PROSPEK_UK) ([]RESPONSE_ADD_PROSPEK_UK,error){		

	responsedata := make([]RESPONSE_ADD_PROSPEK_UK, 0)

	db,err := global.Conn()
	if err != nil {
		return responsedata,err
	}

	valid := validation.Validation{}
	// path := make([]PATH, 0)
	
	var query string	
	var fileimages = make(map[string]string)

	var Tujuan_Transfer_Check string


	start_cek_briva := time.Now()
	
	if (data.TypePencairan =="1" || data.TypePencairan=="2" ){
		Tujuan_Transfer_Check = "1"
	} else {
		var data_bri CHECK_ACCOUNT_BRI

		client := &http.Client{}
		request, err := http.NewRequest("GET",beego.AppConfig.String("BRIAPI")+"/bri/funding/InfoRek/"+data.NomorRekening, nil)
		if err != nil {
			return responsedata,err
		}		

		response, err := client.Do(request)
		if err != nil {
			return responsedata,err
		}
		defer response.Body.Close()

		err = json.NewDecoder(response.Body).Decode(&data_bri)
		if err != nil {
			return responsedata,err
		}

		// fmt.Println(data_bri.Data.SourceAccountName)

		if (strings.TrimSpace(data.NomorRekening)==strings.TrimSpace(data_bri.Data.SourceAccount) && strings.TrimSpace( strings.ToUpper(data.NamaPemilikRekening) )==strings.TrimSpace( strings.ToUpper(data_bri.Data.SourceAccountName) )){
			Tujuan_Transfer_Check = "1"
		}else{
			Tujuan_Transfer_Check = "0"

			// return responsedata,errors.New("No rekening BRI tidak sesuai ")
		}		
	
	}


	elapsed_cek_briva := time.Since(start_cek_briva)
	global.Logging("INFO","models.POST_PROSPEK_UK process_time cek_briva "+elapsed_cek_briva.String()+" ---> ")

	// fmt.Println(Tujuan_Transfer_Check)

	err = db.Begin() //Begin Transaction
	if err != nil {return responsedata,err}	

	// for _, data := range data_arr {	
		
		// check validasi start		
		v, err := valid.Valid(&data)
		if err != nil {return responsedata,err}	
		if !v {
			for _, err := range valid.Errors {
				// fmt.Println(err.Key, err.Message)
				return responsedata,err
			}
		}
		// check validasi end

		start_images := time.Now()

		//ambil last id prospek sebelum di insert untuk bikin path file images/tanggal/id prospek/
		// err = db.RawSQL("SELECT CONCAT('images/',CAST(GETDATE() as DATE),'/',IDENT_CURRENT('INISIASI_PROSPEK')+1,'/') AS PATH_FILE ").Do(&path)
		// if err != nil {		
		// 	errsql := err
		// 	err = db.Rollback() //Rollback Transaction
		// 	if err != nil {return responsedata,err}
		// 	return responsedata,errsql
		// }			

		fileimages["FotoKartuIdentitas"] = data.FotoKartuIdentitas 
		fileimages["FotoSuketDomisili"] = data.FotoSuketDomisili
		fileimages["FotoKK"] = data.FotoKK
		fileimages["FotoKTPSuami"] = data.FotoKTPSuami
		fileimages["FotoKTPPenjamin"] = data.FotoKTPPenjamin
		fileimages["TTD_Nasabah"] = data.TTD_Nasabah
		fileimages["TTD_Penjamin"] = data.TTD_Penjamin
		fileimages["TTD_KSK"] = data.TTD_KSK
		fileimages["TTD_KK"] = data.TTD_KK
		fileimages["TTD_AO"] = data.TTD_AO

		// fileimages,err = global.MapB64SaveFile(fileimages,path[0].FILE)
		fileimages,err = global.MapB64SaveFile(fileimages,beego.AppConfig.String("pathimages"))
		if err != nil { 			
			errfile := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {return responsedata,err}
			return responsedata,errfile 
		}
	
		elapsed_images := time.Since(start_images)
		global.Logging("INFO","models.POST_PROSPEK_UK process_time images "+elapsed_images.String()+" ---> ")

		query = "EXEC ADD_PROSPEK_UK "+
		" @ID_Prospek_revisi ='"+data.ID_Prospek+"' "+	
		",@OurBranchID = '"+data.OurBranchID+"' "+
		",@Sumber = '"+data.Sumber+"' "+
		",@NamaCalonNasabah = '"+data.NamaLengkap+"' "+
		",@NoHP = '"+data.NoHP+"' "+
		",@Sisipan = '"+data.Sisipan+"' "+
		// ",@TanggalSos = '"+data.TanggalSos+"' "+
		",@LokasiSos = '"+data.LokasiSos+"' "+
		",@LanjutUK = '1' "+
		",@LokasiUK = '"+data.LokasiUK+"' "+
		",@CreatedBy = '"+data.CreatedBy+"' "+
		",@CreatedNIP = '"+data.CreatedNIP+"' "+
		",@FotoKartuIdentitas = '"+fileimages["FotoKartuIdentitas"]+"' "+
		",@JenisIdentitas = '"+data.JenisIdentitas+"' "+
		",@NomorIdentitas = '"+data.NomorIdentitas+"' "+
		",@IsSesuaiDukcapil = '"+data.IsSesuaiDukcapil+"' "+
		",@NamaLengkap = '"+data.NamaLengkap+"' "+
		",@TempatLahir = '"+data.TempatLahir+"' "+
		",@TanggalLahir = '"+data.TanggalLahir+"' "+
		// ",@Umur = '"+data.Umur+"' "+
		",@JenisKelamin = 'F' "+
		",@IDAgama = '"+data.IDAgama+"' "+
		",@Status_Perkawinan = '"+data.Status_Perkawinan+"' "+
		",@Alamat = '"+data.Alamat+"' "+
		",@AlamatDomisili = '"+data.AlamatDomisili+"' "+
		",@Latitude = '"+data.Latitude+"' "+
		",@Longitude = '"+data.Longitude+"' "+
		",@FotoSuketDomisili = '"+fileimages["FotoSuketDomisili"]+"' "+
		",@RT = '0' "+
		",@RW = '0' "+
		",@Provinsi = '"+data.Provinsi+"' "+
		",@Kabupaten = '"+data.Kabupaten+"' "+
		",@Kecamatan = '"+data.Kecamatan+"' "+
		",@Kelurahan = '"+data.Kelurahan+"' "+
		// ",@KodePOS = '"+data.KodePOS+"' "+
		",@LokasiTinggal = '1' "+
		",@DaerahTinggal = '1' "+
		",@Pekerjaan = 'ibu rumah tangga' "+
		// ",@TanggalReleaseID = GETDATE() "+
		",@Pendidikan = '4' "+
		",@FotoProspek = 'gak ada' "+
		",@FotoKK = '"+fileimages["FotoKK"]+"' "+
		",@NoKK = '"+data.NoKK+"' "+
		",@NamaAyah = '"+data.NamaAyah+"' "+
		",@NamaGadisIBU = '"+data.NamaGadisIBU+"' "+
		",@NoTelp = '"+data.NoHP+"' "+
		",@JmlhAnak = '"+data.JmlhAnak+"' "+
		",@JmlhTanggungan = '"+data.JmlhTanggungan+"' "+
		// ",@PendidikanAnak = '"+data.PendidikanAnak+"' "+
		",@PendidikanAnak = '0' "+
		",@StatusRumah = '"+data.StatusRumah+"' "+
		",@LamaTinggal = '"+data.LamaTinggal+"' "+
		",@IsAdaRekening = '"+data.IsAdaRekening+"' "+
		",@NamaBank = '"+data.NamaBank+"' "+
		",@NomorRekening = '"+data.NomorRekening+"' "+
		",@NamaPemilikRekening = '"+data.NamaPemilikRekening+"' "+
		",@IsPernyataanDibaca = '"+data.IsPernyataanDibaca+"' "+
		",@Is_DPVerified = '1' "+
		",@NamaSuami = '"+data.NamaSuami+"' "+
		// ",@TglLahirSuami = '"+data.TglLahirSuami+"' "+
		// ",@TempatLahirSuami = '"+data.TempatLahirSuami+"' "+
		// ",@PendidikanTerkahirSuami = '"+data.PendidikanTerkahirSuami+"' "+
		",@PekerjaanSuami = '"+data.PekerjaanSuami+"' "+
		",@JmlTenagaKerja = '"+data.JmlTenagaKerja+"' "+		
		// ",@IDJenisPekerjaanSuami = '"+data.IDJenisPekerjaanSuami+"' "+
		// ",@FotoSuami = '"+fileimages["FotoSuami"]+"' "+
		",@FotoKTPSuami = '"+fileimages["FotoKTPSuami"]+"' "+
		",@IsDitempat = '"+data.IsSuamiDitempat+"' "+
		// ",@IsSuami = '"+data.IsSuami+"' "+
		",@StatusPenjamin = '"+data.StatusPenjamin+"' "+
		",@NamaPenjamin = '"+data.NamaPenjamin+"' "+
		// ",@NoTelpPenjamin = '"+data.NoTelpPenjamin+"' "+
		// ",@NoIdentitasPenjamin = '"+data.NoIdentitasPenjamin+"' "+
		// ",@JenisIdentitasPenjamin = '"+data.JenisIdentitasPenjamin+"' "+
		// ",@AlamatPenjamin = '"+data.AlamatPenjamin+"' "+
		// ",@FotoPenjamin = '"+fileimages["FotoPenjamin"]+"' "+
		",@FotoKTPPenjamin = '"+fileimages["FotoKTPPenjamin"]+"' "+
		",@LuasBangunan = '"+data.LuasBangunan+"' "+
		",@KondisiBangunan = '"+data.KondisiBangunan+"' "+
		",@JenisAtap = '"+data.JenisAtap+"' "+
		",@Dinding = '"+data.Dinding+"' "+
		",@Lantai = '"+data.Lantai+"' "+
		// ",@FotoRumah1 = '"+fileimages["FotoRumah1"]+"' "+
		// ",@FotoRumah2 = '"+fileimages["FotoRumah2"]+"' "+
		// ",@Foto_Jenis_Atap = '"+fileimages["Foto_Jenis_Atap"]+"' "+
		// ",@Foto_Dinding = '"+fileimages["Foto_Dinding"]+"' "+
		// ",@Foto_Lantai = '"+fileimages["Foto_Lantai"]+"' "+
		// ",@Nilai_Total = '"+data.Nilai_Total+"' "+
		// ",@Is_AdaAksesAirBersih = '"+data.Is_AdaAksesAirBersih+"' "+
		// ",@Is_AdaAdaToiletPribadi = '"+data.Is_AdaAdaToiletPribadi+"' "+
		// ",@Is_Renov_Existing = '"+data.Is_Renov_Existing+"' "+
		// ",@Is_REnov_Baru = '"+data.Is_REnov_Baru+"' "+
		// ",@Is_Atap_Galvalum = '"+data.Is_Atap_Galvalum+"' "+
		// ",@Is_Atap_Anyaman = '"+data.Is_Atap_Anyaman+"' "+
		// ",@Is_Atap_None = '"+data.Is_Atap_None+"' "+
		// ",@Is_Lantai_Keramik = '"+data.Is_Lantai_Keramik+"' "+
		// ",@Is_Lantai_Semen = '"+data.Is_Lantai_Semen+"' "+
		// ",@Is_Lantai_KayuBambu = '"+data.Is_Lantai_KayuBambu+"' "+
		// ",@Is_Lantai_None = '"+data.Is_Lantai_None+"' "+
		// ",@Is_Dinding_Tertutup = '"+data.Is_Dinding_Tertutup+"' "+
		// ",@Is_Dinding_TertutupSebagian = '"+data.Is_Dinding_TertutupSebagian+"' "+
		// ",@Is_Dinding_None = '"+data.Is_Dinding_None+"' "+
		// ",@Is_Material_Batako = '"+data.Is_Material_Batako+"' "+
		// ",@Is_Material_GRC = '"+data.Is_Material_GRC+"' "+
		// ",@Is_Material_KayuBambu = '"+data.Is_Material_KayuBambu+"' "+
		// ",@Is_Material_Seng = '"+data.Is_Material_Seng+"' "+
		// ",@Is_Material_None = '"+data.Is_Material_None+"' "+
		// ",@Is_Renov_Lantai = '"+data.Is_Renov_Lantai+"' "+
		// ",@Pj_Lt = '"+data.Pj_Lt+"' "+
		// ",@Lbr_Lt = '"+data.Lbr_Lt+"' "+
		// ",@Is_Renov_Dinding = '"+data.Is_Renov_Dinding+"' "+
		// ",@Pj_Dinding = '"+data.Pj_Dinding+"' "+
		// ",@Lbr_Dinding = '"+data.Lbr_Dinding+"' "+
		// ",@Is_Renov_Atap = '"+data.Is_Renov_Atap+"' "+
		// ",@Pj_Atap = '"+data.Pj_Atap+"' "+
		// ",@Lbr_Atap = '"+data.Lbr_Atap+"' "+
		// ",@Is_Renov_Lainnya = '"+data.Is_Renov_Lainnya+"' "+
		// ",@Ket_Renov_Lainnya = '"+data.Ket_Renov_Lainnya+"' "+
		// ",@Foto_RAB = '"+fileimages["Foto_RAB"]+"' "+
		// ",@Foto_PraRenov = '"+fileimages["Foto_PraRenov"]+"' "+
		// ",@Is_KMdanToiletTerpisah = '"+data.Is_KMdanToiletTerpisah+"' "+
		// ",@Is_AdaKMPribadi = '"+data.Is_AdaKMPribadi+"' "+
		// ",@Is_AdaToiletPribadi = '"+data.Is_AdaToiletPribadi+"' "+
		// ",@Is_MataAir_utkMandi = '"+data.Is_MataAir_utkMandi+"' "+
		// ",@Is_MAtaAir_utkMinum = '"+data.Is_MAtaAir_utkMinum+"' "+
		// ",@Is_BeliAirMandi = '"+data.Is_BeliAirMandi+"' "+
		// ",@Is_BeliAirMinum = '"+data.Is_BeliAirMinum+"' "+
		// ",@Is_Sumur_utkMandi = '"+data.Is_Sumur_utkMandi+"' "+
		// ",@Is_SUmur_utkMinum = '"+data.Is_SUmur_utkMinum+"' "+
		// ",@Is_AirRefill_utkMandi = '"+data.Is_AirRefill_utkMandi+"' "+
		// ",@Is_AirRefill_utkMinum = '"+data.Is_AirRefill_utkMinum+"' "+
		// ",@Is_PDAM_utkMandi = '"+data.Is_PDAM_utkMandi+"' "+
		// ",@Is_PDAM_utkMinum = '"+data.Is_PDAM_utkMinum+"' "+
		// ",@Is_Lainnya_utkMandi = '"+data.Is_Lainnya_utkMandi+"' "+
		// ",@Is_Lainnya_utkMinum = '"+data.Is_Lainnya_utkMinum+"' "+
		// ",@Is_DebitAirBerlimpah = '"+data.Is_DebitAirBerlimpah+"' "+
		// ",@Is_DebitAirNormal = '"+data.Is_DebitAirNormal+"' "+
		// ",@Is_DebitAirKering = '"+data.Is_DebitAirKering+"' "+
		// ",@Is_AirBerbau = '"+data.Is_AirBerbau+"' "+
		// ",@Is_AirBerasa = '"+data.Is_AirBerasa+"' "+
		// ",@Is_AirBerwarna = '"+data.Is_AirBerwarna+"' "+
		// ",@Is_AirBerlumpur = '"+data.Is_AirBerlumpur+"' "+
		// ",@Is_AirNormal = '"+data.Is_AirNormal+"' "+
		// ",@Is_AtapKM_Terbuka = '"+data.Is_AtapKM_Terbuka+"' "+
		// ",@Is_AtapKM_Tertutup = '"+data.Is_AtapKM_Tertutup+"' "+
		// ",@Is_LantaiKM_Tanah = '"+data.Is_LantaiKM_Tanah+"' "+
		// ",@Is_LantaiKM_Semen = '"+data.Is_LantaiKM_Semen+"' "+
		// ",@Is_LantaiKM_KayuBambu = '"+data.Is_LantaiKM_KayuBambu+"' "+
		// ",@Is_LantaiKM_Keramik = '"+data.Is_LantaiKM_Keramik+"' "+
		// ",@Is_DindingKM_Tertutup = '"+data.Is_DindingKM_Tertutup+"' "+
		// ",@Is_DindingKM_TerbukaSebagian = '"+data.Is_DindingKM_TerbukaSebagian+"' "+
		// ",@Is_DindingKM_KayuBambu = '"+data.Is_DindingKM_KayuBambu+"' "+
		// ",@Is_DindingKM_Anyaman = '"+data.Is_DindingKM_Anyaman+"' "+
		// ",@Is_DindingKM_BatakoBata = '"+data.Is_DindingKM_BatakoBata+"' "+
		// ",@Is_BakAir_Ember = '"+data.Is_BakAir_Ember+"' "+
		// ",@Is_BakAir_Semen = '"+data.Is_BakAir_Semen+"' "+
		// ",@Is_BakAir_Fiber = '"+data.Is_BakAir_Fiber+"' "+
		// ",@Is_BakAir_Keramik = '"+data.Is_BakAir_Keramik+"' "+
		// ",@Is_Sewer_Sungai = '"+data.Is_Sewer_Sungai+"' "+
		// ",@Is_Sewer_SaluranAir = '"+data.Is_Sewer_SaluranAir+"' "+
		// ",@Is_Sewer_Kolam = '"+data.Is_Sewer_Kolam+"' "+
		// ",@Is_Sewer_SepticTank = '"+data.Is_Sewer_SepticTank+"' "+
		// ",@Is_WC_Jamban = '"+data.Is_WC_Jamban+"' "+
		// ",@Is_WC_NonJamban = '"+data.Is_WC_NonJamban+"' "+
		// ",@Is_WC_Jongkok = '"+data.Is_WC_Jongkok+"' "+
		// ",@Is_WC_Duduk = '"+data.Is_WC_Duduk+"' "+
		// ",@Is_AtapWC_Terbuka = '"+data.Is_AtapWC_Terbuka+"' "+
		// ",@Is_AtapWC_Tertutup = '"+data.Is_AtapWC_Tertutup+"' "+
		// ",@Is_LantaiWC_Tanah = '"+data.Is_LantaiWC_Tanah+"' "+
		// ",@Is_LantaiWC_Semen = '"+data.Is_LantaiWC_Semen+"' "+
		// ",@Is_LantaiWC_KayuBambu = '"+data.Is_LantaiWC_KayuBambu+"' "+
		// ",@Is_LantaiWC_Keramik = '"+data.Is_LantaiWC_Keramik+"' "+
		// ",@Is_DindingWC_Tertutup = '"+data.Is_DindingWC_Tertutup+"' "+
		// ",@Is_DindingWC_TerbukaSebagian = '"+data.Is_DindingWC_TerbukaSebagian+"' "+
		// ",@Is_DindingWC_KayuBambu = '"+data.Is_DindingWC_KayuBambu+"' "+
		// ",@Is_DindingWC_Anyaman = '"+data.Is_DindingWC_Anyaman+"' "+
		// ",@Is_DindingWC_BatakoBata = '"+data.Is_DindingWC_BatakoBata+"' "+
		// ",@Is_BakAirWC_Ember = '"+data.Is_BakAirWC_Ember+"' "+
		// ",@Is_BakAirWC_Semen = '"+data.Is_BakAirWC_Semen+"' "+
		// ",@Is_BakAirWC_Fiber = '"+data.Is_BakAirWC_Fiber+"' "+
		// ",@Is_BakAirWC_Keramik = '"+data.Is_BakAirWC_Keramik+"' "+
		// ",@Is_SewerWC_Sungai = '"+data.Is_SewerWC_Sungai+"' "+
		// ",@Is_SewerWC_SaluranAir = '"+data.Is_SewerWC_SaluranAir+"' "+
		// ",@Is_SewerWC_Kolam = '"+data.Is_SewerWC_Kolam+"' "+
		// ",@Is_SewerWC_SepticTank = '"+data.Is_SewerWC_SepticTank+"' "+
		// ",@Is_KRVerified = '"+data.Is_KRVerified+"' "+
		",@Siklus = '"+data.Siklus +"' "+
		",@JenisPembiayaan = '"+data.JenisPembiayaan+"' "+
		",@IDProduk = '"+data.IDProduk+"' "+
		",@NamaProduk = '"+data.NamaProduk+"' "+
		",@IDProdukPembiayaan = '"+data.IDProdukPembiayaan+"' "+
		",@ProdukPembiayaan = '"+data.ProdukPembiayaan+"' "+
		",@JumlahPinjaman = '"+data.JumlahPinjaman+"' "+
		",@TermPembiayaan = '"+data.TermPembiayaan+"' "+
		",@KategoriTujuanPembiayaan = '"+data.KategoriTujuanPembiayaan+"' "+
		",@TujuanPembiayaan = '"+data.TujuanPembiayaan+"' "+
		",@TypePencairan = '"+data.TypePencairan+"' "+
		",@FrekuensiPembiayaan = '"+data.FrekuensiPembiayaan+"' "+
		// ",@IsProductVerified = '"+data.IsProductVerified+"' "+
		",@RekeningBank = '1' "+
		",@NamaBankPemilik = '"+data.NamaBank+"' "+
		",@NoRekening = '"+data.NoRekening+"' "+
		",@PemilikRekening = '"+data.PemilikRekening+"' "+
		",@ID_SektorEkonomi = '"+data.ID_SektorEkonomi+"' "+
		",@ID_SubSektorEkonomi = '"+data.ID_SubSektorEkonomi+"' "+
		",@Jenis_Usaha = '"+data.Jenis_Usaha+"' "+
		// ",@Lokasi_Usaha_1 = '"+data.Lokasi_Usaha_1+"' "+
		// ",@Lokasi_Usaha_2 = '"+data.Lokasi_Usaha_2+"' "+
		",@Foto_Usaha = '"+fileimages["Foto_Usaha"]+"' "+
		// ",@Is_SEVerified = '"+data.Is_SEVerified+"' "+
		",@PendapatanKotor_perHari = '"+data.PendapatanKotor_perHari+"' "+
		",@PengeluaranKel_perHari = '"+data.PengeluaranKel_perHari+"' "+
		",@PendapatanBersih_perHari = '"+data.PendapatanBersih_perHari+"' "+
		",@JmlHariUsaha_perBulan = '"+data.JmlHariUsaha_perBulan+"' "+
		",@PendapatanBersih_perBulan = '"+data.PendapatanBersih_perBulan+"' "+
		",@PendapatanBersih_perMinggu = '"+data.PendapatanBersih_perMinggu+"' "+
		",@Is_AdaPembiayaanLain = '"+data.Is_AdaPembiayaanLain+"' "+
		// ",@Nama_Tipe = '"+data.Nama_Tipe+"' "+
		",@Nama_Pembiayaan_Lembaga_Lain = '"+data.Nama_Pembiayaan_Lembaga_Lain+"' "+

		// ",@Jumlah_Pendapatan_Bersih = '"+data.Jumlah_Pendapatan_Bersih+"' "+
		// ",@Tingkat_Pendapatan_Bersih = '"+data.Tingkat_Pendapatan_Bersih+"' "+
		",@Kemampuan_Angsuran = '"+data.Kemampuan_Angsuran+"' "+
		",@Berdasarkan_Tingkat_Pendapatan = '"+data.Berdasarkan_Tingkat_Pendapatan+"' "+
		",@Berdasarkan_Kemampuan_Angsuran = '"+data.Berdasarkan_Kemampuan_Angsuran+"' "+
		",@Berdasarkan_Lembaga_Lain = '"+data.Berdasarkan_Lembaga_Lain+"' "+		

		",@PendapatanKotor_perHari_Suami = '"+data.PendapatanKotor_perHari_Suami+"' "+
		",@PengeluaranKel_perHari_Suami = '"+data.PengeluaranKel_perHari_Suami+"' "+
		",@PendapatanBersih_perHari_Suami = '"+data.PendapatanBersih_perHari_Suami+"' "+
		",@JmlHariUsaha_perBulan_Suami = '"+data.JmlHariUsaha_perBulan_Suami+"' "+
		",@PendapatanBersih_perBulan_Suami = '"+data.PendapatanBersih_perBulan_Suami+"' "+
		",@PendapatanBersih_perMinggu_Suami = '"+data.PendapatanBersih_perMinggu_Suami+"' "+
		// ",@Is_TPVerified = '"+data.Is_TPVerified+"' "+
		",@TTD_Nasabah = '"+fileimages["TTD_Nasabah"]+"' "+
		// ",@TTD_Suami = '"+fileimages["TTD_Suami"]+"' "+
		",@TTD_Penjamin = '"+fileimages["TTD_Penjamin"]+"' "+
		",@TTD_KSK = '"+fileimages["TTD_KSK"]+"' "+
		",@TTD_KK = '"+fileimages["TTD_KK"]+"' "+
		",@TTD_AO = '"+fileimages["TTD_AO"]+"' "+
		// ",@TTD_KC_SAO = '"+fileimages["TTD_KC_SAO"]+"' "+
		",@Tujuan_Transfer_Check = '"+Tujuan_Transfer_Check+"' "+
		",@Nama_TTD_Nasabah = '"+data.Nama_TTD_Nasabah+"' "+
		",@Nama_TTD_Penjamin = '"+data.Nama_TTD_Penjamin+"' "+
		",@Nama_TTD_KSK = '"+data.Nama_TTD_KSK+"' "+
		",@Nama_TTD_KK = '"+data.Nama_TTD_KK+"' "+
		",@Nama_TTD_AO = '"+data.Nama_TTD_AO+"' "+
		// ",@Nama_TTD_KC_SAO = '"+data.Nama_TTD_KC_SAO+"' "

		",@Kelompok_ID = '"+data.Kelompok_ID+"' "+
		",@Nama_Kelompok = '"+data.Nama_Kelompok+"' "+
		",@Sub_Kelompok = '"+data.Sub_Kelompok+"' "+
		",@ClientID = '"+data.ClientID+"' "+
		// ",@Nama_KK = '"+data.Nama_KK+"' "+
		// ",@Nama_KSK = '"+data.Nama_KSK+"' "

		",@KualitasKehadiranNasabah = '"+data.Kehadiran_pkm+"' "+
		",@KualitasPembayaranNasabah = '"+data.Angsuran_Pada_Saat_PKM+"' "

		
		global.Logging("INFO","models.POST_PROSPEK_UK STORE_PROCEDURE ADD_PROSPEK_UK "+query+" ")

		start_sql := time.Now()		

		err = db.RawSQL(query).Do(&responsedata)
		if err != nil {		
			global.Logging("ERROR","models POST_PROSPEK_UK ---> "+query+" ---> " + err.Error())		
			errsql := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {return responsedata,err}
			return responsedata,errsql
		}		

		elapsed_sql := time.Since(start_sql)
		global.Logging("INFO","STORE_PROCEDURE process_time POST_PROSPEK_UK "+elapsed_sql.String()+" ---> ")	
	// }

	err = db.Commit() //Commit Transaction
	if err != nil {return responsedata,err}

	err = db.Close()
	if err != nil {return responsedata,err}	

	return responsedata,nil
}

type SP_ADD_PERSETUJUAN_KELOMPOK struct {
	ID_Prospek string `json:"ID_Prospek"`
	No_Identitas string `json:"No_Identitas"`
	ID_Kelompok string `json:"ID_Kelompok"`
	Tahap_Pembiayaan string `json:"Tahap_Pembiayaan"`
	Plafon_Diajukan string `json:"Plafon_Diajukan"`
	Tenor string `json:"Tenor"`
	Status_Tempat_Tinggal string `json:"Status_Tempat_Tinggal"`
	Is_Perkawinan_Berubah string `json:"Is_Perkawinan_Berubah"`
	Is_Tanggungan_Berubah string `json:"Is_Tanggungan_Berubah"`
	Keterangan_Tanggungan string `json:"Keterangan_Tanggungan"`
	Jml_Kehadiran_PKM string `json:"Jml_Kehadiran_PKM"`
	Jml_Pembayaran string `json:"Jml_Pembayaran"`
	Is_Usaha_Berubah string `json:"Is_Usaha_Berubah"`
	Keterangan_Usaha string `json:"Keterangan_Usaha"`
	Lokasi_Persetujuan string `json:"Lokasi_Persetujuan"`
	Tanggal_Persetujuan string `json:"Tanggal_Persetujuan"`
	TTD_KK string `json:"TTD_KK"`
	TTD_KSK string `json:"TTD_KSK"`
	TTD_AO string `json:"TTD_AO"`
}

func POST_PROSPEK_LAMA(params []SP_ADD_PERSETUJUAN_KELOMPOK)error {
	db,err := global.Conn()
	if err != nil {
		return err
	}	
	var query string
	var fileimages = make(map[string]string)
	// path := make([]PATH, 0)


	err = db.Begin() //Begin Transaction
	if err != nil {return err}		

	for _, data := range params {	

		fileimages["TTD_KK"] = data.TTD_KK
		fileimages["TTD_KSK"] = data.TTD_KSK	
		fileimages["TTD_AO"] = data.TTD_AO

		// err = db.RawSQL("SELECT CONCAT('images/',CAST(GETDATE() as DATE),'/','"+data.ID_Prospek+"','/') AS PATH_FILE ").Do(&path)
		// if err != nil {		
		// 	errsql := err
		// 	err = db.Rollback() //Rollback Transaction
		// 	if err != nil {return err}
		// 	return errsql
		// }		
			
		// fileimages,err = global.MapB64SaveFile(fileimages,path[0].FILE)
		fileimages,err = global.MapB64SaveFile(fileimages,beego.AppConfig.String("pathimages"))
		if err != nil { 			
			errfile := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {return err}
			return errfile 
		}			

		query = "EXEC ADD_PERSETUJUAN_KELOMPOK "+
		"@ID_Prospek = '"+data.ID_Prospek+"' "+
		",@No_Identitas = '"+data.No_Identitas+"' "+
		",@ID_Kelompok = '"+data.ID_Kelompok+"' "+
		",@Tahap_Pembiayaan = '"+data.Tahap_Pembiayaan+"' "+
		",@Plafon_Diajukan = '"+data.Plafon_Diajukan+"' "+
		",@Tenor = '"+data.Tenor+"' "+
		",@Status_Tempat_Tinggal = '"+data.Status_Tempat_Tinggal+"' "+
		",@Is_Perkawinan_Berubah = '"+data.Is_Perkawinan_Berubah+"' "+
		",@Is_Tanggungan_Berubah = '"+data.Is_Tanggungan_Berubah+"' "+
		",@Keterangan_Tanggungan = '"+data.Keterangan_Tanggungan+"' "+
		",@Jml_Kehadiran_PKM = '"+data.Jml_Kehadiran_PKM+"' "+
		",@Jml_Pembayaran = '"+data.Jml_Pembayaran+"' "+
		",@Is_Usaha_Berubah = '"+data.Is_Usaha_Berubah+"' "+
		",@Keterangan_Usaha = '"+data.Keterangan_Usaha+"' "+
		",@Lokasi_Persetujuan = '"+data.Lokasi_Persetujuan+"' "+
		",@Tanggal_Persetujuan = '"+data.Tanggal_Persetujuan+"' "+
		",@TTD_KK = '"+fileimages["TTD_KK"]+"' "+
		",@TTD_KSK = '"+fileimages["TTD_KSK"]+"' "+
		",@TTD_AO = '"+fileimages["TTD_AO"]+"' "	

		global.Logging("INFO","models.POST_PROSPEK_LAMA STORE_PROCEDURE ADD_PERSETUJUAN_KELOMPOK "+query+" ")

		_ , err = db.RawSQL(query).DoWithIterator()
		if err != nil {		
			global.Logging("ERROR","models.POST_PROSPEK_LAMA ---> "+query+" ---> " + err.Error())	
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

type SP_ADD_DATA_KELOMPOK struct {
	OurBranchID string `json:"OurBranchID"`
	IDKelompok string `json:"IDKelompok"`
	NamaKelompok string `json:"NamaKelompok"`
	GroupProduct string `json:"GroupProduct"`
	TanggalPertemuan string `json:"TanggalPertemuan"`
	HariPertemuan string `json:"HariPertemuan"`
	WaktuPertemuan string `json:"WaktuPertemuan"`
	LokasiPertemuan string `json:"LokasiPertemuan"`
	ClientTotal string `json:"ClientTotal"`
	ID_DK_Mobile string `json:"ID_DK_Mobile"`
}

func POST_DATA_KELOMPOK(data SP_ADD_DATA_KELOMPOK)error {
	db,err := global.Conn()
	if err != nil {
		return err
	}	
	var query string

	err = db.Begin() //Begin Transaction
	if err != nil {return err}		

	query = "EXEC ADD_DATA_KELOMPOK "+
	"@OurBranchID = '"+data.OurBranchID+"' "+
	",@IDKelompok = '"+data.IDKelompok+"' "+
	",@NamaKelompok = '"+data.NamaKelompok+"' "+
	",@GroupProduct = '"+data.GroupProduct+"' "+
	",@TanggalPertemuan = '"+data.TanggalPertemuan+"' "+
	",@HariPertemuan = '"+data.HariPertemuan+"' "+
	",@WaktuPertemuan = '"+data.WaktuPertemuan+"' "+
	",@LokasiPertemuan = '"+data.LokasiPertemuan+"' "+
	",@ClientTotal = '"+data.ClientTotal+"' "+
	",@ID_DK_Mobile = '"+data.ID_DK_Mobile+"' "

	global.Logging("INFO","models.POST_DATA_KELOMPOK STORE_PROCEDURE ADD_DATA_KELOMPOK "+query+" ")

	_ , err = db.RawSQL(query).DoWithIterator()
	if err != nil {		
		global.Logging("ERROR","models POST_DATA_KELOMPOK ---> "+query+" ---> " + err.Error())	
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

type SP_ADD_INISIASI_PP struct {

	// OurBranchID string `json:"OurBranchID" valid:"Required"`

	ID_MPP string `json:"ID_MPP"`
	ID_Prospek string `json:"ID_Prospek"`
	ID_Kelompok string `json:"ID_Kelompok"`
	Nama_Kelompok string `json:"Nama_Kelompok"`
	Sub_Kelompok string `json:"Sub_Kelompok"`
	ID_Absen string `json:"ID_Absen"`
	Keterangan_Absen string `json:"Keterangan_Absen"`
	ID_DK_Mobile string `json:"ID_DK_Mobile"`
	// ID_Absen_2 string `json:"ID_Absen_2"`
	// Keterangan_Absen_2 string `json:"Keterangan_Absen_2"`
	// ID_Absen_3 string `json:"ID_Absen_3"`
	// Keterangan_Absen_3 string `json:"Keterangan_Absen_3"`
}

func POST_PP(data SP_ADD_INISIASI_PP)error {
	// GROUPID := make([]RESPONSE_SET_GROUP_BRNET, 0)

	db,err := global.Conn()
	if err != nil {
		return err
	}	
	var query string	


	// check validasi start		
	valid := validation.Validation{}
	v, err := valid.Valid(&data)
	if err != nil {return err}	
	if !v {
		for _, err := range valid.Errors {
			return err
		}
	}
	// check validasi end	

	err = db.Begin() //Begin Transaction
	if err != nil {return err}		

	query = "EXEC ADD_INISIASI_PP "+
	"@ID_MPP = '"+data.ID_MPP+"' "+
	",@ID_Prospek = '"+data.ID_Prospek+"' "+
	",@ID_Kelompok = '"+data.ID_Kelompok+"' "+
	",@Nama_Kelompok = '"+data.Nama_Kelompok+"' "+
	",@Sub_Kelompok = '"+data.Sub_Kelompok+"' "+
	",@ID_Absen = '"+data.ID_Absen+"' "+
	",@Keterangan_Absen_1 = '"+data.Keterangan_Absen+"' "+
	",@ID_DK_Mobile = '"+data.ID_DK_Mobile+"' "
	// ",@ID_Absen_2 = '"+data.ID_Absen_2+"' "+
	// ",@Keterangan_Absen_2 = '"+data.Keterangan_Absen_2+"' "+
	// ",@ID_Absen_3 = '"+data.ID_Absen_3+"' "+
	// ",@Keterangan_Absen_3 = '"+data.Keterangan_Absen_3+"' "



	global.Logging("INFO","models.POST_PP STORE_PROCEDURE ADD_INISIASI_PP "+query+" ")

	_ , err = db.RawSQL(query).DoWithIterator()
	if err != nil {		
		global.Logging("ERROR","models POST_PP ---> "+query+" ---> " + err.Error())			
		errsql := err
		err = db.Rollback() //Rollback Transaction
		if err != nil {return err}
		return errsql
	}			

	if (data.ID_MPP == "3"){			
		query = "UPDATE INISIASI_PROSPEK SET Is_PPPass='1' WHERE ID_Prospek='"+data.ID_Prospek+"' "

		_ , err = db.RawSQL(query).DoWithIterator()
		if err != nil {		
			global.Logging("ERROR","models POST_PP ---> "+query+" ---> " + err.Error())
			errsql := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {return err}
			return errsql
		}				
	}

	// fmt.Println(GROUPID[0].GroupID)


	err = db.Commit() //Commit Transaction
	if err != nil {return err}

	err = db.Close()
	if err != nil {return err}	


	return nil	

}

type SP_SET_VERIF_STATUS struct {
	PostStatus string `json:"PostStatus"`
	Reason string `json:"Reason"`
	ID_PROSPEK string `json:"ID_Prospek"`
	VerifBy string `json:"VerifBy"`
	TTD_KC_SAO string `json:"TTD_KC_SAO"`
	URL_FP4 string `json:"URL_FP4"`
}


func POST_VERIF_STATUS(data SP_SET_VERIF_STATUS)error {
	db,err := global.Conn()
	if err != nil {
		return err
	}	
	var query string
	var fileimages = make(map[string]string)
	// path := make([]PATH, 0)

	err = db.Begin() //Begin Transaction
	if err != nil {return err}		


	if (data.TTD_KC_SAO != "") {
		fileimages["TTD_KC_SAO"] = data.TTD_KC_SAO

		// err = db.RawSQL("SELECT CONCAT('images/',CAST(GETDATE() as DATE),'/',"+data.ID_PROSPEK+",'/') AS PATH_FILE ").Do(&path)
		// if err != nil {		
		// 	errsql := err
		// 	err = db.Rollback() //Rollback Transaction
		// 	if err != nil {return err}
		// 	return errsql
		// }	

		// fileimages,err = global.MapB64SaveFile(fileimages,path[0].FILE)
		fileimages,err = global.MapB64SaveFile(fileimages,beego.AppConfig.String("pathimages"))
		if err != nil { 			
			errfile := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {return err}
			return errfile 
		}	

		query = "EXEC SET_VERIF_STATUS "+
		"@PostStatus = '"+data.PostStatus+"' "+
		",@Reason = '"+data.Reason+"' "+
		",@ID_PROSPEK = '"+data.ID_PROSPEK+"' "+
		",@VerifBy = '"+data.VerifBy+"' "+
		",@TTD_KC_SAO = '"+fileimages["TTD_KC_SAO"]+"' "
	
	}else{
		query = "EXEC SET_VERIF_STATUS "+
		"@PostStatus = '"+data.PostStatus+"' "+
		",@Reason = '"+data.Reason+"' "+
		",@ID_PROSPEK = '"+data.ID_PROSPEK+"' "+
		",@VerifBy = '"+data.VerifBy+"' "

	}


	global.Logging("INFO","models.POST_VERIF_STATUS STORE_PROCEDURE SET_VERIF_STATUS "+query+" ")

	_ , err = db.RawSQL(query).DoWithIterator()
	if err != nil {		
		global.Logging("ERROR","models POST_VERIF_STATUS ---> "+query+" ---> " + err.Error())	
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

type SET_DATA_FROM_INISIASI_MOBILE_BRNET struct {
	ClientID string `json:"ClientID" db:"ClientID"`
	File_FP4 string `json:"File_FP4" db:"File_FP4"`
	Foto_KK string `json:"Foto_KK" db:"Foto_KK"`
	Foto_Kartu_Identitas string `json:"Foto_Kartu_Identitas" db:"Foto_Kartu_Identitas"`
	Nomor_Rekening string `json:"Nomor_Rekening" db:"Nomor_Rekening"`
	Nama_Bank string `json:"Nama_Bank" db:"Nama_Bank"`
	Nama_Pemilik_Rekening string `json:"Nama_Pemilik_Rekening" db:"Nama_Pemilik_Rekening"`
	Nomor_Identitas string `json:"Nomor_Identitas" db:"Nomor_Identitas"`
}

func POST_SET_DATA_FROM_INISIASI_MOBILE_BRNET(param SP_SET_VERIF_STATUS)error {
	data := make([]SET_DATA_FROM_INISIASI_MOBILE_BRNET, 0)

	dbget,err := global.Conn()
	if err != nil {
		return err
	}	
	var query string


	query = "SELECT "+
	"a.ClientID "+
	",'"+param.URL_FP4+"' as File_FP4 "+
	",b.Foto_KK "+
	",b.Foto_Kartu_Identitas "+
	",b.Nomor_Rekening "+
	",b.Nama_Bank "+
	",b.Nama_Pemilik_Rekening "+
	",b.Nomor_Identitas "+
	"FROM INISIASI_PROSPEK a WITH (NOLOCK) "+
	"INNER JOIN INISIASI_UK_DATA_PRIBADI b WITH (NOLOCK) ON a.ID_Prospek=b.ID_Prospek "+
	"WHERE ISNULL(a.ClientID,'')!='' and a.ID_Prospek='"+param.ID_PROSPEK+"' "
	
	err = dbget.RawSQL(query).Do(&data)
	if err != nil {		
		global.Logging("ERROR","models POST_SET_DATA_FROM_INISIASI_MOBILE_BRNET ---> "+query+" ---> " + err.Error())	
		return err
	}	

	if len(data)>0 {
		
		db,err := global.ConnSCORING()
		if err != nil {
			return err
		}	

		err = db.Begin() //Begin Transaction
		if err != nil {return err}	

		query = "EXEC dbo.[SET_DATA_FROM_INISIASI_MOBILE] @ClientID='"+data[0].ClientID+"' "+
				",@File_FP4='"+param.URL_FP4+"' "+
				",@File_KK ='"+data[0].Foto_KK+"' "+
				",@File_KTP ='"+data[0].Foto_Kartu_Identitas+"' "+
				",@No_Rekening ='"+data[0].Nomor_Rekening+"' "+
				",@Nama_Bank ='"+data[0].Nama_Bank+"' "+
				",@Pemilik_Rekening ='"+data[0].Nama_Pemilik_Rekening+"' "+
				",@Tujuan_Transfer = '2' "+
				",@No_KTP ='"+data[0].Nomor_Identitas+"' "

		global.Logging("INFO","models.POST_SET_DATA_FROM_INISIASI_MOBILE_BRNET STORE_PROCEDURE PNM_LIVE.dbo.SET_DATA_FROM_INISIASI_MOBILE "+query+" ")				

		_ , err = db.RawSQL(query).DoWithIterator()
		if err != nil {		
			global.Logging("ERROR","BRNET SET_DATA_FROM_INISIASI_MOBILE UNTUK SCORING ---> "+query+" ---> " + err.Error())	
			errsql := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {return err}
			return errsql
		}		

		err = db.Commit() //Commit Transaction
		if err != nil {return err}

		err = db.Close()
		if err != nil {return err}		

	}

	err = dbget.Close()
	if err != nil {return err}	


	return nil
}


type SP_ADD_KK_KSK_DATA struct {
	ID_Prospek string `json:"ID_Prospek" db:"ID_Prospek"`
	Nama_KK string `json:"Nama_KK" db:"Nama_KK"`
	Nama_KSK string `json:"Nama_KSK" db:"Nama_KSK"`
}

func POST_KK_KSK_DATA(data SP_ADD_KK_KSK_DATA)error {
	db,err := global.Conn()
	if err != nil {
		return err
	}	
	var query string

	err = db.Begin() //Begin Transaction
	if err != nil {return err}	

	query = "EXEC ADD_KK_KSK_DATA "+
	"@ID_Prospek = '"+data.ID_Prospek+"' "+
	",@Nama_KK = '"+data.Nama_KK+"' "+
	",@Nama_KSK = '"+data.Nama_KSK+"' "

	global.Logging("INFO","models.POST_KK_KSK_DATA STORE_PROCEDURE ADD_KK_KSK_DATA "+query+" ")					
	
	_ , err = db.RawSQL(query).DoWithIterator()
	if err != nil {		
		global.Logging("ERROR","models POST_KK_KSK_DATA ---> "+query+" ---> " + err.Error())	
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

//========================================== START PERSETUJUAN PEMBIAYAAN && SYNC BR NET ==========================================//

type SP_ADD_PERSETUJUAN_PEMBIAYAAN struct {
	ID_Prospek string `json:"ID_Prospek" valid:"Required"`
	Keterangan_Pembelian string `json:"Keterangan_Pembelian"`
	TTD_AO string `json:"TTD_AO" valid:"Base64"`
	TTD_KC string `json:"TTD_KC" valid:"Base64"`
	TTD_KK string `json:"TTD_KK"`
	TTD_KSK string `json:"TTD_KSK"`
}

type PARAM_TO_BRNET struct {
	OurBranchID string `json:"OurBranchID" db:"OurBranchID"` 
	Nama_Kelompok string `json:"Nama_Kelompok" db:"Nama_Kelompok"` 
	Kelompok_ID string `json:"Kelompok_ID" db:"Kelompok_ID"` 
	ClientID string `json:"ClientID" db:"ClientID"` 
	GroupProductID string `json:"GroupProductID" db:"GroupProductID"` 
	DefaultLoanSchemeID string `json:"DefaultLoanSchemeID" db:"DefaultLoanSchemeID"` 
	Frekuesi_Pertemuan string `json:"Frekuesi_Pertemuan" db:"Frekuesi_Pertemuan"` 
	Tanggal_PKM_Pertama string `json:"Tanggal_PKM_Pertama" db:"Tanggal_PKM_Pertama"` 
	Hari_Pertemuan string `json:"Hari_Pertemuan" db:"Hari_Pertemuan"` 
	Nama_Lengkap string `json:"Nama_Lengkap" db:"Nama_Lengkap"` 
	Nama_Ayah string `json:"Nama_Ayah" db:"Nama_Ayah"` 
	Nama_Gadis_Ibu_Kandung string `json:"Nama_Gadis_Ibu_Kandung" db:"Nama_Gadis_Ibu_Kandung"` 
	Tanggal_Lahir string `json:"Tanggal_Lahir" db:"Tanggal_Lahir"` 
	Nomor_Identitas string `json:"Nomor_Identitas" db:"Nomor_Identitas"` 
	Alamat_Domisili string `json:"Alamat_Domisili" db:"Alamat_Domisili"` 
	SubGroup string `json:"SubGroup" db:"SubGroup"` 
	Nama_Produk string `json:"Nama_Produk" db:"Nama_Produk"` 
	Jumlah_Pinjaman string `json:"Jumlah_Pinjaman" db:"Jumlah_Pinjaman"` 
	Term_Pembiayaan string `json:"Term_Pembiayaan" db:"Term_Pembiayaan"` 	
	ID_Prospek string `json:"ID_Prospek" db:"ID_Prospek"` 	
	Is_PPPass string `json:"Is_PPPass" db:"Is_PPPass"` 	
	Is_Sisipan string `json:"Is_Sisipan" db:"Is_Sisipan"` 	
	Is_LanjutUK string `json:"Is_LanjutUK" db:"Is_LanjutUK"` 	
	NextMeetingDate string `json:"NextMeetingDate" db:"NextMeetingDate"` 	
	Siklus string `json:"Siklus" db:"Siklus"` 	
	Created_By string `json:"Created_By" db:"Created_By"`
	ID_DK string `json:"ID_DK" db:"ID_DK"`
}

func POST_PERSETUJUAN_PEMBIAYAAN(data_add_persetujuan_pembiayaan SP_ADD_PERSETUJUAN_PEMBIAYAAN)error {
	params := make([]PARAM_TO_BRNET, 0)
	PROJECTION_SANCTION := make([]RESPONSE_POST_PROJECTION_SANCTION, 0)
	// path := make([]PATH, 0)

	db,err := global.Conn()
	if err != nil {
		return err
	}	
	var query string
	var fileimages = make(map[string]string)

	// check validasi start		
	valid := validation.Validation{}
	v, err := valid.Valid(&data_add_persetujuan_pembiayaan)
	if err != nil {return err}	
	if !v {
		for _, err := range valid.Errors {
			return err
		}
	}
	// check validasi end	

	err = db.Begin() //Begin Transaction
	if err != nil {return err}		

	//ambil last id prospek sebelum di insert untuk bikin path file images/tanggal/id prospek/
	// err = db.RawSQL("SELECT CONCAT('images/',CAST(GETDATE() as DATE),'/','"+data.ID_Prospek+"','/') AS PATH_FILE ").Do(&path)
	// if err != nil {		
	// 	global.Logging("ERROR pathfile ---> " + err.Error())			
	// 	errsql := err
	// 	err = db.Rollback() //Rollback Transaction
	// 	if err != nil {return err}
	// 	return errsql
	// }	

	fileimages["TTD_AO"] = data_add_persetujuan_pembiayaan.TTD_AO
	fileimages["TTD_KC"] = data_add_persetujuan_pembiayaan.TTD_KC
	fileimages["TTD_KK"] = data_add_persetujuan_pembiayaan.TTD_KK
	fileimages["TTD_KSK"] = data_add_persetujuan_pembiayaan.TTD_KSK

	// fileimages,err = global.MapB64SaveFile(fileimages,path[0].FILE)
	fileimages,err = global.MapB64SaveFile(fileimages,beego.AppConfig.String("pathimages"))
	if err != nil { 			
		errfile := err
		err = db.Rollback() //Rollback Transaction
		if err != nil {return err}
		return errfile 
	}	

	query = "EXEC ADD_PERSETUJUAN_PEMBIAYAAN "+
	"@ID_Prospek = '"+data_add_persetujuan_pembiayaan.ID_Prospek+"' "+
	",@Keterangan_Pembelian = '"+data_add_persetujuan_pembiayaan.Keterangan_Pembelian+"' "+
	",@TTD_AO = '"+fileimages["TTD_AO"]+"' "+
	",@TTD_KC = '"+fileimages["TTD_KC"]+"' "+
	",@TTD_KK = '"+fileimages["TTD_KK"]+"' "+
	",@TTD_KSK = '"+fileimages["TTD_KSK"]+"' "
	

	// _ , err = db.RawSQL(query).DoWithIterator()
	// if err != nil {	
	// 	global.Logging("ERROR models POST_PERSETUJUAN_PEMBIAYAAN ---> "+query+" ---> " + err.Error())	
	// 	errsql := err
	// 	err = db.Rollback() //Rollback Transaction
	// 	if err != nil {return err}
	// 	return errsql
	// }			

	

	// query = " DECLARE @TGL_PKM_PERTAMA DATETIME=NULL,@ID_PROSPEK INT "+
	// " SET @ID_PROSPEK='"+data_add_persetujuan_pembiayaan.ID_Prospek+"' "+

	// "SELECT TOP 1 @TGL_PKM_PERTAMA=c.Tanggal_PKM_Pertama "+
	// "FROM INISIASIMobile.dbo.INISIASI_PROSPEK a  "+
	// "LEFT JOIN INISIASIMobile.dbo.INISIASI_DATA_KELOMPOK c ON a.OurBranchID=c.OurBranchID AND CONCAT(ISNULL(a.ID_DK,''),ISNULL(a.Kelompok_ID,''))=CONCAT(ISNULL(c.ID_DK_Mobile,''),ISNULL(c.ID_Kelompok,''))    "+
	// "WHERE a.ID_Prospek=@ID_PROSPEK "

	query += "SELECT "+
	"a.OurBranchID "+
	",ISNULL(a.Nama_Kelompok,'') as Nama_Kelompok "+
	",CASE ISNULL(c.ID_Kelompok,'') WHEN '' "+
	"THEN ISNULL(a.Kelompok_ID,'') "+
	"ELSE ISNULL(c.ID_Kelompok,'') END as Kelompok_ID "+
	",ISNULL(ClientID,'') as ClientID "+
	",ISNULL(c.Frekuesi_Pertemuan,'') as GroupProductID "+
	",CASE ISNULL(a.Kelompok_ID,'') WHEN '' THEN '1M2520' ELSE '' END DefaultLoanSchemeID "+
	// ",ISNULL(c.Frekuesi_Pertemuan,'') as Frekuesi_Pertemuan "+
	",'W' as Frekuesi_Pertemuan "+
	",ISNULL(c.Tanggal_PKM_Pertama,'') as Tanggal_PKM_Pertama "+
	",ISNULL(c.Hari_Pertemuan,'') as Hari_Pertemuan "+
	",d.Nama_Lengkap "+
	",d.Nama_Ayah "+
	",d.Nama_Gadis_Ibu_Kandung "+
	",d.Tanggal_Lahir "+
	",d.Nomor_Identitas "+
	",d.Alamat_Domisili "+
	",ISNULL(a.Sub_Kelompok_ID,'') as SubGroup "+
	",f.Nama_Produk "+ 
	",f.Jumlah_Pinjaman "+
	",f.Term_Pembiayaan "+
	",a.ID_Prospek "+
	",a.Is_PPPass "+
	",a.Is_Sisipan "+
	",a.Is_LanjutUK "+
	// ",ISNULL(g.TGL,'') as NextMeetingDate "+
	",ISNULL(DATEADD(day, 7,c.Tanggal_PKM_Pertama),'') as NextMeetingDate "+
	",f.Siklus "+
	",a.Created_By "+
	",ISNULL(a.ID_DK,'') as ID_DK "+
	"FROM INISIASI_PROSPEK a WITH (NOLOCK) "+
	"LEFT JOIN INISIASI_MASTER_CABANG b WITH (NOLOCK) ON a.OurBranchID=b.ID_Cabang "+
	"LEFT JOIN INISIASI_DATA_KELOMPOK c WITH (NOLOCK) ON a.OurBranchID=c.OurBranchID AND CONCAT(ISNULL(a.ID_DK,''),ISNULL(a.Kelompok_ID,''))=CONCAT(ISNULL(c.ID_DK_Mobile,''),ISNULL(c.ID_Kelompok,''))    "+
	"LEFT JOIN INISIASI_UK_DATA_PRIBADI d WITH (NOLOCK) ON a.ID_Prospek=d.ID_Prospek "+
	"LEFT JOIN INISIASI_UKV_PRODUK_PEMBIAYAAN f WITH (NOLOCK) ON a.ID_Prospek=f.ID_Prospek "+
	// "LEFT JOIN dbo.NEXT_HARI_PERTEMUAN_BY_DATE(@TGL_PKM_PERTAMA) g ON c.Hari_Pertemuan=g.ID_HARI_PERTEMUAN "+
	// "WHERE a.ID_Prospek=@ID_PROSPEK and c.Tanggal_PKM_Pertama!='1900-01-01' "
	"WHERE a.ID_Prospek='"+data_add_persetujuan_pembiayaan.ID_Prospek+"' and c.Tanggal_PKM_Pertama!='1900-01-01' "


	global.Logging("SYNC_TO_BRNET","INFO models.POST_PERSETUJUAN_PEMBIAYAAN PARAM_TO_BRNET --->  "+query)

	
	start_sp_persetujuan_pembiayaan := time.Now()

	err = db.RawSQL(query).Do(&params)		
	if err != nil {		
		global.Logging("SYNC_TO_BRNET","ERROR models.POST_PERSETUJUAN_PEMBIAYAAN --->  "+query+" ---> " + err.Error())
		global.Logging("SIMPLE_SYNC_TO_BRNET","ERROR  " + err.Error())
		errsql := err
		err = db.Rollback() //Rollback Transaction
		if err != nil {return err}
		return errsql
	}		

	elapsed_sp_persetujuan_pembiayaan := time.Since(start_sp_persetujuan_pembiayaan)
	global.Logging("SYNC_TO_BRNET","INFO models.POST_PERSETUJUAN_PEMBIAYAAN EXEC ADD_PERSETUJUAN_PEMBIAYAAN dan PARAM_TO_BRNET  process_time "+elapsed_sp_persetujuan_pembiayaan.String()+" ---> ")	
	
	//======================================= SYNC_TO_BRNET =======================================//

	PROJECTION_SANCTION,err = SYNC_TO_BRNET(params[0],data_add_persetujuan_pembiayaan.ID_Prospek)
	if err != nil {					
		global.Logging("SYNC_TO_BRNET","ERROR SET_GROUP_BRNET ---> " + err.Error())
		global.Logging("SIMPLE_SYNC_TO_BRNET","ERROR " + err.Error())
		errsql := err
		err = db.Rollback() //Rollback Transaction
		if err != nil {return err}
		return errsql
	}	
	
	PROJECTION_SANCTIONjson, err := json.Marshal(PROJECTION_SANCTION)
    if err != nil {
		global.Logging("SYNC_TO_BRNET","ERROR JSON PROJECTION_SANCTION ---> " + err.Error())
		global.Logging("SIMPLE_SYNC_TO_BRNET","ERROR " + err.Error())
		errsql := err
		err = db.Rollback() //Rollback Transaction
		if err != nil {return err}
		return errsql
    }		
	global.Logging("SYNC_TO_BRNET","SUCCESS PROSPEK_ID="+data_add_persetujuan_pembiayaan.ID_Prospek+" ---> "+string(PROJECTION_SANCTIONjson))	
	global.Logging("SIMPLE_SYNC_TO_BRNET","INFO SUCCESS PROSPEK_ID="+data_add_persetujuan_pembiayaan.ID_Prospek+" Nasabah="+params[0].Nama_Lengkap)
			
	
	err = db.Commit() //Commit Transaction
	if err != nil {return err}

	err = db.Close()
	if err != nil {return err}	

	return nil	

}

type RESPONSE_DATE_FOR_BRANCH struct {
	OurBranchID string `json:"OurBranchID" db:"OurBranchID"`
	SODDate string `json:"SODDate" db:"SODDate"`
}

type RESPONSE_SET_GROUP_BRNET struct {
	GroupID	string `json:"GroupID" db:"GroupID"`
}

type RESPONSE_POST_CLIENT_BRNET struct {
	ApplicationID string `json:"ApplicationID" db:"ApplicationID"`
}

type RESPONSE_POST_CLIENT_APPROVAL_BRNET struct {
	ClientID string `json:"ClientID" db:"ClientID"`
}

type RESPONSE_POST_PROJECTION_SANCTION struct {
	AccountID string `json:"AccountID" db:"AccountID"`
	IsSelect interface{} `json:"IsSelect" db:"IsSelect"`
	OurBranchID string `json:"OurBranchID" db:"OurBranchID"`
	GroupID string `json:"GroupID" db:"GroupID"`
	GroupName string `json:"GroupName" db:"GroupName"`
	LoanSchemeID interface{} `json:"LoanSchemeID" db:"LoanSchemeID"`
	WFAdvTypeID interface{} `json:"WFAdvTypeID" db:"WFAdvTypeID"`
	WFAdvType interface{} `json:"WFAdvType" db:"WFAdvType"`
	ApplicationID string `json:"ApplicationID" db:"ApplicationID"`
	ApplicationDate interface{} `json:"ApplicationDate" db:"ApplicationDate"`
	ClientID string `json:"ClientID" db:"ClientID"`
	ClientName interface{} `json:"ClientName" db:"ClientName"`
	Term interface{} `json:"Term" db:"Term"`
	RepaymentTerm interface{} `json:"RepaymentTerm" db:"RepaymentTerm"`
	AppliedAmount interface{} `json:"AppliedAmount" db:"AppliedAmount"`
	SanctionedAmount interface{} `json:"SanctionedAmount" db:"SanctionedAmount"`
	LoanPeriod interface{} `json:"LoanPeriod" db:"LoanPeriod"`
	RepaymentFrequency interface{} `json:"RepaymentFrequency" db:"RepaymentFrequency"`
	InterestRate interface{} `json:"InterestRate" db:"InterestRate"`
	DisbursementDate interface{} `json:"DisbursementDate" db:"DisbursementDate"`
	PhotoID interface{} `json:"PhotoID" db:"PhotoID"`
	SignID interface{} `json:"SignID" db:"SignID"`
	KYCDetails interface{} `json:"KYCDetails" db:"KYCDetails"`
	DateOfBirth interface{} `json:"DateOfBirth" db:"DateOfBirth"`
	Age interface{} `json:"Age" db:"Age"`
	AgeAsOn interface{} `json:"AgeAsOn" db:"AgeAsOn"`
	RulesPending interface{} `json:"RulesPending" db:"RulesPending"`
	LoanTypeID interface{} `json:"LoanTypeID" db:"LoanTypeID"`
	LoanType interface{} `json:"LoanType" db:"LoanType"`
	OfficerID interface{} `json:"OfficerID" db:"OfficerID"`
	OfficerName interface{} `json:"OfficerName" db:"OfficerName"`
	LoanCycleNo interface{} `json:"LoanCycleNo" db:"LoanCycleNo"`
	Location interface{} `json:"Location" db:"Location"`
}

func SYNC_TO_BRNET(data_param_to_brnet PARAM_TO_BRNET,ID_Prospek string)([]RESPONSE_POST_PROJECTION_SANCTION,error){
	DATE_BRANCH := make([]RESPONSE_DATE_FOR_BRANCH, 0)
	GROUPID := make([]RESPONSE_SET_GROUP_BRNET, 0)
	CLIENT := make([]RESPONSE_POST_CLIENT_BRNET, 0)
	CLIENT_APPROVAL := make([]RESPONSE_POST_CLIENT_APPROVAL_BRNET, 0)
	PROJECTION_SANCTION := make([]RESPONSE_POST_PROJECTION_SANCTION, 0)	

	db,err := global.ConnBRNET_POST()
	if err != nil {
		return PROJECTION_SANCTION,err
	}	
	var query,groupid,clientid string

	groupid = data_param_to_brnet.Kelompok_ID
	clientid = data_param_to_brnet.ClientID

	err = db.Begin() //Begin Transaction
	if err != nil {return PROJECTION_SANCTION,err}	


	//========================= START DATE_FOR_BRANCH STEP 0 =========================//	

	query = "EXEC [DB_INISIASI].dbo.[GET_OPERATION_DATE_FOR_BRANCH] @OurBranchID='"+data_param_to_brnet.OurBranchID+"' "

	global.Logging("SYNC_TO_BRNET","INFO STEP 0 "+query+"")

	err = db.RawSQL(query).Do(&DATE_BRANCH)
	if err != nil {		
		global.Logging("SYNC_TO_BRNET","ERROR  ---> "+query+" ---> " + err.Error())	
		errsql := err
		err = db.Rollback() //Rollback Transaction
		if err != nil {return PROJECTION_SANCTION,err}
		return PROJECTION_SANCTION,errsql
	}		


	//========================= END DATE_FOR_BRANCH STEP 0 =========================//	

	if (groupid=="" && clientid==""){ //jika groupid dan clientid ada maka tidak usah masuk STEP 1

	//========================= START SET_GROUP_BRNET STEP 1 =========================//

		query = " DECLARE @TGL DATETIME "+
		" SET @TGL=GETDATE() "+
		"EXEC PNM_LIVE.dbo.p_AddEditGroupDetails "+
		"@OurBranchID=N'"+data_param_to_brnet.OurBranchID+"' "

		if (data_param_to_brnet.Kelompok_ID!=""){
			query += ",@GroupID=N'"+data_param_to_brnet.Kelompok_ID+"' "
		}else{
			query += ",@GroupID=NULL "
		}
		
		query += ",@GroupName=N'"+data_param_to_brnet.Nama_Kelompok+"' "+
		",@GroupProductID=N'"+data_param_to_brnet.GroupProductID+"' "

		if (data_param_to_brnet.DefaultLoanSchemeID!=""){
			query += ",@DefaultLoanSchemeID=N'"+data_param_to_brnet.DefaultLoanSchemeID+"' "
		}else{
			query += ",@DefaultLoanSchemeID=NULL "
		}

		query += ",@FormationDate='"+DATE_BRANCH[0].SODDate+"' "+
		// ",@FormationDate='2020-01-09 00:00:00' "+
		",@OpenDate='"+DATE_BRANCH[0].SODDate+"' "+
		// ",@OpenDate='2020-01-09 00:00:00' "+
		",@GroupClassID=N'WKGROUP' "+
		",@CreditOfficerID=N'"+data_param_to_brnet.Created_By+"' "+
		",@GroupFormedBy=N'"+data_param_to_brnet.Created_By+"' "+
		",@CityID=NULL "+
		",@CenterID=NULL "+ 
		",@VillageID=NULL "+
		",@NGOID=NULL "+
		",@MeetingFrequencyID=N'"+data_param_to_brnet.Frekuesi_Pertemuan+"' "+
		",@FirstMeetingDate='"+data_param_to_brnet.Tanggal_PKM_Pertama+"' "+
		// ",@FirstMeetingDate='2020-01-16 00:00:00' "+
		",@FirstDay=NULL "+
		",@NextDay=NULL "+
		",@RepaymentPeriodID=NULL "+
		",@MeetingDayID=N'"+data_param_to_brnet.Hari_Pertemuan+"' "+
		",@MeetingPlace=NULL "+
		",@Address=NULL "+
		",@MeetingTime="+data_param_to_brnet.Hari_Pertemuan+" "+
		",@NextMeetingDate='"+data_param_to_brnet.NextMeetingDate+"' "+
		",@RegistrationNo=NULL "+
		",@GroupLead1=NULL "+
		",@GroupStatusID=N'A' "+
		",@ClosedDate=NULL "+
		",@GrpStatusReasonID=NULL "+
		",@ModeID=NULL "+
		",@CreatedBy=N'"+data_param_to_brnet.Created_By+"' "+
		",@CreatedOn=NULL "+
		",@ModifiedBy=NULL "+
		",@ModifiedOn=NULL "+
		",@SupervisedBy=NULL "+
		",@NewRecord=1 "+
		",@MeetingPlaceCoordinate=NULL "+
		",@IsGroupTransfer=NULL "


		global.Logging("SYNC_TO_BRNET","STEP 1 INFO  "+query+"")

		start_add_group_brnet := time.Now()

		err = db.RawSQL(query).Do(&GROUPID)
		if err != nil {		
			global.Logging("SYNC_TO_BRNET","ERROR  ---> "+query+" ---> " + err.Error())	
			errsql := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {return PROJECTION_SANCTION,err}
			return PROJECTION_SANCTION,errsql
		}		

		elapsed_add_group_brnet := time.Since(start_add_group_brnet)
		global.Logging("SYNC_TO_BRNET","INFO EXEC PNM_LIVE.dbo.p_AddEditGroupDetails process_time "+elapsed_add_group_brnet.String()+" ---> ")

		groupid = GROUPID[0].GroupID
		
	//========================= END SET_GROUP_BRNET STEP 1 =========================//	

	}

	if (data_param_to_brnet.Siklus=="1"){ //jika Nasabah Siklus Lanjutan/Siklus di atas 1, maka tidak usah masuk STEP 2 dan 3

		//========================= START POST_CLIENT_BRNET STEP 2 =========================//	


		query = " DECLARE @TGL DATETIME "+
		" SET @TGL=GETDATE() "+
		" EXEC DB_INISIASI.dbo.PostClient "+
		"@OurBranchID='"+data_param_to_brnet.OurBranchID+"' "+
		",@FirstName='"+data_param_to_brnet.Nama_Lengkap+"' "+
		",@MiddleName='BINTI' "+
		",@LastName='"+data_param_to_brnet.Nama_Ayah+"' "+
		",@FatherName='"+data_param_to_brnet.Nama_Ayah+"' "+
		",@MotherName='"+data_param_to_brnet.Nama_Gadis_Ibu_Kandung+"' "+
		",@DateOfBirth='"+data_param_to_brnet.Tanggal_Lahir+"' "+
		",@ID1='"+data_param_to_brnet.Nomor_Identitas+"' "+
		",@Address='"+data_param_to_brnet.Alamat_Domisili+"' "+
		",@GroupID='"+groupid+"' "+
		",@SubGroupID='"+data_param_to_brnet.SubGroup+"' "+
		",@OpenedBy='"+data_param_to_brnet.Created_By+"' "+
		",@OpenedDate='"+DATE_BRANCH[0].SODDate+"' "+
		// ",@OpenedDate='2020-01-09 00:00:00' "+
		",@CreatedBy='"+data_param_to_brnet.Created_By+"' "+
		",@SektorEkonomi='TEST' "+
		",@JenisUsaha='TEST' "+
		",@ApprovedBy='BM' "

		global.Logging("SYNC_TO_BRNET","INFO STEP 2 "+query+"")

		start_add_client_brnet := time.Now()

		err = db.RawSQL(query).Do(&CLIENT)
		if err != nil {		
			global.Logging("SYNC_TO_BRNET","ERROR  ---> "+query+" ---> " + err.Error())	
			errsql := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {return PROJECTION_SANCTION,err}
			return PROJECTION_SANCTION,errsql
		}			
		
		elapsed_add_client_brnet := time.Since(start_add_client_brnet)
		global.Logging("SYNC_TO_BRNET","INFO EXEC DB_INISIASI.dbo.PostClient process_time "+elapsed_add_client_brnet.String()+" ---> ")		

		//========================= END POST_CLIENT_BRNET STEP 2 =========================//	


		//========================= START POST_CLIENT_APPROVAL_BRNET STEP 3 =========================//	

		query = "EXEC DB_INISIASI.dbo.PostClientApproval "+
		"@OurBranchID='"+data_param_to_brnet.OurBranchID+"' "+
		",@ID='"+CLIENT[0].ApplicationID+"' "+	
		",@ApprovedBy='BM' "

		global.Logging("SYNC_TO_BRNET","INFO STEP 3  "+query+"")

		start_add_client_approval_brnet := time.Now()

		err = db.RawSQL(query).Do(&CLIENT_APPROVAL)
		if err != nil {		
			global.Logging("SYNC_TO_BRNET","ERROR  ---> "+query+" ---> " + err.Error())	
			errsql := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {return PROJECTION_SANCTION,err}
			return PROJECTION_SANCTION,errsql
		}		

		elapsed_add_client_approval_brnet := time.Since(start_add_client_approval_brnet)
		global.Logging("SYNC_TO_BRNET","INFO EXEC DB_INISIASI.dbo.PostClientApproval process_time "+elapsed_add_client_approval_brnet.String()+" ---> ")		

		clientid = CLIENT_APPROVAL[0].ClientID

		//========================= END POST_CLIENT_APPROVAL_BRNET STEP 3 =========================//	

	}

	//========================= START POST_PROJECT_SANCTION_BRNET STEP 4 =========================//	

	query = " DECLARE @TGL DATETIME "+
	" SET @TGL=GETDATE() "+
	"EXEC DB_INISIASI.dbo.PostProjectionSanctionNew "+
	"@OurBranchID='"+data_param_to_brnet.OurBranchID+"' "+
	",@GroupID='"+groupid+"' "+	
	",@ClientID='"+clientid+"' "+	
	",@ProductID='"+data_param_to_brnet.Nama_Produk+"' "+	
	",@LoanAmount='"+data_param_to_brnet.Jumlah_Pinjaman+"' "+	
	",@Term='"+data_param_to_brnet.Term_Pembiayaan+"' "+	
	",@CreatedBy='"+data_param_to_brnet.Created_By+"' "+	
	",@ApprovedBy='BM' "+	
	",@ApprovedDate=@TGL "

	global.Logging("SYNC_TO_BRNET","INFO STEP 4 "+query+"")

	start_add_project_sanction_brnet := time.Now()	

	err = db.RawSQL(query).Do(&PROJECTION_SANCTION)
	if err != nil {		
		global.Logging("SYNC_TO_BRNET","ERROR ---> "+query+" ---> " + err.Error())		
		errsql := err
		err = db.Rollback() //Rollback Transaction
		if err != nil {return PROJECTION_SANCTION,err}
		return PROJECTION_SANCTION,errsql
	}	
	
	elapsed_add_project_sanction_brnet := time.Since(start_add_project_sanction_brnet)
	global.Logging("SYNC_TO_BRNET","INFO EXEC DB_INISIASI.dbo.PostProjectionSanction process_time "+elapsed_add_project_sanction_brnet.String()+" ---> ")	

	//========================= END POST_PROJECT_SANCTION_BRNET STEP 4 =========================//	


	

	//========================= START UPDATE_INISIASI_SETELAH_BRNET_SUCCESS STEP 5 =========================//		
	err = UPDATE_INISIASI_SETELAH_BRNET_SUCCESS(PROJECTION_SANCTION,ID_Prospek,data_param_to_brnet.ID_DK)
	if err != nil {		
		global.Logging("SYNC_TO_BRNET","ERROR UPDATE_INISIASI_SETELAH_BRNET_SUCCESS " + err.Error())		
		errsql := err
		err = db.Rollback() //Rollback Transaction
		if err != nil {return PROJECTION_SANCTION,err}
		return PROJECTION_SANCTION,errsql
	}		
	//========================= END UPDATE_INISIASI_SETELAH_BRNET_SUCCESS STEP 5 =========================//	

	err = db.Commit() //Commit Transaction
	if err != nil {return PROJECTION_SANCTION,err}

	err = db.Close()
	if err != nil {return PROJECTION_SANCTION,err}	

	return PROJECTION_SANCTION,nil		
}

func UPDATE_INISIASI_SETELAH_BRNET_SUCCESS(PROJECTION_SANCTION []RESPONSE_POST_PROJECTION_SANCTION,ID_Prospek string,ID_DK string)error{

	db,err := global.Conn()
	if err != nil {
		return err
	}	
	var query string

	start_update_pp_brnet := time.Now()	

	query = "UPDATE INISIASI_PP SET ID_Kelompok='"+PROJECTION_SANCTION[0].GroupID+"' WHERE	ID_Prospek='"+ID_Prospek+"' "
	_ , err = db.RawSQL(query).DoWithIterator()
	if err != nil {	
		global.Logging("SYNC_TO_BRNET","ERROR models.UPDATE_INISIASI_SETELAH_BRNET_SUCCESS ---> "+query+" ---> " + err.Error())	
		return err
	}		
	
	elapsed_update_pp_brnet := time.Since(start_update_pp_brnet)
	global.Logging("SYNC_TO_BRNET","INFO models.UPDATE_INISIASI_SETELAH_BRNET_SUCCESS "+query+" process_time "+elapsed_update_pp_brnet.String()+" ---> ")		

	if ID_DK!="" { //untuk nasabah baru update kelompok id berdasarkan id_dk_mobile

		start_update_prospek_dk_brnet := time.Now()	

		query = "UPDATE INISIASI_PROSPEK SET Kelompok_ID='"+PROJECTION_SANCTION[0].GroupID+"',ID_DK=NULL WHERE	ID_DK='"+ID_DK+"' "
		_ , err = db.RawSQL(query).DoWithIterator()
		if err != nil {	
			global.Logging("SYNC_TO_BRNET","ERROR ID_DK models.UPDATE_INISIASI_SETELAH_BRNET_SUCCESS ---> "+query+" ---> " + err.Error())	
			return err
		}	
		
		elapsed_update_prospek_dk_brnet := time.Since(start_update_prospek_dk_brnet)
		global.Logging("SYNC_TO_BRNET","INFO models.UPDATE_INISIASI_SETELAH_BRNET_SUCCESS ID_DK!='' "+query+" process_time "+elapsed_update_prospek_dk_brnet.String()+" ---> ")		
			
		start_update_kelompok_brnet := time.Now()

		//update id_dk_mobile jadi null, karena kelompok id sudah ada
		query = "UPDATE INISIASI_DATA_KELOMPOK SET ID_Kelompok='"+PROJECTION_SANCTION[0].GroupID+"',ID_DK_Mobile=NULL WHERE ID_DK_Mobile='"+ID_DK+"' and OurBranchID='"+PROJECTION_SANCTION[0].OurBranchID+"' "
		_ , err = db.RawSQL(query).DoWithIterator()
		if err != nil {	
			global.Logging("SYNC_TO_BRNET","ERROR models.UPDATE_INISIASI_SETELAH_BRNET_SUCCESS ---> "+query+" ---> " + err.Error())	
			return err
		}			

		elapsed_update_kelompok_brnet := time.Since(start_update_kelompok_brnet)
		global.Logging("SYNC_TO_BRNET","INFO models.UPDATE_INISIASI_SETELAH_BRNET_SUCCESS ID_DK!='' "+query+" process_time "+elapsed_update_kelompok_brnet.String()+" ---> ")				
	}	

	start_update_prospek_brnet := time.Now()

	//update semua data ke inisiasi_prospek berdsarkan id_prospek
	query = "UPDATE INISIASI_PROSPEK SET Kelompok_ID='"+PROJECTION_SANCTION[0].GroupID+"',ID_DK=NULL,Is_PPPass=1,Pass_Date=GETDATE(),ClientID='"+PROJECTION_SANCTION[0].ClientID+"',AccountID='"+PROJECTION_SANCTION[0].AccountID+"' WHERE	ID_Prospek='"+ID_Prospek+"' "
	_ , err = db.RawSQL(query).DoWithIterator()
	if err != nil {	
		global.Logging("SYNC_TO_BRNET","ERROR models.UPDATE_INISIASI_SETELAH_BRNET_SUCCESS ---> "+query+" ---> " + err.Error())	
		return err
	}		

	elapsed_update_prospek_brnet := time.Since(start_update_prospek_brnet)
	global.Logging("SYNC_TO_BRNET","INFO models.UPDATE_INISIASI_SETELAH_BRNET_SUCCESS ID_DK!='' "+query+" process_time "+elapsed_update_prospek_brnet.String()+" ---> ")					

	err = db.Close()
	if err != nil {return err}	

	return nil	
}


//========================================== END PERSETUJUAN PEMBIAYAAN && SYNC BR NET ==========================================//