package learn

import "gin-vue-admin/service"

type ApiGroup struct {
	LearnApi
}

var learnService = service.ServiceGroupApp.LearnServiceGroup.CurdService