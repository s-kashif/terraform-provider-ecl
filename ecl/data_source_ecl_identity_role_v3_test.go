package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIdentityV3RoleDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAdminOnly(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccIdentityV3RoleDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIdentityV3DataSourceID("data.ecl_identity_role_v3.role_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_identity_role_v3.role_1", "name", "admin"),
				),
			},
		},
	})
}

func testAccCheckIdentityV3DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find role data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Role data source ID not set")
		}

		return nil
	}
}

const testAccIdentityV3RoleDataSource_basic = `
data "ecl_identity_role_v3" "role_1" {
    name = "admin"
}
`
