import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CompensationPackageService } from '../../../services/CompensationPackage.service';
import { SubBaseComponent } from '../../CompensationPackage/sub.base.component';


@Component({
    selector: 'app-edit-compensationPackage',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCompensationPackageComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CompensationPackage';

    compensationPackageForm: FormGroup;
    compensationPackage: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CompensationPackageService,
        private fb: FormBuilder
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

    
    updateCompensationPackage(effectiveFrom, effectiveTo, currency, Contract, SalaryComponents, BonusPlans, EquityGrants): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCompensationPackage(effectiveFrom, effectiveTo, currency, Contract, SalaryComponents, BonusPlans, EquityGrants, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCompensationPackage']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCompensationPackage(params['id']).subscribe(res => {
                this.compensationPackage = res;
            });
        });
    }
}