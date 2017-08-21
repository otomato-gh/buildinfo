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
        withEnv(["BUILD_NUMBER=${env.BUILD_NUMBER}"]){
            withCredentials([usernamePassword(credentialsId: '18887cf4-97dc-4c1f-90c6-6140f072cc6e', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]) {
              sh("git tag -a v0.${BUILD_NUMBER} -m 'Jenkins tested'")
              sh('git push https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/otomato-gh/buildinfo.git v0.${BUILD_NUMBER}')
            }
        }
      }
    }
  }
}
