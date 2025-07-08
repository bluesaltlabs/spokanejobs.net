class PersonalInformation {
  constructor(data = {}) {
    this.first_name     = data?.first_name || '';
    this.middle_name    = data?.middle_name || '';
    this.last_name      = data?.last_name || '';
    this.preferred_name = data?.preferred_name || '';
    this.email          = data?.email || '';
    this.phone          = data?.phone || '';
    this.avatar_url     = data?.avatar_url || '';
    this.created_at     = data?.created_at || null;
    this.updated_at     = data?.updated_at || null;
    this.urls           = data.urls || [];
  }
}

export default PersonalInformation;
