export interface Movie {
  id: number
  title: string
  overview: string
  poster_path: string | null
  release_date: string
  vote_average: number
  vote_count: number
}

export interface MoviePage {
  page: number
  results: Movie[]
  total_pages: number
  total_results: number
}

export interface Genre {
  id: number
  name: string
}

export interface MovieDetails extends Movie {
  original_title: string
  tagline: string
  runtime: number | null
  genres: Genre[]
  backdrop_path: string | null
}

export async function getMovies(): Promise<MoviePage> {
  const response = await fetch('/api/v1/movies', {
    signal: AbortSignal.timeout(10000),
  })

  if (!response.ok) {
    throw new Error(`Movie request failed: ${response.status}`)
  }

  return response.json()
}

export async function getMovieDetails(id: string): Promise<MovieDetails> {
  const numericId = Number(id)

  if (!/^\d+$/.test(id) || !Number.isSafeInteger(numericId) || numericId <= 0) {
    throw new Error('INVALID_MOVIE_ID')
  }

  const response = await fetch(`/api/v1/movies/${id}`, {
    signal: AbortSignal.timeout(10000),
  })

  if (response.status === 404) {
    throw new Error('MOVIE_NOT_FOUND')
  }

  if (!response.ok) {
    throw new Error('MOVIE_REQUEST_FAILED')
  }

  return response.json()
}