<template>
  <div class="home-guest">
    <div class="hero">
      <img src="/src/assets/icons/rocket.svg" alt="Rocket App Logo" class="hero-logo" />
      <h1>Welcome to the Rocket App</h1>
      <p class="subtitle">
        The gamified step tracker that makes every walk an adventure!
      </p>
      <router-link to="/download" class="download-btn">
        Download for Android
      </router-link>
    </div>
    <div class="features-section">
      <h2>Why Rocket App?</h2>
      <ul class="features-list">
        <li>
          <span class="feature-icon">👣</span>
          <span>
            <strong>Step Tracking</strong> — Track your daily steps and see your progress over time.
          </span>
        </li>
        <li>
          <span class="feature-icon">🗺️</span>
          <span>
            <strong>Route Planning & Tracking</strong> — Plan your runs and walks on a map and record your routes.
          </span>
        </li>
        <li>
          <span class="feature-icon">🏆</span>
          <span>
            <strong>Highscore Lists</strong> — Compete with others and climb the leaderboards!
          </span>
        </li>
        <li>
          <span class="feature-icon">🤝</span>
          <span>
            <strong>Friendlist</strong> — Add friends, follow their progress, and motivate each other.
          </span>
        </li>
        <li>
          <span class="feature-icon">🎯</span>
          <span>
            <strong>Challenges</strong> — Take on daily and weekly challenges to earn Rocket Points.
          </span>
        </li>
        <li>
          <span class="feature-icon">🚀</span>
          <span>
            <strong>Gamified Experience</strong> — Earn Rocket Points, unlock achievements, and make fitness fun!
          </span>
        </li>
      </ul>
    </div>
    <div class="cta-section">
      <p>
        Ready to join the Rocket App community? <br>
        <router-link to="/register" class="cta-link">Create your free account</router-link>
        or
        <router-link to="/login" class="cta-link">log in</router-link>
        to get started!
      </p>
      <router-link to="/download" class="download-btn secondary">
        Download APK
      </router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.home-guest {
  max-width: 700px;
  margin: 2.5rem auto 0 auto;
  padding: 2.5rem 1.5rem 2rem 1.5rem;
  background: #fff;
  border-radius: 1.2rem;
  box-shadow: 0 2px 12px rgba(30,60,114,0.08);
  display: flex;
  flex-direction: column;
  align-items: center;
}

.hero {
  text-align: center;
  margin-bottom: 2.5rem;
}

.hero-logo {
  width: 70px;
  height: 70px;
  margin-bottom: 1rem;
  filter: drop-shadow(0 2px 8px rgba(30,60,114,0.10));
}

h1 {
  font-size: 2.3rem;
  font-weight: 800;
  color: #2a5298;
  margin-bottom: 0.5rem;
}

.subtitle {
  font-size: 1.15rem;
  color: #64748b;
  margin-bottom: 1.5rem;
}

.download-btn {
  display: inline-block;
  background: #4a90e2;
  color: #fff;
  font-size: 1.15rem;
  font-weight: 600;
  border-radius: 0.6rem;
  padding: 0.7rem 2.1rem;
  text-decoration: none;
  margin-top: 0.7rem;
  margin-bottom: 0.5rem;
  transition: background 0.18s;
  box-shadow: 0 2px 8px rgba(30,60,114,0.06);
}
.download-btn:hover {
  background: #2a5298;
}
.download-btn.secondary {
  background: #ffb347;
  color: #1e3c72;
  margin-top: 1.2rem;
}
.download-btn.secondary:hover {
  background: #ffd580;
  color: #1e3c72;
}

.features-section {
  width: 100%;
  margin-bottom: 2.5rem;
}

.features-section h2 {
  text-align: center;
  font-size: 1.5rem;
  color: #1e3c72;
  margin-bottom: 1.2rem;
  font-weight: 700;
}

.features-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.features-list li {
  display: flex;
  align-items: flex-start;
  gap: 1.1rem;
  font-size: 1.08rem;
  margin-bottom: 1.1rem;
  background: #f3f8ff;
  border-radius: 0.7rem;
  padding: 0.85rem 1.1rem;
  box-shadow: 0 1px 4px rgba(30,60,114,0.04);
}

.feature-icon {
  font-size: 1.5rem;
  margin-top: 0.1rem;
}

.cta-section {
  text-align: center;
  margin-top: 1.2rem;
}

.cta-link {
  color: #4a90e2;
  font-weight: 600;
  text-decoration: underline;
  margin: 0 0.2rem;
  transition: color 0.18s;
}
.cta-link:hover {
  color: #2a5298;
}
</style>
