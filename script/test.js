import http from 'k6/http'

export let option = {
    vus: 5,
    duration: '5s'
}

export default function() {
    http.get("http://host.docker.internal:8000/hello")
}