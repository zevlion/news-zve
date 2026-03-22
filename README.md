# News ZVE

A high-performance, minimalist news aggregator that delivers instant global reporting by via the Associated Press internal API.

## Features

- **Asynchronous Ingestion:** HTML parsing via `goquery` with custom sanitization logic for Associated Press data streams.
- **Monochrome UI Engine:** CSS architecture leveraging Tailwind utility classes and the Inter variable font.
- **Stateful Reader:** Vanilla JS implementation for full-viewport modal rendering with history state synchronization.
- **Advanced UX Logic:** \* **Context API Interception:** Custom DOM event listeners overriding default context menus for deep-link generation.
  - **Optimistic UI:** CSS-driven skeleton loaders for API fetch state management.
  - **Media Query Synchronization:** Native `prefers-color-scheme` integration.
- **Containerization:** Multi-stage `Dockerfile` targeting `alpine:latest` for static binary execution.

### Deployment

1. **Clone & Initialize:**

   ```bash
   git clone [https://github.com/zevlion/news-zve.git](https://github.com/zevlion/news-zve.git) && cd news-zve
   ```

2. **Build Image:**

   ```bash
   docker build -t news-zve .
   ```

3. **Runtime Execution:**

   ```bash
   docker run -d -p 8080:8080 --name news-zve-prod news-zve
   ```

# LICENSE

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
