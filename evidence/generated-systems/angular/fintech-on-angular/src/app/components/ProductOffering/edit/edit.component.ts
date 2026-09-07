import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ProductOfferingService } from '../../../services/ProductOffering.service';
import { SubBaseComponent } from '../../ProductOffering/sub.base.component';


@Component({
    selector: 'app-edit-productOffering',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditProductOfferingComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ProductOffering';

    productOfferingForm: FormGroup;
    productOffering: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ProductOfferingService,
        private fb: FormBuilder
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

    
    updateProductOffering(name, productCode, Institution, PricingPlans, Category): void {
        this.route.params.subscribe((params) => {

                        this.service.updateProductOffering(name, productCode, Institution, PricingPlans, Category, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexProductOffering']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getProductOffering(params['id']).subscribe(res => {
                this.productOffering = res;
            });
        });
    }
}