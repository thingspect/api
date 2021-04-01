# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/device.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from common import device_pb2 as common_dot_device__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import field_mask_pb2 as google_dot_protobuf_dot_field__mask__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.api import field_behavior_pb2 as google_dot_api_dot_field__behavior__pb2
from protoc_gen_openapiv2.options import annotations_pb2 as protoc__gen__openapiv2_dot_options_dot_annotations__pb2
from validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='api/device.proto',
  package='thingspect.api',
  syntax='proto3',
  serialized_options=b'Z github.com/thingspect/api/go/api',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x10\x61pi/device.proto\x12\x0ethingspect.api\x1a\x13\x63ommon/device.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a google/protobuf/field_mask.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x17validate/validate.proto\"M\n\x13\x43reateDeviceRequest\x12\x36\n\x06\x64\x65vice\x18\x01 \x01(\x0b\x32\x19.thingspect.common.DeviceB\x0b\xfa\x42\x05\x8a\x01\x02\x10\x01\xe0\x41\x02\"\xf4\x02\n\x1a\x43reateDeviceLoRaWANRequest\x12\x17\n\x02id\x18\x01 \x01(\tB\x0b\xfa\x42\x05r\x03\xb0\x01\x01\xe0\x41\x03\x12q\n\x14gateway_lorawan_type\x18\x02 \x01(\x0b\x32=.thingspect.api.CreateDeviceLoRaWANRequest.GatewayLoRaWANTypeH\x00R\x12gatewayLoRaWANType\x12n\n\x13\x64\x65vice_lorawan_type\x18\x03 \x01(\x0b\x32<.thingspect.api.CreateDeviceLoRaWANRequest.DeviceLoRaWANTypeH\x00R\x11\x64\x65viceLoRaWANType\x1a\x14\n\x12GatewayLoRaWANType\x1a\x31\n\x11\x44\x65viceLoRaWANType\x12\x1c\n\x07\x61pp_key\x18\x01 \x01(\tB\x0b\xfa\x42\x05r\x03\x98\x01 \xe0\x41\x02\x42\x11\n\ntype_oneof\x12\x03\xf8\x42\x01\"+\n\x10GetDeviceRequest\x12\x17\n\x02id\x18\x01 \x01(\tB\x0b\xfa\x42\x05r\x03\xb0\x01\x01\xe0\x41\x02\"~\n\x13UpdateDeviceRequest\x12\x36\n\x06\x64\x65vice\x18\x01 \x01(\x0b\x32\x19.thingspect.common.DeviceB\x0b\xfa\x42\x05\x8a\x01\x02\x10\x01\xe0\x41\x02\x12/\n\x0bupdate_mask\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.FieldMask\"5\n\x1a\x44\x65leteDeviceLoRaWANRequest\x12\x17\n\x02id\x18\x01 \x01(\tB\x0b\xfa\x42\x05r\x03\xb0\x01\x01\xe0\x41\x02\".\n\x13\x44\x65leteDeviceRequest\x12\x17\n\x02id\x18\x01 \x01(\tB\x0b\xfa\x42\x05r\x03\xb0\x01\x01\xe0\x41\x02\"\\\n\x12ListDevicesRequest\x12\x1b\n\tpage_size\x18\x01 \x01(\x05\x42\x08\xfa\x42\x05\x1a\x03\x18\xfa\x01\x12\x12\n\npage_token\x18\x02 \x01(\t\x12\x15\n\x03tag\x18\x03 \x01(\tB\x08\xfa\x42\x05r\x03\x18\xff\x01\"n\n\x13ListDevicesResponse\x12*\n\x07\x64\x65vices\x18\x01 \x03(\x0b\x32\x19.thingspect.common.Device\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\x12\x12\n\ntotal_size\x18\x03 \x01(\x05\x32\x87\x08\n\rDeviceService\x12\xae\x01\n\x0c\x43reateDevice\x12#.thingspect.api.CreateDeviceRequest\x1a\x19.thingspect.common.Device\"^\x82\xd3\xe4\x93\x02\x15\"\x0b/v1/devices:\x06\x64\x65vice\x92\x41@J>\n\x03\x32\x30\x31\x12\x37\n\x16\x41 successful response.\x12\x1d\n\x1b\x1a\x19.thingspect.common.Device\x12\xa4\x01\n\x13\x43reateDeviceLoRaWAN\x12*.thingspect.api.CreateDeviceLoRaWANRequest\x1a\x16.google.protobuf.Empty\"I\x82\xd3\xe4\x93\x02\x1d\"\x18/v1/devices/{id}/lorawan:\x01*\x92\x41#J!\n\x03\x32\x30\x34\x12\x1a\n\x16\x41 successful response.\x12\x00\x12\x62\n\tGetDevice\x12 .thingspect.api.GetDeviceRequest\x1a\x19.thingspect.common.Device\"\x18\x82\xd3\xe4\x93\x02\x12\x12\x10/v1/devices/{id}\x12\x9a\x01\n\x0cUpdateDevice\x12#.thingspect.api.UpdateDeviceRequest\x1a\x19.thingspect.common.Device\"J\x82\xd3\xe4\x93\x02\x44\x1a\x17/v1/devices/{device.id}:\x06\x64\x65viceZ!2\x17/v1/devices/{device.id}:\x06\x64\x65vice\x12\xa1\x01\n\x13\x44\x65leteDeviceLoRaWAN\x12*.thingspect.api.DeleteDeviceLoRaWANRequest\x1a\x16.google.protobuf.Empty\"F\x82\xd3\xe4\x93\x02\x1a*\x18/v1/devices/{id}/lorawan\x92\x41#J!\n\x03\x32\x30\x34\x12\x1a\n\x16\x41 successful response.\x12\x00\x12\x8b\x01\n\x0c\x44\x65leteDevice\x12#.thingspect.api.DeleteDeviceRequest\x1a\x16.google.protobuf.Empty\">\x82\xd3\xe4\x93\x02\x12*\x10/v1/devices/{id}\x92\x41#J!\n\x03\x32\x30\x34\x12\x1a\n\x16\x41 successful response.\x12\x00\x12k\n\x0bListDevices\x12\".thingspect.api.ListDevicesRequest\x1a#.thingspect.api.ListDevicesResponse\"\x13\x82\xd3\xe4\x93\x02\r\x12\x0b/v1/devicesB\"Z github.com/thingspect/api/go/apib\x06proto3'
  ,
  dependencies=[common_dot_device__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,google_dot_protobuf_dot_field__mask__pb2.DESCRIPTOR,google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_api_dot_field__behavior__pb2.DESCRIPTOR,protoc__gen__openapiv2_dot_options_dot_annotations__pb2.DESCRIPTOR,validate_dot_validate__pb2.DESCRIPTOR,])




