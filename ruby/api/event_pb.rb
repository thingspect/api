# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/event.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
require 'google/api/annotations_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("api/event.proto", :syntax => :proto3) do
    add_message "thingspect.api.Event" do
      optional :org_id, :string, 1, json_name: "orgID"
      optional :uniq_id, :string, 2, json_name: "uniqID"
      optional :rule_id, :string, 3, json_name: "ruleID"
      optional :created_at, :message, 4, "google.protobuf.Timestamp"
      optional :trace_id, :string, 5
    end
    add_message "thingspect.api.ListEventsRequest" do
      optional :rule_id, :string, 3, json_name: "ruleID"
      optional :end_time, :message, 4, "google.protobuf.Timestamp"
      optional :start_time, :message, 5, "google.protobuf.Timestamp"
      oneof :id_oneof do
        optional :uniq_id, :string, 1, json_name: "uniqID"
        optional :device_id, :string, 2, json_name: "deviceID"
      end
    end
    add_message "thingspect.api.ListEventsResponse" do
      repeated :events, :message, 1, "thingspect.api.Event"
    end
    add_message "thingspect.api.LatestEventsRequest" do
      optional :rule_id, :string, 1, json_name: "ruleID"
    end
    add_message "thingspect.api.LatestEventsResponse" do
      repeated :events, :message, 1, "thingspect.api.Event"
    end
  end
end

module Thingspect
  module Api
    Event = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("thingspect.api.Event").msgclass
    ListEventsRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("thingspect.api.ListEventsRequest").msgclass
    ListEventsResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("thingspect.api.ListEventsResponse").msgclass
    LatestEventsRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("thingspect.api.LatestEventsRequest").msgclass
    LatestEventsResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("thingspect.api.LatestEventsResponse").msgclass
  end
end
