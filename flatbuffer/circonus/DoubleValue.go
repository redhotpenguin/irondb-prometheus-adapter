// automatically generated by the FlatBuffers compiler, do not modify

package circonus

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DoubleValue struct {
	_tab flatbuffers.Table
}

func GetRootAsDoubleValue(buf []byte, offset flatbuffers.UOffsetT) *DoubleValue {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DoubleValue{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DoubleValue) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DoubleValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DoubleValue) Value() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *DoubleValue) MutateValue(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

func DoubleValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DoubleValueAddValue(builder *flatbuffers.Builder, value float64) {
	builder.PrependFloat64Slot(0, value, 0.0)
}
func DoubleValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
