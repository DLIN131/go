package main

import (
	"github.com/gin-gonic/gin"
  "net/http"
  "database/sql"
  _ "github.com/denisenkom/go-mssqldb" // 導入 MSSQL 驅動程式
  "log"
  "strconv"
  "strings"
)

func main(){
	router := gin.Default()
  connString := "server=192.168.2.182;user id=NTPSLG;password=2B39o182;database=PWB_SLMMS1"
  // 建立連線
  db, err := sql.Open("sqlserver", connString)
  if err != nil {
      log.Fatal("連線失敗: ", err)
  }
  defer db.Close()

  router.GET("/ping", func(c *gin.Context) {
    query := `SELECT TOP (1000) [LID]
                                ,[SLID]
                                ,[LPSN]
                                ,[SENSERID]
                                ,[MATERIAL]
                                ,[WATT]
                                ,[IS_HLAM]
                                ,[QUEUE]
                                ,[RECSTATE]
                                ,[CREATETIME]
                                ,[UPDATETIME]
                            FROM [PWB_SLMMS1].[dbo].[LAMP]
                            where SLID = '16464'`
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal("查詢失敗: ", err)
    }
    defer rows.Close()

   // 儲存查詢結果
    var results []map[string]interface{}
    for rows.Next() {
        var lid int
        var slid, lpsn, senserid, material, watt, isHlam, queue, recstate *string
        var createTime, updateTime sql.NullTime

        // 掃描所有欄位
        if err := rows.Scan(&lid, &slid, &lpsn, &senserid, &material, &watt, &isHlam, &queue, &recstate, &createTime, &updateTime); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "掃描結果失敗: " + err.Error()})
            return
        }

        // 處理時間欄位
        createTimeStr := ""
        if createTime.Valid {
            createTimeStr = createTime.Time.Format("2006-01-02 15:04:05")
        }
        updateTimeStr := ""
        if updateTime.Valid {
            updateTimeStr = updateTime.Time.Format("2006-01-02 15:04:05")
        }

        // 字串拼接（使用 strings.Builder 提高效率）
        var builder strings.Builder
        builder.WriteString(strconv.Itoa(lid))
        builder.WriteString("-")
        builder.WriteString(*lpsn)
        message := builder.String()

        // 儲存結果
        results = append(results, map[string]interface{}{
            "lid":         lid,
            "slid":        slid,
            "lpsn":        lpsn,
            "senserid":    senserid,
            "material":    material,
            "watt":        watt,
            "isHlam":      isHlam,
            "queue":       queue,
            "recstate":    recstate,
            "createTime":  createTimeStr,
            "updateTime":  updateTimeStr,
            "message":     message,
        })
    }

    // 檢查查詢錯誤
    if err = rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "查詢處理錯誤: " + err.Error()})
        return
    }


    // 返回結果
    c.JSON(http.StatusOK, gin.H{
        "data":         results,
    })
  })

  router.Run(":8080") // listen and serve on 0.0.0.0:8080
}