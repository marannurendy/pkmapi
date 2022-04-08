package models

import (
	"fmt"
	"pkmapi/global"
)

func PostTotalTransactionNew(params map[string]interface{}) {
	// fmt.Println(
	//
	db, err := global.ConnPKMPost()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("EXEC ADD_PKM_Transaction @GroupID = '" + fmt.Sprintf("%v", params["GroupID"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["ClientID"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["AccountID"]) + "', @AttendStatus = '" + fmt.Sprintf("%v", params["Attendance"]) + "' , @InstallmentAmount = '" + fmt.Sprintf("%v", params["Angsuran"]) + "', @CreditSaving = '" + fmt.Sprintf("%v", params["Tarikan"]) + "', @DebitSaving = '" + fmt.Sprintf("%v", params["Setoran"]) + "', @CreditUP = '0', @TTD1 = '" + fmt.Sprintf("%v", params["TtdAccountOfficer"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["TtdKetuaKelompok"]) + "', @CreatedDate = '" + fmt.Sprintf("%v", params["CreatedDate"]) + "'")
	fmt.Println("EXEC ADD_PKM_Transaction @GroupID = '" + fmt.Sprintf("%v", params["GroupID"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["ClientID"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["AccountID"]) + "', @AttendStatus = '" + fmt.Sprintf("%v", params["Attendance"]) + "' , @InstallmentAmount = '" + fmt.Sprintf("%v", params["Angsuran"]) + "', @CreditSaving = '" + fmt.Sprintf("%v", params["Tarikan"]) + "', @DebitSaving = '" + fmt.Sprintf("%v", params["Setoran"]) + "', @CreditUP = '0'")

	global.Logging("INFO", "EXEC ADD_PKM_Transaction @GroupID = '"+fmt.Sprintf("%v", params["GroupID"])+"', @ClientID = '"+fmt.Sprintf("%v", params["ClientID"])+"', @AccountID = '"+fmt.Sprintf("%v", params["AccountID"])+"', @AttendStatus = '"+fmt.Sprintf("%v", params["Attendance"])+"' , @InstallmentAmount = '"+fmt.Sprintf("%v", params["Angsuran"])+"', @CreditSaving = '"+fmt.Sprintf("%v", params["Tarikan"])+"', @DebitSaving = '"+fmt.Sprintf("%v", params["Setoran"])+"', @CreditUP = '0'")

	if err != nil {
		fmt.Println(err)
		global.Logging("ERROR", "PKM EXEC ADD_PKM_Transaction @GroupID = '"+fmt.Sprintf("%v", params["GroupID"])+"', @ClientID = '"+fmt.Sprintf("%v", params["ClientID"])+"', @AccountID = '"+fmt.Sprintf("%v", params["AccountID"])+"', @AttendStatus = '"+fmt.Sprintf("%v", params["Attendance"])+"' , @InstallmentAmount = '"+fmt.Sprintf("%v", params["Angsuran"])+"', @CreditSaving = '"+fmt.Sprintf("%v", params["Tarikan"])+"', @DebitSaving = '"+fmt.Sprintf("%v", params["Setoran"])+"', @CreditUP = '0', @TTD1 = '"+fmt.Sprintf("%v", params["TtdAccountOfficer"])+"', @TTD2 = '"+fmt.Sprintf("%v", params["TtdKetuaKelompok"])+"', @CreatedDate = '"+fmt.Sprintf("%v", params["CreatedDate"])+"' ---> "+err.Error())
	}
	// else {
	// 	fmt.Println("TRUE")
	// 	// fmt.Println(path)
	// 	global.SuccessLog("SUCCESS POST PKM EXEC ADD_PKM_Transaction @GroupID = '" + fmt.Sprintf("%v", params["GroupID"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["ClientID"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["AccountID"]) + "', @AttendStatus = '" + fmt.Sprintf("%v", params["Attendance"]) + "' , @InstallmentAmount = '" + fmt.Sprintf("%v", params["Angsuran"]) + "', @CreditSaving = '" + fmt.Sprintf("%v", params["Tarikan"]) + "', @DebitSaving = '" + fmt.Sprintf("%v", params["Setoran"]) + "', @CreditUP = '0', @TTD1 = '" + fmt.Sprintf("%v", params["TtdAccountOfficer"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["TtdKetuaKelompok"]) + "', @CreatedDate = '" + fmt.Sprintf("%v", params["CreatedDate"]) + "'")
	// }
	// db.Close()
}

func PostTotalTransactionPARNew(params map[string]interface{}) {
	// fmt.Println(
	//
	db, err := global.ConnPKMPost()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("EXEC ADD_PKM_Transaction_PAR @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitPAR = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["clientSign"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["AOSign"]) + "', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "'")
	fmt.Println("EXEC ADD_PKM_Transaction_PAR @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitPAR = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["clientSign"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["AOSign"]) + "', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "'")

	global.Logging("INFO", "EXEC ADD_PKM_Transaction_PAR @GroupID = '"+fmt.Sprintf("%v", params["groupid"])+"', @ClientID = '"+fmt.Sprintf("%v", params["clientid"])+"', @AccountID = '"+fmt.Sprintf("%v", params["accountid"])+"', @DebitPAR = '"+fmt.Sprintf("%v", params["jumlahpar"])+"', @TTD1 = '"+fmt.Sprintf("%v", params["clientSign"])+"', @TTD2 = '"+fmt.Sprintf("%v", params["AOSign"])+"', @CreatedBy = '"+fmt.Sprintf("%v", params["createdby"])+"'")

	if err != nil {
		fmt.Println(err)
		global.Logging("ERROR", "PAR EXEC ADD_PKM_Transaction_PAR @GroupID = '"+fmt.Sprintf("%v", params["groupid"])+"', @ClientID = '"+fmt.Sprintf("%v", params["clientid"])+"', @AccountID = '"+fmt.Sprintf("%v", params["accountid"])+"', @DebitPAR = '"+fmt.Sprintf("%v", params["jumlahpar"])+"', @TTD1 = '"+fmt.Sprintf("%v", params["clientSign"])+"', @TTD2 = '"+fmt.Sprintf("%v", params["AOSign"])+"', @CreatedBy = '"+fmt.Sprintf("%v", params["createdby"])+"' ---> "+err.Error())
	}
	// else {
	// 	fmt.Println("TRUE")
	// 	global.SuccessLog("SUCCESS POST PAR EXEC ADD_PKM_Transaction_PAR @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitPAR = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["clientSign"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["AOSign"]) + "', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "'")
	// }
	// db.Close()
}

type ListCollectionNew struct {
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

func GetListCollectionNew(OurBranchID string, AO_ID string) []ListCollectionNew {
	var data []ListCollectionNew

	db, err := global.ConnPKMPost()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	fmt.Println("EXEC GET_PKM_ANGSURAN @AO_ID = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'Angsuran'")
	global.Logging("INFO", "EXEC GET_PKM_ANGSURAN @AO_ID = '"+AO_ID+"', @OurBranchID = '"+OurBranchID+"', @type = 'Angsuran'")

	rows, err := db.Query("EXEC GET_PKM_ANGSURAN @AO_ID = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'Angsuran'")

	if err != nil {
		// fmt.Println(err.Error())
		fmt.Println("ERROR")
		global.Logging("ERROR", "GET PKM EXEC GET_PKM_ANGSURAN @AO_ID = '"+AO_ID+"', @OurBranchID = '"+OurBranchID+"', @type = 'Angsuran' ---> "+err.Error())
	}

	fmt.Println("dibawah")
	// else {
	// 	global.SuccessLog("SUCCESS GET EXEC GET_PKM_ANGSURAN @AO_ID = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'Angsuran' | MATCH")
	// }
	defer rows.Close()
	for rows.Next() {
		var each = ListCollectionNew{}
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

type ListupNew struct {
	OurBranchID    string `json:"OurBranchID"`
	ClientID       string `json:"ClientID"`
	ClientName     string `json:"ClientName"`
	GroupID        string `json:"GroupID"`
	GroupName      string `json:"GroupName"`
	MeetingDay     string `json:"MeetingDay"`
	CompSavingsBal string `json:"CompSavingsBal"`
}

func GetListupNew(OurBranchID string, AO_ID string) []ListupNew {
	var data []ListupNew
	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("EXEC GET_PKM_ANGSURAN @AO_ID = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'UP'")
	global.Logging("INFO", "EXEC GET_PKM_ANGSURAN @AO_ID = '"+AO_ID+"', @OurBranchID = '"+OurBranchID+"', @type = 'UP'")
	if err != nil {
		fmt.Println(err.Error())
		global.Logging("ERROR", "GET UP EXEC GET_PKM_ANGSURAN @AO_ID = '"+AO_ID+"', @OurBranchID = '"+OurBranchID+"', @type = 'UP' ---> "+err.Error())
	}
	// else {
	// 	global.SuccessLog("SUCCESS GET UP EXEC GET_PKM_ANGSURAN @AO_ID = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'UP' | MATCH")
	// }
	defer rows.Close()
	for rows.Next() {
		var each = ListupNew{}
		err = rows.Scan(&each.OurBranchID, &each.ClientID, &each.ClientName, &each.GroupID, &each.GroupName, &each.MeetingDay, &each.CompSavingsBal)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}
	// db.Close()
	return data
}

type GroupListNew struct {
	OurBranchID   string `json:"OurBranchID"`
	GroupID       string `json:"GroupID"`
	GroupName     string `json:"GroupName"`
	MeetingDay    string `json:"MeetingDay"`
	AnggotaAktif  string `json:"AnggotaAktif"`
	JumlahTagihan string `json:"JumlahTagihan"`
	MeetingPlace  string `json:"MeetingPlace"`
	MeetingTime   string `json:"MeetingTime"`
}

func GetListGroupNew(OurBranchID string, USERNAME string) []GroupListNew {
	var data []GroupListNew
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
		global.Logging("ERROR", "EXEC GET_Detail_PKM_LPM @OurBranchID='"+OurBranchID+"', @USERNAME='"+USERNAME+"', @type= 'Process' --->> "+err.Error())
	}
	// else {
	// 	global.SuccessLog("SUCCESS GET EXEC GET_Detail_PKM_LPM @OurBranchID='" + OurBranchID + "', @USERNAME='" + USERNAME + "', @type= 'Process' | MATCH")
	// }

	defer rows.Close()
	for rows.Next() {
		var each = GroupListNew{}
		err = rows.Scan(&each.OurBranchID, &each.GroupID, &each.GroupName, &each.MeetingDay, &each.AnggotaAktif, &each.JumlahTagihan, &each.MeetingPlace, &each.MeetingTime)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}
	// db.Close()
	return data
}

type DetailPKM_PARNew struct {
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
func GetListCollectionPARNew(AO_ID string, OurBranchID string) []DetailPKM_PARNew {
	var data []DetailPKM_PARNew
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
		var each = DetailPKM_PARNew{}
		err = rows.Scan(&each.OurBranchID, &each.ClientID, &each.ClientName, &each.AccountID, &each.ProductID, &each.GroupID, &each.GroupName, &each.ODAmount)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}
	// db.Close()
	return data
}

type DetailGetPKMIndividualNew struct {
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
func GetListPKMIndividualNew(AO_ID string, OurBranchID string) []DetailGetPKMIndividualNew {
	var data []DetailGetPKMIndividualNew
	db, err := global.ConnPKMGet()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("EXEC GET_Detail_PKM_LPM_PAR_INDIVIDUAL @USERNAME = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'Process'")

	if err != nil {
		fmt.Println(err.Error())
		global.Logging("ERROR", "PKM Individu EXEC GET_Detail_PKM_LPM_PAR_INDIVIDUAL @USERNAME = '"+AO_ID+"', @OurBranchID = '"+OurBranchID+"', @type = 'Process' ---> "+err.Error())
	}
	// else {
	// 	global.SuccessLog("SUCCESS PKM Individu EXEC GET_Detail_PKM_LPM_PAR_INDIVIDUAL @USERNAME = '" + AO_ID + "', @OurBranchID = '" + OurBranchID + "', @type = 'Process' | MATCH")
	// }
	defer rows.Close()
	for rows.Next() {
		var each = DetailGetPKMIndividualNew{}
		err = rows.Scan(&each.OurBranchID, &each.GroupID, &each.GroupName, &each.MeetingDay, &each.AccountID, &each.ProductID, &each.ClientID, &each.ClientName, &each.InstallmentAmount, &each.Rill, &each.Ke, &each.VolSavingsBal, &each.PARStatus)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}
	// db.Close()
	return data
}

func PostPKMIndividualNew(params map[string]interface{}) {
	// fmt.Println(
	//
	db, err := global.ConnPKMPost()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("EXEC ADD_PKM_Transaction_PAR_Individual @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitValue = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["clientSign"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["AOSign"]) + "', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "', @CreatedDate = '" + fmt.Sprintf("%v", params["CreatedDate"]) + "'")
	fmt.Println("EXEC ADD_PKM_Transaction_PAR_Individual @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitValue = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["clientSign"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["AOSign"]) + "', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "', @CreatedDate = '" + fmt.Sprintf("%v", params["CreatedDate"]) + "'")

	if err != nil {
		fmt.Println(err)
		global.Logging("ERROR", "POST PKM Individu EXEC ADD_PKM_Transaction_PAR_Individual @GroupID = '"+fmt.Sprintf("%v", params["groupid"])+"', @ClientID = '"+fmt.Sprintf("%v", params["clientid"])+"', @AccountID = '"+fmt.Sprintf("%v", params["accountid"])+"', @DebitValue = '"+fmt.Sprintf("%v", params["jumlahpar"])+"', @TTD1 = '"+fmt.Sprintf("%v", params["clientSign"])+"', @TTD2 = '"+fmt.Sprintf("%v", params["AOSign"])+"', @CreatedBy = '"+fmt.Sprintf("%v", params["createdby"])+"', @CreatedDate = '"+fmt.Sprintf("%v", params["CreatedDate"])+"' ---> "+err.Error())
	}
	// else {
	// 	fmt.Println("TRUE")
	// 	global.SuccessLog("SUCCESS POST PKM Individu EXEC ADD_PKM_Transaction_PAR_Individual @GroupID = '" + fmt.Sprintf("%v", params["groupid"]) + "', @ClientID = '" + fmt.Sprintf("%v", params["clientid"]) + "', @AccountID = '" + fmt.Sprintf("%v", params["accountid"]) + "', @DebitValue = '" + fmt.Sprintf("%v", params["jumlahpar"]) + "', @TTD1 = '" + fmt.Sprintf("%v", params["clientSign"]) + "', @TTD2 = '" + fmt.Sprintf("%v", params["AOSign"]) + "', @CreatedBy = '" + fmt.Sprintf("%v", params["createdby"]) + "', @CreatedDate = '" + fmt.Sprintf("%v", params["CreatedDate"]) + "'")
	// }
	// db.Close()
}
