package v1

import (
	"fmt"
	"reflect"
	"strings"

	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
	"helm.sh/helm/v3/pkg/time"

	"github.com/alauda/component-base/regex"
	"github.com/fatih/structs"
	"github.com/ghodss/yaml"
	"github.com/thoas/go-funk"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
)

const (
	// FinalizerName is the finalizer name we append to each HelmRequest resource
	FinalizerName = "captain.cpaas.io"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Release struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReleaseSpec   `json:"spec"`
	Status ReleaseStatus `json:"status"`
}

// ReleaseSpec describes a deployment of a chart, together with the chart
// and the variables used to deploy that chart.
type ReleaseSpec struct {
	// ChartData is the chart that was released.
	ChartData string `json:"chartData,omitempty"`
	// ConfigData is the set of extra Values added to the chart.
	// These values override the default values inside of the chart.
	ConfigData string `json:"configData,omitempty"`
	// ManifestData is the string representation of the rendered template.
	ManifestData string `json:"manifestData,omitempty"`
	// Hooks are all of the hooks declared for this release.
	HooksData string `json:"hooksData,omitempty"`
	// Version is an int which represents the version of the release.
	Version int `json:"version,omitempty"`

	Name string `json:"name,omitempty"`
}

// Info describes release information.

type ReleaseStatus struct {
	// FirstDeployed is when the release was first deployed.
	FirstDeployed time.Time `json:"first_deployed,omitempty"`
	// LastDeployed is when the release was last deployed.
	LastDeployed time.Time `json:"last_deployed,omitempty"`
	// Deleted tracks when this object was deleted.
	Deleted time.Time `json:"deleted,omitempty"`
	// Description is human-friendly "log entry" about this release.
	Description string `json:"Description,omitempty"`
	// Status is the current state of the release
	Status release.Status `json:"status,omitempty"`
	// Contains the rendered templates/NOTES.txt if available
	Notes string `json:"notes,omitempty"`
}

func (in *ReleaseStatus) CopyFromReleaseInfo(info *release.Info) {
	in.Status = info.Status
	in.Deleted = info.Deleted
	in.Description = info.Description
	in.FirstDeployed = info.FirstDeployed
	in.LastDeployed = info.LastDeployed
	in.Notes = info.Notes
}

func (in *ReleaseStatus) ToReleaseInfo() *release.Info {
	var info release.Info

	info.Status = in.Status
	info.Deleted = in.Deleted
	info.Description = in.Description
	info.FirstDeployed = in.FirstDeployed
	info.LastDeployed = in.LastDeployed
	info.Notes = in.Notes
	return &info
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []Release `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ChartRepoList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []ChartRepo `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ChartRepo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ChartRepoSpec   `json:"spec"`
	Status            ChartRepoStatus `json:"status"`
}

type ChartRepoSpec struct {
	// URL is the repo's url
	URL string `json:"url"`
	// Secret contains information about how to auth to this repo
	Secret *v1.SecretReference `json:"secret,omitempty"`
	// new in v1
	Type string `json:"type"`
	// new in v1.if type is Chart, this is optional and it will provide some compatible with v1alpha1
	Source *ChartRepoSource `json:"source"`
}

type ChartRepoStatus struct {
	// Phase ...
	// After create, this phase will be updated to indicate it's sync status
	// If receive update event, and some field in spec changed, sync agagin.
	Phase ChartRepoPhase `json:"phase,omitempty"`
	// Reason is the failed reason
	Reason string `json:"reason,omitempty"`
}

// ChartRepoSource defines how this ChartRepo is generated  from when it's not a normal chart repo.
// For example, users store some charts source on a VCS, it can be used to generate a helm chart repo
type ChartRepoSource struct {
	// vcs url
	URL string `json:"url"`
	// may be root, may be a subdir
	Path string `json:"path"`
}

// ChartRepoType ...
type ChartRepoType string

const (
	// normal http chart repo
	ChartRepoChart ChartRepoType = "Chart"
	// charts on git
	ChartRepoGit ChartRepoType = "Git"
	// charts on svn
	ChartRepoSvn ChartRepoType = "SVN"
)

type ChartRepoPhase string

const (
	// ChartRepoSynced means is successfully recognized by captain
	ChartRepoSynced ChartRepoPhase = "Synced"

	// ChartRepoFailed means captain is unable to retrieve index info from this repo
	ChartRepoFailed ChartRepoPhase = "Failed"

	// ChartRepoPending means this chartrepo is syncing or pending
	ChartRepoPending ChartRepoPhase = "Pending"
)

func (in *ChartRepo) ValidateCreate() error {
	return nil

}

