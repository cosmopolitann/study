package mvc

import (
	"github.com/cosmopolitann/clouddb/vo"
)

var TopicJoin *vo.TopicJoinMap

func init() {
	TopicJoin = new(vo.TopicJoinMap)
}