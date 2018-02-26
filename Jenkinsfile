pipeline {
    agent{
        docker {image 'golang'}
    }
    stages {
        
        stage('No-op') {
            steps {
                input("Could I yew")
                sh 'ls'
                sh 'wget "www.baidu.com"'
            }
        }
    }
    post {
        always {
            echo 'One way or another, I have finished'
            deleteDir() /* clean up our workspace */
        }
        success {
            echo 'I succeeeded! you are so good I like you '
        }
        unstable {
            echo 'I am unstable :/'
        }
  
        
        failure {
                    echo 'I failed :('

        mail to: 'zpyu@alauda.io',
             subject: "Failed Pipeline: ${currentBuild.fullDisplayName}",
             body: "Something is wrong with ${env.BUILD_URL}"
         }

        changed {
            echo 'Things were different before...'
        }
    }
}
