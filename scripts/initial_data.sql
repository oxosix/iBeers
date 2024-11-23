-- Inserir valores iniciais para beer_type
INSERT INTO beer_type (name)
SELECT name
FROM (VALUES 
    ('Ale'),
    ('Lager'),
    ('Malt'),
    ('Stout')
) AS vt(name)
WHERE NOT EXISTS (
    SELECT 1 FROM beer_type WHERE name = vt.name
);


-- Inserir valores iniciais para beer_style
INSERT INTO beer_style (name)
SELECT name
FROM (VALUES 
    ('Amber'),
    ('Blonde'),
    ('Brown'),
    ('Cream'),
    ('Dark'),
    ('Pale'),
    ('Strong'),
    ('Wheat'),
    ('Red'),
    ('India Pale Ale'),
    ('Lime'),
    ('Pilsner'),
    ('Golden'),
    ('Fruit'),
    ('Honey')
) AS vs(name)
WHERE NOT EXISTS (
    SELECT 1 FROM beer_style WHERE name = vs.name
);
