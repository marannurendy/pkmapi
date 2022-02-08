package models

import (
	// "beego/conf"
	"pkmapi/global"	
	"fmt"
)

type PkmbSurvey struct{
	ClientID string `json:"clientid"`
	AccountID string `json:"accountid"`
	GroupID string `json:"groupid"`
	ID int `json:"id"`
	Pertanyaan string `json:"pertanyaan"`
	Jawaban []string `json:"jawaban"`
	Nilai []int `json:"nilai"`
	// Jawaban1 string `json:"jawaban1"`
	// Jawaban2 string `json:"jawaban2"`
	// Jawaban3 string `json:"jawaban3"`
	// Jawaban4 string `json:"jawaban4"`
}

type PkmbJawaban struct {
	Jawaban string
	Nilai int
}

func Select_PkmbSurvey(AccountID string,ClientID string,GroupID string)[]PkmbSurvey {
	data := make([]PkmbSurvey, 0)
	var query string
	
	db,err := global.ConnPKUold()
	if err != nil {
		return data
	}
	defer db.Close()
	query = "SELECT '"+ClientID+"' as ClientID,'"+AccountID+"' as AccountID,'"+GroupID+"' as GroupID,ID,Pertanyaan FROM PKMB_SURVEY WITH (NOLOCK) "
	
	// fmt.Println("GroupID nya ",GroupID)
	rows, err := db.Query(query)	
	var i int=0	
	if err != nil {fmt.Println(err.Error())}
	defer rows.Close()
	for rows.Next(){					
		var each = PkmbSurvey{}			
		data_jawaban := make([]string,0)
		data_nilai := make([]int,0)

		err = rows.Scan(&each.ClientID,&each.AccountID,&each.GroupID,&each.ID,&each.Pertanyaan)
		if err != nil {fmt.Println(err.Error())}		

		rowsj, errj := db.Query("EXEC [GET_JAWABAN_PKMB]  @IdPertanyaan=?,@AccountID=?,@ClientID=?",each.ID,AccountID,ClientID)

		// fmt.Println(each.ID,AccountID,ClientID)

		var x int=0
		if errj != nil {fmt.Println(errj.Error())}
		defer rowsj.Close()
		for rowsj.Next(){				
			var each_jawaban = PkmbJawaban{}
			errj = rowsj.Scan(&each_jawaban.Jawaban,&each_jawaban.Nilai)

			// fmt.Println(each_jawaban.Jawaban)
			// data_jawaban[x] = each_jawaban.Jawaban
			// data_nilai[x] = each_jawaban.Nilai

			data_jawaban = append(data_jawaban, each_jawaban.Jawaban)
			data_nilai = append(data_nilai, each_jawaban.Nilai)

			x++
		}

		// fmt.Println(data_jawaban)
		
		data = append(data, PkmbSurvey{
			ClientID : each.ClientID,
			AccountID : each.AccountID,
			GroupID : each.GroupID,
			ID 	: each.ID,
			Pertanyaan : each.Pertanyaan,
			Jawaban : data_jawaban,
			Nilai : data_nilai,
		})

		// data[i] = PkmbSurvey{
		// 	ClientID : each.ClientID,
		// 	AccountID : each.AccountID,
		// 	GroupID : each.GroupID,
		// 	ID 	: each.ID,
		// 	Pertanyaan : each.Pertanyaan,
		// 	Jawaban : data_jawaban,
		// 	Nilai : data_nilai,
		// }		

		i++
	}


	return data
}


type PkmbMateri struct{
	ID int `json:"id"`
	Judul string `json:"judul"`
	Materi string `json:"content"`
}


func Select_PkmbMateri()PkmbMateri {
	var data PkmbMateri 
	db,err := global.ConnPKUold()
	if err != nil {
		return data
	}
	defer db.Close()	
	rows, err := db.Query("SELECT ID,Judul,Materi FROM PKMB_MATERI_BY_MONTH")
	
	if err != nil {fmt.Println(err.Error())}
	defer rows.Close()
	for rows.Next(){					
		var each = PkmbMateri{}	

		err = rows.Scan(&each.ID,&each.Judul,&each.Materi)
		if err != nil {fmt.Println(err.Error())}

		data = each //append(data, each)			
	}


	return data
}


func Select_PkmbMingguKe()int {
	var data int
	db,err := global.ConnPKUold()
	if err != nil {
		return data
	}
	defer db.Close()	
	rows, err := db.Query("select MINGGU_KE from CEK_MINGGU_KE_PER_BULAN()")
	
	if err != nil {fmt.Println(err.Error())}
	defer rows.Close()
	for rows.Next(){						
		err = rows.Scan(&data)
		if err != nil {fmt.Println(err.Error())}
		data = data			
	}

	return data
}

