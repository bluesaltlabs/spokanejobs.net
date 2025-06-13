-- Create the table
create table companies (
  id bigint primary key generated always as identity,
  name text not null,
  slug text not null,
  description text,
  website text,
  job_board_url text,
  logo_url text,
  industry text,
  created_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now()
);
-- Insert some sample data into the table
insert into companies (name, slug)
values
  ('Limelyte', 'limelyte'),
  ('CDL Powersuite', 'cdl-powersuite'),
  ('OpenEye', 'openeye');
alter table companies enable row level security;

-- make the companies table data publicly readable
create policy "public can read companies"
on public.companies
for select to anon
using (true);
