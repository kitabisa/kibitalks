/* eslint-disable @typescript-eslint/no-explicit-any */
import { expect, use } from "chai";
import chaiJsonSchema from "chai-json-schema";

use(chaiJsonSchema);

export function expectStatusCode(response: { status: any; }, expectedStatus: any) {
  expect(response.status).to.equal(expectedStatus);
}

export function expectSchema(response: { body: any; }, expectedSchema: any) {
  expect(response.body).to.be.jsonSchema(expectedSchema);
}

export function expectValue(response: any, expectedValue: any) {
  expect(response).to.equal(expectedValue);
}

export function expectDataPaymentMethod(response: any, expectedData: any) {
  expect(response.body.data[0].payment_method.name).to.equal(expectedData);
}
