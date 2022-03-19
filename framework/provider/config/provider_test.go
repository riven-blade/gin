package config

import (
	"testing"

	"github.com/ddh-open/gin/framework"
	"github.com/ddh-open/gin/framework/contract"
	"github.com/ddh-open/gin/framework/provider/app"
	"github.com/ddh-open/gin/framework/provider/env"
	tests "github.com/ddh-open/gin/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNiceConfig_Normal(t *testing.T) {
	Convey("test nice config normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewNiceContainer()
		c.Bind(&app.NiceAppProvider{BaseFolder: basePath})
		c.Bind(&env.NiceEnvProvider{})

		err := c.Bind(&NiceConfigProvider{})
		So(err, ShouldBeNil)

		conf := c.MustMake(contract.ConfigKey).(contract.Config)
		So(conf.GetString("database.default.host"), ShouldEqual, "localhost")
		So(conf.GetInt("database.default.port"), ShouldEqual, 3306)
		//So(conf.GetFloat64("database.default.readtime"), ShouldEqual, 2.3)
		// So(conf.GetString("database.mysql.password"), ShouldEqual, "mypassword")

		maps := conf.GetStringMap("database.default")
		So(maps, ShouldContainKey, "host")
		So(maps["host"], ShouldEqual, "localhost")

		maps2 := conf.GetStringMapString("database.default")
		So(maps2["host"], ShouldEqual, "localhost")

		type Mysql struct {
			Host string `yaml:"host"`
		}
		ms := &Mysql{}
		err = conf.Load("database.default", ms)
		So(err, ShouldBeNil)
		So(ms.Host, ShouldEqual, "localhost")
	})
}
