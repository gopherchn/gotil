package proxy

import "testing"

func TestUserProxy_Login(t *testing.T) {
	type fields struct {
		User *User
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "proxy",
			fields: fields{
				User: &User{
				},
			},
			args: args{
				username: "magatron",
				password: "bert",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserProxy{
				User: tt.fields.User,
			}
			if err := u.Login(tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("UserProxy.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
