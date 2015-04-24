destroy:
	docker ps -a -q | awk '{print $1 }' | xargs docker stop && \
	docker ps -a -q | awk '{print $1 }' | xargs docker rm
