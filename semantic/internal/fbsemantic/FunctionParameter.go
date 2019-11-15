// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FunctionParameter struct {
	_tab flatbuffers.Table
}

func GetRootAsFunctionParameter(buf []byte, offset flatbuffers.UOffsetT) *FunctionParameter {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FunctionParameter{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FunctionParameter) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FunctionParameter) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FunctionParameter) Loc(obj *SourceLocation) *SourceLocation {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SourceLocation)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *FunctionParameter) Key(obj *Identifier) *Identifier {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Identifier)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func FunctionParameterStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func FunctionParameterAddLoc(builder *flatbuffers.Builder, loc flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(loc), 0)
}
func FunctionParameterAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(key), 0)
}
func FunctionParameterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}