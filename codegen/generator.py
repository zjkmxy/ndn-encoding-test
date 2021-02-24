# This is a quick and dirty hack. A better way is to use Go's AST
# SA: https://github.com/golang/tools/blob/master/cmd/stringer/stringer.go

def type_num_len(type_num):
  if type_num <= 0xfc:
    return 1
  elif type_num <= 0xffff:
    return 3
  elif type_num <= 0xffffffff:
    return 5
  else:
    return 9


def is_critical():
  return "((typ <= 31) || ((typ & 1) == 1))"


def natural_number_len(code, is_tlv):
  print(f'\tswitch x := {code}; {{')
  if is_tlv:
    print(f'\tcase x <= 0xfc:')
    print(f'\t\tl += 1')
  else:
    print(f'\tcase x <= 0xff:')
    # TLV Length + The natural number
    print(f'\t\tl += 2')
  print(f'\tcase x <= 0xffff:')
  print(f'\t\tl += 3')
  print(f'\tcase x <= 0xffffffff:')
  print(f'\t\tl += 5')
  print(f'\tdefault:')
  print(f'\t\tl += 9')
  print(f'\t}}')


def natural_number_encode(code, is_tlv):
  print(f'\tswitch x := {code}; {{')
  if is_tlv:
    print(f'\tcase x <= 0xfc:')
    print(f'\t\tbuf[pos] = byte(x)')
    print(f'\t\tpos += 1')
  else:
    print(f'\tcase x <= 0xff:')
    print(f'\t\tbuf[pos] = 1')
    print(f'\t\tbuf[pos+1] = byte(x)')
    print(f'\t\tpos += 2')
  print(f'\tcase x <= 0xffff:')
  if is_tlv:
    print(f'\t\tbuf[pos] = 0xfd')
  else:
    print(f'\t\tbuf[pos] = 2')
  print(f'\t\tbinary.BigEndian.PutUint16(buf[pos+1:], uint16(x))')
  print(f'\t\tpos += 3')
  print(f'\tcase x <= 0xffffffff:')
  if is_tlv:
    print(f'\t\tbuf[pos] = 0xfe')
  else:
    print(f'\t\tbuf[pos] = 4')
  print(f'\t\tbinary.BigEndian.PutUint32(buf[pos+1:], uint32(x))')
  print(f'\t\tpos += 5')
  print(f'\tdefault:')
  if is_tlv:
    print(f'\t\tbuf[pos] = 0xff')
  else:
    print(f'\t\tbuf[pos] = 8')
  print(f'\t\tbinary.BigEndian.PutUint64(buf[pos+1:], uint64(x))')
  print(f'\t\tpos += 9')
  print(f'\t}}')


def tlv_number_decode(code):
  # TODO: Should check the length.
  print(f'\t\tswitch x := buf[pos]; {{')
  print(f'\t\tcase x <= 0xfc:')
  print(f'\t\t\t{code} = uint64(buf[pos])')
  print(f'\t\t\tpos += 1')
  print(f'\t\tcase x == 0xfd:')
  print(f'\t\t\t{code} = uint64(binary.BigEndian.Uint16(buf[pos+1 : pos+3]))')
  print(f'\t\t\tpos += 3')
  print(f'\t\tcase x == 0xfe:')
  print(f'\t\t\t{code} = uint64(binary.BigEndian.Uint32(buf[pos+1 : pos+5]))')
  print(f'\t\t\tpos += 5')
  print(f'\t\tcase x == 0xff:')
  print(f'\t\t\t{code} = uint64(binary.BigEndian.Uint64(buf[pos+1 : pos+9]))')
  print(f'\t\t\tpos += 9')
  print(f'\t\t}}')


def natural_number_decode(code):
  print(f'\t\t\tswitch l {{')
  print(f'\t\t\tcase 1:')
  print(f'\t\t\t\t{code} = uint64(buf[pos])')
  print(f'\t\t\tcase 2:')
  print(f'\t\t\t\t{code} = uint64(binary.BigEndian.Uint16(buf[pos : pos+2]))')
  print(f'\t\t\tcase 4:')
  print(f'\t\t\t\t{code} = uint64(binary.BigEndian.Uint32(buf[pos : pos+4]))')
  print(f'\t\t\tcase 8:')
  print(f'\t\t\t\t{code} = uint64(binary.BigEndian.Uint64(buf[pos : pos+8]))')
  print(f'\t\t\tdefault:')
  # Failed
  print(f'\t\t\t\tfailFlag = true')
  print(f'\t\t\t}}')


