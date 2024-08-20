const competitionsPath = '/competitions';
const competitionPath = `${competitionsPath}/:competitionId`;
const buildCompetitionPath = (id: string) => {
  return `${competitionsPath}/${id}`;
};

const dashboardPath = '';

export const pathnames = { competitionsPath, competitionPath, buildCompetitionPath, dashboardPath };
