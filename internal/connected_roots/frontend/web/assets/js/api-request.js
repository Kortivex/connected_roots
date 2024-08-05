function generateHttpClient(apiHost, apiKey) {
    return axios.create({
        baseURL: apiHost,
        headers: {
            'Authorization': 'Bearer ' + apiKey,
        },
        withCredentials: true
    })
}

function startFetchingData(httpClient, sensorID, interval) {
    fetchSensorData(httpClient, sensorID);
    setInterval(() => fetchSensorData(httpClient, sensorID), interval);
}

async function fetchSensorData(httpClient, sensorID) {
    try {
        let response = await httpClient.get('/sensors/' + sensorID + '/last-data');
        let data = response.data;
        document.getElementById('temperature-in').textContent = `${data.temperature_in.toFixed(2)} °C`;
        document.getElementById('humidity-in').textContent = `${data.humidity_in.toFixed(2)} %`;
        document.getElementById('salt').textContent = `${data.salt} %`;
        document.getElementById('soil').textContent = `${data.soil} %`;
        document.getElementById('light').textContent = `${data.light.toFixed(2)} lx`;
        document.getElementById('altitude').textContent = `${data.altitude.toFixed(2)} m`;
        document.getElementById('pressure').textContent = `${data.pressure.toFixed(2)} hPa`;
        document.getElementById('temperature-out').textContent = `${data.temperature_out.toFixed(2)} °C`;
        document.getElementById('humidity-out').textContent = `${data.humidity_out.toFixed(2)} %`;
        document.getElementById('battery').textContent = `${data.battery.toFixed(2)} %`;
        document.getElementById('battery-level').textContent = `${data.voltage.toFixed(2)} mV`;
    } catch (error) {
        console.error('Error fetching data:', error);
    }
}
