import { getCompetitionsQueryOptions } from '@/features/competitions/api/get-competitions';
import { QueryClient } from '@tanstack/react-query';

export const competitionsLoader = (queryClient: QueryClient) => async () => {
  const query = getCompetitionsQueryOptions();

  return queryClient.getQueryData(query.queryKey) ?? (await queryClient.fetchQuery(query));
};

export const CompetitionsRoute = () => {
  return <div className="text-lg">Competitions</div>;
};
