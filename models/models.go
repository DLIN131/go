package models

// type USR_COMPANY struct {
// 	CPID int `json:"CPID" gorm:"column:ID;primaryKey"`
// 	CP_ACCOUNT string `json:"CP_ACCOUNT", gorm:"column:CP_ACCOUNT"`
// 	CP_PWD string `json:"CP_PWD", gorm:"column:CP_PWD"`
// 	DTID int `json:"DTID", gorm:"column:DTID"`
// 	STATUS int `json:"STATUS", gorm:"column:STATUS"`
// 	MOD_DATETIME string `json:"MOD_DATETIME", gorm:"column:MOD_DATETIME"`
// 	DATELINE_DATETIME string `json:"DATELINE_DATETIME", gorm:"column:DATELINE_DATETIME"`
// 	CREATE_DATETIME string `json:"CREATE_DATETIME", gorm:"column:CREATE_DATETIME"`
// 	ISPFI int `json:"ISPFI", gorm:"column:ISPFI"`
// 	IDOLD string `json:"IDOLD", gorm:"column:IDOLD"`
// 	BIND_AREA_CTYPE string `json:"BIND_AREA_CTYPE", gorm:"column:BIND_AREA_CTYPE"`
// 	TMP_PWD string `json:"TMP_PWD", gorm:"column:TMP_PWD"`
// 	TMP_TIME string `json:"TMP_TIME", gorm:"column:TMP_TIME"` 
// }

type USR_COMPANY struct {
	CPID int 
	CP_ACCOUNT string 
	CP_PWD string 
	DTID int 
	STATUS int 
	MOD_DATETIME string 
	DATELINE_DATETIME string 
	CREATE_DATETIME string 
	ISPFI int 
	IDOLD string 
	BIND_AREA_CTYPE string 
	TMP_PWD string 
	TMP_TIME string 
}

func (USR_COMPANY) TableName() string {
    return "USR_COMPANY" // ⚠️ 改成你在 SQL Server 實際的表名
}

