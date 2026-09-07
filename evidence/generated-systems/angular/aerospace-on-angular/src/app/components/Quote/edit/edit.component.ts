import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { QuoteService } from '../../../services/Quote.service';
import { SubBaseComponent } from '../../Quote/sub.base.component';


@Component({
    selector: 'app-edit-quote',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditQuoteComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Quote';

    quoteForm: FormGroup;
    quote: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: QuoteService,
        private fb: FormBuilder
) {
        super(http);
        this.quoteForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  quoteNumber: ['', Validators.required],
      totalAmount: ['', Validators.required],
      AircraftOrder: ['', ]
        });
    }

    
    updateQuote(quoteNumber, totalAmount, AircraftOrder): void {
        this.route.params.subscribe((params) => {

                        this.service.updateQuote(quoteNumber, totalAmount, AircraftOrder, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexQuote']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getQuote(params['id']).subscribe(res => {
                this.quote = res;
            });
        });
    }
}