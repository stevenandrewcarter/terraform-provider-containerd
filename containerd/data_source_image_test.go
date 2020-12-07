package containerd

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func Test_datasourceImageRead(t *testing.T) {
	diag := dataSourceImageRead(context.Background(), nil, nil)
	if diag == nil {
		t.Error("expected to have a diag result!")
	}
	t.Logf("Results: %v", diag)
	if !diag.HasError() {
		t.Error("should have an error")
	}
}

func TestDataSourceImage(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: "",
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}
