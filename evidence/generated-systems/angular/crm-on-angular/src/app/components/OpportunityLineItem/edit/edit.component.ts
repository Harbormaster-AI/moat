import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { OpportunityLineItemService } from '../../../services/OpportunityLineItem.service';
import { SubBaseComponent } from '../../OpportunityLineItem/sub.base.component';


@Component({
    selector: 'app-edit-opportunityLineItem',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditOpportunityLineItemComponent extends SubBaseComponent implements OnInit {

    title = 'Edit OpportunityLineItem';

    opportunityLineItemForm: FormGroup;
    opportunityLineItem: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: OpportunityLineItemService,
        private fb: FormBuilder
) {
        super(http);
        this.opportunityLineItemForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  quantity: ['', Validators.required],
      unitPrice: ['', Validators.required],
      discountPercent: ['', Validators.required],
      totalPrice: ['', Validators.required],
      Opportunity: ['', ],
      Product: ['', ],
      PriceBookEntry: ['', ]
        });
    }

    
    updateOpportunityLineItem(quantity, unitPrice, discountPercent, totalPrice, Opportunity, Product, PriceBookEntry): void {
        this.route.params.subscribe((params) => {

                        this.service.updateOpportunityLineItem(quantity, unitPrice, discountPercent, totalPrice, Opportunity, Product, PriceBookEntry, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexOpportunityLineItem']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getOpportunityLineItem(params['id']).subscribe(res => {
                this.opportunityLineItem = res;
            });
        });
    }
}