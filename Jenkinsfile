pipeline {
    agent{
        none
    } 

    stages {
        stage('No-op') {
            steps {
                sh 'ls'
                sh 'echo "python pythonfile.py"'
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