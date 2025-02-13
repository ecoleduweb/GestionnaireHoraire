<script lang="ts">
    import { CalendarService } from "../services/calendar.service";
    import { onMount } from "svelte";
    import EventModal from "../Components/Calendar/EventModal.svelte";
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

    async function handleEventSubmit(taskData: Task) {
        if (!selectedDate || !calendarService) return;

        const startDate = new Date(selectedDate.start);

        const [startHours, startMinutes] = taskData.startDateTime
            .toTimeString()
            .split(":");
        startDate.setHours(parseInt(startHours), parseInt(startMinutes), 0);

        const endDate = new Date(startDate);
        const [endHours, endMinutes] = taskData.endDateTime
            .toTimeString()
            .split(":");
        endDate.setHours(parseInt(endHours), parseInt(endMinutes), 0);

        if (endDate < startDate) {
            endDate.setDate(endDate.getDate() + 1);
        }

        if (
            !taskData.name ||
            !taskData.state ||
            !taskData.userId ||
            !taskData.projectId ||
            !taskData.categoryId
        ) {
            alert("Please fill in all required fields");
            return;
        }

        const task: Task = {
            name: taskData.name,
            description: taskData.description || "",
            state: taskData.state,
            billable: taskData.billable || false,
            userId: taskData.userId,
            projectId: taskData.projectId,
            categoryId: taskData.categoryId,
            startDateTime: startDate,
            endDateTime: endDate,
        };

        try {
            const newTask = await TaskApiService.createTask(task);
            calendarService.addEvent({
                id: newTask.id.toString(),
                title: newTask.name,
                start: newTask.startDateTime,
                end: newTask.endDateTime,
                extendedProps: { ...newTask },
            });
        } catch (error) {
            console.error("Error creating task:", error);
            alert("Error creating task");
        }
    }

    async function handleEventUpdate(taskData: Task) {
        if (!calendarService?.calendar) return;

        const startDate = new Date(selectedDate.start);
        const [startHours, startMinutes] = taskData.startDateTime
            .toTimeString()
            .split(":");
        startDate.setHours(parseInt(startHours), parseInt(startMinutes), 0);

        const endDate = new Date(startDate);
        const [endHours, endMinutes] = taskData.endDateTime
            .toTimeString()
            .split(":");
        endDate.setHours(parseInt(endHours), parseInt(endMinutes), 0);

        const updatedTask: Task = {
            id: taskData.id,
            name: taskData.name,
            description: taskData.description || "",
            state: taskData.state,
            billable: taskData.billable || false,
            userId: taskData.userId,
            projectId: taskData.projectId,
            categoryId: taskData.categoryId,
            startDateTime: startDate,
            endDateTime: endDate,
        };

        try {
            const result = await TaskApiService.updateTask(updatedTask);
            const existingEvent = calendarService.calendar.getEventById(
                result.id.toString(),
            );
            if (existingEvent) {
                existingEvent.remove();
                calendarService.addEvent({
                    id: result.id.toString(),
                    title: result.name,
                    start: result.startDateTime,
                    end: result.endDateTime,
                    extendedProps: { ...result },
                });
            }
        } catch (error) {
            console.error("Error updating task:", error);
            alert("Error updating task");
        }
    }

    async function handleEventDelete(taskToDelete: Partial<Task>) {
        if (!calendarService?.calendar) return;

        try {
            await TaskApiService.deleteTask(taskToDelete.id);
            const existingEvent = calendarService.calendar.getEventById(
                taskToDelete.id.toString(),
            );
            if (existingEvent) {
                existingEvent.remove();
            }
        } catch (error) {
            console.error("Error deleting task:", error);
            alert("Error deleting task");
        }
    }
</script>

<div class="calendar-container">
    <div bind:this={calendarEl}></div>
</div>

{#if showModal}
    <EventModal
        show={showModal}
        {users}
        {projects}
        {categories}
        taskToEdit={editTask}
        onDelete={handleEventDelete}
        onClose={() => {
            showModal = false;
        }}
    />
{/if}
