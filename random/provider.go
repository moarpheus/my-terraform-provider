// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package home

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"random_jokes": dataSourceJokes(),
		},
	}
}
