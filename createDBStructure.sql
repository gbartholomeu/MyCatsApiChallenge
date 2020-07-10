CREATE DATABASE cats_api;
USE cats_api;
CREATE TABLE cats_breeds (id VARCHAR(255) NOT NULL, breed_name VARCHAR(255) NOT NULL, weight_imperial VARCHAR(255), 
weight_metric VARCHAR(255), cfa_url VARCHAR(255), vet_street_url VARCHAR(255), vca_hospitals_url VARCHAR(255), 
temperament VARCHAR(255), origin VARCHAR(255), country_codes VARCHAR(2), country_code VARCHAR(2), breed_description VARCHAR(500), 
lifes_span VARCHAR(255), indoor SMALLINT, lap SMALLINT, alt_name VARCHAR(255), adaptability SMALLINT, affection_level SMALLINT, 
child_friendly SMALLINT, dog_friendly SMALLINT, energy_level SMALLINT, grooming SMALLINT, health_issues SMALLINT, intelligence SMALLINT, 
shedding_level SMALLINT, social_needs SMALLINT, stranger_friendly SMALLINT, vocalisation SMALLINT, experimental SMALLINT, 
hairless SMALLINT, breed_natural SMALLINT, rare SMALLINT, rex SMALLINT, suppressed_tail SMALLINT, short_legs SMALLINT, 
wikipedia_url VARCHAR(255), hypoallergenic SMALLINT, PRIMARY KEY (id));