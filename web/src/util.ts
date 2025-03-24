export type To<T, E = Error> = [T, null] | [null, E];

/**
 * An oath will never throw an error, it's just a convention.
 */
export type Oath<T> = Promise<T>;

export async function to<T, E = Error>(promise: Promise<T>): Promise<To<T, E>> {
  try {
    const result = await promise;
    return [result, null];
  } catch (error) {
    return [null, error as E];
  }
}
