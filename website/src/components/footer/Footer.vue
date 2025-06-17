<template>
  <footer class="footer">
    <div class="footer-content">
      <div class="footer-section footer-legal">
        <span class="footer-title">Legal</span>
        <a href="/legal-notice" target="_blank" rel="noopener">Legal Notice</a>
        <a href="/privacy-policy" target="_blank" rel="noopener">Privacy Policy</a>
      </div>
      <div class="footer-section footer-contact">
        <span class="footer-title">Contact</span>
        <a href="mailto:your@email.com">your@email.com</a>
        <a href="https://instagram.com/youraccount" target="_blank" rel="noopener" aria-label="Instagram" class="footer-instagram">
          <span style="font-size: 1.3em;">📸</span>
        </a>
      </div>
    </div>
  </footer>
</template>

<style scoped>
.footer {
  width: 100%;
  background: linear-gradient(90deg, #1e3c72 0%, #2a5298 100%);
  color: #fff;
  border-top: 2px solid #1e3c72;
  padding: 1.2rem 0;
  font-size: 1rem;
  position: relative;
  left: 0;
  bottom: 0;
}

.footer-content {
  max-width: 900px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.footer-section {
  display: flex;
  flex-direction: column;
  gap: 0.3rem;
}

.footer-title {
  font-weight: bold;
  font-size: x-large;
  margin-bottom: 0.2rem;
  letter-spacing: 0.02em;
  color: #ffd580; 
}

.footer-legal a,
.footer-contact a {
  color: #e0e0e0;
  text-decoration: none;
  transition: color 0.2s;
  font-size: 1em;
}

.footer-legal a:hover,
.footer-contact a:hover {
  color: #ffd580; 
  text-decoration: underline;
}

.footer-instagram {
  display: inline-flex;
  align-items: center;
  margin-top: 0.1rem;
}

@media (max-width: 700px) {
  .footer-content {
    flex-direction: column;
    align-items: flex-start;
    gap: 1.2rem;
    padding: 0 1rem;
  }
  .footer-section {
    width: 100%;
  }
}
</style>