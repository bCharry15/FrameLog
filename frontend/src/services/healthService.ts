export interface HealthResponse {
  status: string
  service: string
}

export async function getHealth(): Promise<HealthResponse> {
  const response = await fetch('/api/v1/health/live', {
    signal: AbortSignal.timeout(5000),
  })

  if (!response.ok) {
    throw new Error(`Health request failed: ${response.status}`)
  }

  const data: HealthResponse = await response.json()

  if (data.status !== 'ok' || data.service !== 'framelog-api') {
    throw new Error('Unexpected health response')
  }

  return data
}