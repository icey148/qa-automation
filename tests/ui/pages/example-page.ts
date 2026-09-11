import { Page } from '@playwright/test';
import { Modal } from '../framework/components/modal';

/**
 * Example Page Object.
 * Replace or extend this with real application pages generated from frontend evidence.
 */
export class ExamplePage {
  readonly modal: Modal;

  constructor(private readonly page: Page) {
    this.modal = new Modal(page);
  }

  async open(path = '/') {
    await this.page.goto(path);
  }

  actionButton(name: string | RegExp) {
    return this.page.getByRole('button', { name });
  }

  field(label: string | RegExp) {
    return this.page.getByLabel(label);
  }
}
