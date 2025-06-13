create table job_details (
  job_id uuid not null references jobs(id),
  full_description text not null,
  applicant_count integer not null default 0,
  view_count integer not null default 0,
  requirements text not null,
  responsibilities text not null,
  -- what other detail data could be included here?
  updated_at timestamp not null default now()
);

-- todo: cascade delete when job is deleted
