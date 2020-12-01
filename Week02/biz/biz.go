package biz

import (
	"main/dao"
)

func Biz() error {
	data, err := dao.FetchData()
	if err != nil {
		return err
	}
	// process data
	_ = data

	return nil
}
