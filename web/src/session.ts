export type SessionKey = "authorization-header";

export function sessionGet(key: SessionKey) {
  return sessionStorage.getItem(key);
}

export function sessionSet(key: SessionKey, value: string) {
  sessionStorage.setItem(key, value);
}
