# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mqtt/thingspect_payload.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from common import thingspect_datapoint_pb2 as common_dot_thingspect__datapoint__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dmqtt/thingspect_payload.proto\x12\x0fthingspect.mqtt\x1a!common/thingspect_datapoint.proto\"F\n\x07Payload\x12,\n\x06points\x18\x01 \x03(\x0b\x32\x1c.thingspect.common.DataPoint\x12\r\n\x05token\x18\x02 \x01(\tB%Z#github.com/thingspect/proto/go/mqttb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'mqtt.thingspect_payload_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z#github.com/thingspect/proto/go/mqtt'
  _globals['_PAYLOAD']._serialized_start=85
  _globals['_PAYLOAD']._serialized_end=155
# @@protoc_insertion_point(module_scope)
