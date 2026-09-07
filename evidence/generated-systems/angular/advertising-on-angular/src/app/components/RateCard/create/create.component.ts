import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RateCardService } from '../../../services/RateCard.service';
import { RateCard } from '../../../models/RateCard';
import { SubBaseComponent } from '../../RateCard/sub.base.component';

@Component({
    selector: 'app-create-rateCard',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRateCardComponent extends SubBaseComponent implements OnInit {

    title = 'Add RateCard';

    rateCardForm: FormGroup;
    rateCard: RateCard;

    constructor( http: HttpClient,
        private rateCardService: RateCardService,
        private fb: FormBuilder,
        private router: Router
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

    
    addRateCard(name, effectiveDate, currency, Publisher, Rates): void {
        this.rateCardService
        .addRateCard(name, effectiveDate, currency, Publisher, Rates)
            .subscribe(() => {
                this.router.navigate(['/indexRateCard']);
            });
    }

    ngOnInit(): void {
    }
}