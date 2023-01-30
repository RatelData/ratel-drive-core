export function saveUserCredentials(uid: string, token: string, is_storage_device: boolean) {
    const local_storage = window.localStorage;
    local_storage.setItem('user_id', uid)
    local_storage.setItem('token', token)
    local_storage.setItem('is_storage_device', is_storage_device.toString())
}

export function isStorageDevice(): boolean {
    const local_storage = window.localStorage;
    return Boolean(local_storage.getItem('is_storage_device'))
}