import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OpportunityLineItemService } from '../../../services/OpportunityLineItem.service';
import { OpportunityLineItem } from '../../../models/OpportunityLineItem';
import { SubBaseComponent } from '../../OpportunityLineItem/sub.base.component';

@Component({
    selector: 'app-create-opportunityLineItem',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateOpportunityLineItemComponent extends SubBaseComponent implements OnInit {

    title = 'Add OpportunityLineItem';

    opportunityLineItemForm: FormGroup;
    opportunityLineItem: OpportunityLineItem;

    constructor( http: HttpClient,
        private opportunityLineItemService: OpportunityLineItemService,
        private fb: FormBuilder,
        private router: Router
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

    
    addOpportunityLineItem(quantity, unitPrice, discountPercent, totalPrice, Opportunity, Product, PriceBookEntry): void {
        this.opportunityLineItemService
        .addOpportunityLineItem(quantity, unitPrice, discountPercent, totalPrice, Opportunity, Product, PriceBookEntry)
            .subscribe(() => {
                this.router.navigate(['/indexOpportunityLineItem']);
            });
    }

    ngOnInit(): void {
    }
}