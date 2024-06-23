import { i18n } from "@lingui/core";

export const locales = {
  en: "English",
  cs: "Česky",
} as const;

export const defaultLocale = "en";

export async function dynamicActivate(locale: keyof typeof locales) {
  const { messages } = await import(`./locales/${locale}.po`);

  i18n.load(locale, messages);
  i18n.activate(locale);
}
