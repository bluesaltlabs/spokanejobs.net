const COMPANIES_URL = import.meta.env.COMPANIES_URL

export const githubClient = {
  async fetchCompanies() {
    try {
      // Use local proxy in development, direct URL in production
      const url = import.meta.env.DEV ? '/api/companies.json' : COMPANIES_URL

      const response = await fetch(url)
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      const data = await response.json()
      return { data, error: null }
    } catch (error) {
      console.error('Error fetching companies from GitHub:', error)

      // Fallback to local JSON file in development
      if (import.meta.env.DEV) {
        try {
          console.log('Falling back to local companies.json file...')
          const localResponse = await fetch('/companies.json')
          if (localResponse.ok) {
            const localData = await localResponse.json()
            return { data: localData, error: null }
          }
        } catch (localError) {
          console.error('Error fetching local companies.json:', localError)
        }
      }

      return { data: null, error }
    }
  }
}
