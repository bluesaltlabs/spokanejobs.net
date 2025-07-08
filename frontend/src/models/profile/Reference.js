class Reference {
  constructor(data = {}) {
    this.id                   = data?.id || null;
    this.first_name           = data?.first_name || null;
    this.last_name            = data?.last_name || null;
    this.relationship         = data?.relationship || null;
    this.email                = data?.email || null;
    this.phone                = data?.phone || null;
    this.email_type           = data?.email_type || null;
    this.phone_type           = data?.phone_type || null;
    this.description          = data?.description || null;
    this.employer             = data?.employer || null;
    this.company_slug         = data?.company_slug || null;
    this.job_title            = data?.job_title || null;
    this.created_at           = data?.created_at || null;
    this.updated_at           = data?.updated_at || null;
  }

  //

}

export default Reference;
