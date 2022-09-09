# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'ExchangeSettings',
    'FederationUpstreamDefinition',
    'OperatorPolicyPolicy',
    'PermissionsPermissions',
    'PolicyPolicy',
    'QueueSettings',
    'ShovelInfo',
    'TopicPermissionsPermission',
    'GetExchangeSettingResult',
]

@pulumi.output_type
class ExchangeSettings(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "autoDelete":
            suggest = "auto_delete"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ExchangeSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ExchangeSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ExchangeSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 arguments: Optional[Mapping[str, Any]] = None,
                 auto_delete: Optional[bool] = None,
                 durable: Optional[bool] = None):
        """
        :param str type: The type of exchange.
        :param Mapping[str, Any] arguments: Additional key/value settings for the exchange.
        :param bool auto_delete: Whether the exchange will self-delete when all
               queues have finished using it.
        :param bool durable: Whether the exchange survives server restarts.
               Defaults to `false`.
        """
        pulumi.set(__self__, "type", type)
        if arguments is not None:
            pulumi.set(__self__, "arguments", arguments)
        if auto_delete is not None:
            pulumi.set(__self__, "auto_delete", auto_delete)
        if durable is not None:
            pulumi.set(__self__, "durable", durable)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of exchange.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def arguments(self) -> Optional[Mapping[str, Any]]:
        """
        Additional key/value settings for the exchange.
        """
        return pulumi.get(self, "arguments")

    @property
    @pulumi.getter(name="autoDelete")
    def auto_delete(self) -> Optional[bool]:
        """
        Whether the exchange will self-delete when all
        queues have finished using it.
        """
        return pulumi.get(self, "auto_delete")

    @property
    @pulumi.getter
    def durable(self) -> Optional[bool]:
        """
        Whether the exchange survives server restarts.
        Defaults to `false`.
        """
        return pulumi.get(self, "durable")


