import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PayrollRunService } from '../../../services/PayrollRun.service';
import { PayrollRun } from '../../../models/PayrollRun';
import { SubBaseComponent } from '../../PayrollRun/sub.base.component';

@Component({
    selector: 'app-create-payrollRun',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePayrollRunComponent extends SubBaseComponent implements OnInit {

    title = 'Add PayrollRun';

    payrollRunForm: FormGroup;
    payrollRun: PayrollRun;

    constructor( http: HttpClient,
        private payrollRunService: PayrollRunService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPayrollRun(runNumber, periodStart, periodEnd, paymentDate, PayrollCalendar, PayrollItems, Status): void {
        this.payrollRunService
        .addPayrollRun(runNumber, periodStart, periodEnd, paymentDate, PayrollCalendar, PayrollItems, Status)
            .subscribe(() => {
                this.router.navigate(['/indexPayrollRun']);
            });
    }

    ngOnInit(): void {
    }
}