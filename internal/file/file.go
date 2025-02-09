//
// Copyright (C) 2019-2022 vdaas.org vald team <vald@vdaas.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Package file provides file I/O functionality
package file

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/vdaas/vald/internal/errors"
)

// Open opens the file with the given path, flag and permission.
// If the folder does not exists, create the folder.
// If the file does not exist, create the file.
func Open(path string, flg int, perm fs.FileMode) (file *os.File, err error) {
	if path == "" {
		return nil, errors.ErrPathNotSpecified
	}

	defer func() {
		if err != nil && file != nil {
			err = errors.Wrap(file.Close(), err.Error())
			file = nil
		}
	}()
	if _, err = os.Stat(path); err != nil {
		if _, err = os.Stat(filepath.Dir(path)); err != nil {
			err = os.MkdirAll(filepath.Dir(path), perm)
			if err != nil {
				return nil, err
			}
		}

		if flg&(os.O_CREATE|os.O_APPEND) > 0 {
			file, err = os.Create(path)
			if err != nil {
				return file, err
			}
		}

		if file != nil {
			err = file.Close()
			if err != nil {
				return file, err
			}
		}
	}

	file, err = os.OpenFile(path, flg, perm)
	if err != nil {
		return file, err
	}

	return file, nil
}

// Exists returns file existence
func Exists(path string) (e bool) {
	e, _, _ = exists(path)
	return e
}

// ExistsWithDetail returns file existence
func ExistsWithDetail(path string) (e bool, fi fs.FileInfo, err error) {
	return exists(path)
}

// exists returns file existence with detailed information
func exists(path string) (exists bool, fi fs.FileInfo, err error) {
	fi, err = os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true, fi, nil
		}
		if os.IsNotExist(err) {
			return false, fi, err
		}
		return false, fi, err
	}
	return true, fi, nil
}

// ListInDir returns file list in directory
func ListInDir(path string) ([]string, error) {
	exists, fi, err := exists(path)
	if !exists {
		return nil, err
	}
	if fi.Mode().IsDir() && !strings.HasSuffix(path, string(os.PathSeparator)) {
		path = filepath.Clean(path) + string(os.PathSeparator)
	}
	path = filepath.Dir(path)
	files, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		return nil, err
	}
	return files, nil
}
