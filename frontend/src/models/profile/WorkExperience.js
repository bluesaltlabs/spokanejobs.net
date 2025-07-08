class WorkExperience {
  constructor(data = {}) {
    this.id                 = data?.id || null;
    this.employer           = data?.employer || null;
    this.company_slug       = data?.company_slug || null;
    this.job_title_start    = data?.job_title_start || null;
    this.job_title_end      = data?.job_title_end || null;
    this.start_date         = data?.start_date || null;
    this.end_date           = data?.end_date || null;
    this.is_current         = data?.is_current || null;
    this.responsibilities   = data?.responsibilities || null;
    this.comments           = data?.comments || null;
    this.reason_for_leaving = data?.reason_for_leaving || null;
    this.can_contact        = data?.can_contact || null;
    this.supervisor_name    = data?.supervisor_name || null;
    this.supervisor_email   = data?.supervisor_email || null;
    this.supervisor_phone   = data?.supervisor_phone || null;
    this.supervisor_title   = data?.supervisor_title || null;
    this.created_at         = data?.created_at || null;
    this.updated_at         = data?.updated_at || null;
  }

  //

}

export default WorkExperience;
