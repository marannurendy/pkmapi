package models

import (
	"pkmapi/global"
	// "strconv"
	// "github.com/astaxie/beego/validation"	
	"time"	
	"fmt"	
	// "database/sql"		
	"strings"
)

type PARAM_DATA_SOSIALISASI struct {
	OurBranchID string
	CreatedBy string
	IsPickClient string
	ID_Prospek []string
	Check_ID_Prospek string
	Check_ID_By string
}

type GET_DATA_SOSIALISASI struct {
	Sosialisai []SP_GET_SOSIALISASI_MOBILE `json:"sosialisai"`
	Uk []SP_GET_UK_MOBILE `json:"uk"`
	Uk_detail []SP_GET_UK_DETAIL_MOBILE `json:"uk_detail"`
	Uk_client_data []SP_GET_UK_CLIENT_DATA_MOBILE `json:"uk_client_data"`
	Pp_kelompok []SP_GET_PP_KELOMPOK_MOBILE `json:"pp_kelompok"`
	Pp_2_kelompok []SP_GET_PP_2_KELOMPOK_MOBILE `json:"pp_2_kelompok"`
	Pp_3_kelompok []SP_GET_PP_3_KELOMPOK_MOBILE `json:"pp_3_kelompok"`
	Persetujuan_pembiayaan_kelompok []SP_GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_MOBILE `json:"persetujuan_pembiayaan_kelompok"`
	Persetujuan_pembiayaan_client_kelompok []SP_GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_CLIENT_MOBILE `json:"persetujuan_pembiayaan_client_kelompok"`
	Sosialisasi_prospek_lama []SP_GET_SOSIALISASI_MOBILE_PROSPEK_LAMA `json:"Sosialisasi_prospek_lama"`
}

type SP_GET_SOSIALISASI_MOBILE struct {
	Lokasi_Sos string `json:"Lokasi_Sos" db:"Lokasi_Sos"`
	TanggalSos string `json:"TanggalSos" db:"TanggalSos"`
	CalonNasabah string `json:"CalonNasabah" db:"CalonNasabah"`
}

