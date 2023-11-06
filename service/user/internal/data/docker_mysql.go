package data

import (
	"database/sql"
	"fmt"
	"github.com/ory/dockertest/v3"
	"log"
	"time"
)

func DockerMysql(img, version string) (string, func()) {
	return innerDockerMysql(img, version)
}

func innerDockerMysql(img, version string) (string, func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	pool.MaxWait = time.Minute * 2

	resource, err := pool.Run(img, version, []string{"MYSQL_ROOT_PASSWORD=secret", "MYSQL_ROOT_HOST=%"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	conStr := fmt.Sprintf("root:secret@(localhost:%s)/mysql?parseTime=true", resource.GetPort("3306/tcp"))
	if err := pool.Retry(func() error {
		db, err := sql.Open("mysql", conStr)
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return conStr, func() {
		if err = pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}
}
