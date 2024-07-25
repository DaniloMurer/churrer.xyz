<script setup lang="ts">

</script>

<template>
  <ul class="timeline timeline-vertical timeline-snap-icon">
    <li>
      <div class="timeline-start mb-10">
        <WorkCard class="lg:mr-10" data-aos="fade-right" data-aos-duration="1000" title="Apprentice" company="Würth ITensis" timeframe="01.08.2018 - 31.07.2022" text="Developing as well as maintaining Java applications
      and ABAP / SAPUI5 development"/>
      </div>
      <div class="timeline-middle">
        <span class="iconify carbon--time text-xl"></span>
      </div>
      <hr/>
    </li>
    <li>
      <div class="timeline-end mb-10">
        <WorkCard class="lg:ml-10"
                  data-aos="fade-left"
                  data-aos-duration="1000"
                  title="Junior Software Engineer"
                  company="Würth ITensis"
                  timeframe="01.08.2022 - Current"
                  text="Migrating existing apps to OpenShift Platform. Engineering and maintaining k8s DevOps infrastructure. Developing and maintaining Java applications with Docker. Developing and maintaining SAPUI5 and ABAP applications. Requirements-Engineering with clients"/>
      </div>
      <div class="timeline-middle">
        <span class="iconify carbon--time text-xl"></span>
      </div>
    </li>
  </ul>
</template>

<style scoped>

</style>