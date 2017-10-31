// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

type UpdateCustomerSecretKeyDetails struct {

	// The description you assign to the secret key. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}