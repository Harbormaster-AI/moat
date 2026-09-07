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
                  employeeNumber: ['', Validators.required],
      name: ['', Validators.required],
      workEmail: ['', Validators.required],
      workPhone: ['', Validators.required],
      dateOfHire: ['', Validators.required],
      nationalId: ['', Validators.required],
      Manager: ['', ],
      DirectReports: ['', ],
      Department: ['', ],
      PrimaryLocation: ['', ],
      CostCenter: ['', ],
      EmploymentAssignments: ['', ],
      Contracts: ['', ],
      BenefitEnrollments: ['', ],
      Timesheets: ['', ],
      LeaveRequests: ['', ],
      PerformanceReviews: ['', ],
      TrainingEnrollments: ['', ],
      WorkAuthorizations: ['', ],
      Status: ['', ]
        });
    }

    
    addEmployee(employeeNumber, name, workEmail, workPhone, dateOfHire, nationalId, Manager, DirectReports, Department, PrimaryLocation, CostCenter, EmploymentAssignments, Contracts, BenefitEnrollments, Timesheets, LeaveRequests, PerformanceReviews, TrainingEnrollments, WorkAuthorizations, Status): void {
        this.employeeService
        .addEmployee(employeeNumber, name, workEmail, workPhone, dateOfHire, nationalId, Manager, DirectReports, Department, PrimaryLocation, CostCenter, EmploymentAssignments, Contracts, BenefitEnrollments, Timesheets, LeaveRequests, PerformanceReviews, TrainingEnrollments, WorkAuthorizations, Status)
            .subscribe(() => {
                this.router.navigate(['/indexEmployee']);
            });
    }

    ngOnInit(): void {
    }
}