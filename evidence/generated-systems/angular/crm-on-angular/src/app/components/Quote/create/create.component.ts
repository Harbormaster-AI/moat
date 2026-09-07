import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { QuoteService } from '../../../services/Quote.service';
import { Quote } from '../../../models/Quote';
import { SubBaseComponent } from '../../Quote/sub.base.component';

@Component({
    selector: 'app-create-quote',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateQuoteComponent extends SubBaseComponent implements OnInit {

    title = 'Add Quote';

    quoteForm: FormGroup;
    quote: Quote;

    constructor( http: HttpClient,
        private quoteService: QuoteService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.quoteForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  quoteNumber: ['', Validators.required],
      validityStart: ['', Validators.required],
      validityEnd: ['', Validators.required],
      totalAmount: ['', Validators.required],
      discountPercent: ['', Validators.required],
      taxAmount: ['', Validators.required],
      shippingAmount: ['', Validators.required],
      Organization: ['', ],
      Account: ['', ],
      Opportunity: ['', ],
      Owner: ['', ],
      LineItems: ['', ],
      PriceBook: ['', ],
      Order: ['', ],
      Status: ['', ]
        });
    }

    
    addQuote(quoteNumber, validityStart, validityEnd, totalAmount, discountPercent, taxAmount, shippingAmount, Organization, Account, Opportunity, Owner, LineItems, PriceBook, Order, Status): void {
        this.quoteService
        .addQuote(quoteNumber, validityStart, validityEnd, totalAmount, discountPercent, taxAmount, shippingAmount, Organization, Account, Opportunity, Owner, LineItems, PriceBook, Order, Status)
            .subscribe(() => {
                this.router.navigate(['/indexQuote']);
            });
    }

    ngOnInit(): void {
    }
}