import { Trans } from '@lingui/macro';
import { AppShell } from '@/components/app-shell/app-shell';
import { AppProvider } from './app-provider';
import { useContext, useEffect } from 'react';
import { ApiContext } from '@/services/ApiProvider';
import { CompetitionTypeEnum, GenderEnum } from '@/generated/server';

const Test = () => {
  const api = useContext(ApiContext);

  useEffect(() => {
    const get = async () => {
      const result = await api?.GetCompetitionsCategories();
      console.log(result);
      return result;
    };

    const post = async () => {
      await api?.CreateCompetition({
        categoryId: 'fd051bee-50b2-4207-9700-8c3aa584366f',
        competitionType: CompetitionTypeEnum.international,
        date: new Date().toISOString(),
        federationName: 'test',
        gender: GenderEnum.mixed,
        name: 'test competition',
        organizerName: 'tst org',
        weaponId: 'b5b5eea8-2241-4b4b-9d69-9f8bcd82716d',
      });

      console.log(await get());
    };

    get();
  }, []);

  return null;
};

function App() {
  return (
    <AppProvider>
      <AppShell>
        <Trans>This is a translation</Trans>
        <Test />
      </AppShell>
    </AppProvider>
  );
}

export default App;
