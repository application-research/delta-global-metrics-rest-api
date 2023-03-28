gen --sqltype=postgres --makefile \
--module=github.com/application-research/delta-metrics-rest \
--connstr "host=kEOpaaIn7ZRDLL9IGrKYOH2MpUEavWWg@dpg-cfto8d9a6gdotcfptsrg-a.oregon-postgres.render.com user=deltadb_metrics_user password=kEOpaaIn7ZRDLL9IGrKYOH2MpUEavWWg dbname=deltadb_metrics port=5432" \
--database estuary \
--rest \
--gorm \
--overwrite \
--generate-dao \
--swagger_version=1.0 \
--swagger_path=/ \
--swagger_tos= \
--swagger_contact_name=Me \
--swagger_contact_url=http://me.com/terms.html \
--swagger_contact_email=alvin@protocol.ai \
--json-fmt=lower_camel \
--guregu \
--run-gofmt \
--no-overwrite \
--out .

#--server \