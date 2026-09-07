import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RateService } from '../../../services/Rate.service';
import { SubBaseComponent } from '../../Rate/sub.base.component';


@Component({
    selector: 'app-edit-rate',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRateComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Rate';

    rateForm: FormGroup;
    rate: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RateService,
        private fb: FormBuilder
) {
        super(http);
        this.rateForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  unitPrice: ['', Validators.required],
      RateCard: ['', ],
      AdSlot: ['', ],
      AdFormat: ['', ],
      PricingModel: ['', ]
        });
    }

    
    updateRate(unitPrice, RateCard, AdSlot, AdFormat, PricingModel): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRate(unitPrice, RateCard, AdSlot, AdFormat, PricingModel, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRate']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRate(params['id']).subscribe(res => {
                this.rate = res;
            });
        });
    }
}