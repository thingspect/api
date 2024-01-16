// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mqtt/thingspect_payload.proto

#include "mqtt/thingspect_payload.pb.h"

#include <algorithm>
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/extension_set.h"
#include "google/protobuf/wire_format_lite.h"
#include "google/protobuf/descriptor.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/reflection_ops.h"
#include "google/protobuf/wire_format.h"
#include "google/protobuf/generated_message_tctable_impl.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"
PROTOBUF_PRAGMA_INIT_SEG
namespace _pb = ::google::protobuf;
namespace _pbi = ::google::protobuf::internal;
namespace _fl = ::google::protobuf::internal::field_layout;
namespace thingspect {
namespace mqtt {
        template <typename>
PROTOBUF_CONSTEXPR Payload::Payload(::_pbi::ConstantInitialized)
    : _impl_{
      /*decltype(_impl_.points_)*/ {},
      /*decltype(_impl_.token_)*/ {
          &::_pbi::fixed_address_empty_string,
          ::_pbi::ConstantInitialized{},
      },
      /*decltype(_impl_._cached_size_)*/ {},
    } {}
struct PayloadDefaultTypeInternal {
  PROTOBUF_CONSTEXPR PayloadDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~PayloadDefaultTypeInternal() {}
  union {
    Payload _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 PayloadDefaultTypeInternal _Payload_default_instance_;
}  // namespace mqtt
}  // namespace thingspect
static ::_pb::Metadata file_level_metadata_mqtt_2fthingspect_5fpayload_2eproto[1];
static constexpr const ::_pb::EnumDescriptor**
    file_level_enum_descriptors_mqtt_2fthingspect_5fpayload_2eproto = nullptr;
static constexpr const ::_pb::ServiceDescriptor**
    file_level_service_descriptors_mqtt_2fthingspect_5fpayload_2eproto = nullptr;
const ::uint32_t TableStruct_mqtt_2fthingspect_5fpayload_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(
    protodesc_cold) = {
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::thingspect::mqtt::Payload, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::thingspect::mqtt::Payload, _impl_.points_),
    PROTOBUF_FIELD_OFFSET(::thingspect::mqtt::Payload, _impl_.token_),
};

static const ::_pbi::MigrationSchema
    schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
        {0, -1, -1, sizeof(::thingspect::mqtt::Payload)},
};

static const ::_pb::Message* const file_default_instances[] = {
    &::thingspect::mqtt::_Payload_default_instance_._instance,
};
const char descriptor_table_protodef_mqtt_2fthingspect_5fpayload_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
    "\n\035mqtt/thingspect_payload.proto\022\017thingsp"
    "ect.mqtt\032!common/thingspect_datapoint.pr"
    "oto\"F\n\007Payload\022,\n\006points\030\001 \003(\0132\034.thingsp"
    "ect.common.DataPoint\022\r\n\005token\030\002 \001(\tB%Z#g"
    "ithub.com/thingspect/proto/go/mqttb\006prot"
    "o3"
};
static const ::_pbi::DescriptorTable* const descriptor_table_mqtt_2fthingspect_5fpayload_2eproto_deps[1] =
    {
        &::descriptor_table_common_2fthingspect_5fdatapoint_2eproto,
};
static ::absl::once_flag descriptor_table_mqtt_2fthingspect_5fpayload_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_mqtt_2fthingspect_5fpayload_2eproto = {
    false,
    false,
    202,
    descriptor_table_protodef_mqtt_2fthingspect_5fpayload_2eproto,
    "mqtt/thingspect_payload.proto",
    &descriptor_table_mqtt_2fthingspect_5fpayload_2eproto_once,
    descriptor_table_mqtt_2fthingspect_5fpayload_2eproto_deps,
    1,
    1,
    schemas,
    file_default_instances,
    TableStruct_mqtt_2fthingspect_5fpayload_2eproto::offsets,
    file_level_metadata_mqtt_2fthingspect_5fpayload_2eproto,
    file_level_enum_descriptors_mqtt_2fthingspect_5fpayload_2eproto,
    file_level_service_descriptors_mqtt_2fthingspect_5fpayload_2eproto,
};

