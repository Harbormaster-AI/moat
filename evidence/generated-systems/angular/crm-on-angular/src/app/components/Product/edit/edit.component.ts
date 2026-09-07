import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ProductService } from '../../../services/Product.service';
import { SubBaseComponent } from '../../Product/sub.base.component';


@Component({
    selector: 'app-edit-product',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditProductComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Product';

    productForm: FormGroup;
    product: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ProductService,
        private fb: FormBuilder
) {
        super(http);
        this.productForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  sku: ['', Validators.required],
      name: ['', Validators.required],
      asActive: ['', Validators.required],
      standardPrice: ['', Validators.required],
      description: ['', Validators.required],
      Organization: ['', ],
      PriceBookEntries: ['', ],
      OpportunityLineItems: ['', ],
      QuoteLineItems: ['', ],
      OrderItems: ['', ],
      ProductType: ['', ],
      Uom: ['', ]
        });
    }

    
    updateProduct(sku, name, asActive, standardPrice, description, Organization, PriceBookEntries, OpportunityLineItems, QuoteLineItems, OrderItems, ProductType, Uom): void {
        this.route.params.subscribe((params) => {

                        this.service.updateProduct(sku, name, asActive, standardPrice, description, Organization, PriceBookEntries, OpportunityLineItems, QuoteLineItems, OrderItems, ProductType, Uom, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexProduct']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getProduct(params['id']).subscribe(res => {
                this.product = res;
            });
        });
    }
}