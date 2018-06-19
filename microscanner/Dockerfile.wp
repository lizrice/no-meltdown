FROM wordpress
COPY microscanner /microscanner
RUN chmod +x /microscanner
ARG token
RUN /microscanner --html ${token} > /ms-out.html 
