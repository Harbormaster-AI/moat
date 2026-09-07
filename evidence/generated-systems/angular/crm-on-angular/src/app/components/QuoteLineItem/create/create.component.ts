import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { QuoteLineItemService } from '../../../services/QuoteLineItem.service';
import { QuoteLineItem } from '../../../models/QuoteLineItem';
import { SubBaseComponent } from '../../QuoteLineItem/sub.base.component';

@Component({
    selector: 'app-create-quoteLineItem',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateQuoteLineItemComponent extends SubBaseComponent implements OnInit {

    title = 'Add QuoteLineItem';

    quoteLineItemForm: FormGroup;
    quoteLineItem: QuoteLineItem;

    constructor( http: HttpClient,
        private quoteLineItemService: QuoteLineItemService,
        private fb: FormBuilder,
        private router: Router
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

    
    addQuoteLineItem(quantity, unitPrice, discountAmount, taxAmount, totalAmount, Quote, Product, PriceBookEntry, OpportunityLineItem): void {
        this.quoteLineItemService
        .addQuoteLineItem(quantity, unitPrice, discountAmount, taxAmount, totalAmount, Quote, Product, PriceBookEntry, OpportunityLineItem)
            .subscribe(() => {
                this.router.navigate(['/indexQuoteLineItem']);
            });
    }

    ngOnInit(): void {
    }
}