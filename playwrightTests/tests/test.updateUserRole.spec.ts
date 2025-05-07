import { test, expect } from '@playwright/test';
import { ApiMocker } from '../Helper/mockApi';
import { UserMocks } from '../Helper/Mocks/user.Mock';

test.describe('updateUserRole', () => {

    test.beforeEach(async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            UserMocks.usersSuccess,
            UserMocks.userMeSuccess,
        ])
            .apply();
        await page.clock.install({ time: new Date('2025-03-22T08:00:00') });
        await page.goto('http://localhost:5002/users');
        await page.waitForLoadState('networkidle');
    });

    test('updateRoleSuccess', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            UserMocks.updateUserRoleSuccess,
            UserMocks.usersSuccess,
        ]).apply();
        
        await page.locator('#userSelect').selectOption('1');
        await page.locator('#roleSelect').selectOption('1');
        
        await page.getByText('Confirmer').click();
        
        await expect(page.locator('text=Rôle mis à jour avec succès')).toBeVisible();
    });

    test('updateRoleError', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            UserMocks.updateUserRoleError,
        ]).apply();
        
        await page.locator('#userSelect').selectOption('1');
        await page.locator('#roleSelect').selectOption('1');
        
        await page.getByText('Confirmer').click();
        
        await expect(page.locator('text=Erreur lors de la mise à jour du rôle')).toBeVisible();
    });
});