const companiesUrl = import.meta.env.VITE_COMPANIES_URL ?? '/api/companies.json'

export const githubClient = {
  async fetchCompanies() {
    try {
      const url = companiesUrl

      console.debug(`Fetching companies from ${url}`)

      const response = await fetch(url)
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      const data = await response.json()
      return { data, error: null }
    } catch (error) {
      console.error('Error fetching companies from GitHub:', error)

      return { data: null, error }
    }
  }
}
