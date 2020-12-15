package main

import (
	"AutoModel/global"
	"AutoModel/initialize"
	"AutoModel/core"
	"os"
)

func main()  {
	global.GVA_VP = core.Viper()          // 初始化Viper
	global.GVA_LOG = core.Zap()           // 初始化zap日志库
	global.GVA_DB = initialize.Gorm()
	targetTable:=os.Args[1]
	global.GVA_LOG.Info("table:"+targetTable)
	db, _ := global.GVA_DB.DB()
	defer db.Close()
	buildTableInfo()

}

func buildTableInfo() {


}