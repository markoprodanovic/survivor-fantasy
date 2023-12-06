type EpisodeCastPoints = {
  castId: number;
  points: number;
};

interface Episode {
  id: number;
  episode_number: number;
  episode_date: string;
  points: EpisodeCastPoints[];
}

export type { Episode };
