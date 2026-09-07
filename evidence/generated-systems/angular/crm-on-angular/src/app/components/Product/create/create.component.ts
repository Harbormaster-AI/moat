import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ProductService } from '../../../services/Product.service';
import { Product } from '../../../models/Product';
import { SubBaseComponent } from '../../Product/sub.base.component';

@Component({
    selector: 'app-create-product',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateProductComponent extends SubBaseComponent implements OnInit {

    title = 'Add Product';

    productForm: FormGroup;
    product: Product;

    constructor( http: HttpClient,
        private productService: ProductService,
        private fb: FormBuilder,
        private router: Router
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

    
    addProduct(sku, name, asActive, standardPrice, description, Organization, PriceBookEntries, OpportunityLineItems, QuoteLineItems, OrderItems, ProductType, Uom): void {
        this.productService
        .addProduct(sku, name, asActive, standardPrice, description, Organization, PriceBookEntries, OpportunityLineItems, QuoteLineItems, OrderItems, ProductType, Uom)
            .subscribe(() => {
                this.router.navigate(['/indexProduct']);
            });
    }

    ngOnInit(): void {
    }
}