import { test, expect } from '@playwright/test';
import { ApiMocker } from '../Helper/mockApi';
import { projectMocks } from '../Helper/Mocks/project.mock';
import { userMocks } from '../Helper/Mocks/user.Mock';
import { activityMocks } from '../Helper/Mocks/activity.mock';

test.describe('logOut', () => {

    test.beforeEach(async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getAllActivitiesDefaultWeekSuccess,
            userMocks.userMeSuccess,
        ])
            .apply();
        await page.clock.install({ time: new Date('2025-03-22T08:00:00') });
        await page.goto('http://localhost:5002/calendar');
        await page.waitForLoadState('networkidle');
    });

    test('logoutSucces', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            userMocks.logOutSuccess,
        ])
            .apply();
        
        await page.getByTitle('Se déconnecter').click();
        await expect(page.getByText('Se connecter avec Microsoft')).toBeVisible();

    });
});