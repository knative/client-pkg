// Copyright © 2023 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

// Manifest represents plugin metadata
type Manifest struct {
	// Path to external plugin binary. Always empty for inlined plugins.
	Path string `json:"path,omitempty"`

	// Plugin declares its own manifest to be included in Context Sharing feature
	HasManifest bool `json:"hasManifest"`

	// ProducesContextDataKeys is a list of keys for the ContextData that
	// a plugin can produce. Nil or an empty list declares that this
	// plugin is not ContextDataProducer
	ProducesContextDataKeys []string `json:"producesKeys,omitempty"`

	// ConsumesContextDataKeys is a list of keys from a ContextData that a
	// plugin is interested in to consume. Nil or an empty list declares
	// that this plugin is not a ContextDataConsumer
	ConsumesContextDataKeys []string `json:"consumesKeys,omitempty"`
}

// PluginWithManifest represents extended plugin support for Manifest and Context Sharing feature
type PluginWithManifest interface {
	// Plugin original interface wrapper
	Plugin

	// GetManifest
	GetManifest() *Manifest

	// GetContextData
	GetContextData() map[string]string

	// ExecuteWithContext
	ExecuteWithContext(ctx map[string]string, args []string) error
}
