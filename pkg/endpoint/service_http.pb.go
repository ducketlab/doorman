// Code generated by protoc-gen-go-http. DO NOT EDIT.

package endpoint

import (
	http "github.com/ducketlab/mingo/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				Path: "/auth.endpoint.EndpointService/DescribeEndpoint",
			},
			{
				Path: "/auth.endpoint.EndpointService/QueryEndpoints",
			},
			{
				Path: "/auth.endpoint.EndpointService/Registry",
			},
			{
				Path: "/auth.endpoint.EndpointService/DeleteEndpoint",
			},
		},
	}
	return set
}