type AoGroup struct {
	GroupID string
}

func Select_Ao_Group(username string)[]AoGroup {
	var data []AoGroup 
	db,err := global.ConnPKUold()
	if err != nil {
		return data
	}
	defer db.Close()	
	rows, err := db.Query("SELECT GroupID FROM AO_GROUP WITH (NOLOCK) WHERE Username=?",username)
	
	if err != nil {fmt.Println(err.Error())}
	defer rows.Close()
	for rows.Next(){					
		var each = AoGroup{}	

		err = rows.Scan(&each.GroupID)
		if err != nil {fmt.Println(err.Error())}

		data = append(data, each)			
	}


	return data	
}

type Client struct {
	ClientID string
	AccountID string
}

func Client_By_Group(groupid string)[]Client {
	var data []Client 
	db,err := global.ConnPKUold()
	if err != nil {
		return data
	}
	defer db.Close()

	rows, err := db.Query("SELECT ClientID,AccountID FROM PKM_LPM WITH (NOLOCK) WHERE GroupID=?",groupid)
	
	if err != nil {fmt.Println(err.Error())}
	defer rows.Close()
	for rows.Next(){					
		var each = Client{}	

		err = rows.Scan(&each.ClientID,&each.AccountID)
		if err != nil {fmt.Println(err.Error())}

		data = append(data, each)			
	}


	return data	
}



type PostPkmSurvey struct {
	CABANG_ID string `json:"CABANG_ID"`
	KELOMPOK_ID string `json:"KELOMPOK_ID"`
	NASABAH_ID string `json:"NASABAH_ID"`
	NAMA_NASABAH string `json:"NAMA_NASABAH"`
	NILAI_SURVEY_MENABUNG string `json:"NILAI_SURVEY_MENABUNG"`
	NILAI_SURVEY_PENGELOLAAN_KEUANGAN string `json:"NILAI_SURVEY_PENGELOLAAN_KEUANGAN"`
	NILAI_SURVEY_OMSET string `json:"NILAI_SURVEY_OMSET"`
	NILAI_STRATEGI_PENJUALAN string `json:"NILAI_STRATEGI_PENJUALAN"`
	NILAI_SURVEY_ASSET string `json:"NILAI_SURVEY_ASSET"`
	NILAI_SURVEY_IJIN string `json:"NILAI_SURVEY_IJIN"`
	NILAI_SURVEY_DIVERSIFIKASI string `json:"NILAI_SURVEY_DIVERSIFIKASI"`
	NILAI_SURVEY_TENAGA_KERJA string `json:"NILAI_SURVEY_TENAGA_KERJA"`
}

func Post_Pkm_Survey(params []PostPkmSurvey) (error) {
	db,err := global.ConnPKU()
	if err != nil {
		return err
	}

	err = db.Begin() //Begin Transaction
	if err != nil {return err}		

	for _, data := range params {

		query := "EXEC POST_PKM_SURVEY_MOBILE "+
		"@CABANG_ID = '"+data.CABANG_ID+"' "+
		",@KELOMPOK_ID = '"+data.KELOMPOK_ID+"' "+
		",@NASABAH_ID = '"+data.NASABAH_ID+"' "+
		",@NAMA_NASABAH = '"+data.NAMA_NASABAH+"' "+
		",@NILAI_SURVEY_MENABUNG = '"+data.NILAI_SURVEY_MENABUNG+"' "+
		",@NILAI_SURVEY_PENGELOLAAN_KEUANGAN = '"+data.NILAI_SURVEY_PENGELOLAAN_KEUANGAN+"' "+
		",@NILAI_SURVEY_OMSET = '"+data.NILAI_SURVEY_OMSET+"' "+
		",@NILAI_STRATEGI_PENJUALAN = '"+data.NILAI_STRATEGI_PENJUALAN+"' "+
		",@NILAI_SURVEY_ASSET = '"+data.NILAI_SURVEY_ASSET+"' "+
		",@NILAI_SURVEY_IJIN = '"+data.NILAI_SURVEY_IJIN+"' "+
		",@NILAI_SURVEY_DIVERSIFIKASI = '"+data.NILAI_SURVEY_DIVERSIFIKASI+"' "+
		",@NILAI_SURVEY_TENAGA_KERJA = '"+data.NILAI_SURVEY_TENAGA_KERJA+"' "

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
