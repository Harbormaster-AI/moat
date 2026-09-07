import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PriceBookEntryService } from '../../../services/PriceBookEntry.service';
import { SubBaseComponent } from '../../PriceBookEntry/sub.base.component';


@Component({
    selector: 'app-edit-priceBookEntry',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPriceBookEntryComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PriceBookEntry';

    priceBookEntryForm: FormGroup;
    priceBookEntry: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PriceBookEntryService,
        private fb: FormBuilder
) {
        super(http);
        this.priceBookEntryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  unitPrice: ['', Validators.required],
      effectiveDate: ['', Validators.required],
      expirationDate: ['', Validators.required],
      asActive: ['', Validators.required],
      PriceBook: ['', ],
      Product: ['', ]
        });
    }

    
    updatePriceBookEntry(unitPrice, effectiveDate, expirationDate, asActive, PriceBook, Product): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePriceBookEntry(unitPrice, effectiveDate, expirationDate, asActive, PriceBook, Product, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPriceBookEntry']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPriceBookEntry(params['id']).subscribe(res => {
                this.priceBookEntry = res;
            });
        });
    }
}