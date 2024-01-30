-- initial_data.sql

-- Inserir valores iniciais para beer_type
INSERT INTO beer_type (name) VALUES
    ('Ale'),
    ('Lager'),
    ('Malt'),
    ('Stout')
ON CONFLICT (name) DO NOTHING;

-- Inserir valores iniciais para beer_style
INSERT INTO beer_style (name) VALUES
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
ON CONFLICT (name) DO NOTHING;