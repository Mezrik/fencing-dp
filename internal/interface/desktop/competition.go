package desktop

import "github.com/Mezrik/fencing-dp/internal/app/common"

func (a *Admin) GetCompetitions() []*common.CompetitionResult {

	competitions, _ := a.competitionService.GetAllCompetitions()

	return competitions
}
