import { log } from "@suborbital/runnable";

export const run = (input) => {
  let message = input.split(";").join("-")

  log.info(`👋 log from the runnable: ${input}`);

  return message;
};
