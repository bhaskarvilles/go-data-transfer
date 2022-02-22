package message1_1

import (
	"bytes"
	"io"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	xerrors "golang.org/x/xerrors"

	datatransfer "github.com/filecoin-project/go-data-transfer"
	"github.com/filecoin-project/go-data-transfer/encoding"
	"github.com/filecoin-project/go-data-transfer/message/types"
)

// NewRequest generates a new request for the data transfer protocol
func NewRequest(id datatransfer.TransferID, isRestart bool, isPull bool, vtype datatransfer.TypeIdentifier, voucher encoding.Encodable, baseCid cid.Cid, selector ipld.Node) (datatransfer.Request, error) {
	vbytes, err := encoding.Encode(voucher)
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	builder := basicnode.Prototype.Any.NewBuilder()
	err = dagcbor.Decode(builder, bytes.NewBuffer(vbytes))
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	vnode := builder.Build()

	if baseCid == cid.Undef {
		return nil, xerrors.Errorf("base CID must be defined")
	}
	selBytes, err := encoding.Encode(selector)
	if err != nil {
		return nil, xerrors.Errorf("Error encoding selector")
	}
	builder = basicnode.Prototype.Any.NewBuilder()
	err = dagcbor.Decode(builder, bytes.NewBuffer(selBytes))
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	selnode := builder.Build()

	var typ uint64
	if isRestart {
		typ = uint64(types.RestartMessage)
	} else {
		typ = uint64(types.NewMessage)
	}

	return &transferRequest1_1{
		Type:   typ,
		Pull:   isPull,
		Vouch:  &vnode,
		Stor:   &selnode,
		BCid:   &baseCid,
		VTyp:   vtype,
		XferID: uint64(id),
	}, nil
}

// RestartExistingChannelRequest creates a request to ask the other side to restart an existing channel
func RestartExistingChannelRequest(channelId datatransfer.ChannelID) datatransfer.Request {

	return &transferRequest1_1{Type: uint64(types.RestartExistingChannelRequestMessage),
		RestartChannel: channelId}
}

// CancelRequest request generates a request to cancel an in progress request
func CancelRequest(id datatransfer.TransferID) datatransfer.Request {
	return &transferRequest1_1{
		Type:   uint64(types.CancelMessage),
		XferID: uint64(id),
	}
}

// UpdateRequest generates a new request update
func UpdateRequest(id datatransfer.TransferID, isPaused bool) datatransfer.Request {
	return &transferRequest1_1{
		Type:   uint64(types.UpdateMessage),
		Paus:   isPaused,
		XferID: uint64(id),
	}
}

// VoucherRequest generates a new request for the data transfer protocol
func VoucherRequest(id datatransfer.TransferID, vtype datatransfer.TypeIdentifier, voucher encoding.Encodable) (datatransfer.Request, error) {
	vbytes, err := encoding.Encode(voucher)
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	builder := basicnode.Prototype.Any.NewBuilder()
	err = dagcbor.Decode(builder, bytes.NewBuffer(vbytes))
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	vnode := builder.Build()
	return &transferRequest1_1{
		Type:   uint64(types.VoucherMessage),
		Vouch:  &vnode,
		VTyp:   vtype,
		XferID: uint64(id),
	}, nil
}

// RestartResponse builds a new Data Transfer response
func RestartResponse(id datatransfer.TransferID, accepted bool, isPaused bool, voucherResultType datatransfer.TypeIdentifier, voucherResult encoding.Encodable) (datatransfer.Response, error) {
	vbytes, err := encoding.Encode(voucherResult)
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	builder := basicnode.Prototype.Any.NewBuilder()
	err = dagcbor.Decode(builder, bytes.NewBuffer(vbytes))
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	vnode := builder.Build()
	return &transferResponse1_1{
		Acpt:   accepted,
		Type:   uint64(types.RestartMessage),
		Paus:   isPaused,
		XferID: uint64(id),
		VTyp:   voucherResultType,
		VRes:   &vnode,
	}, nil
}

