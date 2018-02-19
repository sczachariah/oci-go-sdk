// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListPathRouteSetsRequest wrapper for the ListPathRouteSets operation
type ListPathRouteSetsRequest struct {

	// The OCID ({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the load balancer associated with the path route sets
	// to retrieve.
	LoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
}

func (request ListPathRouteSetsRequest) String() string {
	return common.PointerString(request)
}

// ListPathRouteSetsResponse wrapper for the ListPathRouteSets operation
type ListPathRouteSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []PathRouteSet instance
	Items []PathRouteSet `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListPathRouteSetsResponse) String() string {
	return common.PointerString(response)
}
