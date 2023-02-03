export const schCreateUser = {
  type: "object",
  required: ["name", "job", "id", "createdAt"],
  additionalProperties: true,
  properties: {
    name: {
      type: "string",
    },
    job: {
      type: "string",
    },
    id: {
      type: "string",
    },
    createdAt: {
      type: "string",
    },
  },
};

export const schGetAllUsers = {
  type: "object",
  required: ["page", "per_page", "total", "total_pages", "data", "support"],
  additionalProperties: false,
  properties: {
    page: {
      type: "integer",
    },
    per_page: {
      type: "integer",
    },
    total: {
      type: "integer",
    },
    total_pages: {
      type: "integer",
    },
    data: {
      type: "array",
      additionalItems: false,
      items: {
        type: "object",
        required: ["id", "email", "first_name", "last_name", "avatar"],
        additionalProperties: false,
        properties: {
          id: {
            type: "integer",
          },
          email: {
            type: "string",
          },
          first_name: {
            type: "string",
          },
          last_name: {
            type: "string",
          },
          avatar: {
            type: "string",
          },
        },
      },
    },
    support: {
      type: "object",
      required: ["url", "text"],
      additionalProperties: false,
      properties: {
        url: {
          type: "string",
        },
        text: {
          type: "string",
        },
      },
    },
  },
};

export const schEditUser = {
  type: "object",
  required: ["name", "job", "updatedAt"],
  additionalProperties: false,
  properties: {
    name: {
      type: "string",
    },
    job: {
      type: "string",
    },
    updatedAt: {
      type: "string",
    },
  },
};