_CREATEDEVICEREQUEST = _descriptor.Descriptor(
  name='CreateDeviceRequest',
  full_name='thingspect.api.CreateDeviceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='device', full_name='thingspect.api.CreateDeviceRequest.device', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005\212\001\002\020\001\340A\002', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=256,
  serialized_end=333,
)


_CREATEDEVICELORAWANREQUEST_GATEWAYLORAWANTYPE = _descriptor.Descriptor(
  name='GatewayLoRaWANType',
  full_name='thingspect.api.CreateDeviceLoRaWANRequest.GatewayLoRaWANType',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=618,
  serialized_end=638,
)

_CREATEDEVICELORAWANREQUEST_DEVICELORAWANTYPE = _descriptor.Descriptor(
  name='DeviceLoRaWANType',
  full_name='thingspect.api.CreateDeviceLoRaWANRequest.DeviceLoRaWANType',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='app_key', full_name='thingspect.api.CreateDeviceLoRaWANRequest.DeviceLoRaWANType.app_key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005r\003\230\001 \340A\002', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=640,
  serialized_end=689,
)

_CREATEDEVICELORAWANREQUEST = _descriptor.Descriptor(
  name='CreateDeviceLoRaWANRequest',
  full_name='thingspect.api.CreateDeviceLoRaWANRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='thingspect.api.CreateDeviceLoRaWANRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005r\003\260\001\001\340A\003', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='gateway_lorawan_type', full_name='thingspect.api.CreateDeviceLoRaWANRequest.gateway_lorawan_type', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='gatewayLoRaWANType', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='device_lorawan_type', full_name='thingspect.api.CreateDeviceLoRaWANRequest.device_lorawan_type', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='deviceLoRaWANType', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_CREATEDEVICELORAWANREQUEST_GATEWAYLORAWANTYPE, _CREATEDEVICELORAWANREQUEST_DEVICELORAWANTYPE, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='type_oneof', full_name='thingspect.api.CreateDeviceLoRaWANRequest.type_oneof',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[], serialized_options=b'\370B\001'),
  ],
  serialized_start=336,
  serialized_end=708,
)


