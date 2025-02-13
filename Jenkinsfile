pipeline {
    agent any
    stages {
        stage('Checkout Code') {
            steps {
                checkout scm
            }
        }
    

       stage('Setup Go') {
            steps {
                sh 'go version'
            }
        }
    

        stage('Build') {
            steps {
                sh 'go build -o my_app main.go'
                archiveArtifacts artifacts: 'my_app', fingerprint: true
            }
        }
    

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Deploy (Simulation)') {
            steps {
                echo 'Deploying application...'
            }
        }
    }
}