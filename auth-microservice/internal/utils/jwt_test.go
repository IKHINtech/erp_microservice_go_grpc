package utils_test

import (
	"testing"

	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/config"
	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/internal/utils"
)

func TestGenerateJWT(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		userID  string
		cfg     config.Config
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := utils.GenerateJWT(tt.userID, tt.cfg)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GenerateJWT() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GenerateJWT() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GenerateJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
