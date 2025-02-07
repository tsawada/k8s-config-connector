{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}AlloyDBBackup{% endblock %}
{% block body %}


<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>AlloyDB for PostgreSQL</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/alloydb/docs/">/alloydb/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v1.projects.locations.backups</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/alloydb/docs/reference/rest/v1/projects.locations.backups">/alloydb/docs/reference/rest/v1/projects.locations.clusters</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpalloydbbackup<br>gcpalloydbbackups<br>alloydbbackup</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>alloydb.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>alloydbbackups.alloydb.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


<tr>
<td>{{product_name_short}} Default Average Reconcile Interval In Seconds</td>
<td>600</td>
</tr>
</tbody>
</table>

## Custom Resource Definition Properties


### Annotations
<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>cnrm.cloud.google.com/state-into-spec</code></td>
    </tr>
</tbody>
</table>


### Spec
#### Schema
```yaml
clusterNameRef:
  external: string
  name: string
  namespace: string
description: string
encryptionConfig:
  kmsKeyName: string
location: string
projectRef:
  external: string
  name: string
  namespace: string
resourceID: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>clusterNameRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>clusterNameRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `name` field of an `AlloyDBCluster` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>clusterNameRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>clusterNameRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>description</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. User-provided description of the backup.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>encryptionConfig</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>encryptionConfig.kmsKeyName</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The fully-qualified resource name of the KMS key. Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME].{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>location</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The location where the alloydb backup should reside.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The project that this resource belongs to.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `name` field of a `Project` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The backupId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>



### Status
#### Schema
```yaml
conditions:
- lastTransitionTime: string
  message: string
  reason: string
  status: string
  type: string
createTime: string
encryptionInfo:
- encryptionType: string
  kmsKeyVersions:
  - string
etag: string
name: string
observedGeneration: integer
reconciling: boolean
state: string
uid: string
updateTime: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>createTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Time the Backup was created in UTC.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>encryptionInfo</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}EncryptionInfo describes the encryption information of a cluster or a backup.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>encryptionInfo[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>encryptionInfo[].encryptionType</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. Type of encryption.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>encryptionInfo[].kmsKeyVersions</code></td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Output only. Cloud KMS key versions that are being used to protect the database or the backup.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>encryptionInfo[].kmsKeyVersions[]</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>etag</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A hash of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>name</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>reconciling</code></td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}If true, indicates that the service is actively updating the resource. This can happen due to user-triggered updates or system actions like failover or maintenance.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>state</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The current state of the backup.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>uid</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>updateTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Time the Backup was updated in UTC.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Typical Use Case
```yaml
# Copyright 2023 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: alloydb.cnrm.cloud.google.com/v1beta1
kind: AlloyDBBackup
metadata:
  name: alloydbbackup-sample
spec:
  clusterNameRef: 
    external: "projects/${PROJECT_ID?}/locations/us-central1/clusters/alloydbbackup-dep"
  location: us-central1
  encryptionConfig:
    kmsKeyNameRef: 
      external: "projects/${PROJECT_ID?}/locations/us-central1/keyRings/alloydbbackup-dep/cryptoKeys/alloydbbackup-dep"
  projectRef:
    external: ${PROJECT_ID?}
---
apiVersion: alloydb.cnrm.cloud.google.com/v1beta1
kind: AlloyDBCluster
metadata:
  name: alloydbbackup-dep
spec:
  location: us-central1
  networkRef: 
    external: projects/${PROJECT_ID?}/global/networks/alloydbbackup-dep
  projectRef:
    external: ${PROJECT_ID?}
---
apiVersion: alloydb.cnrm.cloud.google.com/v1beta1
kind: AlloyDBInstance
metadata:
  name: alloydbbackup-dep
spec:
  clusterRef: 
    external: projects/${PROJECT_ID?}/locations/us-central1/clusters/alloydbbackup-dep
  instanceType: PRIMARY
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeAddress
metadata:
  name: alloydbbackup-dep
spec:
  location: global
  addressType: INTERNAL
  networkRef:
    name: alloydbbackup-dep
  prefixLength: 16
  purpose: VPC_PEERING
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNetwork
metadata:
  name: alloydbbackup-dep
---
apiVersion: iam.cnrm.cloud.google.com/v1beta1
kind: IAMPartialPolicy
metadata:
  name: alloydbbackup-dep
spec:
  resourceRef:
    apiVersion: kms.cnrm.cloud.google.com/v1beta1
    kind: KMSCryptoKey
    name: alloydbbackup-dep
  bindings:
    - role: roles/cloudkms.cryptoKeyEncrypterDecrypter
      members:
        - member: serviceAccount:service-${PROJECT_NUMBER?}@gcp-sa-alloydb.iam.gserviceaccount.com
---
apiVersion: kms.cnrm.cloud.google.com/v1beta1
kind: KMSCryptoKey
metadata:
  labels:
    source: kcc-alloydbbackup-sample
  name: alloydbbackup-dep
spec:
  keyRingRef:
    name: alloydbbackup-dep
---
apiVersion: kms.cnrm.cloud.google.com/v1beta1
kind: KMSKeyRing
metadata:
  name: alloydbbackup-dep
spec:
  location: us-central1
---
apiVersion: servicenetworking.cnrm.cloud.google.com/v1beta1
kind: ServiceNetworkingConnection
metadata:
  name: alloydbbackup-dep
spec:
  networkRef:
    name: alloydbbackup-dep
  reservedPeeringRanges:
  - external: alloydbbackup-dep
  service: servicenetworking.googleapis.com
```


Note: If you have any trouble with instantiating the resource, refer to <a href="/config-connector/docs/troubleshooting">Troubleshoot Config Connector</a>.

{% endblock %}
