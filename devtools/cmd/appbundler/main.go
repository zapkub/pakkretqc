package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/evanw/esbuild/pkg/api"
	_ "github.com/evanw/esbuild/pkg/api"
	"github.com/zapkub/pakkretqc/internal/fsutil"
)

func main() {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{
			fsutil.PathFromWebDir("app/index.tsx"),
			fsutil.PathFromWebDir("app/login.tsx"),
			fsutil.PathFromWebDir("app/domain.tsx"),
			fsutil.PathFromWebDir("app/project.tsx"),
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
