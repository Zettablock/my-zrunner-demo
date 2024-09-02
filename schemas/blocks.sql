CREATE TABLE blocks
(
    number              integer        NOT NULL,
    hash                VARCHAR(100)   NOT NULL,
    parent_hash         VARCHAR(100)   NOT NULL,
    PRIMARY KEY (number)
);

CREATE TABLE blocks_rpc
(
    number              integer        NOT NULL,
    hash                VARCHAR(100)   NOT NULL,
    parent_hash         VARCHAR(100)   NOT NULL,
    PRIMARY KEY (number)
);