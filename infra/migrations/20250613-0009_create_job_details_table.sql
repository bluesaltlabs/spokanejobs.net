create table job_details (
  job_id uuid primary key not null references jobs(id),
  full_description text not null,
  applicant_count integer not null default 0,
  view_count integer not null default 0,
  requirements text not null,
  responsibilities text not null,
  -- what other detail data could be included here?
  updated_at timestamp with time zone not null default now()
);
alter table job_details enable row level security;

-- make the job_details table data publicly readable
create policy "public can read job_details"
on public.job_details
for select to anon
using (true);
