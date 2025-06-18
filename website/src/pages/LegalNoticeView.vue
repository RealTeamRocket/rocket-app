<template>
  <div class="page-wrapper">
    <Navbar />
    <main class="legal-content">
      <div class="legal-page">
        <h1>Legal Notice</h1>

        <section>
          <h2>Information According to § 5 TMG</h2>
          <p>
            Hochschule für Technik Stuttgart<br />
            Schellingstraße 24<br />
            70174 Stuttgart<br />
            Germany
          </p>
        </section>

        <section>
          <h2>Represented By</h2>
          <p>Prof. Dr. Katja Rade (President)</p>
        </section>

        <section>
          <h2>Contact</h2>
          <p>
            Phone: +49 711 8926 0<br />
            Email: info@hft-stuttgart.de
          </p>
        </section>

        <section>
          <h2>VAT ID</h2>
          <p>
            VAT Identification Number according to § 27 a UStG:<br />
            DE 147794431
          </p>
        </section>

        <section>
          <h2>Supervisory Authority</h2>
          <p>
            Ministry of Science, Research and the Arts Baden-Württemberg<br />
            Konrad-Adenauer-Straße 20<br />
            70173 Stuttgart
          </p>
        </section>

        <section>
          <h2>Disclaimer</h2>
          <p>
            Despite careful content control, we assume no liability for the content of external links. The operators of linked pages are solely responsible for their content.
            <br /><br />
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
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.legal-content {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  padding: 2rem 1rem;
  background: white;
}

.legal-page {
  max-width: 900px;
  width: 100%;
  background: #f5f5f5;
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
</style>