package vault

import (
	"github.com/upbound/provider-vault/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// ConfigureNamespace configures the namespace resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_namespace", func(r *config.Resource) {
		r.Kind = "VaultNamespace"
		r.ShortGroup = "vault"
	})
	p.AddResourceConfigurator("vault_pki_secret_backend_sign", func(r *config.Resource) {
		r.Kind = "SecretBackendSign"
		r.ShortGroup = "pki"
		r.References["csr"] = config.Reference{
			Type: "SecretBackendIntermediateCertRequest",
			Extractor: common.CsrExtractor,
		}
	})
	p.AddResourceConfigurator("vault_pki_secret_backend_intermediate_set_signed", func(r *config.Resource) {
		r.Kind = "SecretBackendIntermediateSetSigned"
		r.ShortGroup = "pki"
		r.References["certificate"] = config.Reference{
			Type: "SecretBackendSign",
			Extractor: common.CrtExtractor,
		}
	})
}
