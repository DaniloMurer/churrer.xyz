<script setup lang="ts">
import {useVisibleStore} from "../store/visible.ts";

const visibleStore = useVisibleStore();
</script>

<template>
    <Toolbar style="padding: 0.5rem;">
        <template #start>
            <div class="flex-container-hor" style="margin-left: 0.5rem;">
              <span class="pi pi-crown"></span>
              <h3 style="word-break: break-all;">churrer.xyz</h3>
            </div>
        </template>
        <template #center>
            <IconField>
                <InputIcon>
                    <span class="pi pi-search"></span>
                </InputIcon>
                <InputText placeholder="What ya lookin' for?"/>
            </IconField>
        </template>
        <template #end>
            <div>
              <Button severity="help" label="Create Card" @click="visibleStore.visible = true"/>
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