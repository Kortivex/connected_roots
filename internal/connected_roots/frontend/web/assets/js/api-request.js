document.addEventListener("DOMContentLoaded", () => {
    let sensorID = document.getElementById('sensor-id')
    if (sensorID) {
        sensorID = sensorID.value;
        console.info('Starting to fetch sensor (', sensorID, ') data periodically...');
        startFetchingData(sensorID, 6000);
    }
});

const apiHost = 'http://localhost:47400';
const httpClient = axios.create({
    baseURL: apiHost,
    headers: {
        'Authorization': 'Bearer 4cae8c84-dd29-42f3-8d58-ed371f1bc8ef',
    },
    withCredentials: true
})

function startFetchingData(sensorID, interval) {
    fetchSensorData(sensorID);
    setInterval(() => fetchSensorData(sensorID), interval);
}

async function fetchSensorData(sensorID) {
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
