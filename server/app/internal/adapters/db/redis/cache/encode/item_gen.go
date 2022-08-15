package encode

// THIS FILE WAS PRODUCED BY THE MSGP CODE GENERATION TOOL (github.com/dchenk/msgp).
// DO NOT EDIT.

import (
	"github.com/dchenk/msgp/msgp"
)

// DecodeMsg implements msgp.Decoder
func (z *Item) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch string(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "data":
			z.Data, err = dc.ReadBytes(z.Data)
			if err != nil {
				return
			}
		case "created":
			z.Created, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "desc":
			z.Description, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encoder
func (z *Item) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "name"
	err = en.Append(0x84, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "data"
	err = en.Append(0xa4, 0x64, 0x61, 0x74, 0x61)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Data)
	if err != nil {
		return
	}
	// write "created"
	err = en.Append(0xa7, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteTime(z.Created)
	if err != nil {
		return
	}
	// write "desc"
	err = en.Append(0xa4, 0x64, 0x65, 0x73, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.Description)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Item) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "name"
	o = append(o, 0x84, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "data"
	o = append(o, 0xa4, 0x64, 0x61, 0x74, 0x61)
	o = msgp.AppendBytes(o, z.Data)
	// string "created"
	o = append(o, 0xa7, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendTime(o, z.Created)
	// string "desc"
	o = append(o, 0xa4, 0x64, 0x65, 0x73, 0x63)
	o = msgp.AppendString(o, z.Description)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Item) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch string(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "data":
			z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)
			if err != nil {
				return
			}
		case "created":
			z.Created, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "desc":
			z.Description, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Item) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.BytesPrefixSize + len(z.Data) + 8 + msgp.TimeSize + 5 + msgp.StringPrefixSize + len(z.Description)
	return
}
