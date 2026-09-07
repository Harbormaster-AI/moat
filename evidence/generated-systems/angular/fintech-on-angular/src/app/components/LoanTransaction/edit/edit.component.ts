import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LoanTransactionService } from '../../../services/LoanTransaction.service';
import { SubBaseComponent } from '../../LoanTransaction/sub.base.component';


@Component({
    selector: 'app-edit-loanTransaction',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLoanTransactionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit LoanTransaction';

    loanTransactionForm: FormGroup;
    loanTransaction: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LoanTransactionService,
        private fb: FormBuilder
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

    
    updateLoanTransaction(transactionId, amount, postingDate, Loan, Type, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLoanTransaction(transactionId, amount, postingDate, Loan, Type, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLoanTransaction']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLoanTransaction(params['id']).subscribe(res => {
                this.loanTransaction = res;
            });
        });
    }
}