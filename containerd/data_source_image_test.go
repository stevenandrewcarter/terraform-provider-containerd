package containerd

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestDataSourceImage(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dataSourceImageConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.containerd_image.test", "name"),
				),
			},
		},
	})
}

const dataSourceImageConfig = `
data "containerd_image" "test" {
	name = "alpine"
}
`
