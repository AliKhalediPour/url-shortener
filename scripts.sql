CREATE TABLE IF NOT EXISTS urls (
  id 	uuid,
  short text not null,
  long text not null,
  primary key (id)
);

CREATE unique index IF NOT EXISTS urls_unique_short on urls (short);
