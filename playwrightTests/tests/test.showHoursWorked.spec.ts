import { test, expect } from "@playwright/test";
import { ApiMocker } from "../Helper/mockApi";
import { activityMocks } from "../Helper/Mocks/activity.mock";
import { projectMocks } from "../Helper/Mocks/project.mock";
import { userMocks } from "../Helper/Mocks/user.Mock";

test.describe("showHoursWorked", () => {
  test.beforeEach(async ({ page }) => {
    const apiMocker = new ApiMocker(page);
    await apiMocker.addMocks([userMocks.userMeSuccess]).apply();
    await page.clock.install({ time: new Date("2025-03-22T08:00:00-04:00") });
  });

  test("showHoursWorkedDefault", async ({ page }) => {
    const apiMocker = new ApiMocker(page);
    await apiMocker
      .addMocks([activityMocks.getAllActivitiesDefaultWeekSuccess])
      .apply();
    await page.goto("http://localhost:5002/calendar");
    await page.waitForSelector(".fc-event", { state: "visible" });
    await expect(
      page.getByText("Vous avez travaillé 8 heures cette semaine.")
    ).toBeVisible();
  });

  test("showHoursWorkedMonth", async ({ page }) => {
    const apiMocker = new ApiMocker(page);
    await apiMocker
      .addMocks([
        activityMocks.getAllActivitiesMonthSuccess,
        activityMocks.getAllActivitiesDefaultWeekSuccess,
      ])
      .apply();
    await page.goto("http://localhost:5002/calendar");
    await page.waitForSelector(".fc-event", { state: "visible" });
    await page.getByRole("button", { name: "Mois", exact: true }).click();
    await page.getByText("Bilan du 1 mars au 31 mars");
    await page.waitForTimeout(1000);
    await page.waitForSelector(
      "text=Vous avez travaillé 21 heures ce mois-ci.",
      { state: "visible" }
    );
    await expect(
      page.getByText("Vous avez travaillé 21 heures ce mois-ci.")
    ).toBeVisible();
  });

  test("showHoursWorkedDay", async ({ page }) => {
    const apiMocker = new ApiMocker(page);
    await apiMocker
      .addMocks([
        activityMocks.getAllActivitiesDaySuccess,
        activityMocks.getAllActivitiesDefaultWeekSuccess,
      ])
      .apply();
    await page.goto("http://localhost:5002/calendar");
    await page.waitForSelector(".fc-event", { state: "visible" });
    await page.getByRole("button", { name: "Jour", exact: true }).click();
    await page.getByText("Bilan du 22 mars");
    await page.waitForTimeout(1000);
    await page.waitForSelector(
      "text=Vous avez travaillé 7 heures aujourd'hui.",
      { state: "visible" }
    );
    await expect(
      page.getByText("Vous avez travaillé 7 heures aujourd'hui.")
    ).toBeVisible();
  });

  test("hoursWorkedNoActivities", async ({ page }) => {
    const apiMocker = new ApiMocker(page);
    await apiMocker
      .addMocks([activityMocks.getAllActivitiesEmptyDefaultWeekSuccess])
      .apply();
    await page.goto("http://localhost:5002/calendar");
    await page.getByRole("button", { name: "Jour", exact: true }).click();
    await page.getByText("Bilan du 22 mars");
    await page.waitForTimeout(1000);
    await page.waitForSelector(
      "text=Vous avez travaillé 0 heures cette semaine.",
      {
        state: "visible",
      }
    );
    await expect(
      page.getByText("Vous avez travaillé 0 heures cette semaine.")
    ).toBeVisible();
  });
});
