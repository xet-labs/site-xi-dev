/*M!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19  Distrib 10.11.13-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: XI
-- ------------------------------------------------------
-- Server version	10.11.13-MariaDB-0ubuntu0.24.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `blog_media`
--

DROP TABLE IF EXISTS `blog_media`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `blog_media` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `blog_id` bigint(20) unsigned NOT NULL,
  `media_id` bigint(20) unsigned NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `blog_id` (`blog_id`),
  KEY `media_id` (`media_id`),
  CONSTRAINT `blog_media_ibfk_1` FOREIGN KEY (`blog_id`) REFERENCES `blogs` (`id`) ON DELETE CASCADE,
  CONSTRAINT `blog_media_ibfk_2` FOREIGN KEY (`media_id`) REFERENCES `media` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_media`
--

LOCK TABLES `blog_media` WRITE;
/*!40000 ALTER TABLE `blog_media` DISABLE KEYS */;
/*!40000 ALTER TABLE `blog_media` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blogs`
--

DROP TABLE IF EXISTS `blogs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `blogs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) unsigned NOT NULL,
  `status` enum('draft','published','published_hidden','archived') NOT NULL DEFAULT 'draft',
  `tags` varchar(255) DEFAULT NULL,
  `title` varchar(255) NOT NULL,
  `short_title` varchar(255) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `featured_img` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `path` varchar(255) DEFAULT NULL,
  `created_at` timestamp(3) NULL DEFAULT current_timestamp(3),
  `updated_at` timestamp(3) NULL DEFAULT current_timestamp(3) ON UPDATE current_timestamp(3),
  `content` longtext DEFAULT NULL,
  `meta` longtext DEFAULT NULL,
  `meta_keywords` text DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`),
  KEY `idx_slug` (`slug`),
  KEY `idx_created_at` (`created_at`),
  FULLTEXT KEY `ft_content` (`content`,`title`,`description`),
  CONSTRAINT `blogs_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blogs`
--

LOCK TABLES `blogs` WRITE;
/*!40000 ALTER TABLE `blogs` DISABLE KEYS */;
INSERT INTO `blogs` VALUES
(1,4,'published','[\"SEO\", \"Marketing Insight\"]','What is SEO? A beginner’s guide to search engine optimization','What is SEO ?','SEO is the strategic art of enhancing a website\'s visibility on search engines like Google, Bing, and Yahoo. It involves optimizing various elements to rank higher in search results...','/media/4/img/what-is-seo.webp','what-is-seo','','2024-05-15 06:35:14.000','2025-08-17 07:37:17.588','<p>\n    In the vast landscape of the internet, standing out is no small feat. Search Engine Optimization (SEO) emerges as the beacon that guides websites through the digital fog, ensuring they are not lost in obscurity. In this blog, we unravel the mysteries of SEO, offering insights and best practices to elevate your website\'s visibility.\n</p>\n<h2>\n    What is SEO?\n</h2>\n<p>\n    <b>SEO is the strategic art of enhancing a website\'s visibility on search engines like Google, Bing, and Yahoo.</b> It involves optimizing various elements to rank higher in search results, driving organic traffic and, ultimately, increasing the chances of converting visitors into customers. \n</p>\n<h2>\n    The Pillars of SEO\n</h2>\n<!-- <h3>\n    The Pillars of SEO\n</h3>\n<h4>\n    The Pillars of SEO\n</h4> -->\n<p>\n    Once you have the foundations of your site ready, you’ll need to pay attention to many small details such as metadata and linking, which can help improve your rankings. This article will cover what it takes to implement those details and make sure that they are meeting SEO standards and eventually SEO ROI.\n</p>\n\n<ol>\n    <li>\n    Share on social media\n    </li>\n    <li>\n    Create a blog newsletter\n    </li>\n    <li>\n    Write for other sites\n    </li>\n    <li>\n    Reach out to an existing community\n    </li>\n    <li>\n    Participate in question and discussion sites\n    </li>\n    <li>\n    Invest in paid ads\n    </li>\n    <li>\n    Try new content formats\n    </li>\n</ol>\n\n<p>\n    At this stage, you have everything you need to start a blog. These last couple of steps will focus on how to spread the word about your blog it into a serious monetization tool.\n</p>\n\n<ul>\n    <li>\n        <b><b>Share on social media:</b></b> Social media is an excellent place to post your content and draw attention to your blog. Whether you promote your blog on Facebook, Instagram, Twitter or LinkedIn, it’s a great way to reach new readers. Learn more about blogging vs instagram in our guide. \n    </li>\n    <li>\n        <b><b>Create a blog newsletter:</b></b> Send out a weekly email newsletter to engage your readers and get them coming back to your blog for more. This will help you sustain a loyal fan base. To get subscribers to your blog email list in the first place, include a prominent Subscribe button in your website’s navigation bar, footer, and within your blog posts.\n    </li>\n</ul>\n\n<div class=\"blog-yt\">\n    <iframe src=\"https://www.youtube-nocookie.com/embed/C-UeBdtopdA?si=VhkHRKvqISKOVGvT?&rel=0\" title=\"XetIndustries YouTube video\"  frameborder=\"0\" allow=\"accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture;\" referrerpolicy=\"strict-origin-when-cross-origin\" allowfullscreen></iframe>\n</div>\n\n<p>\n    In order to get readers, you’ll need to find creative ways to drive traffic to your site. While improving your SEO is an important step, the following methods can also help you promote your blog. Note that most of them are completely free, while a few are paid.\n</p>\n\n<p>\n    If you’re wondering how to create a blog, you’ve come to the right place. As a blogger myself, I can tell you it’s a rewarding way to hone your writing skills, explore new ideas and build an online presence that revolves around your passions and expertise. You’ll get the chance to inspire, educate, and entertain your readers—and as your blog grows, you can even start making money and turn it into a full-time job or use it to start a business.\n</p>\n',NULL,'what is SEO, Search Engine Optimization, what is website optimization, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(2,2,'published','[\"Passive Income\", \"Internet Monetization\"]','Passive income by sharing internet',NULL,'Transform your internet connection into a Passive Income Stream by sharing your internet','/media/2/img/passive-income-by-sharing-internet.webp','passive-income-by-sharing-internet','','2024-05-16 06:36:14.000','2025-08-17 07:37:17.588','<p>\n    The Internet has become an inseparable part of our lives that powers virtually every aspect of modern economy, what if <b>you could transform your Internet connection into a passive income stream</b> by merely sharing your Internet, Intrigued?\n</p>\n<p>\n    Here’s a list of applications sorted by their earning potential (high to low).\n</p>\n\n<h2>\n    <a href=\"https://mystnodes.co/?referral_code=orLHWK7YyqeLyf8G9BL1e7s3ox53YTw6iY8zDGLC\" target=\"_blank\" class=\"inc-content-logo\">\n    Mysterium Node\n    <div class=\"content-logo-wrap\">\n        <div class=\"content-logo\">\n        <img src=\"/media/2/img/myst-favicon-64.png\" alt=\"\">\n        </div>\n    </div>\n    </a>\n</h2>\n<p>\n    Mysterium Network is a decentralized network of computers that provides P2P services. It allows the Node runners to share internet bandwidth and earn “MYST” crypto while simultaneously contributing the mysterium network to grow. \n</p>\n<p>\n    <b>Rate: </b>Varies with services being utilised.\n    <br>\n    <b>Platform: </b>Android, iOS, Windows, Linux, Mac.\n    <br>\n    <b>Payout: </b>Crypto (Myst).\n    <br>\n    <b>Requisite: </b> Internet connection (prefer ethernet).\n</p>\n<p><b>Setup:</b> To start earning, set up MystNodes on Android, iOS, Windows, Linux or Mac.</p>\n<ol class=\"\">\n    <li>\n    <p><a href=\"https://mystnodes.co/?referral_code=orLHWK7YyqeLyf8G9BL1e7s3ox53YTw6iY8zDGLC\" target=\"_blank\"  class=\"th-1\">Sign up for MystNodes</a> and confirm your email.</p>\n    </li>\n\n    <li>\n    <p>On the <a href=\"https://my.mystnodes.com/onboarding\" target=\"_blank\">MystNodes onboarding page</a>, select your platform and continue with node installation.</p>\n    </li>\n    \n    <li>\n    <p>After installation open your NodeUi <a href=\"http://localhost:4449\" target=\"_blank\">http://localhost:4449</a>.</p>\n    </li>\n\n    <li>\n    <p>Complete the Quick Onboarding process by entering your Myst token’s wallet address (<a href=\"https://help.mystnodes.com/en/articles/8004186-adding-myst-token-to-metamask-on-the-polygon-mainnet\" target=\"_blank\">from Metamask</a> or any other crypto wallet).</p>\n    </li>\n    <li>\n    <p>Node setup is now complete, and its stats can be viewed on the <a href=\"https://my.mystnodes.com/nodes\" target=\"_blank\">MystNodes dashboard</a>.</p>\n    </li>\n</ol>\n\n\n<h3>Pro tip</h3>\n<p>To maximize your chances of earning:</p>\n<ul>\n    <li>\n    <b>Prefer ethernet over WiFi</b>.\n    </li>\n    <li>\n    Try <b>enabling additional services</b> on your Node Dashboard.\n    </li>\n    <li>\n    <b>Uptime matters</b>, the moment you are disconnected from a client, they will be reconnected to a more competitive node.\n    </li>\n    <li>\n    Your internet speed & <b>Geographical location matter</b>.\n    </li>\n    <li>\n    <b>Do not use VPNs</b> based on higher earning countries; your node would display as a data center node and wouldn\'t be able to get past geo-restrictions.\n    </li>\n</ul>\n\n\n<!-- Earnapp -->\n<h2>\n    <a href=\"https://earnapp.com/i/r5zpnVZD\" target=\"_blank\" class=\"inc-content-logo\">\n    Earnapp\n    <div class=\"content-logo-wrap\">\n        <div class=\"content-logo\">\n        <img src=\"/media/2/img/earnapp-favicon-64.png\" alt=\"\">\n        </div>\n    </div>\n    </a>\n</h2>\n<p>\n    Earnapp makes you money by sharing your internet bandwidth with businesses in need of network resources. Your device works as a proxy service enabling clients to route traffic through your device.\n</p>\n<p>\n    <b>Rate: </b>$1 / 10GB\n    <br>\n    <b>Platform: </b>Android, iOS, Windows, Linux, Mac.\n    <br>\n    <b>Payout: </b>PayPal, Amazon Gift Card, Wise.\n    <br>\n    <b>Min payout: </b>$2.5  (auto payment supported).\n    <br>\n    <b>Requisite: </b> Internet connection.\n</p>\n<p><b>Setup:</b> <a href=\"https://earnapp.com/i/r5zpnVZD\" target=\"_blank\" class=\"th-1\">Sign up for Earnapp</a> and install the <a href=\"https://earnapp.com/dashboard\" target=\"_blank\">Earnapp app</a>  for your device.</p>\n\n<!-- PacketStream -->\n<h2>\n    <a href=\"https://packetstream.io/?psr=4SxY\" target=\"_blank\" class=\"inc-content-logo\">\n    PacketStream\n    <div class=\"content-logo-wrap\">\n        <div class=\"content-logo\">\n        <img src=\"/media/2/img/packetstream-favicon-64.png\" alt=\"\">\n        </div>\n    </div>\n    </a>\n</h2>\n<p>PacketStream is similar proxy service app that lets you earn by allowing clients to route internet traffic through your device and earn based on the amount of data transferred.</p>\n<p>\n    <b>Rate: </b>$1 / 10GB\n    <br>\n    <b>Platform: </b>Windows, Linux, Mac.\n    <br>\n    <b>Payout: </b>PayPal.\n    <br>\n    <b>Min payout: </b>$5\n    <br>\n    <b>Requisite: </b> Internet connection.\n</p>\n<p><b>Setup:</b> <a href=\"https://packetstream.io/?psr=4SxY\" target=\"_blank\" class=\"th-1\">Sign up for PacketStream</a> and install the app for your device from the “Download” section.</p>\n\n<!-- Honeygain -->\n<h2>\n    <a href=\"https://r.honeygain.me/RISHI0A188\" target=\"_blank\" class=\"inc-content-logo\">\n    Honeygain\n    <div class=\"content-logo-wrap\">\n        <div class=\"content-logo\">\n        <img src=\"/media/2/img/honeygain-favicon-64.png\" alt=\"\">\n        </div>\n    </div>\n    </a>\n</h2>\n<p>Honeygain has been a long-standing platform that enables users to monetize their unused internet bandwidth, users contribute to tasks like content delivery and market research, facilitating peer-to-peer (P2P) traffic through their devices.</p>\n<p>You also get a <b>$3 initial sign-up bonus</b>, link below.</p>\n<p>\n    <b>Rate: </b>$1 / 10GB\n    <br>\n    <b>Platform: </b>Android, Windows, Linux, Mac.\n    <br>\n    <b>Payout: </b>PayPal, Crypto (Jmpt).\n    <br>\n    <b>Min payout: </b> $20 (PayPal), Crypto (no min limit)\n    <br>\n    <b>Requisite: </b> Internet connection.\n</p>\n<p>\n    <b>Setup:</b> <a href=\"https://r.honeygain.me/RISHI0A188\" target=\"_blank\" class=\"th-1\">Sign up for Honeygain</a> and install the app for your device. \n    <br>\n    <b>App:</b>\n    <a href=\"https://download.honeygain.com/android-app/honeygain_app.apk\" target=\"_blank\">Android</a> | \n    <a href=\"https://download.honeygain.com/windows-app/Honeygain_install.exe\" target=\"_blank\">Windows</a> | \n    <a href=\"https://hub.docker.com/r/honeygain/honeygain\" target=\"_blank\">Linux</a> | \n    <a href=\"https://hub.docker.com/r/honeygain/honeygain\" target=\"_blank\">Mac</a>\n</p>\n\n\n<blockquote>\n    <p>\n    <b>For users in India:</b> <br>\n    Other apps with similar payout rates were discovered that seem to be working but may take a considerable time to accumulate substantial earnings. To reach cashout goals efficiently, <b>prioritize apps with proven bandwidth efficiency, like those listed above</b>.\n    <br>\n    For users in other regions, trying all apps and identifying the most rewarding one is recommended.\n    </p>\n</blockquote>\n\n\n<!-- Repocket -->\n<h2>\n    <a href=\"https://link.repocket.com/Txed\" target=\"_blank\" class=\"inc-content-logo\">\n    Repocket\n    <div class=\"content-logo-wrap\">\n        <div class=\"content-logo\">\n        <img src=\"/media/2/img/repocket-favicon-64.png\" alt=\"\">\n        </div>\n    </div>\n    </a>\n</h2>\n<p>Repocket offers a solution to unused internet bandwidth going to waste. Simply install the app and share your spare bandwidth to earn rewards. It is available for most devices and offers a <b>$5 sign-up bonus through the link below</b>.</p>\n<p>\n    <b>Rate: </b>$1 / 10GB\n    <br>\n    <b>Platform: </b>Android, iOS, Windows, Linux, Mac.\n    <br>\n    <b>Payout: </b>PayPal\n    <br>\n    <b>Min payout: </b>$20\n    <br>\n    <b>Requisite: </b> Internet connection.\n</p>\n<p><b>Setup:</b> <a href=\"https://link.repocket.com/Txed\" target=\"_blank\" class=\"th-1\">Sign up for Repocket</a> and install the app for your device.</p>\n\n\n<!-- Pawnsapp -->\n<h2>\n    <a href=\"https://pawns.app/?r=1012365\" target=\"_blank\" class=\"inc-content-logo\">\n    Pawnsapp\n    <div class=\"content-logo-wrap\">\n        <div class=\"content-logo\">\n        <img src=\"/media/2/img/pawnsapp-favicon-64.png\" alt=\"\">\n        </div>\n    </div>\n    </a>\n</h2>\n<p>Pawnsapp is a reliable option for earning passive income. It boasts competitive payout rates and offers multiple options for cashing out your earnings. <b>$3 sign-up bonus through the link below.</b></p>\n<p>\n    <b>Rate: </b>$2 / 10GB\n    <br>\n    <b>Platform: </b>Android, Windows, Linux, Mac.\n    <br>\n    <b>Payout: </b>PayPal, VISA, Crypto (BTC)\n    <br>\n    <b>Min payout: </b>$5\n    <br>\n    <b>Requisite: </b> Internet connection.\n</p>\n<p><b>Setup:</b> <a href=\"https://pawns.app/?r=1012365\" target=\"_blank\" class=\"th-1\">Sign up for Pawnsapp</a> and install the app for your device.</p>\n\n\n<!-- Traffmonetizer -->\n<h2>\n    <a href=\"https://app.traffmonetizer.com/sign-up\" target=\"_blank\" class=\"inc-content-logo\">\n    Traffmonetizer\n    <div class=\"content-logo-wrap\">\n        <div class=\"content-logo\">\n        <img src=\"/media/2/img/traffmonetizer-favicon-64.png\" alt=\"\">\n        </div>\n    </div>\n    </a>\n</h2>\n<p>TraffMonitizer offers a seamless way to earn passive income by selling your unused internet traffic to marketing and advertising agencies. <br><b>However, during my testing, I encountered difficulties with the sign-up process !!</b></p>\n<p>\n    <b>Rate: </b>$1 / 10GB\n    <br>\n    <b>Platform: </b>Android, Windows, Linux, Mac.\n    <br>\n    <b>Payout: </b>Payoneer, BTC, Webmoney, Skrill, Payeer, USDT (TRC20)\n    <br>\n    <b>Min payout: </b>$5\n    <br>\n    <b>Requisite: </b> Internet connection.\n</p>\n<p>\n    <b>Setup:</b> <a href=\"https://app.traffmonetizer.com/sign-up\" target=\"_blank\" class=\"th-1\">Sign up for Traffmonetizer</a> and install the app for your device.\n</p>\n',NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(3,1,'published','[\"mysteriumnode\", \"mysteriumvpn\"]','Mysterium node backup & restore on linux or windows',NULL,'Backup & restore your mysterium node on both windows and linux platforms','/media/1/img/mysterium-node-backup-and-restore.webp','mysterium-node-backup-and-restore','','2024-05-01 06:36:14.000','2025-08-17 07:37:17.588','<p>Backing up your Mysterium node is crucial to ensure that your configs and data are safe in case of system failures or data corruption. This guide will walk you through the steps to back up your Mysterium node on both Windows and Linux platforms.</p>\n\n<h2> Myst node config & data </h2>\n<p> Mysterium node\'s config includes active services, account API, current theme etc, while data includes configs, node’s public address, private key and its encryption details, Message Authentication Code, file id, version etc.</p>\n<p>\n  <b>Config:</b> <code>/etc/mysterium-node/</code>\n  <br>\n  <b>Data:</b> <code>/var/lib/mysterium-node/</code>\n</p>\n\n\n<h2> Backup Myst node on Linux </h2>\n<p>Create a compressed backup with necessary files from <code>/var/lib/mysterium-node/*</code>.</p>\n\n<pre>\n<code class=\"language-bash\"># creates &ltmyst-date.tar.gz&gt compressed backup \nsudo tar --exclude={\"mainnet/*\",\"*/logs/*\"} -cvzf ~/myst-$(date +\"%Y%m%d%H%M%S\").tar.gz -C /var/lib/mysterium-node/</code>\n</pre>\n\n<p>The <code>--exclude={\"mainnet/*\",\"*/logs/*\"}</code> excludes unnecessary files.</p>\n\n\n<h2> Restore Myst node on Linux </h2>\n<p>Replace <code>~/myst-date.tar.gz</code> with actual zip backup of MystNode.</p>\n<pre>\n<code class=\"language-bash\"># extracts &ltmyst-date.tar.gz&gt to \'/var/lib/mysterium-node/\'\nsudo tar -xvzf ~/myst-date.tar.gz -C /var/lib/mysterium-node/ --strip-components=2</code>\n</pre>',NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(4,1,'published_hidden',NULL,'Does my business need a website? 9 reasons why it does',NULL,'The Internet powers virtually every aspect of the modern economy, transform your Internet connection into a passive income stream by merely sharing your Internet connection','https://static.wixstatic.com/media/1f6616_88503e0f296f49afa70a53e3812fd92e~mv2.png/v1/fill/w_740,h_489,al_c,q_90,usm_0.66_1.00_0.01,enc_avif,quality_auto/1f6616_88503e0f296f49afa70a53e3812fd92e~mv2.png','passive-income-by-sharing-internet','','2024-05-17 06:36:14.000','2025-08-17 07:37:17.588',NULL,NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(5,1,'published_hidden','[\"Mysterium Node\"]','	What is a Domain name and why it matters',NULL,'Backup & restore your mysterium node on windows and linux','https://static.wixstatic.com/media/84b06e_74efe24ffe3f4cccbfb4d23e904498eb~mv2.png/v1/fill/w_740,h_423,al_c,q_85,usm_0.66_1.00_0.01,enc_avif,quality_auto/84b06e_74efe24ffe3f4cccbfb4d23e904498eb~mv2.png','mysterium-node-backup-and-restore','','2024-05-01 01:06:14.000','2025-08-17 07:37:17.588',NULL,NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(6,2,'published_hidden','[\"Passive Income\", \"Internet Monetization\"]','10 outstanging site examples',NULL,'Transform your internet connection into a Passive Income Stream by sharing your internet','https://static.wixstatic.com/media/46e2e0_0b737cde10004f8c8a4deade3e45e47e~mv2.png/v1/fill/w_740,h_489,al_c,q_90,usm_0.66_1.00_0.01,enc_avif,quality_auto/46e2e0_0b737cde10004f8c8a4deade3e45e47e~mv2.png','passive-income-by-sharing-internet','','2024-05-16 01:06:14.000','2025-08-17 07:37:17.588',NULL,NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(7,4,'published_hidden','[\"SEO\", \"Marketing Insight\"]','Worlds shortest guide on Ui Ux',NULL,'SEO is the strategic art of enhancing a website\'s visibility on search engines like Google, Bing, and Yahoo. It involves optimizing various elements to rank higher in search results...','https://static.wixstatic.com/media/5af200_7fe19f9ddee04ecfbaafd51a304008e2~mv2.png/v1/fill/w_740,h_423,al_c,q_85,usm_0.66_1.00_0.01,enc_avif,quality_auto/5af200_7fe19f9ddee04ecfbaafd51a304008e2~mv2.png','what-is-seo','','2024-05-15 01:05:14.000','2025-08-17 07:37:17.588',NULL,NULL,'what is SEO, Search Engine Optimization, what is website optimization, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(8,4,'published_hidden','[\"SEO\", \"Marketing Insight\"]','13 vintage websites that showcase timeless retro web design',NULL,'SEO is the strategic art of enhancing a website\'s visibility on search engines like Google, Bing, and Yahoo. It involves optimizing various elements to rank higher in search results...','/media/local/sites_example.avif','what-is-seo','','2024-05-15 01:05:14.000','2025-08-17 07:37:17.588',NULL,NULL,'what is SEO, Search Engine Optimization, what is website optimization, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(9,2,'published_hidden','[\"Passive Income\", \"Internet Monetization\"]','How to build a software engineering, coding and development portfolio',NULL,'Transform your internet connection into a Passive Income Stream by sharing your internet','/media/local/case_study.avif','passive-income-by-sharing-internet','','2024-05-16 01:06:14.000','2025-08-17 07:37:17.588',NULL,NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(10,1,'published_hidden','[\"Mysterium Node\"]','21 social media content ideas, plus best posts that go viral',NULL,'21 social media content ideas, plus best posts that go viral','/media/local/site_building_guide.avif','mysterium-node-backup-and-restore','','2024-05-01 01:06:14.000','2025-08-17 07:37:17.588',NULL,NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(11,4,'published_hidden','[\"SEO\", \"Marketing Insight\"]','What is SEO? A beginner’s guide to search engine optimization','What is SEO ?','SEO is the strategic art of enhancing a website\'s visibility on search engines like Google, Bing, and Yahoo. It involves optimizing various elements to rank higher in search results...','/media/4/img/what-is-seo.webp','what-is-seo','','2024-05-15 06:35:14.000','2025-08-17 07:37:17.588',NULL,NULL,'what is SEO, Search Engine Optimization, what is website optimization, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(12,2,'published_hidden','[\"Passive Income\", \"Internet Monetization\"]','Passive income by sharing internet',NULL,'Transform your internet connection into a Passive Income Stream by sharing your internet','/media/2/img/passive-income-by-sharing-internet.webp','passive-income-by-sharing-internet','','2024-05-16 06:36:14.000','2025-08-17 07:37:17.588',NULL,NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(13,1,'published_hidden','[\"Mysterium Node\"]','Mysterium node backup & restore',NULL,'Backup & restore your mysterium node on windows and linux','/media/1/img/mysterium-node-backup-and-restore.webp','mysterium-node-backup-and-restore','','2024-05-01 06:36:14.000','2025-08-17 07:37:17.588',NULL,NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(14,1,'published_hidden',NULL,'Does my business need a website? 9 reasons why it does',NULL,'The Internet powers virtually every aspect of the modern economy, transform your Internet connection into a passive income stream by merely sharing your Internet connection ','https://static.wixstatic.com/media/1f6616_88503e0f296f49afa70a53e3812fd92e~mv2.png/v1/fill/w_740,h_489,al_c,q_90,usm_0.66_1.00_0.01,enc_avif,quality_auto/1f6616_88503e0f296f49afa70a53e3812fd92e~mv2.png','passive-income-by-sharing-internet','','2024-05-17 06:36:14.000','2025-08-17 07:37:17.588',NULL,NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(15,1,'published_hidden','[\"Mysterium Node\"]','What is a Domain name and why it matters',NULL,'Backup & restore your mysterium node on windows and linux ','https://static.wixstatic.com/media/84b06e_74efe24ffe3f4cccbfb4d23e904498eb~mv2.png/v1/fill/w_740,h_423,al_c,q_85,usm_0.66_1.00_0.01,enc_avif,quality_auto/84b06e_74efe24ffe3f4cccbfb4d23e904498eb~mv2.png','mysterium-node-backup-and-restore','','2024-05-01 01:06:14.000','2025-08-17 07:37:17.588',NULL,NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(16,2,'published_hidden','[\"Passive Income\", \"Internet Monetization\"]','10 outstanging site examples',NULL,'Transform your internet connection into a Passive Income Stream by sharing your internet ','https://static.wixstatic.com/media/46e2e0_0b737cde10004f8c8a4deade3e45e47e~mv2.png/v1/fill/w_740,h_489,al_c,q_90,usm_0.66_1.00_0.01,enc_avif,quality_auto/46e2e0_0b737cde10004f8c8a4deade3e45e47e~mv2.png','passive-income-by-sharing-internet','','2024-05-16 01:06:14.000','2025-08-17 07:37:17.588',NULL,NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(17,4,'published_hidden','[\"SEO\", \"Marketing Insight\"]','Worlds shortest guide on Ui Ux',NULL,'SEO is the strategic art of enhancing a website\'s visibility on search engines like Google, Bing, and Yahoo. It involves optimizing various elements to rank higher in search results... ','https://static.wixstatic.com/media/5af200_7fe19f9ddee04ecfbaafd51a304008e2~mv2.png/v1/fill/w_740,h_423,al_c,q_85,usm_0.66_1.00_0.01,enc_avif,quality_auto/5af200_7fe19f9ddee04ecfbaafd51a304008e2~mv2.png','what-is-seo','','2024-05-15 01:05:14.000','2025-08-17 07:37:17.588',NULL,NULL,'what is SEO, Search Engine Optimization, what is website optimization, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(18,4,'published_hidden','[\"SEO\", \"Marketing Insight\"]','13 vintage websites that showcase timeless retro web design',NULL,'SEO is the strategic art of enhancing a website\'s visibility on search engines like Google, Bing, and Yahoo. It involves optimizing various elements to rank higher in search results... ','/media/local/sites_example.avif','what-is-seo','','2024-05-15 01:05:14.000','2025-08-17 07:37:17.588',NULL,NULL,'what is SEO, Search Engine Optimization, what is website optimization, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(19,2,'published_hidden','[\"Passive Income\", \"Internet Monetization\"]','How to build a software engineering, coding and development portfolio',NULL,'Transform your internet connection into a Passive Income Stream by sharing your internet ','/media/local/case_study.avif','passive-income-by-sharing-internet','','2024-05-16 01:06:14.000','2025-08-17 07:37:17.588',NULL,NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad'),
(20,1,'published_hidden','[\"Mysterium Node\"]','21 social media content ideas, plus best posts that go viral',NULL,'21 social media content ideas, plus best posts that go viral ','/media/local/site_building_guide.avif','mysterium-node-backup-and-restore','','2024-05-01 01:06:14.000','2025-08-17 07:37:17.588',NULL,NULL,'how to make money online, ways to make maney online, xet, xet industries, XetIndustries, XetIndustries blog, Rishikesh Prasad');
/*!40000 ALTER TABLE `blogs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `media`
--

DROP TABLE IF EXISTS `media`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `media` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `type` enum('audio','image','video','document') NOT NULL,
  `hash` varchar(255) NOT NULL,
  `filename` varchar(255) NOT NULL,
  `path` varchar(255) NOT NULL,
  `size` bigint(20) unsigned NOT NULL,
  `format` varchar(50) NOT NULL,
  `width` int(10) unsigned DEFAULT NULL,
  `height` int(10) unsigned DEFAULT NULL,
  `duration` int(10) unsigned DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `hash` (`hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `media`
--

LOCK TABLES `media` WRITE;
/*!40000 ALTER TABLE `media` DISABLE KEYS */;
/*!40000 ALTER TABLE `media` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `refresh_tokens`
--

DROP TABLE IF EXISTS `refresh_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `refresh_tokens` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) unsigned NOT NULL,
  `token_hash` varchar(128) NOT NULL,
  `expires_at` datetime NOT NULL,
  `revoked` tinyint(1) NOT NULL DEFAULT 0,
  `replaced_by_id` bigint(20) unsigned DEFAULT NULL,
  `user_agent` varchar(255) DEFAULT NULL,
  `ip_address` varchar(45) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `token_hash` (`token_hash`),
  KEY `idx_refresh_tokens_uid` (`uid`),
  KEY `idx_refresh_tokens_replaced_by_id` (`replaced_by_id`),
  CONSTRAINT `fk_refresh_token_user` FOREIGN KEY (`uid`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `refresh_tokens`
--

LOCK TABLES `refresh_tokens` WRITE;
/*!40000 ALTER TABLE `refresh_tokens` DISABLE KEYS */;
INSERT INTO `refresh_tokens` VALUES
(1,10,'ae52fad89dabe935e95159f95f098b9e31fda44378b7001eca2a8805ca67825c','2025-11-15 19:27:33',0,NULL,'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','::1','2025-11-08 19:27:33','2025-11-08 19:27:33'),
(2,16,'c10b4c178ed25057314d88cffd2077996d669f13fd38f71a9b7d4d46193f14e9','2025-11-17 00:00:42',0,NULL,'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','::1','2025-11-10 00:00:42','2025-11-10 00:00:42'),
(3,16,'97d064540ede7b9634eda1bab09497fa1dc2f71df253ad4d1be2c525c83e8160','2025-11-17 00:00:46',0,NULL,'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','::1','2025-11-10 00:00:46','2025-11-10 00:00:46'),
(4,17,'284dd5486bc5b8686cc90b641f5d45ef173ee90e4ee0659337990d3dd011e193','2025-11-29 12:48:13',0,NULL,'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','103.10.116.73','2025-11-22 12:48:13','2025-11-22 12:48:13'),
(5,17,'b34cf292b6b52b6935bd4fe332fba5b8f4d9b6158123e50af8b0f56e9c9fb90a','2025-11-30 13:35:56',0,NULL,'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','2405:201:9005:29a2:d467:ec95:b36a:90ff','2025-11-23 13:35:56','2025-11-23 13:35:56');
/*!40000 ALTER TABLE `refresh_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_media`
--

DROP TABLE IF EXISTS `user_media`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_media` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) unsigned NOT NULL,
  `media_id` bigint(20) unsigned NOT NULL,
  `tags` varchar(255) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid` (`uid`,`media_id`),
  KEY `media_id` (`media_id`),
  CONSTRAINT `user_media_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `user_media_ibfk_2` FOREIGN KEY (`media_id`) REFERENCES `media` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_media`
--

LOCK TABLES `user_media` WRITE;
/*!40000 ALTER TABLE `user_media` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_media` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL,
  `name` varchar(50) NOT NULL,
  `email` varchar(120) NOT NULL,
  `verified` tinyint(1) NOT NULL DEFAULT 0,
  `password_hash` varchar(255) NOT NULL,
  `role` enum('admin','dev','user') DEFAULT 'user',
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `last_login` datetime DEFAULT NULL,
  `status` enum('active','inactive','suspended','deleted') DEFAULT 'active',
  `avatar_url` varchar(255) DEFAULT NULL,
  `address` text DEFAULT NULL,
  `phone_no` varchar(20) DEFAULT NULL,
  `dob` date DEFAULT NULL,
  `config` longtext DEFAULT NULL CHECK (json_valid(`config`)),
  `remember_token` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_status` (`status`),
  KEY `idx_role` (`role`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES
(1,'xet','Rishikesh Prasad','rishikeshprasad@xetindustries.com',1,'$2y$12$3zXdoKP91LvYctZUrN9cYOfBV8TyAEeoQK5ADAVQuBINSBt1nJRTK','user','2024-07-27 04:05:01','2025-02-01 13:37:18',NULL,'active','/media/1/profile/xet.jpg',NULL,NULL,NULL,NULL,NULL),
(2,'zet','Zet Ohio','zet@g.com',1,'$2y$12$3zXdoKP91LvYctZUrN9cYOfBV8TyAEeoQK5ADAVQuBINSBt1nJRTK','user','2024-07-27 04:13:44','2025-02-01 13:37:18',NULL,'active','/media/2/profile/zet.jpg',NULL,NULL,NULL,NULL,NULL),
(4,'cristine','Cristine Lepcha','cr@g.com',0,'$2y$12$3zXdoKP91LvYctZUrN9cYOfBV8TyAEeoQK5ADAVQuBINSBt1nJRTK','user','2024-07-27 04:13:44','2025-02-01 13:37:18',NULL,'active','/media/4/profile/cristine.jpg',NULL,NULL,NULL,NULL,NULL),
(5,'t1','t1','t1@g.com',0,'$2y$12$3zXdoKP91LvYctZUrN9cYOfBV8TyAEeoQK5ADAVQuBINSBt1nJRTK','user','2024-07-27 04:13:44','2024-11-29 21:22:51',NULL,'active','/media/3/profile/zet.jpg',NULL,NULL,NULL,NULL,NULL),
(7,'wbeiliq41','UyXhRcENW','wbeiliq41@gmail.com',0,'$2y$12$pTONJnvKco8ksZaLcXvWKuXkzSlMksBvipuq8xueLxfq15RlYDUne','user','2025-03-16 06:02:15','2025-03-16 06:02:15',NULL,'active',NULL,NULL,NULL,NULL,NULL,NULL),
(9,'hey','hey','hey@gmail.com',0,'$2y$12$QMOBTM62Y2YNUkv.wGdBseZuF07BEBZMxr.3PzT5nTm6QDP/hd/z6','user','2025-03-22 14:47:46','2025-03-22 14:47:46',NULL,'active',NULL,NULL,NULL,NULL,NULL,NULL),
(10,'','','itszet1o1@gmail.com',0,'$2a$10$voIXZEWvgYrXAq9d9gcUD.pRprcHZkideKHLYpmyI/1VFr.zfO2iW','user','2025-11-08 19:25:51','2025-11-08 19:27:33','2025-11-08 19:27:33','active',NULL,NULL,NULL,NULL,NULL,NULL),
(16,'luan','','luanka@gmail.com',0,'$2a$10$gcWKRi5kOEbJdz5clge75OzSkSdNzs4Otyb9JW/VXGq2HeShIuGCm','user','2025-11-10 00:00:19','2025-11-10 00:00:46','2025-11-10 00:00:46','active',NULL,NULL,NULL,NULL,NULL,NULL),
(17,'zet1o1','','zet1o1@gmail.com',0,'$2a$10$0wsaYo778lsgnlTcjimSsOwTHupqn.xHXHfS3/z7VcAbnb1mB5P0K','user','2025-11-22 12:48:07','2025-11-23 13:35:56','2025-11-23 13:35:56','active',NULL,NULL,NULL,NULL,NULL,NULL),
(18,'yajvrocoaxsgsmall','','kamufisaru344@gmail.com',0,'$2a$10$1PQmfZV/IdfvjUqgA9ovH.aAuWagWvRGtxFpo/8Gz9S6RbJhBMsGq','user','2025-12-05 16:08:33','2025-12-05 16:08:33',NULL,'active',NULL,NULL,NULL,NULL,NULL,NULL),
(21,'zfvapnbtjutusmgbuck','','qipamamaf07@gmail.com',0,'$2a$10$d6P8ZgQQEC/ZzbvBviPaeeBVHkSXttU7lhQ.Gia6INW8R7xNSHpQq','user','2025-12-14 00:57:44','2025-12-14 00:57:44',NULL,'active',NULL,NULL,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'XI'
--

--
-- Dumping routines for database 'XI'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-12-24 14:14:01
