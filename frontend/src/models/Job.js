class Job {
  constructor(data = {}) {
    this.id             = data?.id || null;
    this.title          = data?.title || null;
    this.description    = data?.description || null;
    this.city           = data?.city || null;
    this.state          = data?.state || null;
    this.url            = data?.url || null;
    this.company_slug   = data?.company_slug || null;
    this.location_slug  = data?.location_slug || null;
    this.is_active      = data?.is_active || false;
    this.created_at     = data?.startDate || null;
    this.updated_at     = data?.updated_at || null;
  }

  //

}

export default Job;
