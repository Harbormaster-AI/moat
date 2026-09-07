import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LoanService } from '../../../services/Loan.service';
import { SubBaseComponent } from '../../Loan/sub.base.component';


@Component({
    selector: 'app-edit-loan',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLoanComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Loan';

    loanForm: FormGroup;
    loan: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LoanService,
        private fb: FormBuilder
) {
        super(http);
        this.loanForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  loanNumber: ['', Validators.required],
      principal: ['', Validators.required],
      interestRate: ['', Validators.required],
      originationDate: ['', Validators.required],
      maturityDate: ['', Validators.required],
      Customer: ['', ],
      Schedule: ['', ],
      Collateral: ['', ],
      Transactions: ['', ],
      RateType: ['', ],
      Status: ['', ]
        });
    }

    
    updateLoan(loanNumber, principal, interestRate, originationDate, maturityDate, Customer, Schedule, Collateral, Transactions, RateType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLoan(loanNumber, principal, interestRate, originationDate, maturityDate, Customer, Schedule, Collateral, Transactions, RateType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLoan']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLoan(params['id']).subscribe(res => {
                this.loan = res;
            });
        });
    }
}