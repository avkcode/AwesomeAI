# Awesomeai

<div style="display: flex; align-items: center; gap: 20px; margin: 10px 0;">
  <img src="https://raw.githubusercontent.com/avkcode/awesomeai/refs/heads/main/favicon.svg" 
       alt="AwesomeAI Logo" 
       width="80">
  <div>
    A visual periodic table-style interface for exploring AI models. Organizes text, image, and multimodal models by category, architecture, and use case with interactive filtering.
  </div>
</div>

## Configuration

### `MODEL_DATA_URL` Environment Variable

Specifies the URL or local path to load the AI models data from.

**Usage:**
```bash
# Use default GitHub URL
make run

# Use a custom remote URL
MODEL_DATA_URL="https://example.com/custom-models.json" make run

# Use a local JSON file
MODEL_DATA_URL="static/models.json" make run
