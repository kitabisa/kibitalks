import { userId } from "../data/user-id-lib";
import { httpStatusCode } from "../data/http-code-lib";
import { createUser } from "../data/create-user";
import { getUser } from "../data/get-user";
import { editUser } from "../data/edit-user";
import * as reqresEndPo from "../endpoint/reqres-endpoint";
import * as expectCondition from "../expect/common-expect";
import * as reqresSchema from "../schema/reqres-sch";

describe("Kibitalk Test", () => {
  it("POST - Create user", async () => {
    const body = createUser;
    const response = await reqresEndPo.createUser(body);
    expectCondition.expectStatusCode(response, httpStatusCode.created);
    expectCondition.expectValue(response.body.name, body.name);
    expectCondition.expectValue(response.body.job, body.job);
    expectCondition.expectSchema(response, reqresSchema.schCreateUser);
  });

  it("GET - Get all user", async () => {
    const param = getUser;
    const response = await reqresEndPo.getAllUsers(param);
    expectCondition.expectStatusCode(response, httpStatusCode.ok);
    expectCondition.expectValue(response.body.page, param.page);
    expectCondition.expectSchema(response, reqresSchema.schGetAllUsers);
  });

  it("PUT - Edit user", async () => {
    const path = userId.userId;
    const body = editUser;
    const response = await reqresEndPo.editUser(path, body);
    expectCondition.expectStatusCode(response, httpStatusCode.ok);
    expectCondition.expectValue(response.body.name, body.name);
    expectCondition.expectValue(response.body.job, body.job);
    expectCondition.expectSchema(response, reqresSchema.schEditUser);
  });

  it("DEL - Delete user", async () => {
    const path = userId.userId;
    const response = await reqresEndPo.deleteUser(path);
    expectCondition.expectStatusCode(response, httpStatusCode.noContent);
  });
});
