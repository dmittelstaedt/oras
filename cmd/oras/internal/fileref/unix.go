//go:build !windows

/*
Copyright The ORAS Authors.
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

package fileref

import (
	"fmt"
	"strings"
)

// Parse parses file reference on unix.
func Parse(reference string, defaultMediaType string) (filePath, mediaType string, err error) {
	i := strings.LastIndex(reference, ":")
	if i < 0 {
		filePath, mediaType = reference, defaultMediaType
	} else {
		filePath, mediaType = reference[:i], reference[i+1:]
	}
	if filePath == "" {
		return "", "", fmt.Errorf("found empty file path in %q", reference)
	}
	return filePath, mediaType, nil
}
