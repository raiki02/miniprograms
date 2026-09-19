package dao

import (
	"MiniPrograms/conf"
	"MiniPrograms/responsity/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MiniProgramsDAO struct {
	db    *gorm.DB
	table string
}

// 初始化 SQLite 数据库连接
func InitDB(conf *conf.Config) (*gorm.DB, error) {
	// 连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: conf.Data.Dsn}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database111: %w", err)
	}
	err = db.Table("miniprograms").AutoMigrate(&model.MiniPrograms{})
	err = db.Table("change_miniprograms").AutoMigrate(&model.MiniPrograms{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// 创建一个新的 MiniPrograms 实例
func NewMiniProgramsDAO(db *gorm.DB, table string) *MiniProgramsDAO {
	return &MiniProgramsDAO{db: db, table: table}
}

// 添加table变量，两个接口操作两个表
func (d *MiniProgramsDAO) Find(name, version string) (*model.MiniPrograms, error) {
	var miniPrograms model.MiniPrograms
	gDB := d.db.Table(d.table).Model(&model.MiniPrograms{}).Where("name= ?", name)
	if version != "" {
		gDB = gDB.Where("version = ?", version)
	}
	err := gDB.First(&miniPrograms).Error
	if err != nil {
		return nil, err
	}
	return &miniPrograms, nil
}

func (d *MiniProgramsDAO) Save(miniPrograms model.MiniPrograms) error {
	gDB := d.db.Table(d.table).Model(&model.MiniPrograms{}).Where("name= ?", miniPrograms.Name)
	if miniPrograms.Version != "" {
		gDB = gDB.Where("version = ?", miniPrograms.Version)
	}

	return gDB.Save(&miniPrograms).Error
}
