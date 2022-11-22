package app

import (
	"context"

	ci_cd_test "gihub.com/sonyamoonglade/ci-cd-go"
)

type Service struct {
	repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUser(ctx context.Context) (*ci_cd_test.User, error) {
	return s.repo.GetUser(ctx)
}
