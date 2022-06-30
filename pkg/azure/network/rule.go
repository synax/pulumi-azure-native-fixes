package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	network "github.com/pulumi/pulumi-azure-native/sdk/go/azure/network"
)

// FirewallPolicyFilterRuleInput is an input type that accepts FirewallPolicyFilterRuleArgs and FirewallPolicyFilterRuleOutput values.
// You can construct a concrete instance of `FirewallPolicyFilterRuleInput` via:
//
//          FirewallPolicyFilterRuleArgs{...}
type FirewallPolicyFilterRuleInput interface {
	pulumi.Input

	ToFirewallPolicyFilterRuleOutput() FirewallPolicyFilterRuleOutput
	ToFirewallPolicyFilterRuleOutputWithContext(context.Context) FirewallPolicyFilterRuleOutput
}

// An application security group in a resource group.
type FirewallPolicyFilterRuleArgs struct {
	// The action type of a Filter rule.
	Action FirewallPolicyFilterRuleActionPtrInput `pulumi:"action"`
	// The name of the rule.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Priority of the Firewall Policy Rule resource.
	Priority pulumi.IntPtrInput `pulumi:"priority"`
	// Collection of rule conditions used by a rule.
	RuleConditions pulumi.ArrayInput `pulumi:"ruleConditions"`
	// The type of the rule.
	// Expected value is 'FirewallPolicyFilterRule'.
	RuleType pulumi.String `pulumi:"ruleType"`
}

func (FirewallPolicyFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*network.FirewallPolicyFilterRule)(nil)).Elem()
}

func (i FirewallPolicyFilterRuleArgs) ToFirewallPolicyFilterRuleOutput() FirewallPolicyFilterRuleOutput {
	return i.ToFirewallPolicyFilterRuleOutputWithContext(context.Background())
}

func (i FirewallPolicyFilterRuleArgs) ToFirewallPolicyFilterRuleOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyFilterRuleOutput)
}

// FirewallPolicyFilterRuleArrayInput is an input type that accepts FirewallPolicyFilterRuleArray and FirewallPolicyFilterRuleArrayOutput values.
// You can construct a concrete instance of `FirewallPolicyFilterRuleArrayInput` via:
//
//          FirewallPolicyFilterRuleArray{ FirewallPolicyFilterRuleArgs{...} }
type FirewallPolicyFilterRuleArrayInput interface {
	pulumi.Input

	ToFirewallPolicyFilterRuleArrayOutput() FirewallPolicyFilterRuleArrayOutput
	ToFirewallPolicyFilterRuleArrayOutputWithContext(context.Context) FirewallPolicyFilterRuleArrayOutput
}

type FirewallPolicyFilterRuleArray []FirewallPolicyFilterRuleInput

func (FirewallPolicyFilterRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]network.FirewallPolicyFilterRule)(nil)).Elem()
}

func (i FirewallPolicyFilterRuleArray) ToFirewallPolicyFilterRuleArrayOutput() FirewallPolicyFilterRuleArrayOutput {
	return i.ToFirewallPolicyFilterRuleArrayOutputWithContext(context.Background())
}

func (i FirewallPolicyFilterRuleArray) ToFirewallPolicyFilterRuleArrayOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyFilterRuleArrayOutput)
}

// An application security group in a resource group.
type FirewallPolicyFilterRuleOutput struct{ *pulumi.OutputState }

func (FirewallPolicyFilterRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*network.FirewallPolicyFilterRule)(nil)).Elem()
}

func (o FirewallPolicyFilterRuleOutput) ToFirewallPolicyFilterRuleOutput() FirewallPolicyFilterRuleOutput {
	return o
}

func (o FirewallPolicyFilterRuleOutput) ToFirewallPolicyFilterRuleOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleOutput {
	return o
}

// Action.
func (o FirewallPolicyFilterRuleOutput) Action() FirewallPolicyFilterRuleActionPtrOutput {
	return o.ApplyT(func(v network.FirewallPolicyFilterRule) *network.FirewallPolicyFilterRuleAction { return v.Action }).(FirewallPolicyFilterRuleActionPtrOutput)
}

// Resource Name.
func (o FirewallPolicyFilterRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v network.FirewallPolicyFilterRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Rule Type.
func (o FirewallPolicyFilterRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v network.FirewallPolicyFilterRule) string { return v.RuleType }).(pulumi.StringOutput)
}

// Rule Priority.
func (o FirewallPolicyFilterRuleOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v network.FirewallPolicyFilterRule) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

// Rule Conditions.
func (o FirewallPolicyFilterRuleOutput) RuleConditions() pulumi.ArrayOutput {
	return o.ApplyT(func(v network.FirewallPolicyFilterRule) []interface{} { return v.RuleConditions }).(pulumi.ArrayOutput)
}

type FirewallPolicyFilterRuleArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicyFilterRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]network.FirewallPolicyFilterRule)(nil)).Elem()
}

func (o FirewallPolicyFilterRuleArrayOutput) ToFirewallPolicyFilterRuleArrayOutput() FirewallPolicyFilterRuleArrayOutput {
	return o
}

func (o FirewallPolicyFilterRuleArrayOutput) ToFirewallPolicyFilterRuleArrayOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleArrayOutput {
	return o
}

func (o FirewallPolicyFilterRuleArrayOutput) Index(i pulumi.IntInput) FirewallPolicyFilterRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) network.FirewallPolicyFilterRule {
		return vs[0].([]network.FirewallPolicyFilterRule)[vs[1].(int)]
	}).(FirewallPolicyFilterRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyFilterRuleOutput{})
}
