import { BasicPageLayout } from '@/components/layouts';
import { Button } from '@/components/ui/button';
import { InputBase } from '@/components/ui/form';
import {
  getCompetitorsQueryOptions,
  useCompetitors,
} from '@/features/competitors/api/get-competitors';
import { CompetitorsTable } from '@/features/competitors/components/competitors-table';
import { msg, Trans } from '@lingui/macro';
import { useLingui } from '@lingui/react';
import { QueryClient } from '@tanstack/react-query';
import { PlusIcon } from 'lucide-react';

export const competitorsLoader = (queryClient: QueryClient) => async () => {
  const query = getCompetitorsQueryOptions();

  return queryClient.getQueryData(query.queryKey) ?? (await queryClient.fetchQuery(query));
};

export const CompetitorsRoute = () => {
  const { _ } = useLingui();
  const competitorsQuery = useCompetitors({});

  if (competitorsQuery.isLoading) {
    return <BasicPageLayout title={_(msg`Competitors`)}>Loading...</BasicPageLayout>;
  }

  if (!competitorsQuery.data || competitorsQuery.data.length <= 0) {
    return <BasicPageLayout title={_(msg`Competitors`)}>No competitors found</BasicPageLayout>;
  }

  return (
    <BasicPageLayout
      title={_(msg`Competitors`)}
      actions={
        <Button className="gap-2" variant="default">
          <PlusIcon className="size-4" />
          <span className="hidden sm:inline">
            <Trans>Add competitor</Trans>
          </span>
        </Button>
      }
      className="flex flex-col min-h-0"
    >
      <div className="mb-4">
        <InputBase placeholder={_(msg`Type here to search`)} className="max-w-56" />
      </div>
      <div className="flex-grow overflow-y-auto">
        <CompetitorsTable data={competitorsQuery.data} />
      </div>
    </BasicPageLayout>
  );
};
