package models

import (
	"fmt"
	"pkmapi/global"
)

type DetailPKM struct {
	OurBranchID   string  `json:"OurBranchID"`
	GroupName     string  `json:"GroupName"`
	MeetingDay    int     `json:"MeetingDay"`
	AnggotaAktif  string  `json:"AnggotaAktif"`
	JumlahTagihan float32 `json:"JumlahTagihan"`
	MeetingPlace  string  `json:"MeetingPlace"`
	MeetingTime   string  `json:"MeetingTime"`
}

// func Select_list()[]DetailPKM {
func Select_list(MeetingDayID string, AO_ID string, tipe string, GroupID string) []DetailPKM {
	var data []DetailPKM
	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("EXEC GET_Detail_PKM_LPM_TYW @MeetingDayID=" + MeetingDayID + ", @AO_ID=" + AO_ID + ", @type=" + tipe + ", @GroupID=" + GroupID)

	// rows, err := global.Db.Query("SELECT TOP 5 OurBranchID, GroupName, MeetingDay, COUNT(ClientID) as AnggotaAktif, SUM(InstallmentAmount) JumlahTagihan, MeetingPlace, MeetingTime FROM PKM_LPM WHERE OurBranchID = '90001'")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var each = DetailPKM{}
		err = rows.Scan(&each.OurBranchID, &each.GroupName, &each.MeetingDay, &each.AnggotaAktif, &each.JumlahTagihan, &each.MeetingPlace, &each.MeetingTime)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}

	// db.Close()

	return data
}

type DetailPKM2 struct {
	OurBranchID  string `json:"OurBranchID"`
	GroupName    string `json:"GroupName"`
	MeetingDay   string `json:"MeetingDay"`
	MeetingPlace string `json:"MeetingPlace"`
	MeetingTime  string `json:"MeetingTime"`
}

func Select_data() []DetailPKM2 {
	var data []DetailPKM2
	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("EXEC SP_Rendy_Test")
	i := 0

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		// var each = DetailPKM2{}
		// _ = rows.Scan(&each.OurBranchID,&each.GroupName,&each.MeetingDay,&each.MeetingPlace,&each.MeetingTime)
		// data[i] = DetailPKM2{\
		var OurBranchID, GroupName, MeetingDay, MeetingPlace, MeetingTime string
		_ = rows.Scan(&OurBranchID, &GroupName, &MeetingDay, &MeetingPlace, &MeetingTime)
		data[i] = DetailPKM2{
			OurBranchID:  OurBranchID,
			GroupName:    GroupName,
			MeetingDay:   MeetingDay,
			MeetingPlace: MeetingPlace,
			MeetingTime:  MeetingTime,
		}
		i++
	}

	// db.Close()

	return data
}

type DetailTest struct {
	OurBranchID     string `json:"OurBranchID"`
	GroupName       string `json:"GroupName"`
	GroupID         string `json:"GroupID"`
	MeetingDay      string `json:"MeetingDay"`
	AnggotaAktif    string `json:"AnggotaAktif"`
	JumlahTagihan   string `json:"JumlahTagihan"`
	MeetingPlace    string `json:"MeetingPlace"`
	MeetingTime     string `json:"MeetingTime"`
	CreditOfficerID string `json:"CreditOfficerID"`
}

func SelectTest() []DetailTest {
	var data []DetailTest

	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	// rows, err := global.Db.Query("SELECT DISTINCT TOP 10 OurBranchID, GroupID, Initial, GroupName, MeetingDay FROM PKM_LPM WHERE GroupName NOT LIKE '%TRF%'")
	// rows, err := global.Db.Query("Select OurBranchID, GroupID, GroupName, MeetingDay, CreditOfficerID from PKM_LPM where CreditOfficerID is not null AND CreditOfficerID = '54545422'")
	rows, err := db.Query("SELECT OurBranchID, GroupName, GroupID, MeetingDay, COUNT(ClientID) as AnggotaAktif, SUM(InstallmentAmount) JumlahTagihan, CASE WHEN MeetingPlace is null then '-' else MeetingDay end as MeetingPlace, MeetingTime, CreditOfficerID FROM PKM_LPM WHERE CreditOfficerID = '54545422' GROUP BY OurBranchID,MeetingDay,MeetingPlace,MeetingTime,GroupName,GroupID,CreditOfficerID")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var each = DetailTest{}
		err = rows.Scan(&each.OurBranchID, &each.GroupName, &each.GroupID, &each.MeetingDay, &each.AnggotaAktif, &each.JumlahTagihan, &each.MeetingPlace, &each.MeetingTime, &each.CreditOfficerID)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}
	// db.Close()
	return data
}

