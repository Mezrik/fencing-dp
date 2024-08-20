import {
  getCompetitionQueryOptions,
  useCompetition,
} from '@/features/competitions/api/get-competition';
import { QueryClient } from '@tanstack/react-query';
import { LoaderFunctionArgs, useParams } from 'react-router-dom';

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

  return <div className="text-lg">{competition.name}</div>;
};
