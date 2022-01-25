package DB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"modTest/utlis/my_log"
)

// gorm的数据封装

// CreateDataDB
// 创建数据库连接
func CreateDataDB() (*gorm.DB, error) {
	var datetimePrecision = 2
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                                                                             // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,                                                                            // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision,                                                              // default datetime precision
		DontSupportRenameIndex:    true,                                                                            // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                            // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                           // smart configure based on used version
	}), &gorm.Config{})

	return db, err
}

// CreateDB
// 创建数据库函数
func CreateDB() (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open("root:hejin0929@tcp(127.0.0.1:3306)/data?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		my_log.WriteLog().Println(err)
		return nil, err
	}

	//db.SetConnMaxLifetime(time.Minute * 3)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)
	if err != nil {
		my_log.WriteLog().Println(err)
		return nil, err
	}

	return db, err
}
