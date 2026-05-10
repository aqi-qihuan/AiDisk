/**
 * Type declarations for spark-md5
 */
declare module 'spark-md5' {
  export default class SparkMD5 {
    /**
     * Create a SparkMD5 instance
     */
    constructor();

    /**
     * Append a string to the hash
     * @param str The string to append
     */
    append(str: string): SparkMD5;

    /**
     * Finish the hash calculation and return the result
     * @param raw Whether to return the raw binary format
     * @returns The hash result
     */
    end(raw?: boolean): string;

    /**
     * Reset the hash state
     */
    reset(): SparkMD5;

    /**
     * Get the current state of the hash
     * @returns The current state
     */
    getState(): SparkMD5.State;

    /**
     * Set the state of the hash
     * @param state The state to set
     */
    setState(state: SparkMD5.State): SparkMD5;

    /**
     * Destroy the SparkMD5 instance
     */
    destroy(): void;

    /**
     * Calculate the hash of a string
     * @param str The string to hash
     * @param raw Whether to return the raw binary format
     * @returns The hash result
     */
    static hash(str: string, raw?: boolean): string;

    /**
     * ArrayBuffer hash calculator class
     */
    static ArrayBuffer: typeof SparkMD5ArrayBuffer;
  }

  namespace SparkMD5 {
    interface State {
      buff: Uint8Array;
      length: number;
      hash: number[];
    }
  }

  /**
   * ArrayBuffer hash calculator
   */
  class SparkMD5ArrayBuffer {
    /**
     * Create a SparkMD5.ArrayBuffer instance
     */
    constructor();

    /**
     * Append an ArrayBuffer to the hash
     * @param arr The ArrayBuffer to append
     */
    append(arr: globalThis.ArrayBuffer): SparkMD5ArrayBuffer;

    /**
     * Finish the hash calculation and return the result
     * @param raw Whether to return the raw binary format
     * @returns The hash result
     */
    end(raw?: boolean): string;

    /**
     * Reset the hash state
     */
    reset(): SparkMD5ArrayBuffer;

    /**
     * Get the current state of the hash
     * @returns The current state
     */
    getState(): SparkMD5.State;

    /**
     * Set the state of the hash
     * @param state The state to set
     */
    setState(state: SparkMD5.State): SparkMD5ArrayBuffer;

    /**
     * Destroy the SparkMD5.ArrayBuffer instance
     */
    destroy(): void;

    /**
     * Calculate the hash of an ArrayBuffer
     * @param arr The ArrayBuffer to hash
     * @param raw Whether to return the raw binary format
     * @returns The hash result
     */
    static hash(arr: globalThis.ArrayBuffer, raw?: boolean): string;
  }
}
