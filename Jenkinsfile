pipeline {
    agent any
    
    environment {
        // Define your Docker image name and tag
        IMAGE_NAME = "go-gitops"
        IMAGE_TAG = "${BUILD_NUMBER}"
    }
    
    stages {
        stage('Checkout') {
            steps {
                // Checkout source code from repository
                checkout scm
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    // Build Docker image
                    def image = docker.build("${IMAGE_NAME}:${IMAGE_TAG}")
                    
                    // Optional: Tag with 'latest' as well
                    sh "docker tag ${IMAGE_NAME}:${IMAGE_TAG} ${IMAGE_NAME}:latest"
                }
            }
        }
        
        stage('List Images') {
            steps {
                // Verify the image was built successfully
                sh 'docker images | grep ${IMAGE_NAME}'
            }
        }
    }
    
    post {
        always {
            // Clean up workspace
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