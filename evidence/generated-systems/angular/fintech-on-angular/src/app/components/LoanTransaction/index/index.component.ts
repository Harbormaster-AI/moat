
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoanTransactionService } from '../../../services/LoanTransaction.service';
import { LoanTransaction } from '../../../models/LoanTransaction';

@Component({
    selector: 'app-index-loanTransaction',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLoanTransactionComponent implements OnInit {

    loanTransactions: LoanTransaction[] = [];

    constructor(
        private router: Router,
        private service: LoanTransactionService
) {}

    ngOnInit(): void {
        this.getLoanTransactions();
}

    getLoanTransactions(): void {
        this.service.getLoanTransactions().subscribe((res) => {
        this.loanTransactions = res;
    });
}

    deleteLoanTransaction(id: any): void {
        this.service.deleteLoanTransaction(id)
            .subscribe(() => {
                this.getLoanTransactions();
            });
    }
}