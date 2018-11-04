# Project Vision

Ambari provides a REST api to manage Hadoop Clusters, in particular for Hortonworks Data Platform
and Hortonworks Data Flow.

But similar to other famous APIs, like AWS, GCP and Azure and also other software components, it does
not provide any command line utility, very useful for automation tasks of devops atcivities.

# Inception

The goal is to create a command line util with intuitive and well thought terminal applicationl,
but for the moment we will re-implement and identical behaviour of the ambari-shell.

# Challenges

One of the main challenge of implementation this terminal application is to handle complex Ambari Blueprints,
and hopefully the command line app will easy that.

# Goals

The main goals are:

* well-designed commnand line utility in terms of User Experience (UX)
* turn Hadoop cluster planning easier
* make the full unattended installation a reality