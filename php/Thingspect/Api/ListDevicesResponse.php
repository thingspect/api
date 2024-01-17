<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/thingspect_device.proto

namespace Thingspect\Api;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * ListDevicesResponse is sent in response to a device list.
 *
 * Generated from protobuf message <code>thingspect.api.ListDevicesResponse</code>
 */
class ListDevicesResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Device array, ordered by ascending created_at timestamp.
     *
     * Generated from protobuf field <code>repeated .thingspect.api.Device devices = 1;</code>
     */
    private $devices;
    /**
     * Pagination token used to retrieve the next page of results. Not returned for the last page.
     *
     * Generated from protobuf field <code>string next_page_token = 2;</code>
     */
    protected $next_page_token = '';
    /**
     * Total number of devices available.
     *
     * Generated from protobuf field <code>int32 total_size = 3;</code>
     */
    protected $total_size = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type array<\Thingspect\Api\Device>|\Google\Protobuf\Internal\RepeatedField $devices
     *           Device array, ordered by ascending created_at timestamp.
     *     @type string $next_page_token
     *           Pagination token used to retrieve the next page of results. Not returned for the last page.
     *     @type int $total_size
     *           Total number of devices available.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Api\ThingspectDevice::initOnce();
        parent::__construct($data);
    }

    /**
     * Device array, ordered by ascending created_at timestamp.
     *
     * Generated from protobuf field <code>repeated .thingspect.api.Device devices = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getDevices()
    {
        return $this->devices;
    }

    /**
     * Device array, ordered by ascending created_at timestamp.
     *
     * Generated from protobuf field <code>repeated .thingspect.api.Device devices = 1;</code>
     * @param array<\Thingspect\Api\Device>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setDevices($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Thingspect\Api\Device::class);
        $this->devices = $arr;

        return $this;
    }

    /**
     * Pagination token used to retrieve the next page of results. Not returned for the last page.
     *
     * Generated from protobuf field <code>string next_page_token = 2;</code>
     * @return string
     */
    public function getNextPageToken()
    {
        return $this->next_page_token;
    }

    /**
     * Pagination token used to retrieve the next page of results. Not returned for the last page.
     *
     * Generated from protobuf field <code>string next_page_token = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setNextPageToken($var)
    {
        GPBUtil::checkString($var, True);
        $this->next_page_token = $var;

        return $this;
    }

    /**
     * Total number of devices available.
     *
     * Generated from protobuf field <code>int32 total_size = 3;</code>
     * @return int
     */
    public function getTotalSize()
    {
        return $this->total_size;
    }

    /**
     * Total number of devices available.
     *
     * Generated from protobuf field <code>int32 total_size = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setTotalSize($var)
    {
        GPBUtil::checkInt32($var);
        $this->total_size = $var;

        return $this;
    }

}

