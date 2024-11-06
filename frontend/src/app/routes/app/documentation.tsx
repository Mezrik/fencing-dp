import { BasicPageLayout } from '@/components/layouts';
import { t } from '@lingui/macro';
import { QueryClient } from '@tanstack/react-query';

export const docsLoader = (queryClient: QueryClient) => async () => {};

export const DocumentationRoute = () => {
  return <BasicPageLayout title={t`Documentation`}>Something</BasicPageLayout>;
};
