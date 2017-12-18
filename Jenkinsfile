pipeline {
    agent any
    stages {
        stage('No-op') {
            steps {
                sh 'ls'
                sh 'go run go.test'
            }
        }
    }
    post {
        always {
            echo 'One way or another, I have finished'
            deleteDir() /* clean up our workspace */
        }
        success {
            echo 'I succeeeded!'
        }
        unstable {
            echo 'I am unstable :/'
        }
  
        
        failure {
                    echo 'I failed :('

        mail to: 'team@example.com',
             subject: "Failed Pipeline: ${currentBuild.fullDisplayName}",
             body: "Something is wrong with ${env.BUILD_URL}"
         }

        changed {
            echo 'Things were different before...'
        }
    }
}