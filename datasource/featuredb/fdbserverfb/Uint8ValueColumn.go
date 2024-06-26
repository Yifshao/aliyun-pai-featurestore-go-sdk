// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fdbserverfb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UInt8ValueColumn struct {
	_tab flatbuffers.Table
}

func GetRootAsUInt8ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *UInt8ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UInt8ValueColumn{}
	x.Init(buf, n+offset)
	return x
}

func FinishUInt8ValueColumnBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsUInt8ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *UInt8ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &UInt8ValueColumn{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedUInt8ValueColumnBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *UInt8ValueColumn) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UInt8ValueColumn) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *UInt8ValueColumn) Value(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *UInt8ValueColumn) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *UInt8ValueColumn) ValueBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *UInt8ValueColumn) MutateValue(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func UInt8ValueColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func UInt8ValueColumnAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func UInt8ValueColumnStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func UInt8ValueColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
