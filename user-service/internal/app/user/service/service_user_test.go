package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"user-service/internal/app/domain"
)

func TestServiceUser_Login(t *testing.T) {
	type fields struct {
		repositoryUser RepositoryUser
		utilBcrypt     UtilBcrypt
		utilJwt        UtilJwt
	}
	type args struct {
		ctx   context.Context
		param domain.LoginParam
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "case email not exist",
			fields: fields{
				repositoryUser: &RepositoryUserMock{
					GetByEmailFunc: func(ctx context.Context, email string) domain.Repository {
						return domain.Repository{
							Error: errors.New("email not exist"),
						}
					},
				},
			},
			args: args{
				ctx: context.Background(),
				param: domain.LoginParam{
					Email:    "test@gmail.com",
					Password: "test",
				},
			},
			want: "GM5EQNy8msHL",
		},
		{
			name: "case wrong email or password",
			fields: fields{
				repositoryUser: &RepositoryUserMock{
					GetByEmailFunc: func(ctx context.Context, email string) domain.Repository {
						return domain.Repository{
							User: domain.User{
								ID:       1,
								Password: "test",
							},
						}
					},
				},
				utilBcrypt: &UtilBcryptMock{
					CheckIfPasswordHashIsEqualFunc: func(repoPassword []byte, paramPassword []byte) error {
						return errors.New("wrong email or password")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				param: domain.LoginParam{
					Email:    "test@gmail.com",
					Password: "test",
				},
			},
			want: "GMz3Y5sNe96O",
		},
		{
			name: "case error when generate hash",
			fields: fields{
				repositoryUser: &RepositoryUserMock{
					GetByEmailFunc: func(ctx context.Context, email string) domain.Repository {
						return domain.Repository{
							User: domain.User{
								ID:       1,
								Password: "test",
							},
						}
					},
				},
				utilBcrypt: &UtilBcryptMock{
					CheckIfPasswordHashIsEqualFunc: func(repoPassword []byte, paramPassword []byte) error {
						return nil
					},
				},
				utilJwt: &UtilJwtMock{
					GenerateJWTFunc: func(id int64) (string, error) {
						return "", errors.New("error when generate hash")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				param: domain.LoginParam{
					Email:    "test@gmail.com",
					Password: "test",
				},
			},
			want: "GMItRDipgI19",
		},
		{
			name: "case no error",
			fields: fields{
				repositoryUser: &RepositoryUserMock{
					GetByEmailFunc: func(ctx context.Context, email string) domain.Repository {
						return domain.Repository{
							User: domain.User{
								ID:       1,
								Password: "test",
							},
						}
					},
				},
				utilBcrypt: &UtilBcryptMock{
					CheckIfPasswordHashIsEqualFunc: func(repoPassword []byte, paramPassword []byte) error {
						return nil
					},
				},
				utilJwt: &UtilJwtMock{
					GenerateJWTFunc: func(id int64) (string, error) {
						return "xxx", nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				param: domain.LoginParam{
					Email:    "test@gmail.com",
					Password: "test",
				},
			},
			want: "GMijx79VQ7bS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := NewServiceUser(tt.fields.repositoryUser, tt.fields.utilBcrypt, tt.fields.utilJwt)
			if got := su.Login(tt.args.ctx, tt.args.param); !reflect.DeepEqual(got.Hash, tt.want) {
				t.Errorf("ServiceUser.Login() = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}

func TestServiceUser_Register(t *testing.T) {
	type fields struct {
		repositoryUser RepositoryUser
		utilBcrypt     UtilBcrypt
		utilJwt        UtilJwt
	}
	type args struct {
		ctx   context.Context
		param domain.RegisterParam
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "case failed generate hash",
			fields: fields{
				utilBcrypt: &UtilBcryptMock{
					GenerateHashFromPlainPasswordFunc: func(paramPassword []byte) (string, error) {
						return "", errors.New("failed generate hash password")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				param: domain.RegisterParam{
					Name:     "test",
					Email:    "test@gmail.com",
					Password: "test",
				},
			},
			want: "GMeLrJKgIHO9",
		},
		{
			name: "case failed create user to database",
			fields: fields{
				repositoryUser: &RepositoryUserMock{
					CreateFunc: func(ctx context.Context, user *domain.User) domain.Repository {
						return domain.Repository{
							Error: errors.New("failed create user to database"),
						}
					},
				},
				utilBcrypt: &UtilBcryptMock{
					GenerateHashFromPlainPasswordFunc: func(paramPassword []byte) (string, error) {
						return "xxx", nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				param: domain.RegisterParam{
					Name:     "test",
					Email:    "test@gmail.com",
					Password: "test",
				},
			},
			want: "GMHWzXXMSk9p",
		},
		{
			name: "case failed generate hash jwt",
			fields: fields{
				repositoryUser: &RepositoryUserMock{
					CreateFunc: func(ctx context.Context, user *domain.User) domain.Repository {
						return domain.Repository{}
					},
				},
				utilBcrypt: &UtilBcryptMock{
					GenerateHashFromPlainPasswordFunc: func(paramPassword []byte) (string, error) {
						return "xxx", nil
					},
				},
				utilJwt: &UtilJwtMock{
					GenerateJWTFunc: func(id int64) (string, error) {
						return "", errors.New("failed generate hash jwt")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				param: domain.RegisterParam{
					Name:     "test",
					Email:    "test@gmail.com",
					Password: "test",
				},
			},
			want: "GMlpLJtRdVdx",
		},
		{
			name: "case no error",
			fields: fields{
				repositoryUser: &RepositoryUserMock{
					CreateFunc: func(ctx context.Context, user *domain.User) domain.Repository {
						return domain.Repository{}
					},
				},
				utilBcrypt: &UtilBcryptMock{
					GenerateHashFromPlainPasswordFunc: func(paramPassword []byte) (string, error) {
						return "xxx", nil
					},
				},
				utilJwt: &UtilJwtMock{
					GenerateJWTFunc: func(id int64) (string, error) {
						return "xxx", nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				param: domain.RegisterParam{
					Name:     "test",
					Email:    "test@gmail.com",
					Password: "test",
				},
			},
			want: "GMmmtY5oiY9J",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := NewServiceUser(tt.fields.repositoryUser, tt.fields.utilBcrypt, tt.fields.utilJwt)
			if got := su.Register(tt.args.ctx, tt.args.param); !reflect.DeepEqual(got.Hash, tt.want) {
				t.Errorf("ServiceUser.Register() = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}

func TestServiceUser_GetByID(t *testing.T) {
	type fields struct {
		repositoryUser RepositoryUser
		utilBcrypt     UtilBcrypt
		utilJwt        UtilJwt
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "case invalid user id",
			args: args{
				ctx: context.Background(),
				id:  0,
			},
			want: "GML8UxsMd5E5",
		},
		{
			name: "case user not exist in database",
			fields: fields{
				repositoryUser: &RepositoryUserMock{
					GetByIDFunc: func(ctx context.Context, id int64) domain.Repository {
						return domain.Repository{
							Error: errors.New("user not exist in database"),
						}
					},
				},
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: "GMO9escS9esE",
		},
		{
			name: "case no error",
			fields: fields{
				repositoryUser: &RepositoryUserMock{
					GetByIDFunc: func(ctx context.Context, id int64) domain.Repository {
						return domain.Repository{
							User: domain.User{
								ID:    1,
								Name:  "test",
								Email: "test@gmail.com",
							},
						}
					},
				},
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: "GMDxylSw7Gnm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			su := NewServiceUser(tt.fields.repositoryUser, tt.fields.utilBcrypt, tt.fields.utilJwt)
			if got := su.GetByID(tt.args.ctx, tt.args.id); !reflect.DeepEqual(got.Hash, tt.want) {
				t.Errorf("ServiceUser.GetByID() = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}
