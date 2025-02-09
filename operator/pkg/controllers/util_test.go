// Copyright 2023 Google LLC
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

package controllers

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestValidateContainerResourceCustomizationValues(t *testing.T) {
	tests := []struct {
		name                           string
		containerResourceCustomization v1.ResourceRequirements
		containerInManifest            map[string]interface{}
		wantErr                        string
	}{
		{
			name: "valid customization - both limit and request are specified",
			containerResourceCustomization: v1.ResourceRequirements{
				Limits: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("400m"),
					v1.ResourceMemory: resource.MustParse("512Mi"),
				},
				Requests: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("200m"),
					v1.ResourceMemory: resource.MustParse("256Mi"),
				},
			},
		},
		{
			name: "valid customization - only request is specified",
			containerResourceCustomization: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("200m"),
					v1.ResourceMemory: resource.MustParse("256Mi"),
				},
			},
			containerInManifest: map[string]interface{}{
				"resources": map[string]interface{}{
					"limits": map[string]any{
						"cpu":    "400m",
						"memory": "512Mi",
					},
				},
			},
		},
		{
			name: "valid customization - only limit is specified",
			containerResourceCustomization: v1.ResourceRequirements{
				Limits: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("400m"),
					v1.ResourceMemory: resource.MustParse("512Mi"),
				},
			},
			containerInManifest: map[string]interface{}{
				"resources": map[string]interface{}{
					"requests": map[string]any{
						"cpu":    "200m",
						"memory": "256Mi",
					},
				},
			},
		},
		{
			name: "valid customization - only request is specified, limit is not found in manifest",
			containerResourceCustomization: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("200m"),
					v1.ResourceMemory: resource.MustParse("256Mi"),
				},
			},
			containerInManifest: map[string]interface{}{},
		},
		{
			name: "valid customization - only limit is specified, request is not found in manifest",
			containerResourceCustomization: v1.ResourceRequirements{
				Limits: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("400m"),
					v1.ResourceMemory: resource.MustParse("512Mi"),
				},
			},
			containerInManifest: map[string]interface{}{},
		},
		{
			name: "invalid customization - memory limit is less than memory request",
			containerResourceCustomization: v1.ResourceRequirements{
				Limits: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("400m"),
					v1.ResourceMemory: resource.MustParse("256Mi"),
				},
				Requests: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("200m"),
					v1.ResourceMemory: resource.MustParse("512Mi"),
				},
			},
			wantErr: "memory limit 256Mi is less than memory request 512Mi",
		},
		{
			name: "invalid customization - memory limit is less than default memory request in manifest",
			containerResourceCustomization: v1.ResourceRequirements{
				Limits: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("400m"),
					v1.ResourceMemory: resource.MustParse("256Mi"),
				},
			},
			containerInManifest: map[string]interface{}{
				"resources": map[string]interface{}{
					"requests": map[string]any{
						"cpu":    "200m",
						"memory": "512Mi",
					},
				},
			},
			wantErr: "memory limit 256Mi is less than the default memory request 512Mi in the manifest",
		},
		{
			name: "invalid customization - the default memory limit in manifest is less than memory request",
			containerResourceCustomization: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("200m"),
					v1.ResourceMemory: resource.MustParse("1Gi"),
				},
			},
			containerInManifest: map[string]interface{}{
				"resources": map[string]interface{}{
					"limits": map[string]any{
						"cpu":    "400m",
						"memory": "512Mi",
					},
				},
			},
			wantErr: "default memory limit 512Mi in the manifest is less than the memory request 1Gi",
		},
		{
			name: "invalid customization - cpu limit is less than cpu request",
			containerResourceCustomization: v1.ResourceRequirements{
				Limits: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("200m"),
					v1.ResourceMemory: resource.MustParse("512Mi"),
				},
				Requests: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("400m"),
					v1.ResourceMemory: resource.MustParse("256Mi"),
				},
			},
			wantErr: "cpu limit 200m is less than cpu request 400m",
		},
		{
			name: "invalid customization - cpu limit is less than default cpu request in manifest",
			containerResourceCustomization: v1.ResourceRequirements{
				Limits: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("100m"),
					v1.ResourceMemory: resource.MustParse("512Mi"),
				},
			},
			containerInManifest: map[string]interface{}{
				"resources": map[string]interface{}{
					"requests": map[string]any{
						"cpu":    "200m",
						"memory": "256Mi",
					},
				},
			},
			wantErr: "cpu limit 100m is less than the default cpu request 200m in the manifest",
		},
		{
			name: "invalid customization - the default cpu limit in manifest is less than cpu request",
			containerResourceCustomization: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceCPU:    resource.MustParse("400m"),
					v1.ResourceMemory: resource.MustParse("256Mi"),
				},
			},
			containerInManifest: map[string]interface{}{
				"resources": map[string]interface{}{
					"limits": map[string]any{
						"cpu":    "200m",
						"memory": "512Mi",
					},
				},
			},
			wantErr: "default cpu limit 200m in the manifest is less than the cpu request 400m",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateContainerResourceCustomizationValues(tc.containerResourceCustomization, tc.containerInManifest)
			if tc.wantErr == "" && err != nil {
				t.Errorf("expect no error, got %s", err.Error())
			} else if tc.wantErr != "" {
				if err == nil {
					t.Errorf("expect error %s, got nil", tc.wantErr)
				} else if tc.wantErr != err.Error() {
					t.Errorf("unexpected error, want %s, got %s", tc.wantErr, err.Error())
				}
			}
		})
	}
}
