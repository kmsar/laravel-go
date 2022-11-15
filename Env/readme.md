
## Usage

Add your application configuration to your `.env` file in the root of your project:

```shell
S3_BUCKET=YOURS3BUCKET
SECRET_KEY=YOURSECRETKEYGOESHERE
```

Then in your Go app you can do something like

```go
package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env-sample file")
  }

  s3Bucket := os.Getenv("S3_BUCKET")
  secretKey := os.Getenv("SECRET_KEY")

  // now do something with s3 or whatever
}
```

If you're even lazier than that, you can just take advantage of the autoload package which will read in `.env` on import

```go
import _ "github.com/joho/godotenv/autoload"
```

While `.env` in the project root is the default, you don't have to be constrained, both examples below are 100% legit

```go
_ = godotenv.Load("somerandomfile")
_ = godotenv.Load("filenumberone.env", "filenumbertwo.env")
```

If you want to be really fancy with your env file you can do comments and exports (below is a valid env file)

```shell
# I am a comment and that is OK
SOME_VAR=someval
FOO=BAR # comments at line end are OK too
export BAR=BAZ
```

Or finally you can do YAML(ish) style

```yaml
FOO: bar
BAR: baz
```

as a final aside, if you don't want godotenv munging your env you can just get a map back instead

```go
var myEnv map[string]string
myEnv, err := godotenv.Read()

s3Bucket := myEnv["S3_BUCKET"]
```

... or from an `io.Reader` instead of a local file

```go
reader := getRemoteFile()
myEnv, err := godotenv.Parse(reader)
```

... or from a `string` if you so desire

```go
content := getRemoteFileContent()
myEnv, err := godotenv.Unmarshal(content)
```

### Precedence & Conventions

Existing envs take precedence of envs that are loaded later.

The [convention](https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use)
for managing multiple environments (i.e. development, test, production)
is to create an env named `{YOURAPP}_ENV` and load envs in this order:

```go
env := os.Getenv("FOO_ENV")
if "" == env {
  env = "development"
}

godotenv.Load(".env-sample." + env + ".local")
if "test" != env {
  godotenv.Load(".env-sample.local")
}
godotenv.Load(".env-sample." + env)
godotenv.Load() // The Original .env-sample
```

If you need to, you can also use `godotenv.Overload()` to defy this convention
and overwrite existing envs instead of only supplanting them. Use with caution.

### Command Mode

Assuming you've installed the command as above and you've got `$GOPATH/bin` in your `$PATH`

```
godotenv -f /some/path/to/.env some_command with some args
```

If you don't specify `-f` it will fall back on the default of loading `.env` in `PWD`

### Writing Env Files

Godotenv can also write a map representing the environment to a correctly-formatted and escaped file

```go
env, err := godotenv.Unmarshal("KEY=value")
err := godotenv.Write(env, "./.env-sample")
```

... or to a string

```go
env, err := godotenv.Unmarshal("KEY=value")
content, err := godotenv.Marshal(env)
```

