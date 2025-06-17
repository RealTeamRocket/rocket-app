<template>
  <div class="page-wrapper">
    <Navbar />
    <main class="legal-content">
      <div class="legal-page">
        <h1>Accessibility</h1>

        <section>
          <h2>Commitment to Accessibility</h2>
          <p>
            We are committed to making our website and application accessible to everyone, including people with disabilities.
            Our goal is to ensure a user-friendly experience for all users, regardless of their physical or technical limitations.
          </p>
        </section>

        <section>
          <h2>Measures to Support Accessibility</h2>
          <ul>
            <li>Responsive design for different devices and screen readers</li>
            <li>Keyboard navigation and focus visibility</li>
            <li>Text alternatives for non-text content</li>
            <li>Consistent structure and headings</li>
          </ul>
        </section>

        <section>
          <h2>Feedback and Contact</h2>
          <p>
            If you encounter accessibility issues or have suggestions for improvement, please contact us:
            <br />
            <strong>Email:</strong> info@hft-stuttgart.de
          </p>
        </section>

        <section>
          <h2>Enforcement Procedure</h2>
          <p>
            In the event of unsatisfactory responses to your accessibility inquiry, you may contact the responsible regulatory authority in your region to enforce your rights under applicable accessibility legislation.
          </p>
        </section>

        <section>
          <h2>Last Updated</h2>
          <p>
            This statement was last updated on 17 June 2025.
          </p>
        </section>
      </div>
    </main>
    <Footer />
  </div>
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'
import Footer from '@/components/footer/Footer.vue'
</script>

<style scoped>
.page-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.legal-content {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  padding: 2rem 1rem;
  background: white; /* wichtig: Seitenhintergrund bleibt weiß */
}

.legal-page {
  max-width: 900px;
  width: 100%;
  background: #f5f5f5; /* grau nur für die Box */
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
  line-height: 1.7;
  color: #333;
}

.legal-page h1 {
  text-align: center;
  margin-bottom: 2rem;
  color: #1e3c72;
  font-size: 2.5rem;
}

.legal-page section {
  margin-bottom: 2rem;
}

.legal-page h2 {
  color: #2a5298;
  font-size: 1.5rem;
  margin-bottom: 0.5rem;
}

.legal-page ul {
  padding-left: 1.5rem;
  margin-top: 0.5rem;
}

.legal-page li {
  margin-bottom: 0.4rem;
}

.footer {
  margin-top: auto;
}
</style>