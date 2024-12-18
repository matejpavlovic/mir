// Code generated by Mir codegen. DO NOT EDIT.

package isspbdsl

import (
	dsl "github.com/matejpavlovic/mir/pkg/dsl"
	types1 "github.com/matejpavlovic/mir/pkg/pb/availabilitypb/types"
	events "github.com/matejpavlovic/mir/pkg/pb/isspb/events"
	types2 "github.com/matejpavlovic/mir/pkg/pb/trantorpb/types"
	types "github.com/matejpavlovic/mir/pkg/trantor/types"
	stdtypes "github.com/matejpavlovic/mir/stdtypes"
)

// Module-specific dsl functions for emitting events.

func PushCheckpoint(m dsl.Module, destModule stdtypes.ModuleID) {
	dsl.EmitMirEvent(m, events.PushCheckpoint(destModule))
}

func SBDeliver(m dsl.Module, destModule stdtypes.ModuleID, sn types.SeqNr, data []uint8, aborted bool, leader stdtypes.NodeID, instanceId stdtypes.ModuleID) {
	dsl.EmitMirEvent(m, events.SBDeliver(destModule, sn, data, aborted, leader, instanceId))
}

func DeliverCert(m dsl.Module, destModule stdtypes.ModuleID, sn types.SeqNr, cert *types1.Cert, empty bool) {
	dsl.EmitMirEvent(m, events.DeliverCert(destModule, sn, cert, empty))
}

func NewConfig(m dsl.Module, destModule stdtypes.ModuleID, epochNr types.EpochNr, membership *types2.Membership) {
	dsl.EmitMirEvent(m, events.NewConfig(destModule, epochNr, membership))
}
