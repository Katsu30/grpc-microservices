import { render } from "@testing-library/react";
import { expect, test, describe } from "vitest";
import { BoardTop } from "@/features/board/components/BoardTop";

describe("BoardTop", () => {
  test("renders board top", () => {
    const wrapper = render(<BoardTop />);
    expect(wrapper.findByText("board top")).toBeTruthy();
  });
});
