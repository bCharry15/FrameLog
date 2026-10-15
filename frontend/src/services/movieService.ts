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

export async function getMovies(): Promise<MoviePage> {
  const response = await fetch('/api/v1/movies', {
    signal: AbortSignal.timeout(10000),
  })

  if (!response.ok) {
    throw new Error(`Movie request failed: ${response.status}`)
  }

  return response.json()
}