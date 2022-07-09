# fly-image-updater
Keeps Docker images up to date in [Fly](https://fly.io)

## How it works
This tool queries the [Fly GraphQL API](https://api.fly.io/graphql) to get the  SHA digest of your image.
If it does not match the latest image available on its associated registry,
the tool updates the image to the latest version in Fly and
optionally notifies you. With the way that deployments work in Fly,
you can specify a Git location to clone and for the tool to work
from in order to update the images.  If you're using the Git method, 
then you will need to specify version numbers with a tool like Dependabot
or Renovate on your Git repository as it will only update if there is a change 
to the version specified there.

This tool is intended to be run as a Docker
image on a schedule like a cron job or Systemd
timer. Or maybe even a Kubernetes job if you're 
feeling like it.

## Configuration
Done with YAML in `config.yaml`.

## Componenets
- YAML Config
- App loop
    - Git cloning and parsing
    - OR
    - Image parsing with GraphQL Fly API
        - Skopeo
    - Update image with Fly remote builder
    - Notification
- KV Storage for Repo Information
