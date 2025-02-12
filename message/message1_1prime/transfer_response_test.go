package message1_1_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	datatransfer "github.com/filecoin-project/go-data-transfer"
	message1_1 "github.com/filecoin-project/go-data-transfer/message/message1_1prime"
	"github.com/filecoin-project/go-data-transfer/testutil"
)

func TestResponseMessageForProtocol(t *testing.T) {
	id := datatransfer.TransferID(rand.Int31())
	voucherResult := testutil.NewFakeDTType()
	response, err := message1_1.NewResponse(id, false, true, voucherResult.Type(), voucherResult) // not accepted
	require.NoError(t, err)

	// v1.2 protocol
	out, err := response.MessageForProtocol(datatransfer.ProtocolDataTransfer1_2)
	require.NoError(t, err)
	require.Equal(t, response, out)

	resp, ok := (out).(datatransfer.Response)
	require.True(t, ok)
	require.True(t, resp.IsPaused())
	require.Equal(t, voucherResult.Type(), resp.VoucherResultType())
	require.True(t, resp.IsVoucherResult())

	// random protocol
	out, err = response.MessageForProtocol("RAND")
	require.Error(t, err)
	require.Nil(t, out)
}
