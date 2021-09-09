// nolint: lll
package nodejs

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/pkg/v3/codegen/internal/test"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/executable"
)

func TestGeneratePackage(t *testing.T) {
	test.TestSDKCodegen(t, "nodejs", GeneratePackage, typeCheckGeneratedPackage)
}

func typeCheckGeneratedPackage(t *testing.T, pwd string) {
	var err error
	var npm string
	npm, err = executable.FindExecutable("npm")
	require.NoError(t, err)

	cmdOptions := integration.ProgramTestOptions{}
	err = integration.RunCommand(t, "npm install", []string{npm, "install"}, pwd, &cmdOptions)
	require.NoError(t, err)
	err = integration.RunCommand(t, "typecheck ts",
		[]string{npm, "exec", "--yes", "--", "ts-node", "--type-check", "."}, pwd, &cmdOptions)
	require.NoError(t, err)
}

func TestGenerateTypeNames(t *testing.T) {
	test.TestTypeNameCodegen(t, "nodejs", func(pkg *schema.Package) test.TypeNameGeneratorFunc {
		modules, info, err := generateModuleContextMap("test", pkg, nil)
		require.NoError(t, err)

		pkg.Language["nodejs"] = info

		root, ok := modules[""]
		require.True(t, ok)

		return func(t schema.Type) string {
			return root.typeString(t, false, nil)
		}
	})
}
