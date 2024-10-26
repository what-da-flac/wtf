// Mantine :
import { Grid } from '@mantine/core';

// Components :

// Models :
import MovieCard from './MovieCard';
import { Movie } from '../models/movie';

type params = { rows: Movie[] };

export default function MovieCards({ rows }: params) {
  return (
    <>
      <h1>Movie List</h1>
      <Grid>
        {rows.map(r => (
          <Grid.Col span={{ base: 12, md: 6, lg: 4, xl: 3 }} key={r.id}>
            <MovieCard m={r} />
          </Grid.Col>
        ))}
      </Grid>
    </>
  );
}
