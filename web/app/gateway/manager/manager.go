package manager

import (
	"go-web/web/app/gateway/conf"
)

// 特殊的service
// 对第三方平台封装的层，预处理返回结果及转化异常信息;
// 对Service层通用能力的下沉，如缓存方案、中间件通用处理;
// 与DAO层交互，对多个DAO的组合复用。
type Manager struct {
	c *conf.Config

}

func New() *Manager {
	return &Manager{

	}
}

func (m *Manager) WechatInfo()  {
	
}