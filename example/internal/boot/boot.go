package boot

import (
	"context"
	//"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/kysion/sms-library/example/router"
	"github.com/kysion/sms-library/sms_global"

	_ "github.com/kysion/sms-library/example/internal/boot/internal"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 导入权限树 别的项目进行导入
			//sys_service.SysPermission().ImportPermissionTree(ctx, sms_consts.PermissionTree, nil)

			// 初始化路由
			apiPrefix := g.Cfg().MustGet(ctx, "service.apiPrefix").String()
			s.Group(apiPrefix, func(group *ghttp.RouterGroup) {
				// 中间件注册
				group.Middleware(ghttp.MiddlewareHandlerResponse)

				// 匿名路由绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					// group.Group("sms", func(group *ghttp.RouterGroup) { group.Bind(sms_controller.Sms) })

					// 路由绑定
					router.ModulesGroup(sms_global.Global.Modules, group)

				})

			})

			s.Run()
			return nil
		},
	}
)
