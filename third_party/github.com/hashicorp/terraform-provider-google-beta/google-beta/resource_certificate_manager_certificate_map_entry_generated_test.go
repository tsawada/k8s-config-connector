// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCertificateManagerCertificateMapEntry_certificateManagerCertificateMapEntryFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCertificateManagerCertificateMapEntryDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateManagerCertificateMapEntry_certificateManagerCertificateMapEntryFullExample(context),
			},
			{
				ResourceName:            "google_certificate_manager_certificate_map_entry.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"map"},
			},
		},
	})
}

func testAccCertificateManagerCertificateMapEntry_certificateManagerCertificateMapEntryFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_certificate_manager_certificate_map" "certificate_map" {
  name        = "tf-test-cert-map-entry%{random_suffix}"
  description = "My acceptance test certificate map"
   labels      = {
    "terraform" : true,
    "acc-test"  : true,
  }
}

resource "google_certificate_manager_certificate_map_entry" "default" {
  name        = "tf-test-cert-map-entry%{random_suffix}"
  description = "My acceptance test certificate map entry"
  map = google_certificate_manager_certificate_map.certificate_map.name 
  labels      = {
    "terraform" : true,
    "acc-test"  : true,
  }
  certificates = [google_certificate_manager_certificate.certificate.id]
  matcher = "PRIMARY"
}

resource "google_certificate_manager_certificate" "certificate" {
  name        = "tf-test-cert-map-entry%{random_suffix}"
  description = "The default cert"
  scope       = "DEFAULT"
  managed {
    domains = [
      google_certificate_manager_dns_authorization.instance.domain,
      google_certificate_manager_dns_authorization.instance2.domain,
      ]
    dns_authorizations = [
      google_certificate_manager_dns_authorization.instance.id,
      google_certificate_manager_dns_authorization.instance2.id,
      ]
  }
}


resource "google_certificate_manager_dns_authorization" "instance" {
  name        = "tf-test-dns-auth%{random_suffix}"
  description = "The default dnss"
  domain      = "subdomain%{random_suffix}.hashicorptest.com"
}

resource "google_certificate_manager_dns_authorization" "instance2" {
  name        = "tf-test-dns-auth2%{random_suffix}"
  description = "The default dnss"
  domain      = "subdomain2%{random_suffix}.hashicorptest.com"
}
`, context)
}

func testAccCheckCertificateManagerCertificateMapEntryDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_certificate_manager_certificate_map_entry" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("CertificateManagerCertificateMapEntry still exists at %s", url)
			}
		}

		return nil
	}
}
