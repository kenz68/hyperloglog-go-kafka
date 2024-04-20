
# Data Engineer Challenge

    The goal of the challenge is to have a tool that is able
    to stream data from [kafka](http://kafka.apache.org/) and count unique things within this data.
    The simplest use case is that we want to calculate unique users per minute, day, week, month, year. 
    For a very first version, bussiness wants us to provide just the unique users per minute.

    - The data consists of (Log)-Frames of JSON data that we streamd into/from apache kafka.
    - Each frame has a timestamp property which is unixtime, the name of the property is `ts`.
    - Each frame has a user id property calles `uid`.
    - You can assume that 99.9% of the frames arrive with maximum latency of 5 seconds.
    - You want to display the results as soon as possible.
    - The results should be forwarded to a new kafka topic (again as json.) choose a suitable structure.
    - For an advanced solution you should assume that you can *not* gurantee that events are always strictly ordered.

## Requirements:
    - provide a Readme that shows how to build and run the code on Linux or OS X
    - write a report: what did you do? what was the reasons you did it like that?
    - measure at least one performance metric (e.g. frames per second)
    - document your approach on how you decide **when** to output the data
    - document the estimated error in counting
    - it should be possible to ingest historical data. e.g. from the last 2 years.

## sample data

    For a quick start you can use the sample data provided at [here](https://tda-public.s3.eu-central-1.amazonaws.com/hire-challenge/stream.jsonl.gz):
    it was generated using `./data-generator -c 1000000 -o stream.json -r 1000 -n 100000`
    `cat stream.json | jq .uid -r | gsort --parallel=4 | guniq | wc -l` yields `99993` unique ids.
    you can just input this file to you kafka topic once by somthing like
    `gzcat stream.gz | kafka-console-produce --broker-list localhost:9092 --topic mytopic`
    You can also generate your own data by using the data-generater tool. it allows you to introdude random-ness if you have a more advanced solution and want to test it. it requires you insall Dlang compiler and D's package manager. then execute `dub build` to build the binary. binaries for os x and linux in the bin directory

## suggested steps:

##### basic solution
1. install kafka
2. create a topic
3. use the kafka producer from kafka itself to send our test data to your topic
4. create a small app that reads this data from kafka and prints it to stdout
5. find a suitable data structure for counting and implement a simple counting mechanism, 
   output the results to stdout
##### advance solution
6. benchmark
7. Output to a new Kafka Topic instead of stdout
8. try to measure performance and optmize
9. write about how you could scale
10. only now think about the edge cases, options and other things

## Evaluation Criteria & Expectations

- ability to break down *business* requirements into simple prototype code
- clean project setup and documentation
- research and use of suitable data structure for a specific use case. explain which and why.
- ability to wire performant code to handle streaming data. measure and document _how_ fast your solution is.
- Understanding how to benchmark and analyze performance bottlenecks. what tools did you use?
- awareness of the mechanism and costs of serialization. Explain (and ideally prove!) why json is an ideal format here or why not and suggest a better solution.
- scalability: _explain_ how you _would_ scale your approach


## Bonus Questions / Challenges:

- how do you scale it to improve throughput.
- you may want count things for different time frames but only do json parsing once.
- explain how you would cope with failure it the app crashes mid day / mid year.
- when creating e.g. per minute statistics, how do you handle frames that arrive late or frames with a random timestamp (e.g. hit by a bit-flip), describe a strategy?
- measure also the performance impact / overhead of the json parse

## Hints:

- tap into other peoples know-how and code.
- expected time is ~8 hours. If you are above that think about which parts to leave out and just document HOW you would do them.
- you should not use any big data framework, just your favourite fast programming language that has a Kafka driver. The most simplest version to archive step 5 can be done in about 10 lines of code (plus the boilerplate to read from kafka).
- check that your last commit compiles.
- be smart and impress us. It does not matte if you impress us with nice clean code or with a very clever hack to archive the business goal in short time.
