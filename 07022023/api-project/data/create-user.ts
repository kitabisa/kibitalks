import { faker } from '@faker-js/faker';

export const createUser = {
  name: faker.name.fullName(),
  job: faker.company.companySuffix(),
};
