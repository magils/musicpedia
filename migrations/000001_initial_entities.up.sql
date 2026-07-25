CREATE TABLE genres (
    id int AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE performers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    bio TEXT NOT NULL,
    country VARCHAR(255) NOT NULL,
    formed_year INT NOT NULL,
    is_band BOOLEAN NOT NULL DEFAULT FALSE,
    picture_url VARCHAR(2048),
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW()
);

CREATE TABLE band_members (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    role VARCHAR(255),
    performer_id INT NOT NULL,
    CONSTRAINT FOREIGN KEY (performer_id) REFERENCES performers(id)
);


CREATE TABLE albums (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    release_year INT not null,
    record_label VARCHAR(255) NOT NULL,
    cover_picture_url VARCHAR(2948) NOT NULL,
    total_tracks INT not null,
    performer_id INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW(),
    CONSTRAINT FOREIGN KEY(performer_id) REFERENCES performers(id)
);

CREATE TABLE album_genres (
    album_id INT NOT NULL,
    genre_id INT NOT NULL,
    PRIMARY KEY (album_id, genre_id),
    CONSTRAINT FOREIGN KEY (album_id) REFERENCES albums(id) ON DELETE CASCADE,
    CONSTRAINT FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
);


CREATE TABLE tracks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    track_number INT,
    duration_seconds INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW(),
    album_id INT NOT NULL,
    CONSTRAINT FOREIGN KEY(album_id) REFERENCES albums(id)
);