package service

import (
	"context"
	"testing"
)

func TestServiceUser_IsUserExists(t *testing.T) {
	type fields struct {
		repositoryUser RepositoryUser
	}
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "case user id is zero",
			fields: fields{
				repositoryUser: &RepositoryUserMock{},
			},
			args: args{
				userID: 0,
			},
			want: false,
		},
		{
			name: "case user id is not zero and repo return false",
			fields: fields{
				repositoryUser: &RepositoryUserMock{
					IsUserExistsFunc: func(ctx context.Context, userID int64) bool {
						return false
					},
				},
			},
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: false,
		},
		{
			name: "case user id is not zero and repo return true",
			fields: fields{
				repositoryUser: &RepositoryUserMock{
					IsUserExistsFunc: func(ctx context.Context, userID int64) bool {
						return true
					},
				},
			},
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := NewServiceUser(tt.fields.repositoryUser)
			if got := su.IsUserExists(tt.args.ctx, tt.args.userID); got != tt.want {
				t.Errorf("ServiceUser.IsUserExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
