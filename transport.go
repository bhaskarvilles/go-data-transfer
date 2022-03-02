package datatransfer

import (
	"context"

	ipld "github.com/ipld/go-ipld-prime"
)

// EventsHandler are semantic data transfer events that happen as a result of graphsync hooks
type EventsHandler interface {
	// OnChannelOpened is called when we send a request for data to the other
	// peer on the given channel ID
	// return values are:
	// - error = ignore incoming data for this channel
	OnChannelOpened(chid ChannelID) error
	// OnResponseReceived is called when we receive a response to a request
	// - nil = continue receiving data
	// - error = cancel this request
	OnResponseReceived(chid ChannelID, msg Response) error

	// OnTransferQueued is called when a new data transfer request is queued in the transport layer.
	OnTransferQueued(chid ChannelID)

	// OnRequestReceived is called when we receive a new request to send data
	// for the given channel ID
	// return values are:
	// message = data transfer message along with reply
	// err = error
	// - nil = proceed with sending data
	// - error = cancel this request
	// - err == ErrPause - pause this request (only for new requests)
	// - err == ErrResume - resume this request (only for update requests)
	OnRequestReceived(chid ChannelID, msg Request) (Response, error)

	// OnTransferCompleted is called when we finish transferring data for the given channel ID
	// Error returns are logged but otherwise have no effect
	OnTransferCompleted(chid ChannelID, err error) error

	// OnRequestCancelled is called when a request we opened (with the given channel Id) to
	// receive data is cancelled by us.
	// Error returns are logged but otherwise have no effect
	OnRequestCancelled(chid ChannelID, err error) error

	// OnRequestDisconnected is called when a network error occurs trying to send a request
	OnRequestDisconnected(chid ChannelID, err error) error

	// OnSendDataError is called when a network error occurs sending data
	// at the transport layer
	OnSendDataError(chid ChannelID, err error) error

	// OnReceiveDataError is called when a network error occurs receiving data
	// at the transport layer
	OnReceiveDataError(chid ChannelID, err error) error

	// OnContextAugment allows the transport to attach data transfer tracing information
	// to its local context, in order to create a hierarchical trace
	OnContextAugment(chid ChannelID) func(context.Context) context.Context

	// OnDataReceive is called when we receive data for the given channel ID
	// return values are:
	// - nil = proceed with sending data
	// - error = cancel this request
	// - err == ErrPause - pause this request if pausable transport
	OnDataReceived(chid ChannelID, link ipld.Link, size uint64, index int64) error

	// OnDataQueued is called when data is queued for sending for the given channel ID
	// return values are:
	// message = data transfer message along with data
	// err = error
	// - nil = proceed with sending data
	// - error = cancel this request
	// - err == ErrPause - pause this request if pauseable transport
	OnDataQueued(chid ChannelID, link ipld.Link, size uint64, index int64) error

	// OnDataSent is called when we send data for the given channel ID
	OnDataSent(chid ChannelID, link ipld.Link, size uint64, index int64) error
}

/*
Transport defines the interface for a transport layer for data
transfer. Where the data transfer manager will coordinate setting up push and
pull requests, validation, etc, the transport layer is responsible for moving
data back and forth, and may be medium specific. For example, some transports
may have the ability to pause and resume requests, while others may not.
Some may support individual data events, while others may only support message
events. Some transport layers may opt to use the actual data transfer network
protocols directly while others may be able to encode messages in their own
data protocol.

Transport is the minimum interface that must be satisfied to serve as a datatransfer
transport layer. Transports must be able to open (open is always called by the receiving peer)
and close channels, and set at an event handler */
type Transport interface {
	// OpenChannel opens a channel on a given transport to move data back and forth
	OpenChannel(
		ctx context.Context,
		channel Channel,
		msg Message,
	) error

	// CloseChannel closes the given channel
	CloseChannel(ctx context.Context, chid ChannelID) error
	// SetEventHandler sets the handler for events on channels
	SetEventHandler(events EventsHandler) error
	// send message sends an arbitrary message over a transport
	SendMessage(ctx context.Context, chid ChannelID, message Message) error
	// CleanupChannel is called on the otherside of a cancel - removes any associated
	// data for the channel
	CleanupChannel(chid ChannelID)
}

// RestartableTransport is a transport that can be restarted in the case of network failure
// or node restart
type RestartableTransport interface {
	Transport
	// PauseChannel paused the given channel ID
	RestartChannel(ctx context.Context,
		channel ChannelState,
		msg Message,
	) error
}

// PauseableTransport is a transport that can also pause and resume channels
type PauseableTransport interface {
	Transport
	// PauseChannel paused the given channel ID
	PauseChannel(ctx context.Context,
		chid ChannelID,
	) error
	// ResumeChannel resumes the given channel
	ResumeChannel(ctx context.Context,
		msg Message,
		chid ChannelID,
	) error
}
