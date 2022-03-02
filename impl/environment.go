package impl

import (
	datatransfer "github.com/filecoin-project/go-data-transfer"
)

type channelEnvironment struct {
	m *manager
}

func (ce *channelEnvironment) CleanupChannel(chid datatransfer.ChannelID) {
	ce.m.transport.CleanupChannel(chid)
	ce.m.spansIndex.EndChannelSpan(chid)
}
