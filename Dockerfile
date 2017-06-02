FROM scratch
ADD urlService /
ENV PORT 8080
EXPOSE 8080
CMD ["/urlService"]