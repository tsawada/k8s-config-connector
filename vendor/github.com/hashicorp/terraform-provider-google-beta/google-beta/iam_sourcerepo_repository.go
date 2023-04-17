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
	"regexp"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var SourceRepoRepositoryIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"repository": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: SourceRepoRepositoryDiffSuppress,
	},
}

func SourceRepoRepositoryDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	oldParts := regexp.MustCompile("projects/[^/]+/repos/").Split(old, -1)
	if len(oldParts) == 2 {
		return oldParts[1] == new
	}
	return new == old
}

type SourceRepoRepositoryIamUpdater struct {
	project    string
	repository string
	d          TerraformResourceData
	Config     *Config
}

func SourceRepoRepositoryIamUpdaterProducer(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	if v, ok := d.GetOk("repository"); ok {
		values["repository"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/repos/(?P<repository>.+)", "(?P<repository>.+)"}, d, config, d.Get("repository").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &SourceRepoRepositoryIamUpdater{
		project:    values["project"],
		repository: values["repository"],
		d:          d,
		Config:     config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("repository", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting repository: %s", err)
	}

	return u, nil
}

func SourceRepoRepositoryIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/repos/(?P<repository>.+)", "(?P<repository>.+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &SourceRepoRepositoryIamUpdater{
		project:    values["project"],
		repository: values["repository"],
		d:          d,
		Config:     config,
	}
	if err := d.Set("repository", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting repository: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *SourceRepoRepositoryIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyRepositoryUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := SendRequest(u.Config, "GET", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *SourceRepoRepositoryIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyRepositoryUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *SourceRepoRepositoryIamUpdater) qualifyRepositoryUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{SourceRepoBasePath}}%s:%s", fmt.Sprintf("projects/%s/repos/%s", u.project, u.repository), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *SourceRepoRepositoryIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/repos/%s", u.project, u.repository)
}

func (u *SourceRepoRepositoryIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-sourcerepo-repository-%s", u.GetResourceId())
}

func (u *SourceRepoRepositoryIamUpdater) DescribeResource() string {
	return fmt.Sprintf("sourcerepo repository %q", u.GetResourceId())
}
