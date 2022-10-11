# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from api import thingspect_rule_alarm_pb2 as api_dot_thingspect__rule__alarm__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class RuleAlarmServiceStub(object):
    """RuleAlarmService contains functions to query and modify rules and alarms.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateRule = channel.unary_unary(
                '/thingspect.api.RuleAlarmService/CreateRule',
                request_serializer=api_dot_thingspect__rule__alarm__pb2.CreateRuleRequest.SerializeToString,
                response_deserializer=api_dot_thingspect__rule__alarm__pb2.Rule.FromString,
                )
        self.CreateAlarm = channel.unary_unary(
                '/thingspect.api.RuleAlarmService/CreateAlarm',
                request_serializer=api_dot_thingspect__rule__alarm__pb2.CreateAlarmRequest.SerializeToString,
                response_deserializer=api_dot_thingspect__rule__alarm__pb2.Alarm.FromString,
                )
        self.GetRule = channel.unary_unary(
                '/thingspect.api.RuleAlarmService/GetRule',
                request_serializer=api_dot_thingspect__rule__alarm__pb2.GetRuleRequest.SerializeToString,
                response_deserializer=api_dot_thingspect__rule__alarm__pb2.Rule.FromString,
                )
        self.GetAlarm = channel.unary_unary(
                '/thingspect.api.RuleAlarmService/GetAlarm',
                request_serializer=api_dot_thingspect__rule__alarm__pb2.GetAlarmRequest.SerializeToString,
                response_deserializer=api_dot_thingspect__rule__alarm__pb2.Alarm.FromString,
                )
        self.UpdateRule = channel.unary_unary(
                '/thingspect.api.RuleAlarmService/UpdateRule',
                request_serializer=api_dot_thingspect__rule__alarm__pb2.UpdateRuleRequest.SerializeToString,
                response_deserializer=api_dot_thingspect__rule__alarm__pb2.Rule.FromString,
                )
        self.UpdateAlarm = channel.unary_unary(
                '/thingspect.api.RuleAlarmService/UpdateAlarm',
                request_serializer=api_dot_thingspect__rule__alarm__pb2.UpdateAlarmRequest.SerializeToString,
                response_deserializer=api_dot_thingspect__rule__alarm__pb2.Alarm.FromString,
                )
        self.DeleteRule = channel.unary_unary(
                '/thingspect.api.RuleAlarmService/DeleteRule',
                request_serializer=api_dot_thingspect__rule__alarm__pb2.DeleteRuleRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.DeleteAlarm = channel.unary_unary(
                '/thingspect.api.RuleAlarmService/DeleteAlarm',
                request_serializer=api_dot_thingspect__rule__alarm__pb2.DeleteAlarmRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.ListRules = channel.unary_unary(
                '/thingspect.api.RuleAlarmService/ListRules',
                request_serializer=api_dot_thingspect__rule__alarm__pb2.ListRulesRequest.SerializeToString,
                response_deserializer=api_dot_thingspect__rule__alarm__pb2.ListRulesResponse.FromString,
                )
        self.ListAlarms = channel.unary_unary(
                '/thingspect.api.RuleAlarmService/ListAlarms',
                request_serializer=api_dot_thingspect__rule__alarm__pb2.ListAlarmsRequest.SerializeToString,
                response_deserializer=api_dot_thingspect__rule__alarm__pb2.ListAlarmsResponse.FromString,
                )
        self.TestRule = channel.unary_unary(
                '/thingspect.api.RuleAlarmService/TestRule',
                request_serializer=api_dot_thingspect__rule__alarm__pb2.TestRuleRequest.SerializeToString,
                response_deserializer=api_dot_thingspect__rule__alarm__pb2.TestRuleResponse.FromString,
                )
        self.TestAlarm = channel.unary_unary(
                '/thingspect.api.RuleAlarmService/TestAlarm',
                request_serializer=api_dot_thingspect__rule__alarm__pb2.TestAlarmRequest.SerializeToString,
                response_deserializer=api_dot_thingspect__rule__alarm__pb2.TestAlarmResponse.FromString,
                )


class RuleAlarmServiceServicer(object):
    """RuleAlarmService contains functions to query and modify rules and alarms.
    """

    def CreateRule(self, request, context):
        """Create a rule. Rules generate events when conditions are met.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateAlarm(self, request, context):
        """Create an alarm. Alarms generate alerts via parent rules.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetRule(self, request, context):
        """Get a rule by ID. Rules generate events when conditions are met.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAlarm(self, request, context):
        """Get an alarm by ID. Alarms generate alerts via parent rules.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateRule(self, request, context):
        """Update a rule. Rules generate events when conditions are met.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateAlarm(self, request, context):
        """Update an alarm. Alarms generate alerts via parent rules.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteRule(self, request, context):
        """Delete a rule by ID. Rules generate events when conditions are met.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteAlarm(self, request, context):
        """Delete an alarm by ID. Alarms generate alerts via parent rules.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListRules(self, request, context):
        """List all rules. Rules generate events when conditions are met.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListAlarms(self, request, context):
        """List alarms. Alarms generate alerts via parent rules.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TestRule(self, request, context):
        """Test a rule. Rules generate events when conditions are met.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TestAlarm(self, request, context):
        """Test an alarm. Alarms generate alerts via parent rules.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_RuleAlarmServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateRule': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateRule,
                    request_deserializer=api_dot_thingspect__rule__alarm__pb2.CreateRuleRequest.FromString,
                    response_serializer=api_dot_thingspect__rule__alarm__pb2.Rule.SerializeToString,
            ),
            'CreateAlarm': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateAlarm,
                    request_deserializer=api_dot_thingspect__rule__alarm__pb2.CreateAlarmRequest.FromString,
                    response_serializer=api_dot_thingspect__rule__alarm__pb2.Alarm.SerializeToString,
            ),
            'GetRule': grpc.unary_unary_rpc_method_handler(
                    servicer.GetRule,
                    request_deserializer=api_dot_thingspect__rule__alarm__pb2.GetRuleRequest.FromString,
                    response_serializer=api_dot_thingspect__rule__alarm__pb2.Rule.SerializeToString,
            ),
            'GetAlarm': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAlarm,
                    request_deserializer=api_dot_thingspect__rule__alarm__pb2.GetAlarmRequest.FromString,
                    response_serializer=api_dot_thingspect__rule__alarm__pb2.Alarm.SerializeToString,
            ),
            'UpdateRule': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateRule,
                    request_deserializer=api_dot_thingspect__rule__alarm__pb2.UpdateRuleRequest.FromString,
                    response_serializer=api_dot_thingspect__rule__alarm__pb2.Rule.SerializeToString,
            ),
            'UpdateAlarm': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateAlarm,
                    request_deserializer=api_dot_thingspect__rule__alarm__pb2.UpdateAlarmRequest.FromString,
                    response_serializer=api_dot_thingspect__rule__alarm__pb2.Alarm.SerializeToString,
            ),
            'DeleteRule': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteRule,
                    request_deserializer=api_dot_thingspect__rule__alarm__pb2.DeleteRuleRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'DeleteAlarm': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteAlarm,
                    request_deserializer=api_dot_thingspect__rule__alarm__pb2.DeleteAlarmRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'ListRules': grpc.unary_unary_rpc_method_handler(
                    servicer.ListRules,
                    request_deserializer=api_dot_thingspect__rule__alarm__pb2.ListRulesRequest.FromString,
                    response_serializer=api_dot_thingspect__rule__alarm__pb2.ListRulesResponse.SerializeToString,
            ),
            'ListAlarms': grpc.unary_unary_rpc_method_handler(
                    servicer.ListAlarms,
                    request_deserializer=api_dot_thingspect__rule__alarm__pb2.ListAlarmsRequest.FromString,
                    response_serializer=api_dot_thingspect__rule__alarm__pb2.ListAlarmsResponse.SerializeToString,
            ),
            'TestRule': grpc.unary_unary_rpc_method_handler(
                    servicer.TestRule,
                    request_deserializer=api_dot_thingspect__rule__alarm__pb2.TestRuleRequest.FromString,
                    response_serializer=api_dot_thingspect__rule__alarm__pb2.TestRuleResponse.SerializeToString,
            ),
            'TestAlarm': grpc.unary_unary_rpc_method_handler(
                    servicer.TestAlarm,
                    request_deserializer=api_dot_thingspect__rule__alarm__pb2.TestAlarmRequest.FromString,
                    response_serializer=api_dot_thingspect__rule__alarm__pb2.TestAlarmResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'thingspect.api.RuleAlarmService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class RuleAlarmService(object):
    """RuleAlarmService contains functions to query and modify rules and alarms.
    """

    @staticmethod
    def CreateRule(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.RuleAlarmService/CreateRule',
            api_dot_thingspect__rule__alarm__pb2.CreateRuleRequest.SerializeToString,
            api_dot_thingspect__rule__alarm__pb2.Rule.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateAlarm(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.RuleAlarmService/CreateAlarm',
            api_dot_thingspect__rule__alarm__pb2.CreateAlarmRequest.SerializeToString,
            api_dot_thingspect__rule__alarm__pb2.Alarm.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetRule(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.RuleAlarmService/GetRule',
            api_dot_thingspect__rule__alarm__pb2.GetRuleRequest.SerializeToString,
            api_dot_thingspect__rule__alarm__pb2.Rule.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAlarm(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.RuleAlarmService/GetAlarm',
            api_dot_thingspect__rule__alarm__pb2.GetAlarmRequest.SerializeToString,
            api_dot_thingspect__rule__alarm__pb2.Alarm.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateRule(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.RuleAlarmService/UpdateRule',
            api_dot_thingspect__rule__alarm__pb2.UpdateRuleRequest.SerializeToString,
            api_dot_thingspect__rule__alarm__pb2.Rule.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateAlarm(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.RuleAlarmService/UpdateAlarm',
            api_dot_thingspect__rule__alarm__pb2.UpdateAlarmRequest.SerializeToString,
            api_dot_thingspect__rule__alarm__pb2.Alarm.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteRule(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.RuleAlarmService/DeleteRule',
            api_dot_thingspect__rule__alarm__pb2.DeleteRuleRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteAlarm(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.RuleAlarmService/DeleteAlarm',
            api_dot_thingspect__rule__alarm__pb2.DeleteAlarmRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListRules(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.RuleAlarmService/ListRules',
            api_dot_thingspect__rule__alarm__pb2.ListRulesRequest.SerializeToString,
            api_dot_thingspect__rule__alarm__pb2.ListRulesResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListAlarms(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.RuleAlarmService/ListAlarms',
            api_dot_thingspect__rule__alarm__pb2.ListAlarmsRequest.SerializeToString,
            api_dot_thingspect__rule__alarm__pb2.ListAlarmsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def TestRule(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.RuleAlarmService/TestRule',
            api_dot_thingspect__rule__alarm__pb2.TestRuleRequest.SerializeToString,
            api_dot_thingspect__rule__alarm__pb2.TestRuleResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def TestAlarm(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/thingspect.api.RuleAlarmService/TestAlarm',
            api_dot_thingspect__rule__alarm__pb2.TestAlarmRequest.SerializeToString,
            api_dot_thingspect__rule__alarm__pb2.TestAlarmResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
