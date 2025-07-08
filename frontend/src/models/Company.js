class Company {
  constructor(data = {}) {
    this.id               = data?.id || null;
    this.slug             = data?.slug || null;
    this.name             = data?.name || null;
    this.description      = data?.description || null;
    this.logo_url         = data?.logo_url || null;
    this.website_url      = data?.website_url || null;
    this.job_board_url    = data?.job_board_url || null;
    this.has_scraper      = data?.has_scraper || false;
    this.last_scraped_at  = data?.last_scraped_at || null;
    this.next_scrape_at   = data?.next_scrape_at || null;
    this.created_at       = data?.startDate || null;
    this.updated_at       = data?.updated_at || null;
  }

  //

}

export default Company;
