<script setup lang="ts">
import { onMounted, ref } from 'vue'
import MovieCard from '@/components/MovieCard.vue'
import { getMovies, type Movie } from '@/services/movieService'

const movies = ref<Movie[]>([])
const isLoading = ref(true)
const errorMessage = ref('')

async function loadMovies() {
  isLoading.value = true
  errorMessage.value = ''

  try {
    const data = await getMovies()
    movies.value = data.results
  } catch {
    errorMessage.value = 'No pudimos cargar las películas. Intenta nuevamente.'
  } finally {
    isLoading.value = false
  }
}

onMounted(loadMovies)
</script>

<template>
  <div class="page">
    <header class="site-header">
      <span class="brand">FrameLog</span>
      <span>Descubre. Mira. Comparte.</span>
    </header>

    <main>
      <section class="hero">
        <p class="eyebrow">TU ESPACIO PARA EL CINE</p>
        <h1>Cada película deja una historia.</h1>
        <p>Encuentra tu próxima película y empieza a construir tu diario de cine.</p>
      </section>

      <section aria-labelledby="movies-heading" :aria-busy="isLoading">
        <h2 id="movies-heading">Películas populares</h2>

        <p v-if="isLoading" role="status">Cargando cartelera...</p>

        <div v-else-if="errorMessage" role="alert">
          <p>{{ errorMessage }}</p>
          <button @click="loadMovies">Reintentar</button>
        </div>

        <p v-else-if="movies.length === 0">No hay películas disponibles.</p>

        <div v-else class="movie-grid">
          <MovieCard
            v-for="movie in movies"
            :key="movie.id"
            :movie="movie"
          />
        </div>
      </section>
    </main>

    <footer>
      <h2>Créditos</h2>
      <p>
        Datos e imágenes de
        <a href="https://www.themoviedb.org/">TMDB</a>.
      </p>
      <p>
        This product uses the TMDB API but is not endorsed or certified by TMDB.
      </p>
    </footer>
  </div>
</template>

<style scoped>
.page {
  max-width: 1280px;
  margin: 0 auto;
  padding: 0 1.5rem;
}

.site-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 1rem;
  padding: 1.5rem 0;
  border-bottom: 1px solid #2a2e3a;
  color: #b7bdcb;
}

.brand {
  color: #f5b942;
  font-size: 1.75rem;
  font-weight: 800;
}

.hero {
  padding: 4rem 0;
  max-width: 760px;
}

.eyebrow {
  color: #f5b942;
  letter-spacing: 0.15em;
  font-size: 0.75rem;
}

h1 {
  margin: 1rem 0;
  font-size: clamp(2rem, 5vw, 4rem);
  line-height: 1.1;
}

.hero > p:last-child {
  color: #b7bdcb;
  line-height: 1.7;
}

h2 {
  margin-bottom: 1.5rem;
}

.movie-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
  gap: 1.5rem;
}

footer {
  margin-top: 4rem;
  padding: 2rem 0;
  border-top: 1px solid #2a2e3a;
  color: #b7bdcb;
  font-size: 0.8rem;
}

footer h2 {
  margin-bottom: 0.75rem;
  font-size: 1rem;
}
</style>