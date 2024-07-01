import {
  CompetitionResult,
  CreateCompetitionCommand,
  getCompetitions,
  postCompetitions,
} from '@/generated/server';
import { Api } from '@/services/Api';

export class RestApi implements Api {
  GetCompetitions(): Promise<Array<CompetitionResult>> {
    return getCompetitions();
  }

  CreateCompetition(data: CreateCompetitionCommand): Promise<void> {
    return postCompetitions(data);
  }
}
