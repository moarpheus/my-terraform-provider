package random

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io"
	"net/http"
	"strconv"
	"time"
)

func dataSourceRandomJoke() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJokesRead,
		Schema: map[string]*schema.Schema{
			"jokes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"setup": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"punchline": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceJokesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := &http.Client{}
	var data map[string]interface{}

	resp, err := client.Get("https://official-joke-api.appspot.com/random_joke")

	if err != nil {
		return diag.FromErr(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return diag.FromErr(err)
	}

	jokes := make([]map[string]interface{}, 0)
	err = json.Unmarshal(body, &data)

	if err != nil {
		err = fmt.Errorf("cannot unmarshal json of API response: %v", err)
	} else if data["id"] == "" {
		err = fmt.Errorf("missing result key in API response: %v", err)
	}

	if err := d.Set("jokes", jokes); err != nil {
		return diag.FromErr(err)
	}

	outputs := make(map[string]interface{})
	outputs["id"] = data["id"]
	outputs["type"] = data["type"]
	outputs["setup"] = data["setup"]
	outputs["punchline"] = data["punchline"]

	jokes = append(jokes, outputs)

	if err := d.Set("jokes", jokes); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
