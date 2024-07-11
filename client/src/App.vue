<script setup lang="ts">
import { ref } from 'vue';
import NavBar from './components/NavBar.vue'

const timelineItems = ref([
  {title: 'Software Developer Apprentice', content: ['Developing and maintaining Java applications', 'ABAP and SAPUI5 development'], timestamp: '01.08.2018 - 01.08.2022'},
  {title: 'Software Engineer',  content: ['Migrating existing apps to OpenShift Platform', 'Engineering and maintaining k8s DevOps infrastructure'], timestamp: '01.08.2022 - 31.07.2024'}
])

</script>

<template>
  <div class="main-container">
    <NavBar style="height: inherit; grid-area: header" ></NavBar>
    <div class="card-container">
      <h2>Who even am I?</h2>
      <p>
        I'm a 22 years old software developer currently working at the Würth Itensis AG in Chur, Switzerland.
      </p>
      <p>
        Programming was a hobby of mine since back in my school days.
        I especially like to challange myself in new technologies, frameworks or programming languages.
      </p>
      <br>
      <p>
        I'm eager to learn new things about pretty much everything
        that software engineering has to offer. I don't back down from the unknown, and I hope this little website just
        shows you that.
      </p>
      <h2>Working Experience</h2>
      <Timeline :value="timelineItems" align="alternate" style="margin-top: 70rem">
        <template #marker>
          <span style="height: 2rem; width: 2rem; border-radius: 1rem; background-color: deeppink; display: flex; justify-content: center; align-items: center;">
            <i class="pi pi-crown"></i>
          </span>
        </template>
        <template #content="slotProps">
          <Card style="height: auto; margin-bottom: 5rem;" v-animateonscroll="{ enterClass: 'animate-fadein', leaveClass: 'animate-fadeout' }">
            <template #title>
              {{ slotProps.item.title }}
            </template>
            <template #subtitle>
              {{ slotProps.item.timestamp}}
            </template>
            <template #content>
              <ul style="text-align: left">
                <li style="margin: 1rem" v-for="item in slotProps.item.content">{{item}}</li>
              </ul>
            </template>
          </Card>
        </template>
      </Timeline>
    </div>
  </div>
</template>

<style scoped>
@keyframes fade-in {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}

@keyframes fade-out {
  0% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}
.animate-fadein {
  animation: fade-in 1s ease-in;
}
.animate-fadeout {
  animation: fade-out 1s ease-out;
}
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
.card-container {
  width: inherit;
  height: inherit;
  display: flex;
  align-content: center;
  flex-direction: column;
  grid-area: main;
}
.input-with-label {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  margin-bottom: 2rem;
}
.message-container {
  position: absolute;
  width: inherit;
  height: inherit;
  display: flex;
  gap: 0.5rem;
  flex-direction: column;
  align-items: end;
  grid-area: sidebar;
  padding: 1rem;
}

.main-container {
  height: 100%;
  width: 100%;
  display: grid;
  grid-template-columns: 25% 25% 25% 25%;
  grid-template-rows: 10% 80% 10%;
  grid-template-areas:
    "header header header header"
    ". main main ."
    "footer footer footer footer";
  padding: 0.25rem;
  position: relative;
}
.test-container {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  background-color: white;
  align-items: center;
  grid-area: footer;
}
p {
  margin: 0;
}
</style>
