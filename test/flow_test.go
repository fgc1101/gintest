package test

import (
	"fmt"
	"gintest/common"
	"testing"
)

func TestFlow(t *testing.T) {
	_config := common.GetConfig()

	fmt.Println(_config.DataSource)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		_config.DataSource.Username,
		_config.DataSource.Password,
		_config.DataSource.Host,
		_config.DataSource.Port,
		_config.DataSource.Database,
		_config.DataSource.Charset)

	fmt.Printf(dsn)
}
