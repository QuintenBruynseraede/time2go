{{define "response"}}
    <div
      class="border border-slate-400 shadow-2xl rounded-3xl bg-stone-100 container max-w-screen-md mx-auto sm p-8 pb-2"
    >
      <div class="">
        <dl class="flex items-center">
          <dt class="text-gray-500 dark:text-gray-400 text-sm font-normal me-1">
            Recommended moment to go out near <b>{{.Location}}:</b>
          </dt>
          <dd class="text-gray-900 text-sm dark:text-white font-semibold">
           {{.RecommendedMoment}}
          </dd>
        </dl>
      </div>
      <div id="column-chart"></div>
    </div>
    <script>
      var actualPrecipitation = {{.PrecipitationList}}
      var actualTemperature = {{.TemperatureList}}
      var actualCloudCover = {{.CloudCoverList}}
      var PrecipitationScore = {{.PrecipitationScoreList}}
      var TemperatureScore = {{.TemperatureScoreList}}
      var CloudCoverScore = {{.CloudCoverScoreList}}
      var TimeRanges = {{.TimeRangeList}}
      var options = {
        series: [
          {
            name: "Temperature",
            data: [...Array(TimeRanges.length).keys()].map(i => {return {"x": TimeRanges[i], "y": TemperatureScore[i]}})
          },
          {
            name: "Cloud cover",
            data: [...Array(TimeRanges.length).keys()].map(i => {return {"x": TimeRanges[i], "y": CloudCoverScore[i]}})
          },
          {
            name: "Precipitation",
            data: [...Array(TimeRanges.length).keys()].map(i => {return {"x": TimeRanges[i], "y": PrecipitationScore[i]}})
          },
        ],
        chart: {
          type: "bar",
          height: "320px",
          stacked: true,
          fontFamily: "Inter, sans-serif",
          toolbar: {
            show: false,
          },
        },
        colors: [
          function ({ value, seriesIndex, dataPointIndex, w }) {
            if ({{.RecommendedTimeRangeIndices}}.includes(seriesIndex)) {
              return "#65a30d";
            } else {
              return "#a1c8ed";
            }
          },
        ],
        plotOptions: {
          bar: {
            horizontal: false,
            columnWidth: "70%",
            borderRadiusApplication: "end",
            borderRadius: 8,
            distributed: true,
          },
        },
        tooltip: {
          shared: true,
          intersect: false,
          style: {
            fontFamily: "Inter, sans-serif",
          },
          custom: function ({ series, seriesIndex, dataPointIndex, w }) {
            let currentTotal = 0;
            series.forEach((s) => {
              currentTotal += s[dataPointIndex];
            });
            return (
              '<div class="rounded-lg p-2 shadow border custom-tooltip"><span>' +
              `<p><b>${TimeRanges[dataPointIndex]}</p></b>` +
              `<p><b>Temperature (${actualTemperature[dataPointIndex]}C)</b>: ${series[0][dataPointIndex]} points</p>` +
              `<p><b>Cloud Cover(${actualCloudCover[dataPointIndex]}%)</b>: ${series[1][dataPointIndex]} points</p>` +
              `<p><b>Precipitation (${actualPrecipitation[dataPointIndex]}mm)</b>: ${series[2][dataPointIndex]} points</p>` +
              "</span></div>"
            );
          },
        },
        states: {
          hover: {
            filter: {
              type: "darken",
              value: 1,
            },
          },
        },
        stroke: {
          show: true,
          width: 0,
          colors: ["transparent"],
        },
        grid: {
          show: false,
          strokeDashArray: 4,
          padding: {
            left: 2,
            right: 2,
            top: -14,
          },
        },
        dataLabels: {
          enabled: false,
        },
        legend: {
          show: false,
        },
        xaxis: {
          tickAmount: 6,
          floating: false,
          labels: {
            show: true,
            style: {
              fontFamily: "Inter, sans-serif",
              cssClass: "text-xs font-normal fill-gray-500 dark:fill-gray-400",
            },
          },
          axisBorder: {
            show: false,
          },
          axisTicks: {
            show: false,
          },
        },
        yaxis: {
          show: true,
        },
        fill: {
          opacity: 1,
        },
      };

      if (
        document.getElementById("column-chart") &&
        typeof ApexCharts !== "undefined"
      ) {
        var chart = new ApexCharts(
          document.getElementById("column-chart"),
          options,
        );
        chart.render();
      }
    </script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/flowbite/2.3.0/flowbite.min.js"></script>
{{end}}
