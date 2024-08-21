import { BasicPageLayout } from '@/components/layouts';
import {
  getCompetitionsQueryOptions,
  useCompetitions,
} from '@/features/competitions/api/get-competitions';
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

  return <BasicPageLayout title={t`Competitions`}>{JSON.stringify(competitions)}</BasicPageLayout>;
};
