import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import { getGenderIcon } from '@/features/competitions/helpers';
import { CompetitorResult } from '@/generated/server';
import { formatUIDate } from '@/utils/date';
import { Trans } from '@lingui/macro';
import { FC } from 'react';

export const CompetitorsTable: FC<{ data: CompetitorResult[] }> = ({ data }) => {
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>
            <Trans>Name</Trans>
          </TableHead>
          <TableHead>
            <Trans>Club</Trans>
          </TableHead>
          <TableHead>
            <Trans>Gender</Trans>
          </TableHead>
          <TableHead>
            <Trans>License</Trans>
          </TableHead>
          <TableHead>
            <Trans>Birthday</Trans>
          </TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {data?.map((comp) => {
          const GenderIcon = getGenderIcon(comp.gender);
          return (
            <TableRow key={comp.id}>
              <TableCell>
                {comp.firstname} {comp.surname}
              </TableCell>
              <TableCell>{comp.club.name}</TableCell>
              <TableCell>{GenderIcon && <GenderIcon className="text-primary" />}</TableCell>
              <TableCell>{comp.license}</TableCell>
              <TableCell>{formatUIDate(comp.birthdate)}</TableCell>
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
};
