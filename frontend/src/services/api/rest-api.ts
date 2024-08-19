import {
  CompetitionCategoryResult,
  CompetitionResult,
  CreateCompetitionCommand,
  getCompetitions,
  getCompetitionsCategories,
  postCompetitions,
} from '@/generated/server';
import { Api } from './api';

export class RestApi implements Api {
  GetCompetitions(): Promise<Array<CompetitionResult>> {
    return getCompetitions();
  }

  CreateCompetition(data: CreateCompetitionCommand): Promise<void> {
    return postCompetitions(data);
  }

  GetCompetitionsCategories(): Promise<Array<CompetitionCategoryResult>> {
    return getCompetitionsCategories();
  }
}
