import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PriceBookService } from '../../../services/PriceBook.service';
import { SubBaseComponent } from '../../PriceBook/sub.base.component';


@Component({
    selector: 'app-edit-priceBook',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPriceBookComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PriceBook';

    priceBookForm: FormGroup;
    priceBook: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PriceBookService,
        private fb: FormBuilder
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

    
    updatePriceBook(name, asActive, description, Organization, Entries, Quotes, Orders): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePriceBook(name, asActive, description, Organization, Entries, Quotes, Orders, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPriceBook']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPriceBook(params['id']).subscribe(res => {
                this.priceBook = res;
            });
        });
    }
}