@pulumi.output_type
class FederationUpstreamDefinition(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ackMode":
            suggest = "ack_mode"
        elif key == "maxHops":
            suggest = "max_hops"
        elif key == "messageTtl":
            suggest = "message_ttl"
        elif key == "prefetchCount":
            suggest = "prefetch_count"
        elif key == "reconnectDelay":
            suggest = "reconnect_delay"
        elif key == "trustUserId":
            suggest = "trust_user_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FederationUpstreamDefinition. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FederationUpstreamDefinition.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FederationUpstreamDefinition.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 uri: str,
                 ack_mode: Optional[str] = None,
                 exchange: Optional[str] = None,
                 expires: Optional[int] = None,
                 max_hops: Optional[int] = None,
                 message_ttl: Optional[int] = None,
                 prefetch_count: Optional[int] = None,
                 queue: Optional[str] = None,
                 reconnect_delay: Optional[int] = None,
                 trust_user_id: Optional[bool] = None):
        """
        :param str uri: The AMQP URI(s) for the upstream. Note that the URI may contain sensitive information, such as a password.
        :param str ack_mode: Determines how the link should acknowledge messages. Valid values are `on-confirm`, `on-publish`, and `no-ack`. Default is `on-confirm`.
        :param str exchange: The name of the upstream exchange.
        :param int expires: The expiry time (in milliseconds) after which an upstream queue for a federated exchange may be deleted if a connection to the upstream is lost.
        :param int max_hops: Maximum number of federation links that messages can traverse before being dropped. Default is `1`.
        :param int message_ttl: The expiry time (in milliseconds) for messages in the upstream queue for a federated exchange (see expires).
        :param int prefetch_count: Maximum number of unacknowledged messages that may be in flight over a federation link at one time. Default is `1000`.
        :param str queue: The name of the upstream queue.
        :param int reconnect_delay: Time in seconds to wait after a network link goes down before attempting reconnection. Default is `5`.
        :param bool trust_user_id: Determines how federation should interact with the validated user-id feature. Default is `false`.
        """
        pulumi.set(__self__, "uri", uri)
        if ack_mode is not None:
            pulumi.set(__self__, "ack_mode", ack_mode)
        if exchange is not None:
            pulumi.set(__self__, "exchange", exchange)
        if expires is not None:
            pulumi.set(__self__, "expires", expires)
        if max_hops is not None:
            pulumi.set(__self__, "max_hops", max_hops)
        if message_ttl is not None:
            pulumi.set(__self__, "message_ttl", message_ttl)
        if prefetch_count is not None:
            pulumi.set(__self__, "prefetch_count", prefetch_count)
        if queue is not None:
            pulumi.set(__self__, "queue", queue)
        if reconnect_delay is not None:
            pulumi.set(__self__, "reconnect_delay", reconnect_delay)
        if trust_user_id is not None:
            pulumi.set(__self__, "trust_user_id", trust_user_id)

    @property
    @pulumi.getter
    def uri(self) -> str:
        """
        The AMQP URI(s) for the upstream. Note that the URI may contain sensitive information, such as a password.
        """
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter(name="ackMode")
    def ack_mode(self) -> Optional[str]:
        """
        Determines how the link should acknowledge messages. Valid values are `on-confirm`, `on-publish`, and `no-ack`. Default is `on-confirm`.
        """
        return pulumi.get(self, "ack_mode")

    @property
    @pulumi.getter
    def exchange(self) -> Optional[str]:
        """
        The name of the upstream exchange.
        """
        return pulumi.get(self, "exchange")

    @property
    @pulumi.getter
    def expires(self) -> Optional[int]:
        """
        The expiry time (in milliseconds) after which an upstream queue for a federated exchange may be deleted if a connection to the upstream is lost.
        """
        return pulumi.get(self, "expires")

    @property
    @pulumi.getter(name="maxHops")
    def max_hops(self) -> Optional[int]:
        """
        Maximum number of federation links that messages can traverse before being dropped. Default is `1`.
        """
        return pulumi.get(self, "max_hops")

    @property
    @pulumi.getter(name="messageTtl")
    def message_ttl(self) -> Optional[int]:
        """
        The expiry time (in milliseconds) for messages in the upstream queue for a federated exchange (see expires).
        """
        return pulumi.get(self, "message_ttl")

    @property
    @pulumi.getter(name="prefetchCount")
    def prefetch_count(self) -> Optional[int]:
        """
        Maximum number of unacknowledged messages that may be in flight over a federation link at one time. Default is `1000`.
        """
        return pulumi.get(self, "prefetch_count")

    @property
    @pulumi.getter
    def queue(self) -> Optional[str]:
        """
        The name of the upstream queue.
        """
        return pulumi.get(self, "queue")

    @property
    @pulumi.getter(name="reconnectDelay")
    def reconnect_delay(self) -> Optional[int]:
        """
        Time in seconds to wait after a network link goes down before attempting reconnection. Default is `5`.
        """
        return pulumi.get(self, "reconnect_delay")

    @property
    @pulumi.getter(name="trustUserId")
    def trust_user_id(self) -> Optional[bool]:
        """
        Determines how federation should interact with the validated user-id feature. Default is `false`.
        """
        return pulumi.get(self, "trust_user_id")


@pulumi.output_type
class OperatorPolicyPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "applyTo":
            suggest = "apply_to"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in OperatorPolicyPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        OperatorPolicyPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        OperatorPolicyPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 apply_to: str,
                 definition: Mapping[str, Any],
                 pattern: str,
                 priority: int):
        """
        :param str apply_to: Can be "queues".
        :param Mapping[str, Any] definition: Key/value pairs of the operator policy definition. See the
               RabbitMQ documentation for definition references and examples.
        :param str pattern: A pattern to match an exchange or queue name.
        :param int priority: The policy with the greater priority is applied first.
        """
        pulumi.set(__self__, "apply_to", apply_to)
        pulumi.set(__self__, "definition", definition)
        pulumi.set(__self__, "pattern", pattern)
        pulumi.set(__self__, "priority", priority)

    @property
    @pulumi.getter(name="applyTo")
    def apply_to(self) -> str:
        """
        Can be "queues".
        """
        return pulumi.get(self, "apply_to")

    @property
    @pulumi.getter
    def definition(self) -> Mapping[str, Any]:
        """
        Key/value pairs of the operator policy definition. See the
        RabbitMQ documentation for definition references and examples.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter
    def pattern(self) -> str:
        """
        A pattern to match an exchange or queue name.
        """
        return pulumi.get(self, "pattern")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        The policy with the greater priority is applied first.
        """
        return pulumi.get(self, "priority")


