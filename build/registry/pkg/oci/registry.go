/*
Copyright (C) 2022 The Falco Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package oci

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"

	"github.com/falcosecurity/plugins/build/registry/pkg/registry"
)

func PrintUpdateStatus(newArtifacts registry.ArtifactsPushStatus, output io.Writer) error {
	bytes, err := json.Marshal(newArtifacts)
	if err != nil {
		return errors.Wrap(err, "error marshaling oci registry push metadata")
	}
	output.Write(bytes)

	return nil
}
