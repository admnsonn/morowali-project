-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 15, 2023 at 05:51 PM
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
-- Database: `morowali`
--

-- --------------------------------------------------------

--
-- Table structure for table `desa`
--

CREATE TABLE `desa` (
  `desa_id` int(11) NOT NULL,
  `Nama_desa` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `desa`
--

INSERT INTO `desa` (`desa_id`, `Nama_desa`) VALUES
(1, 'Bahomoleo');

-- --------------------------------------------------------

--
-- Table structure for table `pengguna`
--

CREATE TABLE `pengguna` (
  `id_pengguna` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role_pengguna` varchar(50) NOT NULL,
  `role_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `created_by` varchar(100) NOT NULL,
  `update_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_by` varchar(100) NOT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `tempat_lahir` varchar(100) DEFAULT NULL,
  `jk_id` int(11) NOT NULL,
  `jenis_kelamin` varchar(20) DEFAULT NULL,
  `status_perkawinan_id` int(11) DEFAULT NULL,
  `status_perkawinan` varchar(255) DEFAULT NULL,
  `profesi` varchar(100) DEFAULT NULL,
  `nik` int(11) DEFAULT NULL,
  `kk` int(11) DEFAULT NULL,
  `umur` int(11) DEFAULT NULL,
  `kategori_financial_id` int(11) DEFAULT NULL,
  `agama` varchar(50) DEFAULT NULL,
  `agama_id` int(11) NOT NULL,
  `desa` varchar(100) NOT NULL,
  `nama_lengkap` varchar(100) NOT NULL,
  `foto_pengguna` text DEFAULT NULL,
  `alamat_id` int(11) NOT NULL,
  `alamat_pengguna` text DEFAULT NULL,
  `rt` varchar(10) NOT NULL,
  `rw` varchar(10) NOT NULL,
  `kode_pos` varchar(10) NOT NULL,
  `no_telp` varchar(20) NOT NULL,
  `token` text NOT NULL,
  `desa_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pengguna`
--

INSERT INTO `pengguna` (`id_pengguna`, `username`, `password`, `role_pengguna`, `role_id`, `created_at`, `created_by`, `update_at`, `updated_by`, `tanggal_lahir`, `tempat_lahir`, `jk_id`, `jenis_kelamin`, `status_perkawinan_id`, `status_perkawinan`, `profesi`, `nik`, `kk`, `umur`, `kategori_financial_id`, `agama`, `agama_id`, `desa`, `nama_lengkap`, `foto_pengguna`, `alamat_id`, `alamat_pengguna`, `rt`, `rw`, `kode_pos`, `no_telp`, `token`, `desa_id`) VALUES
(1, 'user_warga1', 'pass_warga1', 'warga', 1, '2023-11-15 15:06:41', 'admin', '2023-11-15 16:48:23', 'admin', '1990-01-01', 'Bahomoleo', 1, 'Laki-laki', NULL, NULL, 'PNS', 1234567890, 12345, 30, NULL, 'Islam', 1, 'Bahomoleo', 'John Doe', 'foto.jpg', 1, 'Jl. Raya Desa A', '001', '001', '12345', '081234567890', 'token1', 1),
(2, 'user_warga2', 'pass_warga2', 'warga', 1, '2023-11-15 15:06:41', 'admin', '2023-11-15 16:48:23', 'admin', '1995-02-02', 'Bahomoleo', 2, 'Perempuan', NULL, NULL, 'Guru', 2147483647, 23456, 27, NULL, 'Kristen', 2, 'Bahomoleo', 'Jane Smith', 'foto2.jpg', 2, 'Jl. Raya Desa B', '002', '002', '23456', '082345678901', 'token2', 1),
(3, 'user_admin1', 'pass_admin1', 'admin desa', 2, '2023-11-15 15:06:41', 'admin', '2023-11-15 16:48:23', 'admin', '1985-05-05', 'Bahomoleo', 1, 'Laki-laki', NULL, NULL, 'Pegawai', 2147483647, 34567, 36, NULL, 'Katolik', 3, 'Bahomoleo', 'Michael Brown', 'foto3.jpg', 3, 'Jl. Raya Desa C', '003', '003', '34567', '083456789012', 'token3', 1),
(4, 'user_admin2', 'pass_admin2', 'admin desa', 2, '2023-11-15 15:06:41', 'admin', '2023-11-15 16:48:23', 'admin', '1988-08-08', 'Bahomoleo', 2, 'Perempuan', NULL, NULL, 'Dosen', 2147483647, 45678, 33, NULL, 'Buddha', 4, 'Bahomoleo', 'Sarah Johnson', 'foto4.jpg', 4, 'Jl. Raya Desa D', '004', '004', '45678', '084567890123', 'token4', 1),
(5, 'user_admin3', 'pass_admin3', 'admin desa', 2, '2023-11-15 15:06:41', 'admin', '2023-11-15 16:48:23', 'admin', '1980-10-10', 'Bahomoleo', 1, 'Laki-laki', NULL, NULL, 'Wiraswasta', 2147483647, 56789, 41, NULL, 'Hindu', 5, 'Bahomoleo', 'David Lee', 'foto5.jpg', 5, 'Jl. Raya Desa E', '005', '005', '56789', '085678901234', 'token5', 1);

-- --------------------------------------------------------

--
-- Table structure for table `role`
--

CREATE TABLE `role` (
  `id_role` int(11) NOT NULL,
  `nama_role` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `role`
--

INSERT INTO `role` (`id_role`, `nama_role`) VALUES
(1, 'warga'),
(2, 'admin desa');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `desa`
--
ALTER TABLE `desa`
  ADD PRIMARY KEY (`desa_id`);

--
-- Indexes for table `pengguna`
--
ALTER TABLE `pengguna`
  ADD PRIMARY KEY (`id_pengguna`),
  ADD UNIQUE KEY `username` (`username`),
  ADD KEY `fk_pengguna_role` (`role_id`),
  ADD KEY `desa_id` (`desa_id`);

--
-- Indexes for table `role`
--
ALTER TABLE `role`
  ADD PRIMARY KEY (`id_role`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `desa`
--
ALTER TABLE `desa`
  MODIFY `desa_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `pengguna`
--
ALTER TABLE `pengguna`
  MODIFY `id_pengguna` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `role`
--
ALTER TABLE `role`
  MODIFY `id_role` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `pengguna`
--
ALTER TABLE `pengguna`
  ADD CONSTRAINT `fk_pengguna_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id_role`),
  ADD CONSTRAINT `pengguna_ibfk_1` FOREIGN KEY (`desa_id`) REFERENCES `desa` (`desa_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
