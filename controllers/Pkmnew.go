package controllers

import (
	"pkmapi/global"
	"pkmapi/models"

	"github.com/astaxie/beego"

	// "PKM_mekaar/conf"
	// "strings"
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// PKM controllerrr
type PkmController struct {
	beego.Controller
}

// @Title PKM
// @Description pkm
// @Param	OurBranchID path string  true "OurBranchID"
// @Param	AO_ID path string  true "AO_ID"
// @Success 200 models.ListCollection
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /GetCollectionList/:OurBranchID/:AO_ID [get]
func (c *PkmController) GetCollectionList() {
	Response := models.GetListCollectionNew(c.Ctx.Input.Param(":OurBranchID"), c.Ctx.Input.Param(":AO_ID"))

	c.Data["json"] = Response
	c.ServeJSON()
}

// @Title PKM
// @Description pkm
// @Param	OurBranchID path string  true "OurBranchID"
// @Param	AO_ID path string  true "AO_ID"
// @Success 200 models.ListCollection
// @Failure 403 "Error"
// @router /Shit/:OurBranchID/:AO_ID [get]
func (c *PkmController) Getup() {
	Response := models.GetListupNew(c.Ctx.Input.Param(":OurBranchID"), c.Ctx.Input.Param(":AO_ID"))

	c.Data["json"] = Response
	c.ServeJSON()
}

// @Title PKM
// @Description pkm
// @Param	OurBranchID path string  true "OurBranchID"
// @Param	AO_ID path string  true "USERNAME"
// @Success 200 models.ListCollection
// @Failure 403 "Error"
// @router /GetListGroup/:OurBranchID/:USERNAME [get]
func (c *PkmController) GroupList() {
	Response := models.GetListGroupNew(c.Ctx.Input.Param(":OurBranchID"), c.Ctx.Input.Param(":USERNAME"))

	c.Data["json"] = Response
	c.ServeJSON()
}

// // @Title PKM
// // @Description pkm
// // @Param	Username path string  true "Username"
// // @Param	Password path string  true "Password"
// // @Success 200 models.ListCollection
// // @Failure 403 "Error"
// // @router /SetLogin/:Username/:Password [get]
// func (c *PkmController) SetLogin() {
// 	Response := models.GetLogin(c.Ctx.Input.Param(":Username"), c.Ctx.Input.Param(":Password"))

// 	c.Data["json"] = Response
// 	c.ServeJSON()
// }

// @Title PKM
// @Description pkm
// @Param AccountID formData string true "AccountID"
// @Param Angsuran formData string true "Angsuran"
// @Param Attendance formData string true "Attendance"
// @Param ClientID formData string true "ClientID"
// @Param ClientName formData string true "ClientName"
// @Param GroupID formData string true "GroupID"
// @Param IDKetuaKelompok formData string true "IDKetuaKelompok"
// @Param KetuaKelompok formData string true "KetuaKelompok"
// @Param MeetingDay formData string true "MeetingDay"
// @Param ProductID formData string true "ProductID"
// @Param Setoran formData string true "Setoran"
// @Param Tarikan formData string true "Tarikan"
// @Param Titipan formData string true "Titipan"
// @Param TotalAngsuran formData string true "TotalAngsuran"
// @Param TotalSetor formData string true "TotalSetor"
// @Param TotalSetoran formData string true "TotalSetoran"
// @Param TotalTitipan formData string true "TotalTitipan"
// @Param TotalUP formData string true "TotalUP"
// @Param TtdAccountOfficer formData string true "TtdAccountOfficer"
// @Param TtdKetuaKelompok formData string true "TtdKetuaKelompok"
// @Param userName formData string true "userName"
// @Param CreatedDate formData string true "CreatedDate"
// @Success 200 models.ListCollection
// @Failure 403 "Error"
// @router /ReqShit [post]
func (c *PkmController) PostShit() {
	var params = make(map[string]interface{})

	for i := 0; i < len(c.GetStrings("AccountID")); i++ {
		params["AccountID"] = c.GetStrings("AccountID")[i]
		params["Angsuran"] = c.GetStrings("Angsuran")[i]
		params["Attendance"] = c.GetStrings("Attendance")[i]
		params["ClientID"] = c.GetStrings("ClientID")[i]
		params["ClientName"] = c.GetStrings("ClientName")[i]
		params["GroupID"] = c.GetStrings("GroupID")[i]
		params["IDKetuaKelompok"] = c.GetStrings("IDKetuaKelompok")[i]
		params["KetuaKelompok"] = c.GetStrings("KetuaKelompok")[i]
		params["MeetingDay"] = c.GetStrings("MeetingDay")[i]
		params["ProductID"] = c.GetStrings("ProductID")[i]
		params["Setoran"] = c.GetStrings("Setoran")[i]
		params["Tarikan"] = c.GetStrings("Tarikan")[i]
		params["Titipan"] = c.GetStrings("Titipan")[i]
		params["TotalAngsuran"] = c.GetStrings("TotalAngsuran")[i]
		params["TotalSetor"] = c.GetStrings("TotalSetor")[i]
		params["TotalSetoran"] = c.GetStrings("TotalSetoran")[i]
		params["TotalTitipan"] = c.GetStrings("TotalTitipan")[i]
		params["TotalUP"] = c.GetStrings("TotalUP")[i]
		params["TtdAccountOfficer"] = c.GetStrings("TtdAccountOfficer")[i]
		params["TtdKetuaKelompok"] = c.GetStrings("TtdKetuaKelompok")[i]
		params["userName"] = c.GetStrings("userName")[i]
		params["CreatedDate"] = c.GetStrings("CreatedDate")[i]

		models.PostTotalTransactionNew(params)
	}
	c.ServeJSON()
}

// @Title PKM
// @Description pkm
// @Param accountid formData string  true "accountid"
// @param AOSign formData string true "AOSign"
// @param accountid formData string true "accountid"
// @param cabangid formData string true "cabangid"
// @param clientSign formData string true "clientSign"
// @param clientid formData string true "clientid"
// @param createdby formData string true "createdby"
// @param groupid formData string true "groupid"
// @param jumlahpar formData string true "jumlahpar"
// @Success 200 models.ListCollection
// @Failure 403 "Error"
// @router /ReqShitPAR [post]
func (c *PkmController) PostShitPAR() {
	var params = make(map[string]interface{})

	for i := 0; i < len(c.GetStrings("accountid")); i++ {
		params["AOSign"] = c.GetStrings("AOSign")[i]
		params["accountid"] = c.GetStrings("accountid")[i]
		params["cabangid"] = c.GetStrings("cabangid")[i]
		params["clientSign"] = c.GetStrings("clientSign")[i]
		params["clientid"] = c.GetStrings("clientid")[i]
		params["createdby"] = c.GetStrings("createdby")[i]
		params["groupid"] = c.GetStrings("groupid")[i]
		params["jumlahpar"] = c.GetStrings("jumlahpar")[i]

		models.PostTotalTransactionPARNew(params)
	}
	c.ServeJSON()
}

// @Title PKM
// @Description pkm
// @Param	OurBranchID path string  true "OurBranchID"
// @Param	AO_ID path string  true "USERNAME"
// @Success 200 models.ListCollection
// @Failure 403 "Error"
// @router /GetCollectionListPAR/:OurBranchID/:AO_ID [get]
func (c *PkmController) GetCollectionListPAR() {
	Response := models.GetListCollectionPARNew(c.Ctx.Input.Param(":AO_ID"), c.Ctx.Input.Param(":OurBranchID"))
	// Response := models.Select_list()
	c.Data["json"] = Response
	c.ServeJSON()
}

// @Title PKM
// @Description pkm
// @Param	OurBranchID path string  true "OurBranchID"
// @Param	AO_ID path string  true "AO_ID"
// @Success 200 models.ListCollection
// @Failure 403 "Error"
// @router /GetCollectionListPKMIndividual/:OurBranchID/:AO_ID [get]
func (c *PkmController) GetCollectionListPKMIndividual() {
	Response := models.GetListPKMIndividualNew(c.Ctx.Input.Param(":AO_ID"), c.Ctx.Input.Param(":OurBranchID"))
	// Response := models.Select_list()
	c.Data["json"] = Response
	c.ServeJSON()
}

// @Title PKM
// @Description Auth Login
// @Param   data body models.LoginRequest true "login"
// @Success 200 {object} models.LoginResponse
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /AuthLogin [post]
func (c *PkmController) AuthLogin() {

	var dtpkm models.LoginRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &dtpkm)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		var response = models.LoginResponse{
			ResponseStatus: false,
			Status:         statusCode,
			Message:        err.Error(),
		}
		c.Data["json"] = response

		c.ServeJSON()
		return
	}

	db, err := global.ConnPKMGet()
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		var response = models.LoginResponse{
			ResponseStatus: false,
			Status:         statusCode,
			Message:        err.Error(),
		}
		c.Data["json"] = response

		c.ServeJSON()
		return
	}
	defer db.Close()

	rows, err := db.Query(`EXEC GET_User_Mobile_Only_II @Username = '` + dtpkm.Username + `', @Password = '` + b64.StdEncoding.EncodeToString([]byte(dtpkm.Password)) + `'`)

	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		var response = models.LoginResponse{
			ResponseStatus: false,
			Status:         statusCode,
			Message:        err.Error(),
		}
		c.Data["json"] = response

		c.ServeJSON()
		return
	}

	//FZL get token start
	// err = global.Oauth2()
	// if err!=nil{
	// 	c.Ctx.ResponseWriter.WriteHeader(401)
	// 	fmt.Printf(err.Error())
	// 	statcod := c.Ctx.ResponseWriter.Status
	// 	statusCode := strconv.Itoa(statcod)

	// 	var response = models.LoginResponse{
	// 		ResponseStatus: false,
	// 		Status:         statusCode,
	// 		Message:        err.Error(),
	// 	}
	// 	c.Data["json"] = response

	// 	c.ServeJSON()
	// }
	//FZL get token end

	defer rows.Close()

	var each = models.LoginData{}
	var dataArr []models.LoginData
	for rows.Next() {

		err = rows.Scan(&each.UnitKerja, &each.Nip, &each.BranchName, &each.UserName, &each.Password, &each.Name, &each.LoginSession, &each.Jabatan)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(500)
			fmt.Printf(err.Error())
			statcod := c.Ctx.ResponseWriter.Status
			statusCode := strconv.Itoa(statcod)

			var response = models.LoginResponse{
				ResponseStatus: false,
				Status:         statusCode,
				Message:        err.Error(),
			}
			c.Data["json"] = response

			c.ServeJSON()
			return
		}
		dataArr = append(dataArr, each)
	}

	if len(dataArr) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(404)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		var response = models.LoginResponse{
			ResponseStatus: false,
			Status:         statusCode,
			Message:        "Login Failed",
			Data:           each,
		}
		c.Data["json"] = response

		c.ServeJSON()
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	statcod := c.Ctx.ResponseWriter.Status
	statusCode := strconv.Itoa(statcod)

	var response = models.LoginResponse{
		ResponseStatus: true,
		Status:         statusCode,
		Message:        "Login Success",
		Token:          global.GenerateTokenJWT(dtpkm.Username + b64.StdEncoding.EncodeToString([]byte(dtpkm.Password))),
		Data:           each,
	}
	c.Data["json"] = response

	c.ServeJSON()

}

