import { BasicPageLayout } from '@/components/layouts';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { getCompetitionsQueryOptions } from '@/features/competitions/api/get-competitions';
import { RecentCompetitions } from '@/features/competitions/components/recent-competitions';
import { t, Trans } from '@lingui/macro';
import { QueryClient } from '@tanstack/react-query';

export const dashboardLoader = (queryClient: QueryClient) => async () => {
  const query = getCompetitionsQueryOptions();

  return queryClient.getQueryData(query.queryKey) ?? (await queryClient.fetchQuery(query));
};

export const DashboardRoute = () => {
  return (
    <BasicPageLayout title={t`Dashboard`}>
      <Card className="col-span-3">
        <CardHeader>
          <CardTitle>
            <Trans>Recent competitions</Trans>
          </CardTitle>
        </CardHeader>
        <CardContent>
          <RecentCompetitions />
        </CardContent>
      </Card>
    </BasicPageLayout>
  );
};