// This function exists to be marked as weak.
// It can significantly speed up compilation by breaking up LLVM's SCC
// in the .pb.cc translation units. Large translation units see a
// reduction of more than 35% of walltime for optimized builds. Without
// the weak attribute all the messages in the file, including all the
// vtables and everything they use become part of the same SCC through
// a cycle like:
// GetMetadata -> descriptor table -> default instances ->
//   vtables -> GetMetadata
// By adding a weak function here we break the connection from the
// individual vtables back into the descriptor table.
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_mqtt_2fthingspect_5fpayload_2eproto_getter() {
  return &descriptor_table_mqtt_2fthingspect_5fpayload_2eproto;
}
// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2
static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_mqtt_2fthingspect_5fpayload_2eproto(&descriptor_table_mqtt_2fthingspect_5fpayload_2eproto);
namespace thingspect {
namespace mqtt {
// ===================================================================

class Payload::_Internal {
 public:
};

void Payload::clear_points() {
  _internal_mutable_points()->Clear();
}
Payload::Payload(::google::protobuf::Arena* arena)
    : ::google::protobuf::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:thingspect.mqtt.Payload)
}
Payload::Payload(const Payload& from) : ::google::protobuf::Message() {
  Payload* const _this = this;
  (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.points_){from._impl_.points_},
      decltype(_impl_.token_){},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);
  _impl_.token_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.token_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_token().empty()) {
    _this->_impl_.token_.Set(from._internal_token(), _this->GetArenaForAllocation());
  }

  // @@protoc_insertion_point(copy_constructor:thingspect.mqtt.Payload)
}
inline void Payload::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.points_){arena},
      decltype(_impl_.token_){},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _impl_.token_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.token_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}
Payload::~Payload() {
  // @@protoc_insertion_point(destructor:thingspect.mqtt.Payload)
  _internal_metadata_.Delete<::google::protobuf::UnknownFieldSet>();
  SharedDtor();
}
inline void Payload::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.points_.~RepeatedPtrField();
  _impl_.token_.Destroy();
}
void Payload::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

PROTOBUF_NOINLINE void Payload::Clear() {
// @@protoc_insertion_point(message_clear_start:thingspect.mqtt.Payload)
  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _internal_mutable_points()->Clear();
  _impl_.token_.ClearToEmpty();
  _internal_metadata_.Clear<::google::protobuf::UnknownFieldSet>();
}

const char* Payload::_InternalParse(
    const char* ptr, ::_pbi::ParseContext* ctx) {
  ptr = ::_pbi::TcParser::ParseLoop(this, ptr, ctx, &_table_.header);
  return ptr;
}


PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1
const ::_pbi::TcParseTable<1, 2, 1, 37, 2> Payload::_table_ = {
  {
    0,  // no _has_bits_
    0, // no _extensions_
    2, 8,  // max_field_number, fast_idx_mask
    offsetof(decltype(_table_), field_lookup_table),
    4294967292,  // skipmap
    offsetof(decltype(_table_), field_entries),
    2,  // num_field_entries
    1,  // num_aux_entries
    offsetof(decltype(_table_), aux_entries),
    &_Payload_default_instance_._instance,
    ::_pbi::TcParser::GenericFallback,  // fallback
  }, {{
    // string token = 2;
    {::_pbi::TcParser::FastUS1,
     {18, 63, 0, PROTOBUF_FIELD_OFFSET(Payload, _impl_.token_)}},
    // repeated .thingspect.common.DataPoint points = 1;
    {::_pbi::TcParser::FastMtR1,
     {10, 63, 0, PROTOBUF_FIELD_OFFSET(Payload, _impl_.points_)}},
  }}, {{
    65535, 65535
  }}, {{
    // repeated .thingspect.common.DataPoint points = 1;
    {PROTOBUF_FIELD_OFFSET(Payload, _impl_.points_), 0, 0,
    (0 | ::_fl::kFcRepeated | ::_fl::kMessage | ::_fl::kTvTable)},
    // string token = 2;
    {PROTOBUF_FIELD_OFFSET(Payload, _impl_.token_), 0, 0,
    (0 | ::_fl::kFcSingular | ::_fl::kUtf8String | ::_fl::kRepAString)},
  }}, {{
    {::_pbi::TcParser::GetTable<::thingspect::common::DataPoint>()},
  }}, {{
    "\27\0\5\0\0\0\0\0"
    "thingspect.mqtt.Payload"
    "token"
  }},
};

