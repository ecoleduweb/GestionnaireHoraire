import { MockConfig } from "../types";

export const categoryMocks = {
  getCategoriesByProjectSuccess: {
    url: "/project/*/categories",
    method: "GET",
    response: {
      status: 200,
      json: {
        categories: [
          {
            id: 1,
            name: "Test",
            description: "Nouveau format",
            billable: false,
            activities: [],
            createdAt: "2025-04-02T10:29:54-04:00",
            updatedAt: "2025-04-02T10:29:54-04:00",
            userId: 1,
            projectId: 1,
          },
        ],
      },
    },
  },
} satisfies Record<string, MockConfig>;
