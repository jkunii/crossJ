FROM scratch

COPY crossJ /
COPY banner.txt /
COPY .env /
COPY public /public

EXPOSE 1323

CMD ["/crossJ"]