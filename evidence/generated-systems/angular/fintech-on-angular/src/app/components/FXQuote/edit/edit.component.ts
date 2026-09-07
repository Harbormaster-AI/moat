import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { FXQuoteService } from '../../../services/FXQuote.service';
import { SubBaseComponent } from '../../FXQuote/sub.base.component';


@Component({
    selector: 'app-edit-fXQuote',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditFXQuoteComponent extends SubBaseComponent implements OnInit {

    title = 'Edit FXQuote';

    fXQuoteForm: FormGroup;
    fXQuote: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: FXQuoteService,
        private fb: FormBuilder
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

    
    updateFXQuote(baseCurrency, quoteCurrency, rate, quotedAt, expiresAt, RequestedBy, PriceType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateFXQuote(baseCurrency, quoteCurrency, rate, quotedAt, expiresAt, RequestedBy, PriceType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexFXQuote']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getFXQuote(params['id']).subscribe(res => {
                this.fXQuote = res;
            });
        });
    }
}