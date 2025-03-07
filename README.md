Apache Kafka เป็นแพลตฟอร์มสำหรับการสตรีมข้อมูลแบบกระจาย (distributed event streaming platform) ที่ใช้กันอย่างแพร่หลายในระบบที่ต้องการการประมวลผลข้อมูลแบบเรียลไทม์

---

#### 1️⃣ **Kafka คืออะไร?**
Kafka เป็นแพลตฟอร์มที่ช่วยให้ระบบสามารถส่ง, รับ และประมวลผลข้อมูลแบบ event-driven หรือสตรีมมิงได้อย่างมีประสิทธิภาพ โดยมีองค์ประกอบหลักคือ:
- **Producer** – ผู้ส่งข้อมูลเข้า Kafka
- **Topic** – ช่องทางที่ใช้รับส่งข้อมูล
- **Broker** – เซิร์ฟเวอร์ที่เก็บข้อมูลและกระจายไปยัง Consumer
- **Consumer** – ผู้รับข้อมูลจาก Kafka
- **Zookeeper** – ใช้จัดการ metadata และควบคุม cluster

---

#### 2️⃣ **ติดตั้ง Kafka บนเครื่องของคุณ**
Kafka ต้องการ Java และ Zookeeper ทำงานร่วมกัน  
👉 **ขั้นตอนการติดตั้ง**
1. ติดตั้ง **Java** (แนะนำให้ใช้ Java 8 หรือ 11)  
   ```bash
   java -version
   ```
   ถ้ายังไม่มี Java ให้ติดตั้ง:
   ```bash
   sudo apt update && sudo apt install openjdk-11-jdk -y
   ```
   
