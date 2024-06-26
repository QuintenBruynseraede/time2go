{{define "form"}}
    <div
      class="shadow-2xl border border-slate-400 rounded-3xl bg-stone-100 container max-w-screen-md mx-auto sm mt-4 pt-0 p-4"
    >
    <form hx-post="/form" hx-target="#response">
      <div class="text-2xl text-center p-4">
        <span>I want to do a</span>
        <span class="inline-block"
          ><input
            type="number"
            min="1"
            max="24"
            value="1"
            id="duration"
            name="duration"
            class="w-16 bg-gray-50 invalid:border-pink-500 border border-gray-300 text-gray-900 m-3 text-sm rounded-lg focus:ring-slate-200 focus:border-slate-200 block p-2.5"
            required
        /></span>
        <span>hour activity around </span>
        <span class="inline-block">
          <!-- Search input -->
          <div class="relative">
            <div
              class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none"
            >
              <svg
                class="w-4 h-4 text-gray-500 dark:text-gray-400"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 20 20"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"
                />
              </svg>
            </div>
            <input
              type="search"
              id="default-search"
              name="location"
              class="block w-full p-2 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="Search"
              required
              hx-get="/search"
              hx-trigger="keyup changed delay:500ms"
              hx-target="#search_results"
            />
          </div>
          <!-- Dropdown menu -->
          <div id="search_results">
         </div>
        </span>
      </div>
      <div class="grid grid-flow-col place-content-center p-2" role="group">
        <div class="text-base bg-stone-200 has-[:checked]:bg-slate-300 rounded-lg py-3 me-2 mb-2 px-4">
            <input class="hidden peer" type="radio" id="today" name="daterange" value="today" checked>
            <label class="p-4 peer-[:checked]:underline" for="today">Today</label>
        </div>
        <div class="text-base bg-stone-200 has-[:checked]:bg-slate-300 rounded-lg py-3 me-2 mb-2 px-4">
          <input class="hidden peer" type="radio" id="tomorrow" name="daterange" value="tomorrow">
          <label class="p-4 peer-[:checked]:underline" for="tomorrow">Tomorrow</label>
        </div>
        <div class="text-base bg-stone-200 has-[:checked]:bg-slate-300 rounded-lg py-3 me-2 mb-2 px-4">
          <input class="hidden peer" type="radio" id="this_week" name="daterange" value="this_week">
          <label class="p-4 peer-[:checked]:underline" for="this_week">This week</label>
        </div>
      </div>
      <div class="grid p-2 place-content-center">
        <input type="submit"
          class="font-bold bg-stone-300 hover:bg-slate-300 rounded-lg py-3 px-8"
        value="Find the best moment ">
    </input>
      </div>
      </form>
    </div>
{{end}}
