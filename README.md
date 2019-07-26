# Instalation and usage
1. You will need go version 1.12 or above
2. Set up your env variables in conf.yaml
3. Then run `go run main.go` 
4. After that you can go to *localhost:8080/blog-service/graphql* 

# Schema
This is the basic schema used in this demo

```
CREATE DATABASE blog;

USE blog;

CREATE TABLE `AUTHORS` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `POSTS` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `content` varchar(100) NOT NULL,
  `author_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `author_id` (`author_id`),
  CONSTRAINT `POSTS_ibfk_1` FOREIGN KEY (`author_id`) REFERENCES `AUTHORS` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `COMMENTS` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `content` text NOT NULL,
  `post_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `post_id` (`post_id`),
  CONSTRAINT `COMMENTS_ibfk_1` FOREIGN KEY (`post_id`) REFERENCES `POSTS` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

```
