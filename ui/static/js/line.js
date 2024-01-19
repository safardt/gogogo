var latlon = [];

fetch('static/json/weather.json')
    .then(response=>response.json())
    .then(data=>{
        data.forEach(item => {
            var marker=L.marker([item['location']['lat'], item['location']['lon']]).addTo(map);
            marker.bindPopup('<b>'+item['current']['temp_c']+'</b><br>');
            latlon.push([item['location']['lat'], item['location']['lon']]);
        });
        var polyline = L.polyline(latlon,{color: 'red'}).addTo(map);
        console.log(data);
    })
    .catch(error=>{
        console.log('Error:',error);
    });