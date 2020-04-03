// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package rabbitmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ExchangeSettings struct {
	// Additional key/value settings for the exchange.
	Arguments map[string]interface{} `pulumi:"arguments"`
	// Whether the exchange will self-delete when all
	// queues have finished using it.
	AutoDelete *bool `pulumi:"autoDelete"`
	// Whether the exchange survives server restarts.
	// Defaults to `false`.
	Durable *bool `pulumi:"durable"`
	// The type of exchange.
	Type string `pulumi:"type"`
}

type ExchangeSettingsInput interface {
	pulumi.Input

	ToExchangeSettingsOutput() ExchangeSettingsOutput
	ToExchangeSettingsOutputWithContext(context.Context) ExchangeSettingsOutput
}

type ExchangeSettingsArgs struct {
	// Additional key/value settings for the exchange.
	Arguments pulumi.MapInput `pulumi:"arguments"`
	// Whether the exchange will self-delete when all
	// queues have finished using it.
	AutoDelete pulumi.BoolPtrInput `pulumi:"autoDelete"`
	// Whether the exchange survives server restarts.
	// Defaults to `false`.
	Durable pulumi.BoolPtrInput `pulumi:"durable"`
	// The type of exchange.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ExchangeSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExchangeSettings)(nil)).Elem()
}

func (i ExchangeSettingsArgs) ToExchangeSettingsOutput() ExchangeSettingsOutput {
	return i.ToExchangeSettingsOutputWithContext(context.Background())
}

func (i ExchangeSettingsArgs) ToExchangeSettingsOutputWithContext(ctx context.Context) ExchangeSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeSettingsOutput)
}

func (i ExchangeSettingsArgs) ToExchangeSettingsPtrOutput() ExchangeSettingsPtrOutput {
	return i.ToExchangeSettingsPtrOutputWithContext(context.Background())
}

func (i ExchangeSettingsArgs) ToExchangeSettingsPtrOutputWithContext(ctx context.Context) ExchangeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeSettingsOutput).ToExchangeSettingsPtrOutputWithContext(ctx)
}

type ExchangeSettingsPtrInput interface {
	pulumi.Input

	ToExchangeSettingsPtrOutput() ExchangeSettingsPtrOutput
	ToExchangeSettingsPtrOutputWithContext(context.Context) ExchangeSettingsPtrOutput
}

type exchangeSettingsPtrType ExchangeSettingsArgs

func ExchangeSettingsPtr(v *ExchangeSettingsArgs) ExchangeSettingsPtrInput {
	return (*exchangeSettingsPtrType)(v)
}

func (*exchangeSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExchangeSettings)(nil)).Elem()
}

func (i *exchangeSettingsPtrType) ToExchangeSettingsPtrOutput() ExchangeSettingsPtrOutput {
	return i.ToExchangeSettingsPtrOutputWithContext(context.Background())
}

func (i *exchangeSettingsPtrType) ToExchangeSettingsPtrOutputWithContext(ctx context.Context) ExchangeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExchangeSettingsPtrOutput)
}

type ExchangeSettingsOutput struct{ *pulumi.OutputState }

func (ExchangeSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExchangeSettings)(nil)).Elem()
}

func (o ExchangeSettingsOutput) ToExchangeSettingsOutput() ExchangeSettingsOutput {
	return o
}

func (o ExchangeSettingsOutput) ToExchangeSettingsOutputWithContext(ctx context.Context) ExchangeSettingsOutput {
	return o
}

func (o ExchangeSettingsOutput) ToExchangeSettingsPtrOutput() ExchangeSettingsPtrOutput {
	return o.ToExchangeSettingsPtrOutputWithContext(context.Background())
}

func (o ExchangeSettingsOutput) ToExchangeSettingsPtrOutputWithContext(ctx context.Context) ExchangeSettingsPtrOutput {
	return o.ApplyT(func(v ExchangeSettings) *ExchangeSettings {
		return &v
	}).(ExchangeSettingsPtrOutput)
}

// Additional key/value settings for the exchange.
func (o ExchangeSettingsOutput) Arguments() pulumi.MapOutput {
	return o.ApplyT(func(v ExchangeSettings) map[string]interface{} { return v.Arguments }).(pulumi.MapOutput)
}

// Whether the exchange will self-delete when all
// queues have finished using it.
func (o ExchangeSettingsOutput) AutoDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExchangeSettings) *bool { return v.AutoDelete }).(pulumi.BoolPtrOutput)
}

