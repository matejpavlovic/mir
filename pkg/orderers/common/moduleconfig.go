package common

import (
	t "github.com/matejpavlovic/mir/stdtypes"
)

type ModuleConfig struct {
	Self t.ModuleID

	App            t.ModuleID
	Ava            t.ModuleID
	Crypto         t.ModuleID
	Hasher         t.ModuleID
	Net            t.ModuleID
	Ord            t.ModuleID
	PPrepValidator t.ModuleID
	Timer          t.ModuleID
}
