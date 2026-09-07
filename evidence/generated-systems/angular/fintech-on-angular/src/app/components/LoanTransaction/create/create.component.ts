import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LoanTransactionService } from '../../../services/LoanTransaction.service';
import { LoanTransaction } from '../../../models/LoanTransaction';
import { SubBaseComponent } from '../../LoanTransaction/sub.base.component';

@Component({
    selector: 'app-create-loanTransaction',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLoanTransactionComponent extends SubBaseComponent implements OnInit {

    title = 'Add LoanTransaction';

    loanTransactionForm: FormGroup;
    loanTransaction: LoanTransaction;

    constructor( http: HttpClient,
        private loanTransactionService: LoanTransactionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.loanTransactionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  transactionId: ['', Validators.required],
      amount: ['', Validators.required],
      postingDate: ['', Validators.required],
      Loan: ['', ],
      Type: ['', ],
      Status: ['', ]
        });
    }

    
    addLoanTransaction(transactionId, amount, postingDate, Loan, Type, Status): void {
        this.loanTransactionService
        .addLoanTransaction(transactionId, amount, postingDate, Loan, Type, Status)
            .subscribe(() => {
                this.router.navigate(['/indexLoanTransaction']);
            });
    }

    ngOnInit(): void {
    }
}