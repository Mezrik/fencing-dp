import { CompetitionResult, getCompetitions } from "@/generated/server";
import { Api } from "@/services/Api";

export class RestApi implements Api {
  GetCompetitions(): Promise<Array<CompetitionResult>> {
    return getCompetitions();
  }
}
