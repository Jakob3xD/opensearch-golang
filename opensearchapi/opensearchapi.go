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
	"crypto/tls"
	"net/http"
	"strconv"
	"time"

	"github.com/jakob3xd/opensearch-golang"
)

type Config struct {
	Client opensearch.Config
}

type Client struct {
	client *opensearch.Client
	Cat    catClient
}

func clientInit(rootClient *opensearch.Client) *Client {
	return &Client{
		client: rootClient,
		Cat:    newCatClient(rootClient),
	}
}

// NewClient returns a opensearchapi client
func NewClient(config Config) (*Client, error) {
	rootClient, err := opensearch.NewClient(config.Client)
	if err != nil {
		return nil, err
	}
	return clientInit(rootClient), nil
}

// NewDefaultClient returns a opensearchapi client using defauls
func NewDefaultClient() (*Client, error) {
	rootClient, err := opensearch.NewClient(opensearch.Config{
		Username:  "admin",
		Password:  "admin",
		Addresses: []string{"https://localhost:9200"},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	})
	if err != nil {
		return nil, err
	}
	return clientInit(rootClient), nil
}

func (c *Client) Do(ctx context.Context, req opensearch.Request) (*Response, error) {
	resp, err := c.client.Do(ctx, req, nil)
	if err != nil {
		return nil, err
	}

	response := Response{
		statusCode: resp.StatusCode,
		body:       resp.Body,
		header:     resp.Header,
	}
	return &response, nil
}

// formatDuration converts duration to a string in the format
// accepted by Opensearch.
func formatDuration(d time.Duration) string {
	if d < time.Millisecond {
		return strconv.FormatInt(int64(d), 10) + "nanos"
	}
	return strconv.FormatInt(int64(d)/int64(time.Millisecond), 10) + "ms"
}
