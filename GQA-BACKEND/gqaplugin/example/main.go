package example

import (
	"gin-quasar-admin/gqaplugin/example/data"
	"gin-quasar-admin/gqaplugin/example/model"
	"gin-quasar-admin/gqaplugin/example/router/private_router"
	"gin-quasar-admin/gqaplugin/example/router/public_router"
	"github.com/gin-gonic/gin"
)

var PluginExample = new(example)

type example struct{}

func (*example) PluginCode() string {
	return "example"
}

func (*example) PluginName() string {
	return "哈哈哈"
}

func (p *example) PluginRouterPublic(publicGroup *gin.RouterGroup) {
	//实现接口方法，公开路由初始化
	public_router.InitPublicRouter(publicGroup)
}

func (p *example) PluginRouterPrivate(privateGroup *gin.RouterGroup) {
	//实现接口方法，鉴权路由初始化
	private_router.InitPrivateRouter(privateGroup)
}

func (p *example) PluginMigrate() []interface{} {
	//实现接口方法，数据表初始化，填入插件中的数据库模型
	var ModelList = []interface{}{
		model.GqaPluginExampleNews{},
	}
	return ModelList
}

func (p *example) PluginData() []interface{ LoadData() (err error) } {
	//实现接口方法，初始化数据，填入插件中的初始化数据
	var DataList = []interface{ LoadData() (err error) }{
		data.PluginExampleNews,
	}
	return DataList
}