2. ดาวน์โหลด Kafka จาก [Apache Kafka](https://kafka.apache.org/downloads)  
   หรือใช้คำสั่ง:
   ```bash
   wget https://downloads.apache.org/kafka/3.5.1/kafka_2.13-3.5.1.tgz
   tar -xvzf kafka_2.13-3.5.1.tgz
   cd kafka_2.13-3.5.1
   ```

3. เริ่มต้น Zookeeper (Kafka ต้องใช้ Zookeeper)  
   ```bash
   bin/zookeeper-server-start.sh config/zookeeper.properties
   ```

4. เริ่มต้น Kafka Broker  
   ```bash
   bin/kafka-server-start.sh config/server.properties
   ```

---

#### 3️⃣ **ลองใช้งาน Kafka เบื้องต้น**
📌 **สร้าง Topic ใหม่**
```bash
bin/kafka-topics.sh --create --topic my-topic --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
```

📌 **ตรวจสอบ Topic**
```bash
bin/kafka-topics.sh --list --bootstrap-server localhost:9092
```

📌 **ส่งข้อความด้วย Producer**
```bash
bin/kafka-console-producer.sh --topic my-topic --bootstrap-server localhost:9092
```
ลองพิมพ์ข้อความ เช่น:
```
Hello Kafka!
```

📌 **รับข้อความด้วย Consumer**
```bash
bin/kafka-console-consumer.sh --topic my-topic --from-beginning --bootstrap-server localhost:9092
```

### **4️⃣ ทดสอบ Kafka**
📌 **สร้าง Topic ใหม่**  
```bash
docker exec -it kafka kafka-topics --create --topic my-topic --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
```
📌 **ตรวจสอบรายการ Topic**  
```bash
docker exec -it kafka kafka-topics --list --bootstrap-server localhost:9092
```
📌 **ส่งข้อความผ่าน Producer**  
```bash
docker exec -it kafka kafka-console-producer --topic my-topic --bootstrap-server localhost:9092
```
พิมพ์ข้อความแล้วกด Enter เช่น  
```
Hello Kafka!
```
📌 **รับข้อความผ่าน Consumer**  
```bash
docker exec -it kafka kafka-console-consumer --topic my-topic --from-beginning --bootstrap-server localhost:9092
```
ถ้าทุกอย่างทำงานถูกต้อง คุณจะเห็นข้อความที่ Producer ส่งมา 🎉

---

### 🚀 **ขั้นตอนต่อไป**
ถ้าคุณเข้าใจพื้นฐานแล้ว เราสามารถเรียนรู้เรื่อง:
✅ **Kafka Architecture** (Replication, Partitioning, Retention)  
✅ **การตั้งค่า Kafka Cluster** (Multi-broker)  
✅ **Kafka Stream Processing** (ใช้ Kafka Streams, ksqlDB, Flink)  
✅ **การใช้งาน Kafka กับภาษาโปรแกรมต่างๆ เช่น Python, Java, Node.js**  

อยากเจาะลึกเรื่องไหนเพิ่มเติม บอกได้เลย! 😊

ถ้าคุณอยากเป็น **มือหนึ่ง** ด้านการใช้ Kafka ตั้งแต่พื้นฐานจนถึงระดับสูง นี่คือ **Roadmap** ที่แนะนำให้คุณเรียนรู้ทีละขั้นตอน  

---

## **📌 Roadmap การเรียนรู้ Apache Kafka ตั้งแต่พื้นฐานถึงขั้นสูง**
> แบ่งออกเป็น 4 ระดับ: พื้นฐาน → กลาง → สูง → ขั้นเทพ  

---

### **🟢 ระดับพื้นฐาน (Beginner)**
📍 **ทำความเข้าใจแนวคิดของ Kafka**  
✅ Kafka คืออะไร? ใช้ทำอะไร?  
✅ Kafka เหมาะกับงานประเภทไหน (Real-time Data Streaming, Event-driven, Log Aggregation)  
✅ ทำความเข้าใจ **Kafka Components**  
   - **Producer** (ผู้ส่งข้อมูล)  
   - **Consumer** (ผู้รับข้อมูล)  
   - **Broker** (เซิร์ฟเวอร์ Kafka)  
   - **Topic & Partitions** (โครงสร้างของ Kafka)  
   - **Zookeeper** (ตัวจัดการ Metadata)  

📍 **ติดตั้ง & ตั้งค่า Kafka**  
✅ ติดตั้ง Kafka และ Zookeeper บนเครื่อง  
✅ เรียนรู้คำสั่ง CLI พื้นฐาน  
   - สร้าง Topic  
   - ส่งข้อความจาก Producer  
   - อ่านข้อความจาก Consumer  
✅ เข้าใจ **Kafka Message Format** และ **Offset**  

📍 **ทดลองใช้งาน Kafka กับภาษาโปรแกรม**  
✅ ใช้ Python / Java / Node.js เชื่อมต่อ Kafka  
✅ ส่งและรับข้อความผ่าน Kafka Client  

---

### **🟡 ระดับกลาง (Intermediate)**
📍 **เข้าใจ Kafka Internals**  
✅ ทำความเข้าใจ **Replication** และ **Partitioning**  
✅ รู้จัก **Kafka Retention Policy** (เก็บข้อมูลใน Kafka ได้นานแค่ไหน)  
✅ ใช้งาน **Consumer Group** และการทำ Load Balancing  
✅ เข้าใจ **Leader & Follower Concept** ใน Kafka  

📍 **Kafka Cluster & High Availability**  
✅ ตั้งค่า Kafka แบบ **Multi-Broker Cluster**  
✅ ใช้ **ZooKeeper** ในการจัดการคลัสเตอร์  
✅ ทดสอบ Failover และการทำ **Data Replication**  

📍 **Kafka Producer & Consumer ขั้นสูง**  
✅ ตั้งค่า **Producer Acknowledgment & Retries**  
✅ ปรับแต่ง **Consumer Offset Management**  
✅ ใช้ **Kafka Consumer Group** เพื่อรองรับการทำงานแบบกระจาย  

📍 **การทำ Stream Processing**  
✅ ใช้ **Kafka Streams API** ประมวลผลข้อมูลแบบ Real-time  
✅ ทำ Aggregation และ Windowing ใน Kafka Streams  
✅ ใช้ **ksqlDB** สำหรับ Query ข้อมูลจาก Kafka  

---

### **🔴 ระดับสูง (Advanced)**
📍 **การบริหารจัดการ Kafka Cluster**  
✅ ปรับแต่ง **Kafka Performance** (Tuning & Optimization)  
✅ ติดตามการทำงานของ Kafka ด้วย **Monitoring Tools** เช่น Prometheus + Grafana  
✅ ใช้ **Kafka Connect** เพื่อเชื่อม Kafka กับ Database (MySQL, PostgreSQL, MongoDB, Elasticsearch)  
✅ ตั้งค่า **Security** (SSL, ACL, Authentication)  

📍 **Integrating Kafka กับระบบอื่นๆ**  
✅ ใช้ **Kafka กับ Kubernetes** (Kafka on K8s)  
✅ เชื่อม Kafka กับ **Apache Flink / Spark Streaming**  
✅ ใช้ **Kafka กับ Data Lake / Data Warehouse**  

📍 **Deep Dive Kafka Internals**  
✅ เข้าใจ Kafka Storage และ **Log Segments**  
✅ ทำ Partition Rebalancing และ Cluster Scaling  
✅ วิเคราะห์ **Kafka Performance Bottlenecks**  

---

### **🟣 ขั้นเทพ (Expert)**
📍 **Kafka Architecture ระดับลึก**  
✅ ออกแบบ Kafka ให้รองรับ **ล้านข้อความต่อวินาที**  
✅ ทำ **Cross-Data Center Replication** ด้วย **MirrorMaker 2**  
✅ ใช้ **Confluent Schema Registry** และ **Avro** เพื่อจัดการโครงสร้างข้อมูล  
✅ Optimize **Throughput & Latency** ของ Kafka  

📍 **Kafka Use Cases ระดับองค์กร**  
✅ สร้าง **Event-driven Microservices** ด้วย Kafka  
✅ ใช้ Kafka กับ **Machine Learning & AI**  
✅ ทำ **Change Data Capture (CDC)** ด้วย Kafka  

📍 **สร้าง Kafka Solution ที่รองรับ Production**  
✅ ออกแบบ Kafka Infrastructure ให้รองรับการขยายตัว  
✅ ตั้งค่า **Disaster Recovery Plan** สำหรับ Kafka  
✅ วิเคราะห์และแก้ปัญหา Kafka ในระบบขนาดใหญ่  

---

## **🚀 สรุป: เส้นทางสู่ Kafka Master**
📌 **เริ่มจากพื้นฐาน** → ติดตั้งและใช้งาน Kafka เบื้องต้น  
📌 **เข้าใจโครงสร้าง Kafka** → Consumer Group, Partitioning, Replication  
📌 **เรียนรู้ Kafka Streams** → การประมวลผลข้อมูลแบบเรียลไทม์  
📌 **ปรับแต่ง Kafka Cluster** → เพิ่ม Performance และทำ High Availability  
📌 **ใช้ Kafka กับระบบอื่นๆ** → Data Pipeline, AI, Machine Learning  
📌 **พัฒนา Kafka Solution ระดับองค์กร** → ออกแบบ Kafka ที่รองรับปริมาณข้อมูลมหาศาล  

---

🔥 **เป้าหมายของคุณคืออะไร?**  
- คุณต้องการใช้งาน Kafka ในองค์กร หรือเป็นผู้เชี่ยวชาญด้าน Data Engineering?  
- อยากเริ่มจากพื้นฐานไหนก่อน หรือให้ช่วยแนะนำ Workshop? 😊

```
go-kafka
├─ README.md
├─ cmd
│  ├─ consumer
│  │  └─ main.go
│  └─ producer
│     └─ main.go
├─ config
│  ├─ config.go
│  └─ config.yml
├─ docker-compose.yml
├─ go.mod
├─ go.sum
└─ md

```

docker build -t kafka-lag-exporter .
docker run -d --name kafka-lag-exporter -p 8000:8000 kafka-lag-exporter
