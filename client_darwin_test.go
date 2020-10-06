// +build darwin

/*
   Copyright The containerd Authors.

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

package containerd

import (
	"io/ioutil"
	"os"
)

var (
	defaultRoot    = "/var/lib/containerd-test"
	defaultState   = "/var/run/containerd-test"
	defaultAddress = "/var/run/containerd-test/containerd.sock"
	testImage      string
	shortCommand   = withProcessArgs("hello")
	longCommand    = withProcessArgs("ping", "127.0.0.1")
)

func init() {
	testImage = "docker.io/ukontainer/runu-base:0.3-osx-extra"

	tmpDir, _ := ioutil.TempDir("", "containerd-test-")
	defaultRoot = tmpDir + "/root"
	defaultState = tmpDir + "/state"
	defaultAddress = defaultState + "/containerd.sock"
	defer os.RemoveAll(tmpDir)
}
