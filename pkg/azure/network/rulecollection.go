package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	network "github.com/pulumi/pulumi-azure-native/sdk/go/azure/network"
)

// FirewallPolicyFilterRuleCollectionInput is an input type that accepts FirewallPolicyFilterRuleCollectionArgs and FirewallPolicyFilterRuleCollectionOutput values.
// You can construct a concrete instance of `FirewallPolicyFilterRuleCollectionInput` via:
//
//          FirewallPolicyFilterRuleCollectionArgs{...}
type FirewallPolicyFilterRuleCollectionInput interface {
	pulumi.Input

	ToFirewallPolicyFilterRuleCollectionOutput() FirewallPolicyFilterRuleCollectionOutput
	ToFirewallPolicyFilterRuleCollectionOutputWithContext(context.Context) FirewallPolicyFilterRuleCollectionOutput
}

// An application security group in a resource group.
type FirewallPolicyFilterRuleCollectionArgs struct {
	// The action type of a Filter rule collection.
	Action FirewallPolicyFilterRuleCollectionActionPtrInput `pulumi:"action"`
	// The name of the rule collection.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Priority of the Firewall Policy Rule Collection resource.
	Priority pulumi.IntPtrInput `pulumi:"priority"`
	// The type of the rule collection.
	// Expected value is 'FirewallPolicyFilterRuleCollection'.
	RuleCollectionType pulumi.String `pulumi:"ruleCollectionType"`
	// List of rules included in a rule collection.
	Rules pulumi.ArrayInput `pulumi:"rules"`
}

func (FirewallPolicyFilterRuleCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*network.FirewallPolicyFilterRuleCollection)(nil)).Elem()
}

func (i FirewallPolicyFilterRuleCollectionArgs) ToFirewallPolicyFilterRuleCollectionOutput() FirewallPolicyFilterRuleCollectionOutput {
	return i.ToFirewallPolicyFilterRuleCollectionOutputWithContext(context.Background())
}

func (i FirewallPolicyFilterRuleCollectionArgs) ToFirewallPolicyFilterRuleCollectionOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyFilterRuleCollectionOutput)
}

// FirewallPolicyFilterRuleCollectionArrayInput is an input type that accepts FirewallPolicyFilterRuleCollectionArray and FirewallPolicyFilterRuleCollectionArrayOutput values.
// You can construct a concrete instance of `FirewallPolicyFilterRuleCollectionArrayInput` via:
//
//          FirewallPolicyFilterRuleCollectionArray{ FirewallPolicyFilterRuleCollectionArgs{...} }
type FirewallPolicyFilterRuleCollectionArrayInput interface {
	pulumi.Input

	ToFirewallPolicyFilterRuleCollectionArrayOutput() FirewallPolicyFilterRuleCollectionArrayOutput
	ToFirewallPolicyFilterRuleCollectionArrayOutputWithContext(context.Context) FirewallPolicyFilterRuleCollectionArrayOutput
}

type FirewallPolicyFilterRuleCollectionArray []FirewallPolicyFilterRuleCollectionInput

func (FirewallPolicyFilterRuleCollectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]network.FirewallPolicyFilterRuleCollection)(nil)).Elem()
}

func (i FirewallPolicyFilterRuleCollectionArray) ToFirewallPolicyFilterRuleCollectionArrayOutput() FirewallPolicyFilterRuleCollectionArrayOutput {
	return i.ToFirewallPolicyFilterRuleCollectionArrayOutputWithContext(context.Background())
}

func (i FirewallPolicyFilterRuleCollectionArray) ToFirewallPolicyFilterRuleCollectionArrayOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyFilterRuleCollectionArrayOutput)
}

// An application security group in a resource group.
type FirewallPolicyFilterRuleCollectionOutput struct{ *pulumi.OutputState }

func (FirewallPolicyFilterRuleCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*network.FirewallPolicyFilterRuleCollection)(nil)).Elem()
}

func (o FirewallPolicyFilterRuleCollectionOutput) ToFirewallPolicyFilterRuleCollectionOutput() FirewallPolicyFilterRuleCollectionOutput {
	return o
}

func (o FirewallPolicyFilterRuleCollectionOutput) ToFirewallPolicyFilterRuleCollectionOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionOutput {
	return o
}

// Action.
func (o FirewallPolicyFilterRuleCollectionOutput) Action() FirewallPolicyFilterRuleCollectionActionPtrOutput {
	return o.ApplyT(func(v network.FirewallPolicyFilterRuleCollection) *network.FirewallPolicyFilterRuleCollectionAction {
		return v.Action
	}).(FirewallPolicyFilterRuleCollectionActionPtrOutput)
}

// Resource Name.
func (o FirewallPolicyFilterRuleCollectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v network.FirewallPolicyFilterRuleCollection) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Rule Type.
func (o FirewallPolicyFilterRuleCollectionOutput) RuleCollectionType() pulumi.StringOutput {
	return o.ApplyT(func(v network.FirewallPolicyFilterRuleCollection) string { return v.RuleCollectionType }).(pulumi.StringOutput)
}

// Rule Priority.
func (o FirewallPolicyFilterRuleCollectionOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v network.FirewallPolicyFilterRuleCollection) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

// Rule.
func (o FirewallPolicyFilterRuleCollectionOutput) Rules() pulumi.ArrayOutput {
	return o.ApplyT(func(v network.FirewallPolicyFilterRuleCollection) []interface{} { return v.Rules }).(pulumi.ArrayOutput)
}

type FirewallPolicyFilterRuleCollectionArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicyFilterRuleCollectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]network.FirewallPolicyFilterRuleCollection)(nil)).Elem()
}

func (o FirewallPolicyFilterRuleCollectionArrayOutput) ToFirewallPolicyFilterRuleCollectionArrayOutput() FirewallPolicyFilterRuleCollectionArrayOutput {
	return o
}

func (o FirewallPolicyFilterRuleCollectionArrayOutput) ToFirewallPolicyFilterRuleCollectionArrayOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionArrayOutput {
	return o
}

func (o FirewallPolicyFilterRuleCollectionArrayOutput) Index(i pulumi.IntInput) FirewallPolicyFilterRuleCollectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) network.FirewallPolicyFilterRuleCollection {
		return vs[0].([]network.FirewallPolicyFilterRuleCollection)[vs[1].(int)]
	}).(FirewallPolicyFilterRuleCollectionOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyFilterRuleCollectionOutput{})
}
