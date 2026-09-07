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
      totalPremium: ['', Validators.required],
      ratingDate: ['', Validators.required],
      asBound: ['', Validators.required],
      Application: ['', ],
      UnderwritingDecisions: ['', ],
      Policy: ['', ]
        });
    }

    
    addQuote(quoteNumber, totalPremium, ratingDate, asBound, Application, UnderwritingDecisions, Policy): void {
        this.quoteService
        .addQuote(quoteNumber, totalPremium, ratingDate, asBound, Application, UnderwritingDecisions, Policy)
            .subscribe(() => {
                this.router.navigate(['/indexQuote']);
            });
    }

    ngOnInit(): void {
    }
}