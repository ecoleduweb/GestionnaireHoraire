<script lang="ts">
    import { CalendarService } from "../services/calendar.service";
    import { onMount } from "svelte";
    import TaskModal from "../Components/Calendar/TaskModal.svelte";
    import "./calendar.css";
    import { TaskApiService } from "../services/TaskApiService";
    import type { Task } from "../Models";

    let calendarEl = $state<HTMLElement | null>(null);
    let calendarService = $state<CalendarService | null>(null);
    let showModal = $state(false);
    let selectedDate: any = null;
    let editMode = $state(false);
    let editTask = $state(null);

    const users = [];
    const projects = [];
    const categories = [];

    async function loadTasks() {
        try {
            const tasks = await TaskApiService.getTasks();
            tasks.forEach((task) => {
                calendarService?.addEvent({
                    title: task.name,
                    start: task.startDateTime,
                    end: task.endDateTime,
                    extendedProps: { ...task },
                });
            });
        } catch (error) {
            console.error("Error loading tasks", error);
        }
    }

    onMount(() => {
        if (calendarEl) {
            calendarService = new CalendarService();

            calendarService.onDateSelect = (info) => {
                editMode = false;
                editTask = null;
                selectedDate = info;
                showModal = true;
            };

            calendarService.onEventClick = (info) => {
                editMode = true;
                editTask = {
                    id: info.event.extendedProps.id,
                    name: info.event.title,
                    description: info.event.extendedProps.description,
                    state: info.event.extendedProps.state,
                    billable: info.event.extendedProps.billable,
                    userId: info.event.extendedProps.userId,
                    projectId: info.event.extendedProps.projectId,
                    categoryId: info.event.extendedProps.categoryId,
                    startDate: info.event.start,
                    endDate: info.event.end,
                };
                selectedDate = {
                    start: info.event.start,
                    end: info.event.end,
                };
                showModal = true;
            };

            calendarService.initialize(calendarEl);
            calendarService.render();
            loadTasks();
        }
    });

    async function handleTaskSubmit(taskData: Task) {
        if (!calendarService) return;

        calendarService.addEvent({
            id: taskData.id.toString(),
            title: taskData.name,
            start: taskData.startDateTime,
            end: taskData.endDateTime,
            extendedProps: { ...taskData },
        });
    }

    async function handleTaskUpdate(task: Task) {
        if (!calendarService?.calendar) return;
        calendarService.updateEvent(task);
    }

    async function handleTaskDelete(task: Task) {
        if (!calendarService?.calendar || !task.id) return;
        calendarService.deleteTask(task.id.toString());
    }
</script>

<div class="calendar-container">
    <div bind:this={calendarEl}></div>
</div>

{#if showModal}
    <TaskModal
        show={showModal}
        {users}
        {projects}
        {categories}
        taskToEdit={editTask}
        onDelete={handleTaskDelete}
        onSubmit={handleTaskSubmit}
        onUpdate={handleTaskUpdate}
        onClose={() => {
            showModal = false;
        }}
    />
{/if}

<style>
</style>
