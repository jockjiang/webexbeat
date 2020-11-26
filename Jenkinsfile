#!/usr/bn/env groovy

pipeline{
    agent { label 'golang' }

    triggers { pollSCM('H * * * *')}
    stages {
        stage('Checking'){

            steps{
                sh '/usr/bin/pwd'
            }
        }
        stage('Build'){

            steps{
                sh 'make'
            }
        }
    }
    post {
        always {
            echo 'happy learning Jenkins ...'
        }
    }
}