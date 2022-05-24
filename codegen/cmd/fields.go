package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type Field struct {
	Name    string
	TypeNum uint64

	EncodingLength    func(*Field) (string, error)
	EncodeInto        func(*Field) (string, error)
	EncoderStruct     func(*Field) (string, error)
	InitEncoderStruct func(*Field) (string, error)
	DecodeFrom        func(*Field) (string, error)
}

type TlvModel struct {
	Name   string
	Fields []*Field
}

type strErrBuf struct {
	b   strings.Builder
	err error
}

func (m *strErrBuf) printlne(str string, err error) {
	if m.err == nil {
		if err == nil {
			_, m.err = fmt.Fprintln(&m.b, str)
		} else {
			m.err = err
		}
	}
}

func (m *strErrBuf) printlnf(format string, args ...interface{}) {
	if m.err == nil {
		fmt.Fprintf(&m.b, format, args...)
		m.b.WriteRune('\n')
	}
}

func (m *strErrBuf) output() (string, error) {
	return m.b.String(), m.err
}

func (m *strErrBuf) executeTemplate(t *template.Template, data interface{}) {
	if m.err == nil {
		m.err = t.Execute(&m.b, data)
	}
}

func GenTypeNumLen(typeNum uint64) (string, error) {
	var ret uint
	switch {
	case typeNum <= 0xfc:
		ret = 1
	case typeNum <= 0xffff:
		ret = 3
	case typeNum <= 0xffffffff:
		ret = 5
	default:
		ret = 9
	}
	return fmt.Sprintf("\tl += %d", ret), nil
}

func GenEncodeTypeNum(typeNum uint64) (string, error) {
	ret := ""
	switch {
	case typeNum <= 0xfc:
		ret += fmt.Sprintf("\tbuf[pos] = %d\n", typeNum)
		ret += fmt.Sprintf("\tpos += %d", 1)
	case typeNum <= 0xffff:
		ret += fmt.Sprintf("\tbuf[pos] = %d\n", 0xfd)
		ret += fmt.Sprintf("\tbinary.BigEndian.PutUint16(buf[pos+1:], uint16(%d))\n", typeNum)
		ret += fmt.Sprintf("\tpos += %d", 3)
	case typeNum <= 0xffffffff:
		ret += fmt.Sprintf("\tbuf[pos] = %d\n", 0xfe)
		ret += fmt.Sprintf("\tbinary.BigEndian.PutUint32(buf[pos+1:], uint32(%d))\n", typeNum)
		ret += fmt.Sprintf("\tpos += %d", 5)
	default:
		ret += fmt.Sprintf("\tbuf[pos] = %d\n", 0xff)
		ret += fmt.Sprintf("\tbinary.BigEndian.PutUint64(buf[pos+1:], uint64(%d))\n", typeNum)
		ret += fmt.Sprintf("\tpos += %d", 9)
	}
	return ret, nil
}

func GenNaturalNumberLen(code string, isTlv bool) (string, error) {
	const Temp = `switch x := {{.Code}}; {
	{{- if .IsTlv}}
	case x <= 0xfc:
		l += 1
	{{- else}}
	case x <= 0xff:
		l += 2
	{{- end}}
	case x <= 0xffff:
		l += 3
	case x <= 0xffffffff:
		l += 5
	default:
		l += 9
	}`
	t := template.Must(template.New("NaturalNumberLen").Parse(Temp))
	data := struct {
		IsTlv bool
		Code  string
	}{
		IsTlv: isTlv,
		Code:  code,
	}
	b := strings.Builder{}
	err := t.Execute(&b, data)
	return b.String(), err
}

func GenNaturalNumberEncode(code string, isTlv bool) (string, error) {
	const Temp = `switch x := {{.Code}}; {
	{{- if .IsTlv}}
	case x <= 0xfc:
		buf[pos] = byte(x)
		pos += 1
	{{- else}}
	case x <= 0xff:
		buf[pos] = 1
		buf[pos+1] = byte(x)
		pos += 2
	{{- end}}
	case x <= 0xffff:
		buf[pos] = {{if .IsTlv -}} 0xfd {{- else -}} 2 {{- end}}
		binary.BigEndian.PutUint16(buf[pos+1:], uint16(x))
		pos += 3
	case x <= 0xffffffff:
		buf[pos] = {{if .IsTlv -}} 0xfe {{- else -}} 4 {{- end}}
		binary.BigEndian.PutUint32(buf[pos+1:], uint32(x))
		pos += 5
	default:
		buf[pos] = {{if .IsTlv -}} 0xff {{- else -}} 8 {{- end}}
		binary.BigEndian.PutUint64(buf[pos+1:], uint64(x))
		pos += 9
	}`
	t := template.Must(template.New("NaturalNumberEncode").Parse(Temp))
	data := struct {
		IsTlv bool
		Code  string
	}{
		IsTlv: isTlv,
		Code:  code,
	}
	b := strings.Builder{}
	err := t.Execute(&b, data)
	return b.String(), err
}

