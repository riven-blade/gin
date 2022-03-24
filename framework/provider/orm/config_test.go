package orm

import (
	"github.com/ddh-open/gin/framework/contract"
	"github.com/ddh-open/gin/framework/provider/config"
	"github.com/ddh-open/gin/framework/provider/env"
	tests "github.com/ddh-open/gin/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNiceConfig_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&env.NiceEnvProvider{})
	container.Bind(&config.NiceConfigProvider{})

	Convey("test config", t, func() {
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		config := &contract.DBConfig{}
		err := configService.Load("database.mysql", config)
		So(err, ShouldBeNil)
	})

	Convey("test mysql config", t, func() {
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		config := &contract.DBConfig{
			ConnMaxIdle: 10,
		}
		err := configService.Load("database.read", config)
		So(err, ShouldBeNil)
		So(config.ConnMaxIdle, ShouldEqual, 10)
	})

	Convey("test base config", t, func() {
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		config := &contract.DBConfig{
			ConnMaxOpen: 200,
		}
		err := configService.Load("database", config)
		So(err, ShouldBeNil)
		So(config.ConnMaxOpen, ShouldEqual, 200)
	})

}
