
CREATE TABLE domainsV2 (
        id STRING PRIMARY KEY,
        requestDomain STRING,
        requestHash STRING,
        previousGrade STRING,
        created_at int64,
        updated_at int64
);