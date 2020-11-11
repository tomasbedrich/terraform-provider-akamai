// TODO

package netstorage

import (
	"context"
	"sync"

	//storage "github.com/akamai/AkamaiOPEN-edgegrid-golang/storage"
	storage "AkamaiOPEN-edgegrid-golang/storage"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/akamai/terraform-provider-akamai/v2/pkg/akamai"
	"github.com/akamai/terraform-provider-akamai/v2/pkg/config"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type (
	provider struct {
		*schema.Provider
	}
)

var (
	once sync.Once

	inst *provider
)

// Subprovider returns a core sub provider
func Subprovider() akamai.Subprovider {
	once.Do(func() {
		inst = &provider{Provider: Provider()}
	})

	return inst
}

// Provider returns the Akamai terraform.Resource provider.
func Provider() *schema.Provider {

	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"netstorage": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem:     config.Options("netstorage"),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"akamai_netstorage_storage_group": dataSourceNetStorageStorageGroup(),
		},
		ResourcesMap: map[string]*schema.Resource{
			// TODO
		},
	}

	return provider
}

type resourceData interface {
	GetOk(string) (interface{}, bool)
	Get(string) interface{}
}

type set interface {
	List() []interface{}
}

func getConfigNetStorageV1Service(d resourceData) (*edgegrid.Config, error) {
	var netstorageConfig edgegrid.Config
	var err error
	if _, ok := d.GetOk("netstorage"); ok {
		config := d.Get("netstorage").(set).List()[0].(map[string]interface{})

		netstorageConfig = edgegrid.Config{
			Host:         config["host"].(string),
			AccessToken:  config["access_token"].(string),
			ClientToken:  config["client_token"].(string),
			ClientSecret: config["client_secret"].(string),
			MaxBody:      config["max_body"].(int),
		}

		storage.Init(netstorageConfig)
		edgegrid.SetupLogging()
		return &netstorageConfig, nil
	}

	edgerc := d.Get("edgerc").(string)
	section := d.Get("netstorage_section").(string)
	if section == "" {
		section = d.Get("config_section").(string)
	}
	netstorageConfig, err = edgegrid.Init(edgerc, section)
	if err != nil {
		return nil, err
	}

	storage.Init(netstorageConfig)
	return &netstorageConfig, nil
}

func (p *provider) Name() string {
	return "netstorage"
}

// ProviderVersion update version string anytime provider adds new features
const ProviderVersion string = "v0.8.3"

func (p *provider) Version() string {
	return ProviderVersion
}

func (p *provider) Schema() map[string]*schema.Schema {
	return p.Provider.Schema
}

func (p *provider) Resources() map[string]*schema.Resource {
	return p.Provider.ResourcesMap
}

func (p *provider) DataSources() map[string]*schema.Resource {
	return p.Provider.DataSourcesMap
}

func (p *provider) Configure(ctx context.Context, log hclog.Logger, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	log.Named(p.Name()).Debug("START Configure")

	cfg, err := getConfigNetStorageV1Service(d)
	if err != nil {
		return nil, nil
	}

	return cfg, nil
}
