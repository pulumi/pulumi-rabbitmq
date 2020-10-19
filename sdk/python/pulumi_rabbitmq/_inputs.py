# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'ExchangeSettingsArgs',
    'FederationUpstreamDefinitionArgs',
    'PermissionsPermissionsArgs',
    'PolicyPolicyArgs',
    'QueueSettingsArgs',
    'ShovelInfoArgs',
    'TopicPermissionsPermissionArgs',
]

@pulumi.input_type
class ExchangeSettingsArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 arguments: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 auto_delete: Optional[pulumi.Input[bool]] = None,
                 durable: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] type: The type of exchange.
        :param pulumi.Input[Mapping[str, Any]] arguments: Additional key/value settings for the exchange.
        :param pulumi.Input[bool] auto_delete: Whether the exchange will self-delete when all
               queues have finished using it.
        :param pulumi.Input[bool] durable: Whether the exchange survives server restarts.
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
    def type(self) -> pulumi.Input[str]:
        """
        The type of exchange.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def arguments(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Additional key/value settings for the exchange.
        """
        return pulumi.get(self, "arguments")

    @arguments.setter
    def arguments(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "arguments", value)

    @property
    @pulumi.getter(name="autoDelete")
    def auto_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the exchange will self-delete when all
        queues have finished using it.
        """
        return pulumi.get(self, "auto_delete")

    @auto_delete.setter
    def auto_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_delete", value)

    @property
    @pulumi.getter
    def durable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the exchange survives server restarts.
        Defaults to `false`.
        """
        return pulumi.get(self, "durable")

    @durable.setter
    def durable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "durable", value)


@pulumi.input_type
class FederationUpstreamDefinitionArgs:
    def __init__(__self__, *,
                 uri: pulumi.Input[str],
                 ack_mode: Optional[pulumi.Input[str]] = None,
                 exchange: Optional[pulumi.Input[str]] = None,
                 expires: Optional[pulumi.Input[int]] = None,
                 max_hops: Optional[pulumi.Input[int]] = None,
                 message_ttl: Optional[pulumi.Input[int]] = None,
                 prefetch_count: Optional[pulumi.Input[int]] = None,
                 queue: Optional[pulumi.Input[str]] = None,
                 reconnect_delay: Optional[pulumi.Input[int]] = None,
                 trust_user_id: Optional[pulumi.Input[bool]] = None):
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
    def uri(self) -> pulumi.Input[str]:
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "uri", value)

    @property
    @pulumi.getter(name="ackMode")
    def ack_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ack_mode")

    @ack_mode.setter
    def ack_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ack_mode", value)

    @property
    @pulumi.getter
    def exchange(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "exchange")

    @exchange.setter
    def exchange(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exchange", value)

    @property
    @pulumi.getter
    def expires(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "expires")

    @expires.setter
    def expires(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expires", value)

    @property
    @pulumi.getter(name="maxHops")
    def max_hops(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_hops")

    @max_hops.setter
    def max_hops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_hops", value)

    @property
    @pulumi.getter(name="messageTtl")
    def message_ttl(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "message_ttl")

    @message_ttl.setter
    def message_ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_ttl", value)

    @property
    @pulumi.getter(name="prefetchCount")
    def prefetch_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "prefetch_count")

    @prefetch_count.setter
    def prefetch_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "prefetch_count", value)

    @property
    @pulumi.getter
    def queue(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "queue")

    @queue.setter
    def queue(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "queue", value)

    @property
    @pulumi.getter(name="reconnectDelay")
    def reconnect_delay(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "reconnect_delay")

    @reconnect_delay.setter
    def reconnect_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reconnect_delay", value)

    @property
    @pulumi.getter(name="trustUserId")
    def trust_user_id(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "trust_user_id")

    @trust_user_id.setter
    def trust_user_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "trust_user_id", value)


@pulumi.input_type
class PermissionsPermissionsArgs:
    def __init__(__self__, *,
                 configure: pulumi.Input[str],
                 read: pulumi.Input[str],
                 write: pulumi.Input[str]):
        """
        :param pulumi.Input[str] configure: The "configure" ACL.
        :param pulumi.Input[str] read: The "read" ACL.
        :param pulumi.Input[str] write: The "write" ACL.
        """
        pulumi.set(__self__, "configure", configure)
        pulumi.set(__self__, "read", read)
        pulumi.set(__self__, "write", write)

    @property
    @pulumi.getter
    def configure(self) -> pulumi.Input[str]:
        """
        The "configure" ACL.
        """
        return pulumi.get(self, "configure")

    @configure.setter
    def configure(self, value: pulumi.Input[str]):
        pulumi.set(self, "configure", value)

    @property
    @pulumi.getter
    def read(self) -> pulumi.Input[str]:
        """
        The "read" ACL.
        """
        return pulumi.get(self, "read")

    @read.setter
    def read(self, value: pulumi.Input[str]):
        pulumi.set(self, "read", value)

    @property
    @pulumi.getter
    def write(self) -> pulumi.Input[str]:
        """
        The "write" ACL.
        """
        return pulumi.get(self, "write")

    @write.setter
    def write(self, value: pulumi.Input[str]):
        pulumi.set(self, "write", value)


