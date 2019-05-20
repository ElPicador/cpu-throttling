FROM alpine

ADD cpu-throttling /

CMD [ "./cpu-throttling" ]
