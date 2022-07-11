package dbasePostgres

const (
	CreateTableConfig = `CREATE TABLE IF NOT EXISTS config
		(
		id          bigint    NOT NULL PRIMARY KEY, 
		corpname        text NOT NULL,
		dschannel        text NOT NULL,
		tgchannel        bigint NOT NULL,
		wachannel        text NOT NULL,
		mesiddshelp        text NOT NULL,
		mesidtghelp        bigint NOT NULL,
		delmescomplite        bigint NOT NULL,
		guildid        text NOT NULL
		)`
	CreateTableNumkz = `CREATE TABLE IF NOT EXISTS numkz(
		  id          bigserial    NOT NULL PRIMARY KEY,
		lvlkz INT(11) NOT NULL DEFAULT '0',
		number INT(11) NOT NULL DEFAULT '0',
		corpname VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'
		)`
	CreateTableRsevent = `CREATE TABLE IF NOT EXISTS rsevent(
		  id          bigserial    NOT NULL PRIMARY KEY,
		corpname VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		numevent INT(11) NULL DEFAULT NULL,
		activeevent INT(11) NULL DEFAULT NULL,
		number INT(11) NULL DEFAULT NULL
	)`
	CreateTableSborkz = `CREATE TABLE IF NOT EXISTS sborkz(
		  id          bigserial    NOT NULL PRIMARY KEY,
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
	CreateTableSubscribe = `CREATE TABLE IF NOT EXISTS subscribe(
		  id          bigserial    NOT NULL PRIMARY KEY,
		name VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		nameid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		lvlkz INT(11) NULL DEFAULT '0',
		tip INT(11) NULL DEFAULT '0',
		chatid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		timestart VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		timeend VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'
		)`
	CreateTableTimer = `CREATE TABLE IF NOT EXISTS timer(
		  id          bigserial    NOT NULL PRIMARY KEY,
		dsmesid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		dschatid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		tgmesid INT(11) NULL DEFAULT '0',
		tgchatid BIGINT(50) NULL DEFAULT NULL,
		timed INT(11) NULL DEFAULT '0'
		)`
	CreateTableUsers = `CREATE TABLE IF NOT EXISTS users(
		  id          bigserial    NOT NULL PRIMARY KEY,
		tip VARCHAR(30) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		name VARCHAR(30) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		em1 VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		em2 VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		em3 VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		em4 VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'
		)`
	CreateTableTemptopevent = `CREATE TABLE IF NOT EXISTS temptopevent(
		  id          bigserial    NOT NULL PRIMARY KEY,
		name VARCHAR(30) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		numkz INT(11) NULL DEFAULT NULL,
		points INT(11) NULL DEFAULT NULL
	)`
)
