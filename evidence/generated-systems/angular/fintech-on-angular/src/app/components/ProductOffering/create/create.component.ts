import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ProductOfferingService } from '../../../services/ProductOffering.service';
import { ProductOffering } from '../../../models/ProductOffering';
import { SubBaseComponent } from '../../ProductOffering/sub.base.component';

@Component({
    selector: 'app-create-productOffering',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateProductOfferingComponent extends SubBaseComponent implements OnInit {

    title = 'Add ProductOffering';

    productOfferingForm: FormGroup;
    productOffering: ProductOffering;

    constructor( http: HttpClient,
        private productOfferingService: ProductOfferingService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.productOfferingForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      productCode: ['', Validators.required],
      Institution: ['', ],
      PricingPlans: ['', ],
      Category: ['', ]
        });
    }

    
    addProductOffering(name, productCode, Institution, PricingPlans, Category): void {
        this.productOfferingService
        .addProductOffering(name, productCode, Institution, PricingPlans, Category)
            .subscribe(() => {
                this.router.navigate(['/indexProductOffering']);
            });
    }

    ngOnInit(): void {
    }
}