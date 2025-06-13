create table search_logs (
  id uuid not null default uuid_generate_v4(),
  search_term varchar(255) not null,
  filters jsonb not null default '{}',
  searched_at timestamp with time zone not null default now(),
  user_ip inet not null default '0.0.0.0',
  -- user_id uuid not null,
);
alter table search_logs enable row level security;

-- make the search_logs table data publicly readable
create policy "public can read search_logs"
on public.search_logs
for select to anon
using (true);
