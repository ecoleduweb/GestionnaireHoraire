import { test, expect } from '@playwright/test';
import { ApiMocker } from '../Helper/mockApi';
import { activityMocks } from '../Helper/Mocks/activity.mock';




test.describe('checkAddActivity', () => {

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

    test('ajouterUneActivite', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.addActivitySuccess,
        ]).apply();
        
        
        await page.getByText('Nouvelle activité').click();
        await page.getByPlaceholder('Nom de l\'activité...').fill('asd');
        await page.locator("#activity-description").fill('asd');
        await page.locator('#activity-project').selectOption('1');

        await page.getByText ('Créer').click();
        
        await expect(page.getByText('Fête de Jean-Félix et Sherlock')).toBeVisible();
        

    });

    test('ajouterUneActiviteSansNomDescription', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.addActivitySuccessNoNameNoDescription,
        ]).apply();
        
        
        await page.getByText('Nouvelle activité').click();
        await page.locator('#activity-project').selectOption('1');
        await page.getByText ('Créer').click();

        await expect(page.locator('.fc-event-title-container')).toBeVisible(); //Valide l'existence de l'activité (Sans contenir de nom)
    });


});