@pulumi.output_type
class PermissionsPermissions(dict):
    def __init__(__self__, *,
                 configure: str,
                 read: str,
                 write: str):
        """
        :param str configure: The "configure" ACL.
        :param str read: The "read" ACL.
        :param str write: The "write" ACL.
        """
        pulumi.set(__self__, "configure", configure)
        pulumi.set(__self__, "read", read)
        pulumi.set(__self__, "write", write)

    @property
    @pulumi.getter
    def configure(self) -> str:
        """
        The "configure" ACL.
        """
        return pulumi.get(self, "configure")

    @property
    @pulumi.getter
    def read(self) -> str:
        """
        The "read" ACL.
        """
        return pulumi.get(self, "read")

    @property
    @pulumi.getter
    def write(self) -> str:
        """
        The "write" ACL.
        """
        return pulumi.get(self, "write")


@pulumi.output_type
class PolicyPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "applyTo":
            suggest = "apply_to"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PolicyPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PolicyPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PolicyPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 apply_to: str,
                 definition: Mapping[str, Any],
                 pattern: str,
                 priority: int):
        """
        :param str apply_to: Can either be "exchanges", "queues", or "all".
        :param Mapping[str, Any] definition: Key/value pairs of the policy definition. See the
               RabbitMQ documentation for definition references and examples.
        :param str pattern: A pattern to match an exchange or queue name.
        :param int priority: The policy with the greater priority is applied first.
        """
        pulumi.set(__self__, "apply_to", apply_to)
        pulumi.set(__self__, "definition", definition)
        pulumi.set(__self__, "pattern", pattern)
        pulumi.set(__self__, "priority", priority)

    @property
    @pulumi.getter(name="applyTo")
    def apply_to(self) -> str:
        """
        Can either be "exchanges", "queues", or "all".
        """
        return pulumi.get(self, "apply_to")

    @property
    @pulumi.getter
    def definition(self) -> Mapping[str, Any]:
        """
        Key/value pairs of the policy definition. See the
        RabbitMQ documentation for definition references and examples.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter
    def pattern(self) -> str:
        """
        A pattern to match an exchange or queue name.
        """
        return pulumi.get(self, "pattern")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        The policy with the greater priority is applied first.
        """
        return pulumi.get(self, "priority")


@pulumi.output_type
class QueueSettings(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "argumentsJson":
            suggest = "arguments_json"
        elif key == "autoDelete":
            suggest = "auto_delete"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in QueueSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        QueueSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        QueueSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arguments: Optional[Mapping[str, Any]] = None,
                 arguments_json: Optional[str] = None,
                 auto_delete: Optional[bool] = None,
                 durable: Optional[bool] = None):
        """
        :param Mapping[str, Any] arguments: Additional key/value settings for the queue.
               All values will be sent to RabbitMQ as a string. If you require non-string
               values, use `arguments_json`.
        :param str arguments_json: A nested JSON string which contains additional
               settings for the queue. This is useful for when the arguments contain
               non-string values.
        :param bool auto_delete: Whether the queue will self-delete when all
               consumers have unsubscribed.
        :param bool durable: Whether the queue survives server restarts.
               Defaults to `false`.
        """
        if arguments is not None:
            pulumi.set(__self__, "arguments", arguments)
        if arguments_json is not None:
            pulumi.set(__self__, "arguments_json", arguments_json)
        if auto_delete is not None:
            pulumi.set(__self__, "auto_delete", auto_delete)
        if durable is not None:
            pulumi.set(__self__, "durable", durable)

    @property
    @pulumi.getter
    def arguments(self) -> Optional[Mapping[str, Any]]:
        """
        Additional key/value settings for the queue.
        All values will be sent to RabbitMQ as a string. If you require non-string
        values, use `arguments_json`.
        """
        return pulumi.get(self, "arguments")

    @property
    @pulumi.getter(name="argumentsJson")
    def arguments_json(self) -> Optional[str]:
        """
        A nested JSON string which contains additional
        settings for the queue. This is useful for when the arguments contain
        non-string values.
        """
        return pulumi.get(self, "arguments_json")

    @property
    @pulumi.getter(name="autoDelete")
    def auto_delete(self) -> Optional[bool]:
        """
        Whether the queue will self-delete when all
        consumers have unsubscribed.
        """
        return pulumi.get(self, "auto_delete")

    @property
    @pulumi.getter
    def durable(self) -> Optional[bool]:
        """
        Whether the queue survives server restarts.
        Defaults to `false`.
        """
        return pulumi.get(self, "durable")


