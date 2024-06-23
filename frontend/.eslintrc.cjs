module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  extends: [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:react-hooks/recommended",
    "plugin:@typescript-eslint/recommended-requiring-type-checking",
  ],
  ignorePatterns: ["dist", ".eslintrc.cjs", "*.config.*"],
  parser: "@typescript-eslint/parser",
  parserOptions: {
    project: "./frontend/tsconfig.json",
  },
  plugins: ["@typescript-eslint", "react-refresh", "import"],
  settings: {
    "import/parsers": {
      "@typescript-eslint/parser": [".ts", ".tsx"],
    },
    "import/resolver": {
      typescript: {
        project: "./frontend/tsconfig.json",
      },
    },
  },
  rules: {
    "react-refresh/only-export-components": [
      "warn",
      { allowConstantExport: true },
    ],
    "import/no-unresolved": "error",
    "import/no-restricted-paths": [
      "error",
      {
        zones: [
          // enforce unidirectional codebase:
          {
            target: "./src/features",
            from: "./src/app",
          },

          {
            target: [
              "./src/components",
              "./src/hooks",
              "./src/lib",
              "./src/types",
              "./src/utils",
            ],
            from: ["./src/features", "./src/app"],
          },
        ],
      },
    ],
  },
};
