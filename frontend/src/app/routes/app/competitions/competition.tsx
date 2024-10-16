import { BasicPageLayout } from '@/components/layouts';
import {
  getCompetitionQueryOptions,
  useCompetition,
} from '@/features/competitions/api/get-competition';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Overview } from '@/features/competitions/components/competition';

import { t, Trans } from '@lingui/macro';
import { QueryClient } from '@tanstack/react-query';
import { LoaderFunctionArgs, useParams } from 'react-router-dom';
import Groups from '@/features/competitions/components/competition/groups';
import { PencilIcon } from 'lucide-react';
import { Button } from '@/components/ui/button';

export const competitionLoader =
  (queryClient: QueryClient) =>
  async ({ params }: LoaderFunctionArgs) => {
    const competitionId = params.competitionId as string;

    const query = getCompetitionQueryOptions(competitionId);

    return queryClient.getQueryData(query.queryKey) ?? (await queryClient.fetchQuery(query));
  };

export const CompetitionRoute = () => {
  const params = useParams();

  const competitionId = params.competitionId as string;
  const competitionQuery = useCompetition({
    competitionId,
  });

  if (competitionQuery.isLoading) {
    return <div>Loading...</div>;
  }

  const competition = competitionQuery.data;

  if (!competition) {
    return <div>Competition not found</div>;
  }

  return (
    <BasicPageLayout
      title={competition.name}
      subtitle={t`Competition`}
      actions={
        <Button className="gap-2" variant="default">
          <PencilIcon className="size-4" />
          <span className="hidden sm:inline">
            <Trans>Edit</Trans>
          </span>
        </Button>
      }
    >
      <Tabs defaultValue="overview">
        <TabsList className="mb-2">
          <TabsTrigger value="overview">
            <Trans>Overview</Trans>
          </TabsTrigger>
          <TabsTrigger value="groups">
            <Trans>Groups</Trans>
          </TabsTrigger>
          <TabsTrigger value="elimination">
            <Trans>Elimination</Trans>
          </TabsTrigger>
          <TabsTrigger value="referees">
            <Trans>Referees</Trans>
          </TabsTrigger>
        </TabsList>
        <TabsContent value="overview">
          <Overview competition={competition}></Overview>
        </TabsContent>
        <TabsContent value="groups">
          <Groups competitionId={competition.id} />
        </TabsContent>
        <TabsContent value="elimination">WIP</TabsContent>
        <TabsContent value="referees">WIP</TabsContent>
      </Tabs>
    </BasicPageLayout>
  );
};
