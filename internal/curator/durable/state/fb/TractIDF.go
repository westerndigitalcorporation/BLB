// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TractIDF struct {
	_tab flatbuffers.Struct
}

func (rcv *TractIDF) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TractIDF) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *TractIDF) B1() uint16 {
	return rcv._tab.GetUint16(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *TractIDF) MutateB1(n uint16) bool {
	return rcv._tab.MutateUint16(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *TractIDF) B2() uint16 {
	return rcv._tab.GetUint16(rcv._tab.Pos + flatbuffers.UOffsetT(2))
}
func (rcv *TractIDF) MutateB2(n uint16) bool {
	return rcv._tab.MutateUint16(rcv._tab.Pos+flatbuffers.UOffsetT(2), n)
}

func (rcv *TractIDF) B3() uint16 {
	return rcv._tab.GetUint16(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *TractIDF) MutateB3(n uint16) bool {
	return rcv._tab.MutateUint16(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func (rcv *TractIDF) B4() uint16 {
	return rcv._tab.GetUint16(rcv._tab.Pos + flatbuffers.UOffsetT(6))
}
func (rcv *TractIDF) MutateB4(n uint16) bool {
	return rcv._tab.MutateUint16(rcv._tab.Pos+flatbuffers.UOffsetT(6), n)
}

func (rcv *TractIDF) B5() uint16 {
	return rcv._tab.GetUint16(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *TractIDF) MutateB5(n uint16) bool {
	return rcv._tab.MutateUint16(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func CreateTractIDF(builder *flatbuffers.Builder, b1 uint16, b2 uint16, b3 uint16, b4 uint16, b5 uint16) flatbuffers.UOffsetT {
	builder.Prep(2, 10)
	builder.PrependUint16(b5)
	builder.PrependUint16(b4)
	builder.PrependUint16(b3)
	builder.PrependUint16(b2)
	builder.PrependUint16(b1)
	return builder.Offset()
}
