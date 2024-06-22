import { CompetitionResult } from "@/generated/server";

export interface Api {
  GetCompetitions(): Promise<Array<CompetitionResult>>;
}
