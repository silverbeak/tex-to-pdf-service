FROM ubuntu

LABEL maintainer="Kristofer Jarl <kristofer@sootsafe.com>"

RUN apt-get update -y
RUN apt-get install -y latexmk texlive-latex3 texlive-science texlive-latex-extra texlive-fonts-recommended

EXPOSE 50051

RUN mkdir -p /temp

COPY tex-to-pdf-service /

ENTRYPOINT [ "/tex-to-pdf-service" ]