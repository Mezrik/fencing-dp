import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { CompetitionResult } from '@/generated/server';
import { formatUIDate } from '@/utils/date';
import { t, Trans } from '@lingui/macro';
import { FC } from 'react';
import { getCompetionTypeCaption, getGenderCaption } from '../../helpers';
import { InfoList, InfoListItem } from '@/components/ui/info-list';
import { ParticipantsList } from '../participant/participants-list';

export const Overview: FC<{ competition: CompetitionResult }> = ({ competition }) => {
  return (
    <div className="grid grid-cols-4 gap-4 w-full">
      <Card className="col-span-1">
        <CardHeader>
          <CardTitle>
            <Trans>Basic information</Trans>
          </CardTitle>
        </CardHeader>
        <CardContent>
          <InfoList>
            <InfoListItem title={t`Organizer`} description={competition.organizerName} />
            <InfoListItem title={t`Federation`} description={competition.federationName} />
            <InfoListItem
              title={t`Scope of competition`}
              description={getCompetionTypeCaption(competition.competitionType)}
            />
            <InfoListItem title={t`Category`} description={competition.category.name} />
            <InfoListItem title={t`Weapon`} description={competition.weapon.name} />
            <InfoListItem title={t`Gender`} description={getGenderCaption(competition.gender)} />
            <InfoListItem
              title={t`Date of the event`}
              description={formatUIDate(competition.date)}
            />
          </InfoList>
        </CardContent>
      </Card>
      <Card className="col-span-3">
        <CardHeader>
          <CardTitle>
            <Trans>List of competitors</Trans>
          </CardTitle>
        </CardHeader>
        <CardContent>
          <ParticipantsList competitionId={competition.id} />
        </CardContent>
      </Card>
    </div>
  );
};
