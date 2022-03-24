package config

import (
	tests "github.com/ddh-open/gin/test"
	"testing"

	"github.com/ddh-open/gin/framework/contract"
	"github.com/ddh-open/gin/framework/provider/env"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNiceConfig_Normal(t *testing.T) {
	Convey("test nice config normal case", t, func() {
		c := tests.InitBaseContainer()
		c.Bind(&env.NiceEnvProvider{})
		err := c.Bind(&NiceConfigProvider{})
		So(err, ShouldBeNil)

		conf := c.MustMake(contract.ConfigKey).(contract.Config)
		So(conf.GetString("database.mysql.host"), ShouldEqual, "localhost")
		So(conf.GetInt("database.mysql.port"), ShouldEqual, 3306)
		//So(conf.GetFloat64("database.mysql.readtime"), ShouldEqual, 2.3)
		// So(conf.GetString("database.mysql.password"), ShouldEqual, "mypassword")

		maps := conf.GetStringMap("database.mysql")
		So(maps, ShouldContainKey, "host")
		So(maps["host"], ShouldEqual, "localhost")

		maps2 := conf.GetStringMapString("database.mysql")
		So(maps2["host"], ShouldEqual, "localhost")

		type Mysql struct {
			Host string `yaml:"host"`
		}
		ms := &Mysql{}
		err = conf.Load("database.mysql", ms)
		So(err, ShouldBeNil)
		So(ms.Host, ShouldEqual, "localhost")
	})
}
