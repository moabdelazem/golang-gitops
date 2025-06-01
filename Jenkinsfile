pipeline {
    agent any
    
    environment {
        IMAGE_NAME = "go-gitops"
        IMAGE_TAG = "${BUILD_NUMBER}"
    }
    
    stages {
        stage('Verify Docker') {
            steps {
                script {
                    sh 'docker --version'
                    sh 'docker info'
                }
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    sh """
                        docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .
                        docker tag ${IMAGE_NAME}:${IMAGE_TAG} ${IMAGE_NAME}:latest
                    """
                }
            }
        }
        
        stage('List Images') {
            steps {
                sh "docker images | grep ${IMAGE_NAME}"
            }
        }
    }
    
    post {
        always {
            cleanWs()
        }
        success {
            echo 'Docker image built successfully!'
        }
        failure {
            echo 'Docker image build failed!'
        }
    }
}