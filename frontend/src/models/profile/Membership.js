class Membership {
  constructor(data = {}) {
    this.id                   = data?.id || null;
    this.name                 = data?.name || null;
    this.organization         = data?.organization || null;
    this.description          = data?.description || null;
    this.start_date           = data?.start_date || null;
    this.end_date             = data?.end_date || null;
    this.created_at           = data?.created_at || null;
    this.updated_at           = data?.updated_at || null;
  }

  //

}

export default Membership;
