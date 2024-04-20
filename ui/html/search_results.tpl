{{define "search_results"}}
<div
  id="dropdown"
  class="absolute z-10 bg-white divide-y divide-gray-100 rounded-lg shadow w-44"
>
  <ul
    class="flex-1 py-2 text-sm text-gray-700 dark:text-gray-200"
    aria-labelledby="dropdownDefaultButton"
  >
{{if not .}}
    <li>
      <div class="block px-4 py-2" ><i>No results</i></a>
    </li>
{{end}}
{{range .}}
    <li>
      <a href="" class="block px-4 py-2 hover:bg-gray-100">{{title .}}</a>
    </li>
{{end}}
  </ul>
</div>
{{end}}
