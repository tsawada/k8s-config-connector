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

func TestAccComputeSubnetwork_subnetworkBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkBasicExample(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.network-with-private-secondary-ip-ranges",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "tf-test-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkLoggingConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkLoggingConfigExample(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.subnet-with-logging",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkLoggingConfigExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "subnet-with-logging" {
  name          = "tf-test-log-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id

  log_config {
    aggregation_interval = "INTERVAL_10_MIN"
    flow_sampling        = 0.5
    metadata             = "INCLUDE_ALL_METADATA"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-log-test-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkInternalL7lbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkInternalL7lbExample(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.network-for-l7lb",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkInternalL7lbExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-for-l7lb" {
  provider = google-beta

  name          = "tf-test-l7lb-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.0.0.0/22"
  region        = "us-central1"
  purpose       = "INTERNAL_HTTPS_LOAD_BALANCER"
  role          = "ACTIVE"
  network       = google_compute_network.custom-test.id
}

resource "google_compute_network" "custom-test" {
  provider = google-beta

  name                    = "tf-test-l7lb-test-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkIpv6Example(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkIpv6Example(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.subnetwork-ipv6",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkIpv6Example(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "subnetwork-ipv6" {
  name          = "tf-test-ipv6-test-subnetwork%{random_suffix}"
  
  ip_cidr_range = "10.0.0.0/22"
  region        = "us-west2"
  
  stack_type       = "IPV4_IPV6"
  ipv6_access_type = "EXTERNAL"

  network       = google_compute_network.custom-test.id
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-ipv6-test-network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkInternalIpv6Example(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSubnetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkInternalIpv6Example(context),
			},
			{
				ResourceName:            "google_compute_subnetwork.subnetwork-internal-ipv6",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region"},
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkInternalIpv6Example(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "subnetwork-internal-ipv6" {
  name          = "tf-test-internal-ipv6-test-subnetwork%{random_suffix}"
  
  ip_cidr_range = "10.0.0.0/22"
  region        = "us-west2"
  
  stack_type       = "IPV4_IPV6"
  ipv6_access_type = "INTERNAL"

  network       = google_compute_network.custom-test.id
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-internal-ipv6-test-network%{random_suffix}"
  auto_create_subnetworks = false
  enable_ula_internal_ipv6 = true
}
`, context)
}

func testAccCheckComputeSubnetworkDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_subnetwork" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeSubnetwork still exists at %s", url)
			}
		}

		return nil
	}
}
