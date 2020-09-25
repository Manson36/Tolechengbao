package datasource

import "github.com/lechengbao/datamodels"

func init() {
	Rds.initDB()
	PqDB.initDB()

	//自动创建数据库表
	_ = PqDB.AutoMigrate(datamodels.GetModelList()...)	//这里因为没有加 ... 导致了数据表创建错误
}
