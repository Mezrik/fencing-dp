import { Trans } from '@lingui/macro';
import { AppShell } from '@/components/app-shell/app-shell';
import { AppProvider } from './app-provider';
import { useContext, useEffect } from 'react';
import { ApiContext } from '@/services/ApiProvider';
import { CompetitionTypeEnum, GenderEnum } from '@/generated/server';

const Test = () => {
  const api = useContext(ApiContext);

  const get = async () => {
    const result = await api?.GetCompetitionsCategories();
    console.log(result);
    return result;
  };

  const post = async () => {
    await api?.CreateCompetition({
      categoryId: '9ecf8eb0-4101-40fd-b46e-41d9cf39fc22',
      competitionType: CompetitionTypeEnum.international,
      date: new Date().toISOString(),
      federationName: 'test',
      gender: GenderEnum.mixed,
      name: 'test competition',
      organizerName: 'tst org',
      weaponId: '4b0dec55-2a2b-46f7-8fd6-c001761148da',
    });

    console.log(await get());
  };

  return <button onClick={post}>Click</button>;
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
