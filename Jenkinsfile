pipeline {
    agent any
    stages{
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build'){
            steps{
                sh 'echo "Compiling the code"'
                sh 'go build main.go'
                stash includes: 'main.exe', name: 'goTest'
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
                sh 'main.exe'
            }
        }
    }
}
