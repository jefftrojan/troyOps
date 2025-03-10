package kubernetes.admission

deny[msg] {
  input.request.kind.kind == "Deployment"
  not input.request.object.metadata.labels.app
  msg := "Deployment must have an 'app' label"
}

deny[msg] {
  input.request.kind.kind == "Deployment"
  not input.request.object.metadata.labels.environment
  msg := "Deployment must have an 'environment' label"
} 