type DetailShit struct {
	OurBranchID       string `json:"OurBranchID"`
	GroupName         string `json:"GroupName"`
	GroupID           string `json:"GroupID"`
	MeetingDay        string `json:"MeetingDay"`
	MeetingPlace      string `json:"MeetingPlace"`
	MeetingTime       string `json:"MeetingTime"`
	ClientID          string `json:"ClientID"`
	ClientName        string `json:"ClientName"`
	AccountID         string `json:"AccountID"`
	InstallmentAmount string `json:"InstallmentAmount"`
	VolSavingsBal     string `json:"VolSavingsBal"`
	CompSavingsBal    string `json:"CompSavingsBal"`
}

func SelectShit() []DetailShit {
	var data []DetailShit
	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT OurBranchID, GroupName, GroupID, MeetingDay, CASE WHEN MeetingPlace is null then '-' else MeetingDay end as MeetingPlace, MeetingTime, ClientID, ClientName, AccountID, InstallmentAmount, VolSavingsBal, CASE WHEN CompSavingsBal is null then '0' else CompSavingsBal end as CompSavingsBal FROM PKM_LPM WHERE CreditOfficerID = '54545422'")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var each = DetailShit{}
		err = rows.Scan(&each.OurBranchID, &each.GroupName, &each.GroupID, &each.MeetingDay, &each.MeetingPlace, &each.MeetingTime, &each.ClientID, &each.ClientName, &each.AccountID, &each.InstallmentAmount, &each.VolSavingsBal, &each.CompSavingsBal)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}
	// db.Close()
	return data
}

