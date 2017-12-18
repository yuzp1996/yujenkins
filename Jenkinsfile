pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo 'start'
                sh 'echo "Hello World"'
                sh '''
                    echo "Multiline shell steps works too"
                    ls -lah
                '''
                echo 'end'
            }
        }
    }
}