package main

import (
	//_ "github.com/SupenBysz/gf-admin-community"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "github.com/kysion/sms-library/example/internal/boot"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/kysion/sms-library/example/internal/boot"

	_ "github.com/kysion/sms-library/internal/logic"
)

func main() {
	boot.Main.Run(gctx.New())
}
