
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PayrollRunService } from '../../../services/PayrollRun.service';
import { PayrollRun } from '../../../models/PayrollRun';

@Component({
    selector: 'app-index-payrollRun',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPayrollRunComponent implements OnInit {

    payrollRuns: PayrollRun[] = [];

    constructor(
        private router: Router,
        private service: PayrollRunService
) {}

    ngOnInit(): void {
        this.getPayrollRuns();
}

    getPayrollRuns(): void {
        this.service.getPayrollRuns().subscribe((res) => {
        this.payrollRuns = res;
    });
}

    deletePayrollRun(id: any): void {
        this.service.deletePayrollRun(id)
            .subscribe(() => {
                this.getPayrollRuns();
            });
    }
}