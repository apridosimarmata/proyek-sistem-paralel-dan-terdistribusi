FROM python:3.7.9

WORKDIR /ui-service

COPY . .

RUN pip install -r requirements.txt

#RUN apt-get install -y python-pip

#RUN apt-get install -y libmysqlclient-dev

CMD [ "python3", "app.py"]