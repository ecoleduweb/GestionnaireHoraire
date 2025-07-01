import { test, expect } from '@playwright/test';
import { ApiMocker } from '../Helper/mockApi';
import { projectMocks } from '../Helper/Mocks/project.mock';
import { userMocks } from '../Helper/Mocks/user.Mock';
import {activityMocks} from '../Helper/Mocks/activity.mock';

test.describe('checkRemainingHoursColor', () => {

    test.beforeEach(async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            userMocks.userMeSuccess,
            activityMocks.getAllActivitiesDefaultWeekSuccess,
            projectMocks.getProjectsListSuccess,
            projectMocks.getDetailedProjectsByUserSuccess,

        ])
            .apply();
        await page.clock.install({ time: new Date('2025-03-22T08:00:00') });
        await page.goto('http://localhost:5002/calendar');
        await page.waitForLoadState('networkidle');
    });

    test('negativeRemainingHours', async ({ page }) => {
        
        await page.waitForSelector(".text-red-700", { state: "visible" });
        await expect(
          page.locator(".text-red-700").getByText("-5h00")
        ).toBeVisible();
    });
    test('positiveRemainingHours', async ({ page }) => {
        await page.waitForSelector(".text-gray-700", { state: "visible" });
        await expect(
          page.locator(".text-gray-700").getByText("17h30")
        ).toBeVisible();        
    });

    test('noEstimatedHours', async ({ page }) => {
        await page.waitForSelector(".text-gray-400", { state: "visible" });
        await expect(
          page.locator(".text-gray-400").getByText("-1h00")
        ).toBeVisible();
    });

    test('noRemainingHasEstimate', async ({ page }) => {
        await page.waitForSelector(".text-gray-700", { state: "visible" });
        await expect(
          page.locator(".text-gray-700").getByText("-")
        ).toBeVisible();
    });


});