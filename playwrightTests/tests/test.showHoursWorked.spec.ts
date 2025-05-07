import { test, expect } from '@playwright/test';
import { ApiMocker } from '../Helper/mockApi';
import { activityMocks } from '../Helper/Mocks/activity.mock';
import { projectMocks } from '../Helper/Mocks/project.mock';




test.describe('showHoursWorked', () => {

    test.beforeEach(async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getByIdSuccess,      
            activityMocks.getAllActivityEmpty,    
            projectMocks.getProjectsSuccess,  
        ])
            .apply();
        await page.clock.install({ time: new Date('2025-03-22T08:00:00') });
        await page.goto('http://localhost:5002/calendar');
        await page.waitForLoadState('networkidle');
    });




});