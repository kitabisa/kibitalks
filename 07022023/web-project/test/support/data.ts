import { faker } from '@faker-js/faker';

const data = {
  validData: {
    firstName: faker.name.firstName('female'),
    surename: faker.name.lastName(),
    phoneNumber: faker.phone.number('08##########'),
    email: faker.internet.email(),
    newPass: "tryT0Gu3st!",
    dobDate: "1",
    dobMonth: "Jan",
    dobYear: "1995",
    registerUrlSuccess: "checkpoint"
  }
};

export default data;