create table jobs (
  id uuid not null primary key default uuid_generate_v4(),
  title varchar(255) not null,
  slug varchar(255) not null,
  description text not null,
  employment_type employment_type not null,
  salary_min integer not null,
  salary_max integer not null,
  url text,
  company_id bigint not null references companies(id),
  location_id uuid references locations(id),
  category_id bigint references categories(id),
  scraped_at timestamp with time zone default now(),
  posted_at timestamp with time zone not null default now(),
  created_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now()
);
alter table jobs enable row level security;

-- make the jobs table data publicly readable
create policy "public can read jobs"
on public.jobs
for select to anon
using (true);
