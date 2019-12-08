pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'cd cinemahall && go build main.go'
                sh 'cd movies && go build main.go'
                sh 'cd reservation && go build main.go'
                sh 'cd show && go build main.go'
                sh 'cd users && go build main.go'
            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'echo run tests...'
                sh 'cd cinemahall/cinemahall && go test -cover'
                sh 'cd movies/movies && go test -cover'
                sh 'cd reservation/reservation && go test -cover'
                sh 'cd show/show && go test -cover'
                sh 'cd users/users && go test -cover'
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'cd cinemahall'
                sh 'export GO111MODULE=on'
                sh 'golangci-lint run --enable-all'  //--deadline 20m --enable-all--disable-all -E errcheck
            }
        }
        stage('Build Docker Image') {
            agent any
            steps {
                sh "docker-build-and-push -b ${BRANCH_NAME} -f cinemahall/dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -f movies/dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -f reservation/dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -f show/dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -f users/dockerfile"
            }
        }
    }
}