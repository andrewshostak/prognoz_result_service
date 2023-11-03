package service

import (
	"context"

	"github.com/andrewshostak/result-service/client"
	"github.com/andrewshostak/result-service/repository"
)

type AliasRepository interface {
	Find(ctx context.Context, alias string) (*repository.Alias, error)
}

type MatchRepository interface {
	Create(ctx context.Context, match repository.Match) (*repository.Match, error)
	Search(ctx context.Context, search repository.Match) (*repository.Match, error)
}

type FootballAPIClient interface {
	SearchFixtures(ctx context.Context, search client.FixtureSearch) (*client.FixturesResponse, error)
}