func PostTotalTransaction(params map[string]interface{}) {
	// fmt.Println(
	//
	// _, err := global.Db.Exec("EXEC ADD_PKM_Transaction @GroupID = '" + fmt.Sprintf("%v", params["GroupID"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["ClientID"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["AccountID"]) + "', @AttendStatus = '" + fmt.Sprintf("%v", params["Attendance"]) + "' , @InstallmentAmount = '" + fmt.Sprintf("%v", params["Angsuran"]) + "', @CreditSaving = '" + fmt.Sprintf("%v", params["Tarikan"]) + "', @DebitSaving = '" + fmt.Sprintf("%v", params["Setoran"]) + "', @CreditUP = '0', @TTD1 = '" + fmt.Sprintf("%v", params["TtdAccountOfficer"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["TtdKetuaKelompok"]) + "', @CreatedDate = '" + fmt.Sprintf("%v", params["CreatedDate"]) + "'")
	// global.Logging("INFO EXEC ADD_PKM_Transaction @GroupID = '" + fmt.Sprintf("%v", params["GroupID"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["ClientID"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["AccountID"]) + "', @AttendStatus = '" + fmt.Sprintf("%v", params["Attendance"]) + "' , @InstallmentAmount = '" + fmt.Sprintf("%v", params["Angsuran"]) + "', @CreditSaving = '" + fmt.Sprintf("%v", params["Tarikan"]) + "', @DebitSaving = '" + fmt.Sprintf("%v", params["Setoran"]) + "', @CreditUP = '0'")
	db, err := global.ConnPKMPost()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("EXEC ADD_PKM_Transaction @GroupID = '" + fmt.Sprintf("%v", params["GroupID"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["ClientID"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["AccountID"]) + "', @AttendStatus = '" + fmt.Sprintf("%v", params["Attendance"]) + "' , @InstallmentAmount = '" + fmt.Sprintf("%v", params["Angsuran"]) + "', @CreditSaving = '" + fmt.Sprintf("%v", params["Tarikan"]) + "', @DebitSaving = '" + fmt.Sprintf("%v", params["Setoran"]) + "', @CreditUP = '0', @TTD1 = '1', @TTD2 = '2', @CreatedDate = '" + fmt.Sprintf("%v", params["CreatedDate"]) + "'")
	global.Logging("INFO_PKM", "EXEC ADD_PKM_Transaction @GroupID = '"+fmt.Sprintf("%v", params["GroupID"])+"', @ClientID = '"+fmt.Sprintf("%v", params["ClientID"])+"', @AccountID = '"+fmt.Sprintf("%v", params["AccountID"])+"', @AttendStatus = '"+fmt.Sprintf("%v", params["Attendance"])+"' , @InstallmentAmount = '"+fmt.Sprintf("%v", params["Angsuran"])+"', @CreditSaving = '"+fmt.Sprintf("%v", params["Tarikan"])+"', @DebitSaving = '"+fmt.Sprintf("%v", params["Setoran"])+"', @CreditUP = '0', @TTD1 = '1', @TTD2 = '2', @CreatedDate = '"+fmt.Sprintf("%v", params["CreatedDate"])+"'")

	if err != nil {
		fmt.Println(err)
		global.Logging("ERROR_PKM", "EXEC ADD_PKM_Transaction @GroupID = '"+fmt.Sprintf("%v", params["GroupID"])+"', @ClientID = '"+fmt.Sprintf("%v", params["ClientID"])+"', @AccountID = '"+fmt.Sprintf("%v", params["AccountID"])+"', @AttendStatus = '"+fmt.Sprintf("%v", params["Attendance"])+"' , @InstallmentAmount = '"+fmt.Sprintf("%v", params["Angsuran"])+"', @CreditSaving = '"+fmt.Sprintf("%v", params["Tarikan"])+"', @DebitSaving = '"+fmt.Sprintf("%v", params["Setoran"])+"', @CreditUP = '0', @TTD1 = '1', @TTD2 = '2', @CreatedDate = '"+fmt.Sprintf("%v", params["CreatedDate"])+"' ---> "+err.Error())
	}
	// else {
	// 	fmt.Println("TRUE")
	// 	// fmt.Println(path)
	// 	global.SuccessLog("SUCCESS POST PKM EXEC ADD_PKM_Transaction @GroupID = '" + fmt.Sprintf("%v", params["GroupID"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["ClientID"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["AccountID"]) + "', @AttendStatus = '" + fmt.Sprintf("%v", params["Attendance"]) + "' , @InstallmentAmount = '" + fmt.Sprintf("%v", params["Angsuran"]) + "', @CreditSaving = '" + fmt.Sprintf("%v", params["Tarikan"]) + "', @DebitSaving = '" + fmt.Sprintf("%v", params["Setoran"]) + "', @CreditUP = '0', @TTD1 = '" + fmt.Sprintf("%v", params["TtdAccountOfficer"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["TtdKetuaKelompok"]) + "', @CreatedDate = '" + fmt.Sprintf("%v", params["CreatedDate"]) + "'")
	// }
	// db.Close()
}

func PostTotalTransactionUP(params map[string]interface{}) {
	// fmt.Println(
	//
	db, err := global.ConnPKMPost()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("EXEC ADD_PKM_Transaction_UP @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @CreditUP = '" + fmt.Sprintf("%v", params["jumlahup"]) + "'")
	fmt.Println("EXEC ADD_PKM_Transaction_UP @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @CreditUP = '" + fmt.Sprintf("%v", params["jumlahup"]) + "'")

	global.Logging("INFO_PKM", "EXEC ADD_PKM_Transaction_UP @ClientID = '"+fmt.Sprintf("%v", params["clientid"])+"', @CreditUP = '"+fmt.Sprintf("%v", params["jumlahup"])+"'")

	if err != nil {
		fmt.Println(err)
		global.Logging("ERROR_PKM", "UP EXEC ADD_PKM_Transaction_UP @ClientID = '"+fmt.Sprintf("%v", params["clientid"])+"', @CreditUP = '"+fmt.Sprintf("%v", params["jumlahup"])+"' ---> "+err.Error())
	}
	// else {
	// 	fmt.Println("TRUE")
	// 	global.SuccessLog("SUCCESS POST UP EXEC ADD_PKM_Transaction_UP @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @CreditUP = '" + fmt.Sprintf("%v", params["jumlahup"]) + "'")
	// }
	// db.Close()
}

func PostTotalTransactionPAR(params map[string]interface{}) {
	// fmt.Println(
	//
	// _, err := global.Db.Exec("EXEC ADD_PKM_Transaction_PAR @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitPAR = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["clientSign"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["AOSign"]) + "', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "'")
	// global.Logging("INFO EXEC ADD_PKM_Transaction_PAR @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitPAR = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["clientSign"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["AOSign"]) + "', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "'")
	db, err := global.ConnPKMPost()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("EXEC ADD_PKM_Transaction_PAR @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitPAR = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '1', @TTD2 = '1', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "'")
	global.Logging("INFO_PKM", "EXEC ADD_PKM_Transaction_PAR @GroupID = '"+fmt.Sprintf("%v", params["groupid"])+"', @ClientID = '"+fmt.Sprintf("%v", params["clientid"])+"', @AccountID = '"+fmt.Sprintf("%v", params["accountid"])+"', @DebitPAR = '"+fmt.Sprintf("%v", params["jumlahpar"])+"', @TTD1 = '1', @TTD2 = '1', @CreatedBy = '"+fmt.Sprintf("%v", params["createdby"])+"'")

	if err != nil {
		fmt.Println(err)
		global.Logging("ERROR_PKM", "PAR EXEC ADD_PKM_Transaction_PAR @GroupID = '"+fmt.Sprintf("%v", params["groupid"])+"', @ClientID = '"+fmt.Sprintf("%v", params["clientid"])+"', @AccountID = '"+fmt.Sprintf("%v", params["accountid"])+"', @DebitPAR = '"+fmt.Sprintf("%v", params["jumlahpar"])+"', @TTD1 = '1', @TTD2 = '1', @CreatedBy = '"+fmt.Sprintf("%v", params["createdby"])+"' ---> "+err.Error())
	}
	// else {
	// 	fmt.Println("TRUE")
	// 	global.SuccessLog("SUCCESS POST PAR EXEC ADD_PKM_Transaction_PAR @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitPAR = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["clientSign"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["AOSign"]) + "', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "'")
	// }
	// db.Close()
}

type ListCollection struct {
	OurBranchID       string `json:"OurBranchID"`
	GroupID           string `json:"GroupID"`
	GroupName         string `json:"GroupName"`
	MeetingDay        string `json:"MeetingDay"`
	AccountID         string `json:"AccountID"`
	ProductID         string `json:"ProductID"`
	ClientID          string `json:"ClientID"`
	ClientName        string `json:"ClientName"`
	InstallmentAmount string `json:"InstallmentAmount"`
	Rill              string `json:"Rill"`
	Ke                string `json:"Ke"`
	VolSavingsBal     string `json:"VolSavingsBal"`
	StatusPAR         string `json:"StatusPAR"`
}

func GetListCollection(OurBranchID string, AO_ID string) []ListCollection {
	var data []ListCollection
	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	fmt.Println("EXEC GET_PKM_ANGSURAN @AO_ID = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'Angsuran'")
	global.Logging("INFO_PKM", "EXEC GET_PKM_ANGSURAN @AO_ID = '"+AO_ID+"', @OurBranchID = '"+OurBranchID+"', @type = 'Angsuran'")

	rows, err := db.Query("EXEC GET_PKM_ANGSURAN @AO_ID = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'Angsuran'")

	if err != nil {
		// fmt.Println(err.Error())
		fmt.Println("ERROR")
		global.Logging("ERROR_PKM", "EXEC GET_PKM_ANGSURAN @AO_ID = '"+AO_ID+"', @OurBranchID = '"+OurBranchID+"', @type = 'Angsuran' ---> "+err.Error())
	}

	fmt.Println("dibawah")
	// else {
	// 	global.SuccessLog("SUCCESS GET EXEC GET_PKM_ANGSURAN @AO_ID = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'Angsuran' | MATCH")
	// }
	defer rows.Close()
	for rows.Next() {
		var each = ListCollection{}
		err = rows.Scan(
			&each.OurBranchID,
			&each.GroupID,
			&each.GroupName,
			&each.MeetingDay,
			&each.AccountID,
			&each.ProductID,
			&each.ClientID,
			&each.ClientName,
			&each.InstallmentAmount,
			&each.Rill,
			&each.Ke,
			&each.VolSavingsBal,
			&each.StatusPAR)
		if err != nil {
			fmt.Println("ERROR")
		}
		data = append(data, each)
	}

	if data != nil {
		fmt.Println("this")
	} else {
		fmt.Println("that")
	}
	// fmt.Println(data)
	// db.Close()
	return data
}

type Listup struct {
	OurBranchID    string `json:"OurBranchID"`
	ClientID       string `json:"ClientID"`
	ClientName     string `json:"ClientName"`
	GroupID        string `json:"GroupID"`
	GroupName      string `json:"GroupName"`
	MeetingDay     string `json:"MeetingDay"`
	CompSavingsBal string `json:"CompSavingsBal"`
}

func GetListup(OurBranchID string, AO_ID string) []Listup {
	var data []Listup
	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("EXEC GET_PKM_ANGSURAN @AO_ID = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'UP'")
	global.Logging("INFO_PKM", "EXEC GET_PKM_ANGSURAN @AO_ID = '"+AO_ID+"', @OurBranchID = '"+OurBranchID+"', @type = 'UP'")
	if err != nil {
		fmt.Println(err.Error())
		global.Logging("ERROR_PKM", "UP EXEC GET_PKM_ANGSURAN @AO_ID = '"+AO_ID+"', @OurBranchID = '"+OurBranchID+"', @type = 'UP' ---> "+err.Error())
	}
	// else {
	// 	global.SuccessLog("SUCCESS GET UP EXEC GET_PKM_ANGSURAN @AO_ID = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'UP' | MATCH")
	// }
	defer rows.Close()
	for rows.Next() {
		var each = Listup{}
		err = rows.Scan(&each.OurBranchID, &each.ClientID, &each.ClientName, &each.GroupID, &each.GroupName, &each.MeetingDay, &each.CompSavingsBal)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}
	// db.Close()
	return data
}

type GroupList struct {
	OurBranchID   string `json:"OurBranchID"`
	GroupID       string `json:"GroupID"`
	GroupName     string `json:"GroupName"`
	MeetingDay    string `json:"MeetingDay"`
	AnggotaAktif  string `json:"AnggotaAktif"`
	JumlahTagihan string `json:"JumlahTagihan"`
	MeetingPlace  string `json:"MeetingPlace"`
	MeetingTime   string `json:"MeetingTime"`
}

func GetListGroup(OurBranchID string, USERNAME string) []GroupList {
	var data []GroupList
	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("EXEC GET_Detail_PKM_LPM @OurBranchID='" + OurBranchID + "', @USERNAME='" + USERNAME + "', @type= 'Process'")
	fmt.Println("EXEC GET_Detail_PKM_LPM @OurBranchID='" + OurBranchID + "', @USERNAME='" + USERNAME + "', @type= 'Process'")

	// rows, err := global.Db.Query("SELECT TOP 5 OurBranchID, GroupName, MeetingDay, COUNT(ClientID) as AnggotaAktif, SUM(InstallmentAmount) JumlahTagihan, MeetingPlace, MeetingTime FROM PKM_LPM WHERE OurBranchID = '90001'")
	if err != nil {
		fmt.Println(err.Error())
		global.Logging("ERROR_PKM", "EXEC GET_Detail_PKM_LPM @OurBranchID='"+OurBranchID+"', @USERNAME='"+USERNAME+"', @type= 'Process' --->> "+err.Error())
	}
	// else {
	// 	global.SuccessLog("SUCCESS GET EXEC GET_Detail_PKM_LPM @OurBranchID='" + OurBranchID + "', @USERNAME='" + USERNAME + "', @type= 'Process' | MATCH")
	// }

	defer rows.Close()
	for rows.Next() {
		var each = GroupList{}
		err = rows.Scan(&each.OurBranchID, &each.GroupID, &each.GroupName, &each.MeetingDay, &each.AnggotaAktif, &each.JumlahTagihan, &each.MeetingPlace, &each.MeetingTime)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}
	// db.Close()
	return data
}

type DetailPKM_PAR struct {
	OurBranchID string `json:"OurBranchID"`
	ClientID    string `json:"ClientID"`
	ClientName  string `json:"ClientName"`
	AccountID   string `json:"AccountID"`
	ProductID   string `json:"ProductID"`
	GroupID     string `json:"GroupID"`
	GroupName   string `json:"GroupName"`
	ODAmount    string `json:"ODAmount"`
}

// func Select_list()[]DetailPKM {
func GetListCollectionPAR(AO_ID string, OurBranchID string) []DetailPKM_PAR {
	var data []DetailPKM_PAR
	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("EXEC GET_Detail_PKM_LPM_PAR @AO_ID = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'PAR'")

	// rows, err := global.Db.Query("SELECT TOP 5 OurBranchID, GroupName, MeetingDay, COUNT(ClientID) as AnggotaAktif, SUM(InstallmentAmount) JumlahTagihan, MeetingPlace, MeetingTime FROM PKM_LPM WHERE OurBranchID = '90001'")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var each = DetailPKM_PAR{}
		err = rows.Scan(&each.OurBranchID, &each.ClientID, &each.ClientName, &each.AccountID, &each.ProductID, &each.GroupID, &each.GroupName, &each.ODAmount)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}
	// db.Close()
	return data
}

type Login struct {
	OURBRANCHID  string `json:"OURBRANCHID"`
	NIP          string `json:"NIP"`
	BranchName   string `json:"BranchName"`
	USERNAME     string `json:"USERNAME"`
	Password     string `json:"Password"`
	NAMA         string `json:"NAMA"`
	LoginSession string `json:"LoginSession"`
}

type SetLogin struct {
	LoginStat  []Login
	Status     string
	Token      string
	Is_Syariah string
}

func GetLogin(Username string, Password string) SetLogin {
	var data []Login
	var loginStatus = ""
	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("EXEC GET_User_Mobile_Only @Username = '" + Username + "', @Password = '" + Password + "', @Status = '1'")

	fmt.Println("EXEC GET_User_Mobile_Only @Username = '" + Username + "', @Password = '" + Password + "', @Status = '1'")

	if err != nil {
		fmt.Println(err.Error())
		global.Logging("ERROR_PKM", "LOGIN EXEC GET_User_Mobile_Only @Username = '"+Username+"', @Password = '"+Password+"', @Status = '1' | MATCH ---> "+err.Error())
	}
	// else {
	// 	global.SuccessLog("SUCCESS LOGIN EXEC GET_User_Mobile_Only @Username = '" + Username + "', @Password = '" + Password + "', @Status = '1' | MATCH")
	// }
	defer rows.Close()
	for rows.Next() {
		var each = Login{}
		err = rows.Scan(&each.OURBRANCHID, &each.NIP, &each.BranchName, &each.USERNAME, &each.Password, &each.NAMA, &each.LoginSession)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}
	// fmt.Println(data[0].LoginSession)
	if len(data) != 0 {
		if data[0].LoginSession == "0" || data[0].LoginSession == "" {
			loginStatus = "Login Berhasil"
		} else {
			loginStatus = "Login Telah Dipakai"
		}
	} else {
		loginStatus = "Username Atau Password Salah"
	}

	var hasil = SetLogin{
		LoginStat: data,
		Status:    loginStatus,
		Token:     global.GenerateTokenJWT(Username + Password),
	}

	// db.Close()
	return hasil
}

type Logout struct {
	OURBRANCHID  string `json:"OURBRANCHID"`
	NIP          string `json:"NIP"`
	BranchName   string `json:"BranchName"`
	USERNAME     string `json:"USERNAME"`
	Password     string `json:"Password"`
	NAMA         string `json:"NAMA"`
	LoginSession string `json:"LoginSession"`
}

type SetLogout struct {
	LoginStat []Logout
	Status    string
}

func GetLogout(Username string, Password string) SetLogout {
	var data []Logout
	var logoutStatus = ""
	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("EXEC GET_User_Mobile_Only @Username = '" + Username + "', @Password = '" + Password + "', @Status = '0'")

	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var each = Logout{}
		err = rows.Scan(&each.OURBRANCHID, &each.NIP, &each.BranchName, &each.USERNAME, &each.Password, &each.NAMA, &each.LoginSession)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}
	// fmt.Println(data[0].LoginSession)
	if len(data) != 0 {
		if data[0].LoginSession == "0" || data[0].LoginSession == "" {
			logoutStatus = "Logout Berhasil"
		} else {
			logoutStatus = "Logout Telah Dipakai"
		}
	} else {
		logoutStatus = "Username Atau Password Salah"
	}

	var hasil = SetLogout{
		LoginStat: data,
		Status:    logoutStatus,
	}
	// db.Close()
	return hasil
}

type DetailGetPKMIndividual struct {
	OurBranchID       string `json:"OurBranchID"`
	GroupID           string `json:"GroupID"`
	GroupName         string `json:"GroupName"`
	MeetingDay        string `json:"MeetingDay"`
	AccountID         string `json:"AccountID"`
	ProductID         string `json:"ProductID"`
	ClientID          string `json:"ClientID"`
	ClientName        string `json:"ClientName"`
	InstallmentAmount string `json:"InstallmentAmount"`
	Rill              string `json:"Rill"`
	Ke                string `json:"Ke"`
	VolSavingsBal     string `json:"VolSavingsBal"`
	PARStatus         string `json:"PARStatus"`
}

// func Select_list()[]DetailPKM {
func GetListPKMIndividual(AO_ID string, OurBranchID string) []DetailGetPKMIndividual {
	var data []DetailGetPKMIndividual
	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("EXEC GET_Detail_PKM_LPM_PAR_INDIVIDUAL @USERNAME = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'Process'")

	if err != nil {
		fmt.Println(err.Error())
		global.Logging("ERROR_PKM", "Individu EXEC GET_Detail_PKM_LPM_PAR_INDIVIDUAL @USERNAME = '"+AO_ID+"', @OurBranchID = '"+OurBranchID+"', @type = 'Process' ---> "+err.Error())
	}
	// else {
	// 	global.SuccessLog("SUCCESS PKM Individu EXEC GET_Detail_PKM_LPM_PAR_INDIVIDUAL @USERNAME = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'Process' | MATCH")
	// }
	defer rows.Close()
	for rows.Next() {
		var each = DetailGetPKMIndividual{}
		err = rows.Scan(&each.OurBranchID, &each.GroupID, &each.GroupName, &each.MeetingDay, &each.AccountID, &each.ProductID, &each.ClientID, &each.ClientName, &each.InstallmentAmount, &each.Rill, &each.Ke, &each.VolSavingsBal, &each.PARStatus)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}
	// db.Close()
	return data
}

func PostPKMIndividual(params map[string]interface{}) {
	// fmt.Println(
	//
	// _, err := global.Db.Exec("EXEC ADD_PKM_Transaction_PAR_Individual @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitValue = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["clientSign"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["AOSign"]) + "', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "', @CreatedDate = '" + fmt.Sprintf("%v", params["CreatedDate"]) + "'")
	// fmt.Println("EXEC ADD_PKM_Transaction_PAR_Individual @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitValue = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["clientSign"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["AOSign"]) + "', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "', @CreatedDate = '" + fmt.Sprintf("%v", params["CreatedDate"]) + "'")
	db, err := global.ConnPKMPost()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("EXEC ADD_PKM_Transaction_PAR_Individual @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitValue = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '1', @TTD2 = '1', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "', @CreatedDate = '" + fmt.Sprintf("%v", params["CreatedDate"]) + "'")
	global.Logging("INFO_PKM", "EXEC ADD_PKM_Transaction_PAR_Individual @GroupID = '"+fmt.Sprintf("%v", params["groupid"])+"', @ClientID = '"+fmt.Sprintf("%v", params["clientid"])+"', @AccountID = '"+fmt.Sprintf("%v", params["accountid"])+"', @DebitValue = '"+fmt.Sprintf("%v", params["jumlahpar"])+"', @TTD1 = '1', @TTD2 = '1', @CreatedBy = '"+fmt.Sprintf("%v", params["createdby"])+"', @CreatedDate = '"+fmt.Sprintf("%v", params["CreatedDate"])+"'")

	if err != nil {
		fmt.Println(err)
		global.Logging("ERROR_PKM", "EXEC ADD_PKM_Transaction_PAR_Individual @GroupID = '"+fmt.Sprintf("%v", params["groupid"])+"', @ClientID = '"+fmt.Sprintf("%v", params["clientid"])+"', @AccountID = '"+fmt.Sprintf("%v", params["accountid"])+"', @DebitValue = '"+fmt.Sprintf("%v", params["jumlahpar"])+"', @TTD1 = '1', @TTD2 = '1', @CreatedBy = '"+fmt.Sprintf("%v", params["createdby"])+"', @CreatedDate = '"+fmt.Sprintf("%v", params["CreatedDate"])+"' ---> "+err.Error())
	}
	// else {
	// 	fmt.Println("TRUE")
	// 	global.SuccessLog("SUCCESS POST PKM Individu EXEC ADD_PKM_Transaction_PAR_Individual @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitValue = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["clientSign"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["AOSign"]) + "', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "', @CreatedDate = '" + fmt.Sprintf("%v", params["CreatedDate"]) + "'")
	// }
	// db.Close()
}

func PostTotalTransactionUPnew(params map[string]interface{}) {
	// fmt.Println(
	//
	db, err := global.ConnPKMPost()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("EXEC ADD_PKM_Transaction_UP @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @CreditUP = '" + fmt.Sprintf("%v", params["jumlahup"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["TtdAccountOfficer"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["TtdKetuaKelompok"]) + "'")
	fmt.Println("EXEC PKM ADD_PKM_Transaction_UP @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @CreditUP = '" + fmt.Sprintf("%v", params["jumlahup"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["TtdAccountOfficer"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["TtdKetuaKelompok"]) + "'")

	if err != nil {
		fmt.Println(err)
		global.Logging("ERROR_PKM", "POST UP EXEC ADD_PKM_Transaction_UP @ClientID = '"+fmt.Sprintf("%v", params["clientid"])+"', @CreditUP = '"+fmt.Sprintf("%v", params["jumlahup"])+"', @TTD1 = '"+fmt.Sprintf("%v", params["TtdAccountOfficer"])+"', @TTD2 = '"+fmt.Sprintf("%v", params["TtdKetuaKelompok"])+"' ---> "+err.Error())
	}
	// else {
	// 	fmt.Println("TRUE")
	// 	global.SuccessLog("SUCCESS POST UP EXEC ADD_PKM_Transaction_UP @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @CreditUP = '" + fmt.Sprintf("%v", params["jumlahup"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["TtdAccountOfficer"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["TtdKetuaKelompok"]) + "'")
	// }
	// db.Close()
}

type DetailMahasiswa struct {
	Name string
	Age  int
}

func PostListMahasiswa() []DetailMahasiswa {
	var datam []DetailMahasiswa

	// fmt.Sprintf("%v", datam)

	return datam
}

type LoginRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Apk_version string `json:"apk_version"`
}

type LoginResponse struct {
	ResponseStatus bool      `json:"responseStatus"`
	Status         string    `json:"status"`
	Message        string    `json:"message"`
	Token          string    `json:"token"`
	TokenExpired   string    `json:"tokenexpired"`
	Data           LoginData `json:"data"`
}

type LoginData struct {
	UnitKerja    string `json:"unitKerja"`
	Nip          string `json:"nip"`
	BranchName   string `json:"branchName"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	Name         string `json:"nama"`
	Jabatan      string `json:"jabatan"`
	LoginSession string `json:"LoginSession"`
}

type DateResponse struct {
	ResponseStatus bool   `json:"responseStatus"`
	CurrentDate    string `json:"currentDate"`
}

type TransactionAll struct {
	TransactionPKM []TransactionPost `json:"pkm"`
	TransactionUP  []TransactionUP   `json:"up"`
	Signature      Signatures        `json:"sign"`
}

type TransactionPost struct {
	AccountID         string `json:"AccountID"`
	ClientID          string `json:"ClientID"`
	GroupID           string `json:"GroupID"`
	InstallmentAmount string `json:"InstallmentAmount"`
	AttendStatus      string `json:"attendStatus"`
	Savings           string `json:"savings"`
	Trxdate           string `json:"trxdate"`
	UserName          string `json:"userName"`
	WithDraw          string `json:"withDraw"`
}

type TransactionUP struct {
	ClientID string `json:"clientid"`
	JumlahUP string `json:"jumlahup"`
}

type Signatures struct {
	GroupID           string `json:"GroupID"`
	IDKetuaKelompok   string `json:"IDKetuaKelompok"`
	Latitude          string `json:"latitude"`
	Longitude         string `json:"longitude"`
	Trxdate           string `json:"trxdate"`
	UserName          string `json:"userName"`
	TtdKetuaKelompok  string `json:"TtdKetuaKelompok"`
	TtdAccountOfficer string `json:"TtdAccountOfficer"`
}

type TransactionResponse struct {
	ResponseStatus bool
	Status         string
	Message        string
	Statusid       string
}

type IndividualTransaction struct {
	ListIndividualTransaction []ListTransaction `json:"data"`
	DetailTransaction         DetailTransaction `json:"detailTransaction"`
}

type ListTransaction struct {
	AccountID         string `json:"accountid"`
	ClientID          string `json:"clientid"`
	GroupID           string `json:"groupid"`
	BranchID          string `json:"cabangid"`
	InstallmentAmount string `json:"jumlahpar"`
	CreatedBy         string `json:"createdby"`
	AOSign            string `json:"AOSign"`
	ClientSign        string `json:"clientSign"`
	CreatedDate       string `json:"createdDate"`
}

type DetailTransaction struct {
	CreatedBy   string `json:"createdby"`
	BranchID    string `json:"cabangid"`
	CreatedDate string `json:"createdDate"`
}

type GetFailedIndividualTransaction struct {
	OurBranchID string `json:"ourBranchid"`
	GroupID     string `json:"groupId"`
	Initial     string `json:"initial"`
	GroupName   string `json:"groupName"`
	ClientID    string `json:"clientId"`
	AccountID   string `json:"accountId"`
	ClientName  string `json:"clientName"`
}

type IndividualTransactionResponse struct {
	ResponseStatus    bool
	Status            string
	Message           string
	Statusid          string
	FailedTransaction []GetFailedIndividualTransaction `json:"failedTransaction"`
}

type FailedData struct {
	ClientName   string `json:"clientName"`
	FailedStatus string `json:"failedStatus"`
}
