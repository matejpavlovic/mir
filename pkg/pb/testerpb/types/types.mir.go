// Code generated by Mir codegen. DO NOT EDIT.

package testerpbtypes

import (
	mirreflect "github.com/matejpavlovic/mir/codegen/mirreflect"
	testerpb "github.com/matejpavlovic/mir/pkg/pb/testerpb"
	reflectutil "github.com/matejpavlovic/mir/pkg/util/reflectutil"
)

type Tester struct{}

func TesterFromPb(pb *testerpb.Tester) *Tester {
	if pb == nil {
		return nil
	}
	return &Tester{}
}

func (m *Tester) Pb() *testerpb.Tester {
	if m == nil {
		return nil
	}
	pbMessage := &testerpb.Tester{}
	{
	}

	return pbMessage
}

func (*Tester) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*testerpb.Tester]()}
}