// NewResponse builds a new Data Transfer response
func NewResponse(id datatransfer.TransferID, accepted bool, isPaused bool, voucherResultType datatransfer.TypeIdentifier, voucherResult encoding.Encodable) (datatransfer.Response, error) {
	vbytes, err := encoding.Encode(voucherResult)
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	builder := basicnode.Prototype.Any.NewBuilder()
	err = dagcbor.Decode(builder, bytes.NewBuffer(vbytes))
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	vnode := builder.Build()
	return &transferResponse1_1{
		Acpt:   accepted,
		Type:   uint64(types.NewMessage),
		Paus:   isPaused,
		XferID: uint64(id),
		VTyp:   voucherResultType,
		VRes:   &vnode,
	}, nil
}

// VoucherResultResponse builds a new response for a voucher result
func VoucherResultResponse(id datatransfer.TransferID, accepted bool, isPaused bool, voucherResultType datatransfer.TypeIdentifier, voucherResult encoding.Encodable) (datatransfer.Response, error) {
	vbytes, err := encoding.Encode(voucherResult)
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	builder := basicnode.Prototype.Any.NewBuilder()
	err = dagcbor.Decode(builder, bytes.NewBuffer(vbytes))
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	vnode := builder.Build()
	return &transferResponse1_1{
		Acpt:   accepted,
		Type:   uint64(types.VoucherResultMessage),
		Paus:   isPaused,
		XferID: uint64(id),
		VTyp:   voucherResultType,
		VRes:   &vnode,
	}, nil
}

// UpdateResponse returns a new update response
func UpdateResponse(id datatransfer.TransferID, isPaused bool) datatransfer.Response {
	return &transferResponse1_1{
		Type:   uint64(types.UpdateMessage),
		Paus:   isPaused,
		XferID: uint64(id),
	}
}

// CancelResponse makes a new cancel response message
func CancelResponse(id datatransfer.TransferID) datatransfer.Response {
	return &transferResponse1_1{
		Type:   uint64(types.CancelMessage),
		XferID: uint64(id),
	}
}

// CompleteResponse returns a new complete response message
func CompleteResponse(id datatransfer.TransferID, isAccepted bool, isPaused bool, voucherResultType datatransfer.TypeIdentifier, voucherResult encoding.Encodable) (datatransfer.Response, error) {
	vbytes, err := encoding.Encode(voucherResult)
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	builder := basicnode.Prototype.Any.NewBuilder()
	err = dagcbor.Decode(builder, bytes.NewBuffer(vbytes))
	if err != nil {
		return nil, xerrors.Errorf("Creating request: %w", err)
	}
	vnode := builder.Build()
	return &transferResponse1_1{
		Type:   uint64(types.CompleteMessage),
		Acpt:   isAccepted,
		Paus:   isPaused,
		VTyp:   voucherResultType,
		VRes:   &vnode,
		XferID: uint64(id),
	}, nil
}

// FromNet can read a network stream to deserialize a GraphSyncMessage
func FromNet(r io.Reader) (datatransfer.Message, error) {
	builder := Prototype.TransferMessage.Representation().NewBuilder()
	err := dagcbor.Decode(builder, r)
	if err != nil {
		return nil, err
	}
	node := builder.Build()
	tresp := bindnode.Unwrap(node).(*transferMessage1_1)

	if (tresp.IsRequest() && tresp.Request == nil) || (!tresp.IsRequest() && tresp.Response == nil) {
		return nil, xerrors.Errorf("invalid/malformed message")
	}

	if tresp.IsRequest() {
		return tresp.Request, nil
	}
	return tresp.Response, nil
}

// FromNet can read a network stream to deserialize a GraphSyncMessage
func FromIPLD(nd datamodel.Node) (datatransfer.Message, error) {
	buf := new(bytes.Buffer)
	err := dagcbor.Encode(nd, buf)
	if err != nil {
		return nil, err
	}
	return FromNet(buf)
}