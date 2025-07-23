terraform {
  required_providers {
    seeweb = {
      source = "rh363/seeweb"
    }
  }
}

provider "seeweb" {
  ecs_apikey = "YFpwkKflXfkNRE7qyxL5FDUroccZya9Dlz4tzmMHR9UDi87leIOs68ANctNNxZwkCjZ3xAPNiQNksRkkPWOETWhfUf4bEcdwdtortyAde6Tmh7jiY5vp0E3c59QWZ9xj"
}

data "seeweb_regions" "all_regions" {

}

output "regions_output" {
  value = data.seeweb_regions.all_regions.regions
}