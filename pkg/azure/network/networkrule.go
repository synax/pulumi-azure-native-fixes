package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	network "github.com/pulumi/pulumi-azure-native/sdk/go/azure/network"
)

// NetworkRuleInput is an input type that accepts NetworkRuleArgs and NetworkRuleOutput values.
// You can construct a concrete instance of `NetworkRuleInput` via:
//
//          NetworkRuleArgs{...}
type NetworkRuleInput interface {
	pulumi.Input

	ToNetworkRuleOutput() NetworkRuleOutput
	ToNetworkRuleOutputWithContext(context.Context) NetworkRuleOutput
}

type NetworkRuleArgs struct {
	// Description of the rule condition.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// List of destination IP addresses or Service Tags.
	DestinationAddresses pulumi.StringArrayInput `pulumi:"destinationAddresses"`
	// List of destination IpGroups for this rule.
	DestinationIpGroups pulumi.StringArrayInput `pulumi:"destinationIpGroups"`
	// List of destination ports.
	DestinationPorts pulumi.StringArrayInput `pulumi:"destinationPorts"`

	// List of destination FQDNs.
	DestinationFqdns pulumi.StringArrayInput `pulumi:"destinationFqdns"`
	// Array of FirewallPolicyRuleConditionNetworkProtocols.
	IpProtocols pulumi.StringArrayInput `pulumi:"ipProtocols"`
	// Name of the rule condition.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Rule Condition Type.
	// Expected value is 'NetworkRule'.
	RuleType pulumi.StringInput `pulumi:"ruleType"`
	// List of source IP addresses for this rule.
	SourceAddresses pulumi.StringArray `pulumi:"sourceAddresses"`
	// List of source IpGroups for this rule.
	SourceIpGroups pulumi.StringArrayInput `pulumi:"sourceIpGroups"`
}

func (NetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*network.NetworkRule)(nil)).Elem()
}

func (i NetworkRuleArgs) ToNetworkRuleOutput() NetworkRuleOutput {
	return i.ToNetworkRuleOutputWithContext(context.Background())
}

func (i NetworkRuleArgs) ToNetworkRuleOutputWithContext(ctx context.Context) NetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleOutput)
}

func (i NetworkRuleArgs) ToNetworkRulePtrOutput() NetworkRulePtrOutput {
	return i.ToNetworkRulePtrOutputWithContext(context.Background())
}

func (i NetworkRuleArgs) ToNetworkRulePtrOutputWithContext(ctx context.Context) NetworkRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleOutput).ToNetworkRulePtrOutputWithContext(ctx)
}

// NetworkRulePtrInput is an input type that accepts NetworkRuleArgs, NetworkRulePtr and NetworkRulePtrOutput values.
// You can construct a concrete instance of `NetworkRulePtrInput` via:
//
//          NetworkRuleArgs{...}
//
//  or:
//
//          nil
type NetworkRulePtrInput interface {
	pulumi.Input

	ToNetworkRulePtrOutput() NetworkRulePtrOutput
	ToNetworkRulePtrOutputWithContext(context.Context) NetworkRulePtrOutput
}

type NetworkRulePtrType NetworkRuleArgs

func NetworkRulePtr(v *NetworkRuleArgs) NetworkRulePtrInput {
	return (*NetworkRulePtrType)(v)
}

func (*NetworkRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**network.NetworkRule)(nil)).Elem()
}

func (i *NetworkRulePtrType) ToNetworkRulePtrOutput() NetworkRulePtrOutput {
	return i.ToNetworkRulePtrOutputWithContext(context.Background())
}

func (i *NetworkRulePtrType) ToNetworkRulePtrOutputWithContext(ctx context.Context) NetworkRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRulePtrOutput)
}

// NetworkRuleArrayInput is an input type that accepts NetworkRuleArray and NetworkRuleArrayOutput values.
// You can construct a concrete instance of `NetworkRuleArrayInput` via:
//
//          NetworkRuleArray{ NetworkRuleArgs{...} }
type NetworkRuleArrayInput interface {
	pulumi.Input

	ToNetworkRuleArrayOutput() NetworkRuleArrayOutput
	ToNetworkRuleArrayOutputWithContext(context.Context) NetworkRuleArrayOutput
}

type NetworkRuleArray []NetworkRuleInput

func (NetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]network.NetworkRule)(nil)).Elem()
}

func (i NetworkRuleArray) ToNetworkRuleArrayOutput() NetworkRuleArrayOutput {
	return i.ToNetworkRuleArrayOutputWithContext(context.Background())
}

func (i NetworkRuleArray) ToNetworkRuleArrayOutputWithContext(ctx context.Context) NetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleArrayOutput)
}

// An application security group in a resource group.
type NetworkRuleOutput struct{ *pulumi.OutputState }

func (NetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*network.NetworkRule)(nil)).Elem()
}

func (o NetworkRuleOutput) ToNetworkRuleOutput() NetworkRuleOutput {
	return o
}

func (o NetworkRuleOutput) ToNetworkRuleOutputWithContext(ctx context.Context) NetworkRuleOutput {
	return o
}

// Name.
func (o NetworkRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v network.NetworkRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Description.
func (o NetworkRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v network.NetworkRule) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// RuleType.
func (o NetworkRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v network.NetworkRule) string { return v.RuleType }).(pulumi.StringOutput)
}

