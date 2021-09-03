// Copyright 2016-2021, Pulumi Corporation

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

using FluentAssertions;
using NUnit.Framework;
using Moq;

using Pulumi;
using Pulumi.Testing;
using static Pulumi.MadeupPackage.Codegentest.TestHelpers;

using Pulumi.MadeupPackage.Codegentest.Outputs;

namespace Pulumi.MadeupPackage.Codegentest
{
    [TestFixture]
    public class UnitTests
    {
        [Test]
        public async Task FuncWithAllOptionalInputsOutputWorks()
        {
            Func<string,Func<FuncWithAllOptionalInputsOutputArgs?>,Task> check = (
                (expected, args) => Assert
                .Output(() => FuncWithAllOptionalInputs.InvokeOutput(args()).Apply(x => x.R))
                .ResolvesTo(expected)
            );

            await check("a=null b=null", () => null);

            await check("a=null b=null", () => new FuncWithAllOptionalInputsOutputArgs());

            await check("a=my-a b=null", () => new FuncWithAllOptionalInputsOutputArgs
            {
                A = Out("my-a"),
            });

            await check("a=null b=my-b", () => new FuncWithAllOptionalInputsOutputArgs
            {
                B = Out("my-b"),
            });

            await check("a=my-a b=my-b", () => new FuncWithAllOptionalInputsOutputArgs
            {
                A = Out("my-a"),
                B = Out("my-b"),
            });
        }

        [Test]
        public async Task FuncWithDefaultValueOutputWorks()
        {
            Func<string,Func<FuncWithDefaultValueOutputArgs>,Task> check = (
                (expected, args) => Assert
                .Output(() => FuncWithDefaultValue.InvokeOutput(args()).Apply(x => x.R))
                .ResolvesTo(expected)
            );

            // Since A is required, not passing it is an exception.
            // Perhaps this should be rejected by the typechecker
            // instead? Why is A optional statically?
            Func<Task> act = () => check("", () => new FuncWithDefaultValueOutputArgs());
            await act.Should().ThrowAsync<Exception>();

            // Check that default values from the schema work.
            await check("a=my-a b=b-default", () => new FuncWithDefaultValueOutputArgs()
            {
                A = Out("my-a")
            });

            await check("a=my-a b=my-b", () => new FuncWithDefaultValueOutputArgs()
            {
                A = Out("my-a"),
                B = Out("my-b")
            });
        }

        [Test]
        public async Task FuncWithDictParamOutputWorks()
        {
            Func<string,Func<FuncWithDictParamOutputArgs>,Task> check = (
                (expected, args) => Assert
                .Output(() => FuncWithDictParam.InvokeOutput(args()).Apply(x => x.R))
                .ResolvesTo(expected)
            );

            var map = new InputMap<string>();
            map.Add("K1", Out("my-k1"));
            map.Add("K2", Out("my-k2"));

            // Omitted value defaults to empty dict and not null.
            await check("a=[] b=null", () => new FuncWithDictParamOutputArgs());

            await check("a=[K1: my-k1, K2: my-k2] b=null", () => new FuncWithDictParamOutputArgs()
            {
                A = map,
            });

            await check("a=[K1: my-k1, K2: my-k2] b=my-b", () => new FuncWithDictParamOutputArgs()
            {
                A = map,
                B = Out("my-b"),
            });
        }

        [Test]
        public async Task FuncWithListParamOutputWorks()
        {
            Func<string,Func<FuncWithListParamOutputArgs>,Task> check = (
                (expected, args) => Assert
                .Output(() => FuncWithListParam.InvokeOutput(args()).Apply(x => x.R))
                .ResolvesTo(expected)
            );

            var lst = new InputList<string>();
            lst.Add("e1");
            lst.Add("e2");
            lst.Add("e3");

            // Similarly to dicts, omitted value defaults to empty list and not null.
            await check("a=[] b=null", () => new FuncWithListParamOutputArgs());

            await check("a=[e1, e2, e3] b=null", () => new FuncWithListParamOutputArgs()
            {
                A = lst,
            });

            await check("a=[e1, e2, e3] b=my-b", () => new FuncWithListParamOutputArgs()
            {
                A = lst,
                B = Out("my-b"),
            });
        }

        [Test]
        public async Task GetIntegrationRuntimeObjectMetadatumOuputWorks()
        {
            Func<string,Func<GetIntegrationRuntimeObjectMetadatumOutputArgs>,Task> check = (
                (expected, args) => Assert
                .Output(() => GetIntegrationRuntimeObjectMetadatum.InvokeOutput(args()).Apply(x => {
                    var nextLink = x.NextLink ?? "null";
                    var valueRepr = "null";
                    if (x.Value != null)
                    {
                        valueRepr = string.Join(", ", x.Value.Select(e => $"{e}"));
                    }
                    return $"nextLink={nextLink} value=[{valueRepr}]";
                 }))
                .ResolvesTo(expected)
            );

            await check("nextLink=my-next-link value=[factoryName: my-fn, integrationRuntimeName: my-irn, " +
                        "metadataPath: my-mp, resourceGroupName: my-rgn]",
                        () => new GetIntegrationRuntimeObjectMetadatumOutputArgs()
                        {
                            FactoryName = Out("my-fn"),
                            IntegrationRuntimeName = Out("my-irn"),
                            MetadataPath = Out("my-mp"),
                            ResourceGroupName = Out("my-rgn")
                        });
        }

        [Test]
        public async Task TestListStorageAccountsOutputWorks()
        {
            Func<StorageAccountKeyResponse, string> showSAKR = (r) =>
                $"CreationTime={r.CreationTime} KeyName={r.KeyName} Permissions={r.Permissions} Value={r.Value}";

            Func<string,Func<ListStorageAccountKeysOutputArgs>,Task> check = (
                (expected, args) => Assert
                .Output(() => ListStorageAccountKeys.InvokeOutput(args()).Apply(x => {
                    return "[" + string.Join(", ", x.Keys.Select(k => showSAKR(k))) + "]";
                })).ResolvesTo(expected)
            );

            await check("[CreationTime=my-creation-time KeyName=my-key-name Permissions=my-permissions" +
                        " Value=[accountName: my-an, expand: my-expand, resourceGroupName: my-rgn]]",
                        () => new ListStorageAccountKeysOutputArgs()
                        {
                            AccountName = Out("my-an"),
                            ResourceGroupName = Out("my-rgn"),
                            Expand = Out("my-expand")
                        });
        }
    }
}
