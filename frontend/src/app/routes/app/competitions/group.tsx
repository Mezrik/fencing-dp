import { BasicPageLayout } from '@/components/layouts';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import {
  getCompetitionGroupQueryOptions,
  useCompetitionGroup,
} from '@/features/competitions/api/get-group';
import {
  getCompetitionsGroupsQueryOptions,
  useCompetitionsGroups,
} from '@/features/competitions/api/get-groups';
import { mapParticipantsByGroup } from '@/features/competitions/helpers';
import {
  getParticipantsQueryOptions,
  useParticipants,
} from '@/features/competitors/api/get-participants';
import { t } from '@lingui/macro';
import { QueryClient } from '@tanstack/react-query';
import { CZ } from 'country-flag-icons/react/3x2';
import { LoaderFunctionArgs, useParams } from 'react-router-dom';

export const groupLoader =
  (queryClient: QueryClient) =>
  async ({ params }: LoaderFunctionArgs) => {
    const groupId = params.groupId as string;
    const competitionId = params.competitionId as string;

    const query = getCompetitionGroupQueryOptions(competitionId, groupId);
    const participantQuery = getParticipantsQueryOptions(competitionId);

    return (
      (queryClient.getQueryData(query.queryKey) ?? (await queryClient.fetchQuery(query))) &&
      (queryClient.getQueryData(participantQuery.queryKey) ??
        (await queryClient.fetchQuery(participantQuery)))
    );
  };

export const GroupRoute = () => {
  const params = useParams();

  const groupId = params.groupId as string;
  const competitionId = params.competitionId as string;

  const groupQuery = useCompetitionGroup({ groupId, competitionId });
  const participantQuery = useParticipants({ competitionId });

  if (groupQuery.isLoading) {
    return <div>Loading...</div>;
  }

  const group = groupQuery.data;

  if (!group) {
    return <div>Group not found</div>;
  }

  const participantsByGroup = mapParticipantsByGroup(participantQuery.data ?? []);

  return (
    <BasicPageLayout title={group.name} subtitle={t`Group`}>
      <div className="grid grid-cols-9 gap-4">
        <Card className="pt-6 col-span-9 lg:col-span-4">
          <CardContent>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-20">{t`Nation`}</TableHead>
                  <TableHead>{t`Name`}</TableHead>
                  <TableHead>{t`Deployment No.`}</TableHead>
                  <TableHead>{t`Points`}</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {participantsByGroup?.[group.id]?.map((participant) => (
                  <TableRow>
                    <TableCell>
                      <CZ title={t`Czech Republic`} className="size-6" />
                    </TableCell>
                    <TableCell>
                      {participant.competitor.firstname} {participant.competitor.surname}
                    </TableCell>
                    <TableCell>{participant.deploymentNumber}</TableCell>
                    <TableCell>{participant.points}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </CardContent>
        </Card>
        <Card className="col-span-9 lg:col-span-5">
          <CardHeader>
            <CardTitle>{t`Matches`}</CardTitle>
          </CardHeader>
          <CardContent></CardContent>
        </Card>
      </div>
    </BasicPageLayout>
  );
};
