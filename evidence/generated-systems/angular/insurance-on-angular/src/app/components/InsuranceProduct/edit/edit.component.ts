import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InsuranceProductService } from '../../../services/InsuranceProduct.service';
import { SubBaseComponent } from '../../InsuranceProduct/sub.base.component';


@Component({
    selector: 'app-edit-insuranceProduct',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInsuranceProductComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InsuranceProduct';

    insuranceProductForm: FormGroup;
    insuranceProduct: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InsuranceProductService,
        private fb: FormBuilder
) {
        super(http);
        this.insuranceProductForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      productCode: ['', Validators.required],
      Insurer: ['', ],
      CoverageDefinitions: ['', ],
      LineOfBusiness: ['', ]
        });
    }

    
    updateInsuranceProduct(name, productCode, Insurer, CoverageDefinitions, LineOfBusiness): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInsuranceProduct(name, productCode, Insurer, CoverageDefinitions, LineOfBusiness, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInsuranceProduct']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInsuranceProduct(params['id']).subscribe(res => {
                this.insuranceProduct = res;
            });
        });
    }
}