package sample

import "testing"

func TestValidateThailandCitizenID(t *testing.T) {
	type args struct {
		idNo string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Digit",
			args: args{
				idNo: "123456789",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateThailandCitizenID(tt.args.idNo); (err != nil) != tt.wantErr {
				t.Errorf("ValidateThailandCitizenID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
