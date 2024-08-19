import { FC } from 'react';
import { useCompetitions } from '../api/get-competitions';

export const RecentCompetitions: FC<{}> = () => {
  const competitionsQuery = useCompetitions({});

  if (competitionsQuery.isLoading) {
    return <div>Loading...</div>;
  }

  const competitions = competitionsQuery.data || [];

  return (
    <div className="space-y-8">
      {competitions.map((competition) => (
        <div className="flex items-center" key={competition.name}>
          <div className="ml-4 space-y-1">
            <p className="text-sm font-medium leading-none">{competition.name}</p>
            <p className="text-sm text-muted-foreground">Lorem ipsum dolor sit amet</p>
          </div>
          <div className="ml-auto font-medium">99 contestants</div>
        </div>
      ))}
    </div>
  );
};
