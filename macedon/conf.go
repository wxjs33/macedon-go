package macedon

import (
	goconf "github.com/msbranco/goconfig"
	"path/filepath"
	//"time"
)

const DEFAULT_CONF_PATH		= "../conf"
const DEFAULT_CONF			= "macedon.conf"

type Config struct {
	addr	string  /* server bind address */

	create	string	/* create location */
	remove	string	/* remove location */
	update	string  /* update location */
	read	string	/* read location */

	maddr	string	/* mysql addr */

	dbname	string	/* db name */
	dbuser	string	/* db username */
	dbpwd	string	/* db password */

	log		string	/* log file */
	level	string	/* log level */
}

func (conf *Config) ReadConf(file string) (*Config, error) {
	if file == "" {
		file = filepath.Join(DEFAULT_CONF_PATH, DEFAULT_CONF)
	}

	c, err := goconf.ReadConfigFile(file)
	if err != nil {
		return nil, err
	}

	//TODO: check
	conf.addr, _		= c.GetString("default", "addr")
	conf.create, _		= c.GetString("default", "create_location")
	conf.remove, _		= c.GetString("default", "delete_location")
	conf.update, _		= c.GetString("default", "update_location")
	conf.read, _		= c.GetString("default", "read_location")

	conf.log, _			= c.GetString("default", "log")
	conf.level, _		= c.GetString("default", "level")
	conf.maddr, _		= c.GetString("default", "mysql_addr")
	conf.dbname, _		= c.GetString("default", "mysql_dbname")
	conf.dbuser, _		= c.GetString("default", "mysql_dbuser")
	conf.dbpwd, _		= c.GetString("default", "mysql_dbpwd")

	return conf, nil
}
