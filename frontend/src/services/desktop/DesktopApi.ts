import { GetCompetitions } from "@/generated/wailsjs/go/desktop/Admin";
import { query } from "@/generated/wailsjs/go/models";
import { Api } from "@/services/Api";

export class DesktopApi implements Api {
  GetCompetitions(): Promise<Array<query.Competition>> {
    return GetCompetitions();
  }
}
