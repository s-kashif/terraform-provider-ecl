package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDatabaseV1Database_importBasic(t *testing.T) {
	resourceName := "ecl_db_database_v1.basic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckDatabase(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatabaseV1DatabaseDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDatabaseV1DatabaseBasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"region",
				},
			},
		},
	})
}
