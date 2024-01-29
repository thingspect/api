//
//  Generated code. Do not modify.
//  source: api/thingspect_device.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

/// Decoder represents the data payload decoder for a device.
class Decoder extends $pb.ProtobufEnum {
  static const Decoder RAW = Decoder._(0, _omitEnumNames ? '' : 'RAW');
  static const Decoder GATEWAY = Decoder._(1, _omitEnumNames ? '' : 'GATEWAY');
  static const Decoder RADIO_BRIDGE_DOOR_V1 = Decoder._(2, _omitEnumNames ? '' : 'RADIO_BRIDGE_DOOR_V1');
  static const Decoder RADIO_BRIDGE_DOOR_V2 = Decoder._(3, _omitEnumNames ? '' : 'RADIO_BRIDGE_DOOR_V2');
  static const Decoder GLOBALSAT_CO2 = Decoder._(4, _omitEnumNames ? '' : 'GLOBALSAT_CO2');
  static const Decoder GLOBALSAT_CO = Decoder._(5, _omitEnumNames ? '' : 'GLOBALSAT_CO');
  static const Decoder GLOBALSAT_PM25 = Decoder._(6, _omitEnumNames ? '' : 'GLOBALSAT_PM25');
  static const Decoder TEKTELIC_HOME = Decoder._(7, _omitEnumNames ? '' : 'TEKTELIC_HOME');

  static const $core.List<Decoder> values = <Decoder> [
    RAW,
    GATEWAY,
    RADIO_BRIDGE_DOOR_V1,
    RADIO_BRIDGE_DOOR_V2,
    GLOBALSAT_CO2,
    GLOBALSAT_CO,
    GLOBALSAT_PM25,
    TEKTELIC_HOME,
  ];

  static final $core.Map<$core.int, Decoder> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Decoder? valueOf($core.int value) => _byValue[value];

  const Decoder._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
