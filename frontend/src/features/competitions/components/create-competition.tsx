import { FC } from 'react';
import { useMediaQuery } from 'usehooks-ts';
import {
  Drawer,
  DrawerClose,
  DrawerContent,
  DrawerFooter,
  DrawerHeader,
  DrawerTitle,
} from '@/components/ui/drawer';
import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { t, Trans } from '@lingui/macro';
import { useToast } from '@/hooks/ui/use-toast';
import { createCompetitionInputSchema, useCreateCompetition } from '../api/create-competition';
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
  Input,
  RadioGroupFormField,
  RadioOption,
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/form';
import { CompetitionTypeEnum, GenderEnum } from '@/generated/server';
import { useCompetitionCategories } from '../api/get-competition-categories';
import { useWeapons } from '../api/get-weapons';
import { getCompetionTypeCaption, getGenderCaption } from '../helpers';
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import { CalendarIcon } from 'lucide-react';
import { Calendar } from '@/components/ui/calendar';
import { cn } from '@/utils/class-names';
import { formatUIDate } from '@/utils/date';

type CreateCompetitionProps = {
  open: boolean;
  onOpenChange: (open: boolean) => void;
};

const TITLE = t`Create Competition`;

const FORM_ID = 'create-competition-form';

const COMPETITION_TYPE_OPTIONS: RadioOption[] = [
  {
    label: getCompetionTypeCaption(CompetitionTypeEnum.international),
    value: CompetitionTypeEnum.international,
  },
  {
    label: getCompetionTypeCaption(CompetitionTypeEnum.national),
    value: CompetitionTypeEnum.national,
  },
];

const GENDER_OPTIONS: RadioOption[] = [
  {
    label: getGenderCaption(GenderEnum.male),
    value: GenderEnum.male,
  },
  {
    label: getGenderCaption(GenderEnum.female),
    value: GenderEnum.female,
  },
  {
    label: getGenderCaption(GenderEnum.mixed),
    value: GenderEnum.mixed,
  },
];

export const CreateCompetitionForm: FC<{ onSubmit: () => void }> = ({ onSubmit }) => {
  const { toast } = useToast();

  const categoriesQuery = useCompetitionCategories();
  const weaponsQuery = useWeapons();

  const createCompetitionMutation = useCreateCompetition({
    mutationConfig: {
      onSuccess: () => {
        toast({
          description: 'Competition created successfully',
          variant: 'success',
        });
      },
    },
  });

  return (
    <Form
      id={FORM_ID}
      onSubmit={(values) => {
        createCompetitionMutation.mutate({ data: values });
        onSubmit();
      }}
      schema={createCompetitionInputSchema}
    >
      {({ register, formState, control }) => (
        <>
          <Input
            label={t`Competition name`}
            error={formState.errors['name']}
            registration={register('name')}
          />

          <Input
            label={t`Organizer name`}
            error={formState.errors['organizerName']}
            registration={register('organizerName')}
          />

          <Input
            label={t`Federation name`}
            error={formState.errors['federationName']}
            registration={register('federationName')}
          />

          <fieldset className="flex gap-12">
            <RadioGroupFormField
              label={t`Competition type`}
              name="competitionType"
              control={control}
              options={COMPETITION_TYPE_OPTIONS}
            />
            <RadioGroupFormField
              label={t`Gender`}
              name="gender"
              control={control}
              options={GENDER_OPTIONS}
            />
          </fieldset>

          {categoriesQuery.data && (
            <FormField
              control={control}
              name="categoryId"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>
                    <Trans>Category</Trans>
                  </FormLabel>
                  <Select onValueChange={field.onChange} defaultValue={field.value}>
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select a competition category" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      {categoriesQuery.data?.map((c) => (
                        <SelectItem value={c.id} key={c.id}>
                          {c.name}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              )}
            />
          )}

          {weaponsQuery.data && (
            <FormField
              control={control}
              name="weaponId"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>
                    <Trans>Weapon</Trans>
                  </FormLabel>
                  <Select onValueChange={field.onChange} defaultValue={field.value}>
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select a competition weapon" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      {weaponsQuery.data?.map((w) => (
                        <SelectItem value={w.id} key={w.id}>
                          {w.name}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              )}
            />
          )}

          <FormField
            control={control}
            name="date"
            render={({ field }) => (
              <FormItem className="flex flex-col">
                <FormLabel>
                  <Trans>Date of competition</Trans>
                </FormLabel>
                <Popover>
                  <PopoverTrigger asChild>
                    <FormControl>
                      <Button
                        variant={'outline'}
                        className={cn(
                          'w-full pl-3 text-left font-normal',
                          !field.value && 'text-muted-foreground',
                        )}
                      >
                        {field.value ? (
                          formatUIDate(field.value)
                        ) : (
                          <span>
                            <Trans>Pick a date</Trans>
                          </span>
                        )}
                        <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
                      </Button>
                    </FormControl>
                  </PopoverTrigger>
                  <PopoverContent className="w-auto p-0" align="start">
                    <Calendar
                      mode="single"
                      selected={field.value}
                      onSelect={field.onChange}
                      disabled={(date) => date > new Date() || date < new Date('1900-01-01')}
                      initialFocus
                    />
                  </PopoverContent>
                </Popover>
                <FormMessage />
              </FormItem>
            )}
          />
        </>
      )}
    </Form>
  );
};

export const CreateCompetitionDrawer: FC<CreateCompetitionProps> = (props) => {
  return (
    <Drawer open={props.open} onOpenChange={props.onOpenChange}>
      <DrawerContent>
        <DrawerHeader>
          <DrawerTitle>{TITLE}</DrawerTitle>
        </DrawerHeader>

        <div className="p-4 pb-0">
          <CreateCompetitionForm onSubmit={() => props.onOpenChange(false)} />
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

export const CreateCompetitionDialog: FC<CreateCompetitionProps> = (props) => {
  return (
    <Dialog open={props.open} onOpenChange={props.onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{TITLE}</DialogTitle>
        </DialogHeader>

        <CreateCompetitionForm onSubmit={() => props.onOpenChange(false)} />

        <DialogFooter>
          <Button form={FORM_ID}>
            <Trans>Submit</Trans>
          </Button>
          <DialogClose asChild>
            <Button variant="outline">Cancel</Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export const CreateCompetition: FC<CreateCompetitionProps> = (props) => {
  const isDesktop = useMediaQuery('(min-width: 768px)');

  if (isDesktop) {
    return <CreateCompetitionDialog {...props} />;
  }

  return <CreateCompetitionDrawer {...props} />;
};
