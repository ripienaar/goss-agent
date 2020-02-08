package main

func shouldActivate(_ string, config map[string]string) (bool, error) {
	exec := gossExecutable(config)

	return exec != "", nil
}
