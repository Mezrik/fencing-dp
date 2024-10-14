package app

import (
	"github.com/Mezrik/fencing-dp/internal/match/app/query"
	"github.com/Mezrik/fencing-dp/internal/match/domain/repositories"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	AllMatches query.AllMatchesHandler
	GetMatch   query.GetMatchHandler
}

func NewMatchService(matchRepo repositories.MatchRepository, logger *logrus.Entry) Service {
	return Service{
		Commands: Commands{},
		Queries:  Queries{AllMatches: query.NewAllMatchesHandler(matchRepo, logger), GetMatch: query.NewGetMatchHandler(matchRepo, logger)},
	}
}
