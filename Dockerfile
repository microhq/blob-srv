FROM alpine
ADD blob-srv /blob-srv
ENTRYPOINT [ "/blob-srv" ]
