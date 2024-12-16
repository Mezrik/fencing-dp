package entities

import "github.com/Mezrik/fencing-dp/internal/common"

type GroupConfig struct {
	common.Entity
	groupsCount int
	groupSizes  []int
}

func CalculateGroupSizes(totalCompetitors, minGroups, maxGroups int) []GroupConfig {
	var configurations []GroupConfig

	for g := minGroups; g <= maxGroups; g++ {
		baseSize := totalCompetitors / g
		remainder := totalCompetitors % g

		groupSizes := make([]int, g)
		for i := 0; i < g; i++ {
			if i < remainder {
				groupSizes[i] = baseSize + 1
			} else {
				groupSizes[i] = baseSize
			}
		}

		configurations = append(configurations, GroupConfig{
			groupsCount: g,
			groupSizes:  groupSizes,
		})
	}

	return configurations
}

func (g GroupConfig) GroupsCount() int {
	return g.groupsCount
}

func (g GroupConfig) GroupSizes() []int {
	return g.groupSizes
}
