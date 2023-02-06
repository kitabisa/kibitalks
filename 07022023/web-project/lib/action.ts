const { waitFor, ExpectedConditions } = require('webdriverio-explicit-waits')
const Expected = ExpectedConditions

class action {

  async expectToExist(selector: any) {
    waitFor(Expected.visibilityOf(()=> selector), 5000);
    await expect(selector).toExist();
  }

  async clickOn(selector: any) {
    waitFor(Expected.visibilityOf(()=> selector), 5000);
    await selector.click();
  }

  async clickOnIndex(selector: any, index: number) {
    waitFor(Expected.visibilityOf(()=> selector), 5000);
    await expect(selector).toExist();
    await selector[index].click();
  }

  async scrollInto(selector: any) {
    waitFor(Expected.visibilityOf(()=> selector), 5000);
    await selector.scrollIntoView();
  }

  async setValue(selector: any, value: any) {
    waitFor(Expected.visibilityOf(()=> selector), 5000);
    await selector.setValue(value);
  }

  async expectToHaveUrlContaining(value: any) {
    waitFor(Expected.visibilityOf(()=> browser), 5000);
    await browser.pause(800);
    expect(browser).toHaveUrlContaining(value);
  }

}

export default new action()