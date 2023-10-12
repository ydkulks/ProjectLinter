# ProjectLinter

CLI tool that monitors and evaluates the file structure of a directory and its
subdirectories to enforce a specific file structure pattern for a particular
programming language

## Project structure

### Javascript and Typescript (node.js)

```txt
my-node-app/
  ├── node_modules/        (Dependencies installed via npm)
  ├── public/              (Static files like HTML, CSS, client-side JavaScript)
  ├── src/                 (Source code)
  │   ├── controllers/     (Route handlers)
  │   ├── models/          (Database models)
  │   ├── routes/          (Express.js route definitions)
  │   ├── views/           (Template files, if using a template engine)
  ├── config/              (Configuration files)
  │   ├── database.js      (Database configuration)
  │   ├── routes.js        (Route configuration)
  │   ├── environment.js   (Environment-specific settings)
  ├── tests/               (Test files)
  ├── package.json         (Node.js package configuration)
  ├── package-lock.json    (Dependency lock file)
  ├── .gitignore           (Git ignore rules)
  ├── .env                 (Environment variables, not in version control)
  ├── app.js               (Main application file)
  ├── README.md            (Documentation)

```

```text
my-typescript-project/
  ├── node_modules/        (Dependencies installed via npm or yarn)
  ├── src/                 (Source code)
  │   ├── components/      (Reusable UI components)
  │   ├── containers/      (Higher-level components that manage state)
  │   ├── services/        (API services, utilities, and helpers)
  │   ├── styles/          (CSS or SCSS files)
  │   ├── assets/          (Static assets like images)
  ├── public/              (Publicly accessible files, e.g., index.html)
  ├── tests/               (Test files)
  ├── config/              (Configuration files)
  │   ├── env/             (Environment-specific configuration)
  │   ├── constants.ts     (Application constants)
  │   ├── routes.ts        (Route definitions)
  ├── typings/             (Custom type definitions, if needed)
  ├── .gitignore           (Git ignore rules)
  ├── tsconfig.json        (TypeScript configuration)
  ├── package.json         (Node.js package configuration)
  ├── package-lock.json    (Dependency lock file)
  ├── README.md            (Documentation)

```

### Go (Web app)

```txt
my-go-web-app/
  ├── assets/            (Static assets like CSS, JavaScript, and images)
  ├── templates/         (HTML templates if not using a frontend framework)
  ├── cmd/               (Application entry points)
  │   ├── main.go        (Main application entry point)
  ├── internal/          (Internal application packages)
  │   ├── handlers/     (HTTP request handlers)
  │   ├── models/       (Data models)
  │   ├── middleware/   (HTTP middleware)
  │   ├── config/       (Application configuration)
  ├── pkg/               (Reusable packages)
  ├── vendor/            (Vendor directory for dependencies - if not using Go Modules)
  ├── go.mod             (Go module file)
  ├── go.sum             (Dependency checksum file)
  ├── .gitignore         (Git ignore rules)
  ├── README.md          (Documentation)

```

### Go (CLI)

```txt
my-go-cli/
  ├── cmd/               (Application entry points)
  │   ├── mycli/         (CLI application source code)
  │   │   ├── main.go    (Main CLI application entry point)
  ├── internal/          (Internal application packages)
  ├── pkg/               (Reusable packages)
  ├── vendor/            (Vendor directory for dependencies - if not using Go Modules)
  ├── go.mod             (Go module file)
  ├── go.sum             (Dependency checksum file)
  ├── .gitignore         (Git ignore rules)
  ├── README.md          (Documentation)

```

### PHP (Web app)

```txt
my-php-web-app/
  ├── app/
  │   ├── Controllers/     (Controller classes)
  │   ├── Models/          (Model classes)
  │   ├── Views/           (View templates)
  ├── config/              (Application configuration)
  │   ├── database.php     (Database configuration)
  ├── public/              (Publicly accessible files, e.g., index.php, assets)
  ├── resources/           (Non-public resources like language files or raw assets)
  ├── routes/              (Route definitions)
  ├── tests/               (Test files)
  ├── vendor/              (Composer dependencies)
  ├── .env                 (Environment-specific settings)
  ├── .gitignore           (Git ignore rules)
  ├── composer.json        (Composer configuration)
  ├── composer.lock        (Composer lock file)
  ├── README.md            (Documentation)

```

### Lua (Neovim plugin)

```txt
my-neovim-plugin/
  ├── plugin/            (Your plugin's main Lua file)
  ├── lua/               (Additional Lua code, organized by functionality)
  │   ├── my_plugin.lua  (Your plugin's main Lua code)
  │   ├── commands.lua   (Neovim commands)
  │   ├── mappings.lua   (Key mappings)
  │   ├── utils.lua      (Utility functions)
  ├── doc/               (Documentation for your plugin)
  ├── tests/             (Test files for your plugin)
  ├── README.md          (Plugin documentation)

```

## Todo

- [x] Traverse through a specifed directory
- [x] Ignore dot files
- [x] Selection option for different kind of projects
- [x] Restructure this project based on common convention
- [ ] Modules for each kind of project
  - [x] JavaScript (node.js)
  - [x] Go (Web app)
  - [x] Go (CLI)
  - [x] PHP (Web app)
  - [x] Lua (Neovim plugin)
  - [ ] Other
- [x] Sub directory file type validation
