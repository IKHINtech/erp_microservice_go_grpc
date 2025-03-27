package auth

import "google.golang.org/protobuf/runtime/protoimpl"

type (
	LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	ValidateTokenRequest struct {
		Token string `json:"token" binding:"required"`
	}

	ValidateTokenResponse struct {
		state         protoimpl.MessageState `protogen:"open.v1"`
		IsValid       bool                   `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
		UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
		Roles         []string               `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
		unknownFields protoimpl.UnknownFields
		sizeCache     protoimpl.SizeCache
	}

	LoginResponse struct {
		Token string `json:"token"`
	}

	ErrorResponse struct {
		Error   string `json:"error"`
		Details string `json:"details,omitempty"`
	}
)
