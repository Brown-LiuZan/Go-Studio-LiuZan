package stdpkgusage

import (
	"log"
	"os"
	"syscall"
	"testing"
)

var tlog = log.New(os.Stderr, "[stdpkgusage]",
	log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

func TestOsEnvMgr(t *testing.T) {
	tlog.Println("Begin of", t.Name())

	if host, err := os.Hostname(); err != nil {
		tlog.Fatalln("Failed to os.Hostname()")
	} else {
		tlog.Println("Hostname from os.Hostname():", host)
	}

	tlog.Println("Temprary directory from os.Tempdir():", os.TempDir())

	if cacheDir, err := os.UserCacheDir(); err != nil {
		tlog.Fatalln("Failed to os.UserCacheDir()")
	} else {
		tlog.Println("User cache directory from os.UserCacheDir():", cacheDir)
	}

	if homeDir, err := os.UserHomeDir(); err != nil {
		tlog.Fatalln("Failed to os.UserHomeDir()")
	} else {
		tlog.Println("User home directory from os.UserHomeDir():", homeDir)
	}

	tlog.Println("UID:", os.Getuid(), "GID:", os.Getgid())
	if groups, err := os.Getgroups(); err != nil {
		tlog.Fatalln("Failed to os.Getgroups()")
	} else {
		tlog.Println("Groups:", groups)
	}
	tlog.Println("EUID:", os.Geteuid(), "EGID:", os.Getegid())

	if cwd, err := os.Getwd(); err != nil {
		tlog.Fatalln("Failed to get working directory with os.Getwd().")
	} else {
		tlog.Println("Current working directory:", cwd)
	}

	newCwd := "/home/brown"
	if os.Chdir(newCwd) != nil {
		tlog.Fatalf("Failed to os.Chdir('%s').\n", newCwd)
	}
	cwd, _ := os.Getwd()
	tlog.Printf("New working directory: %s Expected: %s\n", cwd, newCwd)

	tlog.Println("PATH from os.ExpandEnv('${PATH}'):", os.ExpandEnv("${PATH}"))
	tlog.Println("PATH from os.ExpandEnv('$PATH'):", os.ExpandEnv("$PATH"))

	tlog.Println("KV pairs in env returned from os.Environ():",
		len(os.Environ()))

	os.Clearenv()
	tlog.Println("After os.ClearEnv(), KV pairs in env:", len(os.Environ()))

	if os.Setenv("PATH", "/home/liuzan.lz") != nil {
		tlog.Fatalln("Failed to os.Setenv('/home/liuzan.lz')")
	}
	if path, ok := os.LookupEnv("PATH"); !ok {
		tlog.Fatalln("Failed to os.LookupEnv('PATH') after os.Setenv('PATH',)")
	} else {
		tlog.Println("After os.Setenv('PATH','/home/liuzan.lz'),",
			"PATH from os.LookupEnv('PATH'):",
			path)
	}
	tlog.Println("After os.Setenv('PATH','/home/liuzan.lz'),",
		"PATH from os.GetEnv('PATH'):", os.Getenv("PATH"))
	tlog.Println("After os.Setenv('PATH','/home/liuzan.lz'),",
		"PATH from os.ExpandEnv('${PATH}'):", os.ExpandEnv("${PATH}"))
	tlog.Println("After os.Setenv('PATH','/home/liuzan.lz'),",
		"PATH from os.ExpandEnv('$PATH'):", os.ExpandEnv("$PATH"))
	tlog.Println("After os.Setenv('PATH','/home/liuzan.lz'),",
		"PATH from os.Expand('${PATH}', os.Getenv()):",
		os.Expand("${PATH}", os.Getenv))
	tlog.Println("After os.Setenv('PATH','/home/liuzan.lz'),",
		"PATH from os.Expand('$PATH', os.Getenv()):",
		os.Expand("$PATH", os.Getenv))

	if os.Unsetenv("PATH") != nil {
		tlog.Fatalln("Failed to os.Unsetenv('PATH')")
	} else if path, ok := os.LookupEnv("PATH"); ok {
		tlog.Fatalln("Should fail to os.LookupEnv('PATH'):", path)
	} else {
		tlog.Println("Succeeded to os.Unsetenv('PATH')")
	}

	tlog.Println("Page size from os.Getpagesize():", os.Getpagesize(), "bytes")

	tlog.Printf("PathSeparator: %v %[1]q\n", os.PathSeparator)
	tlog.Printf("PathListSeparator: %v %[1]q\n", os.PathListSeparator)
	tlog.Println("os.IsPathSeparator('/'):", os.IsPathSeparator('/'))
	tlog.Println("os.IsPathSeparator('\\'):", os.IsPathSeparator('\\'))

	tlog.Println("End of", t.Name())
}

func TestOsProcMgr(t *testing.T) {
	tlog.Println("Begin of", t.Name())

	tlog.Println("PID:", os.Getpid(), "PPID:", os.Getppid())

	if proc, err := os.FindProcess(os.Getppid()); err != nil {
		tlog.Fatalln("Failed to os.FindProcess(os.Getppid())")
	} else {
		tlog.Println("proc.Pid:", proc.Pid, "expected:", os.Getppid())
	}

	procAttr := new(os.ProcAttr)
	procAttr.Dir = os.TempDir()
	if len(procAttr.Env) > 0 {
		tlog.Println("len(procAttr.Env) of newed ProcAttr:", len(procAttr.Env))
	} else {
		tlog.Println("Init newed ProcAttr.Env is empty")
	}
	procAttr.Env = append(procAttr.Env, "UserName=Brown")
	if len(procAttr.Files) > 0 {
		tlog.Println("len(procAttr.Files) of newed ProcAttr:", len(procAttr.Files))
	} else {
		tlog.Println("Init newed ProcAttr.Files is empty")
		procAttr.Files = append(procAttr.Files, os.Stdin)
		procAttr.Files = append(procAttr.Files, os.Stdout)
		procAttr.Files = append(procAttr.Files, os.Stderr)
	}

	newProc, err := os.StartProcess("/usr/bin/touch",
		[]string{"/usr/bin/touch", "brown-go-test.txt"}, procAttr)
	if err != nil {
		tlog.Fatalln("Failed to os.StartProcess()")
	}
	procState, err := newProc.Wait()
	if err != nil {
		tlog.Fatalln("Failed to os.Wait()")
	}
	tlog.Println("Is exited?", procState.Exited())
	tlog.Println("Succeeded?", procState.Success())
	tlog.Println("ExitCode:", procState.ExitCode())
	tlog.Println("PID:", procState.Pid())
	tlog.Println("ProcessState:", procState)
	tlog.Println("Sys:", procState.Sys().(syscall.WaitStatus))
	tlog.Println("SysUsage:", procState.SysUsage().(*syscall.Rusage))
	tlog.Println("SystemTime:", procState.SystemTime())
	tlog.Println("UserTime:", procState.UserTime())

	tlog.Println("End of", t.Name())
}

func TestOsFileMgr(t *testing.T) {
	tlog.Println("Begin of", t.Name())

	tlog.Println("End of", t.Name())
}

func TestOsFileOps(t *testing.T) {
	tlog.Println("Begin of", t.Name())

	tlog.Println("End of", t.Name())
}

func TestOsDirMgr(t *testing.T) {
	tlog.Println("Begin of", t.Name())

	tlog.Println("End of", t.Name())
}

func TestOsDirOps(t *testing.T) {
	tlog.Println("Begin of", t.Name())

	tlog.Println("End of", t.Name())
}

func TestOsHardlink(t *testing.T) {
	tlog.Println("Begin of", t.Name())

	tlog.Println("End of", t.Name())
}

func TestOsSymlink(t *testing.T) {
	tlog.Println("Begin of", t.Name())

	tlog.Println("End of", t.Name())
}

func TestOsPipeline(t *testing.T) {
	tlog.Println("Begin of", t.Name())

	tlog.Println("End of", t.Name())
}
