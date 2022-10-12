package main

import "fmt"

type User struct {
	Id   int
	Name string
}

type ProjectMember struct {
	ProjectId int
	User      User
}

func f01() {
	var u1 User = User{Id: 1, Name: "AAA"}
	fmt.Printf("u1 %v (%p)\n", u1, &u1)

	var u2 User = u1
	fmt.Printf("u2 %v (%p)\n", u2, &u2)
}

func f02_bk() {
	var x = User{Id: 1, Name: "AAA"}

	var u1 *User = &x
	fmt.Printf("u1 %v (%p)\n", u1, &u1)

	var u2 *User = u1
	fmt.Printf("u2 %v (%p)\n", u2, &u2)
}

func f02() {
	var u1 *User = &User{Id: 1, Name: "AAA"}
	fmt.Printf("u1 %v (%p)\n", *u1, u1)

	var u2 *User = u1
	fmt.Printf("u2 %v (%p)\n", *u2, u2)
}

func f03() {
	var u1 *User = &User{Id: 1, Name: "AAA"}
	fmt.Printf("u1 %v (%p)\n", *u1, u1)

	var u2 User = *u1

	fmt.Printf("u2 %v (%p)\n", u2, &u2)
}

func f04() {
	var users1 []User = []User{{Id: 1, Name: "AAA"}, {Id: 2, Name: "BBB"}}

	fmt.Printf("users1 %v (%p)\n", users1, &users1)

	// for i, user := range users1 {
	// 	fmt.Printf("users1[%d] %v (%p)\n", i, user, &user)
	// }

	for i := 0; i < len(users1); i++ {
		fmt.Printf("users1[%d] %v (%p)\n", i, users1[i], &users1[i])
	}

	var users2 []User = users1

	fmt.Printf("users2 %v (%p)\n", users2, &users2)

	// for i, user := range users2 {
	// 	fmt.Printf("users2[%d] %v (%p)\n", i, user, &user)
	// }

	for i := 0; i < len(users2); i++ {
		fmt.Printf("users2[%d] %v (%p)\n", i, users2[i], &users2[i])
	}
}

func f05() {
	var users1 []*User = []*User{{Id: 1, Name: "AAA"}, {Id: 2, Name: "BBB"}}

	fmt.Printf("users1 %v (%p)\n", users1, &users1)

	for i, user := range users1 {
		fmt.Printf("users1[%d] %v (%p)\n", i, user, user)
	}

	var users2 []*User = users1

	fmt.Printf("users2 %v (%p)\n", users2, &users2)

	for i, user := range users2 {
		fmt.Printf("users2[%d] %v (%p)\n", i, user, user)
	}
}

func f06() {

	// var m ProjectMember = ProjectMember{ProjectId: 1, User: User{Id: 1, Name: "AAA"}}
	var projectMembers []ProjectMember = []ProjectMember{{ProjectId: 1, User: User{Id: 1, Name: "AAA"}}, {ProjectId: 1, User: User{Id: 2, Name: "BBB"}}}
	fmt.Printf("projectMembers %v (%p)\n", projectMembers, &projectMembers)

	for i := 0; i < len(projectMembers); i++ {
		fmt.Printf("projectMembers[%d] %v (%p)\n", i, projectMembers[i].User, &projectMembers[i].User)
	}

	users := make([]User, len(projectMembers))
	for i := 0; i < len(users); i++ {
		users[i] = projectMembers[i].User
	}

	for i := 0; i < len(users); i++ {
		fmt.Printf("users[%d] %v (%p)\n", i, users[i], &users[i])
	}
}

func f07() {

	// var m ProjectMember = ProjectMember{ProjectId: 1, User: User{Id: 1, Name: "AAA"}}
	var projectMembers []ProjectMember = []ProjectMember{{ProjectId: 1, User: User{Id: 1, Name: "AAA"}}, {ProjectId: 1, User: User{Id: 2, Name: "BBB"}}}
	fmt.Printf("projectMembers %v (%p)\n", projectMembers, &projectMembers)

	for i := 0; i < len(projectMembers); i++ {
		fmt.Printf("projectMembers[%d] %v (%p)\n", i, projectMembers[i].User, &projectMembers[i].User)
	}

	users := make([]*User, len(projectMembers))
	for i := 0; i < len(users); i++ {
		users[i] = &projectMembers[i].User
	}

	for i := 0; i < len(users); i++ {
		fmt.Printf("users[%d] %v (%p)\n", i, *users[i], users[i])
	}
}

func main() {
	f07()
}
