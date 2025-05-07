import { Page } from '@playwright/test';
import { MockConfig } from './types';

export class ApiMocker {
    private page: Page;
    private mockConfigs: MockConfig[] = [];

    constructor(page: Page) {
        this.page = page;
    }

    // Chainable method to add a single mock configuration
    addMock(config: MockConfig): this {
        this.mockConfigs.push(config);
        return this;
    }

    // Chainable method to add multiple mock configurations
    addMocks(configs: MockConfig[]): this {
        this.mockConfigs.push(...configs);
        return this;
    }

    // Method to apply all collected mock configurations
    async apply(): Promise<void> {
        for (const config of this.mockConfigs) {
            await this.page.route(`http://localhost:8080${config.url}`, async route => {
                await route.fulfill({
                    status: config.response.status,
                    json: config.response.json,
                    headers: config.response.headers
                });
            });
            await this.page.route(`*/**api${config.url}`, async route => {
                await route.fulfill({
                    status: config.response.status,
                    json: config.response.json,
                    headers: config.response.headers
                });
            });
        }

        // Reset mock configurations after application
        this.mockConfigs = [];
    }

    // Optional method to clear all mock routes
    async clearMocks(): Promise<void> {
        await this.page.unroute('**');
        this.mockConfigs = [];
    }
}

