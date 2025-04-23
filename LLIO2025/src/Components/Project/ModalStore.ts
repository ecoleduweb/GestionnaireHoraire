import { writable } from 'svelte/store';

export const modalOuverte = writable<boolean>(false);