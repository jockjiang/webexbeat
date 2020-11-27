#!/usr/bn/env groovy

pipeline{
    agent { label 'golang' }

    triggers { pollSCM('H * * * *')}
    stages {
        stage('Verify Branch'){
            steps{
                echo $GIT_BRANCh
            }
        }
        stage('Checking'){
            steps{
                sh '/usr/bin/pwd'
            }
        }
        stage('Build'){
            steps{
                sh 'export make'
            }
        }
    }
    post {
        always {
            echo 'happy learning Jenkins ...'
        }
    }
}