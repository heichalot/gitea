// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package setting

import (
	"path/filepath"
	"testing"

	"code.gitea.io/gitea/models/db"
)

func TestMain(m *testing.M) {
	db.MainTest(m, filepath.Join("..", "..", "..", ".."))
}
