pipeline {
    agent { label 'master' }

    options {
        buildDiscarder(logRotator(numToKeepStr: '10'))
        timeout(time: 30, unit: 'MINUTES')
        timestamps()
    }

    parameters {
		string(name: 'PROJECT_NAME', defaultValue: 'campaign-service', description: '')
		string(name: 'DOCKER_IMG_NAME', defaultValue: 'platform/campaign-service', description: '')
		string(name: 'ECR_URL', defaultValue: '455516961477.dkr.ecr.ap-south-1.amazonaws.com/', description: '')
		choice(name: 'RELEASE_MODE', choices: ['major', 'minor', 'patch'], description: 'Pick one.')
		gitParameter branchFilter: 'origin/(.*)', defaultValue: 'master', name: 'BRANCH', type: 'PT_BRANCH'

	}

    stages {
       stage ('Checkout') {
            steps {
                script {
                    if ("${params.BRANCH_NAME}" == 'master' || "${params.BRANCH_NAME}".startsWith('release-')) {
                        echo 'Development build not allowed on master or release branch!'
                        sh 'exit 1'
                    }
                }
                checkout scm
            }
        }

        stage('Auto tagging') {
          steps {
            script {
              FINALTAG = sh (script: "bash /opt/jenkins-tag/tag.sh ${params.RELEASE_MODE} ${params.BRANCH}", returnStdout: true).trim()
              echo "Tag is : ${FINALTAG}"
            }
            echo "Returned Tag is : ${FINALTAG}"
          }
        }



		stage('Build') {
            steps {
                sh "cp -p /var/lib/jenkins/tmp/* ./;  docker build -t ${params.DOCKER_IMG_NAME} -f ./Dockerfile ."
            }
        }

        stage('Publish') {
            steps {
                 script {
                    sh "docker tag ${params.DOCKER_IMG_NAME}:latest ${params.ECR_URL}${params.DOCKER_IMG_NAME}:${FINALTAG}"
		            docker.withRegistry("https://${params.ECR_URL}", "ecr:ap-south-1:AWSECR") {
                        sh "docker push ${params.ECR_URL}${params.DOCKER_IMG_NAME}:${FINALTAG}"
                    }
                }
            }
        }

    }

	post {
        success {
            writeFile file: "output/tag.txt", text:"tag=${FINALTAG}"
            archiveArtifacts artifacts: 'output/*.txt'
        }
	}
}
