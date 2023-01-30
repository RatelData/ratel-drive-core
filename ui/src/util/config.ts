import config from '../config/app.json'

export function loginUrl(): string {
    const env = environment();
    const server = env.local_server
    return `${server.protocol}://${server.host}:${server.port}${config.endpoints.login}`
}

function environment() {
    const cur_env = config.current_environment

    if (cur_env === 'development')
        return config.environments.development
    if (cur_env === 'production')
        return config.environments.production
    
    return config.environments.development
}