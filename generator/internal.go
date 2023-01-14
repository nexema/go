package main

/*func gen() {

	// prepare std
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	// read input
	buf, err := reader.ReadBytes('\n')
	if err != nil {
		writer.WriteString(fmt.Sprintf("failure: %s\n", err))
		writer.Flush()
		return
	}

	// unmarshal input
	var m map[string]interface{}
	err = json.Unmarshal(buf, &m)
	if err != nil {
		writer.WriteString(fmt.Sprintf("jfailure: %s\n", err))
		writer.Flush()
		return
	}

	// todo: generate
	writer.WriteString("ok\n")
	writer.Flush()
}
*/
