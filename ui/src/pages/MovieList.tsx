import { useEffect, useState } from 'react';

// Components :
import PreLoader from '../components/PreLoader';
import { notifyErrResponse } from '../components/Errors';

// Models :

// Helpers :
import { ApiMovieList } from '../helpers/api';
import { Movie } from '../models/movie';
import MovieCards from '../components/MovieCards';

export default function MovieList() {
  const [rows, setRows] = useState<Movie[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    async function loadData() {
      try {
        const data = await ApiMovieList();
        setRows(data);
      } catch (err) {
        await notifyErrResponse(err);
      } finally {
        setIsLoading(false);
      }
    }

    loadData();
  }, []);

  return isLoading ? <PreLoader /> : <>{<MovieCards rows={rows} />}</>;
}
