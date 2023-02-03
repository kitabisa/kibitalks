import Page from './page';
import action from '../../lib/action';

class SignupPage extends Page {
    
	get inputFirstName() { return $("[name='firstname']"); }
	get errorEmptyName() { return $("//div[contains(text(),'your name?')]") }
	get errorEmptyNumberEmail() { return $("//div[contains(text(),'when you log in and if you ever need to reset your password.')]") }
	get errorEmptyNewPassword() { return $("//div[contains(text(),'Enter a combination of at least six numbers')]") }
	get errorEmptyDob() { return $("//div[contains(text(),'Please make sure that you use your real date of birth.')]") }
	get inputSurename() { return $("[name='lastname']");}
	get inputNumberEmail() { return $("[name='reg_email__']"); }
	get inputReEnterEmail() { return $("//input[contains(@name,'reg_email_confirmation')]"); }
	get inputPassword() { return $("[name='reg_passwd__']"); }
	get inputDate() { return $("[name='birthday_day']"); }
	get inputMonth() { return $("[name='birthday_month']"); }
	get inputYear() { return $("[name='birthday_year']"); }
	get inputGender() { return $$("[name='sex']"); }
	get linkLearnMore() { return $("#non-users-notice-link"); }
	get buttonSignUp() { return $("(//button)[1]"); }

	async open() {
		await super.open("");
	}

	async chooseDob(value: string) {
		const selector = $("//option[text()='"+value+"']");
		await action.scrollInto(selector);
		await action.expectToExist(selector);
		await action.clickOn(selector);
	};

}

export default new SignupPage();
