import { useParticipants } from '@/features/competitors/api/get-participants';
import { FC } from 'react';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import { t } from '@lingui/macro';
import { CZ } from 'country-flag-icons/react/3x2';

export const ParticipantsList: FC<{ competitionId: UUID }> = ({ competitionId }) => {
  const participantsQuery = useParticipants({ competitionId });

  if (participantsQuery.isLoading) {
    return <div>Loading...</div>;
  }

  const participants = participantsQuery.data;

  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead className="w-20">{t`Nation`}</TableHead>
          <TableHead>{t`Name`}</TableHead>
          <TableHead>{t`Deployment No.`}</TableHead>
          <TableHead>{t`Points`}</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {participants?.map((participant) => (
          <TableRow>
            <TableCell>
              <CZ title={t`Czech Republic`} className="size-6" />
            </TableCell>
            <TableCell>
              {participant.competitor.firstname} {participant.competitor.surname}
            </TableCell>
            <TableCell>{participant.deploymentNumber}</TableCell>
            <TableCell>{participant.points}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};
