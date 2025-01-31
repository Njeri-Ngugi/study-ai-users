data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "./cmd/migration",
  ]
}

variable "url" {
  type = string
  default = getenv("POSTGRES_DSN")
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15/dev"
  url = var.url
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}