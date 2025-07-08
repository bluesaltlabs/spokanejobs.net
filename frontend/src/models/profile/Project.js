class Project {
  constructor(data = {}) {
    this.id                   = data?.id || null;
    this.name                 = data?.name || null;
    this.description          = data?.description || null;
    this.url                  = data?.url || null;
    this.avatar_url           = data?.avatar_url || null;
    this.start_date           = data?.start_date || null;
    this.completion_date      = data?.completion_date || null;
    this.is_completed         = data?.completion_date || this?.completion_date || false;
    this.created_at           = data?.created_at || null;
    this.updated_at           = data?.updated_at || null;
  }

  //

}

export default Project;
