# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from api import tag_pb2 as api_dot_tag__pb2


class TagServiceStub(object):
    """TagService contains functions to query tags.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListTags = channel.unary_unary(
                '/thingspect.api.TagService/ListTags',
                request_serializer=api_dot_tag__pb2.ListTagsRequest.SerializeToString,
                response_deserializer=api_dot_tag__pb2.ListTagsResponse.FromString,
                )


class TagServiceServicer(object):
    """TagService contains functions to query tags.
    """

    def ListTags(self, request, context):
        """List all tags.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TagServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListTags': grpc.unary_unary_rpc_method_handler(
                    servicer.ListTags,
                    request_deserializer=api_dot_tag__pb2.ListTagsRequest.FromString,
                    response_serializer=api_dot_tag__pb2.ListTagsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'thingspect.api.TagService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TagService(object):
    """TagService contains functions to query tags.
    """

    @staticmethod
    def ListTags(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.TagService/ListTags',
            api_dot_tag__pb2.ListTagsRequest.SerializeToString,
            api_dot_tag__pb2.ListTagsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
