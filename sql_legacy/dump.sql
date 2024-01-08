BEGIN;

INSERT INTO episodes VALUES (1, 1, '2023-10-13');
INSERT INTO episodes VALUES (2, 2, '2023-10-20');
INSERT INTO episodes VALUES (3, 3, '2023-10-27');
INSERT INTO episodes VALUES (4, 4, '2023-10-27');
INSERT INTO episodes VALUES (5, 5, '2023-10-27');
INSERT INTO episodes VALUES (6, 5, '2023-10-27');
INSERT INTO episodes VALUES (7, 5, '2023-10-27');
INSERT INTO episodes VALUES (10, 7, '2023-11-01');
INSERT INTO episodes VALUES (11, 20, '2023-11-03');
INSERT INTO episodes VALUES (12, 15, '2023-11-24');


INSERT INTO tribes VALUES (3, 'Lulu', 'yellow');
INSERT INTO tribes VALUES (1, 'Reba', 'red');
INSERT INTO tribes VALUES (2, 'Belo', 'blue');

INSERT INTO players VALUES (2, 'Archie', 'Prodanovic', 12, 1, 'f');
INSERT INTO players VALUES (1, 'Marko', 'Prodanovic', 27, 3, 't');
INSERT INTO players VALUES (6, 'Jim', 'Jimminy', 20, 2, 'f');
INSERT INTO players VALUES (4, 'Bob', 'Apple', 25, 2, 't');
INSERT INTO players VALUES (9, 'New', 'Player', 15, 2, 'f');


INSERT INTO episode_points VALUES (1, 1, 1, 5);
INSERT INTO episode_points VALUES (2, 1, 2, 1);
INSERT INTO episode_points VALUES (3, 1, 4, 20);
INSERT INTO episode_points VALUES (4, 1, 6, 5);
INSERT INTO episode_points VALUES (9, 7, 1, 5);
INSERT INTO episode_points VALUES (10, 7, 2, 1);
INSERT INTO episode_points VALUES (11, 7, 4, 20);
INSERT INTO episode_points VALUES (12, 7, 6, 5);
INSERT INTO episode_points VALUES (21, 10, 2, 10);
INSERT INTO episode_points VALUES (22, 11, 2, 15);
INSERT INTO episode_points VALUES (23, 11, 1, 20);
INSERT INTO episode_points VALUES (24, 11, 6, 5);
INSERT INTO episode_points VALUES (25, 12, 4, 15);
INSERT INTO episode_points VALUES (26, 12, 6, 25);

INSERT INTO user_picks VALUES (1, 1, 1);
INSERT INTO user_picks VALUES (2, 1, 2);
INSERT INTO user_picks VALUES (3, 5, 6);
INSERT INTO user_picks VALUES (4, 6, 1);
INSERT INTO user_picks VALUES (5, 6, 2);
INSERT INTO user_picks VALUES (6, 6, 6);
INSERT INTO user_picks VALUES (7, 7, 1);
INSERT INTO user_picks VALUES (8, 7, 4);
INSERT INTO user_picks VALUES (9, 7, 2);
INSERT INTO user_picks VALUES (10, 7, 6);
INSERT INTO user_picks VALUES (11, 8, 6);
INSERT INTO user_picks VALUES (12, 9, 6);
INSERT INTO user_picks VALUES (13, 9, 1);
INSERT INTO user_picks VALUES (14, 9, 2);


INSERT INTO users VALUES (1, 'Marko Prodanovic', 'marko.prodanovic@outlook.com', 't');
INSERT INTO users VALUES (2, 'Marko Non Admin', NULL, 'f');
INSERT INTO users VALUES (5, 'Marko Prodanovic (FB)', 'marko.prodanovic@live.ca', 'f');

END;