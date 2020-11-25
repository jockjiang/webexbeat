#!/usr/bn/env groovy

pipeline{
    agent any

    triggers { pollSCM('H * * * *')}
    stages {
        stage('Build'){
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