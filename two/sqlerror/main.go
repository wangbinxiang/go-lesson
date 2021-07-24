package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type Row struct {
	Id   string
	Name string
}

var ErrNotFound = errors.New("no found")

func query(row *Row) error {
	return sql.ErrNoRows
}

func dao() (Row, error) {
	var row Row

	err := query(&row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// dao层不处理error
			return row, errors.Wrapf(err, "sql:%s error: %v", "select * from user where id = 1", err)
		} else {
			return row, err
		}
	}
	return row, nil
}

func main() {
	row, err := dao()
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			//
			fmt.Println("未找到数据，返回空数据给上层")
		} else {
			fmt.Println("其他err，返回err给上层")
		}
	}
	fmt.Println("query row: ", row)

}
