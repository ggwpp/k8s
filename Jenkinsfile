node() {

    checkout scm
    
    GIT_COMMIT = git_commit = sh(returnStdout: true, script: 'git rev-parse HEAD').trim().take(7)
    DOCKER_IMAGE = "asia.gcr.io/workshop-mfec/workshop3:${GIT_COMMIT}"
    GCLOUD_CRED  = "mfec-workshop"

    stage('Build Image'){
        sh("docker build -f Dockerfile.multi-stg -t ${DOCKER_IMAGE} .")
    }

    stage('Push Image'){
        withCredentials([file(credentialsId: GCLOUD_CRED, variable: 'KEY')]) {
            sh "gcloud auth activate-service-account --key-file=${KEY}"
            sh("docker push ${DOCKER_IMAGE} ")
        }
    }

    stage('Deploy'){
        sh("kubectl set image deployment workshop3 workshop3=${DOCKER_IMAGE}")
        sh("kubectl rollout status deployment workshop3")
    }
}