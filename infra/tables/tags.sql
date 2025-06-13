create table tags (
  id bigint not null primary key,
  name varchar(255) not null unique,
  slug varchar(255) not null unique,
  description text,
  created_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now()
  );
  alter table tags enable row level security;

  -- make the tags table data publicly readable
  create policy "public can read tags"
  on public.tags
  for select to anon
  using (true);
