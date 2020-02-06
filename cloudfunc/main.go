package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})
}

// func main() {
// 	var cloudFuncPath string
// 	flag.StringVar(&cloudFuncPath, "path", "", "a path to cloudfunction directory")
// 	flag.Parse()

// 	files := packer.NormalizePath(cloudFuncPath)

// 	zip, err := packer.ZipFiles(cloudFuncPath, files)

// 	if err != nil {
// 		fmt.Println("Could not zip the files")
// 	}

// 	fmt.Println(zip)
// }