func GenTlvNumberDecode(code string) (string, error) {
	const Temp = `switch x := buf[pos]; {
		case x <= 0xfc:
			{{.}} = uint64(buf[pos])
			pos += 1
		case x == 0xfd:
			{{.}} = uint64(binary.BigEndian.Uint16(buf[pos+1 : pos+3]))
			pos += 3
		case x == 0xfe:
			{{.}} = uint64(binary.BigEndian.Uint32(buf[pos+1 : pos+5]))
			pos += 5
		case x == 0xff:
			{{.}} = binary.BigEndian.Uint64(buf[pos+1 : pos+9])
			pos += 9
		}`
	t := template.Must(template.New("TlvNumberDecode").Parse(Temp))
	b := strings.Builder{}
	err := t.Execute(&b, code)
	return b.String(), err
}

func GenNaturalNumberDecode(code string) (string, error) {
	const Temp = `switch l {
	case 1:
		{{.}} = uint64(buf[pos])
	case 2:
		{{.}} = uint64(binary.BigEndian.Uint16(buf[pos : pos+2]))
	case 4:
		{{.}} = uint64(binary.BigEndian.Uint32(buf[pos : pos+4]))
	case 8:
		{{.}} = binary.BigEndian.Uint64(buf[pos : pos+8])
	default:
		failFlag = true
	}`
	t := template.Must(template.New("NaturalNumberDecode").Parse(Temp))
	b := strings.Builder{}
	err := t.Execute(&b, code)
	return b.String(), err
}

func NewNaturalField(name string, typeNum uint64) *Field {
	return &Field{
		Name:    name,
		TypeNum: typeNum,
		EncodingLength: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlne(GenTypeNumLen(f.TypeNum))
			g.printlne(GenNaturalNumberLen("v."+f.Name, false))
			return g.output()
		},
		EncodeInto: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlne(GenEncodeTypeNum(f.TypeNum))
			g.printlne(GenNaturalNumberEncode("v."+f.Name, false))
			return g.output()
		},
		DecodeFrom: func(f *Field) (string, error) {
			return GenNaturalNumberDecode("v." + f.Name)
		},
		EncoderStruct:     func(f *Field) (string, error) { return "", nil },
		InitEncoderStruct: func(f *Field) (string, error) { return "", nil },
	}
}

func NewTimeField(name string, typeNum uint64) *Field {
	return &Field{
		Name:    name,
		TypeNum: typeNum,
		EncodingLength: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlne(GenTypeNumLen(f.TypeNum))
			g.printlne(GenNaturalNumberLen(fmt.Sprintf("uint64(v.%s/time.Millisecond)", f.Name), false))
			return g.output()
		},
		EncodeInto: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlne(GenEncodeTypeNum(f.TypeNum))
			g.printlne(GenNaturalNumberEncode(fmt.Sprintf("uint64(v.%s/time.Millisecond)", f.Name), false))
			return g.output()
		},
		DecodeFrom: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlnf("var timeInt uint64 = 0")
			g.printlne(GenNaturalNumberDecode("timeInt"))
			g.printlnf("v.%s = time.Duration(timeInt) * time.Millisecond", f.Name)
			return g.output()
		},
		EncoderStruct:     func(f *Field) (string, error) { return "", nil },
		InitEncoderStruct: func(f *Field) (string, error) { return "", nil },
	}
}

func NewBinaryField(name string, typeNum uint64) *Field {
	return &Field{
		Name:    name,
		TypeNum: typeNum,
		EncodingLength: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlne(GenTypeNumLen(f.TypeNum))
			g.printlne(GenNaturalNumberLen(fmt.Sprintf("len(v.%s)", f.Name), true))
			g.printlnf("l += uint(len(v.%s))", f.Name)
			return g.output()
		},
		EncodeInto: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlne(GenEncodeTypeNum(f.TypeNum))
			g.printlne(GenNaturalNumberEncode(fmt.Sprintf("len(v.%s)", f.Name), true))
			g.printlnf("copy(buf[pos:], v.%s)", f.Name)
			g.printlnf("pos += uint(len(v.%s))", f.Name)
			return g.output()
		},
		DecodeFrom: func(f *Field) (string, error) {
			return fmt.Sprintf("v.%s = buf[pos : pos+uint(l)]\n", f.Name), nil
		},
		EncoderStruct:     func(f *Field) (string, error) { return "", nil },
		InitEncoderStruct: func(f *Field) (string, error) { return "", nil },
	}
}

