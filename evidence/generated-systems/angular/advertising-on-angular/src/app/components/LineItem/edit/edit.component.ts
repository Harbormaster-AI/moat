import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LineItemService } from '../../../services/LineItem.service';
import { SubBaseComponent } from '../../LineItem/sub.base.component';


@Component({
    selector: 'app-edit-lineItem',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLineItemComponent extends SubBaseComponent implements OnInit {

    title = 'Edit LineItem';

    lineItemForm: FormGroup;
    lineItem: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LineItemService,
        private fb: FormBuilder
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

    
    updateLineItem(name, bidAmount, dailyBudget, frequencyCap, Campaign, Placements, TargetingProfile, Deal, Creatives, PerformanceMetrics, Status, PricingModel, BidStrategy, Pacing): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLineItem(name, bidAmount, dailyBudget, frequencyCap, Campaign, Placements, TargetingProfile, Deal, Creatives, PerformanceMetrics, Status, PricingModel, BidStrategy, Pacing, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLineItem']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLineItem(params['id']).subscribe(res => {
                this.lineItem = res;
            });
        });
    }
}