class Location {
  constructor(data = {}) {
    this.id             = data?.id || null;
    this.city           = data?.city || null;
    this.state          = data?.state || null;
    this.street_address = data?.street_address || null;
    this.zip_code       = data?.zip_code || null;
    this.created_at     = data?.startDate || null;
    this.updated_at     = data?.updated_at || null;
  }

  //

}

export default Location;
