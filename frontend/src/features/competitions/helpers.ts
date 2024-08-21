import { CompetitionTypeEnum, GenderEnum } from '@/generated/server';
import { t } from '@lingui/macro';

export const getCompetionTypeCaption = (competitionType: CompetitionTypeEnum) => {
  switch (competitionType) {
    case CompetitionTypeEnum.national:
      return t`National`;
    case CompetitionTypeEnum.international:
      return t`International`;
    default:
      return '???';
  }
};

export const getGenderCaption = (gender: GenderEnum) => {
  switch (gender) {
    case GenderEnum.male:
      return t`Men`;
    case GenderEnum.female:
      return t`Women`;
    case GenderEnum.mixed:
      return t`Mixed`;
    default:
      return '???';
  }
};
