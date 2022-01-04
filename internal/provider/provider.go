package provider

import (
	"context"

	"terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/resources/dap"
	"terraform-provider-tessell/internal/resources/database"
	"terraform-provider-tessell/internal/resources/database_backup"
	"terraform-provider-tessell/internal/resources/vpc"
	"terraform-provider-tessell/internal/resources/vpc_peering"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(terraformVersion string) func() *schema.Provider {
	return func() *schema.Provider {
		provider := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"api_address": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("TESSELL_API_ADDRESS", nil),
				},
				"email_id": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("TESSELL_EMAIL_ID", nil),
				},
				"password": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("TESSELL_PASSWORD", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"tessell_dap":              dap.DataSourceDAP(),
				"tessell_daps":             dap.DataSourceDAPs(),
				"tessell_database":         database.DataSourceDatabase(),
				"tessell_databases":        database.DataSourceDatabases(),
				"tessell_database_backup":  database_backup.DataSourceDatabaseBackup(),
				"tessell_database_backups": database_backup.DataSourceDatabaseBackups(),
				"tessell_vpc":              vpc.DataSourceVPC(),
				"tessell_vpcs":             vpc.DataSourceVPCs(),
				"tessell_vpc_peering":      vpc_peering.DataSourceVPCPeering(),
				"tessell_vpc_peerings":     vpc_peering.DataSourceVPCPeerings(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"tessell_dap":             dap.ResourceDAP(),
				"tessell_database":        database.ResourceDatabase(),
				"tessell_database_backup": database_backup.ResourceDatabaseBackup(),
				"tessell_vpc":             vpc.ResourceVPC(),
				"tessell_vpc_peering":     vpc_peering.ResourceVPCPeering(),
			},
		}

		provider.ConfigureContextFunc = configure(terraformVersion, provider)

		return provider
	}
}

func configure(terraformVersion string, provider *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		emailId := d.Get("email_id").(string)
		password := d.Get("password").(string)
		apiAddress := d.Get("api_address").(string)

		var diags diag.Diagnostics

		c, err := client.NewClient(&apiAddress, &emailId, &password)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return c, diags
	}
}
