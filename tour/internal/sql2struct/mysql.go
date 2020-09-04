package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	)

type DBModel struct{
	DBEngine *sql.DB
	DBInfo *DBInfo
}

type DBInfo struct{
	DBType string
	Host string
	UserNmae string
	Password string
	Charset string
}

type TableColumn struct {
	ColumnName string
	DataType string
	IsNullable string
	ColumnKey string
	ColumnType string
	ColumnComment string
}

var DBTypeToStructType = map[string]string{
	"int": "int32",
	"tinyint": "int8",
	"smallint": "int",
	"mediumint": "int64",
	"bigint": "int64",
	"bit": "int",
	"bool": "bool",
	"enum": "string",
	"set": "string",
	"varchar": "string",
}

func NewDBModel(info *DBInfo) *DBModel{
	return &DBModel{
		DBInfo: info,
	}
}

func (m *DBModel) Connect() error {
	/**
	在连接MySQL数据库时使用的是标准库database/sql的Open方法，第一个参数为驱动名称（如mysql），
	第二个参数为驱动连接数据库的连接信息。需要注意的是，
	在程序中必须导入github.com/go-sql-driver/mysql进行MySQL驱动程序的初始化，否则会出现错误
	 */
	var err error
	dsn := fmt.Sprint(
		"%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local",
		m.DBInfo.UserNmae,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)

	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error){
	query := "SELECT " + "COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ?  AND TABLE_NAME = ? "

	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, errors.New("没有数据")
	}

	defer rows.Close()

	var columns []*TableColumn

	for rows.Next(){
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType,&column.ColumnKey, &column.IsNullable, &column.ColumnType,&column.ColumnComment)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}