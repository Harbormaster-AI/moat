import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LoanApplicationService } from '../../../services/LoanApplication.service';
import { LoanApplication } from '../../../models/LoanApplication';
import { SubBaseComponent } from '../../LoanApplication/sub.base.component';

@Component({
    selector: 'app-create-loanApplication',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLoanApplicationComponent extends SubBaseComponent implements OnInit {

    title = 'Add LoanApplication';

    loanApplicationForm: FormGroup;
    loanApplication: LoanApplication;

    constructor( http: HttpClient,
        private loanApplicationService: LoanApplicationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.loanApplicationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  applicationNumber: ['', Validators.required],
      amountRequested: ['', Validators.required],
      termMonths: ['', Validators.required],
      submittedAt: ['', Validators.required],
      Customer: ['', ],
      RiskAssessment: ['', ],
      Loan: ['', ],
      Product: ['', ],
      Purpose: ['', ],
      Status: ['', ]
        });
    }

    
    addLoanApplication(applicationNumber, amountRequested, termMonths, submittedAt, Customer, RiskAssessment, Loan, Product, Purpose, Status): void {
        this.loanApplicationService
        .addLoanApplication(applicationNumber, amountRequested, termMonths, submittedAt, Customer, RiskAssessment, Loan, Product, Purpose, Status)
            .subscribe(() => {
                this.router.navigate(['/indexLoanApplication']);
            });
    }

    ngOnInit(): void {
    }
}