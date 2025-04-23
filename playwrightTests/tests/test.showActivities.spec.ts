import { test, expect } from '@playwright/test';
import { ApiMocker } from '../Helper/mockApi';
import { activityMocks } from '../Helper/Mocks/activity.mock';




test.describe('showActivities', () => {

    test.beforeEach(async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getByIdSuccess,      
            activityMocks.getAllActivityEmpty,      
        ])
            .apply();
        await page.clock.install({ time: new Date('2025-03-22T08:00:00') });
        await page.goto('http://localhost:5002/calendar');
        await page.waitForLoadState('networkidle');
    });

    test('showActivitiesWeek', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.addActivitySuccess,
        ]).apply();
        
        

    });


});