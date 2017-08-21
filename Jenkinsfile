pipeline {
  agent {
    node {
      label 'golang'
    }
    
  }
  stages {
    stage('Build') {
      steps {
        sh 'go test'
      }
    }
    stage('Push') {
      steps {
        git 'https://github.com/otomato-gh/buildinfo.git'
        sh 'git tag 0.${env.BUILD_NUMBER} && git push origin 0.${env.BUILD_NUMBER}'
      }
    }
  }
}