type STREETLIGHT struct{
	SLID int `json:"SLID" gorm:"column:SLID;primaryKey"`
	CAREA string `json:"CAREA", gorm:"column:CAREA"`
	SLPWBOXID string `json:"SLPWBOXID", gorm:"column:SLPWBOXID"`
	CD_DIRECTION string `json:"CD_DIRECTION", gorm:"column:CD_DIRECTION"`
	CD_SETTYPE string `json:"CD_SETTYPE", gorm:"column:CD_SETTYPE"`
	CD_SLMATERIAL string `json:"CD_SLMATERIAL", gorm:"column:CD_SLMATERIAL"`
	CD_SLWATT string `json:"CD_SLWATT", gorm:"column:CD_SLWATT"`
	CD_STREET string `json:"CD_STREET", gorm:"column:CD_STREET"`
	CD_VILLAGE string `json:"CD_VILLAGE", gorm:"column:CD_VILLAGE"`
	SLSN string `json:"SLSN", gorm:"column:SLSN"`
	SLSNNO string `json:"SLSNNO", gorm:"column:SLSNNO"`
	SLADD string `json:"SLADD", gorm:"column:SLADD"`
	SLADDDUAN string `json:"SLADDDUAN", gorm:"column:SLADDDUAN"`
	SLADDLIN string `json:"SLADDLIN", gorm:"column:SLADDLIN"`
	SLADDALLEY string `json:"SLADDALLEY", gorm:"column:SLADDALLEY"`
	SLADDLANE string `json:"SLADDLANE", gorm:"column:SLADDLANE"`
	SLADDNO string `json:"SLADDNO", gorm:"column:SLADDNO"`
	SLDISTANCE string `json:"SLDISTANCE", gorm:"column:SLDISTANCE"`
	SLHEIGHT string `json:"SLHEIGHT", gorm:"column:SLHEIGHT"`
	SLPOLENO string `json:"SLPOLENO", gorm:"column:SLPOLENO"`
	COMPID string `json:"COMPID", gorm:"column:COMPID"`
	SLSDATE string `json:"SLSDATE", gorm:"column:SLSDATE"`
	SLLIFE string `json:"SLLIFE", gorm:"column:SLLIFE"`
	SLPDATE string `json:"SLPDATE", gorm:"column:SLPDATE"`
	SLBULB string `json:"SLBULB", gorm:"column:SLBULB"`
	SLMEMO string `json:"SLMEMO", gorm:"column:SLMEMO"`
	SLMATER string `json:"SLMATER", gorm:"column:SLMATER"`
	MTTIMES string `json:"MTTIMES", gorm:"column:MTTIMES"`
	SL_CUTOUT string `json:"SL_CUTOUT", gorm:"column:SL_CUTOUT"`
	SL_DM string `json:"SL_DM", gorm:"column:SL_DM"`
	SL_GROUND string `json:"SL_GROUND", gorm:"column:SL_GROUND"`
	SL_LEAKWATER string `json:"SL_LEAKWATER", gorm:"column:SL_LEAKWATER"`
	SL_AIR string `json:"SL_AIR", gorm:"column:SL_AIR"`
	SL_MONITOR string `json:"SL_MONITOR", gorm:"column:SL_MONITOR"`
	SL_LEAKPOWER string `json:"SL_LEAKPOWER", gorm:"column:SL_LEAKPOWER"`
	PRJNID string `json:"PRJNID", gorm:"column:PRJNID"`
	LMSKID string `json:"LMSKID", gorm:"column:LMSKID"`
	COVER string `json:"COVER", gorm:"column:COVER"`
	SL_STYLE string `json:"SL_STYLE", gorm:"column:SL_STYLE"`
	SL_STYLENO string `json:"SL_STYLENO", gorm:"column:SL_STYLENO"`
	DG_CORR string `json:"DG_CORR", gorm:"column:DG_CORR"`
	DG_LOSE string `json:"DG_LOSE", gorm:"column:DG_LOSE"`
	DG_CORD string `json:"DG_CORD", gorm:"column:DG_CORD"`
	DG_SCREEN string `json:"DG_SCREEN", gorm:"column:DG_SCREEN"`
	DG_LTUB string `json:"DG_LTUB", gorm:"column:DG_LTUB"`
	TWD67X string `json:"TWD67X", gorm:"column:TWD67X"`
	TWD67Y string `json:"TWD67Y", gorm:"column:TWD67Y"`
	TWD97X float64 `json:"TWD97X", gorm:"column:TWD97X"`
	TWD97Y float64 `json:"TWD97Y", gorm:"column:TWD97Y"`
	LONGITUDE string `json:"LONGITUDE", gorm:"column:LONGITUDE"`
	LATITUDE string `json:"LATITUDE", gorm:"column:LATITUDE"`
	STATE string `json:"STATE", gorm:"column:STATE"`
	FIXAREA string `json:"FIXAREA", gorm:"column:FIXAREA"`
	RECSTATE string `json:"RECSTATE", gorm:"column:RECSTATE"`
	FLAG string `json:"FLAG", gorm:"column:FLAG"`
	FLAGMEMO string `json:"FLAGMEMO", gorm:"column:FLAGMEMO"`
	DISDATE string `json:"DISDATE", gorm:"column:DISDATE"`
	SLFROM string `json:"SLFROM", gorm:"column:SLFROM"`
	OLDSLSN string `json:"OLDSLSN", gorm:"column:OLDSLSN"`
	HLAM_CATY string `json:"HLAM_CATY", gorm:"column:HLAM_CATY"`
	HLAM_POWER string `json:"HLAM_POWER", gorm:"column:HLAM_POWER"`
	HLAM_CNT string `json:"HLAM_CNT", gorm:"column:HLAM_CNT"`
	HLAM_SEC_DT string `json:"HLAM_SEC_DT", gorm:"column:HLAM_SEC_DT"`
	AD_LINE string `json:"AD_LINE", gorm:"column:AD_LINE"`
	NEWSLSN string `json:"NEWSLSN", gorm:"column:NEWSLSN"`
	SLTYPE string `json:"SLTYPE", gorm:"column:SLTYPE"`
	SLKIND string `json:"SLKIND", gorm:"column:SLKIND"`
	SLEDATE string `json:"SLEDATE", gorm:"column:SLEDATE"`
	LANDMARK string `json:"LANDMARK", gorm:"column:LANDMARK"`
	SET_MATERIAL string `json:"SET_MATERIAL", gorm:"column:SET_MATERIAL"`
	PRT_COMP string `json:"PRT_COMP", gorm:"column:PRT_COMP"`
	ISMAINSTREET string `json:"ISMAINSTREET", gorm:"column:ISMAINSTREET"`
	DENY_MAINTAIN string `json:"DENY_MAINTAIN", gorm:"column:DENY_MAINTAIN"`
	APPLY_FOR_ADD string `json:"APPLY_FOR_ADD", gorm:"column:APPLY_FOR_ADD"`
	DISDESC string `json:"DISDESC", gorm:"column:DISDESC"`
	DISUID string `json:"DISUID", gorm:"column:DISUID"`
	SLADDLONG string `json:"SLADDLONG", gorm:"column:SLADDLONG"`
	SLSNMEMO string `json:"SLSNMEMO", gorm:"column:SLSNMEMO"`
	USLSN string `json:"USLSN", gorm:"column:USLSN"`
	UDATE string `json:"UDATE", gorm:"column:UDATE"`
	IOT_STATUS string `json:"IOT_STATUS", gorm:"column:IOT_STATUS"`
	IOT_REPORT_ID string `json:"IOT_REPORT_ID", gorm:"column:IOT_REPORT_ID"`
	IOT_REPORT_AT string `json:"IOT_REPORT_AT", gorm:"column:IOT_REPORT_AT"`
	IS_ENERGY_SAVING string `json:"IS_ENERGY_SAVING", gorm:"column:IS_ENERGY_SAVING"`
	IS_AI_CONTROL string `json:"IS_AI_CONTROL", gorm:"column:IS_AI_CONTROL"`
	IS_SMART_POLE string `json:"IS_SMART_POLE", gorm:"column:IS_SMART_POLE"`
	EESW string `json:"EESW", gorm:"column:EESW"`
	DIS_OFFICIAL_DOCUMENT string `json:"DIS_OFFICIAL_DOCUMENT", gorm:"column:DIS_OFFICIAL_DOCUMENT"`
}

func (STREETLIGHT) TableName() string {
		return "STREETLIGHT" // ⚠️ 改成你在 SQL Server 實際的表名
}