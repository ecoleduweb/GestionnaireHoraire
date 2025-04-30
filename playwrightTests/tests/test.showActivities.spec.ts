import { test, expect } from '@playwright/test';
import { ApiMocker } from '../Helper/mockApi';
import { activityMocks } from '../Helper/Mocks/activity.mock';
import { UserMocks } from '../Helper/Mocks/user.Mock';




test.describe('showActivities', () => {

    test.afterEach(async ({ page }, testInfo) => {
        if (testInfo.status !== testInfo.expectedStatus) {
          // Get a unique place for the screenshot.
          const screenshotPath = testInfo.outputPath(`failure.png`);
          // Add it to the report.
          testInfo.attachments.push({ name: 'screenshot', path: screenshotPath, contentType: 'image/png' });
          // Take the screenshot itself.
          await page.screenshot({ path: screenshotPath, timeout: 5000 });
        }
      });

    test.beforeEach(async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            UserMocks.userMeSuccess
        ])
        .apply();
            await page.clock.install({ time: new Date('2025-03-22T08:00:00-04:00') });         

        
    });

    
    
    test('showActivitiesWeek', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
           activityMocks.getAllActivitiesDefaultWeekSuccess
        ]).apply();
        // Load la page et fait la requête de base 
        await page.goto('http://localhost:5002/calendar');
        await page.waitForSelector('.fc-event', { state: 'visible' });

        // Vérifie les activités de la semaine
        let activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(2);
        await page.getByText('Toute la journée').click();
        // remets le tableau vide
        await page.waitForSelector('.fc-event', { state: 'visible' });
        activities = []; 
        activities = await page.locator('.fc-event').all();
        console.log(activities);
        expect(activities.length).toBe(3);

    });
    

    test('showActivitiesDay', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getAllActivitiesDaySuccess,
            activityMocks.getAllActivitiesDefaultWeekSuccess
        ]).apply();
        // Load la page et fait la requête de base 
        await page.goto('http://localhost:5002/calendar');
        await page.waitForSelector('.fc-event', { state: 'visible' });
        await page.getByRole('button', { name: 'Jour', exact : true }).click();
        // Vérifie les activités de la journee
        let activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(1);
        await page.getByText('Toute la journée').click();
        // remets le tableau vide
        activities = []; 
        activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(2);

    });
    test('showActivitiesMonth', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getAllActivitiesMonthSuccess,
            activityMocks.getAllActivitiesDefaultWeekSuccess
        ]).apply();
        // Load la page et fait la requête de base 
        await page.goto('http://localhost:5002/calendar');
        await page.waitForSelector('.fc-event', { state: 'visible' });
        await page.getByRole('button', { name: 'Mois', exact : true }).click();
        // Vérifie les activités de la mois
        let activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(5);
    
    });
    test('showActivitiesPreviousWeek', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getAllActivitiesPreviousWeekSuccess,
            activityMocks.getAllActivitiesDefaultWeekSuccess
        ]).apply();
        // Load la page et fait la requête de base 
        await page.goto('http://localhost:5002/calendar');
        await page.waitForSelector('.fc-event', { state: 'visible' });
        await page.locator('button:has(.lucide-chevron-left)').click();
        // Vérifie les activités de la semaine
        let activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(1);
    
    });
    test('showActivitiesNextWeek', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getAllActivitiesNextWeekSuccess,
            activityMocks.getAllActivitiesDefaultWeekSuccess
        ]).apply();
        // Load la page et fait la requête de base 
        await page.goto('http://localhost:5002/calendar');
        await page.waitForSelector('.fc-event', { state: 'visible' });
        await page.locator('button:has(.lucide-chevron-right)').click();
        // Vérifie les activités de la semaine
        let activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(1);
    
    });
    test('showActivitiesPreviousMonth', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getAllActivitiesPreviousMonthSuccess,
            activityMocks.getAllActivitiesDefaultWeekSuccess
        ]).apply();
        // Load la page et fait la requête de base 
        await page.goto('http://localhost:5002/calendar');
        await page.waitForSelector('.fc-event', { state: 'visible' });
        await page.getByRole('button', { name: 'Mois', exact : true }).click();
        await page.locator('button:has(.lucide-chevron-left)').click();

        // Vérifie les activités de la mois
        await page.waitForSelector('.fc-event', { state: 'visible' });
        let activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(1);
    
    });
    test('showActivitiesNextMonth', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getAllActivitiesNextMonthSuccess,
            activityMocks.getAllActivitiesDefaultWeekSuccess
        ]).apply();
        // Load la page et fait la requête de base 
        await page.goto('http://localhost:5002/calendar');
        await page.waitForSelector('.fc-event', { state: 'visible' });
        await page.getByRole('button', { name: 'Mois', exact : true }).click();
        await page.locator('button:has(.lucide-chevron-right)').click();

        // Vérifie les activités de la mois
        await page.waitForSelector('.fc-event', { state: 'visible' });
        let activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(1);
    
    });
    test('showActivitiesPreviousDay', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getAllActivitiesPreviousDaySuccess,
            activityMocks.getAllActivitiesDefaultWeekSuccess
        ]).apply();
        // Load la page et fait la requête de base 
        await page.goto('http://localhost:5002/calendar');
        await page.waitForSelector('.fc-event', { state: 'visible' });
        await page.getByRole('button', { name: 'Jour', exact : true }).click();
        await page.locator('button:has(.lucide-chevron-left)').click();
        
        // Vérifie les activités de la journee
        await page.waitForSelector('.fc-event', { state: 'visible' });
        let activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(1);
        await page.getByText('Toute la journée').click();
        // remets le tableau vide
        await page.waitForSelector('.fc-event', { state: 'visible' });
        activities = []; 
        activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(2);
    
    });
    test('showActivitiesNextDay', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getAllActivitiesNextDaySuccess,
            activityMocks.getAllActivitiesDefaultWeekSuccess
        ]).apply();
        // Load la page et fait la requête de base 
        await page.goto('http://localhost:5002/calendar');
        await page.waitForSelector('.fc-event', { state: 'visible' });
        await page.getByRole('button', { name: 'Jour', exact : true }).click();
        await page.locator('button:has(.lucide-chevron-right)').click();
        // Vérifie les activités de la journee
        await page.waitForSelector('.fc-event', { state: 'visible' });
        let activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(1);
        await page.getByText('Toute la journée').click();
        // remets le tableau vide
        await page.waitForSelector('.fc-event', { state: 'visible' });
        activities = []; 
        activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(2);
    
    });
    test('showActivitiesToday', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getAllActivitiesDaySuccess,
            activityMocks.getAllActivitiesDefaultWeekSuccess
        ]).apply();
        // Load la page et fait la requête de base 
        await page.goto('http://localhost:5002/calendar');
        await page.waitForSelector('.fc-event', { state: 'visible' });
        await page.locator('button:has(.lucide-chevron-right)').click();
        await page.locator('button:has(.lucide-chevron-right)').click();
        
        let activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(0);
    
        await page.getByRole('button', { name: 'Aujourd\'hui', exact : true }).click();
        await page.getByRole('button', { name: 'Jour', exact : true }).click();
        // Vérifie les activités de la journee
        await page.waitForSelector('.fc-event', { state: 'visible' });
        activities = [];
        activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(1);
        await page.getByText('Toute la journée').click();
        // remets le tableau vide
        await page.waitForSelector('.fc-event', { state: 'visible' });
        activities = []; 
        activities = await page.locator('.fc-event').all();
        expect(activities.length).toBe(2);
    });



});