class Field:
  name: str
  type_num: int

  def __init__(self, name, type_num):
      self.name = name
      self.type_num = type_num

  def type_num_len(self):
    print(f'\tl += {type_num_len(self.type_num)}')

  def encode_type_num(self):
    if self.type_num <= 0xfc:
      print(f'\tbuf[pos] = {self.type_num}')
    elif self.type_num <= 0xffff:
      print(f'\tbuf[pos] = 0xfd')
      print(f'\tbinary.BigEndian.PutUint16(buf[pos+1:], uint16({self.type_num}))')
    elif self.type_num <= 0xffffffff:
      print(f'\tbuf[pos] = 0xfe')
      print(f'\tbinary.BigEndian.PutUint32(buf[pos+1:], uint32({self.type_num}))')
    else:
      print(f'\tbuf[pos] = 0xff')
      print(f'\tbinary.BigEndian.PutUint64(buf[pos+1:], uint64({self.type_num}))')
    print(f'\tpos += {type_num_len(self.type_num)}')

  def generate_length(self):
    pass

  def generate_encode_into(self):
    pass

  def generate_encoder_struct(self):
    pass

  def initialize_encoder_struct(self):
    pass

  def generate_decode_from(self):
    pass


class NaturalField(Field):
  def generate_length(self):
    self.type_num_len()
    natural_number_len(f'v.{self.name}', False)

  def generate_encode_into(self):
    self.encode_type_num()
    natural_number_encode(f'v.{self.name}', False)

  def generate_decode_from(self):
    natural_number_decode(f'v.{self.name}')

class TimeField(Field):
  def generate_length(self):
    self.type_num_len()
    natural_number_len(f'uint64(v.{self.name}/time.Millisecond)', False)

  def generate_encode_into(self):
    self.encode_type_num()
    natural_number_encode(f'uint64(v.{self.name}/time.Millisecond)', False)

  def generate_decode_from(self):
    print(f'\t\t\tvar timeInt uint64 = 0')
    natural_number_decode('timeInt')
    print(f'\t\t\tv.{self.name} = time.Duration(timeInt) * time.Millisecond')


class BinaryField(Field):
  def generate_length(self):
    self.type_num_len()
    natural_number_len(f'len(v.{self.name})', True)
    print(f'\tl += uint(len(v.{self.name}))')

  def generate_encode_into(self):
    self.encode_type_num()
    natural_number_encode(f'len(v.{self.name})', True)
    print(f'\tcopy(buf[pos:], v.{self.name})')
    print(f'\tpos += uint(len(v.{self.name}))')

  def generate_decode_from(self):
    print(f'\t\t\tv.{self.name} = buf[pos : pos+uint(l)]')


class SignatureField(Field):
  def generate_length(self):
    self.type_num_len()
    print(f'\tl += 33')

  def generate_encode_into(self):
    self.encode_type_num()
    print(f'\tbuf[pos] = 32')
    print(f'\tpos += 33')

  def generate_decode_from(self):
    print(f'\t\t\tv.{self.name} = buf[pos : pos+uint(l)]')


class NameField(Field):
  def generate_encoder_struct(self):
    print(f'\t{self.name}_length uint')

  def initialize_encoder_struct(self):
    print(f'\te.{self.name}_length = 0')
    print(f'\tfor _, c := range v.{self.name} {{')
    print(f'\t\te.{self.name}_length += uint(len(c))')
    print(f'\t}}')

  def generate_length(self):
    self.type_num_len()
    natural_number_len(f'e.{self.name}_length', True)
    print(f'\tl += e.{self.name}_length')

  def generate_encode_into(self):
    self.encode_type_num()
    natural_number_encode(f'e.{self.name}_length', True)
    print(f'\tfor _, c := range v.{self.name} {{')
    print(f'\t\tcopy(buf[pos:], c)')
    print(f'\t\tpos += uint(len(c))')
    print(f'\t}}')

  def generate_decode_from(self):
    print(f'\t\t\tv.{self.name} = make([][]byte, 0)')
    print(f'\t\t\tendName := pos + uint(l)')
    print(f'\t\t\tstartName := pos')
    print(f'\t\t\tfor pos < endName {{')
    print(f'\t\t\t\tvar componentLen uint64 = 0')
    print(f'\t\t\t\tcomponentStart := pos')
    tlv_number_decode('componentLen')  # ComponentType
    tlv_number_decode('componentLen')
    print(f'\t\t\t\tv.{self.name} = append(v.{self.name}, buf[componentStart : pos+uint(componentLen)])')
    print(f'\t\t\t\tpos += uint(componentLen)')
    print(f'\t\t\t}}')
    print(f'\t\t\tpos = startName')


