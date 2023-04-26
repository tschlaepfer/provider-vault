/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this AuthBackend
func (mg *AuthBackend) GetTerraformResourceType() string {
	return "vault_gcp_auth_backend"
}

// GetConnectionDetailsMapping for this AuthBackend
func (tr *AuthBackend) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"credentials": "spec.forProvider.credentialsSecretRef"}
}

// GetObservation of this AuthBackend
func (tr *AuthBackend) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AuthBackend
func (tr *AuthBackend) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AuthBackend
func (tr *AuthBackend) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AuthBackend
func (tr *AuthBackend) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AuthBackend
func (tr *AuthBackend) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AuthBackend using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AuthBackend) LateInitialize(attrs []byte) (bool, error) {
	params := &AuthBackendParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AuthBackend) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this AuthBackendRole
func (mg *AuthBackendRole) GetTerraformResourceType() string {
	return "vault_gcp_auth_backend_role"
}

// GetConnectionDetailsMapping for this AuthBackendRole
func (tr *AuthBackendRole) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AuthBackendRole
func (tr *AuthBackendRole) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AuthBackendRole
func (tr *AuthBackendRole) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AuthBackendRole
func (tr *AuthBackendRole) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AuthBackendRole
func (tr *AuthBackendRole) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AuthBackendRole
func (tr *AuthBackendRole) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AuthBackendRole using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AuthBackendRole) LateInitialize(attrs []byte) (bool, error) {
	params := &AuthBackendRoleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AuthBackendRole) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this SecretBackend
func (mg *SecretBackend) GetTerraformResourceType() string {
	return "vault_gcp_secret_backend"
}

// GetConnectionDetailsMapping for this SecretBackend
func (tr *SecretBackend) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"credentials": "spec.forProvider.credentialsSecretRef"}
}

// GetObservation of this SecretBackend
func (tr *SecretBackend) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SecretBackend
func (tr *SecretBackend) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SecretBackend
func (tr *SecretBackend) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SecretBackend
func (tr *SecretBackend) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SecretBackend
func (tr *SecretBackend) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SecretBackend using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SecretBackend) LateInitialize(attrs []byte) (bool, error) {
	params := &SecretBackendParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SecretBackend) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SecretImpersonatedAccount
func (mg *SecretImpersonatedAccount) GetTerraformResourceType() string {
	return "vault_gcp_secret_impersonated_account"
}

// GetConnectionDetailsMapping for this SecretImpersonatedAccount
func (tr *SecretImpersonatedAccount) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SecretImpersonatedAccount
func (tr *SecretImpersonatedAccount) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SecretImpersonatedAccount
func (tr *SecretImpersonatedAccount) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SecretImpersonatedAccount
func (tr *SecretImpersonatedAccount) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SecretImpersonatedAccount
func (tr *SecretImpersonatedAccount) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SecretImpersonatedAccount
func (tr *SecretImpersonatedAccount) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SecretImpersonatedAccount using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SecretImpersonatedAccount) LateInitialize(attrs []byte) (bool, error) {
	params := &SecretImpersonatedAccountParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SecretImpersonatedAccount) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SecretRoleset
func (mg *SecretRoleset) GetTerraformResourceType() string {
	return "vault_gcp_secret_roleset"
}

// GetConnectionDetailsMapping for this SecretRoleset
func (tr *SecretRoleset) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SecretRoleset
func (tr *SecretRoleset) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SecretRoleset
func (tr *SecretRoleset) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SecretRoleset
func (tr *SecretRoleset) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SecretRoleset
func (tr *SecretRoleset) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SecretRoleset
func (tr *SecretRoleset) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SecretRoleset using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SecretRoleset) LateInitialize(attrs []byte) (bool, error) {
	params := &SecretRolesetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SecretRoleset) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SecretStaticAccount
func (mg *SecretStaticAccount) GetTerraformResourceType() string {
	return "vault_gcp_secret_static_account"
}

// GetConnectionDetailsMapping for this SecretStaticAccount
func (tr *SecretStaticAccount) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SecretStaticAccount
func (tr *SecretStaticAccount) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SecretStaticAccount
func (tr *SecretStaticAccount) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SecretStaticAccount
func (tr *SecretStaticAccount) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SecretStaticAccount
func (tr *SecretStaticAccount) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SecretStaticAccount
func (tr *SecretStaticAccount) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SecretStaticAccount using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SecretStaticAccount) LateInitialize(attrs []byte) (bool, error) {
	params := &SecretStaticAccountParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SecretStaticAccount) GetTerraformSchemaVersion() int {
	return 0
}