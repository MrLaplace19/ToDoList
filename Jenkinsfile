pipeline {
    agent any

    stages {
        
        stage('Environment') {
            steps {
                sh 'go version'
            }
        }

        stage('Format'){
            steps{
                sh 'test -z "$(gofmt -l .)"'
            }
        }

        stage('Vet') {
            steps {
                sh 'go vet ./...'
            }
        }

        stage('Test') {
            steps {
                sh 'go test -cover ./...'
            }
        }

        stage('Build') {
            steps {
                sh 'go build ./...'
            }
        }
    }
}