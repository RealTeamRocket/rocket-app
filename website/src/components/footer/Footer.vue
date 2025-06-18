<template>
  <footer class="footer">
    <div class="footer-content">
      <div class="footer-section footer-address">
        <span class="footer-title">Address</span>
        <span>Hochschule für Technik Stuttgart</span>
        <span>Schellingstraße 24</span>
        <span>70174 Stuttgart</span>
        <span>Germany</span>
      </div>
      <div class="footer-section footer-legal">
        <span class="footer-title">Legal</span>
        <router-link to="/legal-notice">Legal Notice</router-link>
        <router-link to="/privacy-policy">Privacy Policy</router-link>
        <router-link to="/accessibility">Accessibility</router-link>
      </div>
      <div class="footer-section footer-contact">
        <span class="footer-title">Contact</span>
        <a href="mailto:info@hft-stuttgart.de" id="footer-email">info@hft-stuttgart.de</a>
        <a href="https://www.instagram.com/hft_stuttgart/" target="_blank" rel="noopener" aria-label="Instagram" class="footer-instagram">
            <i class="fab fa-instagram" style="font-size: 1.3em;"></i>
        </a>
        <a href="https://www.youtube.com/c/Hochschulef%C3%BCrTechnikStuttgart" target="_blank" rel="noopener" aria-label="YouTube" class="footer-youtube">
            <i class="fab fa-youtube" style="font-size: 1.3em;"></i>
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
  max-width: 1100px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 2rem;
}

.footer-section {
  display: flex;
  flex-direction: column;
  gap: 0.3rem;
  min-width: 180px;
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

.footer-legal a:hover, #footer-email:hover {
  color: #ffd580; 
  text-decoration: underline;
}

.footer-instagram,
.footer-youtube {
  display: inline-flex;
  align-items: center;
  margin-top: 0.1rem;
  margin-right: 0.6rem;
}

.footer-instagram:hover i {
  color: #f8a8f4;
}

.footer-youtube:hover i {
  color: #ff4c4c;
}

@media (max-width: 900px) {
  .footer-content {
    flex-direction: column;
    align-items: flex-start;
    gap: 1.2rem;
    padding: 0 1rem;
  }
  .footer-section {
    width: 100%;
    min-width: unset;
  }
}
</style>