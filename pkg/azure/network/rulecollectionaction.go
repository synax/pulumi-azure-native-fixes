package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	network "github.com/pulumi/pulumi-azure-native/sdk/go/azure/network"
)

// FirewallPolicyFilterRuleCollectionActionInput is an input type that accepts FirewallPolicyFilterRuleCollectionActionArgs and FirewallPolicyFilterRuleCollectionActionOutput values.
// You can construct a concrete instance of `FirewallPolicyFilterRuleCollectionActionInput` via:
//
//          FirewallPolicyFilterRuleCollectionActionArgs{...}
type FirewallPolicyFilterRuleCollectionActionInput interface {
	pulumi.Input

	ToFirewallPolicyFilterRuleCollectionActionOutput() FirewallPolicyFilterRuleCollectionActionOutput
	ToFirewallPolicyFilterRuleCollectionActionOutputWithContext(context.Context) FirewallPolicyFilterRuleCollectionActionOutput
}

// An application security group in a resource group.
type FirewallPolicyFilterRuleCollectionActionArgs struct {
	// The action type of a Filter rule.
	// The type of action.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (FirewallPolicyFilterRuleCollectionActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*network.FirewallPolicyFilterRuleCollectionAction)(nil)).Elem()
}

func (i FirewallPolicyFilterRuleCollectionActionArgs) ToFirewallPolicyFilterRuleCollectionActionOutput() FirewallPolicyFilterRuleCollectionActionOutput {
	return i.ToFirewallPolicyFilterRuleCollectionActionOutputWithContext(context.Background())
}

func (i FirewallPolicyFilterRuleCollectionActionArgs) ToFirewallPolicyFilterRuleCollectionActionOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyFilterRuleCollectionActionOutput)
}

func (i FirewallPolicyFilterRuleCollectionActionArgs) ToFirewallPolicyFilterRuleCollectionActionPtrOutput() FirewallPolicyFilterRuleCollectionActionPtrOutput {
	return i.ToFirewallPolicyFilterRuleCollectionActionPtrOutputWithContext(context.Background())
}

func (i FirewallPolicyFilterRuleCollectionActionArgs) ToFirewallPolicyFilterRuleCollectionActionPtrOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyFilterRuleCollectionActionOutput).ToFirewallPolicyFilterRuleCollectionActionPtrOutputWithContext(ctx)
}

// FirewallPolicyFilterRuleCollectionActionPtrInput is an input type that accepts FirewallPolicyFilterRuleCollectionActionArgs, FirewallPolicyFilterRuleCollectionActionPtr and FirewallPolicyFilterRuleCollectionActionPtrOutput values.
// You can construct a concrete instance of `FirewallPolicyFilterRuleCollectionActionPtrInput` via:
//
//          FirewallPolicyFilterRuleCollectionActionArgs{...}
//
//  or:
//
//          nil
type FirewallPolicyFilterRuleCollectionActionPtrInput interface {
	pulumi.Input

	ToFirewallPolicyFilterRuleCollectionActionPtrOutput() FirewallPolicyFilterRuleCollectionActionPtrOutput
	ToFirewallPolicyFilterRuleCollectionActionPtrOutputWithContext(context.Context) FirewallPolicyFilterRuleCollectionActionPtrOutput
}

type FirewallPolicyFilterRuleCollectionActionPtrType FirewallPolicyFilterRuleCollectionActionArgs

func FirewallPolicyFilterRuleCollectionActionPtr(v *FirewallPolicyFilterRuleCollectionActionArgs) FirewallPolicyFilterRuleCollectionActionPtrInput {
	return (*FirewallPolicyFilterRuleCollectionActionPtrType)(v)
}

func (*FirewallPolicyFilterRuleCollectionActionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**network.FirewallPolicyFilterRuleCollectionAction)(nil)).Elem()
}

func (i *FirewallPolicyFilterRuleCollectionActionPtrType) ToFirewallPolicyFilterRuleCollectionActionPtrOutput() FirewallPolicyFilterRuleCollectionActionPtrOutput {
	return i.ToFirewallPolicyFilterRuleCollectionActionPtrOutputWithContext(context.Background())
}

func (i *FirewallPolicyFilterRuleCollectionActionPtrType) ToFirewallPolicyFilterRuleCollectionActionPtrOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyFilterRuleCollectionActionPtrOutput)
}

// FirewallPolicyFilterRuleCollectionActionArrayInput is an input type that accepts FirewallPolicyFilterRuleCollectionActionArray and FirewallPolicyFilterRuleCollectionActionArrayOutput values.
// You can construct a concrete instance of `FirewallPolicyFilterRuleCollectionActionArrayInput` via:
//
//          FirewallPolicyFilterRuleCollectionActionArray{ FirewallPolicyFilterRuleCollectionActionArgs{...} }
type FirewallPolicyFilterRuleCollectionActionArrayInput interface {
	pulumi.Input

	ToFirewallPolicyFilterRuleCollectionActionArrayOutput() FirewallPolicyFilterRuleCollectionActionArrayOutput
	ToFirewallPolicyFilterRuleCollectionActionArrayOutputWithContext(context.Context) FirewallPolicyFilterRuleCollectionActionArrayOutput
}

