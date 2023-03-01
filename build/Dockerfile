ARG IMAGE_NAME

FROM ${IMAGE_NAME} 

ARG PORT
ARG APP_NAME

ENV APP_NAME=${APP_NAME}
ENV PORT=${PORT}

WORKDIR /${APP_NAME}

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy

CMD air -c ./config/.air.toml

EXPOSE ${PORT}
