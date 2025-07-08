class Contact {
  constructor(data = {}) {
    this.id             = data?.id || null;
    this.first_name     = data?.first_name || null;
    this.last_name      = data?.last_name || null;
    this.email          = data?.email || null;
    this.phone          = data?.phone || null;
    this.job_title      = data?.job_title || null;
    this.company_slug   = data?.company_slug || null;
    this.created_at     = data?.startDate || null;
    this.updated_at     = data?.updated_at || null;
  }

  //

}

export default Contact;
