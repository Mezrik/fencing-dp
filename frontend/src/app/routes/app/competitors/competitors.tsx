import { BasicPageLayout } from '@/components/layouts';
import { Button } from '@/components/ui/button';
import { InputBase } from '@/components/ui/form';
import {
  getCompetitorsQueryOptions,
  useCompetitors,
} from '@/features/competitors/api/get-competitors';
import { CompetitorsTable } from '@/features/competitors/components/CompetitorsTable';
import { t, Trans } from '@lingui/macro';
import { QueryClient } from '@tanstack/react-query';
import { PlusIcon } from 'lucide-react';

export const competitorsLoader = (queryClient: QueryClient) => async () => {
  const query = getCompetitorsQueryOptions();

  return queryClient.getQueryData(query.queryKey) ?? (await queryClient.fetchQuery(query));
};

export const CompetitorsRoute = () => {
  const competitorsQuery = useCompetitors({});

  if (competitorsQuery.isLoading) {
    return <BasicPageLayout title={t`Competitors`}>Loading...</BasicPageLayout>;
  }

  if (!competitorsQuery.data || competitorsQuery.data.length <= 0) {
    return <BasicPageLayout title={t`Competitors`}>No competitors found</BasicPageLayout>;
  }

  return (
    <BasicPageLayout
      title={t`Competitors`}
      actions={
        <Button className="gap-2" variant="default">
          <PlusIcon className="size-4" />
          <span className="hidden sm:inline">
            <Trans>Add competitor</Trans>
          </span>
        </Button>
      }
    >
      <div className="mb-4">
        <InputBase placeholder={t`Type here to search`} className="max-w-56" />
      </div>
      <CompetitorsTable data={competitorsQuery.data} />
    </BasicPageLayout>
  );
};
