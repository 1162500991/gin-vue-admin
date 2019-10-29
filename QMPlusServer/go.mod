module main

go 1.12

require (
	QMPlusCommon v0.0.0
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-gonic/gin v1.4.0
	github.com/jinzhu/gorm v1.9.11
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/pkg/errors v0.8.1
	github.com/qiniu/api.v7 v7.2.5+incompatible
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/satori/go.uuid v1.2.0
	github.com/shamsher31/goimgext v1.0.0 // indirect
	github.com/shamsher31/goimgtype v1.0.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.4.0
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.5.1
)

replace QMPlusCommon => ../QMPlusCommon
