import { BasicPageLayout } from '@/components/layouts';
import {
  getCompetitionsQueryOptions,
  useCompetitions,
} from '@/features/competitions/api/get-competitions';
import { CompetitionCard } from '@/features/competitions/components/competition-card';
import { t } from '@lingui/macro';
import { QueryClient } from '@tanstack/react-query';

export const competitionsLoader = (queryClient: QueryClient) => async () => {
  const query = getCompetitionsQueryOptions();

  return queryClient.getQueryData(query.queryKey) ?? (await queryClient.fetchQuery(query));
};

export const CompetitionsRoute = () => {
  const competitionsQuery = useCompetitions({});

  if (competitionsQuery.isLoading) {
    return <BasicPageLayout title={t`Competitions`}>Loading...</BasicPageLayout>;
  }

  const competitions = competitionsQuery.data ?? [];

  return (
    <BasicPageLayout title={t`Competitions`}>
      <div className="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
        {competitions.map((comp) => (
          <CompetitionCard
            key={comp.id}
            name={comp.name}
            gender={comp.gender}
            date={comp.date}
            weapon={comp.weapon.name}
            competitionId={comp.id}
          />
        ))}
      </div>
    </BasicPageLayout>
  );
};
