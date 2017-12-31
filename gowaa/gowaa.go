package gowaa

/**
 * gowaa
 * ============================================================================
 * gowaa框架插件
 * ============================================================================
 * author: peter.wang
 * create time: 2016-6-4 20:41
 */

import (
	"fmt"
	"os"
	"strings"
)

var (
	gowaa_args map[string]string
)

/**
 * 解析参数
 * bgocfgdir	配制文件路径,相对于启动文件
 * bgocfgmode	dev/alpha/prod; dev表示加载app_dev配制文件
 * 传递格式如：./appexe bgocfgdir=.. bgocfgmode=dev
 */
func init() {
	gowaa_args = make(map[string]string)
	gowaa_args["bgocfgdir"] = ".."   // 默认值:..
	gowaa_args["bgocfgmode"] = "dev" // 默认值:dev

	for _, arg := range os.Args {
		if !strings.HasPrefix(strings.TrimSpace(arg), "-") {
			kv := strings.Split(arg, "=")
			if len(kv) == 2 {
				gowaa_args[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
			}
		}
	}

	fmt.Println(gowaa_args)
}

func GetArgsCDir() string {
	return gowaa_args["bgocfgdir"]
}
func GetArgsCMode() string {
	return gowaa_args["bgocfgmode"]
}
