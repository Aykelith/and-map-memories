export function createStore() {
    window.and_store = {};
}

export function setStoreItem<T = any>(name: string, data: T) {
    window.and_store[name] = data;
}

export function getStoreItem<T = any>(name: string): T {
    return window.and_store[name];
}