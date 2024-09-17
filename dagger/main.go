// A generated module for Fmysql functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/fmysql/internal/dagger"
	"fmt"
)

type Fmysql struct{}

func (m *Fmysql) Fmysql(source *dagger.Directory) *dagger.Container {
	return dag.Container().From("mysql:8.4").
		WithFile("/docker-entrypoint-initdb.d/fixtures.sql", source.File("/fixtures-timestamp.sql"), dagger.ContainerWithFileOpts{
			Owner: "mysql:mysql",
		}).
		WithEnvVariable("MYSQL_USER", "fmysql").
		WithEnvVariable("MYSQL_PASSWORD", "password").
		WithEnvVariable("MYSQL_DATABASE", "fmysql").
		WithEnvVariable("MYSQL_RANDOM_ROOT_PASSWORD", "1").
		WithEnvVariable("MYSQL_LOG_CONSOLE", "1").
		WithFile("/etc/mysql/conf.d/docker.cnf", source.File("docker.cnf"), dagger.ContainerWithFileOpts{
			Permissions: 0644,
		}).
		WithDefaultArgs([]string{"mysqld", "--mysql-native-password=FORCE", "--sql_mode=NO_ENGINE_SUBSTITUTION"}).
		WithExposedPort(3306)
}

func (m *Fmysql) Testit(source *dagger.Directory) (string, error) {
	mysqlCtr := m.Fmysql(source)
	mysqlSvc := mysqlCtr.AsService()
	ctx := context.Background()

	_, _ = m.Print("overkill")

	return dag.Container().From("mysql:8.4").
		WithServiceBinding("mysqlsvc", mysqlSvc).
		WithFile("/docker-entrypoint-initdb.d/fixtures.sql", source.File("/fixtures.sql"), dagger.ContainerWithFileOpts{
			Owner: "mysql:mysql",
		}).
		WithEnvVariable("MYSQL_USER", "fmysql").
		WithEnvVariable("MYSQL_PASSWORD", "password").
		WithEnvVariable("MYSQL_DATABASE", "fmysql").
		WithEnvVariable("MYSQL_RANDOM_ROOT_PASSWORD", "1").
		WithExec([]string{
			"mysql",
			"-h", "mysqlsvc",
			"-u", "fmysql",
			"-ppassword",
			"-e", "USE fmysql; SELECT @@sql_mode;",
		}).
		WithExec([]string{
			"mysql",
			"-h", "mysqlsvc",
			"-u", "fmysql",
			"-ppassword",
			"-e", "USE fmysql; SELECT * FROM testable ORDER BY created_at DESC LIMIT 5;",
		}).Stdout(ctx)
}

func (m *Fmysql) Print(msg string) (string, error) {
	return dag.Container().From("alpine:3.14").WithExec([]string{"echo", msg}).Stdout(context.Background())
}

func (m *Fmysql) MatrixBuild(
	ctx context.Context,
	// +defaultPath="/"
	src *dagger.Directory,
) *dagger.Directory {
	gooses := []string{"linux", "darwin"}
	goarches := []string{"amd64", "arm64"}

	// create empty directory to put build artifacts
	outputs := dag.Directory()

	golang := dag.Container().
		From("golang:latest").
		WithDirectory("/src", src).
		WithWorkdir("/src")

	for _, goos := range gooses {
		for _, goarch := range goarches {
			// create directory for each OS and architecture
			path := fmt.Sprintf("build/%s/%s/", goos, goarch)

			// build artifact
			build := golang.
				WithEnvVariable("GOOS", goos).
				WithEnvVariable("GOARCH", goarch).
				WithExec([]string{"go", "build", "-o", path})

			// add build to outputs
			outputs = outputs.WithDirectory(path, build.Directory(path))
		}
	}

	// return build directory
	return outputs
}
