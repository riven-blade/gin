package id

import (
	"github.com/ddh-open/gin/framework/provider/env"
	tests "github.com/ddh-open/gin/test"
	"testing"

	"github.com/ddh-open/gin/framework/contract"
	"github.com/ddh-open/gin/framework/provider/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConsoleLog_Normal(t *testing.T) {
	Convey("test nice console log normal case", t, func() {
		c := tests.InitBaseContainer()
		c.Bind(&env.NiceEnvProvider{})
		c.Bind(&config.NiceConfigProvider{})

		err := c.Bind(&NiceIDProvider{})
		So(err, ShouldBeNil)

		idService := c.MustMake(contract.IDKey).(contract.IDService)
		xid := idService.NewID()
		So(xid, ShouldNotBeEmpty)
	})
}
