<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/thingspect_session.proto

namespace Thingspect\Api;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * DeleteKeyRequest is sent to delete an API key.
 *
 * Generated from protobuf message <code>thingspect.api.DeleteKeyRequest</code>
 */
class DeleteKeyRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Key ID (UUID) to delete.
     *
     * Generated from protobuf field <code>string id = 1 [(.google.api.field_behavior) = REQUIRED, (.validate.rules) = {</code>
     */
    protected $id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *           Key ID (UUID) to delete.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Api\ThingspectSession::initOnce();
        parent::__construct($data);
    }

    /**
     * Key ID (UUID) to delete.
     *
     * Generated from protobuf field <code>string id = 1 [(.google.api.field_behavior) = REQUIRED, (.validate.rules) = {</code>
     * @return string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Key ID (UUID) to delete.
     *
     * Generated from protobuf field <code>string id = 1 [(.google.api.field_behavior) = REQUIRED, (.validate.rules) = {</code>
     * @param string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkString($var, True);
        $this->id = $var;

        return $this;
    }

}

