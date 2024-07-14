<script setup lang="ts">
import { ref } from 'vue';
import NavBar from './components/NavBar.vue'

const timelineItems = ref([
  {title: 'Apprentice - Würth Itensis AG', content: [
      'Developing and maintaining Java applications',
      'ABAP and SAPUI5 development'
    ], timestamp: '01.08.2018 - 01.08.2022'},
  {title: 'Software Engineer - Würth Itensis AG',  content: [
      'Migrating existing apps to OpenShift Platform',
      'Engineering and maintaining k8s DevOps infrastructure',
      'Developing and maintaining Java applications with Docker',
      'Developing and maintaining SAPUI5 and ABAP applications',
      'Requirements-Engineering with clients'
    ], timestamp: '01.08.2022 - 31.07.2024'}
])

</script>

<template>
  <div class="main-container">
    <NavBar style="height: inherit; grid-area: header" ></NavBar>
    <div class="skills-container">
      <h1>Skills</h1>
      <div class="skills-item">
        <p style="text-align: start;">Java</p>
        <ProgressBar :value="100" style="flex-grow: 1"/>
      </div>
      <div class="skills-item">
        <p style="text-align: start;">Docker</p>
        <ProgressBar :value="100" style="flex-grow: 1"/>
      </div>
      <div class="skills-item">
        <p style="text-align: start;">DevOps</p>
        <ProgressBar :value="90" style="flex-grow: 1"/>
      </div>
      <div class="skills-item">
        <p style="text-align: start; text-wrap: nowrap">Kubernetes</p>
        <ProgressBar :value="80" style="flex-grow: 1"/>
      </div>
      <div class="skills-item">
        <p style="text-align: start; text-wrap: nowrap">JavaScript / TypeScript</p>
        <ProgressBar :value="80" style="flex-grow: 1"/>
      </div>
      <div class="skills-item">
        <p style="text-align: start; text-wrap: nowrap">GitOps</p>
        <ProgressBar :value="60" style="flex-grow: 1"/>
      </div>
    </div>
    <div class="content-container">
      <h1>Who even am I?</h1>
      <p>
        I'm a 22 years old software developer currently working at the Würth Itensis AG in Chur, Switzerland.
      </p>
      <p style="text-align: start">
        Programming was a hobby of mine since back in my school days.
        I especially like to challenge myself in new technologies, frameworks or programming languages.
      </p>
      <br>
      <p style="text-align: start">
        I'm eager to learn new things about pretty much everything
        that software engineering has to offer. I don't back down from the unknown, and I hope this little website just
        shows you that.
      </p>
      <h2>Working Experience</h2>
    </div>
    <div class="timeline-container">
      <Timeline :value="timelineItems" align="alternate" style="width: inherit">
        <template #marker>
          <span style="height: 2rem; width: 2rem; border-radius: 1rem; background-color: #34d399; display: flex; justify-content: center; align-items: center;">
            <i class="pi pi-crown"></i>
          </span>
        </template>
        <template #content="slotProps">
          <Card style="height: auto; margin-bottom: 5rem" v-animateonscroll="{ enterClass: 'animate-fadein', leaveClass: 'animate-fadeout' }">
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
.content-container {
  width: inherit;
  height: inherit;
  display: flex;
  align-items: start;
  flex-direction: column;
  grid-area: main;
}
.timeline-container {
  width: inherit;
  height: inherit;
  display: flex;
  align-items: center;
  flex-direction: column;
  grid-area: timeline;
  padding-top: 5rem;
}
.skills-container {
  padding-left: 1rem;
  padding-right: 3rem;
  width: inherit;
  height: inherit;
  display: flex;
  align-items: start;
  flex-direction: column;
  grid-area: skills;
}
.skills-item {
  display: flex;
  flex-direction: row;
  margin-bottom: 2rem;
  width: inherit;
  justify-content: space-evenly;
  flex-wrap: wrap;
  gap: 1rem;
}

.main-container {
  height: 100%;
  width: 100%;
  display: grid;
  grid-template-columns: 25% 25% 25% 25%;
  grid-template-rows: 10% 40% 40% 10%;
  grid-template-areas:
    "header header header header"
    "skills main main ."
    "skills timeline timeline ."
    "footer footer footer footer";
  padding: 0.25rem;
  position: relative;
}
p {
  margin: 0;
}
</style>