// Whether the exchange survives server restarts.
// Defaults to `false`.
func (o ExchangeSettingsOutput) Durable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExchangeSettings) *bool { return v.Durable }).(pulumi.BoolPtrOutput)
}

// The type of exchange.
func (o ExchangeSettingsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ExchangeSettings) string { return v.Type }).(pulumi.StringOutput)
}

type ExchangeSettingsPtrOutput struct{ *pulumi.OutputState }

func (ExchangeSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExchangeSettings)(nil)).Elem()
}

func (o ExchangeSettingsPtrOutput) ToExchangeSettingsPtrOutput() ExchangeSettingsPtrOutput {
	return o
}

func (o ExchangeSettingsPtrOutput) ToExchangeSettingsPtrOutputWithContext(ctx context.Context) ExchangeSettingsPtrOutput {
	return o
}

func (o ExchangeSettingsPtrOutput) Elem() ExchangeSettingsOutput {
	return o.ApplyT(func(v *ExchangeSettings) ExchangeSettings { return *v }).(ExchangeSettingsOutput)
}

// Additional key/value settings for the exchange.
func (o ExchangeSettingsPtrOutput) Arguments() pulumi.MapOutput {
	return o.ApplyT(func(v ExchangeSettings) map[string]interface{} { return v.Arguments }).(pulumi.MapOutput)
}

// Whether the exchange will self-delete when all
// queues have finished using it.
func (o ExchangeSettingsPtrOutput) AutoDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExchangeSettings) *bool { return v.AutoDelete }).(pulumi.BoolPtrOutput)
}

// Whether the exchange survives server restarts.
// Defaults to `false`.
func (o ExchangeSettingsPtrOutput) Durable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExchangeSettings) *bool { return v.Durable }).(pulumi.BoolPtrOutput)
}

// The type of exchange.
func (o ExchangeSettingsPtrOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ExchangeSettings) string { return v.Type }).(pulumi.StringOutput)
}

type PermissionsPermissions struct {
	// The "configure" ACL.
	Configure string `pulumi:"configure"`
	// The "read" ACL.
	Read string `pulumi:"read"`
	// The "write" ACL.
	Write string `pulumi:"write"`
}

type PermissionsPermissionsInput interface {
	pulumi.Input

	ToPermissionsPermissionsOutput() PermissionsPermissionsOutput
	ToPermissionsPermissionsOutputWithContext(context.Context) PermissionsPermissionsOutput
}

type PermissionsPermissionsArgs struct {
	// The "configure" ACL.
	Configure pulumi.StringInput `pulumi:"configure"`
	// The "read" ACL.
	Read pulumi.StringInput `pulumi:"read"`
	// The "write" ACL.
	Write pulumi.StringInput `pulumi:"write"`
}

func (PermissionsPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsPermissions)(nil)).Elem()
}

func (i PermissionsPermissionsArgs) ToPermissionsPermissionsOutput() PermissionsPermissionsOutput {
	return i.ToPermissionsPermissionsOutputWithContext(context.Background())
}

func (i PermissionsPermissionsArgs) ToPermissionsPermissionsOutputWithContext(ctx context.Context) PermissionsPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsPermissionsOutput)
}

func (i PermissionsPermissionsArgs) ToPermissionsPermissionsPtrOutput() PermissionsPermissionsPtrOutput {
	return i.ToPermissionsPermissionsPtrOutputWithContext(context.Background())
}

func (i PermissionsPermissionsArgs) ToPermissionsPermissionsPtrOutputWithContext(ctx context.Context) PermissionsPermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsPermissionsOutput).ToPermissionsPermissionsPtrOutputWithContext(ctx)
}

type PermissionsPermissionsPtrInput interface {
	pulumi.Input

	ToPermissionsPermissionsPtrOutput() PermissionsPermissionsPtrOutput
	ToPermissionsPermissionsPtrOutputWithContext(context.Context) PermissionsPermissionsPtrOutput
}

type permissionsPermissionsPtrType PermissionsPermissionsArgs

func PermissionsPermissionsPtr(v *PermissionsPermissionsArgs) PermissionsPermissionsPtrInput {
	return (*permissionsPermissionsPtrType)(v)
}

func (*permissionsPermissionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionsPermissions)(nil)).Elem()
}

