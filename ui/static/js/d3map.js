// Создание контейнера карты
const mapContainer = d3.select("#map");
const width = 960;
const height = 500;


// Создание контейнера для отрисовки карты
const svg = d3.select("#map")
    .append("svg")
    .attr("width", width)
    .attr("height", height);

var projection;
   
// Загрузка географических данных
d3.json("/static/json/custom.geo.json").then(function (data) {
    projection = d3.geoMercator().fitSize([width, height], data);
    const path = d3.geoPath().projection(projection);
    
    // Отрисовка границ стран
      
    svg.selectAll("path")
      .data(data.features)
      .enter()
      .append("path")
      .attr("class", "country")
      .attr("d", path)
      .attr("fill","steelblue");


       
 
}).catch(function (error) {
    console.log(error);
});
