import { CompetitionResult, CreateCompetitionCommand } from '@/generated/server';

export interface Api {
  GetCompetitions(): Promise<Array<CompetitionResult>>;

  CreateCompetition(data: CreateCompetitionCommand): Promise<void>;
}