// DestinationAddresses.
func (o NetworkRuleOutput) DestinationAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRule) []string { return v.DestinationAddresses }).(pulumi.StringArrayOutput)
}

// DestinationFqdns.
func (o NetworkRuleOutput) DestinationFqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRule) []string { return v.DestinationFqdns }).(pulumi.StringArrayOutput)
}

// DestinationIpGroups.
func (o NetworkRuleOutput) DestinationIpGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRule) []string { return v.DestinationIpGroups }).(pulumi.StringArrayOutput)
}

// DestinationPorts.
func (o NetworkRuleOutput) DestinationPorts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRule) []string { return v.DestinationPorts }).(pulumi.StringArrayOutput)
}

// IpProtocols.
func (o NetworkRuleOutput) IpProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRule) []string { return v.IpProtocols }).(pulumi.StringArrayOutput)
}

// SourceAddresses.
func (o NetworkRuleOutput) SourceAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRule) []string { return v.SourceAddresses }).(pulumi.StringArrayOutput)
}

// SourceIpGroups.
func (o NetworkRuleOutput) SourceIpGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v network.NetworkRule) []string { return v.SourceIpGroups }).(pulumi.StringArrayOutput)
}

func (o NetworkRuleOutput) ToNetworkRulePtrOutput() NetworkRulePtrOutput {
	return o.ToNetworkRulePtrOutputWithContext(context.Background())
}

func (o NetworkRuleOutput) ToNetworkRulePtrOutputWithContext(ctx context.Context) NetworkRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v network.NetworkRule) *network.NetworkRule {
		return &v
	}).(NetworkRulePtrOutput)
}

type NetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (NetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]network.NetworkRule)(nil)).Elem()
}

func (o NetworkRuleArrayOutput) ToNetworkRuleArrayOutput() NetworkRuleArrayOutput {
	return o
}

func (o NetworkRuleArrayOutput) ToNetworkRuleArrayOutputWithContext(ctx context.Context) NetworkRuleArrayOutput {
	return o
}

func (o NetworkRuleArrayOutput) Index(i pulumi.IntInput) NetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) network.NetworkRule {
		return vs[0].([]network.NetworkRule)[vs[1].(int)]
	}).(NetworkRuleOutput)
}

type NetworkRulePtrOutput struct{ *pulumi.OutputState }

func (NetworkRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**network.NetworkRule)(nil)).Elem()
}

func (o NetworkRulePtrOutput) ToNetworkRulePtrOutput() NetworkRulePtrOutput {
	return o
}

func (o NetworkRulePtrOutput) ToNetworkRulePtrOutputWithContext(ctx context.Context) NetworkRulePtrOutput {
	return o
}

func (o NetworkRulePtrOutput) Elem() NetworkRuleOutput {
	return o.ApplyT(func(v *network.NetworkRule) network.NetworkRule {
		if v != nil {
			return *v
		}
		var ret network.NetworkRule
		return ret
	}).(NetworkRuleOutput)
}

// Name of the network rule condition.
func (o NetworkRulePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *network.NetworkRule) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// Description of the network rule condition.
func (o NetworkRulePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *network.NetworkRule) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// DestinationIpGroups of the network rule condition.
func (o NetworkRulePtrOutput) DestinationIpGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRule) []string {
		if v == nil {
			return []string{}
		}
		return v.DestinationIpGroups
	}).(pulumi.StringArrayOutput)
}

// DestinationAddresses of the network rule condition.
func (o NetworkRulePtrOutput) DestinationAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRule) []string {
		if v == nil {
			return []string{}
		}
		return v.DestinationAddresses
	}).(pulumi.StringArrayOutput)
}

// DestinationPorts of the network rule condition.
func (o NetworkRulePtrOutput) DestinationPorts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRule) []string {
		if v == nil {
			return []string{}
		}
		return v.DestinationPorts
	}).(pulumi.StringArrayOutput)
}

// IpProtocols of the network rule condition.
func (o NetworkRulePtrOutput) IpProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRule) []string {
		if v == nil {
			return []string{}
		}
		return v.IpProtocols
	}).(pulumi.StringArrayOutput)
}

// SourceAddresses of the network rule condition.
func (o NetworkRulePtrOutput) SourceAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRule) []string {
		if v == nil {
			return []string{}
		}
		return v.SourceAddresses
	}).(pulumi.StringArrayOutput)
}

// DestinationFqdns of the network rule condition.
func (o NetworkRulePtrOutput) DestinationFqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRule) []string {
		if v == nil {
			return []string{}
		}
		return v.DestinationFqdns
	}).(pulumi.StringArrayOutput)
}

// SourceIpGroups of the network rule condition.
func (o NetworkRulePtrOutput) SourceIpGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *network.NetworkRule) []string {
		if v == nil {
			return []string{}
		}
		return v.SourceIpGroups
	}).(pulumi.StringArrayOutput)
}

// RuleType of the network rule condition.
func (o NetworkRulePtrOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v *network.NetworkRule) string {
		if v == nil {
			return ""
		}
		return v.RuleType
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkRuleOutput{})
	pulumi.RegisterOutputType(NetworkRulePtrOutput{})
}