func (i *permissionsPermissionsPtrType) ToPermissionsPermissionsPtrOutput() PermissionsPermissionsPtrOutput {
	return i.ToPermissionsPermissionsPtrOutputWithContext(context.Background())
}

func (i *permissionsPermissionsPtrType) ToPermissionsPermissionsPtrOutputWithContext(ctx context.Context) PermissionsPermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsPermissionsPtrOutput)
}

type PermissionsPermissionsOutput struct{ *pulumi.OutputState }

func (PermissionsPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsPermissions)(nil)).Elem()
}

func (o PermissionsPermissionsOutput) ToPermissionsPermissionsOutput() PermissionsPermissionsOutput {
	return o
}

func (o PermissionsPermissionsOutput) ToPermissionsPermissionsOutputWithContext(ctx context.Context) PermissionsPermissionsOutput {
	return o
}

func (o PermissionsPermissionsOutput) ToPermissionsPermissionsPtrOutput() PermissionsPermissionsPtrOutput {
	return o.ToPermissionsPermissionsPtrOutputWithContext(context.Background())
}

func (o PermissionsPermissionsOutput) ToPermissionsPermissionsPtrOutputWithContext(ctx context.Context) PermissionsPermissionsPtrOutput {
	return o.ApplyT(func(v PermissionsPermissions) *PermissionsPermissions {
		return &v
	}).(PermissionsPermissionsPtrOutput)
}

// The "configure" ACL.
func (o PermissionsPermissionsOutput) Configure() pulumi.StringOutput {
	return o.ApplyT(func(v PermissionsPermissions) string { return v.Configure }).(pulumi.StringOutput)
}

// The "read" ACL.
func (o PermissionsPermissionsOutput) Read() pulumi.StringOutput {
	return o.ApplyT(func(v PermissionsPermissions) string { return v.Read }).(pulumi.StringOutput)
}

// The "write" ACL.
func (o PermissionsPermissionsOutput) Write() pulumi.StringOutput {
	return o.ApplyT(func(v PermissionsPermissions) string { return v.Write }).(pulumi.StringOutput)
}

type PermissionsPermissionsPtrOutput struct{ *pulumi.OutputState }

func (PermissionsPermissionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionsPermissions)(nil)).Elem()
}

func (o PermissionsPermissionsPtrOutput) ToPermissionsPermissionsPtrOutput() PermissionsPermissionsPtrOutput {
	return o
}

func (o PermissionsPermissionsPtrOutput) ToPermissionsPermissionsPtrOutputWithContext(ctx context.Context) PermissionsPermissionsPtrOutput {
	return o
}

func (o PermissionsPermissionsPtrOutput) Elem() PermissionsPermissionsOutput {
	return o.ApplyT(func(v *PermissionsPermissions) PermissionsPermissions { return *v }).(PermissionsPermissionsOutput)
}

// The "configure" ACL.
func (o PermissionsPermissionsPtrOutput) Configure() pulumi.StringOutput {
	return o.ApplyT(func(v PermissionsPermissions) string { return v.Configure }).(pulumi.StringOutput)
}

// The "read" ACL.
func (o PermissionsPermissionsPtrOutput) Read() pulumi.StringOutput {
	return o.ApplyT(func(v PermissionsPermissions) string { return v.Read }).(pulumi.StringOutput)
}

// The "write" ACL.
func (o PermissionsPermissionsPtrOutput) Write() pulumi.StringOutput {
	return o.ApplyT(func(v PermissionsPermissions) string { return v.Write }).(pulumi.StringOutput)
}

type PolicyPolicy struct {
	// Can either be "exchanges", "queues", or "all".
	ApplyTo string `pulumi:"applyTo"`
	// Key/value pairs of the policy definition. See the
	// RabbitMQ documentation for definition references and examples.
	Definition map[string]interface{} `pulumi:"definition"`
	// A pattern to match an exchange or queue name.
	Pattern string `pulumi:"pattern"`
	// The policy with the greater priority is applied first.
	Priority int `pulumi:"priority"`
}

type PolicyPolicyInput interface {
	pulumi.Input

	ToPolicyPolicyOutput() PolicyPolicyOutput
	ToPolicyPolicyOutputWithContext(context.Context) PolicyPolicyOutput
}

