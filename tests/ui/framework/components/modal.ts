import { Page } from '@playwright/test';

export class Modal {
  constructor(private readonly page: Page) {}

  async confirm(name: string | RegExp = /ok|confirm|save/i) {
    await this.page.getByRole('button', { name }).click();
  }

  async cancel(name: string | RegExp = /cancel/i) {
    await this.page.getByRole('button', { name }).click();
  }
}
