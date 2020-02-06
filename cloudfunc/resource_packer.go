package main

import (
	"errors"
	"os"

	"github.com/eclesiomelojunior/terrago/packer"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePacker() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackerCreate,
		Read:   resourcePackerRead,
		Update: resourcePackerUpdate,
		Delete: resourcePackerDelete,

		Schema: map[string]*schema.Schema{
			"function_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"output": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePackerCreate(d *schema.ResourceData, m interface{}) error {
	functionPath := d.Get("function_path").(string)
	output := d.Get("output").(string)

	zipPath, err := createCloudfuncFile(functionPath, output)

	if err != nil {
		return err
	}

	d.SetId(zipPath)

	return resourcePackerRead(d, m)
}

func resourcePackerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePackerUpdate(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("function_path") || d.HasChange("output") {
		zipPath := d.Id()
		if err := packer.RemoveZip(zipPath); err != nil {
			return err
		}
	}

	functionPath := d.Get("function_path").(string)
	output := d.Get("output").(string)

	zipPath, err := createCloudfuncFile(functionPath, output)

	if err != nil {
		return err
	}

	d.SetId(zipPath)

	return resourcePackerRead(d, m)
}

func resourcePackerDelete(d *schema.ResourceData, m interface{}) error {
	zipPath := d.Id()

	if _, err := os.Stat(zipPath); os.IsNotExist(err) {
		d.SetId("")
		return nil
	}

	if err := packer.RemoveZip(zipPath); err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func createCloudfuncFile(functionPath string, output string) (string, error) {
	files := packer.NormalizePath(functionPath)

	if len(files) < 1 {
		return "", errors.New("Cannot pack and non cloud function")
	}

	zipPath, err := packer.ZipFiles(functionPath, output, files)

	if err != nil {
		return "", err
	}

	return zipPath, nil
}
