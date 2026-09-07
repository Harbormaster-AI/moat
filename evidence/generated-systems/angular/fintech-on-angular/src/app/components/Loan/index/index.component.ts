
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoanService } from '../../../services/Loan.service';
import { Loan } from '../../../models/Loan';

@Component({
    selector: 'app-index-loan',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLoanComponent implements OnInit {

    loans: Loan[] = [];

    constructor(
        private router: Router,
        private service: LoanService
) {}

    ngOnInit(): void {
        this.getLoans();
}

    getLoans(): void {
        this.service.getLoans().subscribe((res) => {
        this.loans = res;
    });
}

    deleteLoan(id: any): void {
        this.service.deleteLoan(id)
            .subscribe(() => {
                this.getLoans();
            });
    }
}