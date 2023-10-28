pipeline {
  agent any
  stages {
    stage('Code Checkout') {
      steps {
        git(url: 'https://github.com/ppaprikaa/shorty', branch: 'dev')
      }
    }

    stage('List') {
      parallel {
        stage('List') {
          steps {
            sh 'ls -la'
          }
        }

        stage('Run Unit Testing') {
          steps {
            sh '/var/jenkins_home/go/bin/go test ./...'
          }
        }

      }
    }

  }
}