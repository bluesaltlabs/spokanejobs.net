class Interest {
  constructor(data = {}) {
    this.id                   = data?.id || null;
    this.name                 = data?.name || null;
    this.description          = data?.description || null;
    this.years_of_experience  = data?.years_of_experience || null;
    this.created_at           = data?.created_at || null;
    this.updated_at           = data?.updated_at || null;
  }

  //
  // todo: this.years_of_experience should be selected from a list.

}

export default Interest;
