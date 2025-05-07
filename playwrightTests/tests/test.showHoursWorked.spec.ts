import { test, expect } from "@playwright/test";
import { ApiMocker } from "../Helper/mockApi";
import { activityMocks } from "../Helper/Mocks/activity.mock";
import { projectMocks } from "../Helper/Mocks/project.mock";
import { UserMocks } from "../Helper/Mocks/user.Mock";

test.describe("showHoursWorked", () => {
  test.beforeEach(async ({ page }) => {
    const apiMocker = new ApiMocker(page);
    await apiMocker.addMocks([UserMocks.userMeSuccess]).apply();
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
      page.getByText("Vous avez travaillé 8.00 heures cette semaine.")
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
      "text=Vous avez travaillé 21.00 heures ce mois-ci.",
      { state: "visible" }
    );
    await expect(
      page.getByText("Vous avez travaillé 21.00 heures ce mois-ci.")
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
      "text=Vous avez travaillé 7.00 heures aujourd'hui.",
      { state: "visible" }
    );
    await expect(
      page.getByText("Vous avez travaillé 7.00 heures aujourd'hui.")
    ).toBeVisible();
  });

  test("hoursWorkedNoActivities", async ({ page }) => {
    const apiMocker = new ApiMocker(page);
    await apiMocker
      .addMocks([activityMocks.getAllActivitiesEmptyDefaultWeekSuccess])
      .apply();
    await page.goto("http://localhost:5002/calendar");
    await page.waitForSelector(".fc-event", { state: "visible" });
    await page.getByRole("button", { name: "Jour", exact: true }).click();
    await page.getByText("Bilan du 22 mars");
    await page.waitForTimeout(1000);
    await page.waitForSelector("text=Vous n'avez pas travaillé aujourd'hui.", {
      state: "visible",
    });
    expect(
      page.getByText("Vous n'avez pas travaillé aujourd'hui.")
    ).toBeVisible();
  });
});

