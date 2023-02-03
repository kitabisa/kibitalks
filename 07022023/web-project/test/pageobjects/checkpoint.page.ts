import Page from "./page";

class recoverPage extends Page{

  get buttonContinue() { return $("//*[text()='Continue']"); }
  
}

export default new recoverPage();