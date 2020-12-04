# prometheus
<B>prometheus monitoring demo</B>

#Docker command to run the Prometheus server
<code>
docker run --name prometheus -d -p 9090:9090 -v prometheus:/etc/prometheus prom/prometheus:latest
</code>
#Build and run the web server

git clone https://github.com/Anandanr/prometheus.git

docker build -t hid .

docker run --name greetpeople -d -p 8009:8009 hid:latest

Update the IP address of the server in the target.json file at https://github.com/Anandanr/prometheus/blob/master/prometheus/targets.json


Monitoring Tool
Design and implement a monitoring tool like Prometheus to collect and store the metrices
of the devices.
These devices can be of any type, they are IP connected and expose the metrics on a
common REST endpoint("/metric").
For this exercise, we want to monitor few metrices of IP Switches. And following are the
requirements

<B>#1. Scrape the metrics for every configured time interval(default shall be 5 seconds)</B>

 scrape_interval:     5s // in the prometheus.yml file

<B>#2. Store the data in a file system</B>

The first step is taking snapshots of Prometheus data, which can be done using Prometheus API. In order to use it, Prometheus API must first be enabled, using the CLI command:

./prometheus --storage.tsdb.path=data/ --web.enable-admin-api

The next step is to take the snapshot:

curl -XPOST http://{prometheus}:9090/api/v1/admin/tsdb/snapshot

Snapshots are stored in the data/snapshots directory.

Automating this command, running it as part of a daily or monthly procedure, and moving the snapshots to an external storage solution allows DevOps to retain historical data without having to continuously monitor the storage usage of Prometheus. 

The second step is to configure the required retention time period of time series or the disk size that DevOps maintain. This can be configured easily using the –storage.tsdb.retention.size  and –storage.tsdb.retention.time flags. 

<B>#3. Purge the data points older than x days(x is configurable)</B>

This can be achived by the same like as "2. Store the data in a file system" as follows:

$ curl -X POST \
    -g 'http://localhost:9090/api/v1/admin/tsdb/delete_series?match[]={foo="bar"}

<B>#4. Expose an endpoint to query the historical data of the devices</B>

We can see this in the prometheus dashboard.

<B>#5. The list of devices to be scraped shall come from a device config file(it should be a
configurable parameter)</B>

This is in the mounted volume and we can update or add the devices so it will updated in seconds with prometheus

scrape_configs:
   - job_name: 'node'
     file_sd_configs:
       - files:
          - 'targets.json' 
          
see https://github.com/Anandanr/prometheus/blob/master/prometheus/targets.json


<B>#6. Should scale for hundreds of metrics and thousands of devices</B>

Setting this architecture up is quite simple. A master server is deployed with “targets” that contain a list of slave Prometheus server URLs, like this:
<code>
	
scrape_configs:
	
      - job_name: federated_prometheus
      
        honor_labels: true
	
        metrics_path: /federate
	
        params:
	
          match[]:
	  
          - '{job="the-prom-job"}'
	  
        static_configs:
	
          - targets:
	  
            - prometheus-slave1:9090
	    
            - prometheus-slave2:9090
</code>
												
The match[] param in the configuration instructs Prometheus to accumulate and store all the slave metrics for a specific job. You can also set it as a regex: {__name__=~”^job:.*”}. This will collect metrics from several different jobs that match the expression definition. 

The Prometheus slaves should have the following configuration: 

<code>

global:

  external_labels:
  
    slave: 1
    
  relabel_configs:
  
  - source_labels: [_prometheus_slave]
  
    action: keep
    
    regex: slave1
    
 </code>
