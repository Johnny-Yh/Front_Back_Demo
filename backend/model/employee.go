package model

import (
	"back_go/config"
	"database/sql"
	"fmt"
)

type Employee struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

var DefaultEmployeeDb *EmployeeDb

func NewEmployeeDb() *EmployeeDb {
	return &EmployeeDb{TbName: "employee", Db: config.Db}
}

type EmployeeDb struct {
	TbName string
	Db     *sql.DB
}

func (s *EmployeeDb) Save(employee Employee) error {
	stmt, err := s.Db.Prepare(fmt.Sprintf("INSERT INTO %s(name,age) VALUES(?,?)", s.TbName))
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(employee.Name, employee.Age); err != nil {
		return err
	}
	return nil
}

func (s *EmployeeDb) GetAll() ([]Employee, error) {
	employees := make([]Employee, 0)
	rows, err := s.Db.Query(fmt.Sprintf("SELECT id,name,age from %s", s.TbName))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var employee Employee
		if err := rows.Scan(&employee.Id, &employee.Name, &employee.Age); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (s *EmployeeDb) GetOne(id int32) (Employee, error) {
	var employee Employee
	rows, err := s.Db.Query(fmt.Sprintf("SELECT id,name,age from %s WHERE id=%d", s.TbName, id))
	if err != nil {
		return employee, err
	}
	if rows.Next() {
		if err := rows.Scan(&employee.Id, &employee.Name, &employee.Age); err != nil {
			return employee, err
		}
	}
	return employee, nil
}

func (s *EmployeeDb) Remove(id int32) error {
	stmt, err := s.Db.Prepare(fmt.Sprintf("DELETE FROM %s WHERE id=%d", s.TbName, id))
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

func (s *EmployeeDb) Update(employee Employee) error {
	stmt, err := s.Db.Prepare(fmt.Sprintf("UPDATE %s SET name=?,age=? WHERE id=?", s.TbName))
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(employee.Name, employee.Age, employee.Id); err != nil {
		return err
	}
	return nil
}
