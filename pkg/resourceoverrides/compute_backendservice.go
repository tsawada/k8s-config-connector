// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resourceoverrides

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

var (
	nonReferenceFieldPath = []string{"iap", "oauth2ClientId"}
	referenceFieldPath    = []string{"iap", "oauth2ClientIdRef"}
)

func GetComputeBackendServiceResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "ComputeBackendService",
	}
	// Preserve the legacy non-reference field 'iap.oauth2ClientId' after it is changed to
	// a reference field, 'iap.oauth2ClientIdRef'.
	ro.Overrides = append(ro.Overrides, keepOauth2ClientIdField())
	return ro
}

func keepOauth2ClientIdField() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		if err := PreserveMutuallyExclusiveNonReferenceField(crd, []string{"iap"}, referenceFieldPath[1], nonReferenceFieldPath[1]); err != nil {
			return err
		}
		return nil
	}
	o.PreActuationTransform = func(r *k8s.Resource) error {
		if err := FavorAuthoritativeFieldOverLegacyField(r, nonReferenceFieldPath, referenceFieldPath); err != nil {
			return fmt.Errorf("error handling '%v' and '%v' fields in pre-actuation transformation: %w", strings.Join(nonReferenceFieldPath, "."), strings.Join(referenceFieldPath, "."), err)
		}
		return nil
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource) error {
		if err := PreserveUserSpecifiedLegacyField(original, reconciled, nonReferenceFieldPath...); err != nil {
			return fmt.Errorf("error preserving '%v' in post-actuation transformation: %w", strings.Join(nonReferenceFieldPath, "."), err)
		}
		if err := PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecified(original, reconciled, nonReferenceFieldPath, referenceFieldPath); err != nil {
			return fmt.Errorf("error conditionally pruning '%v' in post-actuation transformation: %w", strings.Join(referenceFieldPath, "."), err)
		}
		return nil
	}
	return o
}
