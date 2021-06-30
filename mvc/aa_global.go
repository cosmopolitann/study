package mvc

import (
	"github.com/cosmopolitann/clouddb/vo"
)

var TopicJoin *vo.TopicJoinMap

func init() {
	TopicJoin = vo.NewTopicJoin()
}

const (
	// RemoteIpnsAddr  = "k51qzi5uqu5dl2hdjuvu5mqlxuvezwe5wbedi6uh7dgu1uiv61vh4p4b71b17v"
	// RemoteIpnsAddr = "k51qzi5uqu5dlkjyn9btb65suntsm74kjj5cqnroad8z380sgv9k1dchu2rcdv"
	RemoteIpnsAddr = "k51qzi5uqu5dgjh05fu67ayu878ejp15okywlh2egemyp23r82ctk6dlwilgvk"
)
