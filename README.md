# huggo

Deployment tool for Hugo

## Installation

The following installation assumes Golang is installed on your remote server.

  1. Install Hugo on your remote server inside `$HOME/bin`: <https://github.com/gohugoio/hugo/releases>
  2. Create a bare git repo: `git init --bare /path/to/example.com.git`
  3. Compile this tool into your hooks directory: `go build -o /path/to/example.com.git/hooks/post-receive huggo.go`
  4. Add the remote to your local git setup: `git remote add deploy username@host:/path/to/example.com.git`
  5. When you're ready to deploy the master branch (swap out branch as necessary), run `git push deploy master`. Don't forget to push to any other remotes (e.g., GitHub) to ensure your source code is centralized!

If you're using NearlyFreeSpeech, `/path/to/example.com.git` can become `~/example.com.git` and the tilde expansion of `~` will yield a final path of `/home/private/example.com.git`. This even applies to adding the remote!
