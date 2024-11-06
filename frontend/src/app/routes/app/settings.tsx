import { BasicPageLayout } from '@/components/layouts';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/form';
import { dynamicActivate, locales } from '@/i18n';
import { t } from '@lingui/macro';
import { QueryClient } from '@tanstack/react-query';

export const settingsLoader = (queryClient: QueryClient) => async () => {};

export const SettingsRoute = () => {
  return (
    <BasicPageLayout title={t`Settings`}>
      <Select
        onValueChange={(value) => {
          dynamicActivate(value as keyof typeof locales);
        }}
      >
        <SelectTrigger>
          <SelectValue placeholder={t`Select a language`} />
        </SelectTrigger>
        <SelectContent>
          {Object.entries(locales).map(([key, value]) => (
            <SelectItem value={key} key={key}>
              {value}
            </SelectItem>
          ))}
        </SelectContent>
      </Select>
    </BasicPageLayout>
  );
};
