import { MenGenderIcon, MixedGenderIcon, WomenGenderIcon } from '@/assets/icons';
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
