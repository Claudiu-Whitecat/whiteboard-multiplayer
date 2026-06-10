// Prettier 3+ ESM config
export default {
  // Common, readable code style
  printWidth: 100,
  tabWidth: 2,
  useTabs: false,
  semi: true,
  singleQuote: true,
  trailingComma: 'all',
  bracketSpacing: true,
  arrowParens: 'always',
  htmlWhitespaceSensitivity: 'ignore',
  vueIndentScriptAndStyle: true,
  endOfLine: 'lf',

  // Plugins — Tailwind must be last
  plugins: [
    'prettier-plugin-organize-imports',
    'prettier-plugin-packagejson',
    'prettier-plugin-tailwindcss',
  ],

  // Optional: keep these consistent across filetypes
  overrides: [
    {
      files: ['*.yml', '*.yaml'],
      options: { singleQuote: false }, // YAML prefers double quotes when needed
    },
    {
      files: ['*.md'],
      options: { proseWrap: 'always' },
    },
  ],
};
