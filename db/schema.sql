CREATE TABLE fetcherjob 
(
    id SERIAL PRIMARY KEY,
    report_title TEXT NOT NULL,
    search_engine VARCHAR(50),
    location VARCHAR(250),
    language VARCHAR(100)
);

CREATE TABLE combinedjob
(
    id SERIAL PRIMARY KEY,
    created TIMESTAMP,
    report_title TEXT NOT NULL,
    grouping_method VARCHAR(50),
    main_keyword_grouping_accuracy INT,
    variant_keyword_grouping_accuracy INT,
    search_engine VARCHAR(50),
    location VARCHAR(250),
    language VARCHAR(100)
);

CREATE TABLE grouperjob
(
    id SERIAL PRIMARY KEY,
    created TIMESTAMP,
    report_title TEXT NOT NULL,
    grouping_method VARCHAR(50),
    main_keyword_grouping_accuracy INT NOT NULL,
    variant_keyword_grouping_accuracy INT NOT NULL
);
