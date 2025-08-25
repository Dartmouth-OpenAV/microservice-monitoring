FROM php:8.2-apache

RUN apt-get update && \
    apt-get install -y --no-install-recommends iputils-ping && \
    rm -rf /var/lib/apt/lists/*

COPY configs/000-default.conf /etc/apache2/sites-available

COPY source/ /var/www/html/

RUN a2enmod rewrite