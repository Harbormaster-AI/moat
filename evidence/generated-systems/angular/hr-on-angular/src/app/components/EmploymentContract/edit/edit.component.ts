import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { EmploymentContractService } from '../../../services/EmploymentContract.service';
import { SubBaseComponent } from '../../EmploymentContract/sub.base.component';


@Component({
    selector: 'app-edit-employmentContract',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditEmploymentContractComponent extends SubBaseComponent implements OnInit {

    title = 'Edit EmploymentContract';

    employmentContractForm: FormGroup;
    employmentContract: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: EmploymentContractService,
        private fb: FormBuilder
) {
        super(http);
        this.employmentContractForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  contractNumber: ['', Validators.required],
      startDate: ['', Validators.required],
      endDate: ['', Validators.required],
      workHoursPerWeek: ['', Validators.required],
      Employee: ['', ],
      CompensationPackage: ['', ],
      WorkSchedule: ['', ],
      Location: ['', ],
      PayrollCalendar: ['', ],
      EmploymentType: ['', ],
      Status: ['', ],
      PayFrequency: ['', ]
        });
    }

    
    updateEmploymentContract(contractNumber, startDate, endDate, workHoursPerWeek, Employee, CompensationPackage, WorkSchedule, Location, PayrollCalendar, EmploymentType, Status, PayFrequency): void {
        this.route.params.subscribe((params) => {

                        this.service.updateEmploymentContract(contractNumber, startDate, endDate, workHoursPerWeek, Employee, CompensationPackage, WorkSchedule, Location, PayrollCalendar, EmploymentType, Status, PayFrequency, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexEmploymentContract']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getEmploymentContract(params['id']).subscribe(res => {
                this.employmentContract = res;
            });
        });
    }
}