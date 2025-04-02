import { test, expect } from '@playwright/test';
import { ApiMocker } from '.././Helper/mockApi';
import { loginMocks } from '../Helper/Mocks/login.mock';



test.describe('checkIfOffersHide', () => {

    test.beforeEach(async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            loginMocks.successModerator])
            .apply();

        await page.goto('http://localhost:5002/login');
        await page.waitForLoadState('networkidle');
        await page.getByLabel('Nom d\'utilisateur').fill('test@gmail.com');
        await page.getByLabel('Mot de passe').fill('test');
        await page.getByRole('button', { name: 'Se connecter' }).click();
    });

    test('nomDuTest', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
        ]).apply();

        await page.goto('http://localhost:5002/dashboard');
        await page.waitForLoadState('networkidle');
        if (await page.locator("#cookieBannerOk")) {
            await page.locator("#cookieBannerOk").click()
        }

    });


});