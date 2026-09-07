import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RateCardService } from '../../../services/RateCard.service';
import { SubBaseComponent } from '../../RateCard/sub.base.component';


@Component({
    selector: 'app-edit-rateCard',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRateCardComponent extends SubBaseComponent implements OnInit {

    title = 'Edit RateCard';

    rateCardForm: FormGroup;
    rateCard: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RateCardService,
        private fb: FormBuilder
) {
        super(http);
        this.rateCardForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      effectiveDate: ['', Validators.required],
      currency: ['', Validators.required],
      Publisher: ['', ],
      Rates: ['', ]
        });
    }

    
    updateRateCard(name, effectiveDate, currency, Publisher, Rates): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRateCard(name, effectiveDate, currency, Publisher, Rates, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRateCard']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRateCard(params['id']).subscribe(res => {
                this.rateCard = res;
            });
        });
    }
}