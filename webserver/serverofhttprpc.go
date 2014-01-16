package main

import (
	"fmt"
	"net/rpc"
	//"errors"
	"bufio"
	"net"
	"os"
	"strings"
)
//Structure for storing key value pairs
type Args struct {
	K, V *string
}

//Type of string on which web services are defined
type BaseOperations string


//Insert procedure for inserting key value pairs
func (t *BaseOperations) Insert(args *Args, reply *string) error {

	f, err := os.OpenFile("kv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		fmt.Printf("File error")
	}

	fmt.Fprintf(f, "<%v,%v>\n", *args.K, *args.V)

	f.Close()

	tmp := "Insertion successful"
	*reply = tmp
	return nil
}

//Remove the key,value pair by using key
func (t *BaseOperations) Removebykey(k *string, reply *string) error {
	file, err := os.Open("kv")
	
	*reply=""
	if err != nil {
		return err
	}
	defer file.Close()
	
	

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

		

	f2, _ := os.OpenFile("tmpp", os.O_WRONLY|os.O_CREATE, 0777)

	str1 := ""
	str2 := ""

	for _, line := range lines {
		str1 = strings.Split(line, ",")[0]
		str2 = strings.Split(line, ",")[1]

		str1 = strings.Split(str1, "<")[1]
		str2 = strings.Split(str2, ">")[0]
		fmt.Println(str1, str2, *k)

		if *k == str1 {
			//fmt.Println("Compared successfully")
			//*reply = str2
			tmp := "Remove by key is successful"
			*reply = tmp

			

		}
		if *k != str1 {
			fmt.Fprintf(f2, "<%v,%v>\n", str1, str2)
		}

	}
	f2.Close()
	err = os.Remove("kv")
	err = os.Rename("tmpp", "kv")
	err = os.Remove("tmpp")

	//tmp := "Remove by key is not successful"
	//*reply = tmp

	return nil

}

//Remove the key,value pair by using value
func (t *BaseOperations) Removebyvalue(v *string, reply *string) error {

	file, err := os.Open("kv")
	*reply=""
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	f2, _ := os.OpenFile("tmpp", os.O_WRONLY|os.O_CREATE, 0777)

	str1 := ""
	str2 := ""

	for _, line := range lines {
		str1 = strings.Split(line, ",")[0]
		str2 = strings.Split(line, ",")[1]

		str1 = strings.Split(str1, "<")[1]
		str2 = strings.Split(str2, ">")[0]
		fmt.Println(str1, str2, *v)

		if *v == str2 {
			//fmt.Println("Compared successfully")
			//*reply = str2
			
			tmp := "Remove by value successful"
			*reply = tmp

			

		}
		if *v != str2 {
			fmt.Fprintf(f2, "<%v,%v>\n", str1, str2)
		}

	}
	f2.Close()
	err = os.Remove("kv")
	err = os.Rename("tmpp", "kv")
	err = os.Remove("tmpp")

	//tmp := "Remove by value not successful"
	//*reply = tmp

	return nil
}

//Search the key,value pair by using key
func (t *BaseOperations) Searchbykey(k *string, reply *string) error {

	file, err := os.Open("kv")
	
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	str1 := ""
	str2 := ""

	for _, line := range lines {
		str1 = strings.Split(line, ",")[0]
		str2 = strings.Split(line, ",")[1]

		str1 = strings.Split(str1, "<")[1]
		str2 = strings.Split(str2, ">")[0]
		fmt.Println(str1, str2, *k)

		if *k == str1 {
			fmt.Println("Compared successfully")
			*reply = str2
			return nil
		}

	}

	return nil

}

//Remove the key,value pair by using value
func (t *BaseOperations) Searchbyvalue(v *string, reply *string) error {

	file, err := os.Open("kv")
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	str1 := ""
	str2 := ""

	for _, line := range lines {
		str1 = strings.Split(line, ",")[0]
		str2 = strings.Split(line, ",")[1]

		str1 = strings.Split(str1, "<")[1]
		str2 = strings.Split(str2, ">")[0]
		fmt.Println(str1, str2, *v)

		if *v == str2 {
			fmt.Println("Compared successfully")
			*reply = str1
			return nil
		}

	}

	return nil
}

//For reading the EXIT to exit from program
func isexit() {
	rd := bufio.NewReader(os.Stdin)
	ip, _ := rd.ReadString('\n')
	ipstr := string([]byte(ip)[0:])
	ipstr = ipstr[:len(ipstr)-1]
	if ipstr == "EXIT" {
		os.Exit(1)
	}
}

func main() {
	baseoperations := new(BaseOperations)
	rpc.Register(baseoperations)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	go isexit()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}

//Checking the error 
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
