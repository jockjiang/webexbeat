#!/usr/bn/env groovy

pipeline{
    agent { label 'golang' }

    triggers { pollSCM('H * * * *')}
    stages {
        stage('Build'){
            steps{
                sh '/usr/bin/pwd'
            }
            steps{
                sh 'mage build'
            }
        }
    }
    post {
        always {
            echo 'happy learning Jenkins ...'
        }
    }
}