# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/tag.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rapi/tag.proto\x12\x0ethingspect.api\x1a\x1cgoogle/api/annotations.proto\"\x11\n\x0fListTagsRequest\" \n\x10ListTagsResponse\x12\x0c\n\x04tags\x18\x01 \x03(\t2m\n\nTagService\x12_\n\x08ListTags\x12\x1f.thingspect.api.ListTagsRequest\x1a .thingspect.api.ListTagsResponse\"\x10\x82\xd3\xe4\x93\x02\n\x12\x08/v1/tagsB\"Z github.com/thingspect/api/go/apib\x06proto3')



_LISTTAGSREQUEST = DESCRIPTOR.message_types_by_name['ListTagsRequest']
_LISTTAGSRESPONSE = DESCRIPTOR.message_types_by_name['ListTagsResponse']
ListTagsRequest = _reflection.GeneratedProtocolMessageType('ListTagsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTTAGSREQUEST,
  '__module__' : 'api.tag_pb2'
  # @@protoc_insertion_point(class_scope:thingspect.api.ListTagsRequest)
  })
_sym_db.RegisterMessage(ListTagsRequest)

ListTagsResponse = _reflection.GeneratedProtocolMessageType('ListTagsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTTAGSRESPONSE,
  '__module__' : 'api.tag_pb2'
  # @@protoc_insertion_point(class_scope:thingspect.api.ListTagsResponse)
  })
_sym_db.RegisterMessage(ListTagsResponse)

_TAGSERVICE = DESCRIPTOR.services_by_name['TagService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z github.com/thingspect/api/go/api'
  _TAGSERVICE.methods_by_name['ListTags']._options = None
  _TAGSERVICE.methods_by_name['ListTags']._serialized_options = b'\202\323\344\223\002\n\022\010/v1/tags'
  _LISTTAGSREQUEST._serialized_start=63
  _LISTTAGSREQUEST._serialized_end=80
  _LISTTAGSRESPONSE._serialized_start=82
  _LISTTAGSRESPONSE._serialized_end=114
  _TAGSERVICE._serialized_start=116
  _TAGSERVICE._serialized_end=225
# @@protoc_insertion_point(module_scope)
