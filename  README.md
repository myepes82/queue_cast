## QueueCast

QueueCast is a general-purpose WebSocket (WS) server built using Golang and its standard libraries, including `http`, `websocket`, and `gzip`. The server is designed to handle client requests following the publisher/subscriber model. It features topics for filtering events, supports event persistence, and offers additional functionalities to enhance the overall experience.

### Features

- **Publisher/Subscriber Model**: QueueCast enables efficient communication between clients using the publisher/subscriber model, allowing clients to publish and subscribe to various topics.
- **Topic Filtering**: Events are filtered based on topics, ensuring that subscribers receive only the relevant events.
- **Event Persistence**: QueueCast supports event persistence, ensuring that events are not lost and can be retrieved when needed.
- **Standard Libraries**: Built using Golang's standard libraries (`http`, `websocket`, and `gzip`), ensuring performance and reliability.

### Getting Started

#### Prerequisites

- Go 1.22.2 or higher installed on your machine.

#### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/myepes82/queue_cast.git
   cd queue_cast


2. Install dependencies:

   ```bash
   go mod tidy

#### Usage

