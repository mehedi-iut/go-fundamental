pipeline {
    agent{
        label 'ats-test-cefalolab'
    } 
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage("Test"){
            steps {
                sh 'hostname'
                // get ip address of the node
                sh "ip addr"
            }
        }
        // stage('Backend') {
        //     steps {
        //         dir('backend') {
        //             sh 'echo "Building backend"'
        //             // Add backend-specific build steps here
        //             // For example:
        //             sh 'npm install'
        //             sh 'npm run build'
        //         }
        //     }
        // }
        // stage('Frontend Internal') {
        //     steps {
        //         dir('frontend/apps/frontend-internal') {
        //             sh 'echo "Building frontend-internal"'
        //             // Add frontend-internal-specific build steps here
        //             // For example:
        //             sh 'npm install'
        //             sh 'npm run build'
        //         }
        //     }
        // }
        // stage('Frontend Public') {
        //     steps {
        //         dir('frontend/apps/frontend-public') {
        //             sh 'echo "Building frontend-public"'
        //             // Add frontend-public-specific build steps here
        //             // For example:
        //             sh 'npm install'
        //             sh 'npm run build'
        //         }
        //     }
        // }
    }
    post {
        success {
            echo 'Build successful! The changes look good.'
        }
        failure {
            echo 'Build failed. Please check the logs and fix the issues.'
        }
    }
}