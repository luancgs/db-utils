package databases

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	POSTGRES_DEFAULT_DOCKER_IMAGE    = "postgres"
	POSTGRES_DEFAULT_PORT            = 5432
	POSTGRES_DEFAULT_PASSWORD        = ""
	POSTGRES_DEFAULT_COMMAND_DUMP    = "pg_dump"
	POSTGRES_DEFAULT_COMMAND_RESTORE = "psql"
	POSTGRES_DEFAULT_COMMAND_QUERY   = "psql"
)

type PostgresDatabase struct {
	Protocol string
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Query    string
}

func (db PostgresDatabase) Dump(isTemporary bool, dir string) (string, error) {
	port := fmt.Sprintf("%d", db.Port)
	passwordInput := fmt.Sprintf("PGPASSWORD=%s", db.Password)

	cmd := exec.Command("docker", "run", "--rm", "--network", "host", "--env", passwordInput, POSTGRES_DEFAULT_DOCKER_IMAGE, POSTGRES_DEFAULT_COMMAND_DUMP, "-U", db.Username, "-h", db.Host, "-p", port, "-d", db.Database, "--clean", "--encoding=UTF8")

	stdout, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed running docker to dump database: %w", err)
	}

	dumpFile := "dump_*.sql"

	if isTemporary {
		dir = os.TempDir()
	} else {
		if dir == "" {
			dir, err = os.UserHomeDir()
			if err != nil {
				return "", fmt.Errorf("failed getting user home directory: %w", err)
			}
		}

		date := time.Now().Local().Format("2006-01-02_15-04-05")
		dumpFile = fmt.Sprintf("dump_%s_%s_*.sql", db.Database, date)
	}

	file, err := os.CreateTemp(dir, dumpFile)
	if err != nil {
		return "", fmt.Errorf("failed creating dump file: %w", err)
	}

	_, err = file.Write(stdout)
	if err != nil {
		return "", fmt.Errorf("failed writing dump to file: %w", err)
	}

	return file.Name(), nil
}

func (db PostgresDatabase) Restore(fileName string) (bool, error) {
	fmt.Printf("Populating database: %s\n", db.Database)

	containerFile := "/input.sql"
	volume := fmt.Sprintf("%s:%s", fileName, containerFile)
	port := fmt.Sprintf("%d", db.Port)
	passwordInput := fmt.Sprintf("PGPASSWORD=%s", db.Password)

	cmd := exec.Command("docker", "run", "--rm", "--network", "host", "--volume", volume, "--env", passwordInput, POSTGRES_DEFAULT_DOCKER_IMAGE, POSTGRES_DEFAULT_COMMAND_RESTORE, "-U", db.Username, "-h", db.Host, "-p", port, "-d", db.Database, "-f", containerFile)

	_, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed running docker to populate database: %w", err)
	}

	return true, nil
}

func (db PostgresDatabase) RunQuery(query string) (string, error) {
	port := fmt.Sprintf("%d", db.Port)
	passwordInput := fmt.Sprintf("PGPASSWORD=%s", db.Password)

	cmd := exec.Command("docker", "run", "--rm", "--network", "host", "--env", passwordInput, POSTGRES_DEFAULT_DOCKER_IMAGE, POSTGRES_DEFAULT_COMMAND_QUERY, "-U", db.Username, "-h", db.Host, "-p", port, "-d", db.Database, "-c", query)

	stdout, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed running query: %s", err)
	}

	return string(stdout), nil
}

func parsePostgres(parsedUrl *url.URL) (Database, error) {

	db := PostgresDatabase{
		Protocol: parsedUrl.Scheme,
		Host:     parsedUrl.Hostname(),
		Port:     POSTGRES_DEFAULT_PORT,
		Database: parsedUrl.Path[1:],
		Query:    parsedUrl.RawQuery,
		Username: parsedUrl.User.Username(),
		Password: POSTGRES_DEFAULT_PASSWORD,
	}

	if parsedUrl.Port() != "" {
		port, err := strconv.Atoi(parsedUrl.Port())
		if err != nil {
			return nil, fmt.Errorf("failed converting port to string (invalid port): %w", err)
		}

		db.Port = port
	}

	if password, ok := parsedUrl.User.Password(); ok {
		db.Password = password
	}

	return db, nil
}
