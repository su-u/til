FROM ubuntu:22.04

RUN apt-get update \
  && apt-get install -y locales percona-toolkit mysql-clinet \
  && locale-gen ja_JP.UTF-8 \
  && echo "export LANG=ja_JP.UTF-8" >> ~/.bashrc