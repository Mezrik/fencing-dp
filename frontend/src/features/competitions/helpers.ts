import { MenGenderIcon, MixedGenderIcon, WomenGenderIcon } from '@/assets/icons';
import { I18n } from '@lingui/core';
import { CompetitionParticipant, CompetitionTypeEnum, GenderEnum } from '@/generated/server';
import { msg } from '@lingui/macro';

export const getCompetionTypeCaption = (competitionType: CompetitionTypeEnum, _: I18n['_']) => {
  switch (competitionType) {
    case CompetitionTypeEnum.national:
      return _(msg`National`);
    case CompetitionTypeEnum.international:
      return _(msg`International`);
    default:
      return '???';
  }
};

export const getGenderCaption = (gender: GenderEnum, _: I18n['_']) => {
  switch (gender) {
    case GenderEnum.male:
      return _(msg`Men`);
    case GenderEnum.female:
      return _(msg`Women`);
    case GenderEnum.mixed:
      return _(msg`Mixed`);
    default:
      return '???';
  }
};

export const getGenderIcon = (gender: GenderEnum) => {
  switch (gender) {
    case GenderEnum.male:
      return MenGenderIcon;
    case GenderEnum.female:
      return WomenGenderIcon;
    case GenderEnum.mixed:
      return MixedGenderIcon;
    default:
      return null;
  }
};

export const getGenderAbbrv = (gender: GenderEnum, _: I18n['_']) => {
  switch (gender) {
    case GenderEnum.male:
      return _(msg`M`);
    case GenderEnum.female:
      return _(msg`F`);
    case GenderEnum.mixed:
      return _(msg`-`);
    default:
      return '???';
  }
};

export const mapParticipantsByGroup = (participants: CompetitionParticipant[]) => {
  const participantsByGroup = participants.reduce<Record<UUID, CompetitionParticipant[]>>(
    (acc, participant) => {
      if (!participant.groupId) return acc;

      if (!acc[participant.groupId]) {
        acc[participant.groupId] = [];
      }

      acc[participant.groupId].push(participant);

      return acc;
    },
    {},
  );

  return participantsByGroup;
};

export const mapParticipantsByCompetitorId = (participants: CompetitionParticipant[]) => {
  return participants.reduce<Record<UUID, CompetitionParticipant>>((acc, participant) => {
    acc[participant.competitor.id] = participant;

    return acc;
  }, {});
};
