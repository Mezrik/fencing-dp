import { MutationConfig } from '@/lib/react-query';
import { api } from '@/services/api';
import { z } from 'zod';
import { getParticipantsQueryOptions } from './get-participants';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { t } from '@lingui/macro';

export const assignParticipantInputSchema = z.object({
  competitionId: z.string().uuid(t`Select correct competition`),
  competitorId: z.string().uuid(t`Select correct competitor`),
});

export type AssignParticipantInput = z.infer<typeof assignParticipantInputSchema>;

export const assignParticipant = ({ data }: { data: AssignParticipantInput }) => {
  return api.AssignParticipant(data.competitorId, data.competitionId);
};

type UseAssignParticipantOptions = {
  mutationConfig?: MutationConfig<typeof assignParticipant>;
};

export const useAssignParticipant = ({ mutationConfig }: UseAssignParticipantOptions = {}) => {
  const queryClient = useQueryClient();

  const { onSuccess, ...rest } = mutationConfig || {};

  return useMutation({
    onSuccess: (...args) => {
      const [, { data }] = args;
      queryClient.invalidateQueries({
        queryKey: getParticipantsQueryOptions(data.competitorId).queryKey,
      });
      onSuccess?.(...args);
    },
    ...rest,
    mutationFn: assignParticipant,
  });
};
