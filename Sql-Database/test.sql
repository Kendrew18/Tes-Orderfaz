-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 10, 2024 at 07:03 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test`
--

-- --------------------------------------------------------

--
-- Table structure for table `logistic`
--

CREATE TABLE `logistic` (
  `id` int(11) NOT NULL,
  `logistic_name` varchar(255) NOT NULL,
  `amount` int(11) NOT NULL,
  `destination_name` varchar(255) NOT NULL,
  `origin_name` varchar(255) NOT NULL,
  `duration` varchar(255) NOT NULL,
  `uuid_user` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `logistic`
--

INSERT INTO `logistic` (`id`, `logistic_name`, `amount`, `destination_name`, `origin_name`, `duration`, `uuid_user`) VALUES
(1, 'JNT', 20000, 'Surabaya', 'Bandung', '2-4', '0cb31b31-ce3e-4fdc-8f9c-f508cf2abe08'),
(2, 'JNE', 20000, 'Surabaya', 'Bandung', '2-4', '0cb31b31-ce3e-4fdc-8f9c-f508cf2abe08');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `uuid` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `msisdn` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`uuid`, `name`, `msisdn`, `username`, `password`) VALUES
('0cb31b31-ce3e-4fdc-8f9c-f508cf2abe08', 'Alex', '6281332226828', 'densu', '2abVt8XWttrzPWeYmZI8xtWYdxgGyX7hSOQEdhnyE31y0b0=');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `logistic`
--
ALTER TABLE `logistic`
  ADD PRIMARY KEY (`id`),
  ADD KEY `content_FK_1` (`uuid_user`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`uuid`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `msisdn` (`msisdn`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `logistic`
--
ALTER TABLE `logistic`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `logistic`
--
ALTER TABLE `logistic`
  ADD CONSTRAINT `content_FK_1` FOREIGN KEY (`uuid_user`) REFERENCES `user` (`uuid`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
