// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterAutomatedBackupPolicy struct {
	/* The length of the time window during which a backup can be taken. If a backup does not succeed within this time window, it will be canceled and considered failed.

	The backup window must be at least 5 minutes long. There is no upper bound on the window. If not set, it will default to 1 hour.

	A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". */
	// +optional
	BackupWindow *string `json:"backupWindow,omitempty"`

	/* Whether automated backups are enabled. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	/* EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key). */
	// +optional
	EncryptionConfig *ClusterEncryptionConfig `json:"encryptionConfig,omitempty"`

	/* Labels to apply to backups created using this configuration. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	/* The location where the backup will be stored. Currently, the only supported option is to store the backup in the same region as the cluster. */
	// +optional
	Location *string `json:"location,omitempty"`

	/* Quantity-based Backup retention policy to retain recent backups. Conflicts with 'time_based_retention', both can't be set together. */
	// +optional
	QuantityBasedRetention *ClusterQuantityBasedRetention `json:"quantityBasedRetention,omitempty"`

	/* Time-based Backup retention policy. Conflicts with 'quantity_based_retention', both can't be set together. */
	// +optional
	TimeBasedRetention *ClusterTimeBasedRetention `json:"timeBasedRetention,omitempty"`

	/* Weekly schedule for the Backup. */
	// +optional
	WeeklySchedule *ClusterWeeklySchedule `json:"weeklySchedule,omitempty"`
}

type ClusterContinuousBackupConfig struct {
	/* Whether continuous backup recovery is enabled. If not set, defaults to true. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	/* EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key). */
	// +optional
	EncryptionConfig *ClusterEncryptionConfig `json:"encryptionConfig,omitempty"`

	/* The numbers of days that are eligible to restore from using PITR. To support the entire recovery window, backups and logs are retained for one day more than the recovery window.

	If not set, defaults to 14 days. */
	// +optional
	RecoveryWindowDays *int64 `json:"recoveryWindowDays,omitempty"`
}

type ClusterEncryptionConfig struct {
	/* (Optional) The fully-qualified resource name of the KMS key. Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME]. */
	// +optional
	KmsKeyNameRef *v1alpha1.ResourceRef `json:"kmsKeyNameRef,omitempty"`
}

type ClusterInitialUser struct {
	/* The initial password for the user. */
	Password ClusterPassword `json:"password"`

	/* The database username. */
	// +optional
	User *string `json:"user,omitempty"`
}

