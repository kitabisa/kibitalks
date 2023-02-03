export default class Page {

  async open(path) {
    await browser.url(path);
  }
}
