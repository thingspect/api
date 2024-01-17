<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/thingspect_user.proto

namespace Thingspect\Api;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * ListUsersResponse is sent in response to a user list.
 *
 * Generated from protobuf message <code>thingspect.api.ListUsersResponse</code>
 */
class ListUsersResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * User array, ordered by ascending created_at timestamp. The completeness of the array will depend on whether the request was made by an admin user, non-admin user, or API key.
     *
     * Generated from protobuf field <code>repeated .thingspect.api.User users = 1;</code>
     */
    private $users;
    /**
     * Pagination token used to retrieve the next page of results. Not returned for the last page.
     *
     * Generated from protobuf field <code>string next_page_token = 2;</code>
     */
    protected $next_page_token = '';
    /**
     * Total number of users available.
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
     *     @type array<\Thingspect\Api\User>|\Google\Protobuf\Internal\RepeatedField $users
     *           User array, ordered by ascending created_at timestamp. The completeness of the array will depend on whether the request was made by an admin user, non-admin user, or API key.
     *     @type string $next_page_token
     *           Pagination token used to retrieve the next page of results. Not returned for the last page.
     *     @type int $total_size
     *           Total number of users available.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Api\ThingspectUser::initOnce();
        parent::__construct($data);
    }

    /**
     * User array, ordered by ascending created_at timestamp. The completeness of the array will depend on whether the request was made by an admin user, non-admin user, or API key.
     *
     * Generated from protobuf field <code>repeated .thingspect.api.User users = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getUsers()
    {
        return $this->users;
    }

    /**
     * User array, ordered by ascending created_at timestamp. The completeness of the array will depend on whether the request was made by an admin user, non-admin user, or API key.
     *
     * Generated from protobuf field <code>repeated .thingspect.api.User users = 1;</code>
     * @param array<\Thingspect\Api\User>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setUsers($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Thingspect\Api\User::class);
        $this->users = $arr;

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
     * Total number of users available.
     *
     * Generated from protobuf field <code>int32 total_size = 3;</code>
     * @return int
     */
    public function getTotalSize()
    {
        return $this->total_size;
    }

    /**
     * Total number of users available.
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

