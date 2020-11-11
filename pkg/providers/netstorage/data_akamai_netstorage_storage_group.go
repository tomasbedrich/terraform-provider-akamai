package netstorage

import (
	"log"

	//storage "github.com/akamai/AkamaiOPEN-edgegrid-golang/storage"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetStorageStorageGroup() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceNetStorageStorageGroupRead,
		Schema: map[string]*schema.Schema{
			// TODO
		},
	}
}

func dataSourceNetStorageStorageGroupRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] dataSourceNetStorageStorageGroupRead")

	// TODO

	return nil
}
