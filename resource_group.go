package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGroup() *schema.Resource {
        return &schema.Resource{
                Create: resourceGroupCreate,
                Read:   resourceGroupRead,
                Update: resourceGroupUpdate,
                Delete: resourceGroupDelete,

                Schema: map[string]*schema.Schema{
                        "name": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {
        return resourceGroupRead(d, m)
}

func resourceGroupRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceGroupUpdate(d *schema.ResourceData, m interface{}) error {
        return resourceGroupRead(d, m)
}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}
