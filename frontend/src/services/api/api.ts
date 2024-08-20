import {
  CompetitionCategoryResult,
  CompetitionResult,
  CreateCompetitionCommand,
  WeaponResult,
} from '@/generated/server';
import { DesktopApi } from './desktop-api';
import { RestApi } from './rest-api';

// TODO: Generate this automatically from generated/server generated/wailsjs
// Could use Hygen for this
export interface Api {
  GetCompetitions(): Promise<Array<CompetitionResult>>;

  GetCompetition(id: UUID): Promise<CompetitionResult>;

  CreateCompetition(data: CreateCompetitionCommand): Promise<void>;

  GetCompetitionsCategories(): Promise<Array<CompetitionCategoryResult>>;

  GetCompetitionsWeapons(): Promise<Array<WeaponResult>>;
}

export const api: Api = import.meta.env.MODE === 'desktop' ? new DesktopApi() : new RestApi();