@pulumi.input_type
class PolicyPolicyArgs:
    def __init__(__self__, *,
                 apply_to: pulumi.Input[str],
                 definition: pulumi.Input[Mapping[str, Any]],
                 pattern: pulumi.Input[str],
                 priority: pulumi.Input[int]):
        """
        :param pulumi.Input[str] apply_to: Can either be "exchanges", "queues", or "all".
        :param pulumi.Input[Mapping[str, Any]] definition: Key/value pairs of the policy definition. See the
               RabbitMQ documentation for definition references and examples.
        :param pulumi.Input[str] pattern: A pattern to match an exchange or queue name.
        :param pulumi.Input[int] priority: The policy with the greater priority is applied first.
        """
        pulumi.set(__self__, "apply_to", apply_to)
        pulumi.set(__self__, "definition", definition)
        pulumi.set(__self__, "pattern", pattern)
        pulumi.set(__self__, "priority", priority)

    @property
    @pulumi.getter(name="applyTo")
    def apply_to(self) -> pulumi.Input[str]:
        """
        Can either be "exchanges", "queues", or "all".
        """
        return pulumi.get(self, "apply_to")

    @apply_to.setter
    def apply_to(self, value: pulumi.Input[str]):
        pulumi.set(self, "apply_to", value)

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Input[Mapping[str, Any]]:
        """
        Key/value pairs of the policy definition. See the
        RabbitMQ documentation for definition references and examples.
        """
        return pulumi.get(self, "definition")

    @definition.setter
    def definition(self, value: pulumi.Input[Mapping[str, Any]]):
        pulumi.set(self, "definition", value)

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Input[str]:
        """
        A pattern to match an exchange or queue name.
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: pulumi.Input[str]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[int]:
        """
        The policy with the greater priority is applied first.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[int]):
        pulumi.set(self, "priority", value)


