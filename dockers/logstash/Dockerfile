FROM logstash:latest
RUN logstash-plugin install logstash-input-syslog
RUN logstash-plugin install logstash-input-gelf
RUN logstash-plugin install logstash-codec-json
RUN logstash-plugin install logstash-filter-json

# Add your logstash plugins setup here
# Example: RUN logstash-plugin install logstash-filter-json 
