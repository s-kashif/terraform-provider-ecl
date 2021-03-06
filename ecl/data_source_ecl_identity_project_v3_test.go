package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIdentityV3ProjectDataSource_basic(t *testing.T) {
	projectName := fmt.Sprintf("tf_test_%s", acctest.RandString(5))
	projectDescription := acctest.RandString(20)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAdminOnly(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccIdentityProjectV3DataSource_project(projectName, projectDescription),
			},
			resource.TestStep{
				Config: testAccIdentityProjectV3DataSource_basic(projectName, projectDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIdentityV3ProjectDataSourceID("data.ecl_identity_project_v3.project_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_identity_project_v3.project_1", "name", projectName),
					resource.TestCheckResourceAttr(
						"ecl_identity_project_v3.project_1", "description", projectDescription),
					resource.TestCheckResourceAttr(
						"ecl_identity_project_v3.project_1", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"ecl_identity_project_v3.project_1", "is_domain", "false"),
				),
			},
		},
	})
}

func testAccCheckIdentityV3ProjectDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find project data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Project data source ID not set")
		}

		return nil
	}
}

func testAccIdentityProjectV3DataSource_project(name, description string) string {
	return fmt.Sprintf(`
	resource "ecl_identity_project_v3" "project_1" {
	  name = "%s"
	  description = "%s"
	}
`, name, description)
}

func testAccIdentityProjectV3DataSource_basic(name, description string) string {
	return fmt.Sprintf(`
	%s

	data "ecl_identity_project_v3" "project_1" {
      name = "${ecl_identity_project_v3.project_1.name}"
	}
`, testAccIdentityProjectV3DataSource_project(name, description))
}
