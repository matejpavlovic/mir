package stats

import trantorpbtypes "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"

type Tracker interface {
	Submit(tx *trantorpbtypes.Transaction)
	Deliver(tx *trantorpbtypes.Transaction)
}
