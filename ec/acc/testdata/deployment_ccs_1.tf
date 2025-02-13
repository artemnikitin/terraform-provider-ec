data "ec_stack" "latest" {
  version_regex = "latest"
  region        = "%s"
}

resource "ec_deployment" "ccs" {
  name                   = "%s"
  region                 = "%s"
  version                = data.ec_stack.latest.version
  deployment_template_id = "%s"

  elasticsearch {
    dynamic "remote_cluster" {
      for_each = ec_deployment.source_ccs
      content {
        deployment_id = remote_cluster.value.id
        alias         = remote_cluster.value.name
      }
    }
  }
}

resource "ec_deployment" "source_ccs" {
  count                  = 3
  name                   = "%s-${count.index}"
  region                 = "%s"
  version                = data.ec_stack.latest.version
  deployment_template_id = "%s"

  elasticsearch {
    topology {
      id         = "hot_content"
      zone_count = 1
      size       = "1g"
    }
  }
}
