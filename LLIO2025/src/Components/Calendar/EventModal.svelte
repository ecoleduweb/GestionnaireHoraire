<script lang="ts">
    import type { Task, User, Project, Category } from "../../Models";
    import { TASK_CONFIG } from "../../config/task.config";

    export let show = false;
    export let users: User[] = [];
    export let projects: Project[] = [];
    export let categories: Category[] = [];
    export let editMode = false;
    export let editTask: Partial<Task> | null = null;

    export let onSubmit: (task: Partial<Task>) => void;
    export let onClose: () => void;
    export let onUpdate: (task: Partial<Task>) => void;
    export let onDelete: (task: Partial<Task>) => void;

    let startHours = "00";
    let startMinutes = "00";
    let endHours = "00";
    let endMinutes = "00";

    let task: Partial<Task> = {
        ...TASK_CONFIG.defaultTask,
    };

    $: if (editMode && editTask) {
        task = { ...editTask };
        const startDate = new Date(editTask.startDate);
        const endDate = new Date(editTask.endDate);

        startHours = startDate.getHours().toString().padStart(2, "0");
        startMinutes = startDate.getMinutes().toString().padStart(2, "0");
        endHours = endDate.getHours().toString().padStart(2, "0");
        endMinutes = endDate.getMinutes().toString().padStart(2, "0");
    }

    const {
        states,
        time: { hours, minutes },
    } = TASK_CONFIG;

    function handleSubmit() {
        if (task.name && task.userId && task.projectId && task.categoryId) {
            const taskWithtime = {
                ...task,
                startTime: `${startHours}:${startMinutes}`,
                endTime: `${endHours}:${endMinutes}`,
            };
            if (editMode) {
                onUpdate(taskWithtime);
            } else {
                onSubmit(taskWithtime);
            }
            resetForm();
        }
    }

    function handleDelete() {
        if (confirm("Supprimer cette tâche ?")) {
            onDelete(task);
            resetForm();
        }
    }

    function resetForm() {
        task = { ...TASK_CONFIG.defaultTask };
        startHours = "00";
        startMinutes = "00";
        endHours = "00";
        endMinutes = "00";
        show = false;
        editMode = false;
        editTask = null;
    }

    function handleClose() {
        resetForm();
        onClose();
    }

    $: {
        if (
            parseInt(endHours) < parseInt(startHours) ||
            (parseInt(endHours) === parseInt(startHours) &&
                parseInt(endMinutes) < parseInt(startMinutes))
        ) {
            endHours = startHours;
            endMinutes = startMinutes;
        }
    }
</script>

