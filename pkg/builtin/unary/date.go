// Copyright 2022 Matrix Origin
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

package unary

import (
	"fmt"
	"github.com/matrixorigin/matrixone/pkg/builtin"
	"github.com/matrixorigin/matrixone/pkg/container/nulls"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/encoding"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec/extend"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec/extend/overload"
	"github.com/matrixorigin/matrixone/pkg/vectorize/date"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func init() {
	extend.FunctionRegistry["date"] = builtin.Date
	overload.AppendFunctionRets(builtin.Date, []types.T{types.T_date}, types.T_date)
	overload.AppendFunctionRets(builtin.Date, []types.T{types.T_datetime}, types.T_date)
	extend.UnaryReturnTypes[builtin.Date] = func(e extend.Extend) types.T {
		return getUnaryReturnType(builtin.Date, e)
	}
	extend.UnaryStrings[builtin.Date] = func(e extend.Extend) string {
		return fmt.Sprintf("date(%s)", e)
	}
	overload.OpTypes[builtin.Date] = overload.Unary
	overload.UnaryOps[builtin.Date] = []*overload.UnaryOp{
		{
			Typ:        types.T_date,
			ReturnType: types.T_date,
			Fn: func(lv *vector.Vector, proc *process.Process, _ bool) (*vector.Vector, error) {
				lvs := lv.Col.([]types.Date)
				size := types.T(types.T_date).TypeLen()
				if lv.Ref == 1 || lv.Ref == 0 {
					lv.Ref = 0
					date.DateToDate(lvs, lvs)
					return lv, nil
				}
				vec, err := process.Get(proc, int64(size)*int64(len(lvs)), types.Type{Oid: types.T_date, Size: int32(size)})
				if err != nil {
					return nil, err
				}
				rs := encoding.DecodeDateSlice(vec.Data)
				rs = rs[:len(lvs)]
				vec.Col = rs
				nulls.Set(vec.Nsp, lv.Nsp)
				vector.SetCol(vec, date.DateToDate(lvs, rs))
				return vec, nil
			},
		},
		{
			Typ:        types.T_datetime,
			ReturnType: types.T_date,
			Fn: func(lv *vector.Vector, proc *process.Process, _ bool) (*vector.Vector, error) {
				lvs := lv.Col.([]types.Datetime)
				size := types.T(types.T_date).TypeLen()
				vec, err := process.Get(proc, int64(size)*int64(len(lvs)), types.Type{Oid: types.T_date, Size: int32(size)})
				if err != nil {
					return nil, err
				}
				rs := encoding.DecodeDateSlice(vec.Data)
				rs = rs[:len(lvs)]
				vec.Col = rs
				nulls.Set(vec.Nsp, lv.Nsp)
				vector.SetCol(vec, date.DateTimeToDate(lvs, rs))
				return vec, nil
			},
		},
	}
}
