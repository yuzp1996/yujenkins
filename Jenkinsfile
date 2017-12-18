pipeline {
    agent { docker 'python:3.5.1' }
    stages {
        stage('build') {
            steps {
                echo 'Start'
                sh 'python --version'
                echo 'End'
            }
        }
    }
}