_GETDEVICEREQUEST = _descriptor.Descriptor(
  name='GetDeviceRequest',
  full_name='thingspect.api.GetDeviceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='thingspect.api.GetDeviceRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005r\003\260\001\001\340A\002', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=710,
  serialized_end=753,
)


_UPDATEDEVICEREQUEST = _descriptor.Descriptor(
  name='UpdateDeviceRequest',
  full_name='thingspect.api.UpdateDeviceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='device', full_name='thingspect.api.UpdateDeviceRequest.device', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005\212\001\002\020\001\340A\002', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='update_mask', full_name='thingspect.api.UpdateDeviceRequest.update_mask', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=755,
  serialized_end=881,
)


_DELETEDEVICELORAWANREQUEST = _descriptor.Descriptor(
  name='DeleteDeviceLoRaWANRequest',
  full_name='thingspect.api.DeleteDeviceLoRaWANRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='thingspect.api.DeleteDeviceLoRaWANRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005r\003\260\001\001\340A\002', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=883,
  serialized_end=936,
)


_DELETEDEVICEREQUEST = _descriptor.Descriptor(
  name='DeleteDeviceRequest',
  full_name='thingspect.api.DeleteDeviceRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='thingspect.api.DeleteDeviceRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005r\003\260\001\001\340A\002', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=938,
  serialized_end=984,
)


_LISTDEVICESREQUEST = _descriptor.Descriptor(
  name='ListDevicesRequest',
  full_name='thingspect.api.ListDevicesRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='page_size', full_name='thingspect.api.ListDevicesRequest.page_size', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005\032\003\030\372\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='page_token', full_name='thingspect.api.ListDevicesRequest.page_token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='tag', full_name='thingspect.api.ListDevicesRequest.tag', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005r\003\030\377\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=986,
  serialized_end=1078,
)


_LISTDEVICESRESPONSE = _descriptor.Descriptor(
  name='ListDevicesResponse',
  full_name='thingspect.api.ListDevicesResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='devices', full_name='thingspect.api.ListDevicesResponse.devices', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='next_page_token', full_name='thingspect.api.ListDevicesResponse.next_page_token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='total_size', full_name='thingspect.api.ListDevicesResponse.total_size', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1080,
  serialized_end=1190,
)

_CREATEDEVICEREQUEST.fields_by_name['device'].message_type = common_dot_device__pb2._DEVICE
_CREATEDEVICELORAWANREQUEST_GATEWAYLORAWANTYPE.containing_type = _CREATEDEVICELORAWANREQUEST
_CREATEDEVICELORAWANREQUEST_DEVICELORAWANTYPE.containing_type = _CREATEDEVICELORAWANREQUEST
_CREATEDEVICELORAWANREQUEST.fields_by_name['gateway_lorawan_type'].message_type = _CREATEDEVICELORAWANREQUEST_GATEWAYLORAWANTYPE
_CREATEDEVICELORAWANREQUEST.fields_by_name['device_lorawan_type'].message_type = _CREATEDEVICELORAWANREQUEST_DEVICELORAWANTYPE
_CREATEDEVICELORAWANREQUEST.oneofs_by_name['type_oneof'].fields.append(
  _CREATEDEVICELORAWANREQUEST.fields_by_name['gateway_lorawan_type'])
_CREATEDEVICELORAWANREQUEST.fields_by_name['gateway_lorawan_type'].containing_oneof = _CREATEDEVICELORAWANREQUEST.oneofs_by_name['type_oneof']
_CREATEDEVICELORAWANREQUEST.oneofs_by_name['type_oneof'].fields.append(
  _CREATEDEVICELORAWANREQUEST.fields_by_name['device_lorawan_type'])
