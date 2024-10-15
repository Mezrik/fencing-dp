import { Match } from '@/generated/server';
import { MatchParticipant } from '../types';
import { FC } from 'react';

type MatchCardProps = {
  match: Match;
  participantOne: MatchParticipant;
  participantTwo: MatchParticipant;
  onClick: VoidFunction;
};

export const MatchCard: FC<MatchCardProps> = ({ participantOne, participantTwo, onClick }) => {
  return (
    <button
      onClick={onClick}
      type="button"
      className="max-w-[33%] bg-gray-100 rounded-lg px-4 py-2 space-y-1.5 hover:bg-gray-200 transition-colors"
    >
      <div className="flex justify-between gap-4">
        <span>
          {participantOne.firstname} {participantOne.surname}
        </span>
        <span>{participantOne.points ?? 0}</span>
      </div>
      <div className="flex justify-between gap-4">
        <span>
          {participantTwo.firstname} {participantTwo.surname}
        </span>
        <span>{participantTwo.points ?? 0}</span>
      </div>
    </button>
  );
};
