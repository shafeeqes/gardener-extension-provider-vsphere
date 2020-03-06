/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type EtherTypeNsServiceEntry struct {
	ResourceType string `json:"resource_type"`

	// Type of the encapsulated protocol
	EtherType int64 `json:"ether_type"`
}

type EtherTypeNsService struct {
	NsService

	NsserviceElement EtherTypeNsServiceEntry `json:"nsservice_element"`
}