type PolicyPolicyArgs struct {
	// Can either be "exchanges", "queues", or "all".
	ApplyTo pulumi.StringInput `pulumi:"applyTo"`
	// Key/value pairs of the policy definition. See the
	// RabbitMQ documentation for definition references and examples.
	Definition pulumi.MapInput `pulumi:"definition"`
	// A pattern to match an exchange or queue name.
	Pattern pulumi.StringInput `pulumi:"pattern"`
	// The policy with the greater priority is applied first.
	Priority pulumi.IntInput `pulumi:"priority"`
}

func (PolicyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyPolicy)(nil)).Elem()
}

func (i PolicyPolicyArgs) ToPolicyPolicyOutput() PolicyPolicyOutput {
	return i.ToPolicyPolicyOutputWithContext(context.Background())
}

func (i PolicyPolicyArgs) ToPolicyPolicyOutputWithContext(ctx context.Context) PolicyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPolicyOutput)
}

func (i PolicyPolicyArgs) ToPolicyPolicyPtrOutput() PolicyPolicyPtrOutput {
	return i.ToPolicyPolicyPtrOutputWithContext(context.Background())
}

func (i PolicyPolicyArgs) ToPolicyPolicyPtrOutputWithContext(ctx context.Context) PolicyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPolicyOutput).ToPolicyPolicyPtrOutputWithContext(ctx)
}

type PolicyPolicyPtrInput interface {
	pulumi.Input

	ToPolicyPolicyPtrOutput() PolicyPolicyPtrOutput
	ToPolicyPolicyPtrOutputWithContext(context.Context) PolicyPolicyPtrOutput
}

type policyPolicyPtrType PolicyPolicyArgs

func PolicyPolicyPtr(v *PolicyPolicyArgs) PolicyPolicyPtrInput {
	return (*policyPolicyPtrType)(v)
}

func (*policyPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyPolicy)(nil)).Elem()
}

func (i *policyPolicyPtrType) ToPolicyPolicyPtrOutput() PolicyPolicyPtrOutput {
	return i.ToPolicyPolicyPtrOutputWithContext(context.Background())
}

func (i *policyPolicyPtrType) ToPolicyPolicyPtrOutputWithContext(ctx context.Context) PolicyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPolicyPtrOutput)
}

type PolicyPolicyOutput struct{ *pulumi.OutputState }

func (PolicyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyPolicy)(nil)).Elem()
}

func (o PolicyPolicyOutput) ToPolicyPolicyOutput() PolicyPolicyOutput {
	return o
}

func (o PolicyPolicyOutput) ToPolicyPolicyOutputWithContext(ctx context.Context) PolicyPolicyOutput {
	return o
}

func (o PolicyPolicyOutput) ToPolicyPolicyPtrOutput() PolicyPolicyPtrOutput {
	return o.ToPolicyPolicyPtrOutputWithContext(context.Background())
}

func (o PolicyPolicyOutput) ToPolicyPolicyPtrOutputWithContext(ctx context.Context) PolicyPolicyPtrOutput {
	return o.ApplyT(func(v PolicyPolicy) *PolicyPolicy {
		return &v
	}).(PolicyPolicyPtrOutput)
}

// Can either be "exchanges", "queues", or "all".
func (o PolicyPolicyOutput) ApplyTo() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyPolicy) string { return v.ApplyTo }).(pulumi.StringOutput)
}

// Key/value pairs of the policy definition. See the
// RabbitMQ documentation for definition references and examples.
func (o PolicyPolicyOutput) Definition() pulumi.MapOutput {
	return o.ApplyT(func(v PolicyPolicy) map[string]interface{} { return v.Definition }).(pulumi.MapOutput)
}

// A pattern to match an exchange or queue name.
func (o PolicyPolicyOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyPolicy) string { return v.Pattern }).(pulumi.StringOutput)
}

// The policy with the greater priority is applied first.
func (o PolicyPolicyOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v PolicyPolicy) int { return v.Priority }).(pulumi.IntOutput)
}

type PolicyPolicyPtrOutput struct{ *pulumi.OutputState }

func (PolicyPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyPolicy)(nil)).Elem()
}

func (o PolicyPolicyPtrOutput) ToPolicyPolicyPtrOutput() PolicyPolicyPtrOutput {
	return o
}

func (o PolicyPolicyPtrOutput) ToPolicyPolicyPtrOutputWithContext(ctx context.Context) PolicyPolicyPtrOutput {
	return o
}

