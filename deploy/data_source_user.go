package deploy

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/wperron/terraform-deploy-provider/client"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"github_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceUserRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)

	res, err := c.CurrentUser()

	if err != nil {
		return fmt.Errorf("Error getting Current User: %w", err)
	}

	log.Printf("[DEBUG] Received Caller Identity: %s %s", res.ID, res.Name)

	d.SetId(res.ID)
	d.Set("id", res.ID)
	d.Set("name", res.Name)
	d.Set("github_id", fmt.Sprint(res.GitHubID))

	return nil
}