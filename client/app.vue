<script setup lang="ts">
const username = ref('');
const password = ref('');
const isWhite = ref(false);
const login = function() {
	const requestBody  = {
		username: username.value,
		password: password.value,
		token: ''
	}
	$fetch('/api/authentication', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Accept': 'application/json',
		},
		body: JSON.stringify(requestBody)
	}).then((data: any) => {
		localStorage.setItem('token', data.token)
		const loginModal = document.getElementById('loginModal');
		if (loginModal) {
			loginModal.close();
			navigateTo('/admin')
		}
	}).catch(err => {
		console.error(err)
	})
}

const toggleTheme = function() {
	console.log(isWhite.value)
	localStorage.setItem("isWhite", String(isWhite.value));
}

onMounted(() => {
	window.addEventListener('keydown', (e) => {
		if (e.ctrlKey && e.altKey && e.key === 'l') {
			document.getElementById('loginModal').showModal();
		}
	});
	console.log("white")
	const isWhiteLocalStorage = localStorage.getItem("isWhite");
	if (isWhiteLocalStorage) {
		isWhite.value = Boolean(isWhiteLocalStorage);
	}
})
</script>
<template>
	<div class="p-3 flex flex-col h-lvh gap-16">
		<div class="navbar bg-base-300 rounded-2xl shadow-2xl">
			<div class="flex-1">
				<a class="btn btn-ghost text-xl" href="/">churrer.xyz</a>
				<a class="btn btn-ghost text-xl" href="/about">about this site</a>
			</div>
			<div class="flex-none p-5">
				<label class="flex cursor-pointer gap-2">
					<span class="label-text">Dark</span>
					<input type="checkbox" value="winter" v-bind:checked="isWhite" v-on:click="toggleTheme" class="toggle theme-controller"/>
					<span class="label-text">Light</span>
				</label>
			</div>
		</div>
		<dialog id="loginModal" class="modal">
			<div class="modal-box p-5">
				<form method="dialog">
					<button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
				</form>
				<div class="flex flex-col align-items-center gap-16">
					<h3 class="text-lg font-bold">Log-In</h3>
					<div>
						<div class="label">
							<span class="label-text">Username</span>
						</div>
						<input class="input input-primary w-11/12" v-model="username"/>
					</div>
					<div>
						<div class="label">
							<span class="label-text">Password</span>
						</div>
						<input type="password" class="input input-primary w-11/12" v-model="password"/>
					</div>
				</div>
				<button class="btn btn-secondary w-24 float-right mt-16 mr-10" v-on:click="login()">Login</button>
			</div>
		</dialog>
		<NuxtRouteAnnouncer/>
		<NuxtPage/>
	</div>
</template>
<style>
html {
	font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;
}
</style>
