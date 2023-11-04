FROM ubuntu:22.04

RUN apt-get update \
  && apt-get install -y locales mysql-client \
  && locale-gen ja_JP.UTF-8 \
  && echo "export LANG=ja_JP.UTF-8" >> ~/.bashrc