import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CompensationPackageService } from '../../../services/CompensationPackage.service';
import { CompensationPackage } from '../../../models/CompensationPackage';
import { SubBaseComponent } from '../../CompensationPackage/sub.base.component';

@Component({
    selector: 'app-create-compensationPackage',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCompensationPackageComponent extends SubBaseComponent implements OnInit {

    title = 'Add CompensationPackage';

    compensationPackageForm: FormGroup;
    compensationPackage: CompensationPackage;

    constructor( http: HttpClient,
        private compensationPackageService: CompensationPackageService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.compensationPackageForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  effectiveFrom: ['', Validators.required],
      effectiveTo: ['', Validators.required],
      currency: ['', Validators.required],
      Contract: ['', ],
      SalaryComponents: ['', ],
      BonusPlans: ['', ],
      EquityGrants: ['', ]
        });
    }

    
    addCompensationPackage(effectiveFrom, effectiveTo, currency, Contract, SalaryComponents, BonusPlans, EquityGrants): void {
        this.compensationPackageService
        .addCompensationPackage(effectiveFrom, effectiveTo, currency, Contract, SalaryComponents, BonusPlans, EquityGrants)
            .subscribe(() => {
                this.router.navigate(['/indexCompensationPackage']);
            });
    }

    ngOnInit(): void {
    }
}