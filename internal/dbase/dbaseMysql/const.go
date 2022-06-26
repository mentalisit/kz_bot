package dbaseMysql

const (
	config = `CREATE TABLE IF NOT EXISTS config(
		id int primary key auto_increment, 
		corpname VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		dschannel VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		tgchannel BIGINT(24) NULL DEFAULT NULL, 
		wachannel VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		mesiddshelp VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		mesidtghelp INT(11) NULL DEFAULT '0',
		delmescomplite INT(11) NULL DEFAULT NULL,
    	guildid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'
		)`
	numkz = `CREATE TABLE IF NOT EXISTS numkz(
		id int primary key auto_increment,
		lvlkz INT(11) NOT NULL DEFAULT '0',
		number INT(11) NOT NULL DEFAULT '0',
		corpname VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'
		)`
	rsevent = `CREATE TABLE IF NOT EXISTS rsevent(
		id int primary key auto_increment,
		corpname VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		numevent INT(11) NULL DEFAULT NULL,
		activeevent INT(11) NULL DEFAULT NULL,
		number INT(11) NULL DEFAULT NULL
	)`
	sborkz = `CREATE TABLE IF NOT EXISTS sborkz(
		id int primary key auto_increment,
		corpname VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		name VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		mention VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		tip VARCHAR(10) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		dsmesid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		tgmesid INT(11) NULL DEFAULT '0',
		wamesid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		time TIME NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		date DATE NULL DEFAULT NULL,
		lvlkz INT(11) NULL DEFAULT NULL,
		numkzn INT(11) NULL DEFAULT NULL,
		numberkz INT(11) NULL DEFAULT NULL,
		numberevent INT(11) NULL DEFAULT NULL,
		eventpoints INT(11) NULL DEFAULT NULL,
		active INT(11) NULL DEFAULT NULL,
		timedown INT(11) UNSIGNED NULL DEFAULT NULL
		)`
	subscribe = `CREATE TABLE IF NOT EXISTS subscribe(
		id int primary key auto_increment, 
		name VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		nameid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		lvlkz INT(11) NULL DEFAULT '0',
		tip INT(11) NULL DEFAULT '0',
		chatid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		timestart VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		timeend VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'
		)`
	timer = `CREATE TABLE IF NOT EXISTS timer(
		id int primary key auto_increment,
		dsmesid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		dschatid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		tgmesid INT(11) NULL DEFAULT '0',
		tgchatid BIGINT(50) NULL DEFAULT NULL,
		timed INT(11) NULL DEFAULT '0'
		)`
	users = `CREATE TABLE IF NOT EXISTS users(
		id int primary key auto_increment,
		tip VARCHAR(30) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		name VARCHAR(30) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		em1 VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		em2 VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		em3 VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		em4 VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'
		)`
	temptopevent = `CREATE TABLE IF NOT EXISTS temptopevent(
		id int primary key auto_increment,
		name VARCHAR(30) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		numkz INT(11) NULL DEFAULT NULL,
		points INT(11) NULL DEFAULT NULL
	)`
)