{#if show}
    <div class="modal-overlay" on:click={handleClose}>
        <div class="modal" on:click|stopPropagation>
            <h2>{editMode ? "Modifier la tâche" : "Nouvelle tâche"}</h2>
            <form on:submit|preventDefault={handleSubmit}>
                <div class="form-group">
                    <label for="task-name">Nom*</label>
                    <input
                        id="task-name"
                        type="text"
                        bind:value={task.name}
                        placeholder="Nom de la tâche..."
                        required
                        autofocus
                    />
                </div>

                <div class="form-group">
                    <label for="task-description">Description</label>
                    <textarea
                        id="task-description"
                        bind:value={task.description}
                        rows="3"
                    ></textarea>
                </div>

                <div class="form-row">
                    <div class="form-group">
                        <label for="task-state">État*</label>
                        <select
                            id="task-state"
                            bind:value={task.state}
                            required
                        >
                            {#each states as state}
                                <option value={state}>{state}</option>
                            {/each}
                        </select>
                    </div>

                    <div class="form-group checkbox">
                        <label>
                            <input
                                type="checkbox"
                                bind:checked={task.billable}
                            />
                            Facturable
                        </label>
                    </div>
                </div>

                <div class="form-group time-inputs">
                    <div class="time-select">
                        <label>Heure de début*</label>
                        <div class="time-pickers">
                            <select bind:value={startHours}>
                                {#each hours as hour}
                                    <option value={hour}>{hour}h</option>
                                {/each}
                            </select>
                            <select bind:value={startMinutes}>
                                {#each minutes as minute}
                                    <option value={minute}>{minute}</option>
                                {/each}
                            </select>
                        </div>
                    </div>
                    <div class="time-select">
                        <label>Heure de fin*</label>
                        <div class="time-pickers">
                            <select bind:value={endHours}>
                                {#each hours as hour}
                                    <option value={hour}>{hour}h</option>
                                {/each}
                            </select>
                            <select bind:value={endMinutes}>
                                {#each minutes as minute}
                                    <option value={minute}>{minute}</option>
                                {/each}
                            </select>
                        </div>
                    </div>
                </div>

                <div class="form-row">
                    <div class="form-group">
                        <label for="task-user">Utilisateur*</label>
                        <select
                            id="task-user"
                            bind:value={task.userId}
                            required
                        >
                            <option value="">Sélectionner...</option>
                            {#each users as user}
                                <option value={user.id}>{user.name}</option>
                            {/each}
                        </select>
                    </div>

                    <div class="form-group">
                        <label for="task-project">Projet*</label>
                        <select
                            id="task-project"
                            bind:value={task.projectId}
                            required
                        >
                            <option value="">Sélectionner...</option>
                            {#each projects as project}
                                <option value={project.id}
                                    >{project.name}</option
                                >
                            {/each}
                        </select>
                    </div>

                    <div class="form-group">
                        <label for="task-category">Catégorie*</label>
                        <select
                            id="task-category"
                            bind:value={task.categoryId}
                            required
                        >
                            <option value="">Sélectionner...</option>
                            {#each categories as category}
                                <option value={category.id}
                                    >{category.name}</option
                                >
                            {/each}
                        </select>
                    </div>
                </div>

                <div class="button-group">
                    {#if editMode}
                        <button
                            type="button"
                            class="btn-danger"
                            on:click={handleDelete}
                        >
                            Supprimer
                        </button>
                        <button type="submit" class="btn-primary">
                            Modifier
                        </button>
                    {:else}
                        <button type="submit" class="btn-primary">
                            Créer
                        </button>
                    {/if}
                    <button
                        type="button"
                        class="btn-secondary"
                        on:click={handleClose}
                    >
                        Annuler
                    </button>
                </div>
            </form>
        </div>
    </div>
{/if}

<style>
    .modal-overlay {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, 0.5);
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 1000;
    }

    .modal {
        background: white;
        padding: 2rem;
        border-radius: 8px;
        width: 90%;
        max-width: 500px;
    }

    h2 {
        margin: 0 0 1.5rem 0;
        color: #333;
        font-size: 1.5rem;
    }

    .form-group {
        margin-bottom: 1.5rem;
    }

    .time-inputs {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 1rem;
    }

    .time-select {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }

    .time-pickers {
        display: flex;
        gap: 0.5rem;
    }

    .checkbox {
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }

    .checkbox input {
        width: auto;
    }

    label {
        display: block;
        margin-bottom: 0.5rem;
        color: #666;
    }

    input {
        width: 100%;
        padding: 0.75rem;
        border: 1px solid #ddd;
        border-radius: 4px;
        font-size: 1rem;
    }

    .button-group {
        display: flex;
        gap: 1rem;
        justify-content: flex-end;
    }

    button {
        padding: 0.75rem 1.5rem;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 1rem;
    }

    .btn-primary {
        background: #1a1a1a;
        color: white;
    }

    .btn-primary:hover {
        background: #333;
    }

    .btn-secondary {
        background: #e0e0e0;
        color: #333;
    }

    .btn-secondary:hover {
        background: #ccc;
    }

    select {
        padding: 0.5rem;
        border: 1px solid #ddd;
        border-radius: 4px;
        background: white;
    }
</style>
