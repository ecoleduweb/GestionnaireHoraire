import { test, expect } from '@playwright/test';
import { ApiMocker } from '../Helper/mockApi';
import { projectMocks } from '../Helper/Mocks/project.mock';
import { userMocks } from '../Helper/Mocks/user.Mock';

test.describe('checkAddProjects', () => {

    test.beforeEach(async ({ page }) => {
        const apiMocker = new ApiMocker(page);
        await apiMocker.addMocks([
            projectMocks.getDetailedProjectsSuccess,
            userMocks.userMeSuccess,
            userMocks.getAllManagersSuccess,
            projectMocks.addProjectSuccess,
        ])
            .apply();
        await page.clock.install({ time: new Date('2025-03-22T08:00:00') });
        await page.goto('http://localhost:5002/projects');
        await page.waitForLoadState('networkidle');
    });

    test('AddProjectToManager', async ({ page }) => {
        await page.getByTitle('Créer un nouveau projet').click();
        await page.locator('#project-name').fill('Nouveau Projet');
        await page.locator('#project-manager').selectOption('2'); 
        await page.locator('#project-description').fill('Description du nouveau projet');
        await page.getByText('Soumettre').click();




    });
    test('AddProjectToAdmin', async ({ page }) => {
               
    });

    test('AddProjectInvalidModal', async ({ page }) => {
               
    });

    test('NotAllowedToAddProject', async ({ page }) => {
               
    });


});