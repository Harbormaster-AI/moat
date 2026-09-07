import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { EmployeeService } from '../../../services/Employee.service';
import { SubBaseComponent } from '../../Employee/sub.base.component';


@Component({
    selector: 'app-edit-employee',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditEmployeeComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Employee';

    employeeForm: FormGroup;
    employee: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: EmployeeService,
        private fb: FormBuilder
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

    
    updateEmployee(employeeNumber, name, workEmail, workPhone, dateOfHire, nationalId, Manager, DirectReports, Department, PrimaryLocation, CostCenter, EmploymentAssignments, Contracts, BenefitEnrollments, Timesheets, LeaveRequests, PerformanceReviews, TrainingEnrollments, WorkAuthorizations, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateEmployee(employeeNumber, name, workEmail, workPhone, dateOfHire, nationalId, Manager, DirectReports, Department, PrimaryLocation, CostCenter, EmploymentAssignments, Contracts, BenefitEnrollments, Timesheets, LeaveRequests, PerformanceReviews, TrainingEnrollments, WorkAuthorizations, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexEmployee']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getEmployee(params['id']).subscribe(res => {
                this.employee = res;
            });
        });
    }
}