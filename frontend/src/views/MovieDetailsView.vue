<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import {
  getMovieDetails,
  type MovieDetails,
} from '@/services/movieService'

const route = useRoute()

const movie = ref<MovieDetails | null>(null)
const isLoading = ref(true)
const errorMessage = ref('')
const imageFailed = ref(false)
const retryCount = ref(0)

const posterUrl = computed(() =>
  movie.value?.poster_path
    ? `https://image.tmdb.org/t/p/w500${movie.value.poster_path}`
    : null,
)

const releaseYear = computed(
  () => movie.value?.release_date?.slice(0, 4) || 'Sin fecha',
)

const duration = computed(() => {
  const minutes = movie.value?.runtime

  if (!minutes || minutes <= 0) {
    return 'Duración no disponible'
  }

  const hours = Math.floor(minutes / 60)
  const remainingMinutes = minutes % 60

  return hours > 0
    ? `${hours} h ${remainingMinutes} min`
    : `${remainingMinutes} min`
})

watch(
  [() => route.params.id, retryCount],
  async ([id], _, onCleanup) => {
    let isCurrentRequest = true

    onCleanup(() => {
      isCurrentRequest = false
    })

    movie.value = null
    errorMessage.value = ''
    imageFailed.value = false
    isLoading.value = true

    try {
      const data = await getMovieDetails(String(id))

      if (isCurrentRequest) {
        movie.value = data
      }
    } catch (error) {
      if (!isCurrentRequest) return

      const code = error instanceof Error ? error.message : ''

      if (code === 'MOVIE_NOT_FOUND') {
        errorMessage.value = 'No encontramos esta película.'
      } else if (code === 'INVALID_MOVIE_ID') {
        errorMessage.value = 'El identificador de la película no es válido.'
      } else {
        errorMessage.value = 'No pudimos cargar la película. Intenta nuevamente.'
      }
    } finally {
      if (isCurrentRequest) {
        isLoading.value = false
      }
    }
  },
  { immediate: true },
)
</script>

<template>
  <main class="details-page">
    <RouterLink to="/" class="back-link">
      ← Volver a la cartelera
    </RouterLink>

    <p v-if="isLoading" class="status-message" role="status">
      Cargando película...
    </p>

    <section v-else-if="errorMessage" class="status-message" role="alert">
      <h1>No pudimos mostrar la ficha</h1>
      <p>{{ errorMessage }}</p>
      <button @click="retryCount++">Reintentar</button>
    </section>

    <article v-else-if="movie" class="movie-details">
      <img
        v-if="posterUrl && !imageFailed"
        :src="posterUrl"
        :alt="`Afiche de ${movie.title}`"
        class="poster"
        width="500"
        height="750"
        @error="imageFailed = true"
      />

      <div v-else class="poster poster-placeholder">
        Sin afiche
      </div>

      <div class="movie-content">
        <p class="eyebrow">FRAMELOG · FICHA DE PELÍCULA</p>

        <h1>{{ movie.title }}</h1>

        <p v-if="movie.original_title !== movie.title" class="muted">
          {{ movie.original_title }}
        </p>

        <p class="metadata">
          {{ releaseYear }} · {{ duration }}
        </p>

        <p v-if="movie.tagline" class="tagline">
          {{ movie.tagline }}
        </p>

        <ul v-if="movie.genres.length" class="genres" aria-label="Géneros">
          <li v-for="genre in movie.genres" :key="genre.id">
            {{ genre.name }}
          </li>
        </ul>

        <section class="rating" aria-label="Puntuación de TMDB">
          <template v-if="movie.vote_count > 0">
            <strong>{{ movie.vote_average.toFixed(1) }} / 10</strong>
            <span>TMDB · {{ movie.vote_count }} votos</span>
          </template>
          <span v-else>TMDB · Sin calificaciones</span>
        </section>

        <section class="synopsis">
          <h2>Sinopsis</h2>
          <p>{{ movie.overview || 'Sinopsis no disponible en este idioma.' }}</p>
        </section>
      </div>
    </article>

    <footer class="credits">
      <p>
        Datos e imágenes de
        <a href="https://www.themoviedb.org/">TMDB</a>.
      </p>
      <p>
        This product uses the TMDB API but is not endorsed or certified by TMDB.
      </p>
    </footer>
  </main>
</template>

<style scoped>
.details-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem 1.5rem;
}

.back-link {
  display: inline-block;
  margin-bottom: 2.5rem;
}

.movie-details {
  display: grid;
  grid-template-columns: minmax(200px, 320px) minmax(0, 1fr);
  gap: 3rem;
  align-items: start;
}

.poster {
  display: block;
  width: 100%;
  height: auto;
  aspect-ratio: 2 / 3;
  object-fit: cover;
  border-radius: 16px;
}

.poster-placeholder {
  display: grid;
  place-items: center;
  background: #242835;
  color: #b7bdcb;
}

.eyebrow {
  color: #f5b942;
  font-size: 0.75rem;
  letter-spacing: 0.12em;
}

h1 {
  margin: 0.75rem 0;
  font-size: clamp(2rem, 4vw, 3.5rem);
  line-height: 1.15;
}

.muted,
.metadata,
.tagline,
.credits {
  color: #b7bdcb;
}

.metadata,
.tagline {
  margin-top: 1rem;
}

.tagline {
  font-style: italic;
}

.genres {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin: 1.5rem 0;
  padding: 0;
  list-style: none;
}

.genres li {
  padding: 0.4rem 0.75rem;
  border: 1px solid #363c4b;
  border-radius: 999px;
  font-size: 0.85rem;
}

.rating {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
  margin: 1.5rem 0;
}

.rating strong {
  color: #f5b942;
  font-size: 1.75rem;
}

.synopsis h2 {
  margin-bottom: 0.75rem;
}

.synopsis p {
  line-height: 1.8;
}

.status-message {
  padding: 3rem 0;
}

.credits {
  margin-top: 3rem;
  padding-top: 1.5rem;
  border-top: 1px solid #2a2e3a;
  font-size: 0.8rem;
}

@media (max-width: 700px) {
  .movie-details {
    grid-template-columns: 1fr;
    gap: 2rem;
  }

  .poster {
    max-width: 280px;
    margin: 0 auto;
  }
}
</style>