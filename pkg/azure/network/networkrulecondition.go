package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	network "github.com/pulumi/pulumi-azure-native/sdk/go/azure/network"
)

// NetworkRuleConditionInput is an input type that accepts NetworkRuleConditionArgs and NetworkRuleConditionOutput values.
// You can construct a concrete instance of `NetworkRuleConditionInput` via:
//
//          NetworkRuleConditionArgs{...}
type NetworkRuleConditionInput interface {
	pulumi.Input

	ToNetworkRuleConditionOutput() NetworkRuleConditionOutput
	ToNetworkRuleConditionOutputWithContext(context.Context) NetworkRuleConditionOutput
}

type NetworkRuleConditionArgs struct {
	// Description of the rule condition.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// List of destination IP addresses or Service Tags.
	DestinationAddresses pulumi.StringArrayInput `pulumi:"destinationAddresses"`
	// List of destination IpGroups for this rule.
	DestinationIpGroups pulumi.StringArrayInput `pulumi:"destinationIpGroups"`
	// List of destination ports.
	DestinationPorts pulumi.StringArrayInput `pulumi:"destinationPorts"`
	// Array of FirewallPolicyRuleConditionNetworkProtocols.
	IpProtocols pulumi.StringArrayInput `pulumi:"ipProtocols"`
	// Name of the rule condition.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Rule Condition Type.
	// Expected value is 'NetworkRuleCondition'.
	RuleConditionType pulumi.StringInput `pulumi:"ruleConditionType"`
	// List of source IP addresses for this rule.
	SourceAddresses pulumi.StringArray `pulumi:"sourceAddresses"`
	// List of source IpGroups for this rule.
	SourceIpGroups pulumi.StringArrayInput `pulumi:"sourceIpGroups"`
}

func (NetworkRuleConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*network.NetworkRuleCondition)(nil)).Elem()
}

func (i NetworkRuleConditionArgs) ToNetworkRuleConditionOutput() NetworkRuleConditionOutput {
	return i.ToNetworkRuleConditionOutputWithContext(context.Background())
}

func (i NetworkRuleConditionArgs) ToNetworkRuleConditionOutputWithContext(ctx context.Context) NetworkRuleConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleConditionOutput)
}

func (i NetworkRuleConditionArgs) ToNetworkRuleConditionPtrOutput() NetworkRuleConditionPtrOutput {
	return i.ToNetworkRuleConditionPtrOutputWithContext(context.Background())
}

func (i NetworkRuleConditionArgs) ToNetworkRuleConditionPtrOutputWithContext(ctx context.Context) NetworkRuleConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleConditionOutput).ToNetworkRuleConditionPtrOutputWithContext(ctx)
}

// NetworkRuleConditionPtrInput is an input type that accepts NetworkRuleConditionArgs, NetworkRuleConditionPtr and NetworkRuleConditionPtrOutput values.
// You can construct a concrete instance of `NetworkRuleConditionPtrInput` via:
//
//          NetworkRuleConditionArgs{...}
//
//  or:
//
//          nil
type NetworkRuleConditionPtrInput interface {
	pulumi.Input

	ToNetworkRuleConditionPtrOutput() NetworkRuleConditionPtrOutput
	ToNetworkRuleConditionPtrOutputWithContext(context.Context) NetworkRuleConditionPtrOutput
}

type NetworkRuleConditionPtrType NetworkRuleConditionArgs

func NetworkRuleConditionPtr(v *NetworkRuleConditionArgs) NetworkRuleConditionPtrInput {
	return (*NetworkRuleConditionPtrType)(v)
}

func (*NetworkRuleConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**network.NetworkRuleCondition)(nil)).Elem()
}

func (i *NetworkRuleConditionPtrType) ToNetworkRuleConditionPtrOutput() NetworkRuleConditionPtrOutput {
	return i.ToNetworkRuleConditionPtrOutputWithContext(context.Background())
}

