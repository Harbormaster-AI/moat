import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { EmployeeService } from '../../../services/Employee.service';
import { Employee } from '../../../models/Employee';
import { SubBaseComponent } from '../../Employee/sub.base.component';

@Component({
    selector: 'app-create-employee',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateEmployeeComponent extends SubBaseComponent implements OnInit {

    title = 'Add Employee';

    employeeForm: FormGroup;
    employee: Employee;

    constructor( http: HttpClient,
        private employeeService: EmployeeService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.employeeForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      WorkCenter: ['', ],
      ShiftAssignments: ['', ],
      CorrectiveActions: ['', ],
      Role: ['', ],
      SkillLevel: ['', ]
        });
    }

    
    addEmployee(firstName, lastName, WorkCenter, ShiftAssignments, CorrectiveActions, Role, SkillLevel): void {
        this.employeeService
        .addEmployee(firstName, lastName, WorkCenter, ShiftAssignments, CorrectiveActions, Role, SkillLevel)
            .subscribe(() => {
                this.router.navigate(['/indexEmployee']);
            });
    }

    ngOnInit(): void {
    }
}