@pulumi.input_type
class QueueSettingsArgs:
    def __init__(__self__, *,
                 arguments: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 arguments_json: Optional[pulumi.Input[str]] = None,
                 auto_delete: Optional[pulumi.Input[bool]] = None,
                 durable: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[Mapping[str, Any]] arguments: Additional key/value settings for the queue.
               All values will be sent to RabbitMQ as a string. If you require non-string
               values, use `arguments_json`.
        :param pulumi.Input[str] arguments_json: A nested JSON string which contains additional
               settings for the queue. This is useful for when the arguments contain
               non-string values.
        :param pulumi.Input[bool] auto_delete: Whether the queue will self-delete when all
               consumers have unsubscribed.
        :param pulumi.Input[bool] durable: Whether the queue survives server restarts.
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
    def arguments(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Additional key/value settings for the queue.
        All values will be sent to RabbitMQ as a string. If you require non-string
        values, use `arguments_json`.
        """
        return pulumi.get(self, "arguments")

    @arguments.setter
    def arguments(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "arguments", value)

    @property
    @pulumi.getter(name="argumentsJson")
    def arguments_json(self) -> Optional[pulumi.Input[str]]:
        """
        A nested JSON string which contains additional
        settings for the queue. This is useful for when the arguments contain
        non-string values.
        """
        return pulumi.get(self, "arguments_json")

    @arguments_json.setter
    def arguments_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arguments_json", value)

    @property
    @pulumi.getter(name="autoDelete")
    def auto_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the queue will self-delete when all
        consumers have unsubscribed.
        """
        return pulumi.get(self, "auto_delete")

    @auto_delete.setter
    def auto_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_delete", value)

    @property
    @pulumi.getter
    def durable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the queue survives server restarts.
        Defaults to `false`.
        """
        return pulumi.get(self, "durable")

    @durable.setter
    def durable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "durable", value)


@pulumi.input_type
class ShovelInfoArgs:
    def __init__(__self__, *,
                 destination_uri: pulumi.Input[str],
                 source_uri: pulumi.Input[str],
                 ack_mode: Optional[pulumi.Input[str]] = None,
                 add_forward_headers: Optional[pulumi.Input[bool]] = None,
                 delete_after: Optional[pulumi.Input[str]] = None,
                 destination_exchange: Optional[pulumi.Input[str]] = None,
                 destination_exchange_key: Optional[pulumi.Input[str]] = None,
                 destination_queue: Optional[pulumi.Input[str]] = None,
                 prefetch_count: Optional[pulumi.Input[int]] = None,
                 reconnect_delay: Optional[pulumi.Input[int]] = None,
                 source_exchange: Optional[pulumi.Input[str]] = None,
                 source_exchange_key: Optional[pulumi.Input[str]] = None,
                 source_queue: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] destination_uri: The amqp uri for the destination .
        :param pulumi.Input[str] source_uri: The amqp uri for the source.
        :param pulumi.Input[str] ack_mode: Determines how the shovel should acknowledge messages.
               Defaults to `on-confirm`.
        :param pulumi.Input[bool] add_forward_headers: Whether to amqp shovel headers.
               Defaults to `false`.
        :param pulumi.Input[str] delete_after: Determines when (if ever) the shovel should delete itself .
               Defaults to `never`.
        :param pulumi.Input[str] destination_exchange: The exchange to which messages should be published.
               Either this or destination_queue must be specified but not both.
        :param pulumi.Input[str] destination_exchange_key: The routing key when using destination_exchange.
        :param pulumi.Input[str] destination_queue: The queue to which messages should be published.
               Either this or destination_exchange must be specified but not both.
        :param pulumi.Input[int] prefetch_count: The maximum number of unacknowledged messages copied over a shovel at any one time.
               Defaults to `1000`.
        :param pulumi.Input[int] reconnect_delay: The duration in seconds to reconnect to a broker after disconnected.
               Defaults to `1`.
        :param pulumi.Input[str] source_exchange: The exchange from which to consume.
               Either this or source_queue must be specified but not both.
        :param pulumi.Input[str] source_exchange_key: The routing key when using source_exchange.
        :param pulumi.Input[str] source_queue: The queue from which to consume.
               Either this or source_exchange must be specified but not both.
        """
        pulumi.set(__self__, "destination_uri", destination_uri)
        pulumi.set(__self__, "source_uri", source_uri)
        if ack_mode is not None:
            pulumi.set(__self__, "ack_mode", ack_mode)
        if add_forward_headers is not None:
            pulumi.set(__self__, "add_forward_headers", add_forward_headers)
        if delete_after is not None:
            pulumi.set(__self__, "delete_after", delete_after)
        if destination_exchange is not None:
            pulumi.set(__self__, "destination_exchange", destination_exchange)
        if destination_exchange_key is not None:
            pulumi.set(__self__, "destination_exchange_key", destination_exchange_key)
        if destination_queue is not None:
            pulumi.set(__self__, "destination_queue", destination_queue)
        if prefetch_count is not None:
            pulumi.set(__self__, "prefetch_count", prefetch_count)
        if reconnect_delay is not None:
            pulumi.set(__self__, "reconnect_delay", reconnect_delay)
        if source_exchange is not None:
            pulumi.set(__self__, "source_exchange", source_exchange)
        if source_exchange_key is not None:
            pulumi.set(__self__, "source_exchange_key", source_exchange_key)
        if source_queue is not None:
            pulumi.set(__self__, "source_queue", source_queue)

    @property
    @pulumi.getter(name="destinationUri")
    def destination_uri(self) -> pulumi.Input[str]:
        """
        The amqp uri for the destination .
        """
        return pulumi.get(self, "destination_uri")

    @destination_uri.setter
    def destination_uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_uri", value)

    @property
    @pulumi.getter(name="sourceUri")
    def source_uri(self) -> pulumi.Input[str]:
        """
        The amqp uri for the source.
        """
        return pulumi.get(self, "source_uri")

    @source_uri.setter
    def source_uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_uri", value)

    @property
    @pulumi.getter(name="ackMode")
    def ack_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Determines how the shovel should acknowledge messages.
        Defaults to `on-confirm`.
        """
        return pulumi.get(self, "ack_mode")

    @ack_mode.setter
    def ack_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ack_mode", value)

    @property
    @pulumi.getter(name="addForwardHeaders")
    def add_forward_headers(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to amqp shovel headers.
        Defaults to `false`.
        """
        return pulumi.get(self, "add_forward_headers")

    @add_forward_headers.setter
    def add_forward_headers(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_forward_headers", value)

    @property
    @pulumi.getter(name="deleteAfter")
    def delete_after(self) -> Optional[pulumi.Input[str]]:
        """
        Determines when (if ever) the shovel should delete itself .
        Defaults to `never`.
        """
        return pulumi.get(self, "delete_after")

    @delete_after.setter
    def delete_after(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_after", value)

    @property
    @pulumi.getter(name="destinationExchange")
    def destination_exchange(self) -> Optional[pulumi.Input[str]]:
        """
        The exchange to which messages should be published.
        Either this or destination_queue must be specified but not both.
        """
        return pulumi.get(self, "destination_exchange")

    @destination_exchange.setter
    def destination_exchange(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_exchange", value)

    @property
    @pulumi.getter(name="destinationExchangeKey")
    def destination_exchange_key(self) -> Optional[pulumi.Input[str]]:
        """
        The routing key when using destination_exchange.
        """
        return pulumi.get(self, "destination_exchange_key")

    @destination_exchange_key.setter
    def destination_exchange_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_exchange_key", value)

    @property
    @pulumi.getter(name="destinationQueue")
    def destination_queue(self) -> Optional[pulumi.Input[str]]:
        """
        The queue to which messages should be published.
        Either this or destination_exchange must be specified but not both.
        """
        return pulumi.get(self, "destination_queue")

    @destination_queue.setter
    def destination_queue(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_queue", value)

    @property
    @pulumi.getter(name="prefetchCount")
    def prefetch_count(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of unacknowledged messages copied over a shovel at any one time.
        Defaults to `1000`.
        """
        return pulumi.get(self, "prefetch_count")

    @prefetch_count.setter
    def prefetch_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "prefetch_count", value)

    @property
    @pulumi.getter(name="reconnectDelay")
    def reconnect_delay(self) -> Optional[pulumi.Input[int]]:
        """
        The duration in seconds to reconnect to a broker after disconnected.
        Defaults to `1`.
        """
        return pulumi.get(self, "reconnect_delay")

    @reconnect_delay.setter
    def reconnect_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reconnect_delay", value)

    @property
    @pulumi.getter(name="sourceExchange")
    def source_exchange(self) -> Optional[pulumi.Input[str]]:
        """
        The exchange from which to consume.
        Either this or source_queue must be specified but not both.
        """
        return pulumi.get(self, "source_exchange")

    @source_exchange.setter
    def source_exchange(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_exchange", value)

    @property
    @pulumi.getter(name="sourceExchangeKey")
    def source_exchange_key(self) -> Optional[pulumi.Input[str]]:
        """
        The routing key when using source_exchange.
        """
        return pulumi.get(self, "source_exchange_key")

    @source_exchange_key.setter
    def source_exchange_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_exchange_key", value)

    @property
    @pulumi.getter(name="sourceQueue")
    def source_queue(self) -> Optional[pulumi.Input[str]]:
        """
        The queue from which to consume.
        Either this or source_exchange must be specified but not both.
        """
        return pulumi.get(self, "source_queue")

    @source_queue.setter
    def source_queue(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_queue", value)


@pulumi.input_type
class TopicPermissionsPermissionArgs:
    def __init__(__self__, *,
                 exchange: pulumi.Input[str],
                 read: pulumi.Input[str],
                 write: pulumi.Input[str]):
        """
        :param pulumi.Input[str] exchange: The exchange to set the permissions for.
        :param pulumi.Input[str] read: The "read" ACL.
        :param pulumi.Input[str] write: The "write" ACL.
        """
        pulumi.set(__self__, "exchange", exchange)
        pulumi.set(__self__, "read", read)
        pulumi.set(__self__, "write", write)

    @property
    @pulumi.getter
    def exchange(self) -> pulumi.Input[str]:
        """
        The exchange to set the permissions for.
        """
        return pulumi.get(self, "exchange")

    @exchange.setter
    def exchange(self, value: pulumi.Input[str]):
        pulumi.set(self, "exchange", value)

    @property
    @pulumi.getter
    def read(self) -> pulumi.Input[str]:
        """
        The "read" ACL.
        """
        return pulumi.get(self, "read")

    @read.setter
    def read(self, value: pulumi.Input[str]):
        pulumi.set(self, "read", value)

    @property
    @pulumi.getter
    def write(self) -> pulumi.Input[str]:
        """
        The "write" ACL.
        """
        return pulumi.get(self, "write")

    @write.setter
    def write(self, value: pulumi.Input[str]):
        pulumi.set(self, "write", value)


