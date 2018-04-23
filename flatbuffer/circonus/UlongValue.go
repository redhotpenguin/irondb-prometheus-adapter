// automatically generated by the FlatBuffers compiler, do not modify

package circonus

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UlongValue struct {
	_tab flatbuffers.Table
}

func GetRootAsUlongValue(buf []byte, offset flatbuffers.UOffsetT) *UlongValue {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UlongValue{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *UlongValue) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UlongValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *UlongValue) Value() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *UlongValue) MutateValue(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func UlongValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func UlongValueAddValue(builder *flatbuffers.Builder, value uint64) {
	builder.PrependUint64Slot(0, value, 0)
}
func UlongValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