func (i *NetworkRuleConditionPtrType) ToNetworkRuleConditionPtrOutputWithContext(ctx context.Context) NetworkRuleConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleConditionPtrOutput)
}

// NetworkRuleConditionArrayInput is an input type that accepts NetworkRuleConditionArray and NetworkRuleConditionArrayOutput values.
// You can construct a concrete instance of `NetworkRuleConditionArrayInput` via:
//
//          NetworkRuleConditionArray{ NetworkRuleConditionArgs{...} }
type NetworkRuleConditionArrayInput interface {
	pulumi.Input

	ToNetworkRuleConditionArrayOutput() NetworkRuleConditionArrayOutput
	ToNetworkRuleConditionArrayOutputWithContext(context.Context) NetworkRuleConditionArrayOutput
}

type NetworkRuleConditionArray []NetworkRuleConditionInput

func (NetworkRuleConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]network.NetworkRuleCondition)(nil)).Elem()
}

func (i NetworkRuleConditionArray) ToNetworkRuleConditionArrayOutput() NetworkRuleConditionArrayOutput {
	return i.ToNetworkRuleConditionArrayOutputWithContext(context.Background())
}

func (i NetworkRuleConditionArray) ToNetworkRuleConditionArrayOutputWithContext(ctx context.Context) NetworkRuleConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleConditionArrayOutput)
}

// An application security group in a resource group.
type NetworkRuleConditionOutput struct{ *pulumi.OutputState }

func (NetworkRuleConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*network.NetworkRuleCondition)(nil)).Elem()
}

func (o NetworkRuleConditionOutput) ToNetworkRuleConditionOutput() NetworkRuleConditionOutput {
	return o
}

func (o NetworkRuleConditionOutput) ToNetworkRuleConditionOutputWithContext(ctx context.Context) NetworkRuleConditionOutput {
	return o
}

// Name.
func (o NetworkRuleConditionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v network.NetworkRuleCondition) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Description.
func (o NetworkRuleConditionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v network.NetworkRuleCondition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// RuleConditionType.
func (o NetworkRuleConditionOutput) RuleConditionType() pulumi.StringOutput {
	return o.ApplyT(func(v network.NetworkRuleCondition) string { return v.RuleConditionType }).(pulumi.StringOutput)
}

// DestinationAddresses.
func (o NetworkRuleConditionOutput) DestinationAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRuleCondition) []string { return v.DestinationAddresses }).(pulumi.StringArrayOutput)
}

// DestinationIpGroups.
func (o NetworkRuleConditionOutput) DestinationIpGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRuleCondition) []string { return v.DestinationIpGroups }).(pulumi.StringArrayOutput)
}

// DestinationPorts.
func (o NetworkRuleConditionOutput) DestinationPorts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRuleCondition) []string { return v.DestinationPorts }).(pulumi.StringArrayOutput)
}

// IpProtocols.
func (o NetworkRuleConditionOutput) IpProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRuleCondition) []string { return v.IpProtocols }).(pulumi.StringArrayOutput)
}

// SourceAddresses.
func (o NetworkRuleConditionOutput) SourceAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRuleCondition) []string { return v.SourceAddresses }).(pulumi.StringArrayOutput)
}

// SourceIpGroups.
func (o NetworkRuleConditionOutput) SourceIpGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRuleCondition) []string { return v.SourceIpGroups }).(pulumi.StringArrayOutput)
}

func (o NetworkRuleConditionOutput) ToNetworkRuleConditionPtrOutput() NetworkRuleConditionPtrOutput {
	return o.ToNetworkRuleConditionPtrOutputWithContext(context.Background())
}

func (o NetworkRuleConditionOutput) ToNetworkRuleConditionPtrOutputWithContext(ctx context.Context) NetworkRuleConditionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v network.NetworkRuleCondition) *network.NetworkRuleCondition {
		return &v
	}).(NetworkRuleConditionPtrOutput)
}