func (o PolicyPolicyPtrOutput) Elem() PolicyPolicyOutput {
	return o.ApplyT(func(v *PolicyPolicy) PolicyPolicy { return *v }).(PolicyPolicyOutput)
}

// Can either be "exchanges", "queues", or "all".
func (o PolicyPolicyPtrOutput) ApplyTo() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyPolicy) string { return v.ApplyTo }).(pulumi.StringOutput)
}

// Key/value pairs of the policy definition. See the
// RabbitMQ documentation for definition references and examples.
func (o PolicyPolicyPtrOutput) Definition() pulumi.MapOutput {
	return o.ApplyT(func(v PolicyPolicy) map[string]interface{} { return v.Definition }).(pulumi.MapOutput)
}

// A pattern to match an exchange or queue name.
func (o PolicyPolicyPtrOutput) Pattern() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyPolicy) string { return v.Pattern }).(pulumi.StringOutput)
}

// The policy with the greater priority is applied first.
func (o PolicyPolicyPtrOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v PolicyPolicy) int { return v.Priority }).(pulumi.IntOutput)
}

type QueueSettings struct {
	// Additional key/value settings for the queue.
	// All values will be sent to RabbitMQ as a string. If you require non-string
	// values, use `argumentsJson`.
	Arguments map[string]interface{} `pulumi:"arguments"`
	// A nested JSON string which contains additional
	// settings for the queue. This is useful for when the arguments contain
	// non-string values.
	ArgumentsJson *string `pulumi:"argumentsJson"`
	// Whether the queue will self-delete when all
	// consumers have unsubscribed.
	AutoDelete *bool `pulumi:"autoDelete"`
	// Whether the queue survives server restarts.
	// Defaults to `false`.
	Durable *bool `pulumi:"durable"`
}

type QueueSettingsInput interface {
	pulumi.Input

	ToQueueSettingsOutput() QueueSettingsOutput
	ToQueueSettingsOutputWithContext(context.Context) QueueSettingsOutput
}

type QueueSettingsArgs struct {
	// Additional key/value settings for the queue.
	// All values will be sent to RabbitMQ as a string. If you require non-string
	// values, use `argumentsJson`.
	Arguments pulumi.MapInput `pulumi:"arguments"`
	// A nested JSON string which contains additional
	// settings for the queue. This is useful for when the arguments contain
	// non-string values.
	ArgumentsJson pulumi.StringPtrInput `pulumi:"argumentsJson"`
	// Whether the queue will self-delete when all
	// consumers have unsubscribed.
	AutoDelete pulumi.BoolPtrInput `pulumi:"autoDelete"`
	// Whether the queue survives server restarts.
	// Defaults to `false`.
	Durable pulumi.BoolPtrInput `pulumi:"durable"`
}

func (QueueSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueSettings)(nil)).Elem()
}

func (i QueueSettingsArgs) ToQueueSettingsOutput() QueueSettingsOutput {
	return i.ToQueueSettingsOutputWithContext(context.Background())
}

func (i QueueSettingsArgs) ToQueueSettingsOutputWithContext(ctx context.Context) QueueSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueSettingsOutput)
}

func (i QueueSettingsArgs) ToQueueSettingsPtrOutput() QueueSettingsPtrOutput {
	return i.ToQueueSettingsPtrOutputWithContext(context.Background())
}

func (i QueueSettingsArgs) ToQueueSettingsPtrOutputWithContext(ctx context.Context) QueueSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueSettingsOutput).ToQueueSettingsPtrOutputWithContext(ctx)
}

type QueueSettingsPtrInput interface {
	pulumi.Input

	ToQueueSettingsPtrOutput() QueueSettingsPtrOutput
	ToQueueSettingsPtrOutputWithContext(context.Context) QueueSettingsPtrOutput
}

type queueSettingsPtrType QueueSettingsArgs

func QueueSettingsPtr(v *QueueSettingsArgs) QueueSettingsPtrInput {
	return (*queueSettingsPtrType)(v)
}

func (*queueSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueSettings)(nil)).Elem()
}

func (i *queueSettingsPtrType) ToQueueSettingsPtrOutput() QueueSettingsPtrOutput {
	return i.ToQueueSettingsPtrOutputWithContext(context.Background())
}

func (i *queueSettingsPtrType) ToQueueSettingsPtrOutputWithContext(ctx context.Context) QueueSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueSettingsPtrOutput)
}

