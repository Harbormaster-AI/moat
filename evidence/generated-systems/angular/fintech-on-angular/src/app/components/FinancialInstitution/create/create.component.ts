import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { FinancialInstitutionService } from '../../../services/FinancialInstitution.service';
import { FinancialInstitution } from '../../../models/FinancialInstitution';
import { SubBaseComponent } from '../../FinancialInstitution/sub.base.component';

@Component({
    selector: 'app-create-financialInstitution',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateFinancialInstitutionComponent extends SubBaseComponent implements OnInit {

    title = 'Add FinancialInstitution';

    financialInstitutionForm: FormGroup;
    financialInstitution: FinancialInstitution;

    constructor( http: HttpClient,
        private financialInstitutionService: FinancialInstitutionService,
        private fb: FormBuilder,
        private router: Router
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

    
    addFinancialInstitution(name, legalName, countryOfIncorporation, bic, website, Branches, Customers, ProductOfferings, PaymentProcessors, CompliancePolicies): void {
        this.financialInstitutionService
        .addFinancialInstitution(name, legalName, countryOfIncorporation, bic, website, Branches, Customers, ProductOfferings, PaymentProcessors, CompliancePolicies)
            .subscribe(() => {
                this.router.navigate(['/indexFinancialInstitution']);
            });
    }

    ngOnInit(): void {
    }
}