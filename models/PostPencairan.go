package models

import (
	"pkmapi/global"

	"github.com/astaxie/beego"
	// "strconv"
	// "github.com/astaxie/beego/validation"
	// "time"
	// "fmt"
)

type SP_ADD_PENCAIRAN struct {
	ID_Prospek      string `json:"ID_Prospek"`
	Is_Batal        string `json:"Is_Batal"`
	TTD_Nasabah     string `json:"TTD_Nasabah"`
	TTD_KC          string `json:"TTD_KC"`
	TTD_Nasabah_2   string `json:"TTD_Nasabah_2"`
	TTD_KSK         string `json:"TTD_KSK"`
	TTD_KK          string `json:"TTD_KK"`
	Foto_Pencairan  string `json:"Foto_Pencairan"`
	Is_Ludin        string `json:"Is_Ludin"`
	Is_PMU          string `json:"Is_PMU"`
	Is_Dicairkan    string `json:"Is_Dicairkan"`
	Jml_UP          string `json:"Jml_UP"`
	Jml_Sisa_UP     string `json:"Jml_Sisa_UP"`
	Jml_Cair_PMU    string `json:"Jml_Cair_PMU"`
	Jml_RealCair    string `json:"Jml_RealCair"`
	LRP_TTD_Nasabah string `json:"LRP_TTD_Nasabah"`
	LRP_TTD_AO      string `json:"LRP_TTD_AO"`
	FP4             string `json:"FP4"`
	Foto_Kegiatan   string `json:"Foto_Kegiatan"`
}

func POST_PENCAIRAN(data_arr []SP_ADD_PENCAIRAN) error {
	db, err := global.ConnPost()
	if err != nil {
		return err
	}
	defer db.Close()

	var query string
	var fileimages = make(map[string]string)
	// var pathfile string

	err = db.Begin() //Begin Transaction
	if err != nil {
		return err
	}

	for _, data := range data_arr {
		// pathfile = "images/"+time.Now().Format("2006-02-01")+"/"+data.ID_Prospek+"/pencairan/"
		fileimages["TTD_Nasabah"] = data.TTD_Nasabah
		fileimages["TTD_KC"] = data.TTD_KC
		fileimages["TTD_Nasabah_2"] = data.TTD_Nasabah_2
		fileimages["TTD_KSK"] = data.TTD_KSK
		fileimages["TTD_KK"] = data.TTD_KK
		fileimages["Foto_Pencairan"] = data.Foto_Pencairan
		fileimages["Foto_Kegiatan"] = data.Foto_Kegiatan
		fileimages["LRP_TTD_Nasabah"] = data.LRP_TTD_Nasabah
		fileimages["LRP_TTD_AO"] = data.LRP_TTD_AO

		fileimages, err = global.MapB64SaveFile(fileimages, beego.AppConfig.String("pathimages"))
		if err != nil {
			errfile := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {
				return err
			}
			return errfile
		}

		query = "EXEC ADD_PENCAIRAN " +
			"@ID_Prospek = '" + data.ID_Prospek + "' " +
			",@Is_Batal = '" + data.Is_Batal + "' " +
			",@TTD_Nasabah = '" + fileimages["TTD_Nasabah"] + "' " +
			",@TTD_KC = '" + fileimages["TTD_KC"] + "' " +
			",@TTD_Nasabah_2 = '" + fileimages["TTD_Nasabah_2"] + "' " +
			",@TTD_KSK = '" + fileimages["TTD_KSK"] + "' " +
			",@TTD_KK = '" + fileimages["TTD_KSK"] + "' " +
			",@Foto_Pencairan = '" + fileimages["Foto_Pencairan"] + "' " +
			",@Is_Ludin = '" + data.Is_Ludin + "' " +
			",@Is_PMU = '" + data.Is_PMU + "' " +
			",@Is_Dicairkan = '" + data.Is_Dicairkan + "' " +
			",@Jml_UP = '" + data.Jml_UP + "' " +
			",@Jml_Sisa_UP = '" + data.Jml_Sisa_UP + "' " +
			",@Jml_Cair_PMU = '" + data.Jml_Cair_PMU + "' " +
			",@Jml_RealCair = '" + data.Jml_RealCair + "' " +
			",@LRP_TTD_Nasabah = '" + fileimages["LRP_TTD_Nasabah"] + "' " +
			",@LRP_TTD_AO = '" + fileimages["LRP_TTD_AO"] + "' " +
			",@FP4 = '" + data.FP4 + "' " +
			",@Foto_Kegiatan = '" + fileimages["Foto_Kegiatan"] + "' "

		_, err = db.RawSQL(query).DoWithIterator()
		if err != nil {
			global.Logging("ERROR", "models.POST_PENCAIRAN "+query+" ---> "+err.Error())
			errsql := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {
				return err
			}
			return errsql
		}

	}

	err = db.Commit() //Commit Transaction
	if err != nil {
		return err
	}

	// err = db.Close()
	// if err != nil {return err}

	return nil

}
