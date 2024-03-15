{{define "response"}}
    <div
      class="border border-slate-400 shadow-2xl rounded-3xl bg-stone-100 container max-w-screen-md mx-auto sm p-8 pb-2"
    >
      <div class="">
        <dl class="flex items-center">
          <dt class="text-gray-500 dark:text-gray-400 text-sm font-normal me-1">
            Recommended moment to go out:
          </dt>
          <dd class="text-gray-900 text-sm dark:text-white font-semibold">
            Friday between 4:00 and 8:00
          </dd>
        </dl>
      </div>
      <div id="column-chart"></div>
    </div>
    <script>
      const actualPrecipitation = [
        1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
        21, 22, 23, 24,
      ];
      const actualTemperature = [
        10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 20, 19, 18, 17, 16, 15,
        14, 13, 12, 11, 10, 9,
      ];
      const actualCloudCover = [
        10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
        10, 10, 10, 10, 10, 10,
      ];
      const options = {
        series: [
          {
            name: "Temperature",
            data: [
              { x: "00:00", y: 10 },
              { x: "01:00", y: 1 },
              { x: "02:00", y: 10 },
              { x: "03:00", y: 10 },
              { x: "04:00", y: 20 },
              { x: "05:00", y: 10 },
              { x: "06:00", y: 10 },
              { x: "07:00", y: 10 },
              { x: "08:00", y: 1 },
              { x: "09:00", y: 10 },
              { x: "10:00", y: 10 },
              { x: "11:00", y: 10 },
              { x: "12:00", y: 10 },
              { x: "13:00", y: 1 },
              { x: "14:00", y: 10 },
              { x: "15:00", y: 10 },
              { x: "16:00", y: 10 },
              { x: "17:00", y: 10 },
              { x: "18:00", y: 10 },
              { x: "19:00", y: 10 },
              { x: "20:00", y: 1 },
              { x: "21:00", y: 10 },
              { x: "22:00", y: 1 },
              { x: "23:00", y: 10 },
            ],
          },
          {
            name: "Cloud cover",
            data: [
              { x: "00:00", y: 1 },
              { x: "01:00", y: 11 },
              { x: "02:00", y: 11 },
              { x: "03:00", y: 11 },
              { x: "04:00", y: 15 },
              { x: "05:00", y: 11 },
              { x: "06:00", y: 11 },
              { x: "07:00", y: 19 },
              { x: "08:00", y: 19 },
              { x: "09:00", y: 11 },
              { x: "10:00", y: 11 },
              { x: "11:00", y: 1 },
              { x: "12:00", y: 11 },
              { x: "13:00", y: 11 },
              { x: "14:00", y: 1 },
              { x: "15:00", y: 11 },
              { x: "16:00", y: 11 },
              { x: "17:00", y: 11 },
              { x: "18:00", y: 11 },
              { x: "19:00", y: 11 },
              { x: "20:00", y: 11 },
              { x: "21:00", y: 11 },
              { x: "22:00", y: 11 },
              { x: "23:00", y: 11 },
            ],
          },
          {
            name: "Precipitation",
            data: [
              { x: "00:00", y: 12 },
              { x: "01:00", y: 12 },
              { x: "02:00", y: 12 },
              { x: "03:00", y: 12 },
              { x: "04:00", y: 17 },
              { x: "05:00", y: 19 },
              { x: "06:00", y: 19 },
              { x: "07:00", y: 21 },
              { x: "08:00", y: 21 },
              { x: "09:00", y: 12 },
              { x: "10:00", y: 12 },
              { x: "11:00", y: 12 },
              { x: "12:00", y: 12 },
              { x: "13:00", y: 12 },
              { x: "14:00", y: 12 },
              { x: "15:00", y: 12 },
              { x: "16:00", y: 12 },
              { x: "17:00", y: 12 },
              { x: "18:00", y: 12 },
              { x: "19:00", y: 12 },
              { x: "20:00", y: 12 },
              { x: "21:00", y: 12 },
              { x: "22:00", y: 12 },
              { x: "23:00", y: 12 },
            ],
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
            if ([4, 5, 6, 7].includes(seriesIndex)) {
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
              `<p><b>${dataPointIndex}:00 - ${(dataPointIndex + 1) % 24}:00</p></b>` +
              `<p><b>Temperature (${actualTemperature[dataPointIndex]}C)</b>: ${series[0][dataPointIndex]} points</p>` +
              `<p><b>Precipitation (${actualPrecipitation[dataPointIndex]}mm)</b>: ${series[1][dataPointIndex]} points</p>` +
              `<p><b>Cloud Cover(${actualCloudCover[dataPointIndex]}%)</b>: ${series[2][dataPointIndex]} points</p>` +
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
        const chart = new ApexCharts(
          document.getElementById("column-chart"),
          options,
        );
        chart.render();
      }
    </script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/flowbite/2.3.0/flowbite.min.js"></script>
{{end}}
