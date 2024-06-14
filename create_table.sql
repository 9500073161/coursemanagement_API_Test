CREATE DATABASE IF NOT EXISTS `coursemanagement` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `coursemanagement`;

CREATE TABLE IF NOT EXISTS `courses` (
    `id` int NOT NULL COMMENT 'Id',
    `name` varchar(255) NOT NULL COMMENT 'Name',
    `teacher_id` int NOT NULL COMMENT 'TeacherID',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=`utf8`;

CREATE TABLE IF NOT EXISTS `students` (
    `id` int NOT NULL COMMENT 'Id',
    `name` varchar(255) NOT NULL COMMENT 'Name',
     PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=`utf8`;

CREATE TABLE IF NOT EXISTS `teachers` (
    `id` int NOT NULL COMMENT 'Id',
    `name` varchar(255) NOT NULL COMMENT 'Name',
     PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=`utf8`;

CREATE TABLE IF NOT EXISTS `entrollments` (
    `id` int NOT NULL COMMENT 'Id',
    `course_id` int NOT NULL COMMENT 'Courseid',
	`student_id` int NOT NULL COMMENT 'Studentid',
     PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=`utf8`;        
