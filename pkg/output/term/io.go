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

package term

import (
	"io"
	"os"

	"golang.org/x/term"
)

// IsReaderTerminal returns true if the given reader is a real terminal.
func IsReaderTerminal(r io.Reader) bool {
	f, ok := r.(*os.File)
	return ok && term.IsTerminal(int(f.Fd()))
}

// IsTerminal returns true if the given writer is a real terminal.
func IsTerminal(w io.Writer) bool {
	if forceColor() {
		return true
	}
	f, ok := w.(*os.File)
	return ok && supportsColor() && term.IsTerminal(int(f.Fd()))
}
