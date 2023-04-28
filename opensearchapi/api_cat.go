// SPDX-License-Identifier: Apache-2.0
//
// The OpenSearch Contributors require contributions made to
// this file be licensed under the Apache-2.0 license or a
// compatible open source license.
//
// Modifications Copyright OpenSearch Contributors. See
// GitHub history for details.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package opensearchapi

import (
	"context"

	"github.com/jakob3xd/opensearch-golang"
)

type catClient struct {
	client *opensearch.Client
}

func newCatClient(client *opensearch.Client) catClient {
	return catClient{client: client}
}

func (c catClient) Indices(ctx context.Context, req *CatIndicesReq) (*CatIndicesResp, error) {
	var data CatIndicesResp
	req.Params.Format = "json"
	req.Params.FilterPath = []string{}
	_, err := c.client.Do(ctx, req, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}