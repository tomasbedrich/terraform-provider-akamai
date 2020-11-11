// +build all netstorage

package netstorage

import "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/registry"

func init() {
	registry.RegisterProvider(Subprovider())
}
