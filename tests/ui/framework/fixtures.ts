import { test as base, expect, Page } from '@playwright/test';

export type QAFixtures = {
  authenticatedPage: Page;
};

export const test = base.extend<QAFixtures>({
  authenticatedPage: async ({ browser }, use) => {
    const storageState = process.env.PLAYWRIGHT_STORAGE_STATE;
    const context = await browser.newContext(
      storageState ? { storageState } : undefined,
    );
    const page = await context.newPage();
    await use(page);
    await context.close();
  },
});

export { expect };
