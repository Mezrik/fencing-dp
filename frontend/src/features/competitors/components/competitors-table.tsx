import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import { getGenderAbbrv, getGenderIcon } from '@/features/competitions/helpers';
import { CompetitorResult } from '@/generated/server';
import { formatUIDate } from '@/utils/date';
import { Trans } from '@lingui/macro';
import { useLingui } from '@lingui/react';
import { FC } from 'react';

export const CompetitorsTable: FC<{ data: CompetitorResult[] }> = ({ data }) => {
  const { _ } = useLingui();

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
          return (
            <TableRow key={comp.id}>
              <TableCell>
                {comp.firstname} {comp.surname}
              </TableCell>
              <TableCell>{comp.club?.name}</TableCell>
              <TableCell>{getGenderAbbrv(comp.gender, _)}</TableCell>
              <TableCell>{comp.license}</TableCell>
              <TableCell>{comp.birthdate ? formatUIDate(comp.birthdate) : '-'}</TableCell>
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
};
