package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIdentityV3Role_importBasic(t *testing.T) {
	resourceName := "ecl_identity_role_v3.role_1"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAdminOnly(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIdentityV3UserDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccIdentityV3Role_basic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