::uint8_t* Payload::_InternalSerialize(
    ::uint8_t* target,
    ::google::protobuf::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:thingspect.mqtt.Payload)
  ::uint32_t cached_has_bits = 0;
  (void)cached_has_bits;

  // repeated .thingspect.common.DataPoint points = 1;
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_points_size()); i < n; i++) {
    const auto& repfield = this->_internal_points().Get(i);
    target = ::google::protobuf::internal::WireFormatLite::
        InternalWriteMessage(1, repfield, repfield.GetCachedSize(), target, stream);
  }

  // string token = 2;
  if (!this->_internal_token().empty()) {
    const std::string& _s = this->_internal_token();
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
        _s.data(), static_cast<int>(_s.length()), ::google::protobuf::internal::WireFormatLite::SERIALIZE, "thingspect.mqtt.Payload.token");
    target = stream->WriteStringMaybeAliased(2, _s, target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target =
        ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
            _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:thingspect.mqtt.Payload)
  return target;
}

::size_t Payload::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:thingspect.mqtt.Payload)
  ::size_t total_size = 0;

  ::uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .thingspect.common.DataPoint points = 1;
  total_size += 1UL * this->_internal_points_size();
  for (const auto& msg : this->_internal_points()) {
    total_size +=
      ::google::protobuf::internal::WireFormatLite::MessageSize(msg);
  }
  // string token = 2;
  if (!this->_internal_token().empty()) {
    total_size += 1 + ::google::protobuf::internal::WireFormatLite::StringSize(
                                    this->_internal_token());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::google::protobuf::Message::ClassData Payload::_class_data_ = {
    ::google::protobuf::Message::CopyWithSourceCheck,
    Payload::MergeImpl
};
const ::google::protobuf::Message::ClassData*Payload::GetClassData() const { return &_class_data_; }


void Payload::MergeImpl(::google::protobuf::Message& to_msg, const ::google::protobuf::Message& from_msg) {
  auto* const _this = static_cast<Payload*>(&to_msg);
  auto& from = static_cast<const Payload&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:thingspect.mqtt.Payload)
  ABSL_DCHECK_NE(&from, _this);
  ::uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_internal_mutable_points()->MergeFrom(from._internal_points());
  if (!from._internal_token().empty()) {
    _this->_internal_set_token(from._internal_token());
  }
  _this->_internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(from._internal_metadata_);
}

void Payload::CopyFrom(const Payload& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:thingspect.mqtt.Payload)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

PROTOBUF_NOINLINE bool Payload::IsInitialized() const {
  return true;
}

void Payload::InternalSwap(Payload* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  _impl_.points_.InternalSwap(&other->_impl_.points_);
  ::_pbi::ArenaStringPtr::InternalSwap(&_impl_.token_, lhs_arena,
                                       &other->_impl_.token_, rhs_arena);
}

::google::protobuf::Metadata Payload::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_mqtt_2fthingspect_5fpayload_2eproto_getter, &descriptor_table_mqtt_2fthingspect_5fpayload_2eproto_once,
      file_level_metadata_mqtt_2fthingspect_5fpayload_2eproto[0]);
}
// @@protoc_insertion_point(namespace_scope)
}  // namespace mqtt
}  // namespace thingspect
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google
// @@protoc_insertion_point(global_scope)
#include "google/protobuf/port_undef.inc"
