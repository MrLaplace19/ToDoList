pipeline {
    agent any

    stages {
        stage('Format'){
            steps{
                sh 'test -z "$(gofmt -l .)"'
            }
        }
        stage('Environment') {
            steps {
                sh 'go version'
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