// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package nested

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Baz struct {
	Hello *string `pulumi:"hello"`
	World *string `pulumi:"world"`
}

// BazInput is an input type that accepts BazArgs and BazOutput values.
// You can construct a concrete instance of `BazInput` via:
//
//          BazArgs{...}
type BazInput interface {
	pulumi.Input

	ToBazOutput() BazOutput
	ToBazOutputWithContext(context.Context) BazOutput
}

type BazArgs struct {
	Hello pulumi.StringPtrInput `pulumi:"hello"`
	World pulumi.StringPtrInput `pulumi:"world"`
}

func (BazArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Baz)(nil)).Elem()
}

func (i BazArgs) ToBazOutput() BazOutput {
	return i.ToBazOutputWithContext(context.Background())
}

func (i BazArgs) ToBazOutputWithContext(ctx context.Context) BazOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BazOutput)
}

func (i BazArgs) ToBazPtrOutput() BazPtrOutput {
	return i.ToBazPtrOutputWithContext(context.Background())
}

func (i BazArgs) ToBazPtrOutputWithContext(ctx context.Context) BazPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BazOutput).ToBazPtrOutputWithContext(ctx)
}

// BazPtrInput is an input type that accepts BazArgs, BazPtr and BazPtrOutput values.
// You can construct a concrete instance of `BazPtrInput` via:
//
//          BazArgs{...}
//
//  or:
//
//          nil
type BazPtrInput interface {
	pulumi.Input

	ToBazPtrOutput() BazPtrOutput
	ToBazPtrOutputWithContext(context.Context) BazPtrOutput
}

type bazPtrType BazArgs

func BazPtr(v *BazArgs) BazPtrInput {
	return (*bazPtrType)(v)
}

func (*bazPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Baz)(nil)).Elem()
}

func (i *bazPtrType) ToBazPtrOutput() BazPtrOutput {
	return i.ToBazPtrOutputWithContext(context.Background())
}

func (i *bazPtrType) ToBazPtrOutputWithContext(ctx context.Context) BazPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BazPtrOutput)
}

type BazOutput struct{ *pulumi.OutputState }

func (BazOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Baz)(nil)).Elem()
}

func (o BazOutput) ToBazOutput() BazOutput {
	return o
}

func (o BazOutput) ToBazOutputWithContext(ctx context.Context) BazOutput {
	return o
}

func (o BazOutput) ToBazPtrOutput() BazPtrOutput {
	return o.ToBazPtrOutputWithContext(context.Background())
}

func (o BazOutput) ToBazPtrOutputWithContext(ctx context.Context) BazPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Baz) *Baz {
		return &v
	}).(BazPtrOutput)
}

func (o BazOutput) Hello() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Baz) *string { return v.Hello }).(pulumi.StringPtrOutput)
}

func (o BazOutput) World() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Baz) *string { return v.World }).(pulumi.StringPtrOutput)
}

type BazPtrOutput struct{ *pulumi.OutputState }

func (BazPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Baz)(nil)).Elem()
}

func (o BazPtrOutput) ToBazPtrOutput() BazPtrOutput {
	return o
}

func (o BazPtrOutput) ToBazPtrOutputWithContext(ctx context.Context) BazPtrOutput {
	return o
}

func (o BazPtrOutput) Elem() BazOutput {
	return o.ApplyT(func(v *Baz) Baz {
		if v != nil {
			return *v
		}
		var ret Baz
		return ret
	}).(BazOutput)
}

func (o BazPtrOutput) Hello() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Baz) *string {
		if v == nil {
			return nil
		}
		return v.Hello
	}).(pulumi.StringPtrOutput)
}

func (o BazPtrOutput) World() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Baz) *string {
		if v == nil {
			return nil
		}
		return v.World
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BazOutput{})
	pulumi.RegisterOutputType(BazPtrOutput{})
}
