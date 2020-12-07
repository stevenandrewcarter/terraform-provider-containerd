package containerd

import (
	"context"
	"errors"
	"strconv"
	"time"

	containerdClient "github.com/containerd/containerd"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceImage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceImageRead,
		Schema: map[string]*schema.Schema{
			"image": &schema.Schema{
				Computed: true,
				Elem:     &schema.Resource{},
			},
		},
	}
}

func dataSourceImageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client, err := containerdClient.New("/run/containerd/containerd.sock")
	if err != nil {
		return diag.FromErr(err)
	}
	defer client.Close()
	imageName := d.Get("name")
	if imageName == nil {
		return diag.FromErr(errors.New("name is required"))
	}
	image, err := client.GetImage(ctx, d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("image", image); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
