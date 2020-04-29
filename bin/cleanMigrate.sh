dropdb later
createdb later

for i in ../sql/[0-9]*; do
    psql -f $i later
done

psql -f ../sql/setup.sql later
psql -f ../sql/seed.sql later
