// *** WARNING: this file was generated by tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codegentest

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Check codegen of functions with all optional inputs.
func FuncWithAllOptionalInputs(ctx *pulumi.Context, args *FuncWithAllOptionalInputsArgs, opts ...pulumi.InvokeOption) (*FuncWithAllOptionalInputsResult, error) {
	var rv FuncWithAllOptionalInputsResult
	err := ctx.Invoke("madeup-package:codegentest:funcWithAllOptionalInputs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type FuncWithAllOptionalInputsArgs struct {
	A *string `pulumi:"a"`
	B *string `pulumi:"b"`
}


type FuncWithAllOptionalInputsResult struct {
	R string `pulumi:"r"`
}


func FuncWithAllOptionalInputsOutput(ctx *pulumi.Context, args FuncWithAllOptionalInputsOutputArgs, opts ...pulumi.InvokeOption) FuncWithAllOptionalInputsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (FuncWithAllOptionalInputsResult, error) {
			args := v.(FuncWithAllOptionalInputsArgs)
			r, err := FuncWithAllOptionalInputs(ctx, &args, opts...)
			return *r, err
		}).(FuncWithAllOptionalInputsResultOutput)
}

type FuncWithAllOptionalInputsOutputArgs struct {
	A pulumi.StringPtrInput `pulumi:"a"`
	B pulumi.StringPtrInput `pulumi:"b"`
}

func (FuncWithAllOptionalInputsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithAllOptionalInputsArgs)(nil)).Elem()
}

type FuncWithAllOptionalInputsResultOutput struct { *pulumi.OutputState }

func (FuncWithAllOptionalInputsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithAllOptionalInputsResult)(nil)).Elem()
}

func (o FuncWithAllOptionalInputsResultOutput) ToFuncWithAllOptionalInputsResultOutput() FuncWithAllOptionalInputsResultOutput {
	return o
}

func (o FuncWithAllOptionalInputsResultOutput) ToFuncWithAllOptionalInputsResultOutputWithContext(ctx context.Context) FuncWithAllOptionalInputsResultOutput {
	return o
}

func (o FuncWithAllOptionalInputsResultOutput) R() pulumi.StringOutput {
	return o.ApplyT(func (v FuncWithAllOptionalInputsResult) string { return v.R }).(pulumi.StringOutput)
}


func init() {
        pulumi.RegisterOutputType(FuncWithAllOptionalInputsResultOutput{})
}

