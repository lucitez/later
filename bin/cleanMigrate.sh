dropdb later
createdb later

for i in ../sql/[0-9]*; do
    psql -f $i later
done

for i in ../seed/*.csv; do
    filename=$(basename $i)
    tablename="${filename%.*}"
    thing="$(cd "$(dirname "$i")"; pwd -P)/$(basename "$i")"

    psql later -c "COPY $tablename FROM '$thing' CSV DELIMITER ',' HEADER;"
done

psql -f ../sql/setup.sql later
