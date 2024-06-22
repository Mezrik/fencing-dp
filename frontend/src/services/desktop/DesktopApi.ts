import { GetCompetitions } from "@/generated/wailsjs/go/desktop/Admin";
import { common } from "@/generated/wailsjs/go/models";
import { Api } from "@/services/Api";

export class DesktopApi implements Api {
  GetCompetitions(): Promise<Array<common.CompetitionResult>> {
    return GetCompetitions();
  }
}
