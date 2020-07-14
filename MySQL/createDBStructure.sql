USE cats_api;
CREATE TABLE cats_breeds (id VARCHAR(255) NOT NULL, breed_name VARCHAR(255) NOT NULL, temperament VARCHAR(255), origin VARCHAR(255), 
breed_description VARCHAR(500),  PRIMARY KEY (id));
CREATE TABLE cats_images (id VARCHAR(255) NOT NULL, breed_id VARCHAR(255) NOT NULL, breed_name VARCHAR(255) NOT NULL,  image_url VARCHAR(255) NOT NULL, PRIMARY KEY (id), INDEX breed_ind (breed_id), FOREIGN KEY (breed_id) REFERENCES cats_breeds(id) ON DELETE CASCADE);
CREATE TABLE stylish_cats_images (id VARCHAR(255) NOT NULL, image_url VARCHAR(255) NOT NULL, has_glasses BOOLEAN, has_hat BOOLEAN, PRIMARY KEY (id));