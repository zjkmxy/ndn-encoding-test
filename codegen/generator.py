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


class NaturalField(Field):
  def generate_length(self):
    self.type_num_len()
    natural_number_len(f'v.{self.name}', False)

  def generate_encode_into(self):
    self.encode_type_num()
    natural_number_encode(f'v.{self.name}', False)


class TimeField(Field):
  def generate_length(self):
    self.type_num_len()
    natural_number_len(f'uint64(v.{self.name}/time.Millisecond)', False)

  def generate_encode_into(self):
    self.encode_type_num()
    natural_number_encode(f'uint64(v.{self.name}/time.Millisecond)', False)



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


class SignatureField(Field):
  def generate_length(self):
    self.type_num_len()
    print(f'\tl += 33')

  def generate_encode_into(self):
    self.encode_type_num()
    print(f'\tbuf[pos] = 32')
    print(f'\tpos += 33')



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

  def generate(self):
    self.generate_encoder_struct()
    print()
    self.initialize_encoder_struct()
    print()
    self.generate_encode_into()
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
