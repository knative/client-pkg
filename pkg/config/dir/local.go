/*
 Copyright 2024 The Knative Authors

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

package dir

import (
	"context"
	"log"
	"os"

	"emperror.dev/errors"
	"github.com/kirsle/configdir"
)

const (
	// ConfigDirEnvName is the name of the environment variable that can be used
	// to override the config directory.
	ConfigDirEnvName = "KN_CONFIG_DIR"

	// CacheDirEnvName is the name of the environment variable that can be used
	// to override the cache directory.
	CacheDirEnvName = "KN_CACHE_DIR"
)

var (
	configDirKey = struct{}{} //nolint:gochecknoglobals
	cacheDirKey  = struct{}{} //nolint:gochecknoglobals
)

// Config returns the path to the config directory. It will be created if it
// does not exist.
func Config(ctx context.Context) string {
	return userPath(ctx, configDirKey, ConfigDirEnvName, func() string {
		return configdir.LocalConfig("kn")
	})
}

// Cache returns the path to the cache directory. It will be created if it
// does not exist.
func Cache(ctx context.Context) string {
	return userPath(ctx, cacheDirKey, CacheDirEnvName, func() string {
		return configdir.LocalCache("kn")
	})
}

func userPath(ctx context.Context, key interface{}, envKey string, fn func() string) string {
	if p, ok := ctx.Value(key).(string); ok {
		return ensurePathExists(p)
	}
	p := os.Getenv(envKey)
	if p == "" {
		p = fn()
	}
	return ensurePathExists(p)
}

func ensurePathExists(p string) string {
	if err := configdir.MakePath(p); err != nil {
		log.Fatal(errors.WithStack(err))
	}
	return p
}
