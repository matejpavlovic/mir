package checkpointpb

import (
	reflect "reflect"
)

func (*Message) ReflectTypeOptions() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*Message_Checkpoint)(nil)),
	}
}

func (*Event) ReflectTypeOptions() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*Event_EpochConfig)(nil)),
		reflect.TypeOf((*Event_StableCheckpoint)(nil)),
		reflect.TypeOf((*Event_EpochProgress)(nil)),
	}
}