class StructField(Field):
  def __init__(self, name, type_num, struct_class):
      super().__init__(name, type_num)
      self.struct_class = struct_class

  def generate_encoder_struct(self):
    print(f'\t{self.name}_encoder {self.struct_class}Encoder')

  def initialize_encoder_struct(self):
    # Shoule do this at the base class: whether it's a pointer or not
    print(f'\te.{self.name}_encoder.init(v.{self.name})')

  def generate_length(self):
    self.type_num_len()
    natural_number_len(f'e.{self.name}_encoder.length', True)
    print(f'\tl += e.{self.name}_encoder.length')

  def generate_encode_into(self):
    self.encode_type_num()
    natural_number_encode(f'e.{self.name}_encoder.length', True)
    print(f'\tpos += e.{self.name}_encoder.encodeInto(v.{self.name}, buf[pos:])')

  def generate_decode_from(self):
    print(f'\t\t\tstruValue, _ := Parse{self.struct_class}(buf[pos : pos+uint(l)], ignoreCritical)')
    print(f'\t\t\tv.{self.name} = struValue')
    print(f'\t\t\tif struValue == nil {{')
    print(f'\t\t\t\tfailFlag = true')
    print(f'\t\t\t}}')


class TLVStruct:
  def __init__(self, name, fields):
    self.fields = fields[:]
    self.name = name

  def generate_encoder_struct(self):
    print(f'type {self.name}Encoder struct {{')
    print(f'\tlength uint')
    for f in self.fields:
      f.generate_encoder_struct()
    print(f'}}')

  def initialize_encoder_struct(self):
    print(f'func (e *{self.name}Encoder) init(v *{self.name}) {{')
    for f in self.fields:
      f.initialize_encoder_struct()
    print(f'\tvar l uint = 0')
    for f in self.fields:
      f.generate_length()
    print(f'\te.length = l')
    print(f'}}')

  def generate_encode_into(self):
    print(f'func (e *{self.name}Encoder) encodeInto(v *{self.name}, buf []byte) uint {{')
    print(f'\tvar pos uint = 0')
    for f in self.fields:
      f.generate_encode_into()
    print(f'\treturn pos')
    print(f'}}')

  def generate_decode_from(self):
    # TODO: Should return an error
    print(f'func Parse{self.name}(buf []byte, ignoreCritical bool) (*{self.name}, uint) {{')
    print(f'\tvar pos uint = 0')
    print(f'\tvar progress uint = -1')
    print(f'\tv := &{self.name}{{}}')
    print(f'\tfor pos < uint(len(buf)) {{')
    print(f'\t\tvar typ uint64 = 0')
    print(f'\t\tvar l uint64 = 0')
    tlv_number_decode('typ')
    tlv_number_decode('l')
    print(f'\t\tfailFlag := false')
    print(f'\t\tswitch {{')
    for i, f in enumerate(self.fields):
      # TODO: sequence
      print(f'\t\tcase progress < {i} && typ == {f.type_num}:')
      print(f'\t\t\tprogress = {i}')
      f.generate_decode_from()
    print(f'\t\tdefault:')
    print(f'\t\t\tfailFlag = true')
    print(f'\t\t}}')  # End switch
    print(f'\t\tif failFlag && !ignoreCritical && {is_critical()} {{')
    print(f'\t\t\treturn nil, 0')
    print(f'\t\t}}')  # End if
    print(f'\t\tpos += uint(l)')
    print(f'\t}}')  # End for
    print(f'\treturn v, pos')
    print(f'}}')  # End func

  def generate(self):
    self.generate_encoder_struct()
    print()
    self.initialize_encoder_struct()
    print()
    self.generate_encode_into()
    print()
    self.generate_decode_from()
    print()


def generate_header(package):
  print(f'// Generated by the generator, DO NOT modify manually')
  print(f'package {package}')
  print()
  print('import (')
  print(f'\t"encoding/binary"')
  print(f'\t"time"')
  print(')')
  print()


def main():
  meta_info = TLVStruct('MetaInfo', [
    NaturalField('ContentType', 0x18),
    TimeField('FreshnessPeriod', 0x19),
    BinaryField('FinalBlockID', 0x1a),
  ])
  signature_info = TLVStruct('SignatureInfo', [NaturalField('SignatureType', 0x1b)])
  data = TLVStruct('Data', [
    NameField('Name', 0x07),
    StructField('MetaInfo', 0x14, 'MetaInfo'),
    BinaryField('Content', 0x15),
    StructField('SignatureInfo', 0x16, 'SignatureInfo'),
    SignatureField('SignatureValue', 0x17),
  ])

  generate_header('codegen')
  meta_info.generate()
  signature_info.generate()
  data.generate()


if __name__ == "__main__":
  main()
