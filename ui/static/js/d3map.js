// Создание контейнера карты
const mapContainer = d3.select("#map");
const width = 1920;
const height = 1080;


// Создание контейнера для отрисовки карты
const svg = d3.select("#map")
    .append("svg")
    .attr("width", width)
    .attr("height", height)
    .call(d3.drag().on("drag",handleDrag));

const sphere = {type: "Sphere"};
var projection;
   
// Загрузка географических данных
d3.json("/static/json/custom.geo.json").then(function (data) {
    projection = d3.geoOrthographic().fitSize([width, height], data);
    const path = d3.geoPath().projection(projection);
    
    // Отрисовка границ стран
      
    svg.selectAll("path")
      .data(data.features)
      .enter()
      .append("path")
      .attr("class", "country")
      .attr("d", path);


       
 
}).catch(function (error) {
    console.log(error);
});
function handleDrag(event) {
  const rotate = projection.rotate();
  const k = sensitivity();
  projection.rotate([
    rotate[0] + event.dx * k,
    rotate[1] - event.dy * k,
  ]);
  svg.selectAll("path").attr("d", path);
}

function sensitivity() {
  return 0.25;
}