# fly-image-updater
Keeps Docker images up to date in [Fly](https://fly.io)

## How it works
This tool queries the [Fly GraphQL API](https://api.fly.io/graphql) to get the  SHA digest of your image.
If it does not match the latest image available on its associated registry,
the tool updates the image to the latest version in Fly and
optionally notifies you. With the way that deployments work in Fly,
you can specify a Git location to clone and for the tool to work
from in order to update the images. In the current release, if you
are using the Git clone functionality, your Dockerfile tag must
be `latest`. I hope to address this in the future through a system
that integrates with Git.

This tool is intended to be run as a Docker
image on a schedule like a cron job or Systemd
timer. Or maybe even a Kubernetes job if you're 
feeling like it.

## Configuration
Done with YAML in `config.yaml`.
