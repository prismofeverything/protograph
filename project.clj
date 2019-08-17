(defproject protograph "0.0.22"
  :description "tranform a stream of messages into a graph"
  :url "http://github.com/prismofeverything/protograph"
  :license {:name "Eclipse Public License"
            :url "http://www.eclipse.org/legal/epl-v10.html"}
  :dependencies [[org.clojure/clojure "1.8.0"]
                 [org.clojure/tools.cli "0.3.5"]
                 [cheshire "5.7.1"]
                 [instaparse "1.4.8"]
                 [org.flatland/ordered "1.5.7"]
                 [io.forward/yaml "1.0.9" :exclusions [org.flatland/ordered]]
                 [com.taoensso/timbre "4.8.0"]
                 [clojurewerkz/propertied "1.2.0"]
                 [org.apache.kafka/kafka_2.11 "0.10.0.1"]]
  :repositories [["sonatype snapshots"
                  "https://oss.sonatype.org/content/repositories/snapshots"]
                 ["sonatype releases"
                  "https://oss.sonatype.org/content/repositories/releases"]]
  :jvm-opts ["-Xmx12g" "-Xms12g" "-XX:-OmitStackTraceInFastThrow"]
  :main protograph.template)
