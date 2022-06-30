package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	network "github.com/pulumi/pulumi-azure-native/sdk/go/azure/network"
)

// FirewallPolicyFilterRuleActionInput is an input type that accepts FirewallPolicyFilterRuleActionArgs and FirewallPolicyFilterRuleActionOutput values.
// You can construct a concrete instance of `FirewallPolicyFilterRuleActionInput` via:
//
//          FirewallPolicyFilterRuleActionArgs{...}
type FirewallPolicyFilterRuleActionInput interface {
	pulumi.Input

	ToFirewallPolicyFilterRuleActionOutput() FirewallPolicyFilterRuleActionOutput
	ToFirewallPolicyFilterRuleActionOutputWithContext(context.Context) FirewallPolicyFilterRuleActionOutput
}

// An application security group in a resource group.
type FirewallPolicyFilterRuleActionArgs struct {
	// The action type of a Filter rule.
	// The type of action.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (FirewallPolicyFilterRuleActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*network.FirewallPolicyFilterRuleAction)(nil)).Elem()
}

func (i FirewallPolicyFilterRuleActionArgs) ToFirewallPolicyFilterRuleActionOutput() FirewallPolicyFilterRuleActionOutput {
	return i.ToFirewallPolicyFilterRuleActionOutputWithContext(context.Background())
}

func (i FirewallPolicyFilterRuleActionArgs) ToFirewallPolicyFilterRuleActionOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyFilterRuleActionOutput)
}

func (i FirewallPolicyFilterRuleActionArgs) ToFirewallPolicyFilterRuleActionPtrOutput() FirewallPolicyFilterRuleActionPtrOutput {
	return i.ToFirewallPolicyFilterRuleActionPtrOutputWithContext(context.Background())
}

func (i FirewallPolicyFilterRuleActionArgs) ToFirewallPolicyFilterRuleActionPtrOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyFilterRuleActionOutput).ToFirewallPolicyFilterRuleActionPtrOutputWithContext(ctx)
}

// FirewallPolicyFilterRuleActionPtrInput is an input type that accepts FirewallPolicyFilterRuleActionArgs, FirewallPolicyFilterRuleActionPtr and FirewallPolicyFilterRuleActionPtrOutput values.
// You can construct a concrete instance of `FirewallPolicyFilterRuleActionPtrInput` via:
//
//          FirewallPolicyFilterRuleActionArgs{...}
//
//  or:
//
//          nil
type FirewallPolicyFilterRuleActionPtrInput interface {
	pulumi.Input

	ToFirewallPolicyFilterRuleActionPtrOutput() FirewallPolicyFilterRuleActionPtrOutput
	ToFirewallPolicyFilterRuleActionPtrOutputWithContext(context.Context) FirewallPolicyFilterRuleActionPtrOutput
}

type FirewallPolicyFilterRuleActionPtrType FirewallPolicyFilterRuleActionArgs

func FirewallPolicyFilterRuleActionPtr(v *FirewallPolicyFilterRuleActionArgs) FirewallPolicyFilterRuleActionPtrInput {
	return (*FirewallPolicyFilterRuleActionPtrType)(v)
}

func (*FirewallPolicyFilterRuleActionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**network.FirewallPolicyFilterRuleAction)(nil)).Elem()
}

func (i *FirewallPolicyFilterRuleActionPtrType) ToFirewallPolicyFilterRuleActionPtrOutput() FirewallPolicyFilterRuleActionPtrOutput {
	return i.ToFirewallPolicyFilterRuleActionPtrOutputWithContext(context.Background())
}

func (i *FirewallPolicyFilterRuleActionPtrType) ToFirewallPolicyFilterRuleActionPtrOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyFilterRuleActionPtrOutput)
}

// FirewallPolicyFilterRuleActionArrayInput is an input type that accepts FirewallPolicyFilterRuleActionArray and FirewallPolicyFilterRuleActionArrayOutput values.
// You can construct a concrete instance of `FirewallPolicyFilterRuleActionArrayInput` via:
//
//          FirewallPolicyFilterRuleActionArray{ FirewallPolicyFilterRuleActionArgs{...} }
type FirewallPolicyFilterRuleActionArrayInput interface {
	pulumi.Input

	ToFirewallPolicyFilterRuleActionArrayOutput() FirewallPolicyFilterRuleActionArrayOutput
	ToFirewallPolicyFilterRuleActionArrayOutputWithContext(context.Context) FirewallPolicyFilterRuleActionArrayOutput
}

