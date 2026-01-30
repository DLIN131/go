package main

import (
	"net/http"
	"fmt"
	"log"
	"gin/models"
	"bytes"
	"strconv"
	"encoding/json"
	"io"
	"runtime"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"github.com/xuri/excelize/v2"
	// "github.com/iancoleman/orderedmap"
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
		var data []models.USR_COMPANY
		// var data []models.STREETLIGHT
		for i := 0; i < 10; i++ {
			if err := db.Find(&data).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}			
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	router.GET("/streetlightExport", func(c *gin.Context) {
		var streetlightData []models.STREETLIGHT
		if err := db.Select("CAREA", "SLSN", "SLADD", "CD_SETTYPE","CD_SLMATERIAL","CD_SLWATT", "SLHEIGHT", "SLSDATE", "TWD97X", "TWD97Y", "SLMEMO").Limit(200000).Find(&streetlightData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		f := excelize.NewFile()
		streamWriter, err := f.NewStreamWriter("Sheet1")
		if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
		}
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()

		// 設定標頭
    headers := []interface{}{"行政區", "路燈編號", "地址", "材質","燈泡材質","瓦數", "高度", "設立日期", "TWD97X", "TWD97Y", "SLMEMO"}
		cell, _ := excelize.CoordinatesToCellName(1, 1)
		streamWriter.SetRow(cell, headers)
    // for i, header := range headers {
    //     cell := fmt.Sprintf("%s1", string(rune(65+i))) // A1, B1, C1
    //     f.SetCellValue("Sheet1", cell, header)
    // }

    // 設定資料
    for rowIdx, row := range streetlightData {
				TWD97XStr := strconv.FormatFloat(row.TWD97X, 'f', 6, 64)
				TWD97YStr := strconv.FormatFloat(row.TWD97Y, 'f', 6, 64)
        values := []interface{}{row.CAREA, row.SLSN, row.SLADD, row.CD_SETTYPE,row.CD_SLMATERIAL,row.CD_SLWATT, row.SLHEIGHT, row.SLSDATE, TWD97XStr, TWD97YStr, row.SLMEMO}
				cell, _ := excelize.CoordinatesToCellName(1, rowIdx+2)
				if err := streamWriter.SetRow(cell, values); err != nil {
						fmt.Println("寫入錯誤：", err)
				}
        // for colIdx, val := range values {
        //     cell := fmt.Sprintf("%s%d", string(rune(65+colIdx)), rowIdx+2) // A2, B2, C2
        //     f.SetCellValue("Sheet1", cell, val)
        // }
    }

		// 結束串流寫入
		if err := streamWriter.Flush(); err != nil {
				fmt.Println("Flush error:", err)
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

	router.POST("/exportWithData", func(c *gin.Context) {
		fmt.Println(c.Request.Header.Get("Content-Type"))
		body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "無法讀取 body"})
        return
    }
    // var data []orderedmap.OrderedMap
    var data [][]interface{}
    if err := json.Unmarshal(body, &data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "JSON 解析失敗", "detail": err.Error()})
        return
    }
		fmt.Println(data)

		// 開始測量記憶體
    var mStart runtime.MemStats
    runtime.ReadMemStats(&mStart)

		f := excelize.NewFile()
		streamWriter, err := f.NewStreamWriter("Sheet1")
		if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
		}
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()

		// headersContentFlag := false
		// var headers []interface{}
		// for rowIdx, row := range data {
		// 	var values []interface{}
		// 	for _, key := range row.Keys() {
		// 			if headersContentFlag == false {
		// 				headers = append(headers, key)
		// 			}
		// 			val, _ := row.Get(key)
		// 			// fmt.Printf("%s: %v\n", key, val)
		// 			values = append(values, val)
		// 	}
		// 	if headersContentFlag == false {
		// 		cell, _ := excelize.CoordinatesToCellName(1, 1)
		// 		streamWriter.SetRow(cell, headers)
		// 	}
		// 	cell, _ := excelize.CoordinatesToCellName(1, rowIdx+2)
		// 	if err := streamWriter.SetRow(cell, values); err != nil {
		// 			fmt.Println("寫入錯誤：", err)
		// 	}
		// 	headersContentFlag = true
		// }

		for rowIdx, row := range data {
				// Excel 座標（從 rowIdx + 1 開始）
				cell, _ := excelize.CoordinatesToCellName(1, rowIdx+1)

				// 直接把整列寫入 Excel
				if err := streamWriter.SetRow(cell, row); err != nil {
						fmt.Println("寫入錯誤：", err)
				}
		}

		// 結束串流寫入
		if err := streamWriter.Flush(); err != nil {
				fmt.Println("Flush error:", err)
		}

    // 儲存檔案
    if err := f.SaveAs("C:\\ApacheServer\\WEBPHP8.3.10\\example-app\\storage\\app\\private\\go-excelize.xlsx"); err != nil {
        fmt.Println(err)
    }
		// 匯出後測量記憶體
    var mEnd runtime.MemStats
    runtime.ReadMemStats(&mEnd)

		
    memUsedMB := float64(mEnd.Alloc-mStart.Alloc) / 1024 / 1024
    memPeakMB := float64(mEnd.TotalAlloc-mStart.TotalAlloc) / 1024 / 1024
		c.JSON(http.StatusOK, gin.H{
			"msg": "JSON 解析成功", 
			"memory_used_mb": memUsedMB,
      "memory_peak_mb": memPeakMB,
		})


	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