@pulumi.output_type
class ShovelInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "destinationUri":
            suggest = "destination_uri"
        elif key == "sourceUri":
            suggest = "source_uri"
        elif key == "ackMode":
            suggest = "ack_mode"
        elif key == "addForwardHeaders":
            suggest = "add_forward_headers"
        elif key == "deleteAfter":
            suggest = "delete_after"
        elif key == "destinationAddForwardHeaders":
            suggest = "destination_add_forward_headers"
        elif key == "destinationAddTimestampHeader":
            suggest = "destination_add_timestamp_header"
        elif key == "destinationAddress":
            suggest = "destination_address"
        elif key == "destinationApplicationProperties":
            suggest = "destination_application_properties"
        elif key == "destinationExchange":
            suggest = "destination_exchange"
        elif key == "destinationExchangeKey":
            suggest = "destination_exchange_key"
        elif key == "destinationProperties":
            suggest = "destination_properties"
        elif key == "destinationProtocol":
            suggest = "destination_protocol"
        elif key == "destinationPublishProperties":
            suggest = "destination_publish_properties"
        elif key == "destinationQueue":
            suggest = "destination_queue"
        elif key == "prefetchCount":
            suggest = "prefetch_count"
        elif key == "reconnectDelay":
            suggest = "reconnect_delay"
        elif key == "sourceAddress":
            suggest = "source_address"
        elif key == "sourceDeleteAfter":
            suggest = "source_delete_after"
        elif key == "sourceExchange":
            suggest = "source_exchange"
        elif key == "sourceExchangeKey":
            suggest = "source_exchange_key"
        elif key == "sourcePrefetchCount":
            suggest = "source_prefetch_count"
        elif key == "sourceProtocol":
            suggest = "source_protocol"
        elif key == "sourceQueue":
            suggest = "source_queue"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ShovelInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ShovelInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ShovelInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 destination_uri: str,
                 source_uri: str,
                 ack_mode: Optional[str] = None,
                 add_forward_headers: Optional[bool] = None,
                 delete_after: Optional[str] = None,
                 destination_add_forward_headers: Optional[bool] = None,
                 destination_add_timestamp_header: Optional[bool] = None,
                 destination_address: Optional[str] = None,
                 destination_application_properties: Optional[str] = None,
                 destination_exchange: Optional[str] = None,
                 destination_exchange_key: Optional[str] = None,
                 destination_properties: Optional[str] = None,
                 destination_protocol: Optional[str] = None,
                 destination_publish_properties: Optional[str] = None,
                 destination_queue: Optional[str] = None,
                 prefetch_count: Optional[int] = None,
                 reconnect_delay: Optional[int] = None,
                 source_address: Optional[str] = None,
                 source_delete_after: Optional[str] = None,
                 source_exchange: Optional[str] = None,
                 source_exchange_key: Optional[str] = None,
                 source_prefetch_count: Optional[int] = None,
                 source_protocol: Optional[str] = None,
                 source_queue: Optional[str] = None):
        """
        :param str destination_uri: The amqp uri for the destination .
        :param str source_uri: The amqp uri for the source.
        :param str ack_mode: Determines how the shovel should acknowledge messages. Possible values are: `on-confirm`, `on-publish` and `no-ack`.
               Defaults to `on-confirm`.
        :param bool add_forward_headers: Whether to add `x-shovelled` headers to shovelled messages.
        :param str delete_after: Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
        :param bool destination_add_forward_headers: Whether to add `x-shovelled` headers to shovelled messages.
        :param str destination_address: The AMQP 1.0 destination link address.
        :param str destination_application_properties: Application properties to set when shovelling messages.
        :param str destination_exchange: The exchange to which messages should be published.
               Either this or `destination_queue` must be specified but not both.
        :param str destination_exchange_key: The routing key when using `destination_exchange`.
        :param str destination_properties: Properties to overwrite when shovelling messages.
        :param str destination_protocol: The protocol (`amqp091` or `amqp10`) to use when connecting to the destination.
               Defaults to `amqp091`.
        :param str destination_publish_properties: A map of properties to overwrite when shovelling messages.
        :param str destination_queue: The queue to which messages should be published.
               Either this or `destination_exchange` must be specified but not both.
        :param int prefetch_count: The maximum number of unacknowledged messages copied over a shovel at any one time.
        :param int reconnect_delay: The duration in seconds to reconnect to a broker after disconnected.
               Defaults to `1`.
        :param str source_address: The AMQP 1.0 source link address.
        :param str source_delete_after: Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
        :param str source_exchange: The exchange from which to consume.
               Either this or `source_queue` must be specified but not both.
        :param str source_exchange_key: The routing key when using `source_exchange`.
        :param int source_prefetch_count: The maximum number of unacknowledged messages copied over a shovel at any one time.
        :param str source_protocol: The protocol (`amqp091` or `amqp10`) to use when connecting to the source.
               Defaults to `amqp091`.
        :param str source_queue: The queue from which to consume.
               Either this or `source_exchange` must be specified but not both.
        """
        pulumi.set(__self__, "destination_uri", destination_uri)
        pulumi.set(__self__, "source_uri", source_uri)
        if ack_mode is not None:
            pulumi.set(__self__, "ack_mode", ack_mode)
        if add_forward_headers is not None:
            pulumi.set(__self__, "add_forward_headers", add_forward_headers)
        if delete_after is not None:
            pulumi.set(__self__, "delete_after", delete_after)
        if destination_add_forward_headers is not None:
            pulumi.set(__self__, "destination_add_forward_headers", destination_add_forward_headers)
        if destination_add_timestamp_header is not None:
            pulumi.set(__self__, "destination_add_timestamp_header", destination_add_timestamp_header)
        if destination_address is not None:
            pulumi.set(__self__, "destination_address", destination_address)
        if destination_application_properties is not None:
            pulumi.set(__self__, "destination_application_properties", destination_application_properties)
        if destination_exchange is not None:
            pulumi.set(__self__, "destination_exchange", destination_exchange)
        if destination_exchange_key is not None:
            pulumi.set(__self__, "destination_exchange_key", destination_exchange_key)
        if destination_properties is not None:
            pulumi.set(__self__, "destination_properties", destination_properties)
        if destination_protocol is not None:
            pulumi.set(__self__, "destination_protocol", destination_protocol)
        if destination_publish_properties is not None:
            pulumi.set(__self__, "destination_publish_properties", destination_publish_properties)
        if destination_queue is not None:
            pulumi.set(__self__, "destination_queue", destination_queue)
        if prefetch_count is not None:
            pulumi.set(__self__, "prefetch_count", prefetch_count)
        if reconnect_delay is not None:
            pulumi.set(__self__, "reconnect_delay", reconnect_delay)
        if source_address is not None:
            pulumi.set(__self__, "source_address", source_address)
        if source_delete_after is not None:
            pulumi.set(__self__, "source_delete_after", source_delete_after)
        if source_exchange is not None:
            pulumi.set(__self__, "source_exchange", source_exchange)
        if source_exchange_key is not None:
            pulumi.set(__self__, "source_exchange_key", source_exchange_key)
        if source_prefetch_count is not None:
            pulumi.set(__self__, "source_prefetch_count", source_prefetch_count)
        if source_protocol is not None:
            pulumi.set(__self__, "source_protocol", source_protocol)
        if source_queue is not None:
            pulumi.set(__self__, "source_queue", source_queue)

    @property
    @pulumi.getter(name="destinationUri")
    def destination_uri(self) -> str:
        """
        The amqp uri for the destination .
        """
        return pulumi.get(self, "destination_uri")

    @property
    @pulumi.getter(name="sourceUri")
    def source_uri(self) -> str:
        """
        The amqp uri for the source.
        """
        return pulumi.get(self, "source_uri")

    @property
    @pulumi.getter(name="ackMode")
    def ack_mode(self) -> Optional[str]:
        """
        Determines how the shovel should acknowledge messages. Possible values are: `on-confirm`, `on-publish` and `no-ack`.
        Defaults to `on-confirm`.
        """
        return pulumi.get(self, "ack_mode")

    @property
    @pulumi.getter(name="addForwardHeaders")
    def add_forward_headers(self) -> Optional[bool]:
        """
        Whether to add `x-shovelled` headers to shovelled messages.
        """
        return pulumi.get(self, "add_forward_headers")

    @property
    @pulumi.getter(name="deleteAfter")
    def delete_after(self) -> Optional[str]:
        """
        Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
        """
        return pulumi.get(self, "delete_after")

    @property
    @pulumi.getter(name="destinationAddForwardHeaders")
    def destination_add_forward_headers(self) -> Optional[bool]:
        """
        Whether to add `x-shovelled` headers to shovelled messages.
        """
        return pulumi.get(self, "destination_add_forward_headers")

    @property
    @pulumi.getter(name="destinationAddTimestampHeader")
    def destination_add_timestamp_header(self) -> Optional[bool]:
        return pulumi.get(self, "destination_add_timestamp_header")

    @property
    @pulumi.getter(name="destinationAddress")
    def destination_address(self) -> Optional[str]:
        """
        The AMQP 1.0 destination link address.
        """
        return pulumi.get(self, "destination_address")

    @property
    @pulumi.getter(name="destinationApplicationProperties")
    def destination_application_properties(self) -> Optional[str]:
        """
        Application properties to set when shovelling messages.
        """
        return pulumi.get(self, "destination_application_properties")

    @property
    @pulumi.getter(name="destinationExchange")
    def destination_exchange(self) -> Optional[str]:
        """
        The exchange to which messages should be published.
        Either this or `destination_queue` must be specified but not both.
        """
        return pulumi.get(self, "destination_exchange")

    @property
    @pulumi.getter(name="destinationExchangeKey")
    def destination_exchange_key(self) -> Optional[str]:
        """
        The routing key when using `destination_exchange`.
        """
        return pulumi.get(self, "destination_exchange_key")

    @property
    @pulumi.getter(name="destinationProperties")
    def destination_properties(self) -> Optional[str]:
        """
        Properties to overwrite when shovelling messages.
        """
        return pulumi.get(self, "destination_properties")

    @property
    @pulumi.getter(name="destinationProtocol")
    def destination_protocol(self) -> Optional[str]:
        """
        The protocol (`amqp091` or `amqp10`) to use when connecting to the destination.
        Defaults to `amqp091`.
        """
        return pulumi.get(self, "destination_protocol")

    @property
    @pulumi.getter(name="destinationPublishProperties")
    def destination_publish_properties(self) -> Optional[str]:
        """
        A map of properties to overwrite when shovelling messages.
        """
        return pulumi.get(self, "destination_publish_properties")

    @property
    @pulumi.getter(name="destinationQueue")
    def destination_queue(self) -> Optional[str]:
        """
        The queue to which messages should be published.
        Either this or `destination_exchange` must be specified but not both.
        """
        return pulumi.get(self, "destination_queue")

    @property
    @pulumi.getter(name="prefetchCount")
    def prefetch_count(self) -> Optional[int]:
        """
        The maximum number of unacknowledged messages copied over a shovel at any one time.
        """
        return pulumi.get(self, "prefetch_count")

    @property
    @pulumi.getter(name="reconnectDelay")
    def reconnect_delay(self) -> Optional[int]:
        """
        The duration in seconds to reconnect to a broker after disconnected.
        Defaults to `1`.
        """
        return pulumi.get(self, "reconnect_delay")

    @property
    @pulumi.getter(name="sourceAddress")
    def source_address(self) -> Optional[str]:
        """
        The AMQP 1.0 source link address.
        """
        return pulumi.get(self, "source_address")

    @property
    @pulumi.getter(name="sourceDeleteAfter")
    def source_delete_after(self) -> Optional[str]:
        """
        Determines when (if ever) the shovel should delete itself. Possible values are: `never`, `queue-length` or an integer.
        """
        return pulumi.get(self, "source_delete_after")

    @property
    @pulumi.getter(name="sourceExchange")
    def source_exchange(self) -> Optional[str]:
        """
        The exchange from which to consume.
        Either this or `source_queue` must be specified but not both.
        """
        return pulumi.get(self, "source_exchange")

    @property
    @pulumi.getter(name="sourceExchangeKey")
    def source_exchange_key(self) -> Optional[str]:
        """
        The routing key when using `source_exchange`.
        """
        return pulumi.get(self, "source_exchange_key")

    @property
    @pulumi.getter(name="sourcePrefetchCount")
    def source_prefetch_count(self) -> Optional[int]:
        """
        The maximum number of unacknowledged messages copied over a shovel at any one time.
        """
        return pulumi.get(self, "source_prefetch_count")

    @property
    @pulumi.getter(name="sourceProtocol")
    def source_protocol(self) -> Optional[str]:
        """
        The protocol (`amqp091` or `amqp10`) to use when connecting to the source.
        Defaults to `amqp091`.
        """
        return pulumi.get(self, "source_protocol")

    @property
    @pulumi.getter(name="sourceQueue")
    def source_queue(self) -> Optional[str]:
        """
        The queue from which to consume.
        Either this or `source_exchange` must be specified but not both.
        """
        return pulumi.get(self, "source_queue")