type FirewallPolicyFilterRuleActionArray []FirewallPolicyFilterRuleActionInput

func (FirewallPolicyFilterRuleActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]network.FirewallPolicyFilterRuleAction)(nil)).Elem()
}

func (i FirewallPolicyFilterRuleActionArray) ToFirewallPolicyFilterRuleActionArrayOutput() FirewallPolicyFilterRuleActionArrayOutput {
	return i.ToFirewallPolicyFilterRuleActionArrayOutputWithContext(context.Background())
}

func (i FirewallPolicyFilterRuleActionArray) ToFirewallPolicyFilterRuleActionArrayOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyFilterRuleActionArrayOutput)
}

// An application security group in a resource group.
type FirewallPolicyFilterRuleActionOutput struct{ *pulumi.OutputState }

func (FirewallPolicyFilterRuleActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*network.FirewallPolicyFilterRuleAction)(nil)).Elem()
}

func (o FirewallPolicyFilterRuleActionOutput) ToFirewallPolicyFilterRuleActionOutput() FirewallPolicyFilterRuleActionOutput {
	return o
}

func (o FirewallPolicyFilterRuleActionOutput) ToFirewallPolicyFilterRuleActionOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleActionOutput {
	return o
}

// Resource ID.
func (o FirewallPolicyFilterRuleActionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v network.FirewallPolicyFilterRuleAction) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o FirewallPolicyFilterRuleActionOutput) ToFirewallPolicyFilterRuleActionPtrOutput() FirewallPolicyFilterRuleActionPtrOutput {
	return o.ToFirewallPolicyFilterRuleActionPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyFilterRuleActionOutput) ToFirewallPolicyFilterRuleActionPtrOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v network.FirewallPolicyFilterRuleAction) *network.FirewallPolicyFilterRuleAction {
		return &v
	}).(FirewallPolicyFilterRuleActionPtrOutput)
}

type FirewallPolicyFilterRuleActionArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicyFilterRuleActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]network.FirewallPolicyFilterRuleAction)(nil)).Elem()
}

func (o FirewallPolicyFilterRuleActionArrayOutput) ToFirewallPolicyFilterRuleActionArrayOutput() FirewallPolicyFilterRuleActionArrayOutput {
	return o
}

func (o FirewallPolicyFilterRuleActionArrayOutput) ToFirewallPolicyFilterRuleActionArrayOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleActionArrayOutput {
	return o
}

func (o FirewallPolicyFilterRuleActionArrayOutput) Index(i pulumi.IntInput) FirewallPolicyFilterRuleActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) network.FirewallPolicyFilterRuleAction {
		return vs[0].([]network.FirewallPolicyFilterRuleAction)[vs[1].(int)]
	}).(FirewallPolicyFilterRuleActionOutput)
}

type FirewallPolicyFilterRuleActionPtrOutput struct{ *pulumi.OutputState }

func (FirewallPolicyFilterRuleActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**network.FirewallPolicyFilterRuleAction)(nil)).Elem()
}

func (o FirewallPolicyFilterRuleActionPtrOutput) ToFirewallPolicyFilterRuleActionPtrOutput() FirewallPolicyFilterRuleActionPtrOutput {
	return o
}

func (o FirewallPolicyFilterRuleActionPtrOutput) ToFirewallPolicyFilterRuleActionPtrOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleActionPtrOutput {
	return o
}

func (o FirewallPolicyFilterRuleActionPtrOutput) Elem() FirewallPolicyFilterRuleActionOutput {
	return o.ApplyT(func(v *network.FirewallPolicyFilterRuleAction) network.FirewallPolicyFilterRuleAction {
		if v != nil {
			return *v
		}
		var ret network.FirewallPolicyFilterRuleAction
		return ret
	}).(FirewallPolicyFilterRuleActionOutput)
}

// Type of the filter rule action.
func (o FirewallPolicyFilterRuleActionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *network.FirewallPolicyFilterRuleAction) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyFilterRuleActionOutput{})
	pulumi.RegisterOutputType(FirewallPolicyFilterRuleActionPtrOutput{})
}
