import { BasicPageLayout } from '@/components/layouts';
import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { InputBase } from '@/components/ui/form';
import {
  getCompetitorsQueryOptions,
  useCompetitors,
} from '@/features/competitors/api/get-competitors';
import { useImportCompetitor } from '@/features/competitors/api/import-competitor';
import { CompetitorsTable } from '@/features/competitors/components/competitors-table';
import { useToast } from '@/hooks/ui/use-toast';
import { msg, Trans } from '@lingui/macro';
import { useLingui } from '@lingui/react';
import { QueryClient } from '@tanstack/react-query';
import { ImportIcon, PlusIcon } from 'lucide-react';
import { useState } from 'react';

export const competitorsLoader = (queryClient: QueryClient) => async () => {
  const query = getCompetitorsQueryOptions();

  return queryClient.getQueryData(query.queryKey) ?? (await queryClient.fetchQuery(query));
};

export const CompetitorsRoute = () => {
  const { _ } = useLingui();
  const { toast } = useToast();
  const competitorsQuery = useCompetitors({});

  const [importDialogOpen, setImportDialogOpen] = useState(false);
  const [importedFile, setImportedFile] = useState<File | null>(null);
  const importCompetitorMutation = useImportCompetitor({
    mutationConfig: {
      onSuccess: () => {
        toast({
          description: _(msg`Competitor imported successfully`),
          variant: 'success',
        });

        setImportDialogOpen(false);

        competitorsQuery.refetch();
      },
    },
  });

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
        <div className="flex gap-2">
          <Button className="gap-2" variant="default" onClick={() => setImportDialogOpen(true)}>
            <ImportIcon className="size-4" />
            <span className="hidden">
              <Trans>Import competitors</Trans>
            </span>
          </Button>
          <Button className="gap-2" variant="secondary">
            <PlusIcon className="size-4" />
            <span className="hidden sm:inline">
              <Trans>Add competitor</Trans>
            </span>
          </Button>
        </div>
      }
      className="flex flex-col min-h-0"
    >
      <div className="mb-4">
        <InputBase placeholder={_(msg`Type here to search`)} className="max-w-56" />
      </div>
      <div className="flex-grow overflow-y-auto">
        <CompetitorsTable data={competitorsQuery.data} />
      </div>

      <Dialog open={importDialogOpen} onOpenChange={setImportDialogOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{_(msg`Import competitors`)}</DialogTitle>
          </DialogHeader>

          <InputBase
            placeholder="Picture"
            type="file"
            accept="text/csv, application/csv"
            onChange={(event) => setImportedFile(event.target.files && event.target.files[0])}
          />

          <DialogFooter>
            <Button
              onClick={() =>
                importedFile && importCompetitorMutation.mutate({ data: { file: importedFile } })
              }
            >
              <Trans>Import</Trans>
            </Button>
            <DialogClose asChild>
              <Button variant="outline">Cancel</Button>
            </DialogClose>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </BasicPageLayout>
  );
};
