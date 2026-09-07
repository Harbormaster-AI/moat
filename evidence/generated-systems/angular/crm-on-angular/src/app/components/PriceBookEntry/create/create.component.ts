import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PriceBookEntryService } from '../../../services/PriceBookEntry.service';
import { PriceBookEntry } from '../../../models/PriceBookEntry';
import { SubBaseComponent } from '../../PriceBookEntry/sub.base.component';

@Component({
    selector: 'app-create-priceBookEntry',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePriceBookEntryComponent extends SubBaseComponent implements OnInit {

    title = 'Add PriceBookEntry';

    priceBookEntryForm: FormGroup;
    priceBookEntry: PriceBookEntry;

    constructor( http: HttpClient,
        private priceBookEntryService: PriceBookEntryService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPriceBookEntry(unitPrice, effectiveDate, expirationDate, asActive, PriceBook, Product): void {
        this.priceBookEntryService
        .addPriceBookEntry(unitPrice, effectiveDate, expirationDate, asActive, PriceBook, Product)
            .subscribe(() => {
                this.router.navigate(['/indexPriceBookEntry']);
            });
    }

    ngOnInit(): void {
    }
}