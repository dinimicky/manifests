package katib_standalone_ibm

import (
	"github.com/kubeflow/manifests/tests"
	"testing"
)

func TestKustomize(t *testing.T) {
	testCase := &tests.KustomizeTestCase{
		Package:  "../../../../katib/installs/katib-standalone-ibm",
		Expected: "test_data/expected",
	}

	tests.RunTestCase(t, testCase)
}
