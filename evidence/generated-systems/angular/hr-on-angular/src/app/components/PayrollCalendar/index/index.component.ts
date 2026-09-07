
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PayrollCalendarService } from '../../../services/PayrollCalendar.service';
import { PayrollCalendar } from '../../../models/PayrollCalendar';

@Component({
    selector: 'app-index-payrollCalendar',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPayrollCalendarComponent implements OnInit {

    payrollCalendars: PayrollCalendar[] = [];

    constructor(
        private router: Router,
        private service: PayrollCalendarService
) {}

    ngOnInit(): void {
        this.getPayrollCalendars();
}

    getPayrollCalendars(): void {
        this.service.getPayrollCalendars().subscribe((res) => {
        this.payrollCalendars = res;
    });
}

    deletePayrollCalendar(id: any): void {
        this.service.deletePayrollCalendar(id)
            .subscribe(() => {
                this.getPayrollCalendars();
            });
    }
}