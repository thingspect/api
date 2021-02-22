# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from api import org_pb2 as api_dot_org__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class OrgServiceStub(object):
    """OrgService contains functions to query and modify organizations.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateOrg = channel.unary_unary(
                '/thingspect.api.OrgService/CreateOrg',
                request_serializer=api_dot_org__pb2.CreateOrgRequest.SerializeToString,
                response_deserializer=api_dot_org__pb2.Org.FromString,
                )
        self.GetOrg = channel.unary_unary(
                '/thingspect.api.OrgService/GetOrg',
                request_serializer=api_dot_org__pb2.GetOrgRequest.SerializeToString,
                response_deserializer=api_dot_org__pb2.Org.FromString,
                )
        self.UpdateOrg = channel.unary_unary(
                '/thingspect.api.OrgService/UpdateOrg',
                request_serializer=api_dot_org__pb2.UpdateOrgRequest.SerializeToString,
                response_deserializer=api_dot_org__pb2.Org.FromString,
                )
        self.DeleteOrg = channel.unary_unary(
                '/thingspect.api.OrgService/DeleteOrg',
                request_serializer=api_dot_org__pb2.DeleteOrgRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.ListOrgs = channel.unary_unary(
                '/thingspect.api.OrgService/ListOrgs',
                request_serializer=api_dot_org__pb2.ListOrgsRequest.SerializeToString,
                response_deserializer=api_dot_org__pb2.ListOrgsResponse.FromString,
                )


class OrgServiceServicer(object):
    """OrgService contains functions to query and modify organizations.
    """

    def CreateOrg(self, request, context):
        """Create an organization.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetOrg(self, request, context):
        """Get an organization by ID.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateOrg(self, request, context):
        """Update an organization.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteOrg(self, request, context):
        """Delete an organization by ID.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListOrgs(self, request, context):
        """List all orgs.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OrgServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateOrg': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateOrg,
                    request_deserializer=api_dot_org__pb2.CreateOrgRequest.FromString,
                    response_serializer=api_dot_org__pb2.Org.SerializeToString,
            ),
            'GetOrg': grpc.unary_unary_rpc_method_handler(
                    servicer.GetOrg,
                    request_deserializer=api_dot_org__pb2.GetOrgRequest.FromString,
                    response_serializer=api_dot_org__pb2.Org.SerializeToString,
            ),
            'UpdateOrg': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateOrg,
                    request_deserializer=api_dot_org__pb2.UpdateOrgRequest.FromString,
                    response_serializer=api_dot_org__pb2.Org.SerializeToString,
            ),
            'DeleteOrg': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteOrg,
                    request_deserializer=api_dot_org__pb2.DeleteOrgRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'ListOrgs': grpc.unary_unary_rpc_method_handler(
                    servicer.ListOrgs,
                    request_deserializer=api_dot_org__pb2.ListOrgsRequest.FromString,
                    response_serializer=api_dot_org__pb2.ListOrgsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'thingspect.api.OrgService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class OrgService(object):
    """OrgService contains functions to query and modify organizations.
    """

    @staticmethod
    def CreateOrg(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.OrgService/CreateOrg',
            api_dot_org__pb2.CreateOrgRequest.SerializeToString,
            api_dot_org__pb2.Org.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetOrg(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.OrgService/GetOrg',
            api_dot_org__pb2.GetOrgRequest.SerializeToString,
            api_dot_org__pb2.Org.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateOrg(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.OrgService/UpdateOrg',
            api_dot_org__pb2.UpdateOrgRequest.SerializeToString,
            api_dot_org__pb2.Org.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteOrg(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.OrgService/DeleteOrg',
            api_dot_org__pb2.DeleteOrgRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListOrgs(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.OrgService/ListOrgs',
            api_dot_org__pb2.ListOrgsRequest.SerializeToString,
            api_dot_org__pb2.ListOrgsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
