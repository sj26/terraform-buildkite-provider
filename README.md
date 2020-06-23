# Terraform provider for Buildkite

Manage [Buildkite](https://www.buildkite.com) pipelines and other configuration via [Terraform](https://www.terraform.io).

Forked from [saymedia](https://github.com/saymedia/terraform-buildkite) via [yougroupteam](https://github.com/yougroupteam/terraform-buildkite). Both teams have done great work. It's been forked here to make some minor improvements and make binary releases available for automated download and install via [terraform-provider-buildkite-plugin](https://github.com/coyainsurance/terraform-provider-buildkite-plugin).

## Installation

**Recommended:** Download a [binary release](https://github.com/sj26/terraform-buildkite-provider/releases) into your local Terraform plugins directory (`~/.terraform.d/plugins/`) or use the [terraform-provider-buildkite-plugin](https://github.com/coyainsurance/terraform-provider-buildkite-plugin) to install it automatically when running on Buildkite.

Or install manually:

```bash
go get github.com/sj26/terraform-provider-buildkite
go install github.com/sj26/terraform-provider-buildkite
mkdir -p ~/.terraform.d/plugins
cp $GOPATH/bin/terraform-provider-buildkite ~/.terraform.d/plugins/
```

See the [Terraform docs on custom plugins](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins) for more options.

## Usage

```terraform
provider "buildkite" {
  # Get an API token from https://buildkite.com/user/api-access-tokens
  # Needs: read_pipelines, write_pipelines
  # Instead of embedding the API token in the .tf file,
  # it can also be passed via env variable BUILDKITE_API_TOKEN
  api_token    = "YOUR_API_TOKEN"
  # This is the part behind https://buildkite.com/, e.g. https://buildkite.com/some-org
  # Instead of embedding the org slug in the .tf file,
  # it can also be passed via env variable BUILDKITE_ORGANIZATION
  organization = "YOUR_ORG_SLUG"
}

resource "buildkite_pipeline" "terraform_test" {
  name       = "terraform-test"
  repository = "git@github.com:you/repo.git"

  step = [
    {
      type    = "script"
      name    = ":llama: Tests"
      command = "echo Hello world!"
    },
  ]
}
```

## Importing existing pipelines

You can import existing pipeline definitions by their slug:

```bash
terraform import buildkite_pipeline.my_name my-pipeline-slug
```

## Local development of this provider

To do local development you will most likely be working in a Github fork of the repository. After creating your fork
you can add it as a remote on your local repository in GOPATH:

* `cd $GOPATH/src/github.com/sj26/terraform-provider-buildkite`
* `git remote add mine git@github.com:yourname/terraform-provider-buildkite`
* `git checkout -b yourbranch`
* `git push -u mine yourbranch`

After this you should be able to `git push` to your fork, and eventually open a PR if you like.

You can build like this:

* `go install github.com/sj26/terraform-provider-buildkite`

This should produce a file at `$GOPATH/bin/terraform-provider-buildkite`. To use this with Terraform you'll need to move that binary to the [third-party plugins directory](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) to help Terraform find this file.

You can see debug output via `TF_LOG=DEBUG terraform plan`
