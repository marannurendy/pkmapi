package controllers

import (
	"pkmapi/models"
	"pkmapi/global"
	"github.com/astaxie/beego"
	// "PKM_mekaar/conf"
	// "strings"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"	
	"strconv"
	"time"
	"context"
)

// PKM controllerrr
type MainController struct {
	beego.Controller
}

func (c *MainController) GetList() {
	Response := models.Select_list(c.Ctx.Input.Param(":MeetingDayID"), c.Ctx.Input.Param(":AO_ID"), c.Ctx.Input.Param(":type"), c.Ctx.Input.Param(":GroupID"))
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
// @Failure 500 "Error"
// @router /GetCollectionList/:OurBranchID/:AO_ID [get]
func (c *MainController) GetCollectionList() {
	Response := models.GetListCollection(c.Ctx.Input.Param(":OurBranchID"), c.Ctx.Input.Param(":AO_ID"))

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
func (c *MainController) Getup() {
	Response := models.GetListup(c.Ctx.Input.Param(":OurBranchID"), c.Ctx.Input.Param(":AO_ID"))

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
func (c *MainController) GroupList() {
	Response := models.GetListGroup(c.Ctx.Input.Param(":OurBranchID"), c.Ctx.Input.Param(":USERNAME"))

	c.Data["json"] = Response
	c.ServeJSON()
}

// @Title PKM
// @Description pkm
// @Param	Username path string  true "Username"
// @Param	Password path string  true "Password"
// @Success 200 models.ListCollection
// @Failure 403 "Error"
// @router /SetLogin/:Username/:Password [get]
func (c *MainController) SetLogin() {
	Response := models.GetLogin(c.Ctx.Input.Param(":Username"), c.Ctx.Input.Param(":Password"))

	c.Data["json"] = Response
	c.ServeJSON()
}

func (c *MainController) SetLogout() {
	Response := models.GetLogout(c.Ctx.Input.Param(":Username"), c.Ctx.Input.Param(":Password"))

	c.Data["json"] = Response
	c.ServeJSON()
}

func (c *MainController) Getdata() {
	Response := models.Select_data()
	// Response := models.Select_list()
	c.Data["json"] = Response
	c.ServeJSON()
}

func (c *MainController) GetTest() {
	Response := models.SelectTest()
	// Response := models.Select_list()
	c.Data["json"] = Response
	c.ServeJSON()
}

func (c *MainController) GetShit() {
	Response := models.SelectShit()
	c.Data["json"] = Response
	c.ServeJSON()
}

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
func (c *MainController) PostShit() {
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

		models.PostTotalTransaction(params)
	}
	c.ServeJSON()
}

func (c *MainController) PostShitUP() {
	var params = make(map[string]interface{})

	for i := 0; i < len(c.GetStrings("clientid")); i++ {
		params["clientid"] = c.GetStrings("clientid")[i]
		params["jumlahup"] = c.GetStrings("jumlahup")[i]

		models.PostTotalTransactionUP(params)
	}
	c.ServeJSON()
}

func (c *MainController) PostShitUPnew() {
	var params = make(map[string]interface{})

	for i := 0; i < len(c.GetStrings("clientid")); i++ {
		params["clientid"] = c.GetStrings("clientid")[i]
		params["jumlahup"] = c.GetStrings("jumlahup")[i]
		params["TtdAccountOfficer"] = c.GetStrings("TtdAccountOfficer")[i]
		params["TtdKetuaKelompok"] = c.GetStrings("TtdKetuaKelompok")[i]

		models.PostTotalTransactionUPnew(params)
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
func (c *MainController) PostShitPAR() {
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

		models.PostTotalTransactionPAR(params)
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
func (c *MainController) GetCollectionListPAR() {
	Response := models.GetListCollectionPAR(c.Ctx.Input.Param(":AO_ID"), c.Ctx.Input.Param(":OurBranchID"))
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
func (c *MainController) GetCollectionListPKMIndividual() {
	Response := models.GetListPKMIndividual(c.Ctx.Input.Param(":AO_ID"), c.Ctx.Input.Param(":OurBranchID"))
	// Response := models.Select_list()
	c.Data["json"] = Response
	c.ServeJSON()
}

// @Title PKM
// @Description pkm
// @Param	accountid formData string  true "accountid"
// @Param AOSign formData string true "AOSign"
// @Param accountid formData string true "accountid"
// @Param cabangid formData string true "cabangid"
// @Param clientSign formData string true "clientSign"
// @Param clientid formData string true "clientid"
// @Param createdby formData string true "createdby"
// @Param groupid formData string true "groupid"
// @Param jumlahpar formData string true "jumlahpar"
// @Param CreatedDate formData string true "CreatedDate"
// @Success 200 models.ListCollection
// @Failure 403 "Error"
// @router /PostPKMIndividual [post]
func (c *MainController) PostPKMIndividual() {
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
		params["CreatedDate"] = c.GetStrings("CreatedDate")[i]

		models.PostPKMIndividual(params)
	}
	c.ServeJSON()
}

// @Title PKM
// @Description Auth Login
// @Param   data body models.LoginRequest true "login"
// @Success 200 {object} models.LoginResponse
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /AuthLogin [post]
func (c *MainController) AuthLogin() {
	var dtpkm models.LoginRequest		
    err := json.Unmarshal(c.Ctx.Input.RequestBody, &dtpkm)	
	if err != nil {
		c.ServeJSON()
	}

	db,err := global.ConnPKM()
	if err != nil {
		fmt.Println(err.Error())
	}	
	defer db.Close()
	// global.Logging("INFO_VERSI","Mobile Version '" + dtpkm.Apk_version + "' API Version '"+beego.AppConfig.String("Version")+"' username='"+dtpkm.Username+"' ")

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

		err = rows.Scan(&each.UnitKerja, &each.Nip, &each.BranchName, &each.UserName, &each.Password, &each.Name, &each.LoginSession, &each.Jabatan )
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
		Token: 			global.GenerateTokenJWT(dtpkm.Username+b64.StdEncoding.EncodeToString([]byte(dtpkm.Password))),
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
func (c *MainController) GetNewDate() {

	// err := global.Upload_file_s3()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	current_time := time.Now().Local()

	var response = models.DateResponse{
		ResponseStatus: true,
		CurrentDate:    current_time.Format("2006-01-02"),
	}
	c.Data["json"] = response

	// fmt.Println(current_time.Format("2006-01-02"))

	c.ServeJSON()
}


func GetDate() string {
	current_time := time.Now().Local()

	var response = models.DateResponse{
		ResponseStatus: true,
		CurrentDate:    current_time.Format("2006-01-02"),
	}

	// fmt.Println(current_time.Format("2006-01-02"))
	return response.CurrentDate
}

// @Title PKM
// @Description pkm
// @Param   data body models.TransactionAll true "PostTransaction"
// @Success 200 {object} models.TransactionResponse
// @Failure 403 "Error"
// @router /PostTransaction [post]
func (c *MainController) PostTransaction() models.TransactionResponse {

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
		global.Logging("ERROR_PKM","Decode Json ---> " + err.Error())

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

	db,err := global.ConnPKM()
	if err != nil {
		fmt.Println(err.Error())
	}	
	defer db.Close()
	// fmt.Println("ini tanggal yang dipake " + tranDate)

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		fmt.Println("begin tran " + err.Error())
		global.Logging("ERROR_PKM","Access Database Begin Transaction ---> " + err.Error())

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

	var queryexec,querysign string
	for i := 0; i < len(dtpkm.TransactionPKM); i++ {

		if (dtpkm.TransactionPKM[i].Savings == "null"){
			dtpkm.TransactionPKM[i].Savings = "0"
		}

		if (dtpkm.TransactionPKM[i].InstallmentAmount == "null"){
			dtpkm.TransactionPKM[i].InstallmentAmount = "0"
		}
		

		queryexec = `IF NOT EXISTS (SELECT TOP 1 * FROM PKM_LPM_Transaction_Individu WITH(NOLOCK) WHERE ClientID = '`+dtpkm.TransactionPKM[i].ClientID+`' AND AccountID = '`+dtpkm.TransactionPKM[i].AccountID+`' AND CAST(TrxDate as DATE) = CAST('`+tranDate+`' as DATE)) 
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
		END`

		_, err = tx.ExecContext(ctx, queryexec)
		// global.Logging("INFO Transaction PKM --> " + queryexec)
		if err != nil {
			// Incase we find any error in the query execution, rollback the transaction
			// tx.Rollback()
			c.Ctx.ResponseWriter.WriteHeader(400)
			statcod := c.Ctx.ResponseWriter.Status
			statusCode := strconv.Itoa(statcod)

			fmt.Println("exec pkm " + err.Error())
			global.Logging("ERROR_PKM","Transaction "+queryexec+" --->" + dtpkm.TransactionPKM[i].GroupID + "  ---> " + err.Error())


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
	
	// querysign = `BEGIN
	// 	INSERT INTO PKM_SIGNATURES VALUES (
	// 		(SELECT TOP 1 OurBranchID FROM AO_GROUP WITH(NOLOCK) WHERE GROUPID = '`+dtpkm.Signature.GroupID+`'),
	// 		'`+dtpkm.Signature.GroupID+`',
	// 		'`+tranDate+`',
	// 		'`+dtpkm.Signature.TtdAccountOfficer+`',
	// 		'`+dtpkm.Signature.TtdKetuaKelompok+`',
	// 		NULL,
	// 		NULL,
	// 		NEWID(),
	// 		'`+dtpkm.Signature.Latitude+`',
	// 		'`+dtpkm.Signature.Longitude+`'
	// 	)
	// 	END`
	
	var fileimages = make(map[string]string)
	fileimages["TtdAccountOfficer"] = dtpkm.Signature.TtdAccountOfficer
	fileimages["TtdKetuaKelompok"] = dtpkm.Signature.TtdKetuaKelompok

	fileimages,err = global.MapB64SaveFile(fileimages,beego.AppConfig.String("pathimages"))
	if err != nil { 			
		// tx.Rollback()
		c.Ctx.ResponseWriter.WriteHeader(400)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		fmt.Println("exec signature " + err.Error())
		global.Logging("ERROR_PKM","images SIGNATURE  ---> " + err.Error())

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
			Message:        err.Error(),
			Statusid:       "3",
		}
		c.Data["json"] = response

		c.ServeJSON()

		return response
	}	

	querysign = `BEGIN
		INSERT INTO PKM_SIGNATURES VALUES (
			(SELECT TOP 1 OurBranchID FROM AO_GROUP WITH(NOLOCK) WHERE GROUPID = '`+dtpkm.Signature.GroupID+`'),
			'`+dtpkm.Signature.GroupID+`',
			'`+tranDate+`',
			'`+fileimages["TtdAccountOfficer"]+`',
			'`+fileimages["TtdKetuaKelompok"]+`',
			NULL,
			NULL,
			NEWID(),
			'`+dtpkm.Signature.Latitude+`',
			'`+dtpkm.Signature.Longitude+`'
		)
		END`

	_, err = tx.ExecContext(ctx, querysign)

	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		// tx.Rollback()
		c.Ctx.ResponseWriter.WriteHeader(400)
		statcod := c.Ctx.ResponseWriter.Status
		statusCode := strconv.Itoa(statcod)

		fmt.Println("exec signature " + err.Error())
		global.Logging("ERROR_PKM","Transaction SIGNATURE " + querysign + "  ---> " + err.Error())

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

		global.Logging("ERROR_PKM","Commit Transaction " + dtpkm.Signature.GroupID + "  ---> " + err.Error())

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