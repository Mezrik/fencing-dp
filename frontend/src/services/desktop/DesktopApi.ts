import { CreateCompetition, GetCompetitions } from '@/generated/wailsjs/go/desktop/Admin';
import { command, query } from '@/generated/wailsjs/go/models';
import { Api } from '@/services/Api';

export class DesktopApi implements Api {
  GetCompetitions(): Promise<Array<query.Competition>> {
    return GetCompetitions();
  }

  CreateCompetition(command: command.CreateCompetition): Promise<void> {
    return CreateCompetition(command);
  }
}