_CREATEDEVICELORAWANREQUEST.fields_by_name['device_lorawan_type'].containing_oneof = _CREATEDEVICELORAWANREQUEST.oneofs_by_name['type_oneof']
_UPDATEDEVICEREQUEST.fields_by_name['device'].message_type = common_dot_device__pb2._DEVICE
_UPDATEDEVICEREQUEST.fields_by_name['update_mask'].message_type = google_dot_protobuf_dot_field__mask__pb2._FIELDMASK
_LISTDEVICESRESPONSE.fields_by_name['devices'].message_type = common_dot_device__pb2._DEVICE
DESCRIPTOR.message_types_by_name['CreateDeviceRequest'] = _CREATEDEVICEREQUEST
DESCRIPTOR.message_types_by_name['CreateDeviceLoRaWANRequest'] = _CREATEDEVICELORAWANREQUEST
DESCRIPTOR.message_types_by_name['GetDeviceRequest'] = _GETDEVICEREQUEST
DESCRIPTOR.message_types_by_name['UpdateDeviceRequest'] = _UPDATEDEVICEREQUEST
DESCRIPTOR.message_types_by_name['DeleteDeviceLoRaWANRequest'] = _DELETEDEVICELORAWANREQUEST
DESCRIPTOR.message_types_by_name['DeleteDeviceRequest'] = _DELETEDEVICEREQUEST
DESCRIPTOR.message_types_by_name['ListDevicesRequest'] = _LISTDEVICESREQUEST
DESCRIPTOR.message_types_by_name['ListDevicesResponse'] = _LISTDEVICESRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateDeviceRequest = _reflection.GeneratedProtocolMessageType('CreateDeviceRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEDEVICEREQUEST,
  '__module__' : 'api.device_pb2'
  # @@protoc_insertion_point(class_scope:thingspect.api.CreateDeviceRequest)
  })
_sym_db.RegisterMessage(CreateDeviceRequest)

CreateDeviceLoRaWANRequest = _reflection.GeneratedProtocolMessageType('CreateDeviceLoRaWANRequest', (_message.Message,), {

  'GatewayLoRaWANType' : _reflection.GeneratedProtocolMessageType('GatewayLoRaWANType', (_message.Message,), {
    'DESCRIPTOR' : _CREATEDEVICELORAWANREQUEST_GATEWAYLORAWANTYPE,
    '__module__' : 'api.device_pb2'
    # @@protoc_insertion_point(class_scope:thingspect.api.CreateDeviceLoRaWANRequest.GatewayLoRaWANType)
    })
  ,

  'DeviceLoRaWANType' : _reflection.GeneratedProtocolMessageType('DeviceLoRaWANType', (_message.Message,), {
    'DESCRIPTOR' : _CREATEDEVICELORAWANREQUEST_DEVICELORAWANTYPE,
    '__module__' : 'api.device_pb2'
    # @@protoc_insertion_point(class_scope:thingspect.api.CreateDeviceLoRaWANRequest.DeviceLoRaWANType)
    })
  ,
  'DESCRIPTOR' : _CREATEDEVICELORAWANREQUEST,
  '__module__' : 'api.device_pb2'
  # @@protoc_insertion_point(class_scope:thingspect.api.CreateDeviceLoRaWANRequest)
  })
_sym_db.RegisterMessage(CreateDeviceLoRaWANRequest)
_sym_db.RegisterMessage(CreateDeviceLoRaWANRequest.GatewayLoRaWANType)
_sym_db.RegisterMessage(CreateDeviceLoRaWANRequest.DeviceLoRaWANType)

GetDeviceRequest = _reflection.GeneratedProtocolMessageType('GetDeviceRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETDEVICEREQUEST,
  '__module__' : 'api.device_pb2'
  # @@protoc_insertion_point(class_scope:thingspect.api.GetDeviceRequest)
  })
_sym_db.RegisterMessage(GetDeviceRequest)

UpdateDeviceRequest = _reflection.GeneratedProtocolMessageType('UpdateDeviceRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEDEVICEREQUEST,
  '__module__' : 'api.device_pb2'
  # @@protoc_insertion_point(class_scope:thingspect.api.UpdateDeviceRequest)
  })
_sym_db.RegisterMessage(UpdateDeviceRequest)

DeleteDeviceLoRaWANRequest = _reflection.GeneratedProtocolMessageType('DeleteDeviceLoRaWANRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEDEVICELORAWANREQUEST,
  '__module__' : 'api.device_pb2'
  # @@protoc_insertion_point(class_scope:thingspect.api.DeleteDeviceLoRaWANRequest)
  })
_sym_db.RegisterMessage(DeleteDeviceLoRaWANRequest)

DeleteDeviceRequest = _reflection.GeneratedProtocolMessageType('DeleteDeviceRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEDEVICEREQUEST,
  '__module__' : 'api.device_pb2'
  # @@protoc_insertion_point(class_scope:thingspect.api.DeleteDeviceRequest)
  })
_sym_db.RegisterMessage(DeleteDeviceRequest)