func NewSignatureField(name string, typeNum uint64) *Field {
	return &Field{
		Name:    name,
		TypeNum: typeNum,
		EncodingLength: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlnf("if v.%s != nil {", f.Name)
			g.printlne(GenTypeNumLen(f.TypeNum))
			g.printlnf("l += 33")
			g.printlnf("}")
			return g.output()
		},
		EncodeInto: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlnf("if v.%s != nil {", f.Name)
			g.printlne(GenEncodeTypeNum(f.TypeNum))
			g.printlnf("buf[pos] = 32")
			g.printlnf("pos += 33")
			g.printlnf("}")
			return g.output()
		},
		DecodeFrom: func(f *Field) (string, error) {
			return fmt.Sprintf("v.%s = buf[pos : pos+uint(l)]\n", f.Name), nil
		},
		EncoderStruct:     func(f *Field) (string, error) { return "", nil },
		InitEncoderStruct: func(f *Field) (string, error) { return "", nil },
	}
}

func NewNameField(name string, typeNum uint64) *Field {
	return &Field{
		Name:    name,
		TypeNum: typeNum,
		EncodingLength: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlne(GenTypeNumLen(f.TypeNum))
			g.printlne(GenNaturalNumberLen(fmt.Sprintf("e.%s_length", f.Name), true))
			g.printlnf("l += e.%s_length", f.Name)
			return g.output()
		},
		EncodeInto: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlne(GenEncodeTypeNum(f.TypeNum))
			g.printlne(GenNaturalNumberEncode(fmt.Sprintf("e.%s_length", f.Name), true))
			const Temp = `for _, c := range v.{{.}} {
				copy(buf[pos:], c)
				pos += uint(len(c))
			}
			`
			t := template.Must(template.New("NameEncodeInto").Parse(Temp))
			g.executeTemplate(t, f.Name)
			return g.output()
		},
		DecodeFrom: func(f *Field) (string, error) {
			var g strErrBuf
			const Temp = `v.{{.Name}} = make([][]byte, 0)
			endName := pos + uint(l)
			startName := pos
			for pos < endName {
				var componentLen uint64 = 0
				componentStart := pos
				{{call .GenTlvNumberDecode "componentLen"}}
				{{call .GenTlvNumberDecode "componentLen"}}
				v.{{.Name}} = append(v.{{.Name}}, buf[componentStart : pos+uint(componentLen)])
				pos += uint(componentLen)
			}
			pos = startName
			`
			data := struct {
				Name               string
				GenTlvNumberDecode func(string) (string, error)
			}{
				Name:               f.Name,
				GenTlvNumberDecode: GenTlvNumberDecode,
			}
			t := template.Must(template.New("NameEncodeInto").Parse(Temp))
			g.executeTemplate(t, data)
			return g.output()
		},
		EncoderStruct: func(f *Field) (string, error) {
			return fmt.Sprintf("%s_length uint", f.Name), nil
		},
		InitEncoderStruct: func(f *Field) (string, error) {
			var g strErrBuf
			const Temp = `e.{{.}}_length = 0
			for _, c := range v.{{.}} {
				e.{{.}}_length += uint(len(c))
			}
			`
			t := template.Must(template.New("NameInitEncoder").Parse(Temp))
			g.executeTemplate(t, f.Name)
			return g.output()
		},
	}
}

func NewStructField(name string, typeNum uint64, structClass string) *Field {
	return &Field{
		Name:    name,
		TypeNum: typeNum,
		EncodingLength: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlne(GenTypeNumLen(f.TypeNum))
			g.printlne(GenNaturalNumberLen(fmt.Sprintf("e.%s_encoder.length", f.Name), true))
			g.printlnf("l += e.%s_encoder.length", f.Name)
			return g.output()
		},
		EncodeInto: func(f *Field) (string, error) {
			var g strErrBuf
			g.printlne(GenEncodeTypeNum(f.TypeNum))
			g.printlne(GenNaturalNumberEncode(fmt.Sprintf("e.%s_encoder.length", f.Name), true))
			g.printlnf("pos += e.%s_encoder.encodeInto(v.%s, buf[pos:])", f.Name, f.Name)
			return g.output()
		},
		DecodeFrom: func(f *Field) (string, error) {
			var g strErrBuf
			const Temp = `struValue, _ := Parse{{.StructClass}}(buf[pos : pos+uint(l)], ignoreCritical)
			v.{{.Name}} = struValue
			if struValue == nil {
				failFlag = true
			}
			`
			data := struct {
				Name        string
				StructClass string
			}{
				Name:        f.Name,
				StructClass: structClass,
			}
			t := template.Must(template.New("StructDecodeFrom").Parse(Temp))
			g.executeTemplate(t, data)
			return g.output()
		},
		EncoderStruct: func(f *Field) (string, error) {
			return fmt.Sprintf("%s_encoder %sEncoder", f.Name, structClass), nil
		},
		InitEncoderStruct: func(f *Field) (string, error) {
			return fmt.Sprintf("e.%s_encoder.init(v.%s)", f.Name, f.Name), nil
		},
	}
}

