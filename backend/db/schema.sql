create table if not exists user (
    user_id int not null auto_increment primary key,
    google_id varchar(128) not null,
    email varchar(128) not null,
    first_name varchar(128),
    last_name varchar(128),
    photo_url varchar(512) 
);

create table if not exists category (
    category_id int not null auto_increment primary key,
    category_name varchar(128) not null
);

create table if not exists saved_category (
    sc_id int not null auto_increment primary key,
    category_id int not null,
    user_id int not null,
    foreign key (user_id) REFERENCES user(user_id),
    foreign key (category_id) REFERENCES category(category_id)
);

ALTER TABLE saved_category ADD UNIQUE `unique_index`(category_id, user_id);

create table if not exists `group` (
    group_id int not null auto_increment primary key,
    user_id int not null,
    category_id int not null,
    group_name varchar(128) not null,
    group_description TEXT,
    created_at DATETIME,
    foreign key (user_id) REFERENCES user(user_id),
    foreign key (category_id) REFERENCES category(category_id)
);

create table if not exists group_comment (
    gc_id int not null auto_increment primary key,
    user_id int not null,
    group_id int not null,
    reply_id int,
    comment_content TEXT not null,
    created_at DATETIME,
    deleted boolean,
    foreign key (user_id) REFERENCES user(user_id),
    foreign key (group_id) REFERENCES `group`(group_id),
    foreign key (reply_id) REFERENCES group_comment(gc_id)
);

create table if not exists blog_post (
    bp_id int not null auto_increment primary key,
    user_id int not null,
    group_id int not null,
    post_title varchar(128) not null,
    post_content TEXT not null,
    created_at DATETIME,
    foreign key (user_id) REFERENCES user(user_id),
    foreign key (group_id) REFERENCES `group`(group_id)
);

create table if not exists blog_comment (
    bc_id int not null auto_increment primary key,
    user_id int not null,
    bp_id int not null,
    reply_id int,
    comment_content TEXT not null,
    created_at DATETIME,
    deleted boolean,
    foreign key (user_id) REFERENCES user(user_id),
    foreign key (bp_id) REFERENCES blog_post(bp_id) ON DELETE CASCADE,
    foreign key (reply_id) REFERENCES blog_comment(bc_id) ON DELETE CASCADE
);

create table if not exists saved_group (
    sg_id int not null auto_increment primary key,
    user_id int not null,
    group_id int not null,
    foreign key (user_id) REFERENCES user(user_id),
    foreign key (group_id) REFERENCES `group`(group_id)
);

ALTER TABLE saved_group ADD UNIQUE `unique_index`(user_id, group_id);

create table if not exists membership (
    membership_id int not null auto_increment primary key,
    user_id int not null,
    group_id int not null,
    updated_at DATETIME,
    state boolean not null, 
    foreign key (user_id) REFERENCES user(user_id),
    foreign key (group_id) REFERENCES `group`(group_id)
);

ALTER TABLE membership ADD UNIQUE `unique_index`(user_id, group_id);


insert into user(google_id, email, first_name, last_name, photo_url)
values('118379264076819254762', 'cahillaw@uw.edu', 'Andy', 'Cahill','https://lh5.googleusercontent.com/-kwF_3vBC36g/AAAAAAAAAAI/AAAAAAAAAAA/AMZuucm_8vhE0mrs6Kzu80ATCl5zV4P8jA/s96-c/photo.jpg');

insert into category(category_name)
values('DIY'),('Movies & TV'),('Nature'),('Camping'),('Card & Board Games'),('Music'),('Cooking'),('Investing'),('Volunteering'),('Yoga & Mindfulness'),('Writing'),('Reading'),('Museums'),('Sports'),('Exercise'),('Hiking'),('Podcasts'),('Cycling'),('Technology'),('Programming'),('Photography & Video'),('Writing'),('Drawing & Painting'),('Animals & Pets'),('Cars'),('Fashion'),('Shopping'),('Collecting'),('Traveling'),('Home Furnishing'),('Bird Watching'),('Crafts'),('Bodybuilding'),('Tutoring'),('Poetry'),('Internships'),('Debate'),('Gaming'),('Content Creation'),('Academics'),('Food'),('Design'),('Social'),('Engineering'),('Architecture');