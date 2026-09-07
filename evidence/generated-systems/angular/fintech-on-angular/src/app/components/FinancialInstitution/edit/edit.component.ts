import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { FinancialInstitutionService } from '../../../services/FinancialInstitution.service';
import { SubBaseComponent } from '../../FinancialInstitution/sub.base.component';


@Component({
    selector: 'app-edit-financialInstitution',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditFinancialInstitutionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit FinancialInstitution';

    financialInstitutionForm: FormGroup;
    financialInstitution: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: FinancialInstitutionService,
        private fb: FormBuilder
) {
        super(http);
        this.financialInstitutionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      legalName: ['', Validators.required],
      countryOfIncorporation: ['', Validators.required],
      bic: ['', Validators.required],
      website: ['', Validators.required],
      Branches: ['', ],
      Customers: ['', ],
      ProductOfferings: ['', ],
      PaymentProcessors: ['', ],
      CompliancePolicies: ['', ]
        });
    }

    
    updateFinancialInstitution(name, legalName, countryOfIncorporation, bic, website, Branches, Customers, ProductOfferings, PaymentProcessors, CompliancePolicies): void {
        this.route.params.subscribe((params) => {

                        this.service.updateFinancialInstitution(name, legalName, countryOfIncorporation, bic, website, Branches, Customers, ProductOfferings, PaymentProcessors, CompliancePolicies, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexFinancialInstitution']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getFinancialInstitution(params['id']).subscribe(res => {
                this.financialInstitution = res;
            });
        });
    }
}