func (m *TlvModel) genEncoderStruct(buf *bytes.Buffer) error {
	const Temp = `type {{.Name}}Encoder struct {
		length uint
		{{- range $f := .Fields}}
		{{call $f.EncoderStruct $f}}
		{{- end}}
	}
	`
	t := template.Must(template.New("ModelEncoderStruct").Parse(Temp))
	return t.Execute(buf, m)
}

func (m *TlvModel) initEncoderStruct(buf *bytes.Buffer) error {
	const Temp = `func (e *{{.Name}}Encoder) init(v *{{.Name}}) {
		{{- range $f := .Fields}}
		{{call $f.InitEncoderStruct $f}}
		{{- end}}
		var l uint = 0
		{{- range $f := .Fields}}
		{{call $f.EncodingLength $f}}
		{{- end}}
		e.length = l
	}
	`
	t := template.Must(template.New("ModelInitEncoderStruct").Parse(Temp))
	return t.Execute(buf, m)
}

func (m *TlvModel) genEncodeInto(buf *bytes.Buffer) error {
	const Temp = `func (e *{{.Name}}Encoder) encodeInto(v *{{.Name}}, buf []byte) uint {
		var pos uint = 0
		{{- range $f := .Fields}}
		{{call $f.EncodeInto $f}}
		{{- end}}
		return pos
	}
	`
	t := template.Must(template.New("ModelEncodeInto").Parse(Temp))
	return t.Execute(buf, m)
}

func (m *TlvModel) genDecodeFrom(buf *bytes.Buffer) error {
	// TODO: Handle sequence field
	const Temp = `func Parse{{.Name}}(buf []byte, ignoreCritical bool) (*{{.Name}}, uint) {
		var pos uint = 0
		var progress int = -1
		v := &{{.Name}}{}
		for pos < uint(len(buf)) {
			var typ uint64 = 0
			var l uint64 = 0
			{{call .GenTlvNumberDecode "typ"}}
			{{call .GenTlvNumberDecode "l"}}
			failFlag := false
			switch {
				{{- range $i, $f := .Fields}}
			case progress < {{$i}} && typ == {{$f.TypeNum}}:
				progress = {{$i}}
				{{call $f.DecodeFrom $f}}
				{{- end}}
			default:
				failFlag = true
			}
			if failFlag && !ignoreCritical && {{.IsCritical}} {
				return nil, 0
			}
			pos += uint(l)
		}
    return v, pos
	}
	`
	data := struct {
		Name               string
		Fields             []*Field
		GenTlvNumberDecode func(string) (string, error)
		IsCritical         string
	}{
		Name:               m.Name,
		Fields:             m.Fields,
		GenTlvNumberDecode: GenTlvNumberDecode,
		IsCritical:         `((typ <= 31) || ((typ & 1) == 1))`,
	}
	t := template.Must(template.New("ModelDecodeFrom").Parse(Temp))
	return t.Execute(buf, data)
}

func (m *TlvModel) Generate(buf *bytes.Buffer) error {
	err := m.genEncoderStruct(buf)
	if err != nil {
		return err
	}
	buf.WriteRune('\n')
	err = m.initEncoderStruct(buf)
	if err != nil {
		return err
	}
	buf.WriteRune('\n')
	err = m.genEncodeInto(buf)
	if err != nil {
		return err
	}
	buf.WriteRune('\n')
	err = m.genDecodeFrom(buf)
	if err != nil {
		return err
	}
	buf.WriteRune('\n')
	return nil
}

func ParseField(name string, typeNum uint64, fieldTypeStr string) *Field {
	switch {
	case strings.HasPrefix(fieldTypeStr, "natural"):
		return NewNaturalField(name, typeNum)
	case strings.HasPrefix(fieldTypeStr, "time"):
		return NewTimeField(name, typeNum)
	case strings.HasPrefix(fieldTypeStr, "binary"):
		return NewBinaryField(name, typeNum)
	case strings.HasPrefix(fieldTypeStr, "name"):
		return NewNameField(name, typeNum)
	case strings.HasPrefix(fieldTypeStr, "signature"):
		return NewSignatureField(name, typeNum)
	case strings.HasPrefix(fieldTypeStr, "struct"):
		const PrefixLen = len("struct:")
		return NewStructField(name, typeNum, fieldTypeStr[PrefixLen:])
	default:
		return nil
	}
}