func GetSosialisasiMobile(params PARAM_DATA_SOSIALISASI)([]SP_GET_SOSIALISASI_MOBILE,error) {
	data := make([]SP_GET_SOSIALISASI_MOBILE, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}		
	defer db.Close()		

	var qmain string = ""

	if params.IsPickClient=="1" {
		qmain = "DECLARE @CLIENT_DATA AS CLIENT_DATA; "

		for _, v := range params.ID_Prospek {
			qmain += " INSERT INTO @CLIENT_DATA (ID_Prospek) VALUES ('"+v+"'); "
		}
	}	

	if params.IsPickClient=="1" {
		if params.CreatedBy=="" {
			qmain += "EXEC GET_SOSIALISASI_MOBILE @OurBranchID='"+params.OurBranchID+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}else{
			qmain += "EXEC GET_SOSIALISASI_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}	
	}else{	
		if params.CreatedBy=="" {
			qmain += "EXEC GET_SOSIALISASI_MOBILE @OurBranchID='"+params.OurBranchID+"' "	
		}else{
			qmain += "EXEC GET_SOSIALISASI_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"' "	
		}			
	}	

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {		
		global.Logging("ERROR","models GetSosialisasiMobile ---> "+qmain+" ---> " + err.Error())	
		return data,err
	}			

	return data,nil
}

type SP_GET_UK_MOBILE struct {
	Lokasi_Sos string `json:"Lokasi_Sos" db:"Lokasi_Sos"`
	DATE string `json:"DATE" db:"DATE"`
}

func GetUkMobile(params PARAM_DATA_SOSIALISASI)([]SP_GET_UK_MOBILE,error) {
	data := make([]SP_GET_UK_MOBILE, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string = ""

	if params.IsPickClient=="1" {
		qmain = "DECLARE @CLIENT_DATA AS CLIENT_DATA; "

		for _, v := range params.ID_Prospek {
			qmain += " INSERT INTO @CLIENT_DATA (ID_Prospek) VALUES ('"+v+"'); "
		}
	}	

	if params.IsPickClient=="1" {
		if params.CreatedBy=="" {
			qmain += "EXEC GET_UK_MOBILE @OurBranchID='"+params.OurBranchID+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}else{
			qmain += "EXEC GET_UK_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}	
	}else{		
		if params.CreatedBy=="" {
			qmain = "EXEC GET_UK_MOBILE @OurBranchID='"+params.OurBranchID+"' "
		}else{
			qmain = "EXEC GET_UK_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"' "
		}		
	}
	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetUkMobile ---> "+qmain+" ---> " + err.Error())	
		return data,err
	}			

	return data,nil
}

type SP_GET_UK_DETAIL_MOBILE struct {
	Lokasi_Sos string `json:"Lokasi_Sos" db:"Lokasi_Sos"`
	DATE string `json:"DATE" db:"DATE"`
	Nama_Prospek string `json:"Nama_Prospek" db:"Nama_Prospek"`
}

func GetUkDetailMobile(params PARAM_DATA_SOSIALISASI)([]SP_GET_UK_DETAIL_MOBILE,error) {
	data := make([]SP_GET_UK_DETAIL_MOBILE, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string = ""

	if params.IsPickClient=="1" {
		qmain = "DECLARE @CLIENT_DATA AS CLIENT_DATA; "

		for _, v := range params.ID_Prospek {
			qmain += " INSERT INTO @CLIENT_DATA (ID_Prospek) VALUES ('"+v+"'); "
		}
	}	

	if params.IsPickClient=="1" {
		if params.CreatedBy=="" {
			qmain += "EXEC GET_UK_DETAIL_MOBILE @OurBranchID='"+params.OurBranchID+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}else{
			qmain += "EXEC GET_UK_DETAIL_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}	
	}else{		
		if params.CreatedBy=="" {
			qmain = "EXEC GET_UK_DETAIL_MOBILE @OurBranchID='"+params.OurBranchID+"' "
		}else{
			qmain = "EXEC GET_UK_DETAIL_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"' "
		}		
	}	
	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetUkDetailMobile ---> "+qmain+" ---> " + err.Error())	
		return data,err
	}			

	return data,nil
}

type SP_GET_UK_CLIENT_DATA_MOBILE struct {
	Tanggal_Sos string `json:"Tanggal_Sos" db:"Tanggal_Sos"`

	Tanggal_Input string `json:"Tanggal_Input" db:"Tanggal_Input"`
	Nama_Gadis_Ibu_Kandung string `json:"Nama_Gadis_Ibu_Kandung" db:"Nama_Gadis_Ibu_Kandung"`
	Is_Pernyataan_Dibaca string `json:"Is_Pernyataan_Dibaca" db:"Is_Pernyataan_Dibaca"`
	Is_Syariah string `json:"Is_Syariah" db:"Is_Syariah"`

	Lokasi_Sos string `json:"Lokasi_Sos" db:"Lokasi_Sos"`
	ID_Sumber string `json:"ID_Sumber" db:"ID_Sumber"`
	ID_Prospek string `json:"ID_Prospek" db:"ID_Prospek"`
	Foto_Kartu_Identitas string `json:"Foto_Kartu_Identitas" db:"Foto_Kartu_Identitas"`
	Jenis_Identitas string `json:"Jenis_Identitas" db:"Jenis_Identitas"`
	Nomor_Identitas string `json:"Nomor_Identitas" db:"Nomor_Identitas"`
	Nama_Lengkap string `json:"Nama_Lengkap" db:"Nama_Lengkap"`
	Tempat_Lahir string `json:"Tempat_Lahir" db:"Tempat_Lahir"`
	Tanggal_Lahir string `json:"Tanggal_Lahir" db:"Tanggal_Lahir"`
	Status_Perkawinan string `json:"Status_Perkawinan" db:"Status_Perkawinan"`
	Alamat_Sesuai_ID string `json:"Alamat_Sesuai_ID" db:"Alamat_Sesuai_ID"`
	Alamat_Domisili string `json:"Alamat_Domisili" db:"Alamat_Domisili"`
	Latitude interface{} `json:"Latitude" db:"Latitude"`
	Longitude interface{} `json:"Longitude" db:"Longitude"`
	Foto_Suket_Domisili string `json:"Foto_Suket_Domisili" db:"Foto_Suket_Domisili"`
	Provinsi string `json:"Provinsi" db:"Provinsi"`
	Kabupaten string `json:"Kabupaten" db:"Kabupaten"`
	Kecamatan string `json:"Kecamatan" db:"Kecamatan"`
	Kelurahan string `json:"Kelurahan" db:"Kelurahan"`

	ID_Agama string `json:"ID_Agama" db:"ID_Agama"`
	Pekerjaan string `json:"Pekerjaan" db:"Pekerjaan"`
	Jml_Tenaga_Kerja string `json:"Jml_Tenaga_Kerja" db:"Jml_Tenaga_Kerja"`



	Nama_Provinsi string `json:"Nama_Provinsi" db:"Nama_Provinsi"`
	Nama_KabKot string `json:"Nama_KabKot" db:"Nama_KabKot"`
	Nama_Kecamatan string `json:"Nama_Kecamatan" db:"Nama_Kecamatan"`
	Nama_KelurahanDesa string `json:"Nama_KelurahanDesa" db:"Nama_KelurahanDesa"`
	Foto_KK string `json:"Foto_KK" db:"Foto_KK"`
	No_KK string `json:"No_KK" db:"No_KK"`
	Nama_Ayah string `json:"Nama_Ayah" db:"Nama_Ayah"`
	No_Telp string `json:"No_Telp" db:"No_Telp"`
	Jml_Anak string `json:"Jml_Anak" db:"Jml_Anak"`
	Jml_Tanggungan string `json:"Jml_Tanggungan" db:"Jml_Tanggungan"`
	StatusRumah string `json:"StatusRumah" db:"StatusRumah"`
	Lama_Tinggal string `json:"Lama_Tinggal" db:"Lama_Tinggal"`
	Nama_Suami string `json:"Nama_Suami" db:"Nama_Suami"`
	Foto_KTP_Suami string `json:"Foto_KTP_Suami" db:"Foto_KTP_Suami"`
	Is_Ditempat string `json:"Is_Ditempat" db:"Is_Ditempat"`
	Status_Penjamin string `json:"Status_Penjamin" db:"Status_Penjamin"`
	Nama_Penjamin string `json:"Nama_Penjamin" db:"Nama_Penjamin"`
	Foto_KTP_Penjamin string `json:"Foto_KTP_Penjamin" db:"Foto_KTP_Penjamin"`
	Siklus_Pembiayaan string `json:"Siklus_Pembiayaan" db:"Siklus_Pembiayaan"`
	Siklus string `json:"Siklus" db:"Siklus"`
	Jenis_Pembiayaan string `json:"Jenis_Pembiayaan" db:"Jenis_Pembiayaan"`
	ID_Produk string `json:"ID_Produk" db:"ID_Produk"`
	Nama_Produk string `json:"Nama_Produk" db:"Nama_Produk"`
	ID_Produk_Pembiayaan string `json:"ID_Produk_Pembiayaan" db:"ID_Produk_Pembiayaan"`
	Produk_Pembiayaan string `json:"Produk_Pembiayaan" db:"Produk_Pembiayaan"`
	Jumlah_Pinjaman string `json:"Jumlah_Pinjaman" db:"Jumlah_Pinjaman"`
	Term_Pembiayaan string `json:"Term_Pembiayaan" db:"Term_Pembiayaan"`
	Kategori_Tujuan_Pembiayaan string `json:"Kategori_Tujuan_Pembiayaan" db:"Kategori_Tujuan_Pembiayaan"`
	Tujuan_Pembiayaan string `json:"Tujuan_Pembiayaan" db:"Tujuan_Pembiayaan"`
	Type_Pencairan string `json:"Type_Pencairan" db:"Type_Pencairan"`
	Frekuensi_Pembiayaan string `json:"Frekuensi_Pembiayaan" db:"Frekuensi_Pembiayaan"`
	Is_Ada_Rekening_Bank interface{} `json:"Is_Ada_Rekening_Bank" db:"Is_Ada_Rekening_Bank"`
	Nama_Bank interface{} `json:"Nama_Bank" db:"Nama_Bank"`
	No_Rekening interface{} `json:"No_Rekening" db:"No_Rekening"`
	Pemilik_Rekening interface{} `json:"Pemilik_Rekening" db:"Pemilik_Rekening"`
	Luas_Bangunan string `json:"Luas_Bangunan" db:"Luas_Bangunan"`
	Kondisi_Bangunan string `json:"Kondisi_Bangunan" db:"Kondisi_Bangunan"`
	Jenis_Atap string `json:"Jenis_Atap" db:"Jenis_Atap"`
	Dinding string `json:"Dinding" db:"Dinding"`
	Lantai string `json:"Lantai" db:"Lantai"`
	Is_AdaAksesAirBersih interface{} `json:"Is_AdaAksesAirBersih" db:"Is_AdaAksesAirBersih"`
	Is_AdaAdaToiletPribadi interface{} `json:"Is_AdaAdaToiletPribadi" db:"Is_AdaAdaToiletPribadi"`
	ID_SektorEkonomi string `json:"ID_SektorEkonomi" db:"ID_SektorEkonomi"`
	ID_SubSektorEkonomi string `json:"ID_SubSektorEkonomi" db:"ID_SubSektorEkonomi"`
	Jenis_Usaha string `json:"Jenis_Usaha" db:"Jenis_Usaha"`
	PendapatanKotor_perHari string `json:"PendapatanKotor_perHari" db:"PendapatanKotor_perHari"`
	PengeluaranKel_perHari string `json:"PengeluaranKel_perHari" db:"PengeluaranKel_perHari"`
	PendapatanBersih_perHari string `json:"PendapatanBersih_perHari" db:"PendapatanBersih_perHari"`
	JmlHariUsaha_perBulan string `json:"JmlHariUsaha_perBulan" db:"JmlHariUsaha_perBulan"`
	PendapatanBersih_perBulan string `json:"PendapatanBersih_perBulan" db:"PendapatanBersih_perBulan"`
	PendapatanBersih_perMinggu string `json:"PendapatanBersih_perMinggu" db:"PendapatanBersih_perMinggu"`
	Is_AdaPembiayaanLain string `json:"Is_AdaPembiayaanLain" db:"Is_AdaPembiayaanLain"`
	Nama_Pembiayaan_Lembaga_Lain interface{} `json:"Nama_Pembiayaan_Lembaga_Lain" db:"Nama_Pembiayaan_Lembaga_Lain"`
	Kemampuan_Angsuran string `json:"Kemampuan_Angsuran" db:"Kemampuan_Angsuran"`
	PendapatanKotor_perHari_Suami string `json:"PendapatanKotor_perHari_Suami" db:"PendapatanKotor_perHari_Suami"`
	PengeluaranKel_perHari_Suami string `json:"PengeluaranKel_perHari_Suami" db:"PengeluaranKel_perHari_Suami"`
	PendapatanBersih_perHari_Suami string `json:"PendapatanBersih_perHari_Suami" db:"PendapatanBersih_perHari_Suami"`
	JmlHariUsaha_perBulan_Suami string `json:"JmlHariUsaha_perBulan_Suami" db:"JmlHariUsaha_perBulan_Suami"`
	PendapatanBersih_perBulan_Suami string `json:"PendapatanBersih_perBulan_Suami" db:"PendapatanBersih_perBulan_Suami"`
	PendapatanBersih_perMinggu_Suami string `json:"PendapatanBersih_perMinggu_Suami" db:"PendapatanBersih_perMinggu_Suami"`
	TTD_Nasabah string `json:"TTD_Nasabah" db:"TTD_Nasabah"`
	// TTD_Suami string `json:"TTD_Suami" db:"TTD_Suami"`
	TTD_Penjamin string `json:"TTD_Penjamin" db:"TTD_Penjamin"`
	TTD_KSK string `json:"TTD_KSK" db:"TTD_KSK"`
	TTD_KK string `json:"TTD_KK" db:"TTD_KK"`
	TTD_AO string `json:"TTD_AO" db:"TTD_AO"`
	TTD_KC_SAO string `json:"TTD_KC_SAO" db:"TTD_KC_SAO"`

	Nama_TTD_Nasabah string `json:"Nama_TTD_Nasabah" db:"Nama_TTD_Nasabah"`
	Nama_TTD_Penjamin string `json:"Nama_TTD_Penjamin" db:"Nama_TTD_Penjamin"`
	Nama_TTD_KSK string `json:"Nama_TTD_KSK" db:"Nama_TTD_KSK"`
	Nama_TTD_KK string `json:"Nama_TTD_KK" db:"Nama_TTD_KK"`
	Nama_TTD_AO string `json:"Nama_TTD_AO" db:"Nama_TTD_AO"`
	Nama_TTD_KC_SAO string `json:"Nama_TTD_KC_SAO" db:"Nama_TTD_KC_SAO"`


	OurBranchID string `json:"OurBranchID" db:"OurBranchID"`
	Created_By string `json:"Created_By" db:"Created_By"`
	Is_UKPass string `json:"Is_UKPass" db:"Is_UKPass"`
	Is_VerifPass string `json:"Is_VerifPass" db:"Is_VerifPass"`
	Is_Layak string `json:"Is_Layak" db:"Is_Layak"`

	Kelompok_ID interface{} `json:"Kelompok_ID" db:"Kelompok_ID"`
	Sub_Kelompok interface{} `json:"Sub_Kelompok" db:"Sub_Kelompok"`
	Nama_Kelompok interface{} `json:"Nama_Kelompok" db:"Nama_Kelompok"`
	Is_Ketua_Kelompok string `json:"Is_Ketua_Kelompok" db:"Is_Ketua_Kelompok"`
	Is_Ketua_Sub_Kelompok string `json:"Is_Ketua_Sub_Kelompok" db:"Is_Ketua_Sub_Kelompok"`

	PostStatus interface{} `json:"PostStatus" db:"PostStatus"`
	Reason interface{} `json:"Reason" db:"Reason"`
	Tanggal_Verif interface{} `json:"Tanggal_Verif" db:"Tanggal_Verif"`
	Jasa string `json:"Jasa" db:"Jasa"`
	Angsuran_Perminggu string `json:"Angsuran_Perminggu" db:"Angsuran_Perminggu"`
	Is_Sisipan string `json:"Is_Sisipan" db:"Is_Sisipan"`
	Is_Pencairan string `json:"Is_Pencairan" db:"Is_Pencairan"`
	Is_LanjutUK string `json:"Is_LanjutUK" db:"Is_LanjutUK"`
	ID_MPP string `json:"ID_MPP" db:"ID_MPP"`

	Kualitas_KehadiranPKM string `json:"Kualitas_KehadiranPKM" db:"Kualitas_KehadiranPKM"`
	Kualitas_Pembayaran string `json:"Kualitas_Pembayaran" db:"Kualitas_Pembayaran"`
	Pendidikan_Anak_Tertinggi string `json:"Pendidikan_Anak" db:"Pendidikan_Anak_Tertinggi"`
	// Nama_Pendidikan string `json:"Nama_Pendidikan" db:"Nama_Pendidikan"`
}

func GetUkClientDataMobile(params PARAM_DATA_SOSIALISASI)([]SP_GET_UK_CLIENT_DATA_MOBILE,error) {
	data := make([]SP_GET_UK_CLIENT_DATA_MOBILE, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}		
	defer db.Close()

	var qmain string = ""

	if params.IsPickClient=="1" {
		qmain = "DECLARE @CLIENT_DATA AS CLIENT_DATA; "

		for _, v := range params.ID_Prospek {
			qmain += " INSERT INTO @CLIENT_DATA (ID_Prospek) VALUES ('"+v+"'); "
		}
	}	

	if params.IsPickClient=="1" {
		if params.CreatedBy=="" {
			qmain += "EXEC GET_UK_CLIENT_DATA_MOBILE @OurBranchID='"+params.OurBranchID+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}else{
			qmain += "EXEC GET_UK_CLIENT_DATA_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}	
	}else{		
		if params.CreatedBy=="" {
			qmain = "EXEC GET_UK_CLIENT_DATA_MOBILE @OurBranchID='"+params.OurBranchID+"' "
		}else{
			qmain = "EXEC GET_UK_CLIENT_DATA_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"' "
		}		
	}	

	start_uk_client_data_mobile := time.Now()

	// fmt.Println(qmain)
	err = db.RawSQL(qmain).Do(&data)	

	if err != nil {
		global.Logging("ERROR","models GetUkClientDataMobile ---> "+qmain+" ---> " + err.Error())	
		return data,err
	}			

	elapsed_uk_client_data_mobile := time.Since(start_uk_client_data_mobile)
	global.Logging("INFO","STORE_PROCEDURE GET_UK_CLIENT_DATA_MOBILE process_time "+elapsed_uk_client_data_mobile.String()+" ")


	start_images := time.Now()

	if len(data)>0 {
		for k, images := range data {
			// if strings.Contains(images.Foto_Kartu_Identitas,"images") == true {
			if images.Foto_Kartu_Identitas != "" {					
				// data[k].Foto_Kartu_Identitas, err = global.FileB64(images.Foto_Kartu_Identitas)
				// if err != nil {
				// 	global.Logging("ERROR models.GetUkClientDataMobile Foto_Kartu_Identitas ---> " + err.Error())
				// 	return data,err
				// }	

				data[k].Foto_Kartu_Identitas = ""
			}
			// if strings.Contains(images.Foto_Suket_Domisili,"images") == true {
			if images.Foto_Suket_Domisili != "" {								
				// data[k].Foto_Suket_Domisili, err = global.FileB64(images.Foto_Suket_Domisili)
				// if err != nil {
				// 	global.Logging("ERROR models.GetUkClientDataMobile Foto_Suket_Domisili ---> " + err.Error())					
				// 	return data,err
				// }					

				data[k].Foto_Suket_Domisili = ""
			}	
			// if strings.Contains(images.Foto_KK,"images") == true {
			if images.Foto_KK != "" {								
				// data[k].Foto_KK, err = global.FileB64(images.Foto_KK)
				// if err != nil {
				// 	global.Logging("ERROR models.GetUkClientDataMobile Foto_KK ---> " + err.Error())					
				// 	return data,err
				// }				
				
				data[k].Foto_KK = ""
			}
			// if strings.Contains(images.Foto_KTP_Suami,"images") == true {
			if images.Foto_KTP_Suami != "" {			
				// data[k].Foto_KTP_Suami, err = global.FileB64(images.Foto_KTP_Suami)
				// if err != nil {
				// 	global.Logging("ERROR models.GetUkClientDataMobile Foto_KTP_Suami ---> " + err.Error())					
				// 	return data,err
				// }			
				
				data[k].Foto_KTP_Suami = ""
			}
			// if strings.Contains(images.Foto_KTP_Penjamin,"images") == true {	
			if images.Foto_KTP_Penjamin != "" {	
				// data[k].Foto_KTP_Penjamin, err = global.FileB64(images.Foto_KTP_Penjamin)
				// if err != nil {
				// 	global.Logging("ERROR models.GetUkClientDataMobile Foto_KTP_Penjamin ---> " + err.Error())					
				// 	return data,err
				// }					

				data[k].Foto_KTP_Penjamin = ""
			}			
			// if strings.Contains(images.TTD_Nasabah,"images") == true {	
			if images.TTD_Nasabah != "" {		
				// data[k].TTD_Nasabah, err = global.FileB64(images.TTD_Nasabah)
				// if err != nil {
				// 	global.Logging("ERROR models.GetUkClientDataMobile TTD_Nasabah ---> " + err.Error())					
				// 	return data,err
				// }			
				
				data[k].TTD_Nasabah = ""
			}				
			// if strings.Contains(images.TTD_KSK,"images") == true {	
			if images.TTD_KSK != "" {		
				// data[k].TTD_KSK, err = global.FileB64(images.TTD_KSK)
				// if err != nil {
				// 	global.Logging("ERROR models.GetUkClientDataMobile TTD_KSK ---> " + err.Error())					
				// 	return data,err
				// }				
				
				data[k].TTD_KSK = ""
			}		
			// if strings.Contains(images.TTD_KK,"images") == true {	
			if images.TTD_KK != "" {		
				// data[k].TTD_KK, err = global.FileB64(images.TTD_KK)
				// if err != nil {
				// 	global.Logging("ERROR models.GetUkClientDataMobile TTD_KK ---> " + err.Error())					
				// 	return data,err
				// }	
				
				data[k].TTD_KK = ""
			}	
			// if strings.Contains(images.TTD_Penjamin,"images") == true {	
			if images.TTD_Penjamin != "" {	
				// data[k].TTD_Penjamin, err = global.FileB64(images.TTD_Penjamin)
				// if err != nil {
				// 	global.Logging("ERROR models.GetUkClientDataMobile TTD_Penjamin ---> " + err.Error())					
				// 	return data,err
				// }					
				data[k].TTD_Penjamin = ""
			}								
			// if strings.Contains(images.TTD_AO,"images") == true {	
			if images.TTD_AO != "" {					
				// data[k].TTD_AO, err = global.FileB64(images.TTD_AO)
				// if err != nil {
				// 	global.Logging("ERROR models.GetUkClientDataMobile TTD_AO ---> " + err.Error())					
				// 	return data,err
				// }			
				data[k].TTD_AO = ""
			}				
			// if strings.Contains(images.TTD_KC_SAO,"images") == true {	
			if images.TTD_KC_SAO != "" {	
				// data[k].TTD_KC_SAO, err = global.FileB64(images.TTD_KC_SAO)
				// if err != nil {
				// 	global.Logging("ERROR models.GetUkClientDataMobile TTD_KC_SAO ---> " + err.Error())					
				// 	return data,err
				// }			
				data[k].TTD_KC_SAO = ""
			}				
		}
	}
	
	elapsed_images := time.Since(start_images)
	global.Logging("INFO","models.GetUkClientDataMobile images process_time "+elapsed_images.String()+" ")	

	return data,nil
}

type SP_GET_PP_KELOMPOK_MOBILE struct {
	Nama_Kelompok string `json:"Nama_Kelompok" db:"Nama_Kelompok"`
	Total string `json:"Total" db:"Total"`
	Created_Date string `json:"Created_Date" db:"Created_Date"`
}

func GetPpKelompokMobile(params PARAM_DATA_SOSIALISASI)([]SP_GET_PP_KELOMPOK_MOBILE,error) {
	data := make([]SP_GET_PP_KELOMPOK_MOBILE, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string = ""

	if params.IsPickClient=="1" {
		qmain = "DECLARE @CLIENT_DATA AS CLIENT_DATA; "

		for _, v := range params.ID_Prospek {
			qmain += " INSERT INTO @CLIENT_DATA (ID_Prospek) VALUES ('"+v+"'); "
		}
	}	

	if params.IsPickClient=="1" {
		if params.CreatedBy=="" {
			qmain += "EXEC GET_PP_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}else{
			qmain += "EXEC GET_PP_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}	
	}else{		
		if params.CreatedBy=="" {
			qmain = "EXEC GET_PP_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"' "
		}else{
			qmain = "EXEC GET_PP_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"' "
		}		
	}	

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetPpKelompokMobile ---> "+qmain+" ---> " + err.Error())	
		return data,err
	}			

	return data,nil
}

type SP_GET_PP_2_KELOMPOK_MOBILE struct {
	Nama_Kelompok string `json:"Nama_Kelompok" db:"Nama_Kelompok"`
	Total string `json:"Total" db:"Total"`
	Created_Date string `json:"Created_Date" db:"Created_Date"`
}

func GetPp2KelompokMobile(params PARAM_DATA_SOSIALISASI)([]SP_GET_PP_2_KELOMPOK_MOBILE,error) {
	data := make([]SP_GET_PP_2_KELOMPOK_MOBILE, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string = ""

	if params.IsPickClient=="1" {
		qmain = "DECLARE @CLIENT_DATA AS CLIENT_DATA; "

		for _, v := range params.ID_Prospek {
			qmain += " INSERT INTO @CLIENT_DATA (ID_Prospek) VALUES ('"+v+"'); "
		}
	}	

	if params.IsPickClient=="1" {
		if params.CreatedBy=="" {
			qmain += "EXEC GET_PP_2_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}else{
			qmain += "EXEC GET_PP_2_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}	
	}else{		
		if params.CreatedBy=="" {
			qmain = "EXEC GET_PP_2_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"' "
		}else{
			qmain = "EXEC GET_PP_2_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"' "
		}		
	}	

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetPp2KelompokMobile ---> "+qmain+" ---> " + err.Error())			
		return data,err
	}			

	return data,nil
}

type SP_GET_PP_3_KELOMPOK_MOBILE struct {
	Nama_Kelompok string `json:"Nama_Kelompok" db:"Nama_Kelompok"`
	Total string `json:"Total" db:"Total"`
	Created_Date string `json:"Created_Date" db:"Created_Date"`
}

func GetPp3KelompokMobile(params PARAM_DATA_SOSIALISASI)([]SP_GET_PP_3_KELOMPOK_MOBILE,error) {
	data := make([]SP_GET_PP_3_KELOMPOK_MOBILE, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string = ""

	if params.IsPickClient=="1" {
		qmain = "DECLARE @CLIENT_DATA AS CLIENT_DATA; "

		for _, v := range params.ID_Prospek {
			qmain += " INSERT INTO @CLIENT_DATA (ID_Prospek) VALUES ('"+v+"'); "
		}
	}	

	if params.IsPickClient=="1" {
		if params.CreatedBy=="" {
			qmain += "EXEC GET_PP_3_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}else{
			qmain += "EXEC GET_PP_3_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}	
	}else{		
		if params.CreatedBy=="" {
			qmain = "EXEC GET_PP_3_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"' "
		}else{
			qmain = "EXEC GET_PP_3_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"' "
		}		
	}	

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetPp3KelompokMobile ---> "+qmain+" ---> " + err.Error())					
		return data,err
	}			

	return data,nil
}

type SP_GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_MOBILE struct {
	Kelompok_ID string `json:"Kelompok_ID" db:"Kelompok_ID"`
	Nama_Kelompok string `json:"Nama_Kelompok" db:"Nama_Kelompok"`
	Total string `json:"Total" db:"Total"`
}

func GetPersetujuanPembiayaanKelompokMobile(params PARAM_DATA_SOSIALISASI)([]SP_GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_MOBILE,error) {
	data := make([]SP_GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_MOBILE, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string = ""

	if params.IsPickClient=="1" {
		qmain = "DECLARE @CLIENT_DATA AS CLIENT_DATA; "

		for _, v := range params.ID_Prospek {
			qmain += " INSERT INTO @CLIENT_DATA (ID_Prospek) VALUES ('"+v+"'); "
		}
	}			

	if params.IsPickClient=="1" {
		if params.CreatedBy=="" {
			qmain += "EXEC GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}else{
			qmain += "EXEC GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}	
	}else{		
		if params.CreatedBy=="" {
			qmain = "EXEC GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"' "
		}else{
			qmain = "EXEC GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"' "
		}		
	}		

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetPersetujuanPembiayaanKelompokMobile ---> "+qmain+" ---> " + err.Error())					
		return data,err
	}			

	return data,nil
}

type SP_GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_CLIENT_MOBILE struct {
	Kelompok_ID string `json:"Kelompok_ID" db:"Kelompok_ID"`
	Nama_Prospek string `json:"Nama_Prospek" db:"Nama_Prospek"`
	Jumlah_Pinjaman string `json:"Jumlah_Pinjaman" db:"Jumlah_Pinjaman"`
	Term_Pembiayaan string `json:"Term_Pembiayaan" db:"Term_Pembiayaan"`
	Jasa string `json:"Jasa" db:"Jasa"`
	Angsuran_Per_Minggu string `json:"Angsuran_Per_Minggu" db:"Angsuran_Per_Minggu"`
	TTD_AO string `json:"TTD_AO" db:"TTD_AO"`
	TTD_KC string `json:"TTD_KC" db:"TTD_KC"`	
}
func GetPersetujuanPembiayaanKelompokClientMobile(params PARAM_DATA_SOSIALISASI)([]SP_GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_CLIENT_MOBILE,error) {
	data := make([]SP_GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_CLIENT_MOBILE, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string = ""

	if params.IsPickClient=="1" {
		qmain = "DECLARE @CLIENT_DATA AS CLIENT_DATA; "

		for _, v := range params.ID_Prospek {
			qmain += " INSERT INTO @CLIENT_DATA (ID_Prospek) VALUES ('"+v+"'); "
		}
	}	

	if params.IsPickClient=="1" {
		if params.CreatedBy=="" {
			qmain += "EXEC GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_CLIENT_MOBILE @OurBranchID='"+params.OurBranchID+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}else{
			qmain += "EXEC GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_CLIENT_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}	
	}else{		
		if params.CreatedBy=="" {
			qmain = "EXEC GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_CLIENT_MOBILE @OurBranchID='"+params.OurBranchID+"' "
		}else{
			qmain = "EXEC GET_PERSETUJUAN_PEMBIAYAAN_KELOMPOK_CLIENT_MOBILE @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"' "
		}		
	}		

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetPersetujuanPembiayaanKelompokClientMobile ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	if len(data)>0 {
		for k, images := range data {
			if strings.Contains(images.TTD_AO,"images") == true {	
				data[k].TTD_AO, _ = global.FileB64("./"+images.TTD_AO)
			}
			if strings.Contains(images.TTD_KC,"images") == true {	
				data[k].TTD_KC, _ = global.FileB64("./"+images.TTD_KC)
			}			
		}
	}

	return data,nil
}

type SP_GET_LIST_OF_CLIENT struct {
	ID_Prospek string `json:"ID_Prospek" db:"ID_Prospek"`
	Nama string `json:"Nama" db:"Nama"`
}

func GetListOfClient(params map[string]string)([]SP_GET_LIST_OF_CLIENT,error) {
	data := make([]SP_GET_LIST_OF_CLIENT, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	if params["CreatedBy"]=="undefined" {
		qmain = "EXEC GET_LIST_OF_CLIENT @OurBranchID='"+params["OurBranchID"]+"' "
	}else{
		qmain = "EXEC GET_LIST_OF_CLIENT @OurBranchID='"+params["OurBranchID"]+"',@CreatedBy='"+params["CreatedBy"]+"' "
	}	

	if params["Search"]!="undefined" {
		qmain += " ,@Search = '"+params["Search"]+"' "
	}

	qmain += " ,@PAGE = '"+params["PAGE"]+"',@LIMIT = '"+params["LIMIT"]+"' "

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetListOfClient ---> "+qmain+" ---> " + err.Error())		
		return data,err
	}			

	return data,nil
}

type SP_GET_SCORING_STATUS struct {
	ID_Prospek string `json:"ID_Prospek" db:"ID_Prospek"`
	Status_Kelayakan interface{} `json:"Status_Kelayakan" db:"Status_Kelayakan"`
}

func GetScoringStatus(id_prospek string) ([]SP_GET_SCORING_STATUS,error) {
	data := make([]SP_GET_SCORING_STATUS, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "EXEC GET_SCORING_STATUS @ID_Prospek='"+id_prospek+"' "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetScoringStatus ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type GET_DATA_PENCAIRAN struct {
	Pencairan []SP_GET_PENCAIRAN_MOBILE `json:"pencairan"`
	Pencairan_nasabah []SP_GET_PENCAIRAN_NASABAH_MOBILE `json:"pencairan_nasabah"`
}

type SP_GET_PENCAIRAN_MOBILE struct {
	Kelompok_ID string `json:"Kelompok_ID" db:"Kelompok_ID"`
	Nama_Kelompok string `json:"Nama_Kelompok" db:"Nama_Kelompok"`
	Jml_ID_Prospek string `json:"Jml_ID_Prospek" db:"Jml_ID_Prospek"`
}

func GetPencairanMobile(params map[string]string)([]SP_GET_PENCAIRAN_MOBILE,error) {
	data := make([]SP_GET_PENCAIRAN_MOBILE, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}		
	defer db.Close()

	var qmain string

	if params["CreatedBy"]=="undefined" {
		qmain = "EXEC dbo.[GET_PENCAIRAN_MOBILE] @OurBranchID='"+params["OurBranchID"]+"' "	
	}else{
		qmain = "EXEC GET_PENCAIRAN_MOBILE @OurBranchID='"+params["OurBranchID"]+"',@CreatedBy='"+params["CreatedBy"]+"' "	
	}	

	// fmt.Println(qmain)

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetPencairanMobile ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil
}

type SP_GET_PENCAIRAN_NASABAH_MOBILE struct {
	Kelompok_ID string `json:"Kelompok_ID" db:"Kelompok_ID"`
	Nama_Kelompok string `json:"Nama_Kelompok" db:"Nama_Kelompok"`
	Nama_Prospek string `json:"Nama_Prospek" db:"Nama_Prospek"`
	Alamat_Domisili string `json:"Alamat_Domisili" db:"Alamat_Domisili"`
	Nomor_Identitas string `json:"Nomor_Identitas" db:"Nomor_Identitas"`
	Nama_Tipe_Pencairan string `json:"Nama_Tipe_Pencairan" db:"Nama_Tipe_Pencairan"`
	Nama_Penjamin string `json:"Nama_Penjamin" db:"Nama_Penjamin"`
	Jumlah_Pinjaman string `json:"Jumlah_Pinjaman" db:"Jumlah_Pinjaman"`
	Nama_Produk string `json:"Nama_Produk" db:"Nama_Produk"`
	Jasa string `json:"Jasa" db:"Jasa"`
	Term_Pembiayaan string `json:"Term_Pembiayaan" db:"Term_Pembiayaan"`
	Angsuran_Per_Minggu string `json:"Angsuran_Per_Minggu" db:"Angsuran_Perminggu"`
	// TTD_Nasabah interface{} `json:"TTD_Nasabah" db:"TTD_Nasabah"`
	// TTD_KC interface{} `json:"TTD_KC" db:"TTD_KC"`
	// TTD_Nasabah_2 interface{} `json:"TTD_Nasabah_2" db:"TTD_Nasabah_2"`
	// TTD_KSK interface{} `json:"TTD_KSK" db:"TTD_KSK"`
	// TTD_KK interface{} `json:"TTD_KK" db:"TTD_KK"`
	// Foto_Pencairan interface{} `json:"Foto_Pencairan" db:"Foto_Pencairan"`
	// LRP_TTD_Nasabah interface{} `json:"LRP_TTD_Nasabah" db:"LRP_TTD_Nasabah"`
	// LRP_TTD_AO interface{} `json:"LRP_TTD_AO" db:"LRP_TTD_AO"`
	ClientID interface{} `json:"ClientID" db:"ClientID"`
	ID_Prospek string `json:"ID_Prospek" db:"ID_Prospek"`
}

func GetPencairanNasabahMobile(params map[string]string)([]SP_GET_PENCAIRAN_NASABAH_MOBILE,error) {
	data := make([]SP_GET_PENCAIRAN_NASABAH_MOBILE, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}		
	defer db.Close()

	var qmain string

	if params["CreatedBy"]=="undefined" {
		qmain = "EXEC GET_PENCAIRAN_NASABAH_MOBILE @OurBranchID='"+params["OurBranchID"]+"' "	
	}else{
		qmain = "EXEC GET_PENCAIRAN_NASABAH_MOBILE @OurBranchID='"+params["OurBranchID"]+"',@CreatedBy='"+params["CreatedBy"]+"' "	
	}	
	
	// fmt.Println(qmain)

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetPencairanNasabahMobile ---> "+qmain+" ---> " + err.Error())		
		return data,err
	}			

	return data,nil
}

type SP_GET_SOSIALISASI_MOBILE_PROSPEK_LAMA struct {
	Nama_Lengkap string `json:"Nama_Lengkap" db:"Nama_Lengkap"`
	No_Telp string `json:"No_Telp" db:"No_Telp"`
	Last_Status string `json:"Last_Status" db:"Last_Status"`
}

func GetSosialisasiMobileProspekLama(params PARAM_DATA_SOSIALISASI)([]SP_GET_SOSIALISASI_MOBILE_PROSPEK_LAMA,error){
	data := make([]SP_GET_SOSIALISASI_MOBILE_PROSPEK_LAMA, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}		
	defer db.Close()

	var qmain string = ""

	if params.IsPickClient=="1" {
		qmain = "DECLARE @CLIENT_DATA AS CLIENT_DATA; "

		for _, v := range params.ID_Prospek {
			qmain += " INSERT INTO @CLIENT_DATA (ID_Prospek) VALUES ('"+v+"'); "
		}
	}	

	if params.IsPickClient=="1" {
		if params.CreatedBy=="" {
			qmain += "EXEC GET_SOSIALISASI_MOBILE_PROSPEK_LAMA @OurBranchID='"+params.OurBranchID+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}else{
			qmain += "EXEC GET_SOSIALISASI_MOBILE_PROSPEK_LAMA @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"',@IsPickClient='1',@ID_Prospek = @CLIENT_DATA "	
		}	
	}else{	
		if params.CreatedBy=="" {
			qmain += "EXEC GET_SOSIALISASI_MOBILE_PROSPEK_LAMA @OurBranchID='"+params.OurBranchID+"' "	
		}else{
			qmain += "EXEC GET_SOSIALISASI_MOBILE_PROSPEK_LAMA @OurBranchID='"+params.OurBranchID+"',@CreatedBy='"+params.CreatedBy+"' "	
		}			
	}	

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetSosialisasiMobileProspekLama ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil
}


type GetMasterDataResponse struct {
	ResponseCode        string      `json:"responseCode"`
	ResponseDescription string      `json:"responseDescription"`
	Status              string      `json:"status"`
	Data                *MasterData `json:"data"`
}


type MasterData struct {
	Absent             		 []MasterAbsent             		`json:"absent"`
	Religion           		 []MasterReligion           		`json:"religion"`
	Branch             		 []MasterBranch             		`json:"branch"`
	LivingType         		 []MasterLivingType         		`json:"livingType"`
	IdentityType       		 []MasterIdentityType       		`json:"identityType"`
	PartnerJob         		 []MasterPartnerJob         		`json:"partnerJob"`
	DwellingCondition  		 []MasterDwellingCondition  		`json:"dwellingCondition"`
	ResidenceLocation  		 []MasterResidenceLocation  		`json:"residenceLocation"`
	PembiayaanLain     		 []MasterPembiayaanLain     		`json:"pembiayaanLain"`
	Education          		 []MasterEducation          		`json:"education"`
	Product            		 []MasterProduct            		`json:"product"`
	EconomicSector     		 []MasterEconomicSector     		`json:"economicSector"`
	RelationStatus     		 []MasterRelationStatus     		`json:"relationStatus"`
	MarriageStatus     		 []MasterMarriageStatus     		`json:"marriageStatus"`
	HomeStatus         		 []MasterHomeStatus         		`json:"homeStatus"`
	Referral           		 []MasterReferral           		`json:"referral"`
	TransFund          		 []MasterTransFund          		`json:"transFund"`
	JenisPembiayaan	   		 []MasterJenisPembiayaan	 		`json:"jenisPembiayaan"`
	SubJenisPembiayaan 		 []MasterSubJenisPembiayaan  		`json:"subjenisPembiayaan"`
	TujuanPembiayaan  		 []MasterTujuanPembiayaan			`json:"tujuanPembiayaan"`
	KategoriTujuanPembiayaan []MasterKategoriTujuanPembiayaan	`json:"kategoritujuanPembiayaan"`
	Frekuensi 		 		 []MasterFrekuensi					`json:"Frekuensi"`
	WilayahMobile 			 []MasterWilayahMobile				`json:"WilayahMobile"`
	AvailableSubGroup	 	 []MasterAvailableSubGroup			`json:"MasterAvailableSubGroup"`
	GroupProduct			 []MasterGroupProduct				`json:"MasterGroupProduct"`
	SetUKtimeOut			 string								`json:"SetUKtimeOut"`
}


type MasterAbsent struct {
	Id           string `json:"id" db:"ID_Absen"`
	AbsentDetail string `json:"absent" db:"Keterangan_Absen"`
}

func GetMasterAbsent() ([]MasterAbsent,error) {
	data := make([]MasterAbsent, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "EXEC GET_MASTER_ABSEN "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterAbsent ---> "+qmain+" ---> " + err.Error())		
		return data,err
	}			

	return data,nil	
}

type MasterReligion struct {
	Id             string `json:"id" db:"ID_Agama"`
	ReligionDetail string `json:"religion" db:"Nama_Agama"`
}

func GetMasterReligion() ([]MasterReligion,error) {
	data := make([]MasterReligion, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "EXEC GET_MASTER_AGAMA "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterReligion ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterBranch struct {
	Id           string `json:"id" db:"ID_Daerah_Tempat_Tinggal"`
	BranchDetail string `json:"branchDetail" `
}

func GetMasterBranch() ([]MasterBranch,error) {
	data := make([]MasterBranch, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()


	var qmain string

	qmain = "EXEC GET_MASTER_CABANG @type = 'CabangAsal', @OurBranchID = '' "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterBranch ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterLivingType struct {
	Id               string `json:"id" db:"ID_Daerah_Tempat_Tinggal"`
	LivingTypeDetail string `json:"livingTypeDetail" db:"Nama_Daerah_Tempat_Tinggal"`
}


func GetMasterLivingType() ([]MasterLivingType,error) {
	data := make([]MasterLivingType, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()


	var qmain string

	qmain = "SELECT ID_Daerah_Tempat_Tinggal, Nama_Daerah_Tempat_Tinggal from INISIASI_MASTER_DAERAH_TEMPAT_TINGGAL with (NOLOCK)"
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterLivingType ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterIdentityType struct {
	Id           string `json:"id" db:"ID_Jenis_Kartu_Identitas"`
	IdentityCard string `json:"identityCard" db:"Nama_Jenis_Kartu_Identitas"`
}

func GetMasterIdentityType() ([]MasterIdentityType,error) {
	data := make([]MasterIdentityType, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT ID_Jenis_Kartu_Identitas, Nama_Jenis_Kartu_Identitas from INISIASI_MASTER_JENIS_KARTU_IDENTITAS with (NOLOCK)"
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterIdentityType ---> "+qmain+" ---> " + err.Error())	
		return data,err
	}			

	return data,nil	
}

type MasterPartnerJob struct {
	Id         string `json:"id" db:"ID_Jenis_Pekerjaan_Suami"`
	PartnerJob string `json:"partnerJob" db:"Nama_Jenis_Pekerjaan_Suami"`
}

func GetMasterPartnerJob() ([]MasterPartnerJob,error) {
	data := make([]MasterPartnerJob, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT ID_Jenis_Pekerjaan_Suami, Nama_Jenis_Pekerjaan_Suami from INISIASI_MASTER_JENIS_PEKERJAAN_SUAMI with (NOLOCK)"
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterPartnerJob ---> "+qmain+" ---> " + err.Error())	
		return data,err
	}			

	return data,nil	
}

type MasterDwellingCondition struct {
	Id             string `json:"id" db:"ID_KondisiRumah"`
	Category       string `json:"category" db:"Nama_Kondisi"`
	CategoryDetail string `json:"categoryDetail" db:"Nama_SubKondisi"`
	Score          string `json:"score" db:"Nilai_SubKondisi"`
}

func GetMasterDwellingCondition() ([]MasterDwellingCondition,error) {
	data := make([]MasterDwellingCondition, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "EXEC GET_MASTER_KONDISI_RUMAH "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterDwellingCondition ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterResidenceLocation struct {
	Id             string `json:"id" db:"ID_Lokasi_Tempat_Tinggal"`
	LocationDetail string `json:"locationDetail" db:"Nama_Lokasi_Tempat_Tinggal"`
}

func GetMasterResidenceLocation() ([]MasterResidenceLocation,error) {
	data := make([]MasterResidenceLocation, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT ID_Lokasi_Tempat_Tinggal, Nama_Lokasi_Tempat_Tinggal from INISIASI_MASTER_LOKASI_TEMPAT_TINGGAL with (NOLOCK)"
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterResidenceLocation ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterPembiayaanLain struct {
	Id                    string `json:"id" db:"ID_Pembiyaan_Lembaga_Lain"`
	JenisPembiayaanDetail string `json:"jenisPembiayaanDetail" db:"Nama_Pembiayan_Lembaga_Lain"`
}

func GetMasterPembiayaanLain() ([]MasterPembiayaanLain,error) {
	data := make([]MasterPembiayaanLain, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT ID_Pembiyaan_Lembaga_Lain, Nama_Pembiayan_Lembaga_Lain from INISIASI_MASTER_PEMBIAYAAN_LEMBAGA_LAIN with (NOLOCK)"
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterPembiayaanLain ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterEducation struct {
	Id              string `json:"id" db:"ID_Pendidikan"`
	EducationDetail string `json:"educationDetail" db:"Nama_Pendidikan"`
}

func GetMasterEducation() ([]MasterEducation,error) {
	data := make([]MasterEducation, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT ID_Pendidikan, Nama_Pendidikan from INISIASI_MASTER_PENDIDIKAN with (NOLOCK)"
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterEducation ---> "+qmain+" ---> " + err.Error())	
		return data,err
	}			

	return data,nil	
}

type MasterProduct struct {
	Id          string `json:"id" db:"ID_Produk"`
	ProductName string `json:"productName" db:"Nama_Produk"`
	IsReguler   string `json:"isReguler" db:"Is_Reguler"`
	Is_MP   	string `json:"IsMP" db:"Is_MP"`
	MinPlafond  string `json:"minPlafond" db:"Min_Plafon"`
	MaxPlafond  string `json:"maxPlafond" db:"Max_Plafon"`
	PaymentTerm string `json:"paymentTerm" db:"Term_Pembiayaan"`
	Interest    string `json:"interest" db:"Bunga"`
	IsSyariah   string `json:"isSyariah" db:"Is_Syariah"`	
}

func GetMasterProduct() ([]MasterProduct,error) {
	data := make([]MasterProduct, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT ID_Produk, Nama_Produk, Is_Reguler, Is_MP, Min_Plafon, Max_Plafon, Term_Pembiayaan, Bunga, Is_Syariah from INISIASI_MASTER_PRODUK with (NOLOCK)"
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterProduct ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterEconomicSector struct {
	Id                   string `json:"id" db:"ID_SektorEkonomi"`
	EconomicSectorDetail string `json:"economicSectorDetail" db:"Nama_SektorEkonomi"`
	IdSubSector          string `json:"idSubsector" db:"ID_SubSektorEkonomi"`
	SubSectorDetail      string `json:"subSectorDetail" db:"Nama_SubSektorEkonomi"`
}

func GetMasterEconomicSector() ([]MasterEconomicSector,error) {
	data := make([]MasterEconomicSector, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT ID_SektorEkonomi, Nama_SektorEkonomi, ID_SubSektorEkonomi, Nama_SubSektorEkonomi from INISIASI_MASTER_SEKTOR_EKONOMI with (NOLOCK) "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterEconomicSector ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterRelationStatus struct {
	Id             string `json:"id" db:"ID_Status_Hubungan_Keluarga"`
	RelationStatus string `json:"relationStatus" db:"Nama_Status_Hubungan_Keluarga"`
}


func GetMasterRelationStatus() ([]MasterRelationStatus,error) {
	data := make([]MasterRelationStatus, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT ID_Status_Hubungan_Keluarga, Nama_Status_Hubungan_Keluarga from INISIASI_MASTER_STATUS_HUBUNGAN_KELUARGA with (NOLOCK) "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {		
		global.Logging("ERROR","models GetMasterRelationStatus ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterMarriageStatus struct {
	Id                   string `json:"id" db:"ID_Status_Perkawinan"`
	MarriageStatusDetail string `json:"marriageStatusDetail" db:"Nama_Status_Perkawinan"`
}

func GetMasterMarriageStatus() ([]MasterMarriageStatus,error) {
	data := make([]MasterMarriageStatus, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT ID_Status_Perkawinan, Nama_Status_Perkawinan from INISIASI_MASTER_STATUS_PERKAWINAN with (NOLOCK) "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterMarriageStatus ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterHomeStatus struct {
	Id               string `json:"id" db:"ID_Status_Rumah_Tinggal"`
	HomeStatusDetail string `json:"homeStatusDetail" db:"Nama_Status_Rumah_Tinggal"`
}

func GetMasterHomeStatus() ([]MasterHomeStatus,error) {
	data := make([]MasterHomeStatus, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT ID_Status_Rumah_Tinggal, Nama_Status_Rumah_Tinggal from INISIASI_MASTER_STATUS_RUMAH_TINGGAL with (NOLOCK) "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {		
		global.Logging("ERROR","models GetMasterHomeStatus ---> "+qmain+" ---> " + err.Error())		
		return data,err
	}			

	return data,nil	
}

type MasterReferral struct {
	Id             string `json:"id" db:"ID_Sumber"`
	ReferralDetail string `json:"referralDetail" db:"Nama_Sumber"`
}

func GetMasterReferral() ([]MasterReferral,error) {
	data := make([]MasterReferral, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT ID_Sumber, Nama_Sumber from INISIASI_MASTER_SUMBER with (NOLOCK) "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterReferral ---> "+qmain+" ---> " + err.Error())		
		return data,err
	}			

	return data,nil	
}

type MasterTransFund struct {
	Id              string `json:"id" db:"ID_Tipe_Pencairan"`
	TransfundDetail string `json:"transfundDetail" db:"Nama_Tipe_Pencairan"`
}

func GetMasterTransFund() ([]MasterTransFund,error) {
	data := make([]MasterTransFund, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT ID_Tipe_Pencairan, Nama_Tipe_Pencairan from INISIASI_MASTER_TIPE_PENCAIRAN with (NOLOCK)"
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterTransFund ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterJenisPembiayaan struct {
	ID_Jenis_Pembiayaan string `json:"id" db:"ID_Jenis_Pembiayaan"`
	Nama_Jenis_Pembiayaan string `json:"namajenispembiayaan" db:"Nama_Jenis_Pembiayaan"`
}

func GetMasterJenisPembiayaan() ([]MasterJenisPembiayaan,error) {
	data := make([]MasterJenisPembiayaan, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "EXEC GET_MASTER_JENIS_PEMBIAYAAN "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterJenisPembiayaan ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterSubJenisPembiayaan struct {
	ID_Sub_Jenis_Pembiayaan string `json:"id" db:"ID_Sub_Jenis_Pembiayaan"`
	ID_Jenis_Pembiayaan string `json:"idjenispembiayaan" db:"ID_Jenis_Pembiayaan"`
	Nama_Jenis_Pembiayaan string `json:"namajenispembiayaan" db:"Nama_Jenis_Pembiayaan"`
}

func GetMasterSubJenisPembiayaan() ([]MasterSubJenisPembiayaan,error) {
	data := make([]MasterSubJenisPembiayaan, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "EXEC GET_MASTER_SUB_JENIS_PEMBIAYAAN "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterSubJenisPembiayaan ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterTujuanPembiayaan struct {
	ID_Tujuan_Pembiayaan string `json:"id" db:"ID_Tujuan_Pembiayaan"`
	Nama_Tujuan_Pembiayaan string `json:"namatujuanpembiayaan" db:"Nama_Tujuan_Pembiayaan"`
}

func GetMasterTujuanPembiayaan() ([]MasterTujuanPembiayaan,error) {
	data := make([]MasterTujuanPembiayaan, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "EXEC GET_MASTER_TUJUAN_PEMBIAYAAN "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterTujuanPembiayaan ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterKategoriTujuanPembiayaan struct {
	ID_Kategori_Tujuan_Pembiayaan string `json:"id" db:"ID_Kategori_Tujuan_Pembiayaan"`
	Nama_Kategori_Tujuan_Pembiayaan string `json:"namakategoritujuanpembiayaan" db:"Nama_Kategori_Tujuan_Pembiayaan"`
}

func GetMasterKategoriTujuanPembiayaan() ([]MasterKategoriTujuanPembiayaan,error) {
	data := make([]MasterKategoriTujuanPembiayaan, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "EXEC GET_MASTER_KATEGORI_TUJUAN_PEMBIAYAAN "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterKategoriTujuanPembiayaan ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterFrekuensi struct {
	ID string `json:"id" db:"ID_Frekuensi"`
	Nama_Frekuensi string `json:"namafrekuensi" db:"Nama_Frekuensi"`
}

func GetMasterFrekuensi() ([]MasterFrekuensi,error) {
	data := make([]MasterFrekuensi, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "EXEC [GET_MASTER_FREKUENSI]"
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterFrekuensi ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterWilayahMobile struct {
	ID_Provinsi string `json:"ID_Provinsi" db:"ID_Provinsi"`
	Nama_Provinsi string `json:"Nama_Provinsi" db:"Nama_Provinsi"`
	ID_Kabkot string `json:"ID_Kabkot" db:"ID_Kabkot"`
	Nama_KabKot string `json:"Nama_KabKot" db:"Nama_KabKot"`
	ID_Kecamatan string `json:"ID_Kecamatan" db:"ID_Kecamatan"`
	Nama_Kecamatan string `json:"Nama_Kecamatan" db:"Nama_Kecamatan"`
	ID_KelDes string `json:"ID_KelDes" db:"ID_KelDes"`
	Nama_KelurahanDesa string `json:"Nama_KelurahanDesa" db:"Nama_KelurahanDesa"`
}

func GetMasterWilayahMobile(OurBranchID string) ([]MasterWilayahMobile,error) {
	data := make([]MasterWilayahMobile, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "EXEC [GET_MASTER_WILAYAH_MOBILE] @OurBranchID='"+OurBranchID+"' "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models.GetMasterWilayahMobile ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

func GetAllMasterWilayahMobile() ([]MasterWilayahMobile,error) {
	data := make([]MasterWilayahMobile, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "SELECT  e.ID_Provinsi,e.Nama_Provinsi,b.ID_Kabkot, b.Nama_KabKot, c.ID_Kecamatan, c.Nama_Kecamatan, d.ID_KelDes, d.Nama_KelurahanDesa "+
	"FROM INISIASI_MASTER_KABKOT b "+
	"INNER JOIN INISIASI_MASTER_KECAMATAN c ON b.ID_Kabkot = c.ID_KabKot "+
	"INNER JOIN INISIASI_MASTER_KELDES d ON c.ID_Kecamatan = d.ID_Kecamatan "+
	"INNER JOIN INISIASI_MASTER_PROVINSI e ON b.ID_Provinsi = e.ID_Provinsi "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models.GetAllMasterWilayahMobile ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type MasterGroupProduct struct {
	ID_Group_Product string `json:"ID_Group_Product" db:"ID_Group_Product"`
	Group_Product_Name string `json:"Group_Product_Name" db:"Group_Product_Name"`
	Description string `json:"Description" db:"Description"`
}

func GetMasterGroupProduct()([]MasterGroupProduct,error){
	data := make([]MasterGroupProduct, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}			
	defer db.Close()

	var qmain string

	qmain = "EXEC [GET_MASTER_GROUP_PRODUCT] "
	// fmt.Println(qmain)	
	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models.GetMasterGroupProduct ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil	
}

type SP_GET_LIST_OF_CLIENT_BRNET struct {
	ClientID string `json:"ClientID" db:"ClientID"`
	Name string `json:"Name" db:"Name"`
}

func GetListOfClientBRNET(params map[string]string)([]SP_GET_LIST_OF_CLIENT_BRNET,error) {
	data := make([]SP_GET_LIST_OF_CLIENT_BRNET, 0)
	db,err := global.ConnBRNET_GET()
	if err != nil {
		return data,err
	}			
	var qmain string

	qmain = "EXEC GET_LIST_OF_CLIENT_BRNET @OurBranchID='"+params["OurBranchID"]+"' "

	if params["Search"]!="undefined" {
		qmain += " ,@Search = '"+params["Search"]+"' "
	}

	qmain += " ,@PAGE = '"+params["PAGE"]+"',@LIMIT = '"+params["LIMIT"]+"' "

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetListOfClientBRNET ---> "+qmain+" ---> " + err.Error())
		return data,err
	}			

	return data,nil
}

type PARAM_GET_SIKLUS_NASABAH_BRNET struct {
	ClientID []string
}


type SP_GET_SIKLUS_NASABAH_BRNET struct {
	ClientID string `json:"ClientID" db:"ClientID"`
	ClientName string `json:"ClientName" db:"ClientName"`
	IdentityNumber string `json:"IdentityNumber" db:"IdentityNumber"`
	GroupID string `json:"GroupID" db:"GroupID"`
	SubGroup string `json:"SubGroup" db:"SubGroup"`
	GroupName string `json:"GroupName" db:"GroupName"`
	LoanSeries string `json:"LoanSeries" db:"LoanSeries"`
}

func GetSiklusNasabahBRNET(params PARAM_GET_SIKLUS_NASABAH_BRNET)([]SP_GET_SIKLUS_NASABAH_BRNET,error){
	data := make([]SP_GET_SIKLUS_NASABAH_BRNET, 0)
	db,err := global.ConnBRNET_GET()
	if err != nil {
		return data,err
	}		
	var qmain string = ""

	qmain = "DECLARE @CLIENT_DATA AS CLIENT_DATA; "

	for _, v := range params.ClientID {
		qmain += " INSERT INTO @CLIENT_DATA (ClientID) VALUES ('"+v+"'); "
	}

	qmain += "EXEC [GET_SIKLUS_NASABAH] @ClientID = @CLIENT_DATA "	
	
	fmt.Println(qmain)

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetSiklusNasabahBRNET ---> "+qmain+" ---> " + err.Error())	
		return data,err
	}			

	return data,nil
}

type MasterAvailableSubGroup struct {
	OurBranchID string `json:"OurBranchID" db:"OurBranchID"`
	GroupID string `json:"GroupID" db:"GroupID"`
	SubGroupID string `json:"SubGroupID" db:"SubGroupID"`
	Total string `json:"Total" db:"TOTAL"`
}

func GetMasterAvailableSubGroupBRNET(OurBranchID string)([]MasterAvailableSubGroup,error){
	data := make([]MasterAvailableSubGroup, 0)
	db,err := global.ConnBRNET_GET()
	if err != nil {
		return data,err
	}		
	var qmain string	

	qmain = "EXEC [DB_INISIASI].dbo.[GET_AVAILABLE_SUB_GROUP] @OurBranchID= '"+OurBranchID+"' "		

	err = db.RawSQL(qmain).Do(&data)

	if err != nil {
		global.Logging("ERROR","models GetMasterAvailableSubGroupBRNET ---> "+qmain+" ---> " + err.Error())			
		return data,err
	}			

	return data,nil
}