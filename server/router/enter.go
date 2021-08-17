package router

import (
	"gin-vue-admin/router/autocode"
	"gin-vue-admin/router/example"
	"gin-vue-admin/router/learn"
	"gin-vue-admin/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
	Learn    learn.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
