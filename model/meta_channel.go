package model

import "simple-tracking/support/db"

type MetaArray map[string] interface{}

/**
  渠道
*/
type MetaChannel struct {
	ChannelName         string    `db:"channelName"`
}

func (this MetaChannel) value() db.DataMap {
	v := make(db.DataMap)
	v["channelName"] = this.ChannelName
	return v
}
