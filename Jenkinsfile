pipeline {
  agent any
  environment {
    DOCKER_IMAGE = "github.com/surdiana/primasys:latest"
    REGISTRY_CREDENTIALS = 'dockerhub-credentials'
    HELM_RELEASE = "simple-goserver"
    KUBE_CONFIG = credentials('kubeconfig')
  }
  stages {
    stage('Build & Push Docker Image') {
      steps {
        script {
          docker.withRegistry('', REGISTRY_CREDENTIALS) {
            def app = docker.build(DOCKER_IMAGE)
            app.push()
          }
        }
      }
    }
    stage('Deploy to Kubernetes with Helm') {
      steps {
        withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG')]) {
          sh '''
            helm upgrade --install $HELM_RELEASE ./helm-chart \
              --set image.repository=YOUR_DOCKERHUB_USERNAME/simple-goserver \
              --set image.tag=latest
          '''
        }
      }
    }
  }
}