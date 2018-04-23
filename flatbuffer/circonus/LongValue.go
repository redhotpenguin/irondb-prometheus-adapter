// automatically generated by the FlatBuffers compiler, do not modify

package circonus

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LongValue struct {
	_tab flatbuffers.Table
}

func GetRootAsLongValue(buf []byte, offset flatbuffers.UOffsetT) *LongValue {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LongValue{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *LongValue) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LongValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LongValue) Value() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LongValue) MutateValue(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func LongValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func LongValueAddValue(builder *flatbuffers.Builder, value int64) {
	builder.PrependInt64Slot(0, value, 0)
}
func LongValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
