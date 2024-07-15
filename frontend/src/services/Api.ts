import {
  CompetitionCategoryResult,
  CompetitionResult,
  CreateCompetitionCommand,
} from '@/generated/server';

// TODO: Generate this automatically from generated/server generated/wailsjs
// Could use Hygen for this
export interface Api {
  GetCompetitions(): Promise<Array<CompetitionResult>>;

  CreateCompetition(data: CreateCompetitionCommand): Promise<void>;

  GetCompetitionsCategories(): Promise<Array<CompetitionCategoryResult>>;
}
