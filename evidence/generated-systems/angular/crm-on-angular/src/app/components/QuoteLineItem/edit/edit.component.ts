import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { QuoteLineItemService } from '../../../services/QuoteLineItem.service';
import { SubBaseComponent } from '../../QuoteLineItem/sub.base.component';


@Component({
    selector: 'app-edit-quoteLineItem',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditQuoteLineItemComponent extends SubBaseComponent implements OnInit {

    title = 'Edit QuoteLineItem';

    quoteLineItemForm: FormGroup;
    quoteLineItem: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: QuoteLineItemService,
        private fb: FormBuilder
) {
        super(http);
        this.quoteLineItemForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  quantity: ['', Validators.required],
      unitPrice: ['', Validators.required],
      discountAmount: ['', Validators.required],
      taxAmount: ['', Validators.required],
      totalAmount: ['', Validators.required],
      Quote: ['', ],
      Product: ['', ],
      PriceBookEntry: ['', ],
      OpportunityLineItem: ['', ]
        });
    }

    
    updateQuoteLineItem(quantity, unitPrice, discountAmount, taxAmount, totalAmount, Quote, Product, PriceBookEntry, OpportunityLineItem): void {
        this.route.params.subscribe((params) => {

                        this.service.updateQuoteLineItem(quantity, unitPrice, discountAmount, taxAmount, totalAmount, Quote, Product, PriceBookEntry, OpportunityLineItem, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexQuoteLineItem']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getQuoteLineItem(params['id']).subscribe(res => {
                this.quoteLineItem = res;
            });
        });
    }
}