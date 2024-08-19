import { api } from '@/services/api';
import { CompetitionTypeEnum, GenderEnum } from '@/generated/server';

import { AppRouter } from './router';
import { AppProvider } from './app-provider';

const Test = () => {
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

export const App = () => {
  return (
    <AppProvider>
      <AppRouter />
    </AppProvider>
  );
};
