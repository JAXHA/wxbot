package sqlite

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Orm *gorm.DB
}

// Open 创建数据库连接
func Open(dbPath string, db *DB, opts ...gorm.Option) error {
	d, err := gorm.Open(sqlite.Open(dbPath), opts...)
	if err != nil {
		return err
	}
	db.Orm = d
	return nil
}

// Create 创建数据表
func (d *DB) Create(table string, dst ...interface{}) error {
	return d.Orm.Table(table).AutoMigrate(dst...)
}

// CreateAndFirstOrCreate 创建数据表并创建第一条数据，如果已有一条数据则不创建
func (d *DB) CreateAndFirstOrCreate(table string, dest interface{}, conds ...interface{}) error {
	if err := d.Orm.Table(table).AutoMigrate(dest); err != nil {
		return err
	}
	return d.Orm.Table(table).FirstOrCreate(dest, conds...).Error
}