type ClusterNetworkConfig struct {
	/* The name of the allocated IP range for the private IP AlloyDB cluster. For example: "google-managed-services-default".
	If set, the instance IPs for this cluster will be created in the allocated range. */
	// +optional
	AllocatedIpRange *string `json:"allocatedIpRange,omitempty"`

	/* (Required) The relative resource name of the VPC network on which
	the instance can be accessed. It is specified in the following form:
	projects/{project}/global/networks/{network_id}. */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`
}

type ClusterPassword struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *ClusterValueFrom `json:"valueFrom,omitempty"`
}

type ClusterQuantityBasedRetention struct {
	/* The number of backups to retain. */
	// +optional
	Count *int64 `json:"count,omitempty"`
}

type ClusterRestoreBackupSource struct {
	/* (Required) The name of the backup that this cluster is restored from. */
	BackupNameRef v1alpha1.ResourceRef `json:"backupNameRef"`
}

type ClusterRestoreContinuousBackupSource struct {
	/* (Required) The name of the source cluster that this cluster is restored from. */
	ClusterRef v1alpha1.ResourceRef `json:"clusterRef"`

	/* Immutable. The point in time that this cluster is restored to, in RFC 3339 format. */
	PointInTime string `json:"pointInTime"`
}

type ClusterSecondaryConfig struct {
	/* Name of the primary cluster must be in the format
	'projects/{project}/locations/{location}/clusters/{cluster_id}' */
	PrimaryClusterNameRef v1alpha1.ResourceRef `json:"primaryClusterNameRef"`
}

type ClusterStartTimes struct {
	/* Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time. */
	// +optional
	Hours *int64 `json:"hours,omitempty"`

	/* Minutes of hour of day. Currently, only the value 0 is supported. */
	// +optional
	Minutes *int64 `json:"minutes,omitempty"`

	/* Fractions of seconds in nanoseconds. Currently, only the value 0 is supported. */
	// +optional
	Nanos *int64 `json:"nanos,omitempty"`

	/* Seconds of minutes of the time. Currently, only the value 0 is supported. */
	// +optional
	Seconds *int64 `json:"seconds,omitempty"`
}

type ClusterTimeBasedRetention struct {
	/* The retention period.
	A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". */
	// +optional
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
}

type ClusterValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.SecretKeyRef `json:"secretKeyRef,omitempty"`
}

type ClusterWeeklySchedule struct {
	/* The days of the week to perform a backup. At least one day of the week must be provided. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]. */
	// +optional
	DaysOfWeek []string `json:"daysOfWeek,omitempty"`

	/* The times during the day to start a backup. At least one start time must be provided. The start times are assumed to be in UTC and to be an exact hour (e.g., 04:00:00). */
	StartTimes []ClusterStartTimes `json:"startTimes"`
}

type AlloyDBClusterSpec struct {
	/* The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default. */
	// +optional
	AutomatedBackupPolicy *ClusterAutomatedBackupPolicy `json:"automatedBackupPolicy,omitempty"`

	/* The type of cluster. If not set, defaults to PRIMARY. Default value: "PRIMARY" Possible values: ["PRIMARY", "SECONDARY"]. */
	// +optional
	ClusterType *string `json:"clusterType,omitempty"`

	/* The continuous backup config for this cluster.

	If no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days. */
	// +optional
	ContinuousBackupConfig *ClusterContinuousBackupConfig `json:"continuousBackupConfig,omitempty"`

	/* Policy to determine if the cluster should be deleted forcefully.
	Deleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster.
	Deleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = "FORCE" otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance. */
	// +optional
	DeletionPolicy *string `json:"deletionPolicy,omitempty"`

	/* User-settable and human-readable display name for the Cluster. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key). */
	// +optional
	EncryptionConfig *ClusterEncryptionConfig `json:"encryptionConfig,omitempty"`

	/* Initial user to setup during cluster creation. */
	// +optional
	InitialUser *ClusterInitialUser `json:"initialUser,omitempty"`

	/* Immutable. The location where the alloydb cluster should reside. */
	Location string `json:"location"`

	/* Metadata related to network configuration. */
	// +optional
	NetworkConfig *ClusterNetworkConfig `json:"networkConfig,omitempty"`

	/* (Required) The relative resource name of the VPC network on which
	the instance can be accessed. It is specified in the following form:
	projects/{project}/global/networks/{network_id}. */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The clusterId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The source when restoring from a backup. Conflicts with 'restore_continuous_backup_source', both can't be set together. */
	// +optional
	RestoreBackupSource *ClusterRestoreBackupSource `json:"restoreBackupSource,omitempty"`

	/* Immutable. The source when restoring via point in time recovery (PITR). Conflicts with 'restore_backup_source', both can't be set together. */
	// +optional
	RestoreContinuousBackupSource *ClusterRestoreContinuousBackupSource `json:"restoreContinuousBackupSource,omitempty"`

	/* Configuration of the secondary cluster for Cross Region Replication. This should be set if and only if the cluster is of type SECONDARY. */
	// +optional
	SecondaryConfig *ClusterSecondaryConfig `json:"secondaryConfig,omitempty"`
}

type ClusterBackupSourceStatus struct {
	/* The name of the backup resource. */
	// +optional
	BackupName *string `json:"backupName,omitempty"`
}

type ClusterContinuousBackupInfoStatus struct {
	/* The earliest restorable time that can be restored to. Output only field. */
	// +optional
	EarliestRestorableTime *string `json:"earliestRestorableTime,omitempty"`

	/* When ContinuousBackup was most recently enabled. Set to null if ContinuousBackup is not enabled. */
	// +optional
	EnabledTime *string `json:"enabledTime,omitempty"`

	/* Output only. The encryption information for the WALs and backups required for ContinuousBackup. */
	// +optional
	EncryptionInfo []ClusterEncryptionInfoStatus `json:"encryptionInfo,omitempty"`

	/* Days of the week on which a continuous backup is taken. Output only field. Ignored if passed into the request. */
	// +optional
	Schedule []string `json:"schedule,omitempty"`
}

type ClusterEncryptionInfoStatus struct {
	/* Output only. Type of encryption. */
	// +optional
	EncryptionType *string `json:"encryptionType,omitempty"`

	/* Output only. Cloud KMS key versions that are being used to protect the database or the backup. */
	// +optional
	KmsKeyVersions []string `json:"kmsKeyVersions,omitempty"`
}

type ClusterMigrationSourceStatus struct {
	/* The host and port of the on-premises instance in host:port format. */
	// +optional
	HostPort *string `json:"hostPort,omitempty"`

	/* Place holder for the external source identifier(e.g DMS job name) that created the cluster. */
	// +optional
	ReferenceId *string `json:"referenceId,omitempty"`

	/* Type of migration source. */
	// +optional
	SourceType *string `json:"sourceType,omitempty"`
}

type AlloyDBClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   AlloyDBCluster's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Cluster created from backup. */
	// +optional
	BackupSource []ClusterBackupSourceStatus `json:"backupSource,omitempty"`

	/* ContinuousBackupInfo describes the continuous backup properties of a cluster. */
	// +optional
	ContinuousBackupInfo []ClusterContinuousBackupInfoStatus `json:"continuousBackupInfo,omitempty"`

	/* The database engine major version. This is an output-only field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation. */
	// +optional
	DatabaseVersion *string `json:"databaseVersion,omitempty"`

	/* EncryptionInfo describes the encryption information of a cluster or a backup. */
	// +optional
	EncryptionInfo []ClusterEncryptionInfoStatus `json:"encryptionInfo,omitempty"`

	/* Cluster created via DMS migration. */
	// +optional
	MigrationSource []ClusterMigrationSourceStatus `json:"migrationSource,omitempty"`

	/* The name of the cluster resource. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The system-generated UID of the resource. */
	// +optional
	Uid *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpalloydbcluster;gcpalloydbclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AlloyDBCluster is the Schema for the alloydb API
// +k8s:openapi-gen=true
type AlloyDBCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlloyDBClusterSpec   `json:"spec,omitempty"`
	Status AlloyDBClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlloyDBClusterList contains a list of AlloyDBCluster
type AlloyDBClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlloyDBCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlloyDBCluster{}, &AlloyDBClusterList{})
}
