<script setup lang="ts">

import {definePageMeta} from "#imports";

definePageMeta({
	middleware: ['auth'],
})

const experience = ref({
	company: '',
	position: '',
	timeFrame: '',
	responsibilities: ''
});

const technology = ref({
	name: '',
	experience: '',
	description: '',
	logoClass: ''
});

let authenticationToken: string | null = '';

onMounted(() => {
	authenticationToken = localStorage.getItem("token");
})

const saveExperience = function() {
	$fetch('/api/experiences', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Accept': 'application/json',
			'Authorization': `Basic ${authenticationToken}`
		},
		body: JSON.stringify(experience.value)
	});
}

const saveTechnology = function() {
	$fetch('/api/technologies', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Accept': 'application/json',
			'Authorization': `Basic ${authenticationToken}`
		},
		body: JSON.stringify(technology.value)
	})
}
</script>

<template>
	<h1 class="text-2xl">Welcome to the admin page</h1>

	<div class="flex flex-col items-center gap-10">
		<h2 class="font-bold text-4xl">Experience</h2>
		<div>
			<div class="label">
				<span class="label-text">Company</span>
			</div>
			<input type="text" class="input input-primary" v-model="experience.company"/>
		</div>

		<div>
			<div class="label">
				<span class="label-text">Position</span>
			</div>
			<input type="text" class="input input-primary" v-model="experience.position"/>
		</div>

		<div>
			<div class="label">
				<span class="label-text">TimeFrame</span>
			</div>
			<input type="text" class="input input-primary" v-model="experience.timeFrame"/>
		</div>

		<div>
			<div class="label">
				<span class="label-text">Responsibilities</span>
			</div>
			<input type="text" class="input input-primary" v-model="experience.responsibilities"/>
		</div>
		<button class="btn btn-secondary" v-on:click="saveExperience">Save</button>
	</div>
	<hr/>
	<div class="flex flex-col items-center gap-10">
		<h2 class="font-bold text-4xl">Technology</h2>
		<div>
			<div class="label">
				<span class="label-text">Name</span>
			</div>
			<input type="text" class="input input-primary" v-model="technology.name"/>
		</div>

		<div>
			<div class="label">
				<span class="label-text">Experience</span>
			</div>
			<input type="text" class="input input-primary" v-model="technology.experience"/>
		</div>

		<div>
			<div class="label">
				<span class="label-text">Description</span>
			</div>
			<input type="text" class="input input-primary" v-model="technology.description"/>
		</div>

		<div>
			<div class="label">
				<span class="label-text">Tailwind Logo</span>
			</div>
			<input type="text" class="input input-primary" v-model="technology.logoClass"/>
		</div>
		<button class="btn btn-secondary" v-on:click="saveTechnology">Save</button>
	</div>
</template>
