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
                bat 'echo "Compiling the code"'
                bat 'go build main.go'
                stash includes: 'main.exe', name: 'goTest'
            }
        }
        stage('test'){
            steps{
                bat 'echo "Testing the go file"'
                bat 'go test'
            }
        }
        stage('deploy'){
            steps{
                bat 'echo "running the file"'
                unstash 'goTest'
                bat 'main.exe'
            }
        }
    }
}
