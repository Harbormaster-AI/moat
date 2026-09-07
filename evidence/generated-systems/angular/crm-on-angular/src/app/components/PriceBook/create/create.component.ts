import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PriceBookService } from '../../../services/PriceBook.service';
import { PriceBook } from '../../../models/PriceBook';
import { SubBaseComponent } from '../../PriceBook/sub.base.component';

@Component({
    selector: 'app-create-priceBook',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePriceBookComponent extends SubBaseComponent implements OnInit {

    title = 'Add PriceBook';

    priceBookForm: FormGroup;
    priceBook: PriceBook;

    constructor( http: HttpClient,
        private priceBookService: PriceBookService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.priceBookForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      asActive: ['', Validators.required],
      description: ['', Validators.required],
      Organization: ['', ],
      Entries: ['', ],
      Quotes: ['', ],
      Orders: ['', ]
        });
    }

    
    addPriceBook(name, asActive, description, Organization, Entries, Quotes, Orders): void {
        this.priceBookService
        .addPriceBook(name, asActive, description, Organization, Entries, Quotes, Orders)
            .subscribe(() => {
                this.router.navigate(['/indexPriceBook']);
            });
    }

    ngOnInit(): void {
    }
}