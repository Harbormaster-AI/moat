import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { EmploymentContractService } from '../../../services/EmploymentContract.service';
import { EmploymentContract } from '../../../models/EmploymentContract';
import { SubBaseComponent } from '../../EmploymentContract/sub.base.component';

@Component({
    selector: 'app-create-employmentContract',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateEmploymentContractComponent extends SubBaseComponent implements OnInit {

    title = 'Add EmploymentContract';

    employmentContractForm: FormGroup;
    employmentContract: EmploymentContract;

    constructor( http: HttpClient,
        private employmentContractService: EmploymentContractService,
        private fb: FormBuilder,
        private router: Router
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

    
    addEmploymentContract(contractNumber, startDate, endDate, workHoursPerWeek, Employee, CompensationPackage, WorkSchedule, Location, PayrollCalendar, EmploymentType, Status, PayFrequency): void {
        this.employmentContractService
        .addEmploymentContract(contractNumber, startDate, endDate, workHoursPerWeek, Employee, CompensationPackage, WorkSchedule, Location, PayrollCalendar, EmploymentType, Status, PayFrequency)
            .subscribe(() => {
                this.router.navigate(['/indexEmploymentContract']);
            });
    }

    ngOnInit(): void {
    }
}