# Awesomeai

<table>
  <tr>
    <td valign="top" width="100">
      <img src="https://raw.githubusercontent.com/avkcode/awesomeai/refs/heads/main/favicon.svg" 
           alt="AwesomeAI Logo" 
           width="80">
    </td>
    <td valign="middle">
      A visual periodic table-style interface for exploring AI models. Organizes text, image, and multimodal models by category, architecture, and use case with interactive filtering.
    </td>
  </tr>
</table>

<div style="height: 3px; background: linear-gradient(90deg, #fff8d8, #f7e394, #e6c35a, #c59b3a, #a67c00); margin: 20px 0;"></div>

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
