package View

type Engine interface {
	/**
	 * Get the evaluated contents of the view.
	 *
	 * @param  string  path
	 * @param  array  data
	 * @return string
	 */
	get(path string, array map[string]interface{}) string
}
