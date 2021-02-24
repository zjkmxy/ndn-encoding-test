// Generated by the generator, DO NOT modify manually
package codegen

import (
	"encoding/binary"
	"time"
)

type MetaInfoEncoder struct {
	length uint
}

func (e *MetaInfoEncoder) init(v *MetaInfo) {
	var l uint = 0
	l += 1
	switch x := v.ContentType; {
	case x <= 0xff:
		l += 2
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += 1
	switch x := uint64(v.FreshnessPeriod / time.Millisecond); {
	case x <= 0xff:
		l += 2
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += 1
	switch x := len(v.FinalBlockID); {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += uint(len(v.FinalBlockID))
	e.length = l
}

func (e *MetaInfoEncoder) encodeInto(v *MetaInfo, buf []byte) uint {
	var pos uint = 0
	buf[pos] = 24
	pos += 1
	switch x := v.ContentType; {
	case x <= 0xff:
		buf[pos] = 1
		buf[pos+1] = byte(x)
		pos += 2
	case x <= 0xffff:
		buf[pos] = 2
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = 4
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = 8
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}
	buf[pos] = 25
	pos += 1
	switch x := uint64(v.FreshnessPeriod / time.Millisecond); {
	case x <= 0xff:
		buf[pos] = 1
		buf[pos+1] = byte(x)
		pos += 2
	case x <= 0xffff:
		buf[pos] = 2
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = 4
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = 8
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}
	buf[pos] = 26
	pos += 1
	switch x := len(v.FinalBlockID); {
	case x <= 0xfc:
		buf[pos] = byte(x)
		pos += 1
	case x <= 0xffff:
		buf[pos] = 0xfd
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = 0xfe
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = 0xff
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}
	copy(buf[pos:], v.FinalBlockID)
	pos += uint(len(v.FinalBlockID))
	return pos
}

func ParseMetaInfo(buf []byte, ignoreCritical bool) (*MetaInfo, uint) {
	var pos uint = 0
	var progress uint = 0
	v := &MetaInfo{}
	for pos < uint(len(buf)) {
		var typ uint64 = 0
		var l uint64 = 0
		switch x := buf[pos]; {
		case x <= 0xfc:
			typ = uint64(buf[pos])
			pos += 1
		case x == 0xfd:
			typ = uint64(binary.BigEndian.Uint16(buf[pos+1 : pos+3]))
			pos += 3
		case x == 0xfe:
			typ = uint64(binary.BigEndian.Uint32(buf[pos+1 : pos+5]))
			pos += 5
		case x == 0xff:
			typ = uint64(binary.BigEndian.Uint64(buf[pos+1 : pos+9]))
			pos += 9
		}
		switch x := buf[pos]; {
		case x <= 0xfc:
			l = uint64(buf[pos])
			pos += 1
		case x == 0xfd:
			l = uint64(binary.BigEndian.Uint16(buf[pos+1 : pos+3]))
			pos += 3
		case x == 0xfe:
			l = uint64(binary.BigEndian.Uint32(buf[pos+1 : pos+5]))
			pos += 5
		case x == 0xff:
			l = uint64(binary.BigEndian.Uint64(buf[pos+1 : pos+9]))
			pos += 9
		}
		failFlag := false
		switch {
		case progress <= 0 && typ == 24:
			progress = 0
			switch l {
			case 1:
				v.ContentType = uint64(buf[pos])
			case 2:
				v.ContentType = uint64(binary.BigEndian.Uint16(buf[pos : pos+2]))
			case 4:
				v.ContentType = uint64(binary.BigEndian.Uint32(buf[pos : pos+4]))
			case 8:
				v.ContentType = uint64(binary.BigEndian.Uint64(buf[pos : pos+8]))
			default:
				failFlag = true
			}
		case progress <= 1 && typ == 25:
			progress = 1
			var timeInt uint64 = 0
			switch l {
			case 1:
				timeInt = uint64(buf[pos])
			case 2:
				timeInt = uint64(binary.BigEndian.Uint16(buf[pos : pos+2]))
			case 4:
				timeInt = uint64(binary.BigEndian.Uint32(buf[pos : pos+4]))
			case 8:
				timeInt = uint64(binary.BigEndian.Uint64(buf[pos : pos+8]))
			default:
				failFlag = true
			}
			v.FreshnessPeriod = time.Duration(timeInt) * time.Millisecond
		case progress <= 2 && typ == 26:
			progress = 2
			v.FinalBlockID = buf[pos : pos+uint(l)]
		default:
			failFlag = true
		}
		if failFlag && !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
			return nil, 0
		}
		pos += uint(l)
	}
	return v, pos
}

type SignatureInfoEncoder struct {
	length uint
}

func (e *SignatureInfoEncoder) init(v *SignatureInfo) {
	var l uint = 0
	l += 1
	switch x := v.SignatureType; {
	case x <= 0xff:
		l += 2
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	e.length = l
}

func (e *SignatureInfoEncoder) encodeInto(v *SignatureInfo, buf []byte) uint {
	var pos uint = 0
	buf[pos] = 27
	pos += 1
	switch x := v.SignatureType; {
	case x <= 0xff:
		buf[pos] = 1
		buf[pos+1] = byte(x)
		pos += 2
	case x <= 0xffff:
		buf[pos] = 2
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = 4
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = 8
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}
	return pos
}

func ParseSignatureInfo(buf []byte, ignoreCritical bool) (*SignatureInfo, uint) {
	var pos uint = 0
	var progress uint = 0
	v := &SignatureInfo{}
	for pos < uint(len(buf)) {
		var typ uint64 = 0
		var l uint64 = 0
		switch x := buf[pos]; {
		case x <= 0xfc:
			typ = uint64(buf[pos])
			pos += 1
		case x == 0xfd:
			typ = uint64(binary.BigEndian.Uint16(buf[pos+1 : pos+3]))
			pos += 3
		case x == 0xfe:
			typ = uint64(binary.BigEndian.Uint32(buf[pos+1 : pos+5]))
			pos += 5
		case x == 0xff:
			typ = uint64(binary.BigEndian.Uint64(buf[pos+1 : pos+9]))
			pos += 9
		}
		switch x := buf[pos]; {
		case x <= 0xfc:
			l = uint64(buf[pos])
			pos += 1
		case x == 0xfd:
			l = uint64(binary.BigEndian.Uint16(buf[pos+1 : pos+3]))
			pos += 3
		case x == 0xfe:
			l = uint64(binary.BigEndian.Uint32(buf[pos+1 : pos+5]))
			pos += 5
		case x == 0xff:
			l = uint64(binary.BigEndian.Uint64(buf[pos+1 : pos+9]))
			pos += 9
		}
		failFlag := false
		switch {
		case progress <= 0 && typ == 27:
			progress = 0
			switch l {
			case 1:
				v.SignatureType = uint64(buf[pos])
			case 2:
				v.SignatureType = uint64(binary.BigEndian.Uint16(buf[pos : pos+2]))
			case 4:
				v.SignatureType = uint64(binary.BigEndian.Uint32(buf[pos : pos+4]))
			case 8:
				v.SignatureType = uint64(binary.BigEndian.Uint64(buf[pos : pos+8]))
			default:
				failFlag = true
			}
		default:
			failFlag = true
		}
		if failFlag && !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
			return nil, 0
		}
		pos += uint(l)
	}
	return v, pos
}

type DataEncoder struct {
	length                uint
	Name_length           uint
	MetaInfo_encoder      MetaInfoEncoder
	SignatureInfo_encoder SignatureInfoEncoder
}

func (e *DataEncoder) init(v *Data) {
	e.Name_length = 0
	for _, c := range v.Name {
		e.Name_length += uint(len(c))
	}
	e.MetaInfo_encoder.init(v.MetaInfo)
	e.SignatureInfo_encoder.init(v.SignatureInfo)
	var l uint = 0
	l += 1
	switch x := e.Name_length; {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += e.Name_length
	l += 1
	switch x := e.MetaInfo_encoder.length; {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += e.MetaInfo_encoder.length
	l += 1
	switch x := len(v.Content); {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += uint(len(v.Content))
	l += 1
	switch x := e.SignatureInfo_encoder.length; {
	case x <= 0xfc:
		l += 1
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}
	l += e.SignatureInfo_encoder.length
	l += 1
	l += 33
	e.length = l
}

func (e *DataEncoder) encodeInto(v *Data, buf []byte) uint {
	var pos uint = 0
	buf[pos] = 7
	pos += 1
	switch x := e.Name_length; {
	case x <= 0xfc:
		buf[pos] = byte(x)
		pos += 1
	case x <= 0xffff:
		buf[pos] = 0xfd
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = 0xfe
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = 0xff
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}
	for _, c := range v.Name {
		copy(buf[pos:], c)
		pos += uint(len(c))
	}
	buf[pos] = 20
	pos += 1
	switch x := e.MetaInfo_encoder.length; {
	case x <= 0xfc:
		buf[pos] = byte(x)
		pos += 1
	case x <= 0xffff:
		buf[pos] = 0xfd
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = 0xfe
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = 0xff
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}
	pos += e.MetaInfo_encoder.encodeInto(v.MetaInfo, buf[pos:])
	buf[pos] = 21
	pos += 1
	switch x := len(v.Content); {
	case x <= 0xfc:
		buf[pos] = byte(x)
		pos += 1
	case x <= 0xffff:
		buf[pos] = 0xfd
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = 0xfe
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = 0xff
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}
	copy(buf[pos:], v.Content)
	pos += uint(len(v.Content))
	buf[pos] = 22
	pos += 1
	switch x := e.SignatureInfo_encoder.length; {
	case x <= 0xfc:
		buf[pos] = byte(x)
		pos += 1
	case x <= 0xffff:
		buf[pos] = 0xfd
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = 0xfe
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = 0xff
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}
	pos += e.SignatureInfo_encoder.encodeInto(v.SignatureInfo, buf[pos:])
	buf[pos] = 23
	pos += 1
	buf[pos] = 32
	pos += 33
	return pos
}

func ParseData(buf []byte, ignoreCritical bool) (*Data, uint) {
	var pos uint = 0
	var progress uint = 0
	v := &Data{}
	for pos < uint(len(buf)) {
		var typ uint64 = 0
		var l uint64 = 0
		switch x := buf[pos]; {
		case x <= 0xfc:
			typ = uint64(buf[pos])
			pos += 1
		case x == 0xfd:
			typ = uint64(binary.BigEndian.Uint16(buf[pos+1 : pos+3]))
			pos += 3
		case x == 0xfe:
			typ = uint64(binary.BigEndian.Uint32(buf[pos+1 : pos+5]))
			pos += 5
		case x == 0xff:
			typ = uint64(binary.BigEndian.Uint64(buf[pos+1 : pos+9]))
			pos += 9
		}
		switch x := buf[pos]; {
		case x <= 0xfc:
			l = uint64(buf[pos])
			pos += 1
		case x == 0xfd:
			l = uint64(binary.BigEndian.Uint16(buf[pos+1 : pos+3]))
			pos += 3
		case x == 0xfe:
			l = uint64(binary.BigEndian.Uint32(buf[pos+1 : pos+5]))
			pos += 5
		case x == 0xff:
			l = uint64(binary.BigEndian.Uint64(buf[pos+1 : pos+9]))
			pos += 9
		}
		failFlag := false
		switch {
		case progress <= 0 && typ == 7:
			progress = 0
			v.Name = make([][]byte, 0)
			endName := pos + uint(l)
			startName := pos
			for pos < endName {
				var componentLen uint64 = 0
				componentStart := pos
				switch x := buf[pos]; {
				case x <= 0xfc:
					componentLen = uint64(buf[pos])
					pos += 1
				case x == 0xfd:
					componentLen = uint64(binary.BigEndian.Uint16(buf[pos+1 : pos+3]))
					pos += 3
				case x == 0xfe:
					componentLen = uint64(binary.BigEndian.Uint32(buf[pos+1 : pos+5]))
					pos += 5
				case x == 0xff:
					componentLen = uint64(binary.BigEndian.Uint64(buf[pos+1 : pos+9]))
					pos += 9
				}
				switch x := buf[pos]; {
				case x <= 0xfc:
					componentLen = uint64(buf[pos])
					pos += 1
				case x == 0xfd:
					componentLen = uint64(binary.BigEndian.Uint16(buf[pos+1 : pos+3]))
					pos += 3
				case x == 0xfe:
					componentLen = uint64(binary.BigEndian.Uint32(buf[pos+1 : pos+5]))
					pos += 5
				case x == 0xff:
					componentLen = uint64(binary.BigEndian.Uint64(buf[pos+1 : pos+9]))
					pos += 9
				}
				v.Name = append(v.Name, buf[componentStart:pos+uint(componentLen)])
				pos += uint(componentLen)
			}
			pos = startName
		case progress <= 1 && typ == 20:
			progress = 1
			struValue, _ := ParseMetaInfo(buf[pos:pos+uint(l)], ignoreCritical)
			v.MetaInfo = struValue
			if struValue == nil {
				failFlag = true
			}
		case progress <= 2 && typ == 21:
			progress = 2
			v.Content = buf[pos : pos+uint(l)]
		case progress <= 3 && typ == 22:
			progress = 3
			struValue, _ := ParseSignatureInfo(buf[pos:pos+uint(l)], ignoreCritical)
			v.SignatureInfo = struValue
			if struValue == nil {
				failFlag = true
			}
		case progress <= 4 && typ == 23:
			progress = 4
			v.SignatureValue = buf[pos : pos+uint(l)]
		default:
			failFlag = true
		}
		if failFlag && !ignoreCritical && ((typ <= 31) || ((typ & 1) == 1)) {
			return nil, 0
		}
		pos += uint(l)
	}
	return v, pos
}
