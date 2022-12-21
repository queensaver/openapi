UID := $(shell id -u)

all: openapi protos grpc

wd := /mnt/home/wogri/queensaver/openapi

openapi: 
	docker run --rm -v "${wd}:/local" --user ${UID} openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g html2 -o /local/html
	docker run --rm -v "${wd}:/local" --user ${UID} openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g protobuf-schema -o /local/proto --package-name=openapi
	docker run --rm -v "${wd}:/local" --user ${UID} openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g typescript-angular -o /local/angular
	docker run --rm -v "${wd}:/local" --user ${UID} openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g go-server -o /local/golang --git-user-id queensaver --git-repo-id openapi/golang
	# TODO: this is generated as root - we can't do that.
	rm golang/go/api_default.go
	rm golang/main.go

protos:
	protoc -I proto --proto_path proto \
		--go_out=paths=source_relative:golang/proto \
		--go_opt=Mmodels/hive.proto.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/bbox_bhive.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/bhive.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/bbox.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/hive.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/stand.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/alerts.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/bbox_bhive_inner.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/user.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/weight.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/temperature.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/get_temperature_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/varroa_scan.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/varroa_scan_metadata.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/varroa_scan_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/get_stands_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/post_stands_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/put_stand_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/get_bbox_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/put_bbox_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/post_bbox_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/get_hives_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/put_hive_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/post_hives_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/bbox_config_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/registration_id.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/login_post_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/authenticate_post_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/generic_post_response.proto=github.com/queensaver/openapi/golang/proto/models \
		\
		proto/models/hive.proto \
		proto/models/bbox_bhive.proto \
		proto/models/bbox.proto \
		proto/models/bhive.proto \
		proto/models/hive.proto \
		proto/models/stand.proto \
		proto/models/alerts.proto \
		proto/models/bbox_bhive_inner.proto \
		proto/models/user.proto \
		proto/models/weight.proto \
		proto/models/temperature.proto \
		proto/models/get_temperature_response.proto \
		proto/models/varroa_scan.proto \
		proto/models/varroa_scan_metadata.proto \
		proto/models/varroa_scan_response.proto \
		proto/models/get_stands_response.proto \
		proto/models/post_stands_response.proto \
		proto/models/put_stand_response.proto \
		proto/models/get_bbox_response.proto \
		proto/models/put_bbox_response.proto \
		proto/models/post_bbox_response.proto \
		proto/models/get_hives_response.proto \
		proto/models/put_hive_response.proto \
		proto/models/post_hives_response.proto \
		proto/models/bbox_config_response.proto \
		proto/models/registration_id.proto \
		proto/models/login_post_response.proto \
		proto/models/authenticate_post_response.proto \
		proto/models/generic_post_response.proto

grpc:
	~/bin/protoc --proto_path proto \
		--go_out=paths=source_relative:golang/proto \
		--go_opt=paths=source_relative \
		--go_grpc_out=paths=source_relative:golang/proto \
		--go_grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go_grpc=/home/wogri/go/bin/protoc-gen-go-grpc \
		\
		--go_opt=Mservices/default_service.proto=github.com/queensaver/openapi/golang/proto/services \
		--go_opt=Mmodels/hive.proto.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/bbox_bhive.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/bbox.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/bhive.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/hive.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/stand.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/alerts.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/bbox_bhive_inner.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/user.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/weight.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/temperature.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/get_temperature_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/varroa_scan.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/varroa_scan_metadata.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/varroa_scan_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/get_stands_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/post_stands_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/put_stand_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/get_bbox_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/put_bbox_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/post_bbox_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/get_hives_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/put_hive_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/post_hives_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/bbox_config_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/registration_id.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/login_post_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/authenticate_post_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_opt=Mmodels/generic_post_response.proto=github.com/queensaver/openapi/golang/proto/models \
		\
		--go_grpc_opt=Mservices/default_service.proto=github.com/queensaver/openapi/golang/proto/services \
		--go_grpc_opt=Mmodels/hive.proto.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/bbox_bhive.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/bbox.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/bhive.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/hive.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/stand.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/alerts.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/bbox_bhive_inner.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/user.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/weight.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/temperature.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/get_temperature_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/varroa_scan.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/varroa_scan_metadata.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/varroa_scan_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/get_stands_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/post_stands_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/put_stand_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/get_bbox_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/put_bbox_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/post_bbox_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/get_hives_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/put_hive_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/post_hives_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/bbox_config_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/registration_id.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/login_post_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/authenticate_post_response.proto=github.com/queensaver/openapi/golang/proto/models \
		--go_grpc_opt=Mmodels/generic_post_response.proto=github.com/queensaver/openapi/golang/proto/models \
		\
		proto/services/default_service.proto

update_and_push: all
	git commit -am 'openapi updates'
	git push
