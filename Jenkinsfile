pipeline{
    agent any
    stages{
        stages('Jenkins environment'){
            steps{
                sh 'go version'
            }
        }
        stage('CI started'){
            steps{
                sh 'echo CI for ToDoList started'
                sh 'pwd'
                sh 'ls -la'
            }
        }
    }
}