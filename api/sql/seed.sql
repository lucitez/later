INSERT INTO users (username, name, email, phone_number, signed_up_at, password)
VALUES ('lucitez', 'Lucas Gregory', 'lucgreggs@gmail.com', '3109130692', now(), crypt('pass', gen_salt('bf')));

INSERT INTO users (username, name, email, phone_number, signed_up_at, password)
VALUES ('test', 'Test McTestFace', 'test@test.com', '1111111111', now(), crypt('pass', gen_salt('bf')));

INSERT INTO users (username, name, email, phone_number, signed_up_at, password)
VALUES ('boat', 'Boaty McBoatFace', 'angel@dust.com', '2222222222', now(), crypt('pass', gen_salt('bf')));

INSERT INTO users (username, name, email, phone_number, signed_up_at, password)
VALUES ('glump', 'Glumpo McSqueefy', 'glump@glump.co', '3333333333', now(), crypt('pass', gen_salt('bf')));

INSERT INTO users (username, name, email, phone_number, signed_up_at, password)
VALUES ('steven', 'Steve McQueen', 'smq@gmail.com', '9999999999', now(), crypt('pass', gen_salt('bf')));

-- create table content_types (id int, t text);
-- insert into content_types values (1, 'watch'), (2, 'read'), (3, 'listen');

-- CREATE OR REPLACE FUNCTION random_between(low INT ,high INT) 
--    RETURNS INT AS
-- $$
-- BEGIN
--    RETURN floor(random()* (high-low + 1) + low);
-- END;
-- $$ language 'plpgsql' STRICT;


-- DO
-- $do$
-- BEGIN
-- for i in 1..10000 LOOP
-- INSERT INTO user_content (share_id, content_id, user_id, sent_by_user_id, content_type, tag, saved_at)
-- VALUES (
--     uuid_generate_v4(),
--     uuid_generate_v4(),
--     uuid_generate_v4(),
--     uuid_generate_v4(),
--     (select random_between(0,3)::text),
--     (select random_between(0,15)::text),
--     now()
-- );
-- END LOOP;
-- END
-- $do$;

-- DO
-- $do$
-- BEGIN
-- for i in 1..1000 LOOP
-- INSERT INTO user_content (share_id, content_id, user_id, sent_by_user_id, content_type, tag, saved_at)
-- VALUES (
--     uuid_generate_v4(),
--     uuid_generate_v4(),
--     '1ef551dc-0faa-4c5c-bce8-a9cf80540de0',
--     uuid_generate_v4(),
--     (select random_between(0,3)::text),
--     (select random_between(0,15)::text),
--     now()
-- );
-- END LOOP;
-- END
-- $do$;

-- DO
-- $do$
-- BEGIN
-- for i in 1..1000 LOOP
-- INSERT INTO user_content (share_id, content_id, user_id, sent_by_user_id, content_type, tag)
-- VALUES (
--     uuid_generate_v4(),
--     uuid_generate_v4(),
--     '1ef551dc-0faa-4c5c-bce8-a9cf80540de0',
--     uuid_generate_v4(),
--     (select random_between(0,3)::text),
--     (select random_between(0,15)::text)
-- );
-- END LOOP;
-- END
-- $do$;

-- DO
-- $do$
-- BEGIN
-- for i in 1..10000 LOOP
-- INSERT INTO user_content (share_id, content_id, user_id, sent_by_user_id, content_type, tag)
-- VALUES (
--     uuid_generate_v4(),
--     uuid_generate_v4(),
--     uuid_generate_v4(),
--     uuid_generate_v4(),
--     (select random_between(0,3)::text),
--     (select random_between(0,15)::text)
-- );
-- END LOOP;
-- END
-- $do$;