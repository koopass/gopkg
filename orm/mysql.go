package orm

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Entity interface {
	GetID() int64
	GetKV() map[string]interface{}
}

type EasyCRUD[T Entity] interface {
	SelectOne(ctx context.Context, query interface{}, args ...interface{}) (T, error)
	Select(ctx context.Context, query interface{}, args ...interface{}) ([]T, error)
	SelectLimit(ctx context.Context, limit int, query interface{}, args ...interface{}) ([]T, error)
	SelectPage(ctx context.Context, offset, limit int, query interface{}, args ...interface{}) ([]T, error)
	InsertOne(ctx context.Context, entity T, cols []string) (int64, error)
	UpdateOne(ctx context.Context, entity T, cols []string) (int64, error)
	DeleteOne(ctx context.Context, id int64) error
}

func NewEasyGORM[T Entity](db *gorm.DB) EasyCRUD[T] {
	return &EasyGORM[T]{
		db: db,
	}
}

type EasyGORM[T Entity] struct {
	db *gorm.DB
}

func (e *EasyGORM[T]) SelectLimit(ctx context.Context, limit int, query interface{}, args ...interface{}) ([]T, error) {
	var rows []T
	err := e.db.Where(query, args...).Limit(limit).Find(&rows).Error
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	return rows, nil
}

func (e *EasyGORM[T]) SelectPage(ctx context.Context, offset, limit int, query interface{}, args ...interface{}) ([]T, error) {
	var rows []T
	err := e.db.Where(query, args...).Offset(offset).Limit(limit).Find(&rows).Error
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	return rows, nil
}

func (e *EasyGORM[T]) SelectOne(ctx context.Context, query interface{}, args ...interface{}) (row T, err error) {
	var rows []T
	err = e.db.Where(query, args...).Limit(1).Find(&rows).Error
	if err != nil {
		return row, err
	}
	if len(rows) == 0 {
		return row, nil
	}
	return rows[0], nil
}

func (e *EasyGORM[T]) Select(ctx context.Context, query interface{}, args ...interface{}) ([]T, error) {
	var rows []T
	err := e.db.Where(query, args...).Find(&rows).Error
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, nil
	}
	return rows, nil
}

func (e *EasyGORM[T]) InsertOne(ctx context.Context, entity T, cols []string) (int64, error) {
	if len(cols) == 0 {
		err := e.db.Create(entity).Error
		return entity.GetID(), err
	}
	err := e.db.Select(cols).Create(entity).Error
	return entity.GetID(), err
}

func (e *EasyGORM[T]) UpdateOne(ctx context.Context, entity T, cols []string) (int64, error) {
	if len(cols) == 0 {
		return 0, nil
	}
	result := e.db.Model(entity).Select(cols).Updates(entity.GetKV())
	return result.RowsAffected, result.Error
}

func (e *EasyGORM[T]) DeleteOne(ctx context.Context, id int64) error {
	row := new(T)
	return e.db.Delete(row, id).Error
}

type MySQLConfig struct {
	DBName   string `json:"db_name"`
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Charset  string `json:"charset"`
	PoolSize int    `json:"pool_size"`
}

func (c *MySQLConfig) buildDSN() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.DBName, c.Charset)
}

func NewMySQL(c *MySQLConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(c.buildDSN()))
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(c.PoolSize)
	sqlDB.SetMaxOpenConns(c.PoolSize)
	return db, nil
}
