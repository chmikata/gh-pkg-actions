FROM chmikata/gh-pkg-cli:0.4.0
COPY entrypoint.sh /app
ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["--help"]
