import { FC } from 'react';
import { useCompetitionsGroups } from '../../api/get-groups';
import { useParticipants } from '@/features/competitors/api/get-participants';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Link } from 'react-router-dom';
import { pathnames } from '@/app/pathnames';
import { mapParticipantsByGroup } from '../../helpers';
import { Trans } from '@lingui/macro';

type GroupsProps = {
  competitionId: UUID;
};

const Groups: FC<GroupsProps> = ({ competitionId }) => {
  const groupsQuery = useCompetitionsGroups({ competitionId });
  const participantQuery = useParticipants({ competitionId });

  if (groupsQuery.isLoading || participantQuery.isLoading) {
    return <div>Loading...</div>;
  }

  const participantsByGroup = mapParticipantsByGroup(participantQuery.data ?? []);

  return (
    <div className="grid grid-cols-4 gap-4">
      {groupsQuery.data?.map((group) => (
        <Link to={pathnames.buildCompetitionGroupPath(group.id, competitionId)} key={group.id}>
          <Card className="col-span-4 sm:col-span-2 md:col-span-1 hover:bg-gray-50 transition-colors">
            <CardHeader className="flex flex-row items-center justify-between space-y-0">
              <CardTitle>{group.name}</CardTitle>
              <CardDescription>
                <Trans>Piste: {group.pisteNumber}</Trans>
              </CardDescription>
            </CardHeader>

            <CardContent>
              <ul className="divide-y divide-gray-200">
                {participantsByGroup?.[group.id]?.map((participant) => (
                  <li className="py-1">
                    {participant.competitor.firstname} {participant.competitor.surname}
                  </li>
                ))}
              </ul>
            </CardContent>
          </Card>
        </Link>
      ))}
    </div>
  );
};

export default Groups;
