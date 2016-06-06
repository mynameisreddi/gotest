/* very eeasy and not optimized table, only for playing */
DROP TABLE listing;

CREATE TABLE listing (
    uid TEXT,
    minprice INT,
    title TEXT,
    country TEXT,
    region TEXT,
    city TEXT,
    street TEXT,
    postalcode TEXT,
    subarea TEXT,
    housenumber TEXT,
    housetype TEXT)
