import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import {
  Drawer,
  DrawerClose,
  DrawerContent,
  DrawerFooter,
  DrawerHeader,
  DrawerTitle,
} from '@/components/ui/drawer';
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/form';
import {
  assignParticipantInputSchema,
  useAssignParticipant,
} from '@/features/competitors/api/assign-participant';
import { useCompetitors } from '@/features/competitors/api/get-competitors';
import { useToast } from '@/hooks/ui/use-toast';
import { msg, t, Trans } from '@lingui/macro';
import { useLingui } from '@lingui/react';
import { FC } from 'react';
import { useMediaQuery } from 'usehooks-ts';

const FORM_ID = 'assign-competitor-form';

type AssignParticipantProps = {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  competitionId: UUID;
};

export const AssignParticipantForm: FC<{ onSubmit: () => void; competitionId: UUID }> = ({
  onSubmit,
  competitionId,
}) => {
  const { toast } = useToast();
  const { _ } = useLingui();

  const competitorsQuery = useCompetitors();

  const assignParticipantMutation = useAssignParticipant({
    mutationConfig: {
      onSuccess: () => {
        toast({
          description: 'Competitor assigned successfully',
          variant: 'success',
        });
      },
    },
  });

  if (competitorsQuery.isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <Form
      id={FORM_ID}
      onSubmit={(values) => {
        assignParticipantMutation.mutate({ data: values });
        onSubmit();
      }}
      schema={assignParticipantInputSchema}
      options={{
        defaultValues: {
          competitionId: competitionId,
          competitorId: '',
        },
      }}
    >
      {({ control }) => (
        <>
          <FormField
            control={control}
            name="competitorId"
            render={({ field }) => (
              <FormItem>
                <FormLabel>
                  <Trans>Competitor</Trans>
                </FormLabel>
                <Select onValueChange={field.onChange} defaultValue={field.value}>
                  <FormControl>
                    <SelectTrigger>
                      <SelectValue placeholder={_(msg`Select a competitor`)} />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    {competitorsQuery.data?.map((w) => (
                      <SelectItem value={w.id} key={w.id}>
                        {w.firstname} {w.surname}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                <FormMessage />
              </FormItem>
            )}
          />
        </>
      )}
    </Form>
  );
};

export const AssignParticipantDrawer: FC<AssignParticipantProps> = (props) => {
  const { _ } = useLingui();

  return (
    <Drawer open={props.open} onOpenChange={props.onOpenChange}>
      <DrawerContent>
        <DrawerHeader>
          <DrawerTitle>{_(msg`Assign Competitor`)}</DrawerTitle>
        </DrawerHeader>

        <div className="p-4 pb-0">
          <AssignParticipantForm
            onSubmit={() => props.onOpenChange(false)}
            competitionId={props.competitionId}
          />
        </div>

        <DrawerFooter>
          <Button form={FORM_ID}>
            <Trans>Submit</Trans>
          </Button>
          <DrawerClose asChild>
            <Button variant="outline">
              <Trans>Cancel</Trans>
            </Button>
          </DrawerClose>
        </DrawerFooter>
      </DrawerContent>
    </Drawer>
  );
};

export const AssignParticipantDialog: FC<AssignParticipantProps> = (props) => {
  const { _ } = useLingui();

  return (
    <Dialog
      open={props.open}
      onOpenChange={props.onOpenChange}
      aria-describedby={_(msg`Assign Competitor`)}
    >
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{_(msg`Assign Competitor`)}</DialogTitle>
        </DialogHeader>

        <AssignParticipantForm
          onSubmit={() => props.onOpenChange(false)}
          competitionId={props.competitionId}
        />

        <DialogFooter>
          <Button form={FORM_ID}>
            <Trans>Submit</Trans>
          </Button>
          <DialogClose asChild>
            <Button variant="outline">
              <Trans>Cancel</Trans>
            </Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export const AssignParticipant: FC<AssignParticipantProps> = (props) => {
  const isDesktop = useMediaQuery('(min-width: 768px)');

  if (isDesktop) {
    return <AssignParticipantDialog {...props} />;
  }

  return <AssignParticipantDrawer {...props} />;
};
