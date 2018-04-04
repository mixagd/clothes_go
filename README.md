# clothes_go

To create the postgres database:

    docker run -it --rm --name Clothes -e POSTGRES_PASSWORD=clothes -e POSTGRES_USER=clothes -e POSTGRES_DB=clothes  -v  $(pwd)/pgdata:/var/lib/postgresql/data -p 5432:5432 -d postgres

    docker exec -it Clothes /bin/bash

    psql -U clothes clothes

    CREATE TABLE CLOTHES(ID CHAR PRIMARY KEY NOT NULL, TYPE TEXT NOT NULL, COLOUR TEXT NOT NULL, FIT TEXT, OWNER CHAR(50));

    INSERT INTO CLOTHES (ID, TYPE, COLOUR, FIT, OWNER) 
    VALUES 
    ('1', 'Blouse', 'Blue', 'Tight', 'Michaila'), 
    ('2', 'Blouse', 'White', 'Loose', 'Maria'), 
    ('3', 'Blouse', 'Black', 'Loose', 'Maria'), 
    ('4', 'Jeans', 'Blue', 'Tight', 'Michaila'), 
    ('5', 'Trousers', 'Beige', 'Tight', 'Katerina');
