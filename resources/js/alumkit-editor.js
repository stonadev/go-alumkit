import EditorJS from '@editorjs/editorjs';
import Header from '@editorjs/header';
import List from '@editorjs/list';
import Paragraph from '@editorjs/paragraph';
import Table from '@editorjs/table';
import ImageTool from '@editorjs/image';

class AlumkitEditor {
  constructor(elementId, options = {}) {
    this.elementId = elementId;
    this.options = options;
    this.editor = null;
    this.init();
  }

  init() {
    const element = document.getElementById(this.elementId);
    if (!element) {
      console.error(`Element #${this.elementId} not found`);
      return;
    }

    this.editor = new EditorJS({
      holder: this.elementId,
      placeholder: this.options.placeholder || 'Start writing...',
      tools: {
        header: {
          class: Header,
          inlineToolbar: true,
          config: {
            levels: [1, 2, 3, 4],
            defaultLevel: 2
          }
        },
        list: {
          class: List,
          inlineToolbar: true
        },
        paragraph: {
          class: Paragraph,
          inlineToolbar: true
        },
        table: {
          class: Table,
          inlineToolbar: true
        },
        image: {
          class: ImageTool,
          config: {
            endpoints: {
              byFile: this.options.uploadUrl || '/admin/api/upload'
            },
            types: 'image/*'
          }
        }
      },
      data: this.options.data || {},
      onChange: (api, event) => {
        if (this.options.onChange) {
          this.options.onChange(api, event);
        }
      },
      onReady: () => {
        if (this.options.onReady) {
          this.options.onReady(this.editor);
        }
      }
    });
  }

  async save() {
    if (!this.editor) return null;
    return await this.editor.save();
  }

  async render(data) {
    if (!this.editor) return;
    await this.editor.render(data);
  }

  destroy() {
    if (this.editor) {
      this.editor.destroy();
      this.editor = null;
    }
  }
}

// Export for use
if (typeof window !== 'undefined') {
  window.AlumkitEditor = AlumkitEditor;
}

export default AlumkitEditor;
