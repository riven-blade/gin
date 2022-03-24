package config

import (
	"github.com/ddh-open/gin/framework/provider/env"
	"path/filepath"
	"testing"

	"github.com/ddh-open/gin/framework/contract"
	tests "github.com/ddh-open/gin/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNiceConfig_GetInt(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&env.NiceEnvProvider{})
	Convey("test nice env normal case", t, func() {
		appService := container.MustMake(contract.AppKey).(contract.App)
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		folder := filepath.Join(appService.ConfigFolder(), envService.AppEnv())

		serv, err := NewNiceConfig(container, folder, map[string]string{})
		So(err, ShouldBeNil)
		conf := serv.(*NiceConfig)
		timeout := conf.GetString("database.mysql.timeout")
		So(timeout, ShouldEqual, "10s")
	})
}
