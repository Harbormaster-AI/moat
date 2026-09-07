import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TransactionService } from '../../../services/Transaction.service';
import { Transaction } from '../../../models/Transaction';
import { SubBaseComponent } from '../../Transaction/sub.base.component';

@Component({
    selector: 'app-create-transaction',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTransactionComponent extends SubBaseComponent implements OnInit {

    title = 'Add Transaction';

    transactionForm: FormGroup;
    transaction: Transaction;

    constructor( http: HttpClient,
        private transactionService: TransactionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.transactionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  amount: ['', Validators.required],
      fee: ['', Validators.required],
      exchangeRate: ['', Validators.required],
      createdAt: ['', Validators.required],
      completedAt: ['', Validators.required],
      narrative: ['', Validators.required],
      Account: ['', ],
      Wallet: ['', ],
      PaymentOrder: ['', ],
      Merchant: ['', ],
      Card: ['', ],
      RelatedTransactions: ['', ],
      Alerts: ['', ],
      TransactionType: ['', ],
      Status: ['', ]
        });
    }

    
    addTransaction(amount, fee, exchangeRate, createdAt, completedAt, narrative, Account, Wallet, PaymentOrder, Merchant, Card, RelatedTransactions, Alerts, TransactionType, Status): void {
        this.transactionService
        .addTransaction(amount, fee, exchangeRate, createdAt, completedAt, narrative, Account, Wallet, PaymentOrder, Merchant, Card, RelatedTransactions, Alerts, TransactionType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexTransaction']);
            });
    }

    ngOnInit(): void {
    }
}