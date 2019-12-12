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
                //sh 'cd reservation/reservation && go test -cover'
                sh 'cd show/show && go test -cover'
                sh 'cd users/users && go test -cover'
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                //--deadline 20m --enable-all; --disable-all -E errcheck
                sh 'cd cinemahall && golangci-lint run --enable-all --skip-dirs proto -D wsl -D lll' 
                sh 'cd movies && golangci-lint run --enable-all --skip-dirs proto -D wsl -D lll -D gosimple' 
                sh 'cd reservation && golangci-lint run --enable-all --skip-dirs proto -D wsl -D lll -D funlen -D unparam'
                sh 'cd show && golangci-lint run --enable-all --skip-dirs proto -D wsl -D lll -D golint'
                sh 'cd users && golangci-lint run --enable-all --skip-dirs proto -D wsl -D lll'
            }
        }
        stage('Build Docker Image') {
            agent any
            steps {
                sh "docker-build-and-push -b ${BRANCH_NAME} -s cinemahall -f cinemahall/dockerfile ."
                sh "docker-build-and-push -b ${BRANCH_NAME} -s movies -f movies/dockerfile ."
                sh "docker-build-and-push -b ${BRANCH_NAME} -s reservation -f reservation/dockerfile ."
                sh "docker-build-and-push -b ${BRANCH_NAME} -s show -f show/dockerfile ."
                sh "docker-build-and-push -b ${BRANCH_NAME} -s users -f users/dockerfile ."
            }
        }
    }
}