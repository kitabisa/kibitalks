import { faker } from '@faker-js/faker';

export const editUser = {
  name: faker.name.fullName(),
  job: faker.company.companySuffix(),
};
