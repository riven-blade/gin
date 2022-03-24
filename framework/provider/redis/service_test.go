package redis

import (
	"context"
	"github.com/ddh-open/gin/framework/provider/config"
	"github.com/ddh-open/gin/framework/provider/env"
	"github.com/ddh-open/gin/framework/provider/log"
	tests "github.com/ddh-open/gin/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestNiceService_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&env.NiceEnvProvider{})
	container.Bind(&config.NiceConfigProvider{})
	container.Bind(&log.NiceLogServiceProvider{})

	Convey("test get client", t, func() {
		niceRedis, err := NewNiceRedis(container)
		So(err, ShouldBeNil)
		service, ok := niceRedis.(*NiceRedis)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("redis.write"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		ctx := context.Background()
		err = client.Set(ctx, "foo", "bar", 1*time.Hour).Err()
		So(err, ShouldBeNil)
		val, err := client.Get(ctx, "foo").Result()
		So(err, ShouldBeNil)
		So(val, ShouldEqual, "bar")
		err = client.Del(ctx, "foo").Err()
		So(err, ShouldBeNil)
	})
}
