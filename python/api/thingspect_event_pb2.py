# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/thingspect_event.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1a\x61pi/thingspect_event.proto\x12\x0ethingspect.api\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x17validate/validate.proto\"\x92\x01\n\x05\x45vent\x12\x15\n\x06org_id\x18\x01 \x01(\tR\x05orgID\x12\x17\n\x07uniq_id\x18\x02 \x01(\tR\x06uniqID\x12\x17\n\x07rule_id\x18\x03 \x01(\tR\x06ruleID\x12.\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x10\n\x08trace_id\x18\x05 \x01(\t\"\xef\x01\n\x11ListEventsRequest\x12\x19\n\x07uniq_id\x18\x01 \x01(\tH\x00R\x06uniqID\x12*\n\tdevice_id\x18\x02 \x01(\tB\x0b\xfa\x42\x08r\x06\xb0\x01\x01\xd0\x01\x01H\x00R\x08\x64\x65viceID\x12$\n\x07rule_id\x18\x03 \x01(\tB\x0b\xfa\x42\x08r\x06\xb0\x01\x01\xd0\x01\x01R\x06ruleID\x12,\n\x08\x65nd_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nstart_time\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x0f\n\x08id_oneof\x12\x03\xf8\x42\x01\";\n\x12ListEventsResponse\x12%\n\x06\x65vents\x18\x01 \x03(\x0b\x32\x15.thingspect.api.Event\";\n\x13LatestEventsRequest\x12$\n\x07rule_id\x18\x01 \x01(\tB\x0b\xfa\x42\x08r\x06\xb0\x01\x01\xd0\x01\x01R\x06ruleID\"=\n\x14LatestEventsResponse\x12%\n\x06\x65vents\x18\x01 \x03(\x0b\x32\x15.thingspect.api.Event2\xed\x01\n\x0c\x45ventService\x12g\n\nListEvents\x12!.thingspect.api.ListEventsRequest\x1a\".thingspect.api.ListEventsResponse\"\x12\x82\xd3\xe4\x93\x02\x0c\x12\n/v1/events\x12t\n\x0cLatestEvents\x12#.thingspect.api.LatestEventsRequest\x1a$.thingspect.api.LatestEventsResponse\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/v1/events/latestB$Z\"github.com/thingspect/proto/go/apib\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'api.thingspect_event_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\"github.com/thingspect/proto/go/api'
  _globals['_LISTEVENTSREQUEST'].oneofs_by_name['id_oneof']._options = None
  _globals['_LISTEVENTSREQUEST'].oneofs_by_name['id_oneof']._serialized_options = b'\370B\001'
  _globals['_LISTEVENTSREQUEST'].fields_by_name['device_id']._options = None
  _globals['_LISTEVENTSREQUEST'].fields_by_name['device_id']._serialized_options = b'\372B\010r\006\260\001\001\320\001\001'
  _globals['_LISTEVENTSREQUEST'].fields_by_name['rule_id']._options = None
  _globals['_LISTEVENTSREQUEST'].fields_by_name['rule_id']._serialized_options = b'\372B\010r\006\260\001\001\320\001\001'
  _globals['_LATESTEVENTSREQUEST'].fields_by_name['rule_id']._options = None
  _globals['_LATESTEVENTSREQUEST'].fields_by_name['rule_id']._serialized_options = b'\372B\010r\006\260\001\001\320\001\001'
  _globals['_EVENTSERVICE'].methods_by_name['ListEvents']._options = None
  _globals['_EVENTSERVICE'].methods_by_name['ListEvents']._serialized_options = b'\202\323\344\223\002\014\022\n/v1/events'
  _globals['_EVENTSERVICE'].methods_by_name['LatestEvents']._options = None
  _globals['_EVENTSERVICE'].methods_by_name['LatestEvents']._serialized_options = b'\202\323\344\223\002\023\022\021/v1/events/latest'
  _globals['_EVENT']._serialized_start=135
  _globals['_EVENT']._serialized_end=281
  _globals['_LISTEVENTSREQUEST']._serialized_start=284
  _globals['_LISTEVENTSREQUEST']._serialized_end=523
  _globals['_LISTEVENTSRESPONSE']._serialized_start=525
  _globals['_LISTEVENTSRESPONSE']._serialized_end=584
  _globals['_LATESTEVENTSREQUEST']._serialized_start=586
  _globals['_LATESTEVENTSREQUEST']._serialized_end=645
  _globals['_LATESTEVENTSRESPONSE']._serialized_start=647
  _globals['_LATESTEVENTSRESPONSE']._serialized_end=708
  _globals['_EVENTSERVICE']._serialized_start=711
  _globals['_EVENTSERVICE']._serialized_end=948
# @@protoc_insertion_point(module_scope)
