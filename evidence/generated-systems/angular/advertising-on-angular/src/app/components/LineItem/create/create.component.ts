import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LineItemService } from '../../../services/LineItem.service';
import { LineItem } from '../../../models/LineItem';
import { SubBaseComponent } from '../../LineItem/sub.base.component';

@Component({
    selector: 'app-create-lineItem',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLineItemComponent extends SubBaseComponent implements OnInit {

    title = 'Add LineItem';

    lineItemForm: FormGroup;
    lineItem: LineItem;

    constructor( http: HttpClient,
        private lineItemService: LineItemService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.lineItemForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      bidAmount: ['', Validators.required],
      dailyBudget: ['', Validators.required],
      frequencyCap: ['', Validators.required],
      Campaign: ['', ],
      Placements: ['', ],
      TargetingProfile: ['', ],
      Deal: ['', ],
      Creatives: ['', ],
      PerformanceMetrics: ['', ],
      Status: ['', ],
      PricingModel: ['', ],
      BidStrategy: ['', ],
      Pacing: ['', ]
        });
    }

    
    addLineItem(name, bidAmount, dailyBudget, frequencyCap, Campaign, Placements, TargetingProfile, Deal, Creatives, PerformanceMetrics, Status, PricingModel, BidStrategy, Pacing): void {
        this.lineItemService
        .addLineItem(name, bidAmount, dailyBudget, frequencyCap, Campaign, Placements, TargetingProfile, Deal, Creatives, PerformanceMetrics, Status, PricingModel, BidStrategy, Pacing)
            .subscribe(() => {
                this.router.navigate(['/indexLineItem']);
            });
    }

    ngOnInit(): void {
    }
}