import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PayrollRunService } from '../../../services/PayrollRun.service';
import { SubBaseComponent } from '../../PayrollRun/sub.base.component';


@Component({
    selector: 'app-edit-payrollRun',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPayrollRunComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PayrollRun';

    payrollRunForm: FormGroup;
    payrollRun: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PayrollRunService,
        private fb: FormBuilder
) {
        super(http);
        this.payrollRunForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  runNumber: ['', Validators.required],
      periodStart: ['', Validators.required],
      periodEnd: ['', Validators.required],
      paymentDate: ['', Validators.required],
      PayrollCalendar: ['', ],
      PayrollItems: ['', ],
      Status: ['', ]
        });
    }

    
    updatePayrollRun(runNumber, periodStart, periodEnd, paymentDate, PayrollCalendar, PayrollItems, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePayrollRun(runNumber, periodStart, periodEnd, paymentDate, PayrollCalendar, PayrollItems, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPayrollRun']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPayrollRun(params['id']).subscribe(res => {
                this.payrollRun = res;
            });
        });
    }
}