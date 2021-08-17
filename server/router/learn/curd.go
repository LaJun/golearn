package learn
import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type learnRouter struct {
}

func (s *learnRouter) InitLearnRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("learn").Use(middleware.OperationRecord())
	var curdRouterApi = v1.ApiGroupApp.LearnApiCroup.LearnApi
	{
		apiRouter.POST("curd", curdRouterApi.CreateApi)               // 创建Api
	}
}