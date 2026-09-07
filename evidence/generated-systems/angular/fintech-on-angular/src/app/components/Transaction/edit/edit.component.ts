import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TransactionService } from '../../../services/Transaction.service';
import { SubBaseComponent } from '../../Transaction/sub.base.component';


@Component({
    selector: 'app-edit-transaction',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTransactionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Transaction';

    transactionForm: FormGroup;
    transaction: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TransactionService,
        private fb: FormBuilder
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

    
    updateTransaction(amount, fee, exchangeRate, createdAt, completedAt, narrative, Account, Wallet, PaymentOrder, Merchant, Card, RelatedTransactions, Alerts, TransactionType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTransaction(amount, fee, exchangeRate, createdAt, completedAt, narrative, Account, Wallet, PaymentOrder, Merchant, Card, RelatedTransactions, Alerts, TransactionType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTransaction']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTransaction(params['id']).subscribe(res => {
                this.transaction = res;
            });
        });
    }
}