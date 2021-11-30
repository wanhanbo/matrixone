// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package extend

import (
	"github.com/matrixorigin/matrixone/pkg/container/batch"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func (e *ParenExtend) IsLogical() bool {
	return e.E.IsLogical()
}

func (_ *ParenExtend) IsConstant() bool {
	return false
}

func (e *ParenExtend) ReturnType() types.T {
	return e.E.ReturnType()
}

func (e *ParenExtend) Attributes() []string {
	return e.E.Attributes()
}

func (e *ParenExtend) Eval(bat *batch.Batch, proc *process.Process) (*vector.Vector, types.T, error) {
	return e.E.Eval(bat, proc)
}

func (a *ParenExtend) Eq(b Extend) bool {
	return a.E.Eq(b)
}

func (e *ParenExtend) String() string {
	return "(" + e.E.String() + ")"
}