// @Title PKM
// @Description pkm
// @Success 200 models.DateResponse
// @Failure 403 "Error"
// @router /GetDate [get]
func (c *PkmController) GetNewDate() {

	current_time := time.Now().Local()

	var response = models.DateResponse{
		ResponseStatus: true,
		CurrentDate:    current_time.Format("2006-01-02"),
	}
	c.Data["json"] = response

	// fmt.Println(current_time.Format("2006-01-02"))

	c.ServeJSON()
}

// @Title PKM
// @Description pkm
// @Param   data body models.TransactionAll true "PostTransaction"
// @Success 200 {object} models.TransactionResponse
// @Failure 403 "Error"
// @router /PostTransaction [post]
func (c *PkmController) PostTransaction() models.TransactionResponse {

	decodeJson := json.NewDecoder(c.Ctx.Request.Body)

	servDate := GetDate()

	// fmt.Println(servDate)

	var dtpkm models.TransactionAll
	err := decodeJson.Decode(&dtpkm)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		fmt.Println("json decode " + err.Error())
		global.Logging("ERROR", "Decode Json ---> "+err.Error())

		var response = models.TransactionResponse{
			ResponseStatus: false,
			Status:         statusCode,
			Message:        "Accessing Database Error",
			Statusid:       "2",
		}
		c.Data["json"] = response

		c.ServeJSON()

		return response
	}

	var tranDate = dtpkm.Signature.Trxdate
	if servDate == tranDate {
		tranDate = dtpkm.Signature.Trxdate
	} else {
		tranDate = servDate
	}

	// fmt.Println("ini tanggal yang dipake " + tranDate)
	db, err := global.ConnPKMPost()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		fmt.Println("begin tran " + err.Error())
		global.Logging("ERROR", "Access Database Begin Transaction ---> "+err.Error())

		var response = models.TransactionResponse{
			ResponseStatus: false,
			Status:         statusCode,
			Message:        "Accessing Database Error",
			Statusid:       "2",
		}
		c.Data["json"] = response

		c.ServeJSON()

		return response
	}

	for i := 0; i < len(dtpkm.TransactionPKM); i++ {

		_, err = tx.ExecContext(ctx, `IF NOT EXISTS (SELECT TOP 1 * FROM PKM_LPM_Transaction_Individu WITH(NOLOCK) WHERE ClientID = '`+dtpkm.TransactionPKM[i].ClientID+`' AND AccountID = '`+dtpkm.TransactionPKM[i].AccountID+`' AND CAST(TrxDate as DATE) = CAST('`+tranDate+`' as DATE)) 
		BEGIN
			INSERT INTO PKM_LPM_Transaction
			SELECT
				NEWID(),
				Cabang_Mekaar.AreaID,
				PKM_LPM.OurBranchID,
				Initial,
				'`+dtpkm.TransactionPKM[i].GroupID+`',
				GroupName,
				MeetingDay,
				NULL,
				'`+dtpkm.TransactionPKM[i].ClientID+`',
				ClientName,
				'`+dtpkm.TransactionPKM[i].AccountID+`',
				NULL,
				ProductID,
				'`+dtpkm.TransactionPKM[i].AttendStatus+`',
				'`+dtpkm.TransactionPKM[i].InstallmentAmount+`',
				VolAccountID,
				'`+dtpkm.TransactionPKM[i].Savings+`',
				'`+dtpkm.TransactionPKM[i].WithDraw+`',
				'0',
				NULL,
				MeetingPlace,
				MeetingTime,
				AO_GROUP.MappingBy,
				AO_GROUP.MappingDate,
				NULL,
				NULL,
				NULL,
				NULL,
				'0',
				NULL,
				NULL,
				NULL,
				NULL,
				NULL,
				NULL,
				NULL,
				NULL,
				CASE
					WHEN AO_GROUP.AlternateUser IS NOT NULL THEN AO_GROUP.AlternateUser
					ELSE AO_GROUP.Username
				END,
				AO.NAMA,				
				CAST('`+tranDate+`' as DATETIME),
				CAST('`+tranDate+`' as DATETIME),
				NULL,
				0,
				CASE
					WHEN NoOfArrearDays = 0 THEN 'NO'
					ELSE 'YES'
				END PAR_Status,
				NULL,
				NULL,
				NULL,
				NULL,
				NULL,
				NULL,
				NULL
			FROM PKM_LPM WITH(NOLOCK) INNER JOIN AO_GROUP
			ON PKM_LPM.GroupID = AO_GROUP.GROUPID AND PKM_LPM.OurBranchID = AO_GROUP.OURBRANCHID
			INNER JOIN Cabang_Mekaar ON PKM_LPM.OurBranchID = Cabang_Mekaar.OurBranchID
			INNER JOIN AO ON AO_GROUP.Username = AO.USERNAME
			WHERE ClientID = '`+dtpkm.TransactionPKM[i].ClientID+`' AND AccountID = '`+dtpkm.TransactionPKM[i].AccountID+`'

			PRINT 'Successful'
		END
	ELSE
		BEGIN
			PRINT 'Cannot insert duplicate transaction with same Transaction Date'
		END`)

		if err != nil {
			// Incase we find any error in the query execution, rollback the transaction
			// tx.Rollback()
			c.Ctx.ResponseWriter.WriteHeader(400)
			statcod := c.Ctx.ResponseWriter.Status
			statusCode := strconv.Itoa(statcod)

			fmt.Println("exec pkm " + err.Error())
			global.Logging("ERROR", "Transaction PKM "+dtpkm.TransactionPKM[i].GroupID+"  ---> "+err.Error())

			errrollback := tx.Rollback()
			if errrollback != nil {
				var response = models.TransactionResponse{
					ResponseStatus: false,
					Status:         statusCode,
					Message:        errrollback.Error(),
					Statusid:       "3",
				}
				c.Data["json"] = response

				c.ServeJSON()

				return response
			}

			var response = models.TransactionResponse{
				ResponseStatus: false,
				Status:         statusCode,
				Message:        "Submit PKM to Database Error",
				Statusid:       "3",
			}
			c.Data["json"] = response

			c.ServeJSON()

			return response
		}
	}

	// for a := 0; a < len(dtpkm.TransactionUP); a++ {
	// 	// Create a new context, and begin a transaction
	// 	// Here, the query is executed on the transaction instance, and not applied to the database yet
	// 	_, err = tx.ExecContext(ctx, `IF NOT EXISTS (SELECT * FROM PKM_LPM WITH(NOLOCK) WHERE ClientID = '`+dtpkm.TransactionUP[a].ClientID+`')
	// 		BEGIN

	// 			INSERT INTO PKM_LPM_Transaction_UP
	// 			SELECT
	// 				NEWID(),
	// 				b.AreaID,
	// 				a.OurBranchID,
	// 				b.ShortName,
	// 				a.GroupID,
	// 				a.GroupName,
	// 				a.MeetingDay,
	// 				'`+dtpkm.TransactionUP[a].ClientID+`',
	// 				a.AccountID,
	// 				a.ClientName,
	// 				'`+dtpkm.TransactionUP[a].JumlahUP+`',
	// 				'`+dtpkm.TransactionUP[a].JumlahUP+`',
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				'0',
	// 				ISNULL(c.AlternateUser, c.Username),
	// 				GETDATE(),
	// 				CAST(GETDATE() as DATETIME),
	// 				NULL,
	// 				NULL,
	// 				NULL
	// 			FROM PKM_LPM_UP a WITH(NOLOCK) INNER JOIN Cabang_Mekaar b
	// 			ON a.OurBranchID = b.OurBranchID
	// 			INNER JOIN AO_GROUP c
	// 			ON a.GroupID = c.GROUPID
	// 			INNER JOIN AO d
	// 			ON c.Username = d.USERNAME
	// 			WHERE ClientID = '`+dtpkm.TransactionUP[a].ClientID+`'

	// 			INSERT INTO PKM_LPM_Transaction
	// 			SELECT
	// 				NEWID(),
	// 				Cabang_Mekaar.AreaID,
	// 				PKM_LPM_UP.OurBranchID,
	// 				Cabang_Mekaar.ShortName,
	// 				PKM_LPM_UP.GroupID,
	// 				GroupName,
	// 				MeetingDay,
	// 				NULL,
	// 				'`+dtpkm.TransactionUP[a].ClientID+`',
	// 				ClientName,
	// 				AccountID,
	// 				NULL,
	// 				'UP',
	// 				'1',
	// 				'0',
	// 				'0',
	// 				'0',
	// 				'0',
	// 				'0',
	// 				NULL,
	// 				MeetingPlace,
	// 				MeetingTime,
	// 				AO_GROUP.MappingBy,
	// 				AO_GROUP.MappingDate,
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				'0',
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				CASE
	// 					WHEN AO_GROUP.AlternateUser IS NOT NULL THEN AO_GROUP.AlternateUser
	// 					ELSE AO_GROUP.Username
	// 				END,
	// 				AO.NAMA,
	// 				CAST(GETDATE() as DATETIME),
	// 				CAST(GETDATE() as DATETIME),
	// 				NEWID(),
	// 				0,
	// 				'NO',
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				NULL
	// 			FROM PKM_LPM_UP WITH(NOLOCK) INNER JOIN AO_GROUP
	// 			ON PKM_LPM_UP.GroupID = AO_GROUP.GROUPID AND PKM_LPM_UP.OurBranchID = AO_GROUP.OURBRANCHID
	// 			INNER JOIN Cabang_Mekaar ON PKM_LPM_UP.OurBranchID = Cabang_Mekaar.OurBranchID
	// 			INNER JOIN AO ON AO_GROUP.Username = AO.USERNAME
	// 			WHERE ClientID = '`+dtpkm.TransactionUP[a].ClientID+`'
	// 		END
	// 	ELSE
	// 		BEGIN
	// 			INSERT INTO PKM_LPM_Transaction_UP
	// 			SELECT
	// 				NEWID(),
	// 				b.AreaID,
	// 				a.OurBranchID,
	// 				b.ShortName,
	// 				a.GroupID,
	// 				a.GroupName,
	// 				a.MeetingDay,
	// 				'`+dtpkm.TransactionUP[a].ClientID+`',
	// 				a.AccountID,
	// 				a.ClientName,
	// 				'`+dtpkm.TransactionUP[a].JumlahUP+`',
	// 				'`+dtpkm.TransactionUP[a].JumlahUP+`',
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				NULL,
	// 				'0',
	// 				ISNULL(c.AlternateUser, c.Username),
	// 				GETDATE(),
	// 				CAST(GETDATE() as DATETIME),
	// 				NULL,
	// 				NULL,
	// 				NULL
	// 			FROM PKM_LPM_UP a WITH(NOLOCK) INNER JOIN Cabang_Mekaar b
	// 			ON a.OurBranchID = b.OurBranchID
	// 			INNER JOIN AO_GROUP c
	// 			ON a.GroupID = c.GROUPID
	// 			INNER JOIN AO d
	// 			ON c.Username = d.USERNAME
	// 			WHERE ClientID = '`+dtpkm.TransactionUP[a].ClientID+`'
	// 		END`)

	// 	if err != nil {
	// 		// Incase we find any error in the query execution, rollback the transaction
	// 		tx.Rollback()
	// 		c.Ctx.ResponseWriter.WriteHeader(400)
	// 		statcod := c.Ctx.ResponseWriter.Status
	// 		statusCode := strconv.Itoa(statcod)

	// 		fmt.Println(err.Error())
	// 		global.Logging("ERROR Transaction UP " + dtpkm.TransactionUP[a].ClientID + "  ---> " + err.Error())

	// 		var response = models.TransactionResponse{
	// 			ResponseStatus: false,
	// 			Status:         statusCode,
	// 			Message:        "Submit UP to Database Error",
	// 			Statusid:       "3",
	// 		}
	// 		c.Data["json"] = response

	// 		c.ServeJSON()

	// 		return response
	// 	}
	// }

	_, err = tx.ExecContext(ctx, `BEGIN
		INSERT INTO PKM_SIGNATURES VALUES (
			(SELECT TOP 1 OurBranchID FROM AO_GROUP WITH(NOLOCK) WHERE GROUPID = '`+dtpkm.Signature.GroupID+`'),
			'`+dtpkm.Signature.GroupID+`',
			'`+tranDate+`',
			'`+dtpkm.Signature.TtdAccountOfficer+`',
			'`+dtpkm.Signature.TtdKetuaKelompok+`',
			NULL,
			NULL,
			NEWID(),
			'`+dtpkm.Signature.Latitude+`',
			'`+dtpkm.Signature.Longitude+`'
		)
	END`)

	fmt.Println(`BEGIN
	INSERT INTO PKM_SIGNATURES VALUES (
		(SELECT TOP 1 OurBranchID FROM AO_GROUP WITH(NOLOCK) WHERE GROUPID = '` + dtpkm.Signature.GroupID + `'),
		'` + dtpkm.Signature.GroupID + `',
		'` + tranDate + `',
		'` + dtpkm.Signature.TtdAccountOfficer + `',
		'` + dtpkm.Signature.TtdKetuaKelompok + `',
		NULL,
		NULL,
		NEWID(),
		'` + dtpkm.Signature.Latitude + `',
		'` + dtpkm.Signature.Longitude + `'
	)
	END`)

	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		// tx.Rollback()

		c.Ctx.ResponseWriter.WriteHeader(400)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		fmt.Println("exec signature " + err.Error())
		// global.Logging("ERROR Transaction SIGNATURE " + dtpkm.Signature.GroupID + "  ---> " + err.Error())
		global.Logging("INFO", `BEGIN
		INSERT INTO PKM_SIGNATURES VALUES (
			(SELECT TOP 1 OurBranchID FROM AO_GROUP WITH(NOLOCK) WHERE GROUPID = '`+dtpkm.Signature.GroupID+`'),
			'`+dtpkm.Signature.GroupID+`',
			'`+tranDate+`',
			'`+dtpkm.Signature.TtdAccountOfficer+`',
			'`+dtpkm.Signature.TtdKetuaKelompok+`',
			NULL,
			NULL,
			NEWID(),
			'`+dtpkm.Signature.Latitude+`',
			'`+dtpkm.Signature.Longitude+`'
			)
		END`)

		errrollback := tx.Rollback()
		if errrollback != nil {
			var response = models.TransactionResponse{
				ResponseStatus: false,
				Status:         statusCode,
				Message:        errrollback.Error(),
				Statusid:       "3",
			}
			c.Data["json"] = response

			c.ServeJSON()

			return response
		}

		var response = models.TransactionResponse{
			ResponseStatus: false,
			Status:         statusCode,
			Message:        "Submit Signature to Database Error",
			Statusid:       "3",
		}
		c.Data["json"] = response

		c.ServeJSON()

		return response
	}

	err = tx.Commit()
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(400)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		global.Logging("ERROR", "Commit Transaction "+dtpkm.Signature.GroupID+"  ---> "+err.Error())

		var response = models.TransactionResponse{
			ResponseStatus: false,
			Status:         statusCode,
			Message:        "Submit Data to Database Error",
			Statusid:       "3",
		}
		c.Data["json"] = response

		c.ServeJSON()

		return response
	}

	// defer
	c.Ctx.ResponseWriter.WriteHeader(200)
	statcod := c.Ctx.ResponseWriter.Status
	statusCode := strconv.Itoa(statcod)

	var response = models.TransactionResponse{
		ResponseStatus: true,
		Status:         statusCode,
		Message:        "Submit PKM Berhasil",
		Statusid:       "1",
	}
	c.Data["json"] = response

	c.ServeJSON()

	return response
}
