//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package plan

import "encoding/json"

type Stream struct {
	readonly

	cost        float64
	cardinality float64
}

func NewStream(cost, cardinality float64) *Stream {
	return &Stream{
		cost:        cost,
		cardinality: cardinality,
	}
}

func (this *Stream) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitStream(this)
}

func (this *Stream) New() Operator {
	return &Stream{}
}

func (this *Stream) Cost() float64 {
	return this.cost
}

func (this *Stream) Cardinality() float64 {
	return this.cardinality
}

func (this *Stream) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.MarshalBase(nil))
}

func (this *Stream) MarshalBase(f func(map[string]interface{})) map[string]interface{} {
	r := map[string]interface{}{"#operator": "Stream"}
	if this.cost > 0.0 {
		r["cost"] = this.cost
	}
	if this.cardinality > 0.0 {
		r["cardinality"] = this.cardinality
	}
	if f != nil {
		f(r)
	}
	return r
}

func (this *Stream) UnmarshalJSON(body []byte) error {
	var _unmarshalled struct {
		_           string  `json:"#operator"`
		Cost        float64 `json:"cost"`
		Cardinality float64 `json:"cardinality"`
	}

	err := json.Unmarshal(body, &_unmarshalled)
	if err != nil {
		return err
	}

	this.cost = getCost(_unmarshalled.Cost)
	this.cardinality = getCardinality(_unmarshalled.Cardinality)

	return nil
}
