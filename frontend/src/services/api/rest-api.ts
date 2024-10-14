import {
  CompetitionCategoryResult,
  CompetitionResult,
  CreateCompetitionCommand,
  getCompetitions,
  getCompetitionsCategories,
  postCompetitions,
  getCompetitionsCompetitionId,
  WeaponResult,
  getCompetitionsWeapons,
  CompetitorResult,
  CreateCompetitorCommand,
  getCompetitors,
  postCompetitors,
  CompetitionGroup,
  MatchDetail,
  Match,
  CompetitionParticipant,
  getCompetitionsCompetitionIdGroups,
  getCompetitionsCompetitionIdGroupsGroupId,
  getMatchesMatchId,
  getMatchesGroupIdAll,
  getCompetitorsAllCompetitionId,
} from '@/generated/server';
import { Api } from './api';

export class RestApi implements Api {
  GetCompetitions(): Promise<Array<CompetitionResult>> {
    return getCompetitions();
  }

  GetCompetition(id: UUID): Promise<CompetitionResult> {
    return getCompetitionsCompetitionId(id);
  }

  CreateCompetition(data: CreateCompetitionCommand): Promise<void> {
    return postCompetitions(data);
  }

  GetCompetitionsCategories(): Promise<Array<CompetitionCategoryResult>> {
    return getCompetitionsCategories();
  }

  GetCompetitionsWeapons(): Promise<Array<WeaponResult>> {
    return getCompetitionsWeapons();
  }

  GetCompetitors(): Promise<Array<CompetitorResult>> {
    return getCompetitors();
  }

  CreateCompetitor(data: CreateCompetitorCommand): Promise<void> {
    return postCompetitors(data);
  }

  GetCompetitionsGroups(competitionId: UUID): Promise<Array<CompetitionGroup>> {
    return getCompetitionsCompetitionIdGroups(competitionId);
  }

  GetGroup(groupId: UUID, competitionId: UUID): Promise<CompetitionGroup> {
    return getCompetitionsCompetitionIdGroupsGroupId(competitionId, groupId);
  }

  GetMatch(id: UUID): Promise<MatchDetail> {
    return getMatchesMatchId(id);
  }

  GetMatches(groupId: UUID): Promise<Array<Match>> {
    return getMatchesGroupIdAll(groupId);
  }

  GetParticipants(competitionId: UUID): Promise<Array<CompetitionParticipant>> {
    return getCompetitorsAllCompetitionId(competitionId);
  }
}