type FirewallPolicyFilterRuleCollectionActionArray []FirewallPolicyFilterRuleCollectionActionInput

func (FirewallPolicyFilterRuleCollectionActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]network.FirewallPolicyFilterRuleCollectionAction)(nil)).Elem()
}

func (i FirewallPolicyFilterRuleCollectionActionArray) ToFirewallPolicyFilterRuleCollectionActionArrayOutput() FirewallPolicyFilterRuleCollectionActionArrayOutput {
	return i.ToFirewallPolicyFilterRuleCollectionActionArrayOutputWithContext(context.Background())
}

func (i FirewallPolicyFilterRuleCollectionActionArray) ToFirewallPolicyFilterRuleCollectionActionArrayOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyFilterRuleCollectionActionArrayOutput)
}

// An application security group in a resource group.
type FirewallPolicyFilterRuleCollectionActionOutput struct{ *pulumi.OutputState }

func (FirewallPolicyFilterRuleCollectionActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*network.FirewallPolicyFilterRuleCollectionAction)(nil)).Elem()
}

func (o FirewallPolicyFilterRuleCollectionActionOutput) ToFirewallPolicyFilterRuleCollectionActionOutput() FirewallPolicyFilterRuleCollectionActionOutput {
	return o
}

func (o FirewallPolicyFilterRuleCollectionActionOutput) ToFirewallPolicyFilterRuleCollectionActionOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionOutput {
	return o
}

// Resource ID.
func (o FirewallPolicyFilterRuleCollectionActionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v network.FirewallPolicyFilterRuleCollectionAction) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o FirewallPolicyFilterRuleCollectionActionOutput) ToFirewallPolicyFilterRuleCollectionActionPtrOutput() FirewallPolicyFilterRuleCollectionActionPtrOutput {
	return o.ToFirewallPolicyFilterRuleCollectionActionPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyFilterRuleCollectionActionOutput) ToFirewallPolicyFilterRuleCollectionActionPtrOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v network.FirewallPolicyFilterRuleCollectionAction) *network.FirewallPolicyFilterRuleCollectionAction {
		return &v
	}).(FirewallPolicyFilterRuleCollectionActionPtrOutput)
}

type FirewallPolicyFilterRuleCollectionActionArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicyFilterRuleCollectionActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]network.FirewallPolicyFilterRuleCollectionAction)(nil)).Elem()
}

func (o FirewallPolicyFilterRuleCollectionActionArrayOutput) ToFirewallPolicyFilterRuleCollectionActionArrayOutput() FirewallPolicyFilterRuleCollectionActionArrayOutput {
	return o
}

func (o FirewallPolicyFilterRuleCollectionActionArrayOutput) ToFirewallPolicyFilterRuleCollectionActionArrayOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionArrayOutput {
	return o
}

func (o FirewallPolicyFilterRuleCollectionActionArrayOutput) Index(i pulumi.IntInput) FirewallPolicyFilterRuleCollectionActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) network.FirewallPolicyFilterRuleCollectionAction {
		return vs[0].([]network.FirewallPolicyFilterRuleCollectionAction)[vs[1].(int)]
	}).(FirewallPolicyFilterRuleCollectionActionOutput)
}

type FirewallPolicyFilterRuleCollectionActionPtrOutput struct{ *pulumi.OutputState }

func (FirewallPolicyFilterRuleCollectionActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**network.FirewallPolicyFilterRuleCollectionAction)(nil)).Elem()
}

func (o FirewallPolicyFilterRuleCollectionActionPtrOutput) ToFirewallPolicyFilterRuleCollectionActionPtrOutput() FirewallPolicyFilterRuleCollectionActionPtrOutput {
	return o
}

func (o FirewallPolicyFilterRuleCollectionActionPtrOutput) ToFirewallPolicyFilterRuleCollectionActionPtrOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionPtrOutput {
	return o
}

func (o FirewallPolicyFilterRuleCollectionActionPtrOutput) Elem() FirewallPolicyFilterRuleCollectionActionOutput {
	return o.ApplyT(func(v *network.FirewallPolicyFilterRuleCollectionAction) network.FirewallPolicyFilterRuleCollectionAction {
		if v != nil {
			return *v
		}
		var ret network.FirewallPolicyFilterRuleCollectionAction
		return ret
	}).(FirewallPolicyFilterRuleCollectionActionOutput)
}

// Type of the filter rule action.
func (o FirewallPolicyFilterRuleCollectionActionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *network.FirewallPolicyFilterRuleCollectionAction) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyFilterRuleCollectionActionOutput{})
	pulumi.RegisterOutputType(FirewallPolicyFilterRuleCollectionActionPtrOutput{})
}