ListDevicesRequest = _reflection.GeneratedProtocolMessageType('ListDevicesRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTDEVICESREQUEST,
  '__module__' : 'api.device_pb2'
  # @@protoc_insertion_point(class_scope:thingspect.api.ListDevicesRequest)
  })
_sym_db.RegisterMessage(ListDevicesRequest)

ListDevicesResponse = _reflection.GeneratedProtocolMessageType('ListDevicesResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTDEVICESRESPONSE,
  '__module__' : 'api.device_pb2'
  # @@protoc_insertion_point(class_scope:thingspect.api.ListDevicesResponse)
  })
_sym_db.RegisterMessage(ListDevicesResponse)


DESCRIPTOR._options = None
_CREATEDEVICEREQUEST.fields_by_name['device']._options = None
_CREATEDEVICELORAWANREQUEST_DEVICELORAWANTYPE.fields_by_name['app_key']._options = None
_CREATEDEVICELORAWANREQUEST.oneofs_by_name['type_oneof']._options = None
_CREATEDEVICELORAWANREQUEST.fields_by_name['id']._options = None
_GETDEVICEREQUEST.fields_by_name['id']._options = None
_UPDATEDEVICEREQUEST.fields_by_name['device']._options = None
_DELETEDEVICELORAWANREQUEST.fields_by_name['id']._options = None
_DELETEDEVICEREQUEST.fields_by_name['id']._options = None
_LISTDEVICESREQUEST.fields_by_name['page_size']._options = None
_LISTDEVICESREQUEST.fields_by_name['tag']._options = None

_DEVICESERVICE = _descriptor.ServiceDescriptor(
  name='DeviceService',
  full_name='thingspect.api.DeviceService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1193,
  serialized_end=2224,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateDevice',
    full_name='thingspect.api.DeviceService.CreateDevice',
    index=0,
    containing_service=None,
    input_type=_CREATEDEVICEREQUEST,
    output_type=common_dot_device__pb2._DEVICE,
    serialized_options=b'\202\323\344\223\002\025\"\013/v1/devices:\006device\222A@J>\n\003201\0227\n\026A successful response.\022\035\n\033\032\031.thingspect.common.Device',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateDeviceLoRaWAN',
    full_name='thingspect.api.DeviceService.CreateDeviceLoRaWAN',
    index=1,
    containing_service=None,
    input_type=_CREATEDEVICELORAWANREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=b'\202\323\344\223\002\035\"\030/v1/devices/{id}/lorawan:\001*\222A#J!\n\003204\022\032\n\026A successful response.\022\000',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetDevice',
    full_name='thingspect.api.DeviceService.GetDevice',
    index=2,
    containing_service=None,
    input_type=_GETDEVICEREQUEST,
    output_type=common_dot_device__pb2._DEVICE,
    serialized_options=b'\202\323\344\223\002\022\022\020/v1/devices/{id}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateDevice',
    full_name='thingspect.api.DeviceService.UpdateDevice',
    index=3,
    containing_service=None,
    input_type=_UPDATEDEVICEREQUEST,
    output_type=common_dot_device__pb2._DEVICE,
    serialized_options=b'\202\323\344\223\002D\032\027/v1/devices/{device.id}:\006deviceZ!2\027/v1/devices/{device.id}:\006device',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteDeviceLoRaWAN',
    full_name='thingspect.api.DeviceService.DeleteDeviceLoRaWAN',
    index=4,
    containing_service=None,
    input_type=_DELETEDEVICELORAWANREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=b'\202\323\344\223\002\032*\030/v1/devices/{id}/lorawan\222A#J!\n\003204\022\032\n\026A successful response.\022\000',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteDevice',
    full_name='thingspect.api.DeviceService.DeleteDevice',
    index=5,
    containing_service=None,
    input_type=_DELETEDEVICEREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=b'\202\323\344\223\002\022*\020/v1/devices/{id}\222A#J!\n\003204\022\032\n\026A successful response.\022\000',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='ListDevices',
    full_name='thingspect.api.DeviceService.ListDevices',
    index=6,
    containing_service=None,
    input_type=_LISTDEVICESREQUEST,
    output_type=_LISTDEVICESRESPONSE,
    serialized_options=b'\202\323\344\223\002\r\022\013/v1/devices',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_DEVICESERVICE)

DESCRIPTOR.services_by_name['DeviceService'] = _DEVICESERVICE

# @@protoc_insertion_point(module_scope)
