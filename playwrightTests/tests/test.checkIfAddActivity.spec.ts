import { test, expect } from '@playwright/test';
import { ApiMocker } from '../Helper/mockApi';
import { activityMocks } from '../Helper/Mocks/activity.mock';
import { projectMocks } from '../Helper/Mocks/project.mock';




test.describe('checkAddActivity', () => {
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

    test('ajouterUneActivite', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        let requestSent = false;
        let responseReceived = false;
        
        page.on('request', request => {
            if (request.url().includes('/activity') && request.method() === 'POST') {
                requestSent = true;
            }
        });
        
        page.on('response', response => {
            if (response.url().includes('/activity') && response.request().method() === 'POST') {
                responseReceived = true;
            }
        });
        
        await apiMocker.addMocks([
            activityMocks.addActivitySuccess,
            activityMocks.getAllActivitiesAfterAdd
        ]).apply();
        
        // Ouvrir la modale et remplir le formulaire
        await page.getByText('Nouvelle activité').click();
        await page.waitForSelector('#activity-project');
        await page.getByPlaceholder('Nom de l\'activité...').fill('asd');
        await page.locator("#activity-description").fill('asd');
        await page.locator('#activity-project').selectOption({index: 1});
        
        // Soumettre le formulaire et attendre la réponse
        await Promise.all([
            page.getByText('Créer').click(),
            page.waitForResponse(response => 
                response.url().includes('/activity') && 
                response.request().method() === 'POST'
            )
        ]);
        
        // Vérifier que la requête a été envoyée et traitée
        expect(requestSent).toBe(true);
        expect(responseReceived).toBe(true);
        
        // Vérifier l'affichage dans le calendrier
        await page.waitForSelector('.fc-event', { state: 'visible' });
        await expect(page.locator('a').filter({ hasText: '08:00 - 09:00Fête de Jean-Fé' })).toBeVisible();
    });

    test('ajouterUneActiviteSansNomDescription', async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        let requestSent = false;
        let responseReceived = false;
        
        page.on('request', request => {
            if (request.url().includes('/activity') && request.method() === 'POST') {
                requestSent = true;
            }
        });
        
        page.on('response', response => {
            if (response.url().includes('/activity') && response.request().method() === 'POST') {
                responseReceived = true;
            }
        });
        
        await apiMocker.addMocks([
            activityMocks.addActivitySuccessNoNameNoDescription,
            activityMocks.getAllActivitiesAfterAddNoName
        ]).apply();
        
        // Ouvrir la modale et sélectionner un projet
        await page.getByText('Nouvelle activité').click();
        await page.waitForSelector('#activity-project');
        await page.locator('#activity-project').selectOption({index: 1});
        
        // Soumettre le formulaire et attendre la réponse
        await Promise.all([
            page.getByText('Créer').click(),
            page.waitForResponse(response => 
                response.url().includes('/activity') && 
                response.request().method() === 'POST'
            )
        ]);
        
        // Vérifier que la requête a été envoyée et traitée
        expect(requestSent).toBe(true);
        expect(responseReceived).toBe(true);
    });
});