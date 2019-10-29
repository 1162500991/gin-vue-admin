package dbModel

import (
	"QMPlusCommon/init/qmsql"
	"github.com/jinzhu/gorm"
)

type ExampleModule struct {
	gorm.Model
	ExampleName string `json:"exampleName"`
}

func (e *ExampleModule) CreateExampleModule() (err error) {
	err = qmsql.DEFAULTDB.Create(e).Error
	return
}
