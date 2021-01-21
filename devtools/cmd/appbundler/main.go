package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/evanw/esbuild/pkg/api"
	_ "github.com/evanw/esbuild/pkg/api"
	"github.com/zapkub/pakkretqc/internal/fsutil"

	"github.com/bmatcuk/doublestar"
	"github.com/radovskyb/watcher"
)

func build() {
	log.Println("rebuild tsx file...")
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{
			fsutil.PathFromWebDir("app/index.tsx"),
			fsutil.PathFromWebDir("app/login.tsx"),
			fsutil.PathFromWebDir("app/domain.tsx"),
			fsutil.PathFromWebDir("app/project.tsx"),
			fsutil.PathFromWebDir("app/defect.tsx"),
		},
		Outdir:    fsutil.PathFromWebDir("dist"),
		Bundle:    true,
		Write:     true,
		Splitting: true,
		Format:    api.FormatESModule,
		Define: map[string]string{
			"process.env.NODE_ENV": "'development'",
		},
		// MinifySyntax:      true,
		// MinifyIdentifiers: true,
		// MinifyWhitespace:  true,
		Platform: api.PlatformBrowser,
		Tsconfig: fsutil.PathFromWebDir("app/tsconfig.json"),
	})

	if len(result.Errors) > 0 {
		log.Fatalf("build error: %+v", result.Errors)
	}

	exec.Command("/bin/bash", "-c", fmt.Sprintf("cp -R %s %s", fsutil.PathFromWebDir("styles"), fsutil.PathFromWebDir("dist"))).Run()

}

func main() {
	var sig = make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	var reload = make(chan struct{}, 1)
	build()
	go watchGlob(reload, strings.Join([]string{
		fsutil.PathFromWebDir("app/*/**"),
		fsutil.PathFromWebDir("app/*"),
		fsutil.PathFromWebDir("styles/*"),
		fsutil.PathFromWebDir("common/*"),
	}, ","))
	for {
		select {
		case <-sig:
			fmt.Println("signal for kill this command...")
			os.Exit(0)
		case <-reload:
			build()
		}
	}

}

// code from gorogoso https://github.com/zapkub/gorogoso/blob/master/runner/lib.go
func watchGlob(reload chan struct{}, glob string) {
	w := watcher.New()
	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			select {
			case event := <-w.Event:
				fmt.Printf("\n[gosogoso] watcher tigger...\n")
				fmt.Printf("[gorogoso] %s\n", event)
				reload <- struct{}{}
			case <-w.Closed:
				return
			}
		}
	}()

	w.SetMaxEvents(1)
	w.FilterOps(watcher.Write)
	globList := strings.Split(glob, ",")
	fmt.Println("[gorogoso] watch file list")
	for _, g := range globList {
		paths, _ := doublestar.Glob(g)
		for _, path := range paths {
			fmt.Println(path)
			if err := w.Add(path); err != nil {
				panic(err)
			}
		}
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