type NetworkRuleConditionArrayOutput struct{ *pulumi.OutputState }

func (NetworkRuleConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]network.NetworkRuleCondition)(nil)).Elem()
}

func (o NetworkRuleConditionArrayOutput) ToNetworkRuleConditionArrayOutput() NetworkRuleConditionArrayOutput {
	return o
}

func (o NetworkRuleConditionArrayOutput) ToNetworkRuleConditionArrayOutputWithContext(ctx context.Context) NetworkRuleConditionArrayOutput {
	return o
}

func (o NetworkRuleConditionArrayOutput) Index(i pulumi.IntInput) NetworkRuleConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) network.NetworkRuleCondition {
		return vs[0].([]network.NetworkRuleCondition)[vs[1].(int)]
	}).(NetworkRuleConditionOutput)
}

type NetworkRuleConditionPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**network.NetworkRuleCondition)(nil)).Elem()
}

func (o NetworkRuleConditionPtrOutput) ToNetworkRuleConditionPtrOutput() NetworkRuleConditionPtrOutput {
	return o
}

func (o NetworkRuleConditionPtrOutput) ToNetworkRuleConditionPtrOutputWithContext(ctx context.Context) NetworkRuleConditionPtrOutput {
	return o
}

func (o NetworkRuleConditionPtrOutput) Elem() NetworkRuleConditionOutput {
	return o.ApplyT(func(v *network.NetworkRuleCondition) network.NetworkRuleCondition {
		if v != nil {
			return *v
		}
		var ret network.NetworkRuleCondition
		return ret
	}).(NetworkRuleConditionOutput)
}

// Name of the network rule condition.
func (o NetworkRuleConditionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *network.NetworkRuleCondition) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// Description of the network rule condition.
func (o NetworkRuleConditionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *network.NetworkRuleCondition) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// DestinationIpGroups of the network rule condition.
func (o NetworkRuleConditionPtrOutput) DestinationIpGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRuleCondition) []string {
		if v == nil {
			return []string{}
		}
		return v.DestinationIpGroups
	}).(pulumi.StringArrayOutput)
}

// DestinationAddresses of the network rule condition.
func (o NetworkRuleConditionPtrOutput) DestinationAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRuleCondition) []string {
		if v == nil {
			return []string{}
		}
		return v.DestinationAddresses
	}).(pulumi.StringArrayOutput)
}

// DestinationPorts of the network rule condition.
func (o NetworkRuleConditionPtrOutput) DestinationPorts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRuleCondition) []string {
		if v == nil {
			return []string{}
		}
		return v.DestinationPorts
	}).(pulumi.StringArrayOutput)
}

// IpProtocols of the network rule condition.
func (o NetworkRuleConditionPtrOutput) IpProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRuleCondition) []string {
		if v == nil {
			return []string{}
		}
		return v.IpProtocols
	}).(pulumi.StringArrayOutput)
}

// SourceAddresses of the network rule condition.
func (o NetworkRuleConditionPtrOutput) SourceAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRuleCondition) []string {
		if v == nil {
			return []string{}
		}
		return v.SourceAddresses
	}).(pulumi.StringArrayOutput)
}

// SourceIpGroups of the network rule condition.
func (o NetworkRuleConditionPtrOutput) SourceIpGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRuleCondition) []string {
		if v == nil {
			return []string{}
		}
		return v.SourceIpGroups
	}).(pulumi.StringArrayOutput)
}

// RuleConditionType of the network rule condition.
func (o NetworkRuleConditionPtrOutput) RuleConditionType() pulumi.StringOutput {
	return o.ApplyT(func(v *network.NetworkRuleCondition) string {
		if v == nil {
			return ""
		}
		return v.RuleConditionType
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkRuleConditionOutput{})
	pulumi.RegisterOutputType(NetworkRuleConditionPtrOutput{})
}
