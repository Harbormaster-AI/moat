import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RateService } from '../../../services/Rate.service';
import { Rate } from '../../../models/Rate';
import { SubBaseComponent } from '../../Rate/sub.base.component';

@Component({
    selector: 'app-create-rate',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRateComponent extends SubBaseComponent implements OnInit {

    title = 'Add Rate';

    rateForm: FormGroup;
    rate: Rate;

    constructor( http: HttpClient,
        private rateService: RateService,
        private fb: FormBuilder,
        private router: Router
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

    
    addRate(unitPrice, RateCard, AdSlot, AdFormat, PricingModel): void {
        this.rateService
        .addRate(unitPrice, RateCard, AdSlot, AdFormat, PricingModel)
            .subscribe(() => {
                this.router.navigate(['/indexRate']);
            });
    }

    ngOnInit(): void {
    }
}