type QueueSettingsOutput struct{ *pulumi.OutputState }

func (QueueSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueSettings)(nil)).Elem()
}

func (o QueueSettingsOutput) ToQueueSettingsOutput() QueueSettingsOutput {
	return o
}

func (o QueueSettingsOutput) ToQueueSettingsOutputWithContext(ctx context.Context) QueueSettingsOutput {
	return o
}

func (o QueueSettingsOutput) ToQueueSettingsPtrOutput() QueueSettingsPtrOutput {
	return o.ToQueueSettingsPtrOutputWithContext(context.Background())
}

func (o QueueSettingsOutput) ToQueueSettingsPtrOutputWithContext(ctx context.Context) QueueSettingsPtrOutput {
	return o.ApplyT(func(v QueueSettings) *QueueSettings {
		return &v
	}).(QueueSettingsPtrOutput)
}

// Additional key/value settings for the queue.
// All values will be sent to RabbitMQ as a string. If you require non-string
// values, use `argumentsJson`.
func (o QueueSettingsOutput) Arguments() pulumi.MapOutput {
	return o.ApplyT(func(v QueueSettings) map[string]interface{} { return v.Arguments }).(pulumi.MapOutput)
}

// A nested JSON string which contains additional
// settings for the queue. This is useful for when the arguments contain
// non-string values.
func (o QueueSettingsOutput) ArgumentsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueSettings) *string { return v.ArgumentsJson }).(pulumi.StringPtrOutput)
}

// Whether the queue will self-delete when all
// consumers have unsubscribed.
func (o QueueSettingsOutput) AutoDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v QueueSettings) *bool { return v.AutoDelete }).(pulumi.BoolPtrOutput)
}

// Whether the queue survives server restarts.
// Defaults to `false`.
func (o QueueSettingsOutput) Durable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v QueueSettings) *bool { return v.Durable }).(pulumi.BoolPtrOutput)
}

type QueueSettingsPtrOutput struct{ *pulumi.OutputState }

func (QueueSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueSettings)(nil)).Elem()
}

func (o QueueSettingsPtrOutput) ToQueueSettingsPtrOutput() QueueSettingsPtrOutput {
	return o
}

func (o QueueSettingsPtrOutput) ToQueueSettingsPtrOutputWithContext(ctx context.Context) QueueSettingsPtrOutput {
	return o
}

func (o QueueSettingsPtrOutput) Elem() QueueSettingsOutput {
	return o.ApplyT(func(v *QueueSettings) QueueSettings { return *v }).(QueueSettingsOutput)
}

// Additional key/value settings for the queue.
// All values will be sent to RabbitMQ as a string. If you require non-string
// values, use `argumentsJson`.
func (o QueueSettingsPtrOutput) Arguments() pulumi.MapOutput {
	return o.ApplyT(func(v QueueSettings) map[string]interface{} { return v.Arguments }).(pulumi.MapOutput)
}

// A nested JSON string which contains additional
// settings for the queue. This is useful for when the arguments contain
// non-string values.
func (o QueueSettingsPtrOutput) ArgumentsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueSettings) *string { return v.ArgumentsJson }).(pulumi.StringPtrOutput)
}

// Whether the queue will self-delete when all
// consumers have unsubscribed.
func (o QueueSettingsPtrOutput) AutoDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v QueueSettings) *bool { return v.AutoDelete }).(pulumi.BoolPtrOutput)
}

// Whether the queue survives server restarts.
// Defaults to `false`.
func (o QueueSettingsPtrOutput) Durable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v QueueSettings) *bool { return v.Durable }).(pulumi.BoolPtrOutput)
}

type TopicPermissionsPermission struct {
	// The exchange to set the permissions for.
	Exchange string `pulumi:"exchange"`
	// The "read" ACL.
	Read string `pulumi:"read"`
	// The "write" ACL.
	Write string `pulumi:"write"`
}

type TopicPermissionsPermissionInput interface {
	pulumi.Input

	ToTopicPermissionsPermissionOutput() TopicPermissionsPermissionOutput
	ToTopicPermissionsPermissionOutputWithContext(context.Context) TopicPermissionsPermissionOutput
}

type TopicPermissionsPermissionArgs struct {
	// The exchange to set the permissions for.
	Exchange pulumi.StringInput `pulumi:"exchange"`
	// The "read" ACL.
	Read pulumi.StringInput `pulumi:"read"`
	// The "write" ACL.
	Write pulumi.StringInput `pulumi:"write"`
}

