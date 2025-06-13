-- Create the table
create table companies (
  id uuid not null primary key default uuid_generate_v4(),
  name varchar(255) not null,
  slug varchar(255) not null,
  description text,
  website text,
  job_board_url text,
  logo_url text,
  industry text,
  created_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now()
);
alter table companies enable row level security;

-- make the companies table data publicly readable
create policy "public can read companies"
on public.companies
for select to anon
using (true);
