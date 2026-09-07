import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InsuranceProductService } from '../../../services/InsuranceProduct.service';
import { InsuranceProduct } from '../../../models/InsuranceProduct';
import { SubBaseComponent } from '../../InsuranceProduct/sub.base.component';

@Component({
    selector: 'app-create-insuranceProduct',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInsuranceProductComponent extends SubBaseComponent implements OnInit {

    title = 'Add InsuranceProduct';

    insuranceProductForm: FormGroup;
    insuranceProduct: InsuranceProduct;

    constructor( http: HttpClient,
        private insuranceProductService: InsuranceProductService,
        private fb: FormBuilder,
        private router: Router
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

    
    addInsuranceProduct(name, productCode, Insurer, CoverageDefinitions, LineOfBusiness): void {
        this.insuranceProductService
        .addInsuranceProduct(name, productCode, Insurer, CoverageDefinitions, LineOfBusiness)
            .subscribe(() => {
                this.router.navigate(['/indexInsuranceProduct']);
            });
    }

    ngOnInit(): void {
    }
}