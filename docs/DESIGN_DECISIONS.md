# Design Decisions

## Architecture

### Microservices Over Monolith

**Why microservices**:
- **Modular scenarios**: Different products need different service combinations
- **Independent scaling**: CHISG (CPU-heavy) vs Behavior Recording (I/O-heavy)
- **Technology flexibility**: Go for APIs, Python for ML, Node for rapid dev
- **Team autonomy**: Different services = different development speeds
- **Fault isolation**: One service failure doesn't crash entire system
- **Product flexibility**: Can compose services differently per customer need

**Composition patterns**:
```yaml
AI Tutor (MVP Priority):
  services: [humanos-api, chisg-api, esp-assist-api, weaviate, mongodb]
  target: Individual students, families
  
GCSE Revision Tool (Quick Revenue):
  services: [humanos-api, chisg-api, pdf-rag, skillsmap-api, esp-assist-api]
  target: Exam students (Year 10-11)

Skills Tree Rising (Collaborator):
  services: [skillsmap-api, esp-assist-api, ms-graph-integration]
  target: Leadership training programs

Full ESP Suite (Long-term):
  services: [ALL]
  target: Schools (comprehensive MIS)
  note: Only product requiring all services
```

**Critical Insight**: Full ESP Suite is NOT the priority. Focus on high-value service subsets that can launch faster and generate revenue sooner.

### Service Communication Strategy

**Synchronous (REST)**: Knowledge queries, student lookups
**Asynchronous (Events)**: Behavior triggers, escalation workflows  
**WebSocket (Real-time)**: Live behavior updates, scoreboards

**Anti-pattern avoided**: Direct database access between services
**Pattern used**: API contracts + event bus

### Go Backend for Core APIs

**Why Go**:
- **Performance**: High throughput, low latency
- **Concurrency**: Efficient handling of thousands of requests
- **Stability**: Reliable and predictable behavior
- **Ecosystem**: Rich set of tools and libraries

**Service language matrix**:
- **Go**: HumanOS, CHISG, Skills Map (performance-critical)
- **Python**: ESP Assist, PDF RAG (ML/NLP libraries)
- **Node.js**: ESP Behavior, Planning (rapid iteration)

### Docker Compose → Kubernetes Path

**Current**: Docker Compose (development + small deployments)
**Future**: Kubernetes (production scale)

**Why start with Docker Compose**:
- Faster iteration
- Simpler local development
- Easier testing of service interactions
- Natural progression to K8s (same container images)

## Cross-Project Integration (Updated)

### CHISG ↔ HumanOS ↔ Skills Map Triangle

### Federated Learning as Core Innovation

**The Big Idea**: Every deployment contributes to collective intelligence while maintaining complete privacy.

**Why This Matters**:
- **Network effects**: System gets smarter as more people use it
- **Privacy-first**: No PII ever leaves local deployment
- **Competitive moat**: Unique approach in EdTech
- **Continuous improvement**: Models improve without manual updates

**How It Works**:
````
<userPrompt>
Provide the fully rewritten file, incorporating the suggested code change. You must produce the complete file.
</userPrompt>
