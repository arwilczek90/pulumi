// *** WARNING: this file was generated by . ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.MadeupPackage.Codegentest
{
    public static class FuncWithAllOptionalInputs
    {
        /// <summary>
        /// Check codegen of functions with all optional inputs.
        /// </summary>
        public static Task<FuncWithAllOptionalInputsResult> InvokeAsync(FuncWithAllOptionalInputsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<FuncWithAllOptionalInputsResult>("madeup-package:codegentest:funcWithAllOptionalInputs", args ?? new FuncWithAllOptionalInputsArgs(), options.WithVersion());

        public static Output<FuncWithAllOptionalInputsResult> Invoke(FuncWithAllOptionalInputsOutputArgs? args = null, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.A.Box(),
                args.B.Box()
            ).Apply(a => {
                    var args = new FuncWithAllOptionalInputsArgs();
                    a[0].Set(args, nameof(args.A));
                    a[1].Set(args, nameof(args.B));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class FuncWithAllOptionalInputsArgs : Pulumi.InvokeArgs
    {
        [Input("a")]
        public string? A { get; set; }

        [Input("b")]
        public string? B { get; set; }

        public FuncWithAllOptionalInputsArgs()
        {
        }
    }

    public sealed class FuncWithAllOptionalInputsOutputArgs : Pulumi.InvokeArgs
    {
        [Input("a")]
        public string? A { get; set; }

        [Input("b")]
        public string? B { get; set; }

        public FuncWithAllOptionalInputsOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class FuncWithAllOptionalInputsResult
    {
        public readonly string R;

        [OutputConstructor]
        private FuncWithAllOptionalInputsResult(string r)
        {
            R = r;
        }
    }
}
