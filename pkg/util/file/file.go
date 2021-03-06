/*
Copyright 2017 caicloud authors. All rights reserved.

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

package file

import (
	"os"
)

// DirExists checks the existence of the dir, true if exists, or false if not exist.
func DirExists(path string) bool {
	fileInfo, err := os.Stat(path)
	return (err == nil || os.IsExist(err)) && fileInfo.IsDir()
}

// FileExists checks the existence of the file, true if exists, or false if not exist.
func FileExists(path string) bool {
	// First check the err, if the file exists, err is nil
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}
