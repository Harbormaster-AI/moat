
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EmployeeService } from '../../../services/Employee.service';
import { Employee } from '../../../models/Employee';

@Component({
    selector: 'app-index-employee',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexEmployeeComponent implements OnInit {

    employees: Employee[] = [];

    constructor(
        private router: Router,
        private service: EmployeeService
) {}

    ngOnInit(): void {
        this.getEmployees();
}

    getEmployees(): void {
        this.service.getEmployees().subscribe((res) => {
        this.employees = res;
    });
}

    deleteEmployee(id: any): void {
        this.service.deleteEmployee(id)
            .subscribe(() => {
                this.getEmployees();
            });
    }
}