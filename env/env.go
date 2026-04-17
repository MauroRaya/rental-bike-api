package env

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Env struct {
	DSN  string
	Port int32
}

func parse(file io.Reader) map[string]string {
	scanner := bufio.NewScanner(file)
	envs := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		i := strings.Index(line, "=")

		if i == -1 {
			continue
		}

		key := line[:i]
		value := line[i+1:]

		envs[key] = value
	}

	return envs
}

func sets(envs map[string]string) error {
	for k, v := range envs {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}

	return nil
}

func validate(envs map[string]string, env *Env) error {
	value, ok := envs["DSN"]
	if !ok {
		return fmt.Errorf("env DSN is not set")
	}
	env.DSN = value

	value, ok = envs["PORT"]
	if !ok {
		return fmt.Errorf("env PORT is not set")
	}

	port, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return fmt.Errorf("env PORT is not int32")
	}
	env.Port = int32(port)

	return nil
}

func Load(fileName string, shouldValidate bool) (Env, error) {
	var env Env

	file, err := os.Open(fileName)
	if err != nil {
		return env, err
	}
	defer file.Close()

	envs := parse(file)

	if err := sets(envs); err != nil {
		return env, err
	}

	if err := validate(envs, &env); err != nil {
		return env, err
	}

	return env, nil
}
