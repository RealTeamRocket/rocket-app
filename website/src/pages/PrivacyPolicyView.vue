<template>
  <Navbar />
  <div class="legal-page">
    <h1>Privacy Policy</h1>

    <section>
      <h2>1. Introduction</h2>
      <p>
        We value your privacy and are committed to protecting your personal data. This Privacy Policy explains how we collect, use, and store your information when you use our website and mobile application.
      </p>
    </section>

    <section>
      <h2>2. Data We Collect</h2>
      <ul>
        <li>Account information (e.g., username, email address)</li>
        <li>Fitness and activity data (e.g., steps, routes, challenges)</li>
        <li>Device information (e.g., browser type, IP address)</li>
      </ul>
    </section>

    <section>
      <h2>3. How We Use Your Data</h2>
      <p>
        We use your data to:
      </p>
      <ul>
        <li>Provide and improve our services</li>
        <li>Track your activity and progress</li>
        <li>Enable social features like leaderboards and friend connections</li>
        <li>Respond to your inquiries and support requests</li>
      </ul>
    </section>

    <section>
      <h2>4. Data Sharing</h2>
      <p>
        We do not sell or rent your personal information. Data may be shared with third parties only when required by law or to provide essential services (e.g., hosting providers).
      </p>
    </section>

    <section>
      <h2>5. Your Rights</h2>
      <p>You have the right to:</p>
      <ul>
        <li>Access and review your personal data</li>
        <li>Request data correction or deletion</li>
        <li>Withdraw your consent at any time</li>
      </ul>
    </section>

    <section>
      <h2>6. Contact</h2>
      <p>
        For questions regarding this policy, please contact us at:
        <br />
        <strong>Email:</strong> info@hft-stuttgart.de
      </p>
    </section>

    <section>
      <h2>7. Changes to This Policy</h2>
      <p>
        We may update this Privacy Policy from time to time. Changes will be published on this page. This statement was last updated on 17 June 2025.
      </p>
    </section>
  </div>
  <Footer />
</template>

<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'
import Footer from '@/components/footer/Footer.vue'
</script>

<style scoped>
.legal-page {
  max-width: 900px;
  margin: 0 auto;
  padding: 2rem;
  line-height: 1.7;
  color: #333;
  background: #fdfdfd;
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
</style>