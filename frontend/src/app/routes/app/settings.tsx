import { BasicPageLayout } from '@/components/layouts';
import { t } from '@lingui/macro';
import { QueryClient } from '@tanstack/react-query';

export const settingsLoader = (queryClient: QueryClient) => async () => {};

export const SettingsRoute = () => {
  return <BasicPageLayout title={t`Settings`}>Something</BasicPageLayout>;
};
