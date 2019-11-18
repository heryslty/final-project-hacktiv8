-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 18, 2019 at 09:39 AM
-- Server version: 10.4.8-MariaDB
-- PHP Version: 7.1.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gamedb`
--

-- --------------------------------------------------------

--
-- Table structure for table `article`
--

CREATE TABLE `article` (
  `id` int(11) NOT NULL,
  `title` varchar(100) NOT NULL,
  `content` text NOT NULL,
  `thumbnail` text NOT NULL,
  `image` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_published` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `article`
--

INSERT INTO `article` (`id`, `title`, `content`, `thumbnail`, `image`, `created_at`, `updated_at`, `deleted_at`, `is_published`) VALUES
(3, 'Super Mario Oddysey', '<p>Super Mario Odyssey is a platform game published by Nintendo for the Nintendo Switch on October 27, 2017. An entry in the Super Mario series, it follows Mario and Cappy, a sentient hat that allows Mario to control other characters and objects, as they journey across various worlds to save Princess Peach from his nemesis Bowser, who plans to forcibly marry her. In contrast to the linear gameplay of prior entries, the game returns to the primarily open-ended, exploration-based gameplay featured in Super Mario 64 and Super Mario Sunshine.</p>\n\n<p>The game was developed by Nintendo\'s Entertainment Planning &amp; Development division, and began development soon after the release of Super Mario 3D World in 2013. Various ideas were suggested during development, and to incorporate them all, the team decided to employ a sandbox-style of gameplay. Unlike previous installments such as New Super Mario Bros. and Super Mario 3D World, which were aimed at a casual audience, the team designed Super Mario Odyssey to appeal to the series\' core fans. The game also features a vocal theme song, \"Jump Up, Super Star!\", a first for the series.</p>\n\n<p>Super Mario Odyssey received universal acclaim from critics who called it one of the best games in the series, with particular praise towards its inventiveness and originality. It also won numerous awards, including for game of the year. The game was also a commercial success, selling over 15.38 million copies by September 2019, which makes it one of the best-selling Switch games.</p>', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/icons/4.jpeg', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/banners/4.jpeg', '2019-11-10 22:39:57', '2019-11-17 12:22:06', NULL, 1),
(4, 'The Legend of Zelda: Majora\'s Mask 3D', '<p>The Legend of Zelda: Majora\'s Mask 3D is an action-adventure game developed by Grezzo and published by Nintendo for the Nintendo 3DS handheld game console. The game is an enhanced remake of The Legend of Zelda: Majora\'s Mask, which was originally released for the Nintendo 64 home console in 2000. The game was released worldwide in February 2015, coinciding with the North American and European releases of the New Nintendo 3DS.</p>\n\n<p>Much like its predecessor, The Legend of Zelda: Ocarina of Time 3D, Majora\'s Mask 3D is a remaster of the original Majora\'s Mask game, featuring enhanced stereoscopic 3D graphics, touchscreen controls, and gyroscopic features. Like the original, the game follows Link. Link is given only three days to save the land of Termina from being crushed by its moon, using various abilities obtained by wearing different masks. In addition, Nintendo changed the time mechanic of the original game, giving the player more time to explore, and added fishing.</p>', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/icons/5.jpeg', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/banners/5.jpeg', '2019-11-10 22:41:42', '2019-11-17 12:22:18', NULL, 1),
(5, 'The Legend of Zelda: The Wind Waker', '<p>The Legend of Zelda: The Wind Waker is a 2002 action-adventure game developed and published by Nintendo for the GameCube. The tenth installment in The Legend of Zelda series, The Wind Waker is set on a group of islands in a vast sea, a departure for the series. It follows series protagonist Link as he attempts to save his sister from the sorcerer Ganon and becomes embroiled in a struggle for the Triforce, a sacred wish-granting relic. Aided by allies including pirate captain Tetra—an incarnation of Princess Zelda—and a talking boat named the King of Red Lions, Link sails the ocean, explores islands, and traverses dungeons to acquire the power necessary to defeat Ganon.</p>\n\n<p>Gameplay is presented from the third-person perspective. The player controls Link, who fights with a sword and shield, in addition to other items. Like previous Zelda games, Link explores dungeons to fight enemies, find items, and solve puzzles. The player must also explore the oceans, and wind, which facilitates sailing, plays a prominent role and can be controlled with a magic conductor\'s baton called the Wind Waker. While the game retains the basic 3D gameplay of its predecessors, Ocarina of Time (1998) and Majora\'s Mask (2000), The Wind Waker features a distinctive cartoon-like art style created through cel shading.</p>\n\n<p>Development began before Majora\'s Mask was completed in 2000 and lasted until late 2002. Eiji Aonuma reprised his directorial duties for The Wind Waker, while Shigeru Miyamoto and Takashi Tezuka produced. Nintendo\'s Zelda team did not want to continue using the realistic graphics of previous Zelda games, instead choosing a cartoonish aesthetic that would offer new gameplay and combat possibilities. Kenta Nagata, Hajime Wakai, Toru Minegishi, and Koji Kondo composed the soundtrack, which consists of original tracks inspired by traditional Irish music and reworked pieces from older Zelda games. Nintendo released The Wind Waker in Japan in December 2002, in North America in March 2003, and in Europe in May 2003.</p>\n\n<p>Although The Wind Waker received critical acclaim—with praise directed towards its visuals, gameplay, design, and story—and won several Game of the Year accolades, its art direction proved divisive among players. This contributed to comparatively weak sales; the game sold 4.6 million copies, far below the 7.6 million Ocarina of Time sold. As a result, Nintendo changed directions with the next Zelda installment, the more realistically styled Twilight Princess (2006). However, The Wind Waker\'s reputation improved over time, and it is now widely considered one of the greatest video games ever made. The Wind Waker originated the \"Toon Link\" character and received two sequels for the Nintendo DS, Phantom Hourglass (2007), and Spirit Tracks (2009). A high-definition remaster, The Legend of Zelda: The Wind Waker HD, was released for the Wii U in 2013.</p>', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/icons/19.jpeg', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/banners/19.jpeg', '2019-11-11 06:31:15', '2019-11-17 12:23:50', NULL, 1),
(6, 'The Legend of Zelda: A Link Between Worlds', '<p>The Legend of Zelda: A Link Between Worlds is an action-adventure game developed and published by Nintendo for the Nintendo 3DS. It is the 17th installment in The Legend of Zelda series and the sequel to the 1991 Super Nintendo Entertainment System game A Link to the Past. Announced in April 2013, A Link Between Worlds was released in Australia, Europe, and North America in November 2013. A month later, it was released in Japan as The Legend of Zelda: Triforce of the Gods 2[a] (Triforce of the Gods being the Japanese title of A Link to the Past).</p>\r\n\r\n<p>Players play as a young adventurer, Link, who is tasked with restoring peace to the kingdom of Hyrule after the evil sorcerer Yuga captures Princess Zelda and escapes through a rift to the ruined world of Lorule. Yuga seeks to kidnap the Seven Sages and use their power to resurrect the Demon King Ganon. Link is granted the ability to merge onto walls as a painting after obtaining a magic bracelet and encountering Yuga, which allows Link to reach previously inaccessible areas and travel between Hyrule and Lorule.</p>\r\n\r\n<p>Concept development began with a small team in 2009. During this phase, the game mechanic of Link merging onto walls was prototyped. However, development suffered several setbacks and ceased entirely in late 2010 as core team members were reassigned to different projects. A year later, development restarted and, after several failed pitches to series creator, Shigeru Miyamoto, the game entered full production in 2012. Changing the established conventions of the series became a goal for the game\'s designers; this led to a change in the game\'s structure, allowing players to clear the majority of dungeons in any order they choose, and the introduction of the item rental system.</p>\r\n\r\n<p>A Link Between Worlds received critical acclaim and sold over 2.5 million copies worldwide within five months. The audio, dungeon and puzzle design, open structure, and level of difficulty were praised by critics. The introduction of features such as the wall-merging mechanics and item rental system were well received, with reviewers complimenting how well they integrated with the existing gameplay formula. The game also received multiple awards and nominations.</p>', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/icons/31.jpeg', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/banners/31.jpeg', '2019-11-11 06:31:57', NULL, NULL, 1),
(7, 'The Legend of Zelda: Breath of the Wild', '<p>The Legend of Zelda: Breath of the Wild is an action-adventure game developed and published by Nintendo, released for the Nintendo Switch and Wii U consoles on March 3, 2017. Breath of the Wild is set at the end of the Zelda timeline; the player controls Link, who awakens from a hundred-year slumber to defeat Calamity Ganon before it can destroy the kingdom of Hyrule.</p>\r\n\r\n<p>Similarly to the original Legend of Zelda (1986), players are given little instruction and can explore the open world freely. Tasks include collecting multipurpose items to aid in objectives or solving puzzles and side quests for rewards. The world is unstructured and designed to reward experimentation, and the story can be completed in a nonlinear fashion.</p>\r\n\r\n<p>Development of Breath of the Wild lasted five years. Wanting to reinvent the series, Nintendo introduced elements such as a detailed physics engine, high-definition visuals, and voice acting. Monolith Soft, known for their work in the open-world Xenoblade Chronicles series, assisted in designing landscapes and topography. The game was planned for release in 2015 as a Wii U exclusive, but was delayed twice due to problems with the physics engine. Breath of the Wild was a launch game for the Switch and the final game published by Nintendo for the Wii U. Two downloadable content packs were released later in 2017.</p>\r\n\r\n<p>Breath of the Wild received acclaim for its open-ended gameplay and attention to detail, and has been called one of the greatest video games of all time. Critics called it a landmark in open-world design, despite minor criticism for its technical performance at launch. It won numerous awards, including several game of the year honors. By 2019, Breath of the Wild had sold over 16.04 million copies worldwide, making it the bestselling Zelda game. A sequel was announced at E3 2019.</p>', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/icons/3.jpeg', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/banners/3.jpeg', '2019-11-11 06:32:44', '2019-11-18 06:06:42', NULL, 1),
(8, 'Granblue Fantasy', '<p>Granblue Fantasy is a Japanese media franchise and a role-playing video game developed by Cygames for Android, iOS and web browsers, which first released in Japan in March 2014. The game is notable for reuniting music composer Nobuo Uematsu and art director Hideo Minaba, who previously collaborated on Final Fantasy V (1992), Final Fantasy VI (1994), Final Fantasy IX (2000), and Lost Odyssey (2007).</p>\r\n\r\n<p>The game plays as a role-playing video game with turn-based battles.[2] The game also contains summons and a class system that alters the main character\'s move-set and growth.[3] Characters gain levels and abilities by accruing experience; summons and weapons equipped also confer characters with bonuses on attack power and HP. The characters themselves are gained either via quests (the main story quests or special event quests) or by using in-game currency to receive random crystal fragments, which may contain special weapons that add specific characters to the party. Characters, summons, and weapons are ranked (from best to worst) as SSR, SR, R, or N; each is also of type wind, water, fire, earth, light, or darkness. Voice actors provide voices for all of the characters in battle, and for much of the main and event storylines. </p>', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/icons/256.jpeg', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/banners/256.jpeg', '2019-11-11 06:36:35', '2019-11-17 13:12:21', NULL, 1),
(9, 'Super Mario 64', '<p>Super Mario 64 is a 1996 platform video game for the Nintendo 64 and the first in the Super Mario series to feature three-dimensional (3D) gameplay. As Mario, the player explores Princess Peach\'s castle and must rescue her from Bowser. Super Mario 64 features open-world playability, degrees of freedom through all three axes in space, and relatively large areas which are composed primarily of true 3D polygons as opposed to only two-dimensional (2D) sprites. It emphasizes exploration within vast worlds, which require the player to complete various missions in addition to the occasional linear obstacle courses (as in traditional platform games). It preserves many gameplay elements and characters of earlier Mario games as well as the visual style.<br><br>Producer/director and Mario creator Shigeru Miyamoto conceived a 3D Mario game during the production of Star Fox (1993). Super Mario 64\'s development, handled by Nintendo EAD, lasted approximately three years; one was spent on designing while the next two on direct work. The visuals were created using the Nichimen N-World toolkit and Miyamoto aimed to include more details than earlier games. The score was composed by Koji Kondo. A multiplayer mode featuring Luigi as a playable character was planned but cut. Along with Pilotwings 64, Super Mario 64 was one of the launch games for Nintendo 64. Nintendo released it in Japan on June 23, 1996, and later in North America, Europe, and Australia. A remake, Super Mario 64 DS, was released for the Nintendo DS in 2004, and the original version was re-released for Nintendo\'s Virtual Console service on the Wii and Wii U in 2006 and 2015, respectively.<br><br>Super Mario 64 is acclaimed as one of the greatest video games of all time, and was the first game to receive a perfect score from Edge magazine. Reviewers praised its ambition, visuals, gameplay, and music, although they criticized its unreliable camera system. It is the Nintendo 64\'s bestseller, with more than eleven million copies sold by 2003. The game left a lasting impression on the field of 3D game design, featuring a dynamic camera system and 360-degree analog control, and established a new archetype for the 3D genre, much as Super Mario Bros. did for 2D side-scrolling platformers. Numerous developers cited Super Mario 64 as an influence on their later games.<br></p>', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/icons/171.jpeg', 'https://mythril.nyc3.cdn.digitaloceanspaces.com/games/banners/171.jpeg', '2019-11-17 13:12:13', '2019-11-18 06:06:40', NULL, 1);

-- --------------------------------------------------------

--
-- Table structure for table `feedback`
--

CREATE TABLE `feedback` (
  `id` int(11) NOT NULL,
  `sender` varchar(100) NOT NULL,
  `subject` varchar(100) NOT NULL,
  `text` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `feedback`
--

INSERT INTO `feedback` (`id`, `sender`, `subject`, `text`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'djodi.ramadhan@gmail.com', 'Webnya bagus', 'webnya bagus gan', '2019-11-17 08:40:16', NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `privilege` enum('admin','user') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `name`, `email`, `password`, `privilege`) VALUES
(1, 'Administrator', 'djodi.ramadhan@gmail.com', '$2a$10$J0NSfGIrCJ57M.szRWOut.IZSg0QZ6UcI/D.HSZ/zmDkyPN1Ez86W', 'admin');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `article`
--
ALTER TABLE `article`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `feedback`
--
ALTER TABLE `feedback`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `article`
--
ALTER TABLE `article`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `feedback`
--
ALTER TABLE `feedback`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
