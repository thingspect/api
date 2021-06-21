# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: api/session.proto for package 'thingspect.api'

require 'grpc'
require 'api/session_pb'

module Thingspect
  module Api
    module SessionService
      # SessionService contains functions to create sessions and keys.
      class Service

        include ::GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'thingspect.api.SessionService'

        # Log in a user. Login tokens are time-limited and accompanied by an expiration.
        rpc :Login, ::Thingspect::Api::LoginRequest, ::Thingspect::Api::LoginResponse
        # Create an API key. API keys are persistent and do not expire until revoked.
        rpc :CreateKey, ::Thingspect::Api::CreateKeyRequest, ::Thingspect::Api::CreateKeyResponse
        # Delete an API key by ID. API keys are persistent and do not expire until revoked.
        rpc :DeleteKey, ::Thingspect::Api::DeleteKeyRequest, ::Google::Protobuf::Empty
        # List all API keys. API keys are persistent and do not expire until revoked.
        rpc :ListKeys, ::Thingspect::Api::ListKeysRequest, ::Thingspect::Api::ListKeysResponse
      end

      Stub = Service.rpc_stub_class
    end
  end
end
