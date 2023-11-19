# :metal:	Dyte Log-Ingestor-Service
- Log-Ingestor helps to monitor the application and pods/container on UI dashboard like kibana, where you can filter/query the logs(Production). 


## High Level Design:
![LocoTransaction](https://github-production-user-asset-6210df.s3.amazonaws.com/40069230/284015822-0b0a5d47-80e1-4bd0-a352-03634b288323.png)

## :computer: Tech Stack

* [Golang](https://go.dev/)
* [ElasticSearch](https://www.elastic.co/elasticsearch)
* [Kibana](https://www.elastic.co/kibana)
* [Filebeat](https://www.elastic.co/beats/filebeat)
* [Docker](https://www.docker.com/)


## :running_woman: Local Installation Guide :

1. Clone the repository by using Git Client:

         git clone https://github.com/palrohitg/log-ingestor

2. To Setup and Run Log-Ingestor:

        docker-compose up 

3. To Manually Generating logs :

        curl --location 'http://localhost:8080/api' \
        --header 'accept: application/json, text/plain, */*' \
        --header 'cache-control: no-cache' \
        --header 'Cookie: csrftoken=jH2qzTXI1DfSRW01HlHj0mJq49lqiMJ5BAGf6pZKhIfiz6YoLim14bD6V06tde7v'

4. Visit to Kibana URL Dashboard and Set the Index Pattern Mention
      
         http://0.0.0.0:5601/app/kibana#/home?_g=() 


## :computer: Logs of Docker Shell
![Kibana1](https://github.com/palrohitg/log-ingestor/assets/40069230/2569e9d3-6d0f-4f68-ad68-a39cb90c4fa4)
![Terminal2](https://github.com/palrohitg/log-ingestor/assets/40069230/c2c690f4-b4ba-48ef-a2e1-ff04a1bad954)


## :computer: Logs of Docker Shell
![Kibana](https://github.com/palrohitg/log-ingestor/assets/40069230/a33a9fbb-be2c-4e9a-8812-338e0f2bbc81)
![Terminal](https://github.com/palrohitg/log-ingestor/assets/40069230/b11e1d11-94be-493f-90e1-ae3d39e66557)

## 📜 LICENSE

[MIT](https://github.com/palrohitg/log-ingestor) 