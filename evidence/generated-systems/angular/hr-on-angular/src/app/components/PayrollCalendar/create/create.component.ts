import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PayrollCalendarService } from '../../../services/PayrollCalendar.service';
import { PayrollCalendar } from '../../../models/PayrollCalendar';
import { SubBaseComponent } from '../../PayrollCalendar/sub.base.component';

@Component({
    selector: 'app-create-payrollCalendar',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePayrollCalendarComponent extends SubBaseComponent implements OnInit {

    title = 'Add PayrollCalendar';

    payrollCalendarForm: FormGroup;
    payrollCalendar: PayrollCalendar;

    constructor( http: HttpClient,
        private payrollCalendarService: PayrollCalendarService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.payrollCalendarForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      country: ['', Validators.required],
      Organization: ['', ],
      PayrollRuns: ['', ],
      Employees: ['', ],
      PayFrequency: ['', ]
        });
    }

    
    addPayrollCalendar(name, country, Organization, PayrollRuns, Employees, PayFrequency): void {
        this.payrollCalendarService
        .addPayrollCalendar(name, country, Organization, PayrollRuns, Employees, PayFrequency)
            .subscribe(() => {
                this.router.navigate(['/indexPayrollCalendar']);
            });
    }

    ngOnInit(): void {
    }
}