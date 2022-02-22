package message1_1

import (
	"bytes"
	"io"

	datatransfer "github.com/filecoin-project/go-data-transfer"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/node/bindnode"
)

//go:generate cbor-gen-for --map-encoding transferMessage1_1

// transferMessage1_1 is the transfer message for the 1.1 Data Transfer Protocol.
type transferMessage1_1 struct {
	IsRq bool

	Request  *transferRequest1_1
	Response *transferResponse1_1
}

// ========= datatransfer.Message interface

// IsRequest returns true if this message is a data request
func (tm *transferMessage1_1) IsRequest() bool {
	return tm.IsRq
}

// TransferID returns the TransferID of this message
func (tm *transferMessage1_1) TransferID() datatransfer.TransferID {
	if tm.IsRequest() {
		return tm.Request.TransferID()
	}
	return tm.Response.TransferID()
}

// ToNet serializes a transfer message type. It is simply a wrapper for MarshalCBOR, to provide
// symmetry with FromNet
func (tm *transferMessage1_1) ToIPLD() (datamodel.Node, error) {
	buf := new(bytes.Buffer)
	err := tm.ToNet(buf)
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

// ToNet serializes a transfer message type. It is simply a wrapper for MarshalCBOR, to provide
// symmetry with FromNet
func (tm *transferMessage1_1) ToNet(w io.Writer) error {
	node := bindnode.Wrap(&tm, Prototype.TransferMessage.Type())
	return dagcbor.Encode(node.Representation(), w)
}