import Cropper from 'cropperjs';

class AlumkitCropper {
  constructor(options = {}) {
    this.options = {
      aspectRatio: options.aspectRatio || 16 / 9,
      viewMode: options.viewMode || 1,
      autoCropArea: options.autoCropArea || 0.8,
      ...options
    };
    this.modal = null;
    this.image = null;
    this.cropper = null;
    this.resolve = null;
  }

  open(imageSrc) {
    return new Promise((resolve) => {
      this.resolve = resolve;
      this.createModal(imageSrc);
    });
  }

  createModal(imageSrc) {
    // Create modal
    this.modal = document.createElement('div');
    this.modal.className = 'alumkit-cropper-modal';
    this.modal.innerHTML = `
      <div class="alumkit-cropper-container">
        <div style="margin-bottom: 1rem; display: flex; justify-content: space-between; align-items: center;">
          <h3 style="margin: 0; font-size: 1.125rem; font-weight: 600;">Crop Image</h3>
          <button class="alumkit-cropper-cancel" style="background: none; border: none; cursor: pointer; font-size: 1.5rem;">&times;</button>
        </div>
        <div style="max-height: 70vh; overflow: hidden;">
          <img id="alumkit-cropper-image" src="${imageSrc}" style="max-width: 100%; display: block;">
        </div>
        <div style="margin-top: 1rem; display: flex; gap: 0.5rem; justify-content: flex-end;">
          <button class="alumkit-cropper-cancel btn" style="background: #e2e8f0; color: #1e293b;">Cancel</button>
          <button class="alumkit-cropper-confirm btn btn-primary">Crop</button>
        </div>
      </div>
    `;

    document.body.appendChild(this.modal);

    // Initialize cropper
    this.image = document.getElementById('alumkit-cropper-image');
    this.cropper = new Cropper(this.image, this.options);

    // Event listeners
    this.modal.querySelector('.alumkit-cropper-cancel').addEventListener('click', () => this.cancel());
    this.modal.querySelector('.alumkit-cropper-confirm').addEventListener('click', () => this.confirm());
    this.modal.addEventListener('click', (e) => {
      if (e.target === this.modal) this.cancel();
    });
  }

  cancel() {
    if (this.cropper) {
      this.cropper.destroy();
    }
    if (this.modal) {
      this.modal.remove();
    }
    if (this.resolve) {
      this.resolve(null);
    }
  }

  confirm() {
    if (!this.cropper) return;

    const canvas = this.cropper.getCroppedCanvas();
    canvas.toBlob((blob) => {
      if (this.cropper) {
        this.cropper.destroy();
      }
      if (this.modal) {
        this.modal.remove();
      }
      if (this.resolve) {
        this.resolve(blob);
      }
    }, 'image/jpeg', 0.9);
  }

  destroy() {
    if (this.cropper) {
      this.cropper.destroy();
    }
    if (this.modal) {
      this.modal.remove();
    }
  }
}

// Export for use
if (typeof window !== 'undefined') {
  window.AlumkitCropper = AlumkitCropper;
}

export default AlumkitCropper;
