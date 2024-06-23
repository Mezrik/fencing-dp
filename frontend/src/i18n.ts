import { i18n } from "@lingui/core";

export const locales = {
  en: "English",
  cs: "Česky",
} as const;

export const defaultLocale = "en";

export async function dynamicActivate(locale: keyof typeof locales) {
  // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
  const { messages } = await import(`./locales/${locale}.po`);

  // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
  i18n.load(locale, messages);
  i18n.activate(locale);
}
