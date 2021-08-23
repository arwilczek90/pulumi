// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ArgFunction(ctx *pulumi.Context, args *ArgFunctionArgs, opts ...pulumi.InvokeOption) (*ArgFunctionResult, error) {
	var rv ArgFunctionResult
	err := ctx.Invoke("example::argFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ArgFunctionArgs struct {
	Name *random.RandomPet `pulumi:"name"`
}

type ArgFunctionResult struct {
	Age *int `pulumi:"age"`
}

func ArgFunctionOutput(ctx *pulumi.Context, args ArgFunctionOutputArgs, opts ...pulumi.InvokeOption) ArgFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ArgFunctionResult, error) {
			args := v.(ArgFunctionArgs)
			r, err := ArgFunction(ctx, &args, opts...)
			return *r, err
		}).(ArgFunctionResultOutput)
}

type ArgFunctionOutputArgs struct {
	Name random.RandomPetInput `pulumi:"name"`
}

func (ArgFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArgFunctionArgs)(nil)).Elem()
}

type ArgFunctionResultOutput struct{ *pulumi.OutputState }

func (ArgFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArgFunctionResult)(nil)).Elem()
}

func (o ArgFunctionResultOutput) ToArgFunctionResultOutput() ArgFunctionResultOutput {
	return o
}

func (o ArgFunctionResultOutput) ToArgFunctionResultOutputWithContext(ctx context.Context) ArgFunctionResultOutput {
	return o
}

func (o ArgFunctionResultOutput) Age() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArgFunctionResult) *int { return v.Age }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ArgFunctionResultOutput{})
}
