FROM scratch
COPY orderbish /
EXPOSE 8080
CMD ["./orderbish"]