func (TopicPermissionsPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicPermissionsPermission)(nil)).Elem()
}

func (i TopicPermissionsPermissionArgs) ToTopicPermissionsPermissionOutput() TopicPermissionsPermissionOutput {
	return i.ToTopicPermissionsPermissionOutputWithContext(context.Background())
}

func (i TopicPermissionsPermissionArgs) ToTopicPermissionsPermissionOutputWithContext(ctx context.Context) TopicPermissionsPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicPermissionsPermissionOutput)
}

type TopicPermissionsPermissionArrayInput interface {
	pulumi.Input

	ToTopicPermissionsPermissionArrayOutput() TopicPermissionsPermissionArrayOutput
	ToTopicPermissionsPermissionArrayOutputWithContext(context.Context) TopicPermissionsPermissionArrayOutput
}

type TopicPermissionsPermissionArray []TopicPermissionsPermissionInput

func (TopicPermissionsPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicPermissionsPermission)(nil)).Elem()
}

func (i TopicPermissionsPermissionArray) ToTopicPermissionsPermissionArrayOutput() TopicPermissionsPermissionArrayOutput {
	return i.ToTopicPermissionsPermissionArrayOutputWithContext(context.Background())
}

func (i TopicPermissionsPermissionArray) ToTopicPermissionsPermissionArrayOutputWithContext(ctx context.Context) TopicPermissionsPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicPermissionsPermissionArrayOutput)
}

type TopicPermissionsPermissionOutput struct{ *pulumi.OutputState }

func (TopicPermissionsPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicPermissionsPermission)(nil)).Elem()
}

func (o TopicPermissionsPermissionOutput) ToTopicPermissionsPermissionOutput() TopicPermissionsPermissionOutput {
	return o
}

func (o TopicPermissionsPermissionOutput) ToTopicPermissionsPermissionOutputWithContext(ctx context.Context) TopicPermissionsPermissionOutput {
	return o
}

// The exchange to set the permissions for.
func (o TopicPermissionsPermissionOutput) Exchange() pulumi.StringOutput {
	return o.ApplyT(func(v TopicPermissionsPermission) string { return v.Exchange }).(pulumi.StringOutput)
}

// The "read" ACL.
func (o TopicPermissionsPermissionOutput) Read() pulumi.StringOutput {
	return o.ApplyT(func(v TopicPermissionsPermission) string { return v.Read }).(pulumi.StringOutput)
}

// The "write" ACL.
func (o TopicPermissionsPermissionOutput) Write() pulumi.StringOutput {
	return o.ApplyT(func(v TopicPermissionsPermission) string { return v.Write }).(pulumi.StringOutput)
}

type TopicPermissionsPermissionArrayOutput struct{ *pulumi.OutputState }

func (TopicPermissionsPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicPermissionsPermission)(nil)).Elem()
}

func (o TopicPermissionsPermissionArrayOutput) ToTopicPermissionsPermissionArrayOutput() TopicPermissionsPermissionArrayOutput {
	return o
}

func (o TopicPermissionsPermissionArrayOutput) ToTopicPermissionsPermissionArrayOutputWithContext(ctx context.Context) TopicPermissionsPermissionArrayOutput {
	return o
}

func (o TopicPermissionsPermissionArrayOutput) Index(i pulumi.IntInput) TopicPermissionsPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TopicPermissionsPermission {
		return vs[0].([]TopicPermissionsPermission)[vs[1].(int)]
	}).(TopicPermissionsPermissionOutput)
}

func init() {
	pulumi.RegisterOutputType(ExchangeSettingsOutput{})
	pulumi.RegisterOutputType(ExchangeSettingsPtrOutput{})
	pulumi.RegisterOutputType(PermissionsPermissionsOutput{})
	pulumi.RegisterOutputType(PermissionsPermissionsPtrOutput{})
	pulumi.RegisterOutputType(PolicyPolicyOutput{})
	pulumi.RegisterOutputType(PolicyPolicyPtrOutput{})
	pulumi.RegisterOutputType(QueueSettingsOutput{})
	pulumi.RegisterOutputType(QueueSettingsPtrOutput{})
	pulumi.RegisterOutputType(TopicPermissionsPermissionOutput{})
	pulumi.RegisterOutputType(TopicPermissionsPermissionArrayOutput{})
}
