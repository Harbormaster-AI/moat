import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { FXQuoteService } from '../../../services/FXQuote.service';
import { FXQuote } from '../../../models/FXQuote';
import { SubBaseComponent } from '../../FXQuote/sub.base.component';

@Component({
    selector: 'app-create-fXQuote',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateFXQuoteComponent extends SubBaseComponent implements OnInit {

    title = 'Add FXQuote';

    fXQuoteForm: FormGroup;
    fXQuote: FXQuote;

    constructor( http: HttpClient,
        private fXQuoteService: FXQuoteService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.fXQuoteForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  baseCurrency: ['', Validators.required],
      quoteCurrency: ['', Validators.required],
      rate: ['', Validators.required],
      quotedAt: ['', Validators.required],
      expiresAt: ['', Validators.required],
      RequestedBy: ['', ],
      PriceType: ['', ]
        });
    }

    
    addFXQuote(baseCurrency, quoteCurrency, rate, quotedAt, expiresAt, RequestedBy, PriceType): void {
        this.fXQuoteService
        .addFXQuote(baseCurrency, quoteCurrency, rate, quotedAt, expiresAt, RequestedBy, PriceType)
            .subscribe(() => {
                this.router.navigate(['/indexFXQuote']);
            });
    }

    ngOnInit(): void {
    }
}