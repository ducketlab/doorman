// Code generated by protoc-gen-go-http. DO NOT EDIT.

package role

import (
	http "github.com/ducketlab/mingo/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				Path: "/auth.role.RoleService/CreateRole",
			},
			{
				Path: "/auth.role.RoleService/DescribeRole",
			},
			{
				Path: "/auth.role.RoleService/DescribePermission",
			},
		},
	}
	return set
}
