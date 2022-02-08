package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["pkmapi/controllers:ImagesController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:ImagesController"],
        beego.ControllerComments{
            Method: "PkmImages",
            Router: "/home/PKM/:data",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:ImagesController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:ImagesController"],
        beego.ControllerComments{
            Method: "PkmImagesProd",
            Router: "/home/PKM_PROD/:data",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MainController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MainController"],
        beego.ControllerComments{
            Method: "AuthLogin",
            Router: "/AuthLogin",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MainController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MainController"],
        beego.ControllerComments{
            Method: "GetCollectionList",
            Router: "/GetCollectionList/:OurBranchID/:AO_ID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MainController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MainController"],
        beego.ControllerComments{
            Method: "GetCollectionListPAR",
            Router: "/GetCollectionListPAR/:OurBranchID/:AO_ID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MainController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MainController"],
        beego.ControllerComments{
            Method: "GetCollectionListPKMIndividual",
            Router: "/GetCollectionListPKMIndividual/:OurBranchID/:AO_ID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MainController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MainController"],
        beego.ControllerComments{
            Method: "GetNewDate",
            Router: "/GetDate",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MainController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MainController"],
        beego.ControllerComments{
            Method: "GroupList",
            Router: "/GetListGroup/:OurBranchID/:USERNAME",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MainController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MainController"],
        beego.ControllerComments{
            Method: "PostPKMIndividual",
            Router: "/PostPKMIndividual",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MainController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MainController"],
        beego.ControllerComments{
            Method: "PostTransaction",
            Router: "/PostTransaction",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MainController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MainController"],
        beego.ControllerComments{
            Method: "PostShit",
            Router: "/ReqShit",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MainController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MainController"],
        beego.ControllerComments{
            Method: "PostShitPAR",
            Router: "/ReqShitPAR",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MainController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MainController"],
        beego.ControllerComments{
            Method: "SetLogin",
            Router: "/SetLogin/:Username/:Password",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MainController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MainController"],
        beego.ControllerComments{
            Method: "Getup",
            Router: "/Shit/:OurBranchID/:AO_ID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"],
        beego.ControllerComments{
            Method: "Get_All_Wilayah",
            Router: "/GetAllWilayah",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"],
        beego.ControllerComments{
            Method: "GetListClient",
            Router: "/GetListClient/:OurBranchID/:CreatedBy/:Search/:PAGE/:LIMIT",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"],
        beego.ControllerComments{
            Method: "GetListClientBRNET",
            Router: "/GetListClientBRNET/:OurBranchID/:CreatedBy/:Search/:PAGE/:LIMIT",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"],
        beego.ControllerComments{
            Method: "Get_Master_Data",
            Router: "/GetMasterData/:OurBranchID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"],
        beego.ControllerComments{
            Method: "Get_Pencairan_Mobile",
            Router: "/GetPencairanMobile/:OurBranchID/:CreatedBy",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"],
        beego.ControllerComments{
            Method: "Get_Siklus_Nasabah_Brnet",
            Router: "/GetSiklusNasabahBrnet",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"],
        beego.ControllerComments{
            Method: "Get_Sosialisasi_Mobile",
            Router: "/GetSosialisasiMobile",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:MasterDataController"],
        beego.ControllerComments{
            Method: "ScoringStatus",
            Router: "/ScoringStatus/:prospek_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:OtherController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:OtherController"],
        beego.ControllerComments{
            Method: "AuthLogout",
            Router: "/AuthLogout",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:OtherController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:OtherController"],
        beego.ControllerComments{
            Method: "PostAuditTrail",
            Router: "/post_audit_trail",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmController"],
        beego.ControllerComments{
            Method: "AuthLogin",
            Router: "/AuthLogin",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmController"],
        beego.ControllerComments{
            Method: "GetCollectionList",
            Router: "/GetCollectionList/:OurBranchID/:AO_ID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmController"],
        beego.ControllerComments{
            Method: "GetCollectionListPAR",
            Router: "/GetCollectionListPAR/:OurBranchID/:AO_ID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmController"],
        beego.ControllerComments{
            Method: "GetCollectionListPKMIndividual",
            Router: "/GetCollectionListPKMIndividual/:OurBranchID/:AO_ID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmController"],
        beego.ControllerComments{
            Method: "GetNewDate",
            Router: "/GetDate",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmController"],
        beego.ControllerComments{
            Method: "GroupList",
            Router: "/GetListGroup/:OurBranchID/:USERNAME",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmController"],
        beego.ControllerComments{
            Method: "PostPKMIndividual",
            Router: "/PostPKMIndividual",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmController"],
        beego.ControllerComments{
            Method: "PostTransaction",
            Router: "/PostTransaction",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmController"],
        beego.ControllerComments{
            Method: "PostShit",
            Router: "/ReqShit",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmController"],
        beego.ControllerComments{
            Method: "PostShitPAR",
            Router: "/ReqShitPAR",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmController"],
        beego.ControllerComments{
            Method: "Getup",
            Router: "/Shit/:OurBranchID/:AO_ID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmbController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmbController"],
        beego.ControllerComments{
            Method: "GetPkmb",
            Router: "/get_pkmb/:Username",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PkmbController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PkmbController"],
        beego.ControllerComments{
            Method: "PostPkmb",
            Router: "/post_pkmb",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"],
        beego.ControllerComments{
            Method: "PersetujuanPembiayaan",
            Router: "/persetujuan_pembiayaan",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"],
        beego.ControllerComments{
            Method: "PostDataKelompok",
            Router: "/post_data_kelompok",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"],
        beego.ControllerComments{
            Method: "PostKetuaSubKetua",
            Router: "/post_ketua_subketua",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"],
        beego.ControllerComments{
            Method: "PostPencairan",
            Router: "/post_pencairan",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"],
        beego.ControllerComments{
            Method: "PostPP",
            Router: "/post_pp",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"],
        beego.ControllerComments{
            Method: "PostProspek",
            Router: "/post_prospek",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"],
        beego.ControllerComments{
            Method: "PostProspekLama",
            Router: "/post_prospek_lama",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"],
        beego.ControllerComments{
            Method: "PostProspekUK",
            Router: "/post_prospek_uk",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"] = append(beego.GlobalControllerRouter["pkmapi/controllers:PostInisiasiController"],
        beego.ControllerComments{
            Method: "PostVerifStatus",
            Router: "/post_verif_status",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
