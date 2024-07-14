<script setup lang="ts">
import {useVisibleStore} from "../store/visible.ts";

const visibleStore = useVisibleStore();
</script>

<template>
    <Toolbar>
        <template #center>
          <div class="flex-container-hor" style="margin-left: 0.5rem;">
            <span class="pi pi-crown"></span>
            <h3 style="word-break: break-all;">churrer.xyz</h3>
          </div>
        </template>
    </Toolbar>
</template>
<style>
.flex-container-hor {
  display: flex;
  align-items: center;
  flex-wrap: nowrap;
  gap: 1rem;
  width: 10rem;
}
</style>