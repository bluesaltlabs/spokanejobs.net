class LicenseCertification {
  constructor(data = {}) {
    this.id                   = data?.id || null;
    this.name                 = data?.name || null;
    this.organization         = data?.organization || null;
    this.description          = data?.description || null;
    this.date_earned          = data?.date_earned || null;
    this.date_expires         = data?.date_expires || null;
    this.created_at           = data?.created_at || null;
    this.updated_at           = data?.updated_at || null;
  }

  //

}

export default LicenseCertification;
