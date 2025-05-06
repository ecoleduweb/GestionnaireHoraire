import { test, expect } from '@playwright/test';
import { ApiMocker } from '../Helper/mockApi';
import { activityMocks } from '../Helper/Mocks/activity.mock';
import { projectMocks } from '../Helper/Mocks/project.mock';
import { UserMocks } from '../Helper/Mocks/user.Mock';
import { categoryMocks } from '../Helper/Mocks/category.mock';

test.describe('checkAddActivity', () => {
    test.beforeEach(async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.getByIdSuccess,      
            activityMocks.getAllActivityEmpty,    
            projectMocks.getProjectsSuccess,
            UserMocks.userMeSuccess,
            categoryMocks.getCategoriesByProjectSuccess 
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
            activityMocks.getAllActivitiesAfterAdd
        ]).apply();
        
        // Ouvrir la modale
        await page.getByText('Nouvelle activité').click();
        
        // Attendre que le formulaire soit chargé
        await page.waitForSelector('#activity-project');
        
        // Remplir le formulaire
        await page.getByPlaceholder('Nom de l\'activité...').fill('asd');
        await page.locator("#activity-description").fill('asd');
        await page.locator('#activity-project').selectOption({index: 0});
        
        // Cliquer sur Créer sans attendre la réponse
        await page.getByText('Créer').click();
        
        // Attendre que l'événement apparaisse dans le calendrier
        await page.waitForSelector('.fc-event', { state: 'visible', timeout: 10000 });
        
        // Vérifier que l'activité apparaît dans le calendrier
        await expect(page.locator('a').filter({ hasText: 'Fête de Jean-Félix et Sherlock' })).toBeVisible();
    });

    test('ajouterUneActiviteSansNomDescription', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            activityMocks.addActivitySuccessNoNameNoDescription,
            activityMocks.getAllActivitiesAfterAddNoName
        ]).apply();
        
        // Ouvrir la modale
        await page.getByText('Nouvelle activité').click();
        
        // Attendre que le formulaire soit chargé
        await page.waitForSelector('#activity-project');
        
        // Sélectionner un projet
        await page.locator('#activity-project').selectOption({index: 0});
        
        // Cliquer sur Créer sans attendre la réponse
        await page.getByText('Créer').click();
        
        // Vérification de l'élément dans le calendrier
        await expect(page.locator('.fc-event-title-container')).toBeVisible({ timeout: 10000 });
    });
});