import {
  CreateCompetition,
  GetCompetitions,
  GetCompetition,
  GetCompetitionsCategories,
  GetCompetitionsWeapons,
  CreateCompetitor,
  GetCompetitors,
  GetCompetitionsGroups,
  GetGroup,
  GetMatch,
  GetMatches,
  GetParticipants,
  AssignCompetitor,
  UpdateCompetitor,
} from '@/generated/wailsjs/go/desktop/Admin';
import { command, query } from '@/generated/wailsjs/go/models';
import { Api } from './api';

export class DesktopApi implements Api {
  GetCompetitions(): Promise<Array<query.Competition>> {
    return GetCompetitions();
  }

  GetCompetition(id: UUID): Promise<query.Competition> {
    return GetCompetition(id);
  }

  CreateCompetition(command: command.CreateCompetition): Promise<void> {
    return CreateCompetition(command);
  }

  GetCompetitionsCategories(): Promise<Array<query.CompetitionCategory>> {
    return GetCompetitionsCategories();
  }

  GetCompetitionsWeapons(): Promise<Array<query.Weapon>> {
    return GetCompetitionsWeapons();
  }

  CreateCompetitor(command: command.CreateCompetitor): Promise<void> {
    return CreateCompetitor(command);
  }

  UpdateCompetitor(id: UUID, command: Omit<command.UpdateCompetitor, 'id'>): Promise<void> {
    return UpdateCompetitor({ id, ...command });
  }

  GetCompetitors(): Promise<Array<query.Competitor>> {
    return GetCompetitors();
  }

  GetCompetitionsGroups(competitionId: UUID): Promise<Array<query.Group>> {
    return GetCompetitionsGroups(competitionId);
  }

  GetGroup(groupId: UUID): Promise<query.Group> {
    return GetGroup(groupId);
  }

  GetMatch(id: UUID): Promise<query.MatchDetail> {
    return GetMatch(id);
  }

  GetMatches(groupId: UUID): Promise<Array<query.Match>> {
    return GetMatches(groupId);
  }

  GetParticipants(competitionId: UUID): Promise<Array<query.Participant>> {
    return GetParticipants(competitionId);
  }

  AssignParticipant(competitorId: UUID, competitionId: UUID): Promise<void> {
    return AssignCompetitor(competitorId, competitionId);
  }

  ImportCompetitor(): Promise<void> {
    return Promise.resolve();
  }
}
