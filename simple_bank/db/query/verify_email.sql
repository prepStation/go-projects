-- name: CreateVerifyEmail :one
INSERT INTO verify_email(
    username, email, secret_code
)VALUES(
    $1, $2, $3
)RETURNING *;

-- name: UpdateVerifyEmail :one
UPDATE verify_email
SET
    is_used = TRUE
WHERE 
    id = @id
    AND secret_code = @secret_code
    AND is_used = FALSE
    AND expired_at > now()
RETURNING *;

-- SELECT genre.name, genre.id, COUNT(track.id) as track_count, (SELECT COUNT(*) FROM (
--     SELECT genre.name, genre.id, COUNT(track.id) as track_count
--     FROM genre
--     INNER JOIN track ON track.genre_id = genre.id
--     GROUP BY genre.id
-- ) as subquery) as total_records
-- FROM genre
-- INNER JOIN track ON track.genre_id = genre.id
-- GROUP BY genre.id;


-- SELECT
--     genre.name,
--     genre.id,
--     COUNT(track.id) as track_count,
--     SUM(COUNT(track.id)) OVER() as total_records
-- FROM
--     genre
-- INNER JOIN
--     track ON track.genre_id = genre.id
-- GROUP BY
--     genre.id, genre.name;



