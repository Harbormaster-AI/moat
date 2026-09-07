import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LoanApplicationService } from '../../../services/LoanApplication.service';
import { SubBaseComponent } from '../../LoanApplication/sub.base.component';


@Component({
    selector: 'app-edit-loanApplication',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLoanApplicationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit LoanApplication';

    loanApplicationForm: FormGroup;
    loanApplication: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LoanApplicationService,
        private fb: FormBuilder
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

    
    updateLoanApplication(applicationNumber, amountRequested, termMonths, submittedAt, Customer, RiskAssessment, Loan, Product, Purpose, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLoanApplication(applicationNumber, amountRequested, termMonths, submittedAt, Customer, RiskAssessment, Loan, Product, Purpose, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLoanApplication']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLoanApplication(params['id']).subscribe(res => {
                this.loanApplication = res;
            });
        });
    }
}