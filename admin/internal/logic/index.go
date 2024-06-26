package logic

import (
	"github.com/issueye/pitaya_admin/internal/common/model"
	"github.com/issueye/pitaya_admin/internal/common/nodb"
	"github.com/issueye/pitaya_admin/internal/repository"
)

type Logic struct{}

// 查询 http 报文数据
func (l *Logic) GetHttpMessages(req *repository.HttpRequest) (datas []*model.TrafficStatistics, err error) {
	list, err := nodb.GetDB().FindByFilter(func(key string, value *model.TrafficStatistics) bool {
		return true
	}, 1, 10)

	return list, err
}
