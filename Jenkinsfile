pipeline {
    agent {
        docker {
            image 'golang'
        }
    }
    stages{
        stage('Build'){
            steps{
                sh 'echo "Compiling the code"'
                sh 'go biuld main.go'
                stash includes: 'main', name: 'goTest'
            }
        }
        stage('test'){
            steps{
                sh 'echo "Testing the go file"'
                sh 'go test'
            }
        }
        stage('deploy'){
            steps{
                sh 'echo "running the file"'
                unstash 'goTest'
                sh './main'
            }
        }
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
    }
}