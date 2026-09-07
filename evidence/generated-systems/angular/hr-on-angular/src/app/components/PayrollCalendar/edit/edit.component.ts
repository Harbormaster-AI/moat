import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PayrollCalendarService } from '../../../services/PayrollCalendar.service';
import { SubBaseComponent } from '../../PayrollCalendar/sub.base.component';


@Component({
    selector: 'app-edit-payrollCalendar',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPayrollCalendarComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PayrollCalendar';

    payrollCalendarForm: FormGroup;
    payrollCalendar: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PayrollCalendarService,
        private fb: FormBuilder
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

    
    updatePayrollCalendar(name, country, Organization, PayrollRuns, Employees, PayFrequency): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePayrollCalendar(name, country, Organization, PayrollRuns, Employees, PayFrequency, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPayrollCalendar']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPayrollCalendar(params['id']).subscribe(res => {
                this.payrollCalendar = res;
            });
        });
    }
}