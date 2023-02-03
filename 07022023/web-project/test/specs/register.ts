import action from '../../lib/action';
import checkpointPage from '../pageobjects/checkpoint.page';
import signupPage from '../pageobjects/signup.page';
import data from '../support/data';

describe('Register Facebook', function () {
  this.retries(1);

  beforeEach( async() => {
    await signupPage.open();
  });

  it('NEGATIVE - empty all field', async () => {
    await action.clickOn(signupPage.buttonSignUp);
    await action.expectToExist(signupPage.errorEmptyName);
    await action.clickOn(signupPage.inputSurename);
    await action.expectToExist(signupPage.errorEmptyName);
    await action.clickOn(signupPage.inputNumberEmail);
    await action.expectToExist(signupPage.errorEmptyNumberEmail);
    await action.clickOn(signupPage.inputPassword);
    await action.expectToExist(signupPage.errorEmptyNewPassword);
  });

  it('POSITIVE - successfully registered with email', async () => {
    await action.setValue(signupPage.inputFirstName, data.validData.firstName);
    await action.setValue(signupPage.inputSurename, data.validData.surename);
    await action.setValue(signupPage.inputNumberEmail, data.validData.email);
    await action.setValue(signupPage.inputReEnterEmail, data.validData.email);
    await action.setValue(signupPage.inputPassword, data.validData.newPass);
    await action.clickOn(signupPage.inputDate);
    await signupPage.chooseDob(data.validData.dobDate);
    await action.clickOn(signupPage.inputMonth);
    await signupPage.chooseDob(data.validData.dobMonth);
    await action.clickOn(signupPage.inputYear);
    await signupPage.chooseDob(data.validData.dobYear);
    await action.clickOnIndex(signupPage.inputGender, 0);
    await action.clickOn(signupPage.buttonSignUp);
    await action.expectToHaveUrlContaining(data.validData.registerUrlSuccess);
  });
});


