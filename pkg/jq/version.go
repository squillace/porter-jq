package jq

import (
	"github.com/deislabs/porter/pkg/mixin"
	"github.com/deislabs/porter/pkg/porter/version"
	"github.com/squillace/porter-jq/pkg"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	metadata := mixin.Metadata{
		Name: "jq",
		VersionInfo: mixin.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "ralph squillace",
		},
	}
	return version.PrintVersion(m.Context, opts, metadata)
}
