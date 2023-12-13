# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/thingspect_alert.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1a\x61pi/thingspect_alert.proto\x12\x0ethingspect.api\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x17validate/validate.proto\"\xe9\x01\n\x05\x41lert\x12\x15\n\x06org_id\x18\x01 \x01(\tR\x05orgID\x12\x17\n\x07uniq_id\x18\x02 \x01(\tR\x06uniqID\x12\x19\n\x08\x61larm_id\x18\x03 \x01(\tR\x07\x61larmID\x12\x17\n\x07user_id\x18\x04 \x01(\tR\x06userID\x12+\n\x06status\x18\x05 \x01(\x0e\x32\x1b.thingspect.api.AlertStatus\x12\r\n\x05\x65rror\x18\x06 \x01(\t\x12.\n\ncreated_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x10\n\x08trace_id\x18\x08 \x01(\t\"\x92\x02\n\x11ListAlertsRequest\x12\x19\n\x07uniq_id\x18\x01 \x01(\tH\x00R\x06uniqID\x12*\n\tdevice_id\x18\x02 \x01(\tB\x0b\xfa\x42\x08r\x06\xb0\x01\x01\xd0\x01\x01H\x00R\x08\x64\x65viceID\x12&\n\x08\x61larm_id\x18\x03 \x01(\tB\x0b\xfa\x42\x08r\x06\xb0\x01\x01\xd0\x01\x01R\x07\x61larmID\x12$\n\x07user_id\x18\x04 \x01(\tB\x0b\xfa\x42\x08r\x06\xb0\x01\x01\xd0\x01\x01R\x06userID\x12,\n\x08\x65nd_time\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nstart_time\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\n\n\x08id_oneof\";\n\x12ListAlertsResponse\x12%\n\x06\x61lerts\x18\x01 \x03(\x0b\x32\x15.thingspect.api.Alert*@\n\x0b\x41lertStatus\x12\x1c\n\x18\x41LERT_STATUS_UNSPECIFIED\x10\x00\x12\x08\n\x04SENT\x10\x01\x12\t\n\x05\x45RROR\x10\x02\x32w\n\x0c\x41lertService\x12g\n\nListAlerts\x12!.thingspect.api.ListAlertsRequest\x1a\".thingspect.api.ListAlertsResponse\"\x12\x82\xd3\xe4\x93\x02\x0c\x12\n/v1/alertsB\"Z github.com/thingspect/api/go/apib\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'api.thingspect_alert_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z github.com/thingspect/api/go/api'
  _globals['_LISTALERTSREQUEST'].fields_by_name['device_id']._options = None
  _globals['_LISTALERTSREQUEST'].fields_by_name['device_id']._serialized_options = b'\372B\010r\006\260\001\001\320\001\001'
  _globals['_LISTALERTSREQUEST'].fields_by_name['alarm_id']._options = None
  _globals['_LISTALERTSREQUEST'].fields_by_name['alarm_id']._serialized_options = b'\372B\010r\006\260\001\001\320\001\001'
  _globals['_LISTALERTSREQUEST'].fields_by_name['user_id']._options = None
  _globals['_LISTALERTSREQUEST'].fields_by_name['user_id']._serialized_options = b'\372B\010r\006\260\001\001\320\001\001'
  _globals['_ALERTSERVICE'].methods_by_name['ListAlerts']._options = None
  _globals['_ALERTSERVICE'].methods_by_name['ListAlerts']._serialized_options = b'\202\323\344\223\002\014\022\n/v1/alerts'
  _globals['_ALERTSTATUS']._serialized_start=708
  _globals['_ALERTSTATUS']._serialized_end=772
  _globals['_ALERT']._serialized_start=135
  _globals['_ALERT']._serialized_end=368
  _globals['_LISTALERTSREQUEST']._serialized_start=371
  _globals['_LISTALERTSREQUEST']._serialized_end=645
  _globals['_LISTALERTSRESPONSE']._serialized_start=647
  _globals['_LISTALERTSRESPONSE']._serialized_end=706
  _globals['_ALERTSERVICE']._serialized_start=774
  _globals['_ALERTSERVICE']._serialized_end=893
# @@protoc_insertion_point(module_scope)
