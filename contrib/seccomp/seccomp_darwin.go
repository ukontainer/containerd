// +build darwin

package seccomp

import (
	"github.com/containerd/containerd/oci"
	"github.com/opencontainers/runtime-spec/specs-go"
)

// WithDefaultProfile sets the default seccomp profile to the spec.
// darwin not concerned about seccomp
func WithDefaultProfile() oci.SpecOpts {
	return nil
}

// WithProfile receives the name of a file stored on disk comprising a json
// formatted seccomp profile, as specified by the opencontainers/runtime-spec.
// The profile is read from the file, unmarshaled, and set to the spec.
// darwin not concerned about seccomp
func WithProfile(profile string) oci.SpecOpts {
	return nil
}

// DefaultProfile defines the whitelist for the default seccomp profile.
// darwin not concerned about seccomp
func DefaultProfile(sp *specs.Spec) *specs.LinuxSeccomp {
	return nil
}