@pulumi.output_type
class TopicPermissionsPermission(dict):
    def __init__(__self__, *,
                 exchange: str,
                 read: str,
                 write: str):
        """
        :param str exchange: The exchange to set the permissions for.
        :param str read: The "read" ACL.
        :param str write: The "write" ACL.
        """
        pulumi.set(__self__, "exchange", exchange)
        pulumi.set(__self__, "read", read)
        pulumi.set(__self__, "write", write)

    @property
    @pulumi.getter
    def exchange(self) -> str:
        """
        The exchange to set the permissions for.
        """
        return pulumi.get(self, "exchange")

    @property
    @pulumi.getter
    def read(self) -> str:
        """
        The "read" ACL.
        """
        return pulumi.get(self, "read")

    @property
    @pulumi.getter
    def write(self) -> str:
        """
        The "write" ACL.
        """
        return pulumi.get(self, "write")


@pulumi.output_type
class GetExchangeSettingResult(dict):
    def __init__(__self__, *,
                 type: str,
                 arguments: Optional[Mapping[str, Any]] = None,
                 auto_delete: Optional[bool] = None,
                 durable: Optional[bool] = None):
        pulumi.set(__self__, "type", type)
        if arguments is not None:
            pulumi.set(__self__, "arguments", arguments)
        if auto_delete is not None:
            pulumi.set(__self__, "auto_delete", auto_delete)
        if durable is not None:
            pulumi.set(__self__, "durable", durable)

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def arguments(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "arguments")

    @property
    @pulumi.getter(name="autoDelete")
    def auto_delete(self) -> Optional[bool]:
        return pulumi.get(self, "auto_delete")

    @property
    @pulumi.getter
    def durable(self) -> Optional[bool]:
        return pulumi.get(self, "durable")


