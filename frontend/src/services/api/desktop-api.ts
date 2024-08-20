import {
  CreateCompetition,
  GetCompetitions,
  GetCompetition,
  GetCompetitionsCategories,
  GetCompetitionsWeapons,
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
}
