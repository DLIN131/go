package main

import (
	"net/http"
	"fmt"
	"log"
	"gin/models"
	"bytes"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"github.com/xuri/excelize/v2"
)
// 資料庫連線變數
var db *gorm.DB



// 初始化資料庫
func initDB() {

	// ⚠️ 修改為你的連線字串
	server := "192.168.2.182"
	port := 1433
	user := "sa"
	password := "#10357#YuLin"
	database := "PWB_SLMMS1"

	dsn := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%d;database=%s;encrypt=disable",
		server, user, password, port, database,
	)

	var err error
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
            TablePrefix:   "",   // 無前綴
            // SingularTable: true, // 單數表名
            NoLowerCase:   true, // 禁用小寫轉換
        },
	})
	if err != nil {
		log.Fatal("❌ 連線 MSSQL 失敗：", err)
	}

	sqlDB, _ := db.DB()
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("❌ 資料庫 Ping 失敗：", err)
	}

	fmt.Println("✅ 成功連線 MSSQL！")
}


func main() {
	initDB()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{})

	router.GET("/mssqlSearchTest", func(c *gin.Context) {
		// name := c.Query("CP_ACCOUNT")
		var usr_companies []models.USR_COMPANY
		if err := db.Find(&usr_companies).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": usr_companies})
	})

	router.GET("/streetlightExport", func(c *gin.Context) {
		var streetlightData []models.STREETLIGHT
		if err := db.Select("CAREA", "SLSN", "SLADD", "CD_SETTYPE","CD_SLMATERIAL","CD_SLWATT", "SLHEIGHT", "SLSDATE", "TWD97X", "TWD97Y", "SLMEMO").Limit(100000).Find(&streetlightData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(streetlightData[0])
		f := excelize.NewFile()
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()

		// 設定標頭
    headers := []string{"行政區", "路燈編號", "地址", "材質","燈泡材質","瓦數", "高度", "設立日期", "TWD97X", "TWD97Y", "SLMEMO"}
    for i, header := range headers {
        cell := fmt.Sprintf("%s1", string(rune(65+i))) // A1, B1, C1
        f.SetCellValue("Sheet1", cell, header)
    }

    // 設定資料
    for rowIdx, row := range streetlightData {
				TWD97XStr := strconv.FormatFloat(row.TWD97X, 'f', 6, 64)
				TWD97YStr := strconv.FormatFloat(row.TWD97Y, 'f', 6, 64)
        values := []interface{}{row.CAREA, row.SLSN, row.SLADD, row.CD_SETTYPE,row.CD_SLMATERIAL,row.CD_SLWATT, row.SLHEIGHT, row.SLSDATE, TWD97XStr, TWD97YStr, row.SLMEMO}
        for colIdx, val := range values {
            cell := fmt.Sprintf("%s%d", string(rune(65+colIdx)), rowIdx+2) // A2, B2, C2
            f.SetCellValue("Sheet1", cell, val)
        }
    }

    // 儲存檔案
    if err := f.SaveAs("C:\\ApacheServer\\WEBPHP8.3.10\\SL_Main\\storage\\app\\public\\tmpdownload\\output.xlsx"); err != nil {
        fmt.Println(err)
    }

		// 將 Excel 檔案寫入緩衝區
		var buf bytes.Buffer
		if err := f.Write(&buf); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate Excel file"})
			return
		}
		// 設置 HTTP 標頭以觸發下載
		c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Header("Content-Disposition", "attachment; filename=streetlight_data.xlsx")
		c.Header("Content-Length", fmt.Sprintf("%d", buf.Len()))

    fmt.Println("Excel 檔案已匯出：output.xlsx")

		// 回傳二進位資料給前端
		c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())

		// c.JSON(http.StatusOK, gin.H{"msg": "Excel 檔案已匯出"})
	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
