const competitionsPath = '/competitions';
const competitionPath = `${competitionsPath}/:competitionId`;
const buildCompetitionPath = (id: UUID) => {
  return `${competitionsPath}/${id}`;
};

const competitionGroupPath = `${competitionPath}/groups/:groupId`;

const buildCompetitionGroupPath = (id: UUID, competitionId: UUID) => {
  return `${buildCompetitionPath(competitionId)}/groups/${id}`;
};

const dashboardPath = '';

const competitorsPath = '/competitors';

const settingsPath = '/settings';

const docsPath = '/docs';

export const pathnames = {
  competitionsPath,
  competitionPath,
  buildCompetitionPath,
  competitionGroupPath,
  buildCompetitionGroupPath,
  dashboardPath,
  competitorsPath,
  settingsPath,
  docsPath,
};
