<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/thingspect_status.proto

namespace GPBMetadata\Api;

class ThingspectStatus
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
api/thingspect_status.protothingspect.api*:
Status
STATUS_UNSPECIFIED 

ACTIVE
DISABLEDB$Z"github.com/thingspect/proto/go/apibproto3'
        , true);

        static::$is_initialized = true;
    }
}

