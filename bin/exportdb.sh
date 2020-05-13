SCHEMA="public"
DB="later"

psql -Atc "select tablename from pg_tables where schemaname='$SCHEMA'" $DB |\
  while read TBL; do
    psql -c "COPY $SCHEMA.$TBL TO STDOUT WITH CSV DELIMITER ',' HEADER" $DB > $TBL.csv
  done