func (in *ChartRepo) ValidateUpdate(old runtime.Object) error {
	klog.V(4).Info("validate chartrepo update: ", in.GetName())

	oldRepo, ok := old.(*ChartRepo)
	if !ok {
		return fmt.Errorf("expect old object to be a %T instead of %T", oldRepo, old)
	}

	if in.Spec.URL != oldRepo.Spec.URL {
		return fmt.Errorf(".spec.url is immutable")
	}
	return nil
}

func (in *ChartRepo) ValidateDelete() error {
	return nil
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Chart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ChartSpec `json:"spec"`
}

type ChartSpec struct {
	Versions []*ChartVersion `json:"versions,omitempty"`
}

type ChartVersion struct {
	repo.ChartVersion
}

func (in *ChartVersion) DeepCopyInto(out *ChartVersion) {
	if in == nil {
		return
	}

	b, err := yaml.Marshal(in.ChartVersion)
	if err != nil {
		return
	}
	var r repo.ChartVersion
	err = yaml.Unmarshal(b, &r)
	if err != nil {
		return
	}

	out.ChartVersion = r
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ChartList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []Chart `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HelmRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelmRequestSpec   `json:"spec"`
	Status HelmRequestStatus `json:"status"`
}

type HelmRequestSpec struct {
	// ClusterName is the cluster where the chart will be installed. If InstallToAllClusters=true,
	// this field will be ignored
	ClusterName string `json:"clusterName,omitempty"`

	// InstallToAllClusters will install this chart to all available clusters, even the cluster was
	// created after this chart. If this field is true, ClusterName will be ignored(useless)
	InstallToAllClusters bool `json:"installToAllClusters,omitempty"`

	// The following 3 are the new fields added
	// TargetClusters defines the target clusters which the chart will be installed
	TargetClusters *TargetClusters `json:"targetClusters,omitempty"`
	// Source defines the source of chart, If this field is set, Chart and Version field will be ignored(useless)
	Source *ChartSource `json:"source,omitempty"`

	// Dependencies is the dependencies of this HelmRequest, it's a list of there names
	// THe dependencies must lives in the same namespace, and each of them must be in Synced status
	// before we sync this HelmRequest
	Dependencies []string `json:"dependencies,omitempty"`

	// ReleaseName is the Release name to be generated, default to HelmRequest.Name. If we want to manually
	// install this chart to multi clusters, we may have different HelmRequest name(with cluster prefix or suffix)
	// and same release name
	ReleaseName string `json:"releaseName,omitempty"`
	Chart       string `json:"chart,omitempty"`
	Version     string `json:"version,omitempty"`
	// Namespace is the namespace where the Release object will be lived in. Notes this should be used with
	// the values defined in the chart， otherwise the install will failed
	Namespace string `json:"namespace,omitempty"`
	// ValuesFrom represents values from ConfigMap/Secret...
	ValuesFrom []ValuesFromSource `json:"valuesFrom,omitempty"`
	// values is a map
	HelmValues `json:",inline"`
}

type TargetClusters struct {
	// ClusterNames store a list of cluster names
	ClusterNames *TargetClusterNames `json:"clusterNames,omitempty"`
	// ClusterLabels store the key value of labels to match clusters
	ClusterLabels *TargetClusterLabels `json:"clusterLabels,omitempty"`
}

type TargetClusterNames struct {
	Names []string `json:"names"`
}

type TargetClusterLabels struct {
	LabelSelector map[string]string `json:"labelSelector"`
}

type ChartSource struct {
	HTTP *ChartSourceHTTP `json:"http,omitempty"`
	OCI  *ChartSourceOCI  `json:"oci,omitempty"`
}

type ChartSourceHTTP struct {
	// URL is the URL of the http(s) endpoint
	URL string `json:"url"`
	// SecretRef A Secret reference, the secret should contain accessKeyId (user name) base64 encoded, and secretKey (password) also base64 encoded
	// +optional
	SecretRef string `json:"secretRef,omitempty"`
}

type ChartSourceOCI struct {
	// Repo is the repo of the oci artifact
	Repo string `json:"repo"`
	// SecretRef A Secret reference, the secret should contain accessKeyId (user name) base64 encoded, and secretKey (password) also base64 encoded
	// +optional
	SecretRef string `json:"secretRef,omitempty"`
}

//ValuesFromSource represents a source of values, only one of it's fields may be set
type ValuesFromSource struct {
	// ConfigMapKeyRef selects a key of a ConfigMap
	ConfigMapKeyRef *v1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
	// SecretKeyRef selects a key of a Secret
	SecretKeyRef *v1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

//HelmValues embeds helm values so we can add deepcopy on it
type HelmValues struct {
	chartutil.Values `json:"values,omitempty"`
}

func (in *HelmValues) DeepCopyInto(out *HelmValues) {
	if in == nil {
		return
	}

	b, err := yaml.Marshal(in.Values)
	if err != nil {
		return
	}
	var values chartutil.Values
	err = yaml.Unmarshal(b, &values)
	if err != nil {
		return
	}
	out.Values = values
}

// HelmRequestPhase is a label for the condition of a HelmRequest at the current time.
type HelmRequestPhase string

// These are the valid statuses of pods.
const (
	HelmRequestSynced HelmRequestPhase = "Synced"

	// HelmRequestPartialSynced means the HelmRequest is partial synced to target clusters
	HelmRequestPartialSynced HelmRequestPhase = "PartialSynced"

	HelmRequestFailed HelmRequestPhase = "Failed"

	// HelmRequestPending is when helm request is syncing...
	HelmRequestPending HelmRequestPhase = "Pending"

	HelmRequestUnknown HelmRequestPhase = "Unknown"

	// HelmRequestUninstalling is telling that current helmrequest should be uninstalled
	HelmRequestUninstalling HelmRequestPhase = "Uninstalling"
)

// HelmRequestConditionType is a valid value for HelmRequestCondition.Type
type HelmRequestConditionType string

// These are valid conditions of HelmRequestConditionType.
const (
	// ConditionReady indicates than this hr is synced.
	ConditionReady HelmRequestConditionType = "Ready"

	// ConditionValidated means target chart has been downloaded, and permission check passed
	ConditionValidated HelmRequestConditionType = "Validated"

	// ConditionInitialized means this helmrequest has been initialized (chart processed)
	ConditionInitialized HelmRequestConditionType = "Initialized"
)

type HelmRequestCondition struct {
	// Type is the type of the condition.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Type HelmRequestConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=HelmRequestConditionType"`
	// Status is the status of the condition.
	// Can be True, False, Unknown.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Status v1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
	// Last time we probed the condition.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty" protobuf:"bytes,3,opt,name=lastProbeTime"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

type HelmRequestStatus struct {
	Phase HelmRequestPhase `json:"phase,omitempty"`
	// LastSpecHash store the has value of the synced spec, if this value not equal to the current one,
	// means we need to do a update for the chart
	LastSpecHash string `json:"lastSpecHash,omitempty"`
	// SyncedClusters will store the synced clusters if InstallToAllClusters is true
	SyncedClusters []string `json:"syncedClusters,omitempty"`

	// The following 1 is the new field added
	// TargetClusterSyncResults will store the chart sync result of every target cluster
	TargetClusterSyncResults map[string]interface{} `json:"targetClusterSyncResults,omitempty"`

	// Notes is the contents from helm (after helm install successfully it will be printed to the console
	Notes string `json:"notes,omitempty"`

	Conditions []HelmRequestCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,2,rep,name=conditions"`

	// Verions is the real version that installed
	Version string `json:"version,omitempty"`

	// Reason will store the reason why the HelmRequest deploy failed
	Reason string `json:"reason,omitempty"`
}

type ClusterSyncResult struct {
	// Name store the name of cluster
	Name string `json:"name,omitempty"`
	// Endpoint store the apiserver's endpoint of the cluster
	Endpoint string `json:"endpoint,omitempty"`
	// Phase store the phase of the chart which installed into current cluster
	Phase HelmRequestPhase `json:"phase,omitempty"`
	// The following 1 is the new field added
	// AppStatus store the status of the application
	AppStatus AppStatus `json:"appStatus,omitempty"`
	// Reason store the reason why the HelmRequest deploy failed
	Reason string `json:"reason,omitempty"`
	// LastUpdateAt store the last update time
	LastUpdateAt metav1.Time `json:"lastUpdateAt,omitempty"`
}

type AppStatus string

const (
	AppPending        AppStatus = "Pending"
	AppPartialRunning AppStatus = "PartialRunning"
	AppFailed         AppStatus = "Failed"
	AppEmpty          AppStatus = "Empty"
	AppRunning        AppStatus = "Running"
	AppStopped        AppStatus = "Stopped"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HelmRequestList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []HelmRequest `json:"items"`
}

// nameRegexError is thin wrapper to create regex validate error
// key is the field name and value is it's value
func (in *HelmRequest) nameRegexError(key, value string) error {
	return regex.DefaultResourceNameRegexError("HelmRequest", in.GetName(), key, value)
}

// Default makes HelmRequest an mutating webhook
// When delete, if error occurs, finalizer is a good options for us to retry and
// record the events.
func (in *HelmRequest) Default() {
	if !in.DeletionTimestamp.IsZero() {
		return
	}

	// If no releaseName applied, use HelmRequest's name
	if in.Spec.ReleaseName == "" {
		in.Spec.ReleaseName = in.GetName()
		klog.Info("use helmrequest name as release name: ", in.GetName())
	}

	// If no namespace applied, use HelmRequest's namespace
	if in.Spec.Namespace == "" {
		in.Spec.Namespace = in.GetNamespace()
		klog.Info("use helmrequest namespace as release namespace: ", in.GetNamespace())
	}

	in.Finalizers = []string{FinalizerName}
	klog.V(4).Info("append finalizers to helmrequest: ", in.GetName())
}

//ValidateCreate implements webhook.Validator
// 1. check filed regex
func (in *HelmRequest) ValidateCreate() error {
	klog.V(4).Info("validate HelmRequest create: ", in.GetName())
	if in.Spec.ClusterName != "" && !regex.IsValidResourceName(in.Spec.ClusterName) {
		return in.nameRegexError(".spec.clusterName", in.Spec.ClusterName)
	}

	if in.Spec.ReleaseName != "" && !regex.IsValidResourceName(in.Spec.ReleaseName) {
		return in.nameRegexError(".spec.releaseName", in.Spec.ReleaseName)
	}

	if len(in.Spec.Dependencies) > 0 {
		for _, name := range in.Spec.Dependencies {
			if !regex.IsValidResourceName(name) {
				return in.nameRegexError(".spec.dependencies.[]", name)
			}
		}
	}

	if len(in.Spec.ValuesFrom) > 0 {
		for _, item := range in.Spec.ValuesFrom {
			if item.ConfigMapKeyRef != nil && item.SecretKeyRef != nil {
				return fmt.Errorf("cannot set configmap ref and secret ref in the same source")
			}
		}
	}

	return nil
}

//ValidateUpdate validate HelmRequest update request
// immutable fields:
// 1. clusterName
// 2. installToAllCluster
// 3. releaseName
// 4. chart
// 5. namespace
func (in *HelmRequest) ValidateUpdate(old runtime.Object) error {
	klog.V(4).Info("validate HelmRequest update: ", in.GetName())

	oldHR, ok := old.(*HelmRequest)
	if !ok {
		return fmt.Errorf("expect old object to be a %T instead of %T", oldHR, old)
	}

	// check chart name
	_, oldChart := ParseChartName(oldHR.Spec.Chart)
	_, newChart := ParseChartName(in.Spec.Chart)

	if oldChart != newChart {
		return fmt.Errorf("chart name cannot be updated after create")
	}

	// check dependency
	if !reflect.DeepEqual(oldHR.Spec.Dependencies, in.Spec.Dependencies) {
		return fmt.Errorf("dependencies cannot be updated after create")
	}

	o := structs.New(oldHR.Spec)
	n := structs.New(in.Spec)

	for _, key := range []string{"ClusterName", "InstallToAllClusters", "ReleaseName", "Namespace"} {
		kind := o.Field(key).Kind().String()
		if kind == "string" {
			if o.Field(key).Value().(string) != n.Field(key).Value().(string) {
				return fmt.Errorf("field .spec.%s can not update after created", key)
			}
		}
		if kind == "bool" {
			if o.Field(key).Value().(bool) != n.Field(key).Value().(bool) {
				return fmt.Errorf("field .spec.%s can not update after created", key)
			}
		}
	}

	if len(in.Spec.ValuesFrom) > 0 {
		for _, item := range in.Spec.ValuesFrom {
			if item.ConfigMapKeyRef != nil && item.SecretKeyRef != nil {
				return fmt.Errorf("cannot set configmap ref and secret ref in the same source")
			}
		}
	}

	return nil

}

func (in *HelmRequest) ValidateDelete() error {
	return nil
}

//IsClusterSynced check if this HelmRequest has been synced to cluster
func (in *HelmRequest) IsClusterSynced(name string) bool {
	if !in.Spec.InstallToAllClusters {
		return name == in.Spec.ClusterName && in.Status.Phase == HelmRequestSynced
	}

	clusters := in.Status.SyncedClusters
	if len(clusters) > 0 {
		return funk.Contains(clusters, name)
	}

	return false
}

// GetReleaseName get release name. If it's empty in spec, use hr's name
func (in *HelmRequest) GetReleaseName() string {
	name := in.GetName()
	if in.Spec.ReleaseName != "" {
		name = in.Spec.ReleaseName
	}
	return name
}

// GetReleaseNamespace get release namespace. If it's not set, use hr's namespace
func (in *HelmRequest) GetReleaseNamespace() string {
	ns := in.GetNamespace()
	if in.Spec.Namespace != "" {
		ns = in.Spec.Namespace
	}
	return ns
}

// ParseChartName is a simple function that parse chart name
func ParseChartName(name string) (repo, chart string) {
	data := strings.Split(name, "/")
	if len(data) == 1 {
		return "", name
	}
	return data[0], data[1]
}
