class EducationExperience {
  constructor(data = {}) {
    this.id                 = data?.id || null;
    this.institution        = data?.institution || null;
    this.location           = data?.location || null;
    this.major              = data?.major || null;
    this.minor              = data?.minor || null;
    this.area_of_study      = data?.area_of_study || null;
    this.credits_earned     = data?.credits_earned || null;
    this.credit_type        = data?.credit_type || null;
    this.gpa                = data?.gpa || null;
    this.start_date         = data?.start_date || null;
    this.end_date           = data?.end_date || null;
    this.completion_date    = data?.completion_date || null;
    this.description        = data?.description || null;
    this.honors_recognition = data?.honors_recognition || [];
  }

  //

}

export default EducationExperience;
