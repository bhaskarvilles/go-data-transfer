// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package internal

import (
	"fmt"
	"io"
	"sort"

	datatransfer "github.com/filecoin-project/go-data-transfer"
	cid "github.com/ipfs/go-cid"
	peer "github.com/libp2p/go-libp2p-core/peer"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = sort.Sort

func (t *ChannelState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{180}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.SelfPeer (peer.ID) (string)
	if len("SelfPeer") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"SelfPeer\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("SelfPeer"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("SelfPeer")); err != nil {
		return err
	}

	if len(t.SelfPeer) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.SelfPeer was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.SelfPeer))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.SelfPeer)); err != nil {
		return err
	}

	// t.TransferID (datatransfer.TransferID) (uint64)
	if len("TransferID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TransferID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("TransferID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("TransferID")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TransferID)); err != nil {
		return err
	}

	// t.Initiator (peer.ID) (string)
	if len("Initiator") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Initiator\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Initiator"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Initiator")); err != nil {
		return err
	}

	if len(t.Initiator) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Initiator was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Initiator))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Initiator)); err != nil {
		return err
	}

	// t.Responder (peer.ID) (string)
	if len("Responder") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Responder\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Responder"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Responder")); err != nil {
		return err
	}

	if len(t.Responder) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Responder was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Responder))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Responder)); err != nil {
		return err
	}

	// t.BaseCid (cid.Cid) (struct)
	if len("BaseCid") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"BaseCid\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("BaseCid"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("BaseCid")); err != nil {
		return err
	}

	if err := cbg.WriteCidBuf(scratch, w, t.BaseCid); err != nil {
		return xerrors.Errorf("failed to write cid field t.BaseCid: %w", err)
	}

	// t.Selector (typegen.Deferred) (struct)
	if len("Selector") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Selector\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Selector"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Selector")); err != nil {
		return err
	}

	if err := t.Selector.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Sender (peer.ID) (string)
	if len("Sender") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Sender\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Sender"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Sender")); err != nil {
		return err
	}

	if len(t.Sender) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Sender was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Sender))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Sender)); err != nil {
		return err
	}

	// t.Recipient (peer.ID) (string)
	if len("Recipient") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Recipient\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Recipient"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Recipient")); err != nil {
		return err
	}

	if len(t.Recipient) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Recipient was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Recipient))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Recipient)); err != nil {
		return err
	}

	// t.TotalSize (uint64) (uint64)
	if len("TotalSize") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TotalSize\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("TotalSize"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("TotalSize")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TotalSize)); err != nil {
		return err
	}

	// t.Status (datatransfer.Status) (uint64)
	if len("Status") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Status\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Status"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Status")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Status)); err != nil {
		return err
	}

	// t.Queued (uint64) (uint64)
	if len("Queued") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Queued\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Queued"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Queued")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Queued)); err != nil {
		return err
	}

	// t.Sent (uint64) (uint64)
	if len("Sent") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Sent\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Sent"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Sent")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Sent)); err != nil {
		return err
	}

	// t.Received (uint64) (uint64)
	if len("Received") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Received\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Received"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Received")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Received)); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len("Message") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Message\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Message"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Message")); err != nil {
		return err
	}

	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Message)); err != nil {
		return err
	}

	// t.Vouchers ([]internal.EncodedVoucher) (slice)
	if len("Vouchers") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Vouchers\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Vouchers"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Vouchers")); err != nil {
		return err
	}

	if len(t.Vouchers) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Vouchers was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Vouchers))); err != nil {
		return err
	}
	for _, v := range t.Vouchers {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.VoucherResults ([]internal.EncodedVoucherResult) (slice)
	if len("VoucherResults") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"VoucherResults\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("VoucherResults"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("VoucherResults")); err != nil {
		return err
	}

	if len(t.VoucherResults) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.VoucherResults was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.VoucherResults))); err != nil {
		return err
	}
	for _, v := range t.VoucherResults {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.ReceivedBlocksTotal (int64) (int64)
	if len("ReceivedBlocksTotal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"ReceivedBlocksTotal\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("ReceivedBlocksTotal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("ReceivedBlocksTotal")); err != nil {
		return err
	}

	if t.ReceivedBlocksTotal >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.ReceivedBlocksTotal)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.ReceivedBlocksTotal-1)); err != nil {
			return err
		}
	}

	// t.QueuedBlocksTotal (int64) (int64)
	if len("QueuedBlocksTotal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"QueuedBlocksTotal\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("QueuedBlocksTotal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("QueuedBlocksTotal")); err != nil {
		return err
	}

	if t.QueuedBlocksTotal >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.QueuedBlocksTotal)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.QueuedBlocksTotal-1)); err != nil {
			return err
		}
	}

	// t.SentBlocksTotal (int64) (int64)
	if len("SentBlocksTotal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"SentBlocksTotal\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("SentBlocksTotal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("SentBlocksTotal")); err != nil {
		return err
	}

	if t.SentBlocksTotal >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.SentBlocksTotal)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.SentBlocksTotal-1)); err != nil {
			return err
		}
	}

	// t.Stages (datatransfer.ChannelStages) (struct)
	if len("Stages") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Stages\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Stages"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Stages")); err != nil {
		return err
	}

	if err := t.Stages.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *ChannelState) UnmarshalCBOR(r io.Reader) error {
	*t = ChannelState{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ChannelState: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.SelfPeer (peer.ID) (string)
		case "SelfPeer":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.SelfPeer = peer.ID(sval)
			}
			// t.TransferID (datatransfer.TransferID) (uint64)
		case "TransferID":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.TransferID = datatransfer.TransferID(extra)

			}
			// t.Initiator (peer.ID) (string)
		case "Initiator":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Initiator = peer.ID(sval)
			}
			// t.Responder (peer.ID) (string)
		case "Responder":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Responder = peer.ID(sval)
			}
			// t.BaseCid (cid.Cid) (struct)
		case "BaseCid":

			{

				c, err := cbg.ReadCid(br)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.BaseCid: %w", err)
				}

				t.BaseCid = c

			}
			// t.Selector (typegen.Deferred) (struct)
		case "Selector":

			{

				t.Selector = new(cbg.Deferred)

				if err := t.Selector.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("failed to read deferred field: %w", err)
				}
			}
			// t.Sender (peer.ID) (string)
		case "Sender":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Sender = peer.ID(sval)
			}
			// t.Recipient (peer.ID) (string)
		case "Recipient":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Recipient = peer.ID(sval)
			}
			// t.TotalSize (uint64) (uint64)
		case "TotalSize":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.TotalSize = uint64(extra)

			}
			// t.Status (datatransfer.Status) (uint64)
		case "Status":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Status = datatransfer.Status(extra)

			}
			// t.Queued (uint64) (uint64)
		case "Queued":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Queued = uint64(extra)

			}
			// t.Sent (uint64) (uint64)
		case "Sent":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Sent = uint64(extra)

			}
			// t.Received (uint64) (uint64)
		case "Received":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Received = uint64(extra)

			}
			// t.Message (string) (string)
		case "Message":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Message = string(sval)
			}
			// t.Vouchers ([]internal.EncodedVoucher) (slice)
		case "Vouchers":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.Vouchers: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.Vouchers = make([]EncodedVoucher, extra)
			}

			for i := 0; i < int(extra); i++ {

				var v EncodedVoucher
				if err := v.UnmarshalCBOR(br); err != nil {
					return err
				}

				t.Vouchers[i] = v
			}

			// t.VoucherResults ([]internal.EncodedVoucherResult) (slice)
		case "VoucherResults":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.VoucherResults: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.VoucherResults = make([]EncodedVoucherResult, extra)
			}

			for i := 0; i < int(extra); i++ {

				var v EncodedVoucherResult
				if err := v.UnmarshalCBOR(br); err != nil {
					return err
				}

				t.VoucherResults[i] = v
			}

			// t.ReceivedBlocksTotal (int64) (int64)
		case "ReceivedBlocksTotal":
			{
				maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative oveflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.ReceivedBlocksTotal = int64(extraI)
			}
			// t.QueuedBlocksTotal (int64) (int64)
		case "QueuedBlocksTotal":
			{
				maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative oveflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.QueuedBlocksTotal = int64(extraI)
			}
			// t.SentBlocksTotal (int64) (int64)
		case "SentBlocksTotal":
			{
				maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
				var extraI int64
				if err != nil {
					return err
				}
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative oveflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.SentBlocksTotal = int64(extraI)
			}
			// t.Stages (datatransfer.ChannelStages) (struct)
		case "Stages":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					t.Stages = new(datatransfer.ChannelStages)
					if err := t.Stages.UnmarshalCBOR(br); err != nil {
						return xerrors.Errorf("unmarshaling t.Stages pointer: %w", err)
					}
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *EncodedVoucher) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Type (datatransfer.TypeIdentifier) (string)
	if len("Type") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Type\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Type"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Type")); err != nil {
		return err
	}

	if len(t.Type) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Type was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Type))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Type)); err != nil {
		return err
	}

	// t.Voucher (typegen.Deferred) (struct)
	if len("Voucher") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Voucher\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Voucher"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Voucher")); err != nil {
		return err
	}

	if err := t.Voucher.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *EncodedVoucher) UnmarshalCBOR(r io.Reader) error {
	*t = EncodedVoucher{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("EncodedVoucher: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Type (datatransfer.TypeIdentifier) (string)
		case "Type":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Type = datatransfer.TypeIdentifier(sval)
			}
			// t.Voucher (typegen.Deferred) (struct)
		case "Voucher":

			{

				t.Voucher = new(cbg.Deferred)

				if err := t.Voucher.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("failed to read deferred field: %w", err)
				}
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *EncodedVoucherResult) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Type (datatransfer.TypeIdentifier) (string)
	if len("Type") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Type\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Type"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Type")); err != nil {
		return err
	}

	if len(t.Type) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Type was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Type))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Type)); err != nil {
		return err
	}

	// t.VoucherResult (typegen.Deferred) (struct)
	if len("VoucherResult") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"VoucherResult\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("VoucherResult"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("VoucherResult")); err != nil {
		return err
	}

	if err := t.VoucherResult.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *EncodedVoucherResult) UnmarshalCBOR(r io.Reader) error {
	*t = EncodedVoucherResult{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("EncodedVoucherResult: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Type (datatransfer.TypeIdentifier) (string)
		case "Type":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Type = datatransfer.TypeIdentifier(sval)
			}
			// t.VoucherResult (typegen.Deferred) (struct)
		case "VoucherResult":

			{

				t.VoucherResult = new(cbg.Deferred)

				if err := t.VoucherResult.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("failed to read deferred field: %w", err)
				}
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
