//  Copyright (c) 2018 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package semantics

import (
	"github.com/couchbase/query/algebra"
	"github.com/couchbase/query/datastore"
	"github.com/couchbase/query/errors"
)

func (this *SemChecker) VisitCreatePrimaryIndex(stmt *algebra.CreatePrimaryIndex) (interface{}, error) {
	return nil, stmt.MapExpressions(this)
}

func (this *SemChecker) VisitCreateIndex(stmt *algebra.CreateIndex) (interface{}, error) {
	gsi := stmt.Using() == datastore.GSI || stmt.Using() == datastore.DEFAULT
	for i, term := range stmt.Keys() {
		if term.HasAttribute(algebra.IK_MISSING) && (i > 0 || !gsi) {
			return nil, errors.NewSemanticsError(nil, "MISSING attribute only allowed on GSI index leading key")
		}
	}
	return nil, stmt.MapExpressions(this)
}

func (this *SemChecker) VisitDropIndex(stmt *algebra.DropIndex) (interface{}, error) {
	return nil, stmt.MapExpressions(this)
}

func (this *SemChecker) VisitAlterIndex(stmt *algebra.AlterIndex) (interface{}, error) {
	return nil, stmt.MapExpressions(this)
}

func (this *SemChecker) VisitBuildIndexes(stmt *algebra.BuildIndexes) (interface{}, error) {
	return nil, stmt.MapExpressions(this)
}

func (this *SemChecker) VisitCreateScope(stmt *algebra.CreateScope) (interface{}, error) {
	return nil, stmt.MapExpressions(this)
}

func (this *SemChecker) VisitDropScope(stmt *algebra.DropScope) (interface{}, error) {
	return nil, stmt.MapExpressions(this)
}

func (this *SemChecker) VisitCreateCollection(stmt *algebra.CreateCollection) (interface{}, error) {
	return nil, stmt.MapExpressions(this)
}

func (this *SemChecker) VisitDropCollection(stmt *algebra.DropCollection) (interface{}, error) {
	return nil, stmt.MapExpressions(this)
}

func (this *SemChecker) VisitFlushCollection(stmt *algebra.FlushCollection) (interface{}, error) {
	return nil, stmt.MapExpressions(this)
}
