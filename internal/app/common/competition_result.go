package common

type CompetitionResult struct {
	Result
	Name     string `json:"name"`
	Category *CompetitionCategoryResult
	Weapon   *WeaponResult
}
