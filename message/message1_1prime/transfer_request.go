package message1_1

import (
	"bytes"
	"io"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/libp2p/go-libp2p-core/protocol"
	xerrors "golang.org/x/xerrors"

	datatransfer "github.com/filecoin-project/go-data-transfer"
	"github.com/filecoin-project/go-data-transfer/encoding"
	"github.com/filecoin-project/go-data-transfer/message/types"
)

//go:generate cbor-gen-for --map-encoding transferRequest1_1

// transferRequest1_1 is a struct for the 1.1 Data Transfer Protocol that fulfills the datatransfer.Request interface.
// its members are exported to be used by cbor-gen
type transferRequest1_1 struct {
	BCid   *cid.Cid
	Type   uint64
	Paus   bool
	Part   bool
	Pull   bool
	Stor   *datamodel.Node
	Vouch  *datamodel.Node
	VTyp   datatransfer.TypeIdentifier
	XferID uint64

	RestartChannel datatransfer.ChannelID
}

func (trq *transferRequest1_1) MessageForProtocol(targetProtocol protocol.ID) (datatransfer.Message, error) {
	switch targetProtocol {
	case datatransfer.ProtocolDataTransfer1_2:
		return trq, nil
	default:
		return nil, xerrors.Errorf("protocol not supported")
	}
}

// IsRequest always returns true in this case because this is a transfer request
func (trq *transferRequest1_1) IsRequest() bool {
	return true
}

func (trq *transferRequest1_1) IsRestart() bool {
	return trq.Type == uint64(types.RestartMessage)
}

func (trq *transferRequest1_1) IsRestartExistingChannelRequest() bool {
	return trq.Type == uint64(types.RestartExistingChannelRequestMessage)
}

func (trq *transferRequest1_1) RestartChannelId() (datatransfer.ChannelID, error) {
	if !trq.IsRestartExistingChannelRequest() {
		return datatransfer.ChannelID{}, xerrors.New("not a restart request")
	}
	return trq.RestartChannel, nil
}

func (trq *transferRequest1_1) IsNew() bool {
	return trq.Type == uint64(types.NewMessage)
}

func (trq *transferRequest1_1) IsUpdate() bool {
	return trq.Type == uint64(types.UpdateMessage)
}

func (trq *transferRequest1_1) IsVoucher() bool {
	return trq.Type == uint64(types.VoucherMessage) || trq.Type == uint64(types.NewMessage)
}

func (trq *transferRequest1_1) IsPaused() bool {
	return trq.Paus
}

func (trq *transferRequest1_1) TransferID() datatransfer.TransferID {
	return datatransfer.TransferID(trq.XferID)
}

// ========= datatransfer.Request interface
// IsPull returns true if this is a data pull request
func (trq *transferRequest1_1) IsPull() bool {
	return trq.Pull
}

// VoucherType returns the Voucher ID
func (trq *transferRequest1_1) VoucherType() datatransfer.TypeIdentifier {
	return trq.VTyp
}

// Voucher returns the Voucher bytes
func (trq *transferRequest1_1) Voucher(decoder encoding.Decoder) (encoding.Encodable, error) {
	if trq.Vouch == nil {
		return nil, xerrors.New("No voucher present to read")
	}
	buf := new(bytes.Buffer)
	dagcbor.Encode(*trq.Vouch, buf)
	return decoder.DecodeFromCbor(buf.Bytes())
}

func (trq *transferRequest1_1) EmptyVoucher() bool {
	return trq.VTyp == datatransfer.EmptyTypeIdentifier
}

// BaseCid returns the Base CID
func (trq *transferRequest1_1) BaseCid() cid.Cid {
	if trq.BCid == nil {
		return cid.Undef
	}
	return *trq.BCid
}

// Selector returns the message Selector bytes
func (trq *transferRequest1_1) Selector() (datamodel.Node, error) {
	if trq.Stor == nil {
		return nil, xerrors.New("No selector present to read")
	}
	return *trq.Stor, nil
}

// IsCancel returns true if this is a cancel request
func (trq *transferRequest1_1) IsCancel() bool {
	return trq.Type == uint64(types.CancelMessage)
}

// IsPartial returns true if this is a partial request
func (trq *transferRequest1_1) IsPartial() bool {
	return trq.Part
}

func (trq *transferRequest1_1) ToIPLD() (datamodel.Node, error) {
	buf := new(bytes.Buffer)
	err := trq.ToNet(buf)
	if err != nil {
		return nil, err
	}
	nb := basicnode.Prototype.Any.NewBuilder()
	err = dagcbor.Decode(nb, buf)
	if err != nil {
		return nil, err
	}
	return nb.Build(), nil
}

// ToNet serializes a transfer request. It's a wrapper for MarshalCBOR to provide
// symmetry with FromNet
func (trq *transferRequest1_1) ToNet(w io.Writer) error {
	msg := transferMessage1_1{
		IsRq:     true,
		Request:  trq,
		Response: nil,
	}
	node := bindnode.Wrap(&msg, Prototype.TransferMessage.Type())
	return dagcbor.Encode(node.Representation(), w)
}