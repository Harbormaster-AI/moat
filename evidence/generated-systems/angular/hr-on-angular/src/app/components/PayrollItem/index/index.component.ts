
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PayrollItemService } from '../../../services/PayrollItem.service';
import { PayrollItem } from '../../../models/PayrollItem';

@Component({
    selector: 'app-index-payrollItem',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPayrollItemComponent implements OnInit {

    payrollItems: PayrollItem[] = [];

    constructor(
        private router: Router,
        private service: PayrollItemService
) {}

    ngOnInit(): void {
        this.getPayrollItems();
}

    getPayrollItems(): void {
        this.service.getPayrollItems().subscribe((res) => {
        this.payrollItems = res;
    });
}

    deletePayrollItem(id: any): void {
        this.service.deletePayrollItem(id)
            .subscribe(() => {
                this.getPayrollItems();
            });
    }
}