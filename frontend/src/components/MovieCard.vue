<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { RouterLink } from 'vue-router'
import type { Movie } from '@/services/movieService'

const props = defineProps<{
  movie: Movie
}>()

const imageFailed = ref(false)

const posterUrl = computed(() =>
  props.movie.poster_path
    ? `https://image.tmdb.org/t/p/w500${props.movie.poster_path}`
    : null,
)

const releaseYear = computed(
  () => props.movie.release_date?.slice(0, 4) || 'Sin fecha',
)

watch(posterUrl, () => {
  imageFailed.value = false
})
</script>

<template>
  <article class="movie-card">
    <RouterLink
      :to="{ name: 'movie-details', params: { id: movie.id } }"
      class="movie-link"
    >
      <img
        v-if="posterUrl && !imageFailed"
        :src="posterUrl"
        :alt="`Afiche de ${movie.title}`"
        class="movie-poster"
        loading="lazy"
        width="500"
        height="750"
        @error="imageFailed = true"
      />

      <div v-else class="movie-poster poster-placeholder">
        Sin afiche
      </div>

      <div class="movie-info">
        <h3>{{ movie.title }}</h3>
        <p>{{ releaseYear }}</p>

        <p class="movie-rating">
          {{
            movie.vote_count > 0
              ? `TMDB · ${movie.vote_average.toFixed(1)} / 10`
              : 'TMDB · Sin calificaciones'
          }}
        </p>
      </div>
    </RouterLink>
  </article>
</template>

<style scoped>
.movie-card {
  overflow: hidden;
  background: #191c24;
  border: 1px solid #2a2e3a;
  border-radius: 14px;
}

.movie-link {
  display: block;
  height: 100%;
  color: inherit;
  text-decoration: none;
}

.movie-link:hover {
  background: #202430;
}

.movie-link:focus-visible {
  outline: 3px solid #f5b942;
  outline-offset: -3px;
  border-radius: 14px;
}

.movie-poster {
  display: block;
  width: 100%;
  height: auto;
  aspect-ratio: 2 / 3;
  object-fit: cover;
}

.poster-placeholder {
  display: grid;
  place-items: center;
  background: #242835;
  color: #b7bdcb;
}

.movie-info {
  padding: 1rem;
}

h3 {
  margin-bottom: 0.5rem;
  font-size: 1rem;
}

p {
  color: #b7bdcb;
  font-size: 0.875rem;
}

.movie-rating {
  margin-top: 0.75rem;
  color: #f5b942;
}
</style>