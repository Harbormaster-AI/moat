import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LoanService } from '../../../services/Loan.service';
import { Loan } from '../../../models/Loan';
import { SubBaseComponent } from '../../Loan/sub.base.component';

@Component({
    selector: 'app-create-loan',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLoanComponent extends SubBaseComponent implements OnInit {

    title = 'Add Loan';

    loanForm: FormGroup;
    loan: Loan;

    constructor( http: HttpClient,
        private loanService: LoanService,
        private fb: FormBuilder,
        private router: Router
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

    
    addLoan(loanNumber, principal, interestRate, originationDate, maturityDate, Customer, Schedule, Collateral, Transactions, RateType, Status): void {
        this.loanService
        .addLoan(loanNumber, principal, interestRate, originationDate, maturityDate, Customer, Schedule, Collateral, Transactions, RateType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexLoan']);
            });
    }

    ngOnInit(): void {
    }
}