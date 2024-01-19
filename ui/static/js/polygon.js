
var latlon=[];


fetch('static/json/weather.json')
    .then(response => response.json())
    .then(data => {
    data.forEach(item => {
        latlon.push([item.location.lat, item.location.lon]);
    });
    console.log(latlon);
    latlon.sort()
    const polygon = L.polygon(latlon,{color:'red'}).addTo(map);
})
