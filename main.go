package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/peter-evans/patience"
)

func main() {
	original := strings.Split(originalContent, "\n")
	updated := strings.Split(updatedContent, "\n")

	diffs := patience.Diff(original, updated)

	// Combined diff
	// diff := patience.DiffText(diffs)
	// fmt.Println(originalContent, updatedContent)
	// Unified diff with options
	unidiffs := patience.UnifiedDiffTextWithOptions(
		diffs,
		patience.UnifiedDiffOptions{
			Precontext:  2,
			Postcontext: 2,
			SrcHeader:   "a/rollup.config.ts",
			DstHeader:   "b/rollup.config.ts",
		},
	)

	os.WriteFile("patience.patch", []byte(unidiffs), 0644)
	fmt.Println(unidiffs)
}

const originalContent = `import babel from '@rollup/plugin-babel';
import { nodeResolve } from '@rollup/plugin-node-resolve';
import { terser } from 'rollup-plugin-terser';

import pkg from './package.json';


const extensions = ['.ts'];

export default {
  input: 'src/index.ts',
  plugins: [
    nodeResolve({ extensions }),
    babel({
      extensions,
      babelHelpers: 'bundled',
      exclude: 'node_modules/**'
    }),
    terser()
  ],
  output: [
    {
      file: pkg.main,
      format: 'cjs',
      exports: 'auto',
      sourcemap: true
    },
    {
      file: pkg.module,
      format: 'es',
      sourcemap: true
    },
    {
      name: pkg.name,
      file: pkg.umd,
      format: 'umd',
      sourcemap: true
    }
  ]
};
`

const updatedContent = `import babel from '@rollup/plugin-babel'
import { nodeResolve } from '@rollup/plugin-node-resolve'
import { terser } from 'rollup-plugin-terser'

import pkg from './package.json'

const extensions = ['.ts']

export default {
    input: 'src/index.ts',
    plugins: [
        nodeResolve({ extensions }),
        babel({
            extensions,
            babelHelpers: 'bundled',
            exclude: 'node_modules/**',
        }),
        terser(),
    ],
    output: [
        {
            file: pkg.main,
            format: 'cjs',
            exports: 'auto',
            sourcemap: true,
        },
        {
            file: pkg.module,
            format: 'es',
            sourcemap: true,
        },
        {
            name: pkg.name,
            file: pkg.umd,
            format: 'umd',
            sourcemap: true,
        },
    ],
}
`
